// Copyright (c) The Thanos Authors.
// Licensed under the Apache License 2.0.

//go:build !ignoretests

package query

import (
	"context"
	"fmt"
	"os"
	"path/filepath"
	"testing"
	"time"

	"github.com/efficientgo/core/testutil"
	"github.com/go-kit/log"
	"github.com/oodle-ai/thanos/pkg/component"
	"github.com/oodle-ai/thanos/pkg/store"
	"github.com/oodle-ai/thanos/pkg/store/storepb"
	storetestutil "github.com/oodle-ai/thanos/pkg/store/storepb/testutil"
	"github.com/oodle-ai/thanos/pkg/testutil/custom"
	"github.com/prometheus/prometheus/storage"
)

func TestMain(m *testing.M) {
	custom.TolerantVerifyLeakMain(m)
}

func TestQuerier_Proxy(t *testing.T) {
	files, err := filepath.Glob("testdata/promql/**/*.test")
	testutil.Ok(t, err)
	testutil.Equals(t, 10, len(files), "%v", files)

	logger := log.NewLogfmtLogger(os.Stderr)
	t.Run("proxy", func(t *testing.T) {
		var clients []store.Client
		q := NewQueryableCreator(
			logger,
			nil,
			store.NewProxyStore(logger, nil, func() []store.Client { return clients },
				component.Debug, nil, 5*time.Minute, store.EagerRetrieval),
			1000000,
			5*time.Minute,
		)

		createQueryableFn := func(stores []*testStore) storage.Queryable {
			clients = clients[:0]
			for i, st := range stores {
				m, err := storepb.PromMatchersToMatchers(st.matchers...)
				testutil.Ok(t, err)

				// TODO(bwplotka): Parse external labels.
				clients = append(clients, &storetestutil.TestClient{
					Name:        fmt.Sprintf("store number %v", i),
					StoreClient: storepb.ServerAsClient(selectedStore(store.NewTSDBStore(logger, st.storage.DB, component.Debug, nil), m, st.mint, st.maxt)),
					MinTime:     st.mint,
					MaxTime:     st.maxt,
				})
			}
			return q(true,
				nil,
				nil,
				0,
				false,
				false,
				nil,
				NoopSeriesStatsReporter,
			)
		}

		for _, fn := range files {
			t.Run(fn, func(t *testing.T) {
				te, err := newTestFromFile(t, fn)
				testutil.Ok(t, err)
				testutil.Ok(t, te.run(createQueryableFn))
				te.close()
			})
		}
	})
}

// selectStore allows wrapping another storeEndpoints with additional time and matcher selection.
type selectStore struct {
	matchers []storepb.LabelMatcher

	storepb.StoreServer
	mint, maxt int64
}

// selectedStore wraps given store with selectStore.
func selectedStore(wrapped storepb.StoreServer, matchers []storepb.LabelMatcher, mint, maxt int64) *selectStore {
	return &selectStore{
		StoreServer: wrapped,
		matchers:    matchers,
		mint:        mint,
		maxt:        maxt,
	}
}

func (s *selectStore) Info(ctx context.Context, r *storepb.InfoRequest) (*storepb.InfoResponse, error) {
	resp, err := s.StoreServer.Info(ctx, r)
	if err != nil {
		return nil, err
	}
	if resp.MinTime < s.mint {
		resp.MinTime = s.mint
	}
	if resp.MaxTime > s.maxt {
		resp.MaxTime = s.maxt
	}
	// TODO(bwplotka): Match labelsets and expose only those?
	return resp, nil
}

func (s *selectStore) Series(r *storepb.SeriesRequest, srv storepb.Store_SeriesServer) error {
	if r.MinTime < s.mint {
		r.MinTime = s.mint
	}
	if r.MaxTime > s.maxt {
		r.MaxTime = s.maxt
	}

	matchers := make([]storepb.LabelMatcher, 0, len(r.Matchers))
	matchers = append(matchers, r.Matchers...)

	req := *r
	req.Matchers = matchers
	return s.StoreServer.Series(&req, srv)
}
