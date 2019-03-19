# This file is part of the Resultra project.
#
# This Source Code Form is subject to the terms of the Mozilla Public
# License, v. 2.0. If a copy of the MPL was not distributed with this
# file, You can obtain one at https://mozilla.org/MPL/2.0/.

DEPTH = ..
include $(DEPTH)/build/common.mk

.PHONY: resultraServer resultraLocalBackend

all: build export

build: resultraServer resultraLocalBackend

clean:
	$(RM) resultraServer resultraLocalBackend
	
resultraServer: 
	$(GOBUILD) resultraServer.go
	
resultraLocalBackend:
	$(GOBUILD) resultraLocalBackend.go

windows:
	GOOS=windows GOARCH=amd64 CC=/usr/bin/x86_64-w64-mingw32-gcc CGO_ENABLED=1 $(GOBUILD) -o resultraLocalBackend.exe resultraLocalBackend.go
	$(INSTALL_FILE) resultraLocalBackend.exe $(PKG_BIN_DIR)
	
	
export:
	$(INSTALL_FILE) resultraServer $(PKG_BIN_DIR)
	$(INSTALL_FILE) resultraLocalBackend $(PKG_BIN_DIR)
