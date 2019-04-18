# This file is part of the Resultra project.
#
# This Source Code Form is subject to the terms of the Mozilla Public
# License, v. 2.0. If a copy of the MPL was not distributed with this
# file, You can obtain one at https://mozilla.org/MPL/2.0/.
DEPTH = .
include $(DEPTH)/build/common.mk
	
realclean:
	$(RM_DIR) vendor
	

# Only use Go's dep tool when inside GOPATH, but use
# modules otherwise.
# 
# The Makefile pattern match is constructed per the 
# following: https://stackoverflow.com/questions/2741708/makefile-contains-string
INGOPATH := $(shell $(DEPTH)/build/testRepoInGoPath.py)	
	
install:
ifneq (,$(findstring TRUE,$(VARIABLE)))
	go get -u github.com/golang/dep/cmd/dep
endif

prebuild:
ifneq (,$(findstring TRUE,$(VARIABLE)))
	dep ensure
endif
	
all: prebuild
