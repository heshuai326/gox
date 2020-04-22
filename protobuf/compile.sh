#!/usr/bin/env bash

DIR=$(dirname "$0")

protoc  -I ${DIR}/  ${DIR}/base.proto --go_out=plugins=grpc:${DIR}/
rm -r ${DIR}/base/*.pb.go
mv ${DIR}/github.com/gopub/gox/protobuf/base ${DIR}/
rm -r ${DIR}/github.com

