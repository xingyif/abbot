# Copyright 2020 The arhat.dev Authors.
#
# Licensed under the Apache License, Version 2.0 (the "License");
# you may not use this file except in compliance with the License.
# You may obtain a copy of the License at
#
#     http://www.apache.org/licenses/LICENSE-2.0
#
# Unless required by applicable law or agreed to in writing, software
# distributed under the License is distributed on an "AS IS" BASIS,
# WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
# See the License for the specific language governing permissions and
# limitations under the License.

# IMAGE_REPOS is the comma separated list of image registries
IMAGE_REPOS ?= docker.io/arhatdev,ghcr.io/arhat-dev

export IMAGE_REPOS

DEFAULT_IMAGE_MANIFEST_TAG ?= latest

include scripts/lint.mk

GOMOD := GOPROXY=direct GOSUMDB=off go mod
.PHONY: vendor
vendor:
	${GOMOD} tidy
	${GOMOD} vendor

# testing
include scripts/test/unit.mk

# binary build
include scripts/build/abbot.mk

# image
include scripts/image/abbot.mk

image.build.linux.all: \
	image.build.abbot.linux.all

image.build.windows.all: \
	image.build.abbot.windows.all

image.push.linux.all: \
	image.push.abbot.linux.all

image.push.windows.all: \
	image.push.abbot.windows.all

# manifest gen
include scripts/gen/manifest.mk

# packaging
include scripts/package/abbot.mk

# optional private scripts
-include private/scripts.mk
