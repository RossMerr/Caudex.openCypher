Cypher is a graph query language that allows for expressive and efficient querying of graph data. Cypher is intuitive, powerful and easy to learn.

It lets you write graph queries by describing patterns in your data. Because graphs already describe your domain, Cypher lets you focus on your domain instead of getting lost in the mechanics of data access.

Designed to be a human readable query language, Cypher is suitable for developers and operations professionals alike. The guiding principle behind Cypher is to make simple things easy and complex things possible. Cypher combines English prose and intuitive iconography, making queries more self-explanatory.

Cypher is declarative, which means it lets users express what data to retrieve, letting the engine underneath take care of seamlessly retrieving it. In contrast, imperative or scripting languages, including early-generation graph languages, require query writers to have a deep technical understanding of physical implementation details. Cypher eliminates this burden.

The expressive querying of Cypher is inspired by a number of different approaches and established practices. Most of the keywords, such as WHERE and ORDER BY , are inspired by SQL, while pattern matching borrows from SPARQL. In addition, some of the collection semantics have been borrowed from languages such as Haskell and Python.

# Setup

[Install Antlr](https://github.com/antlr/antlr4/blob/master/doc/getting-started.md) first then running somthing like to generate the parser/lexer


* java -jar /usr/local/lib/antlr-4.7-complete.jar -Dlanguage=Go -package cypher Cypher.g4
* java -jar "C:\Program Files\antlr\antlr-4.7-complete.jar" -Dlanguage=Go -package cypher Cypher.g4
