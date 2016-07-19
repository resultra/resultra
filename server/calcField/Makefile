
DEPTH = ../..
include $(DEPTH)/build/common.mk

all: prebuild

	
clean:
	$(RM) formulaParser_gen.go y.output

test:
	$(GOTEST)
	
prebuild: formulaParser_gen.go

formulaParser_gen.go: formulaParser.y
	go tool yacc -o formulaParser_gen.go -p formula formulaParser.y
