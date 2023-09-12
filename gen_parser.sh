# doc https://github.com/tunnelvisionlabs/antlr4/blob/master/doc/tool-options.md
antlr4 ./lib/TypeScriptLexer.g4 -Dlanguage=Go -package lib
antlr4 ./lib/TypeScriptParser.g4 -Dlanguage=Go -package lib -visitor
# java -jar C:\Users\TheoSun\Documents\bin\antlr-4.13.1-complete.jar .\lib\TypeScriptLexer.g4 -Dlanguage=Go -package lib
# java -jar C:\Users\TheoSun\Documents\bin\antlr-4.13.1-complete.jar .\lib\TypeScriptParser.g4 -Dlanguage=Go -package lib -visitor