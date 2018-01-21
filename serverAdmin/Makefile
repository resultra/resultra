
DEPTH = ..
include $(DEPTH)/build/common.mk

.PHONY: addResultraAccount

all: build export

build: addResultraAccount

clean:
	$(RM) addResultraAccount
	
addResultraAccount: 
	$(GOBUILD) addResultraAccount.go
	
export:
	$(INSTALL_FILE) addResultraAccount $(PKG_BIN_DIR)