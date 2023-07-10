# GetMark Home Assignment

В Makefile'е описаны все доступные цели для билда и запуска проекта. 

Сборка:

```shell
make build
```

Сборка Docker образа:

```shell
make docker-build
```

Отдельно существует одна цель, чтобы править всеми. Для запуска:

```shell
make run
```

Документация по API доступна в `./cmd/api/README.md`

После запуска API доступно по адресу http://localhost:8081