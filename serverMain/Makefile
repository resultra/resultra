
DEPTH = ..
include $(DEPTH)/build/common.mk

.PHONY: datasheetServer resultraLocalBackend

all: build export

build: datasheetServer resultraLocalBackend

clean:
	$(RM) datasheetServer
	
datasheetServer: 
	$(GOBUILD) datasheetServer.go
	
resultraLocalBackend:
	$(GOBUILD) resultraLocalBackend.go
	
	
export:
	$(INSTALL_FILE) datasheetServer $(PKG_BIN_DIR)
	$(INSTALL_FILE) resultraLocalBackend $(PKG_BIN_DIR)
