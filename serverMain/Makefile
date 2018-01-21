
DEPTH = ..
include $(DEPTH)/build/common.mk

.PHONY: datasheetServer

all: build export

build: datasheetServer

clean:
	$(RM) datasheetServer
	
datasheetServer: 
	$(GOBUILD) datasheetServer.go
	
export:
	$(INSTALL_FILE) datasheetServer $(PKG_BIN_DIR)
