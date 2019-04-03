# This file is part of the Resultra project.
#
# This Source Code Form is subject to the terms of the Mozilla Public
# License, v. 2.0. If a copy of the MPL was not distributed with this
# file, You can obtain one at https://mozilla.org/MPL/2.0/.

DEPTH = ../..
include $(DEPTH)/build/common.mk

all: prebuild

clean:
	$(RM) formulaParser_gen.go y.output

test:
	$(GOTEST)
	
# The calcField package needs Go's yacc to build the parser. This parser is obtained using the following
installyacc:
	go get -u golang.org/x/tools/cmd/goyacc

install: installyacc
	
prebuild: formulaParser_gen.go

formulaParser_gen.go: formulaParser.y
	goyacc -o formulaParser_gen.go -p formula formulaParser.y
