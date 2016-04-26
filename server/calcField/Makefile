
all: formulaParser_gen.go

test:
	goapp test -v

formulaParser_gen.go: formulaParser.y
	go tool yacc -o formulaParser_gen.go -p formula formulaParser.y
