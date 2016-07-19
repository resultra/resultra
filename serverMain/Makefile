
DEPTH = ..
include $(DEPTH)/build/common.mk

.PHONY: datasheetServer

all: build package

build: datasheetServer

clean:
	$(RM) datasheetServer
	
datasheetServer: 
	$(GOBUILD) datasheetServer.go
	
package:
	$(INSTALL_FILE) datasheetServer $(PKG_BIN_DIR)
