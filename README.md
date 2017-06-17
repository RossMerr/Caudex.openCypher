### Antlr does not yet generated the go code correctly

### Setup

[Install Antlr](https://github.com/antlr/antlr4/blob/master/doc/getting-started.md) first then running somthing like to generate the parser/lexer


* java -jar /usr/local/lib/antlr-4.7-complete.jar -Dlanguage=Go -package cypher Cypher.g4
* java -jar "C:\Program Files\antlr\antlr-4.7-complete.jar" -Dlanguage=Go -package cypher Cypher.g4