# Developer

## Prerequisites

* Go 1.15
  * [https://golang.org/](https://golang.org/)
* gvm
  * [https://github.com/moovweb/gvm](https://github.com/moovweb/gvm)
* Docker
  * [https://docs.docker.com/install/](https://docs.docker.com/install/)
* psql (Postgresql client)
  * [https://www.postgresql.org/download/macosx/](https://www.postgresql.org/download/macosx/)
  * [https://www.postgresql.org/download/linux/ubuntu/](https://www.postgresql.org/download/linux/ubuntu/)

## Services

Create new service

```bash
./script/generate service [name]
```

* Update `~/.gitignore` to ignore binaries
* Update `service/[service name]/.env.development` server port to an available port
* Add configuration to `~/server/nginx.conf` and `~/server/nginx.macos.conf`
* Add configuration to `~/server/script/load-test-data` for loading test data
* Add configuration and types to `~/server/schema`

## Database Migrations

```bash
RUN go get -u -d github.com/golang-migrate/migrate/cmd/migrate github.com/lib/pq github.com/hashicorp/go-multierror \
    && go build -tags 'postgres' -o $GOPATH/bin/migrate github.com/golang-migrate/migrate/cmd/migrate
```

Create new migration

```bash
./script/db-migrate-create [service] [description]
```

Migrate up

```bash
./script/db-migrate-up
```

Migrate down

```bash
./script/db-migrate-down
```

## Testing

Test all packages

```bash
./script/test
```

## Running

Start all services

```bash
./script/start
```

Stop all services

```bash
./script/stop
```

## Upgrading Go Version

* Install `gvm install go1.X.X && gvm use go1.X.X --default`
* Update all `./server/service/*/build/docker/Dockerfile` configurations
* Update `./server/script` scripts
* Update `./.gitlab-ci.yml`
