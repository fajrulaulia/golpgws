# golpgws
Golang Webservice boilerplate


# development mode
- clone repo into local
``` bash
git clone git@github.com:fajrulaulia/golpgws.git
```

- check your setup base URL to DBMS PG server
``` bash
export DB_URL=postgres://golpgwsuser:golpgwspass@localhost:5432/golpgwsdb?sslmode=disable
```
- migrate to exec sql file to database
``` bash
make local-migrate
```
- Running into local
``` bash
make local-run
```
