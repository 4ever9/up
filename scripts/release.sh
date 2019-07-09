#!/bin/sh

ARCHES="darwin linux freebsd windows"

echo "Generating ${1} release binaries..."
for arch in ${ARCHES}; do
    CGO_ENABLED=0 GOOS=${arch} GOARCH=amd64 go build -ldflags "${3}" -o bin/${1}-${arch} ${4}
    tar -C bin/ -czvf bin/${1}_v${2}_${arch}.tar.gz ${1}-${arch}
done
