# Collection research product

This document explains how to scrap the different products step by step

## Design

## Process

The following process explains how setup the information to start to scraping the research products from GrupLACs using the API developed here
 
1- **Setup the GrupLAC definitions**, the idea in this point is to store the  GrupLAC's URLs where the application should find the content of it. 
The following is the request that we should do to create new GrupLAC definition. We need the following

Endpoint:
```
http://host/v1/gruplac-definitions
```

Parameters:

| Name   |      Description      |
|----------|:-------------|
| name |  GrupLAC's name | 
| URL |    It's the url to find the GrupLAC's content    |  

Request Example:

```curl
curl --request POST \
  --url http://localhost:8080/v1/gruplac-definitions \
  --header 'Content-Type: application/json' \
  --data '{
	"name": "Gestión en Cultura y Educación Ambiental",
	"url": "https://scienti.minciencias.gov.co/gruplac/jsp/visualiza/visualizagr.jsp?nro=00000000002584"
}'
```

2- **Setup the products definitions**, in this step should confiured the definitions of the products that should be parsed in the next steps. 
The following is the request that should be used to creating a new product definition:

Endpoint
```curl
```
Parameters

| Name   |      Description      |
|----------|:-------------|
| name |  Product definition name | 
| definition | The yaml definition in raw string |  

Request example:
```curl
curl --request POST \
  --url http://localhost:8080/v1/product-definitions \
  --header 'Content-Type: application/json' \
  --data '{
	"name" : "articles",
	"definition" : 		"name: Artículos\npath: \"/html[1]/body[1]/table[14]/tbody[1]/tr\"\nparser:\n  lookahead-manual: true\nfields:\n  -\n    name: type\n    ex: (<strong>).*(:</strong>)\n  -\n    name: title\n    ex: (</strong>).*(<br/>)\n  -\n    name: published_country\n    ex: (<br/>)(\\s).*?(,)\n  -\n    name: maganize_name\n    ex: (,).*(ISSN:)\n  -\n    name: maganize_issn\n    ex: \"[0-9]+-[0-9]+X?\"\n  -\n    name: published_year\n    ex: \\d{4}( vol:)\n  -\n    name: vol\n    ex: (vol:).*( fasc:)\n  -\n    name: fasc\n    ex: (fasc:).*( págs:)\n  -\n    name: pags\n    ex: (págs:).*(, <strong>)\n  -\n    name: DOI\n    ex: (DOI:</strong>).*(\\s{2,} <br/>) #No elimina salto de linea ni etiqueta br\n  -\n    name: autores\n    ex: (Autores:).*, \n"
}'
```

3- **Running the worker**, in charge to process the data in the queue.
It's a process manually right now. It should be executed with:
```go
go run /path/rpfc/app/cmd/products.scrapers/*.go
```
   
**Note**: should be the backend up at the same time in which it would be executed.
 
 
4- **Retrieving the contents**, in this step the idea is get all the GrupLAC definitions stored in the application
and take the URL and scrap the content on it. After that the content should be sent to queue configured in
 environment variable`COLLECTOR_QUEUE_NAME`. The request should be the following to active the process
 
```
curl --request POST \
--url http://localhost:8080/v1/gruplacs/scraping
```  
The payload data saved in the queue is represented by the following attributes:

| Name   |      Description      |
|----------|:-------------|
| code |  The GrupLAC code extracted from the URL | 
| content | The full html scraped   | 

The following is JSON example 
```json
{
"code": "00000000002584",
"content" : "\n\n\u003c!DOCTYPE HTML PUBLIC..."
}
```