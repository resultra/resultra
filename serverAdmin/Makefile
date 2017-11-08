
DEPTH = ..
include $(DEPTH)/build/common.mk

.PHONY: addResultraAccount

all: build package

build: addResultraAccount

clean:
	$(RM) addResultraAccount
	
addResultraAccount: 
	$(GOBUILD) addResultraAccount.go
	
package:
	$(INSTALL_FILE) addResultraAccount $(PKG_BIN_DIR)