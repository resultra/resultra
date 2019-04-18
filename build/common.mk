# This file is part of the Resultra project.
#
# This Source Code Form is subject to the terms of the Mozilla Public
# License, v. 2.0. If a copy of the MPL was not distributed with this
# file, You can obtain one at https://mozilla.org/MPL/2.0/.

# Default is a debug build. This can be set explicitely on the command line with DEBUG=0 or DEBUG=1
DEBUG ?= 1

RM = rm -f
RM_DIR = rm -Rf 
INSTALL_FILE = $(DEPTH)/build/installFile.sh
INSTALL_RENAMED_FILE = $(DEPTH)/build/installRenameFile.sh
PKG_BIN_DIR = $(DEPTH)/build/dest/bin
PKG_TEMPLATES_DIR = $(DEPTH)/build/dest/factoryTemplates
PKG_SCHEMA_DIR = $(DEPTH)/build/dest/schema

DOCKER_DIST_ROOT_DIR = $(DEPTH)/build/dest/docker
DOCKER_DIST_DIR = $(DOCKER_DIST_ROOT_DIR)/resultra
DOCKER_BIN_DIR = $(DOCKER_DIST_DIR)/bin
DOCKER_IMAGE_DIR = $(DOCKER_DIST_DIR)/dockerImage

GO = go
GOBUILD = $(GO) build 
GOTEST = $(GO) test .

# Tell Go to use modules, even when the repo is inside the GOPATH.
# This is needed until Go's modeule support is turned on in Go 1.13.
# See https://blog.golang.org/modules2019 for more info.
export GO111MODULE = on

# Default, no-op/empty rules for different build phases. Makefiles can override these
# to implement build rules for different phases.

realclean:

clean:

install:

prebuild:

build:

export:

package:

windows:

winpkg:

dockerdist:

dockerpkg:

test:

systest:

