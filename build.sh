#!/usr/bin/env bash

function build() {
    target_os=$1
    target_arch=$2
    GOOS=${target_os} GOARCH=${target_arch} go build -o bin/ipinfo-sender-${target_os}-${target_arch} github.com/ismdeep/ipinfo-sender/sender
    echo BUILD PASS ${target_os} ${target_arch}
}


build linux amd64
build linux 386
build linux arm
build linux arm64
build darwin amd64
build windows amd64
build windows 386
build windows arm
build windows amd64
build freebsd amd64
build freebsd 386
build freebsd arm
build freebsd arm64