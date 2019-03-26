# This file is part of the Resultra project.
#
# This Source Code Form is subject to the terms of the Mozilla Public
# License, v. 2.0. If a copy of the MPL was not distributed with this
# file, You can obtain one at https://mozilla.org/MPL/2.0/.
DEPTH = ..
include $(DEPTH)/build/common.mk

.PHONY: addResultraAccount exportTrackerTemplate resultraSetup

all: build

build: addResultraAccount exportTrackerTemplate resultraSetup

clean:
	$(RM) addResultraAccount exportTrackerTemplate
	
addResultraAccount: 
	$(GOBUILD) addResultraAccount.go
	$(INSTALL_FILE) addResultraAccount $(PKG_BIN_DIR)
	
exportTrackerTemplate:
	$(GOBUILD) exportTrackerTemplate.go
	$(INSTALL_FILE) exportTrackerTemplate $(PKG_BIN_DIR)


resultraSetup:
	$(GOBUILD) resultraSetup.go
	$(INSTALL_FILE) resultraSetup $(PKG_BIN_DIR)
