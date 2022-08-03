# RPCF
Research product collector framework


## Requirements
Docker

Docker-compose

MySqlWorkbench (To load authors and products definitions)

Run frontend (more info: https://github.com/JairoAndresR/rpcf-web)

## Development

The following are the instructions to run it in local environment

1. You should up the database with docker compose using the next command:
```cassandraql
docker-compose -f docker-compose-local.yml up
```
2. Set author definition using MySqlWorkbench for author_definitions table with following SQL sentence
```cassandraql
INSERT INTO rpcf.author_definitions (id, definition) VALUES ('1', 'name: Libros\npath: \"/html/body/table[5]/tbody[1]/tr\"\nparser:\n lookahead-manual: true\n skip-results: true\n skip-number-results: 2\nfields:\n -\n name: names\n ex: (>).*(</a>)\n\n - name: link\n ex: (href=\").*(\")');
```
3. Create new admin user using this url and filling the form 
`localhost:4200/accounts/new`
4. Change user rol using MySqlWorkbench with following SQL sentence
```UPDATE rpcf.accounts SET role = 'admin' WHERE (email = '{your email}');```
5. Login in admin module using this url `localhost:4200/admin`
6. Create new gruplac definition using admin module and set the group name and Scienti url
![image](https://user-images.githubusercontent.com/25800176/178641949-2b5378dd-40ca-4374-b769-3b67c2d21ac9.png)

7. Create new product definition using admin module and set product name and definition (definitions examples: https://github.com/JairoAndresR/collector/tree/master/examples)
Product names: (articles, books, companies, doctoral_theses)

![image](https://user-images.githubusercontent.com/25800176/178642008-5b806c07-5f24-40cd-8bb0-0636d90daacb.png)

![image](https://user-images.githubusercontent.com/25800176/178642026-441868fb-6b9f-4ff9-8ab1-fcce4e20d8bd.png)

9. Go to admin module home and execute authors scrapping. Wait two minutes to execute products scrapper
![image](https://user-images.githubusercontent.com/25800176/178642469-4cb67ccd-9f5c-447e-a888-c77b71284d79.png)

10. Now you can go to reports module in `http://localhost:4200`
