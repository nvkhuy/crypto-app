# Setup enviroment

## Flyway 5.2.4

go to dir conf/flyway

flyway migrate

## SQL BOILER

cd /volatiletech

git clone https://github.com/volatiletech/sqlboiler.git

cd /sqlboiler, then run go get

cd /sqlboiler/drivers/sqlboier-mysql, then run go get

sqlboiler -c .\sqlboiler.toml mysql --no-context

## Generate Unit Test

gotests [options] PATH ...

-all generate tests for all functions and methods

-w write output to (test) files instead of stdout

## GO MOD

go mod init

## Heroku

create app then go to setting -> reveal config then add key GOVERSION value 1.13