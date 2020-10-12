# RPCF
Research product collector framework

## Development

The following are the instructions to run it in local environment

1. You should up the database with docker compose using the next command:
```cassandraql
docker-compose up
```
2. Setup the environment variables that are set in the `.env` file
3. Running the application without IDE in the root directory using the following command: 
```cassandraql
go run *.go
``` 
## Testing

```
go test ./...
```
