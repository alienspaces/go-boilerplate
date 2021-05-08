# Holy Raging Mages - Server

Server application for the [Holy Raging Mages](https://go-mono-api-boilerplate.com) game.

## Start

Start all services

```bash
./script/start
```

## API Documentation

* [Player](http://localhost:8082/players/documentation)
* [Characters](http://localhost:8082/entities/documentation)
* [Spells](http://localhost:8082/spells/documentation)

## Stop

Stop all services

```bash
./script/stop
```

## Technical

Server

* Separate domain oriented services
  * Go
  * Postgres
* Player creation with OAuth/third party providers only

```plantuml
left to right direction
[Public API] ..> [Mage]
[Public API] ..> [Item]
[Public API] ..> [Spell]
[Public API] ..> [Fight]
[Public API] ..> [Tactic]
[Fight] ..> [Mage]
[Fight] ..> [Spell]
[Fight] ..> [Item]
[Fight] ..> [Tactic]
```

## License

COPYRIGHT 2021 ALIENSPACES alienspaces@gmail.com
