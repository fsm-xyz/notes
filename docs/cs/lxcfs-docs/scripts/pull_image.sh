#!/bin/bash
# 下载所有 Distroless 镜像版本

IMAGES=(
    "static-debian13"
    "base-debian13" 
    "base-nossl-debian13"
    "cc-debian13"
)

TAGS=("latest" "nonroot" "debug" "debug-nonroot")

for img in "${IMAGES[@]}"; do
    for tag in "${TAGS[@]}"; do
        echo "正在下载: gcr.io/distroless/$img:$tag"
        docker pull "gcr.io/distroless/$img:$tag" 2>&1 | grep -E "(Downloaded|Pulling|Already exists|Error)" || true
    done
done

echo "完成! 已下载所有镜像。"
docker images | grep gcr.io/distroless
