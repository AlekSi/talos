#!/bin/bash

# Sync changes with configuring-pull-through-cache.md.
# com.github.AlekSi.docker-ps.group label is for https://github.com/AlekSi/docker-ps-bitbar

set -e

docker run -d -p 5000:5000 \
    -e REGISTRY_PROXY_REMOTEURL=https://registry-1.docker.io \
    --restart always \
    --label com.github.AlekSi.docker-ps.group=talos-registry \
    --name registry-docker.io registry:2

docker run -d -p 5001:5000 \
    -e REGISTRY_PROXY_REMOTEURL=https://k8s.gcr.io \
    --restart always \
    --label com.github.AlekSi.docker-ps.group=talos-registry \
    --name registry-k8s.gcr.io registry:2

docker run -d -p 5002:5000 \
    -e REGISTRY_PROXY_REMOTEURL=https://quay.io \
    --restart always \
    --label com.github.AlekSi.docker-ps.group=talos-registry \
    --name registry-quay.io registry:2.5

docker run -d -p 5003:5000 \
    -e REGISTRY_PROXY_REMOTEURL=https://gcr.io \
    --restart always \
    --label com.github.AlekSi.docker-ps.group=talos-registry \
    --name registry-gcr.io registry:2

docker run -d -p 5004:5000 \
    -e REGISTRY_PROXY_REMOTEURL=https://ghcr.io \
    --restart always \
    --label com.github.AlekSi.docker-ps.group=talos-registry \
    --name registry-ghcr.io registry:2
