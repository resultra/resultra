
DEPTH = ..
include $(DEPTH)/build/common.mk

.PHONY: datasheetServer resultraLocalBackend

all: build export

build: datasheetServer resultraLocalBackend

clean:
	$(RM) datasheetServer resultraLocalBackend
	
datasheetServer: 
	$(GOBUILD) datasheetServer.go
	
resultraLocalBackend:
	$(GOBUILD) resultraLocalBackend.go

windows:
	GOOS=windows GOARCH=386 CC=/usr/bin/i686-w64-mingw32-gcc CGO_ENABLED=1 $(GOBUILD) -o resultraLocalBackend.exe resultraLocalBackend.go
	$(INSTALL_FILE) resultraLocalBackend.exe $(PKG_BIN_DIR)
	
	
export:
	$(INSTALL_FILE) datasheetServer $(PKG_BIN_DIR)
	$(INSTALL_FILE) resultraLocalBackend $(PKG_BIN_DIR)
