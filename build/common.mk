
# Default is a debug build. This can be set explicitely on the command line with DEBUG=0 or DEBUG=1
DEBUG ?= 1

RM = rm -f
INSTALL_FILE = $(DEPTH)/build/installFile.sh
INSTALL_RENAMED_FILE = $(DEPTH)/build/installRenameFile.sh
PKG_BIN_DIR = $(DEPTH)/build/dest/bin
PKG_SCHEMA_DIR = $(DEPTH)/build/dest/schema

GO = go
GOBUILD = $(GO) build 
GOTEST = $(GO) test -v .