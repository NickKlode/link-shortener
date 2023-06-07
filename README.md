Тестовое задание OZON | Сервис для сокращения ссылок

Запуск in-memory:
1. в docker-compose.yml "command: ./wait-for-postgres.sh db ./ozontest inmemory"
2. make build
3. make run

Запуск postgres:
1. в docker-compose.yml "command: ./wait-for-postgres.sh db ./ozontest postgres"
2. make build
3. make run

Пример:
POST запрос для получения короткой ссылки.


<img width="400" alt="Снимок экрана 2023-06-07 в 22 28 46" src="https://github.com/NickKlode/ozon-urlshortener/assets/83373008/81784161-1906-42e9-b97b-c5c3466e7297">


GET запрос для получения оригинального url.


<img width="400" alt="Снимок экрана 2023-06-07 в 22 29 12" src="https://github.com/NickKlode/ozon-urlshortener/assets/83373008/b2d11d16-7c01-4dbf-be60-1fb11e358f24">
