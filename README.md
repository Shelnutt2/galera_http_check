# Galera HTTP Check

This is a small script to provide the status of a galera cluster over http.

With a (mysql) galera node if it is in a donor or joiner state, the port will
be up but the server is not accepting queries.

This provides an http endpoint (/check-galera-status) which will return 500
if the server is down or galera is in a non-serving state. It will return
200 otherwise.

The body returned provides details.

## Running

This package can be build and run from source or in a docker container

### Docker

This is the recommended method for running. Provided is a docker compose file
for use. This must be updated with the proper variables for your setup.

Once the docker-compose has been updated it can be run as such:

```bash
docker-compse up
```

### Source

The second method of running this code is to build it from source

```bash
go get -u Shelnutt2/galera_http_check
```

To connect to mysql, the username and password are read from environmental
variables. You must set these before you run the application

```bash
export MYSQL_USERNAME=user
export MYSQL_PASSWORD=password
```

If $GOPATH/bin is in your PATH then you can invoke it directly.

```bash
galera_http_check --host localhost
```

Else run it from $GOPATH

```bash
$GOPATH/bin/galera_http_check --host localhost
```
