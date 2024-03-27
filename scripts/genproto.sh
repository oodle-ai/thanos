#!/usr/bin/env bash
#
# Generate all protobuf bindings.
# Run from repository root.
set -e
set -u

PROTOC_VERSION=${PROTOC_VERSION:-3.20.1}
PROTOC_BIN=${PROTOC_BIN:-protoc}
GOIMPORTS_BIN=${GOIMPORTS_BIN:-goimports}

if ! [[ "scripts/genproto.sh" =~ $0 ]]; then
  echo "must be run from repository root"
  exit 255
fi

if ! [[ $(${PROTOC_BIN} --version) == *"${PROTOC_VERSION}"* ]]; then
  echo "could not find protoc ${PROTOC_VERSION}, is it installed + in PATH?"
  exit 255
fi

mkdir -p /tmp/protobin/
PATH=${PATH}:/tmp/protobin

GOOGLE_PROTO="$(go list  -f '{{.Dir}}' -m github.com/golang/protobuf@v1.3.2)"
VT_PROTO="$(go list  -f '{{.Dir}}' -m github.com/planetscale/vtprotobuf\@v0.6.0)"

DIRS="store/storepb/ store/storepb/prompb/ store/labelpb rules/rulespb targets/targetspb store/hintspb queryfrontend metadata/metadatapb exemplars/exemplarspb info/infopb api/query/querypb"
echo "generating code"
pushd "pkg"
for dir in ${DIRS}; do
  ${PROTOC_BIN} \
    --plugin protoc-gen-go="${PROTOC_GEN_GO_BIN}" \
    --go_out=. \
    --go-vtproto_out=. --plugin protoc-gen-go-vtproto="${PROTOC_GEN_GO_VTPROTO_BIN}" \
    -I=. \
    -I="${GOOGLE_PROTO}" \
    -I="${VT_PROTO}"/include \
    -I="/usr/local/include" \
    -I="/opt/homebrew/include" \
    ${dir}/*.proto

  pushd ${dir}
  ${GOIMPORTS_BIN} -w *.pb.go
  popd
done
popd

cp -rf pkg/github.com/oodle-ai/thanos/* ./
rm -rf pkg/github.com

## Generate vendored Cortex protobufs.
#CORTEX_DIRS="cortex/querier/queryrange/"
#pushd "internal"
#for dir in ${CORTEX_DIRS}; do
#  ${PROTOC_BIN} --gogofast_out=Mgoogle/protobuf/any.proto=github.com/gogo/protobuf/types,plugins=grpc:. \
#    -I=../pkg \
#    -I="${GOGOPROTO_PATH}" \
#    -I=. \
#    ${dir}/*.proto
#
#  pushd ${dir}
#  sed -i.bak -E 's/import _ \"gogoproto\"//g' *.pb.go
#  sed -i.bak -E 's/_ \"google\/protobuf\"//g' *.pb.go
#  sed -i.bak -E 's/\"cortex\/cortexpb\"/\"github.com\/thanos-io\/thanos\/internal\/cortex\/cortexpb\"/g' *.pb.go
#  rm -f *.bak
#  ${GOIMPORTS_BIN} -w *.pb.go
#  popd
#done
#popd
