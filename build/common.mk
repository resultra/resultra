
# Default is a debug build. This can be set explicitely on the command line with DEBUG=0 or DEBUG=1
DEBUG ?= 1

RM = rm -f
RM_DIR = rm -Rf 
INSTALL_FILE = $(DEPTH)/build/installFile.sh
INSTALL_RENAMED_FILE = $(DEPTH)/build/installRenameFile.sh
PKG_BIN_DIR = $(DEPTH)/build/dest/bin
PKG_TEMPLATES_DIR = $(DEPTH)/build/dest/factoryTemplates
PKG_SCHEMA_DIR = $(DEPTH)/build/dest/schema

GO = go
GOBUILD = $(GO) build 
GOTEST = $(GO) test -v .

# Default, no-op/empty rules for different build phases. Makefiles can override these
# to implement build rules for different phases.

realclean:

clean:

install:

prebuild:

build:

export:

package:

windows:

winpkg:

