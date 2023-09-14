# doc https://github.com/tunnelvisionlabs/antlr4/blob/master/doc/tool-options.md
antlr4 ./lib/TypeScriptLexer.g4 -Dlanguage=Go -package typescript_ast
antlr4 ./lib/TypeScriptParser.g4 -Dlanguage=Go -package typescript_ast -visitor
# java -jar C:\Users\TheoSun\Documents\bin\antlr-4.13.1-complete.jar .\lib\TypeScriptLexer.g4 -Dlanguage=Go -package typescript_ast
# java -jar C:\Users\TheoSun\Documents\bin\antlr-4.13.1-complete.jar .\lib\TypeScriptParser.g4 -Dlanguage=Go -package typescript_ast -visitor