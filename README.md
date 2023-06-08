Тестовое задание OZON | Сервис для сокращения ссылок

Запуск in-memory:
1. в docker-compose.yml "command: ./wait-for-postgres.sh db ./ozontest inmemory"
2. make build
3. make run

Запуск postgres:
1. в docker-compose.yml "command: ./wait-for-postgres.sh db ./ozontest postgres"
2. make build
3. make run
4. make migration

Пример:
POST запрос для получения короткой ссылки.

<img width="400" alt="Снимок экрана 2023-06-08 в 11 19 31" src="https://github.com/NickKlode/ozon-urlshortener/assets/83373008/a133bb02-8d7c-4234-b15c-09f8c62d94ad">


GET запрос для получения оригинального url.

<img width="400" alt="Снимок экрана 2023-06-08 в 11 19 53" src="https://github.com/NickKlode/ozon-urlshortener/assets/83373008/b9a6a173-0a5c-4f1c-8281-29aafc84e30e">
