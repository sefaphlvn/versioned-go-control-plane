#!/bin/bash -e

# Script to generate Envoy proto files
set -o pipefail

SRCS=(envoy contrib)
IMPORT_BASE="github.com/sefaphlvn/versioned-go-control-plane"
ENVOY_SRC_DIR="${ENVOY_SRC_DIR:-}"
ENVOY_VERSION="${ENVOY_VERSION:-}"
RATELIMIT_VERSION="${RATELIMIT_VERSION:-v0.1.0}"  # Default if not provided

if [[ -z "$ENVOY_SRC_DIR" ]]; then
    echo "ENVOY_SRC_DIR not set, it should point to a cloned Envoy repo" >&2
    exit 1
elif [[ ! -d "$ENVOY_SRC_DIR" ]]; then
    echo "ENVOY_SRC_DIR ($ENVOY_SRC_DIR) not found, did you clone it?" >&2
    exit 1
fi

# Create required directories
mkdir -p envoy contrib

build_protos () {
    echo "Building go protos ..."
    cd "${ENVOY_SRC_DIR}" || exit 1
    ./ci/do_ci.sh api.go
    cd - || exit 1
}

sync_protos () {
    local src envoy_src
    echo "Syncing go protos ..."
    
    # First clean existing directories
    for src in "${SRCS[@]}"; do
        rm -rf "${src:?}/*"
        mkdir -p "${src}"
    done
    
    # Copy and organize proto files
    for src in "${SRCS[@]}"; do
        envoy_src="${ENVOY_SRC_DIR}/build_go/${src}"
        echo "Copying ${envoy_src}/* -> ${src}"
        cp -r "${envoy_src}"/* "${src}/"
        
        # Update import paths in all .go files
        find "${src}" -type f -name "*.go" -exec sed -i "s|github.com/envoyproxy/go-control-plane|${IMPORT_BASE}|g" {} \;
    done
    
    # Create and update go.mod files
    printf "module github.com/sefaphlvn/versioned-go-control-plane\n\ngo 1.22\n\nreplace (\n  github.com/sefaphlvn/versioned-go-control-plane/envoy => ./envoy\n  github.com/sefaphlvn/versioned-go-control-plane/ratelimit => ./ratelimit\n)\n\nrequire (\n  github.com/sefaphlvn/versioned-go-control-plane/envoy v%s\n  github.com/sefaphlvn/versioned-go-control-plane/ratelimit %s\n  github.com/google/go-cmp v0.6.0\n  github.com/stretchr/testify v1.10.0\n  go.uber.org/goleak v1.3.0\n  google.golang.org/genproto/googleapis/rpc v0.0.0-20241015192408-796eee8c2d53\n  google.golang.org/grpc v1.69.4\n  google.golang.org/protobuf v1.36.3\n)\n" "${ENVOY_VERSION}" "${RATELIMIT_VERSION}" > go.mod

    printf "module github.com/sefaphlvn/versioned-go-control-plane/envoy\n\ngo 1.22\n\nrequire (\n  google.golang.org/genproto/googleapis/rpc v0.0.0-20241015192408-796eee8c2d53\n  google.golang.org/protobuf v1.36.3\n)\n" > envoy/go.mod

    printf "module github.com/sefaphlvn/versioned-go-control-plane/ratelimit\n\ngo 1.22\n\nrequire (\n  github.com/sefaphlvn/versioned-go-control-plane/envoy v%s\n  google.golang.org/protobuf v1.36.3\n)\n\nreplace github.com/sefaphlvn/versioned-go-control-plane/envoy => ../envoy\n" "${ENVOY_VERSION}" > ratelimit/go.mod
}

build_protos
sync_protos

echo "Proto generation completed" 