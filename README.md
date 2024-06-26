## Статия
Сервис ещё не завершён. Есть ошибки и недоработки, но можно понять задумку автора.

## Общее описание решения

- Сервис реализован на языке `Golang`.
- В качестве web фреймворка используется [chi](github.com/go-chi/chi/v5).
- Логирование операций в файл `file.log` или консоль, настраиваемое в файле конфигураций.
- Реализованы `Middlewares` для: логирования, проверки авторизации и прав доступа,
  а также кэширования.
- Сервис поднимается в `Docker` контейнерах: база данных, хранилище кэша и основное приложение.
- Контейнеры конфигурируются в  `docker-compose`.
- В качестве СУБД используется `PostgreSQL`. В качестве библиотеки для работы с запросами к `PostgreSQL` используется
  [pgxpool](https://github.com/jackc/pgx), а в качестве драйвера [pgx](https://github.com/jackc/pgx), позволяющие быстро обрабатывать запросы.
- В качестве хранилища кэша используется `Redis`. В качестве библиотеки для работы с `Redis` используется
  [rueidis](github.com/redis/rueidis).
- Взаимодействие с проектом организовано посредством `Makefile`.

### Сборка контейнера

```bash
make build-docker-banner
```
### Запуск всей системы

Для запуска необходимо выполнить следующую команду:

```bash
make run
```

### Остановка

Для остановки с сохранением контейнеров необходимо выполнить следующую команду:

```bash
make stop
```

Для полной остановки необходимо выполнить следующую команду:

```bash
make down
```