

RM = rm -f
INSTALL_FILE = $(DEPTH)/build/installFile.sh
INSTALL_RENAMED_FILE = $(DEPTH)/build/installRenameFile.sh
PKG_BIN_DIR = $(DEPTH)/build/dest/bin
PKG_SCHEMA_DIR = $(DEPTH)/build/dest/schema

GO = go
GOBUILD = $(GO) build 
GOTEST = $(GO) test -v .