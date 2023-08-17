# Traveland-REST-API
# API Для Туристического приложения
## Первый опыт в создании api
---
* Подход Чистой Архитектуры в построении структуры приложения. Техника внедрения зависимости.
* Работа с фреймворком gin-gonic/gin.
* Работа с БД Postgres. Запуск из Docker. Генерация файлов миграций.
## Список Возможностей:
*  Регистрация пользователя 
*  Аутентификация пользователя через jwt и refresh токены
*  CRUD по сущиностям проекта
*  Работа с DOCKER + DOCKER COMPOSE
# Для запуска приложения:
```
docker-compose up --build
```
Если приложение запускается впервые, необходимо применить миграции к базе данных:
```
migrate -path ./db/migrations -database 'postgres://postgres:<password>@localhost:5432/postgres?sslmode=disable' up
```
