# This file is part of the Resultra project.
#
# This Source Code Form is subject to the terms of the Mozilla Public
# License, v. 2.0. If a copy of the MPL was not distributed with this
# file, You can obtain one at https://mozilla.org/MPL/2.0/.
DEPTH = ../..
include $(DEPTH)/build/common.mk


# There's a conflict between lower versions of gulp and node version 10
# See: https://teamtreehouse.com/community/assertion-args1isstring-failed-gulp
# This is why a new version of gulp needs to be installed locally to support
# the build.

install:
	npm install
	npm install gulp@3.9.1 --save-dev

clean:
	
realclean:
	$(RM_DIR) node_modules
	
all: prebuild
