
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
	
prebuild: installyacc formulaParser_gen.go

formulaParser_gen.go: formulaParser.y
	goyacc -o formulaParser_gen.go -p formula formulaParser.y
