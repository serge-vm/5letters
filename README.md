# 5 Букв

Помошник в решении задании игры 5 букв

## Сборка

```bash
make
```

Для генерации документации при внесении изменений необходим пакет swag

```
go install github.com/swaggo/swag/cmd/swag@latest
```

## Параметры запуска

* host - на котором сервис принимает входящие соединения
* port - на котором сервис принимает входящие соединения
* base-url - если сервис необходимо разместить за http-сервером с определенным путём, например 5letters для работы на адресе http://localhost/5letters

## Документация OpenAPI

Доступ к интерфейсу swagger запущенного сервиса по http://localhost/swagger/index.html
Вместо localhost можно использовать внешний адрес