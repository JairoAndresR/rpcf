## Products-definitions

The following documentation contains examples of products definitions that can
be created based on the supported products.

### Articles

```shell script
curl --request POST \
  --url http://localhost:8080/v1/product-definitions \
  --header 'Content-Type: application/json' \
  --data '{
	"name" : "articles",
	"definition" : 		"name: Artículos\npath: \"/html[1]/body[1]/table[14]/tbody[1]/tr\"\nparser:\n  lookahead-manual: true\nfields:\n  -\n    name: type\n    ex: (<strong>).*(:</strong>)\n  -\n    name: title\n    ex: (</strong>).*(<br/>)\n  -\n    name: published_country\n    ex: (<br/>)(\\s).*?(,)\n  -\n    name: maganize_name\n    ex: (,).*(ISSN:)\n  -\n    name: maganize_issn\n    ex: \"[0-9]+-[0-9]+X?\"\n  -\n    name: published_year\n    ex: \\d{4}( vol:)\n  -\n    name: vol\n    ex: (vol:).*( fasc:)\n  -\n    name: fasc\n    ex: (fasc:).*( págs:)\n  -\n    name: pags\n    ex: (págs:).*(, <strong>)\n  -\n    name: DOI\n    ex: (DOI:</strong>).*(\\s{2,} <br/>) #No elimina salto de linea ni etiqueta br\n  -\n    name: autores\n    ex: (Autores:).*, \n"
}'
```

### Books

```shell script
curl --request POST \
  --url http://localhost:8080/v1/product-definitions \
  --header 'Content-Type: application/json' \
  --data '{
	"name" : "books",
	"definition" : 			"name: Libros\npath: \"/html[1]/body[1]/table[15]/tbody[1]/tr\"\nparser:\n  lookahead-manual: true\nfields:\n  -\n    name: type\n    ex: (<strong>).*(</strong>)\n  -\n    name: name\n    ex: (</strong>).*(<br/>)\n\n  - name: isbn\n    ex: (ISBN:).*(vol)\n  -\n    name: published_country\n    ex: (<br/>\\s+).*(,[0-9]+, ISBN)\n  -\n    name: published_year\n    ex: \\d{4}\n  -\n    name: vol\n    ex: (vol:).*( págs:)\n  -\n    name: pags\n    ex: (págs:).*(, Ed.)\n  -\n    name: editorial\n    ex: (Ed.)(.*)\\s+(<br/>)\n  -\n    name: autores\n    ex: (Autores:)(.*),\n"
}'
```
