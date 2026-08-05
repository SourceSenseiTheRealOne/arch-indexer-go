#!/usr/bin/env bash
set -euo pipefail

project_root="$(cd "$(dirname "${BASH_SOURCE[0]}")/.." && pwd)"
cd "$project_root"

shopt -s globstar nullglob
go_files=(**/*.go)

if ((${#go_files[@]} > 0)); then
  unformatted="$(gofmt -l "${go_files[@]}")"
  if [[ -n "$unformatted" ]]; then
    printf '%s\n%s\n' 'Go files require gofmt:' "$unformatted" >&2
    exit 1
  fi
fi

go test -count=1 ./...
go test -count=1 -race ./...
go vet ./...
mkdir -p bin
build_goos="${BUILD_GOOS:-linux}"
build_goarch="${BUILD_GOARCH:-amd64}"
binary="bin/indexerd-${build_goos}-${build_goarch}"
GOOS="$build_goos" GOARCH="$build_goarch" CGO_ENABLED=0 \
  go build -trimpath -o "$binary" ./cmd/indexerd
