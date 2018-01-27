The calcField package needs Go's yacc to build the parser. This parser is obtained using the following:

	$ go get -u golang.org/x/tools/cmd/goyacc

After which the following is available:
	$ goyacc ...