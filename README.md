# Online song library

Микросервис онлайн библиотеки песен на базе `REST API`. <br>
Программа разработана на языке программирования Go с генерацией `Swagger` документацией на реализованное API. Сборка и запуск проекта осуществляется в `Docker` контейнере.<br>
Миграция базы данных на основе SQL скрипта с использованием утилиты `migrate`.<br>

API микросервиса имеет три маршрута: `/swagger/*any`, `/api/songs`, `/api/song`.<br>
Маршрут `/api/song` имеет 4 операции: `GET`, `DELETE`, `PATCH`, `POST`.


## /swagger/*any
Swagger документация на реализованное API микросервиса. <br>

## /api/songs
Получение данных библиотеки с фильтрацией по всем полям и пагинацией.<br>

### Требуемые данные: <br>
`page` - обязательно: "номер страницы"<br>
`sizePage` - обязательно: "размер одной страницы"<br>
`group` - опционально: "музыкальная группа"<br>
`song` - опционально: "название песни"<br>
`releaseDate` - опционально: "дата релиза песни"<br>
`text` - опционально: "текст песни"<br>
`link` - опционально: "ссылка на песню"<br>

## GET /api/song
Получение текста песни с пагинацией по куплетам.<br>

### Требуемые данные: <br>
`id` - обязательно: "id песни"<br>
`page` - обязательно: "номер страницы"<br>
`sizePage` - обязательно: "размер одной страницы"<br>

## DELETE /api/song
Удаление песни.<br>

### Требуемые данные: <br>
`id` - обязательно: "id песни"<br>

## PATCH /api/song
Изменение данных песни.<br>

### Требуемые данные: <br>
`id` - обязательно: "id песни"<br>
`group` - опционально: "музыкальная группа"<br>
`song` - опционально: "название песни"<br>
`releaseDate` - опционально: "дата релиза песни"<br>
`text` - опционально: "текст песни"<br>
`link` - опционально: "ссылка на песню"<br>

## POST /api/song
Добавление новой песни с обогащением информации через сторонее API.<br>

### Требуемые данные: <br>
`group` - обязательно: "музыкальная группа"<br>
`song` - обязательно: "название песни"<br>

## Installation
Для сборки с помощью Makefile необходим make: `brew install make`<br>
Сборка и запуск микросервиса осуществляется через команду Makefile: `make dockerCompose`.<br>
Остановка микросервиса: `make dockerStop`.<br>