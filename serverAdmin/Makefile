DEPTH = ..
include $(DEPTH)/build/common.mk

.PHONY: addResultraAccount exportTrackerTemplate

all: build

build: addResultraAccount exportTrackerTemplate

clean:
	$(RM) addResultraAccount exportTrackerTemplate
	
addResultraAccount: 
	$(GOBUILD) addResultraAccount.go
	$(INSTALL_FILE) addResultraAccount $(PKG_BIN_DIR)
	
exportTrackerTemplate:
	$(GOBUILD) exportTrackerTemplate.go
	$(INSTALL_FILE) exportTrackerTemplate $(PKG_BIN_DIR)
