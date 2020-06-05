#!/bin/bash
# Copyright 2020 The Go Authors. All rights reserved.
# Use of this source code is governed by a BSD-style
# license that can be found in the LICENSE file.

# NOTE: The integration scripts deliberately do not check to
# make sure that the test protos have been regenerated.
# It is intentional that older versions of the .pb.go files
# are checked in to ensure that they continue to function.
#
# Versions used:
#	protoc:        v3.9.1
#	protoc-gen-go: v1.3.2

base_home='github.com/cargod-bj/b2c-proto-common/common'
mkdir -p ${base_home}
curl -o ${base_home}/common.proto https://raw.githubusercontent.com/cargod-bj/b2c-proto-common/master/common/common.proto

for X in $(find . -name "*.proto" | sed "s|^\./||"); do
	protoc -I$(pwd) --go_out=paths=source_relative:.  --micro_out=paths=source_relative:. $X
done

rm -rf 'github.com'
