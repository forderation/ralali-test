# Welcome to RalaliCakeApp!
Go project demo of cake api, config for golang are located at config.toml

# How to start
## Install go-migrate tools
this tools is only for running migration, the migration sql are located at db/migration directory
[Release Downloads](https://github.com/golang-migrate/migrate/releases)

```bash
$ curl -L https://github.com/golang-migrate/migrate/releases/download/$version/migrate.$os-$arch.tar.gz | tar xvz
```
### MacOS
```bash
$ brew install golang-migrate
```
### Windows
Using [scoop](https://scoop.sh/)
```bash
$ scoop install migrate
```
### Linux (*.deb package)
```bash
$ curl -L https://packagecloud.io/golang-migrate/migrate/gpgkey | apt-key add -
$ echo "deb https://packagecloud.io/golang-migrate/migrate/ubuntu/ $(lsb_release -sc) main" > /etc/apt/sources.list.d/migrate.list
$ apt-get update
$ apt-get install -y migrate
```

## Running thorough docker compose
### spawn required docker container services
```bash
docker compose -f docker-compose.yaml up -d --build
```
### running migration
db_url parameter is dsn related after docker container is spawned
```bash
make migrateup db_url="mysql://root:root@tcp(127.0.0.1:52000)/ralali?x-tls-insecure-skip-verify=true"
```
### docker status
type docker ps to see status
```bash
docker ps
2262ddb04fa8   mysql:8.0             "docker-entrypoint.s…"   3 minutes ago   Up 3 minutes   3306/tcp, 33060/tcp, 0.0.0.0:52000->52000/tcp   ralali-mysql_db_ralali-1
0834bacd4c22   api:ralali            "./main"                 3 minutes ago   Up 1 second    0.0.0.0:8081->8081/tcp                          api-ralali
```