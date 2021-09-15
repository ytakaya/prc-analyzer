all: build-parser

build-parser:
	cd parser && antlr -Dlanguage=Go TSqlLexer.g4
	cd parser && antlr -Dlanguage=Go -visitor TSqlParser.g4
