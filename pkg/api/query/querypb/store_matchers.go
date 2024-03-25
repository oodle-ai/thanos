// Copyright (c) The Thanos Authors.
// Licensed under the Apache License 2.0.

package querypb

import (
	"github.com/oodle-ai/thanos/pkg/store/storepb"
	"github.com/prometheus/prometheus/model/labels"
)

func StoreMatchersToLabelMatchers(matchers []*StoreMatchers) ([][]*labels.Matcher, error) {
	if len(matchers) == 0 {
		return nil, nil
	}

	labelMatchers := make([][]*labels.Matcher, len(matchers))
	for i, storeMatcher := range matchers {
		storeMatchers, err := storepb.MatchersToPromMatchers(storeMatcher.LabelMatchers...)
		if err != nil {
			return nil, err
		}
		labelMatchers[i] = storeMatchers
	}

	return labelMatchers, nil
}
