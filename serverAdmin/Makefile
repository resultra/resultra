DEPTH = ..
include $(DEPTH)/build/common.mk

.PHONY: addResultraAccount

all: build export

build: addResultraAccount exportTrackerTemplate

clean:
	$(RM) addResultraAccount exportTrackerTemplate
	
addResultraAccount: 
	$(GOBUILD) addResultraAccount.go
	
exportTrackerTemplate: exportTrackerTemplate.go
	$(GOBUILD) exportTrackerTemplate.go
	
	
export:
	$(INSTALL_FILE) addResultraAccount $(PKG_BIN_DIR)