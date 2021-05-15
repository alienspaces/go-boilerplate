# Go Boilerplate

Contains *mostly* everything you might need to build a set of independent REST API's in Go.

As this is a **boilerplate** with supporting scripts and configuration files it is highly opinionated. As such, I won't feel bad if you don't like it, and neither should you 👍

## Overview

- Layered architecture
  - Application
  - Model
  - Data
- Multiple independent services
- Configuration driven API handlers
  - JSON schema API validation
  - Auto generated API documentation
- Middleware
  - Authentication
  - Authorization
  - Validation
  - Database
- Database
  - Migrations
  - Data management tools
- Example services
  - Complete CRUD handlers
  - Complete unit tests
- Developer `./scripts` for the *most common* development tasks
- Gitlab CI integration
- Kubernetes deployment configuration

## Documentation

Best practice:

- All source directories are documented
- All source files are documented

Within each directory there will be a `README.md` document of a standard format.

```bash
# /directory/path

An explanation of why this directory exists and its contents.

## Subdirectories

| Directory | Description         |
| --------- | ------------------- |
| ./next    | Short description.. |

```

Where possible all source files will be documented according to the standard for that source files language.

Otherwise, as close to the top of the file as possible will be a description, along with all variables and functions documented when necessary.

## Prerequisites

- Go 1.16.4
  - [https://golang.org/](https://golang.org/)
- gvm
  - [https://github.com/moovweb/gvm](https://github.com/moovweb/gvm)
- Docker
  - [https://docs.docker.com/install/](https://docs.docker.com/install/)
- psql (Postgresql client)
  - [https://www.postgresql.org/download/macosx/](https://www.postgresql.org/download/macosx/)
  - [https://www.postgresql.org/download/linux/ubuntu/](https://www.postgresql.org/download/linux/ubuntu/)

## Running

Start all services:

```bash
./script/start
```

### Examples API Documentation

- [Player](http://localhost:8082/players/documentation)
- [Characters](http://localhost:8082/characters/documentation)

Stop all services:

```bash
./script/stop
```

## Testing

Test all packages:

```bash
./script/test
```

Test specific service:

```bash
./script/test-service ./service/example/player
```

## Generators

Generator scripts are available for creating new services based on the existing `./service/template` service.

Example, generate a new service:

```bash
./script/generate service [servicename]
```

After a new service has been created the following manual steps are required to complete its integration.

- Git
  - Update `~/.gitignore` to ignore binaries
- Environment
  - Update `service/[service name]/.env.development` server port to an available port
- JSON Schema
  - Add JSON schema specfications and Go types to `~/server/schema`
- Data
  - Add configuration to `~/server/script/load-seed-data`
  - Add configuration to `~/server/script/load-test-data`
- API Gateway
  - Add configuration to `~/server/krakend.conf`

## Database Migrations

Database migrations are managed with [golang-migrate](https://github.com/golang-migrate/migrate) and is installed automatically when running migrations or may be installed manually with the following script.

Install `golang-migrate`:

```bash
./script/db-migrate-install
```

Create new migration:

```bash
./script/db-migrate-create [service] [description]
```

Migrate up:

```bash
./script/db-migrate-up
```

Migrate down:

```bash
./script/db-migrate-down
```

## Upgrading Go Version

- Install `gvm install go1.X.X && gvm use go1.X.X --default`
- Update `APP_SERVER_GO_VERSION` in `.env.development` and `.env.ci`
- Update `./.gitlab-ci.yml` Docker images
- Update all `./server/service/**/build/docker/Dockerfile` configurations

## TODO

- Fulfill the documentation promise
- Add `./core/server/middleware` directory
- Add architecture diagrams
- Demonstrate interfaces/mocks for models and repositories and how tests might make use of them
- Add support for event publishing and subscribing
