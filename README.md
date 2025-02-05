# Building Service (Golang + PostgreSQL + Gin)

## 📌 Описание

Этот сервис предоставляет API для управления строениями. Реализованы следующие функции:

- Добавление нового строения в базу данных.
- Получение списка строений с возможностью фильтрации по городу, году (диапазон) и количеству этажей.
- Автоматическая генерация документации OpenAPI (Swagger).

## 🚀 Технологии

- **Golang** (Gin, GORM)
- **PostgreSQL** (ORM: GORM)
- **Swagger** (Документация: `swaggo/swag`)
- **YAML** (Конфигурация)

---

## 📂 Структура проекта

```
building-service/
│
├── config/                  # Конфигурация приложения
│   ├── config.go            # Чтение конфигурации
│   ├── config.yaml          # Файл конфигурации
│
├── database/                # Подключение к базе данных
│   ├── database.go          # Логика подключения к PostgreSQL
│
├── models/                  # Определение моделей
│   ├── building.go          # Модель Building
│
├── repository/              # Логика работы с БД
│   ├── building_repo.go     # Репозиторий для Building
│
├── handlers/                # HTTP-обработчики (контроллеры)
│   ├── building_handler.go  # Контроллер для Building
│
├── docs/                    # Сгенерированная OpenAPI документация
│
├── main.go                  # Точка входа в приложение
├── go.mod                   # Go модули
├── go.sum                   # Файл зависимостей
└── README.md                # Документация проекта
```

---

## 🔧 Установка и запуск

### 1️⃣ Клонируйте репозиторий

```sh
git clone https://github.com/your-repo/building-service.git
cd building-service
```

### 2️⃣ Установите зависимости

```sh
go mod tidy
```

### 3️⃣ Настройте базу данных

Перед запуском убедитесь, что PostgreSQL запущен и доступен.
Настройки подключения указываются в `config/config.yaml`:

```yaml
database:
  host: "localhost"
  port: 5432
  user: "postgres"
  password: "password"
  db: "buildings_db"
```

### 4️⃣ Запустите сервис

```sh
go run main.go
```

---

## 📌 API Эндпоинты

### 1️⃣ Добавление строения

**POST** `/buildings`

```json
{
  "name": "Empire State Building",
  "city": "New York",
  "year": 1931,
  "floors": 102
}
```

Ответ:

```json
{
  "id": 1,
  "name": "Empire State Building",
  "city": "New York",
  "year": 1931,
  "floors": 102
}
```

### 2️⃣ Получение списка строений с фильтрацией

**GET** `/buildings?city=New York&year_from=2000&year_to=2010&floors=102`
Ответ:

```json
[
  {
    "id": 1,
    "name": "Empire State Building",
    "city": "New York",
    "year": 2006,
    "floors": 102
  }
]
```

---

## 📜 Swagger-документация

Swagger-документация доступна по адресу:
[http://localhost:8080/swagger/index.html](http://localhost:8080/swagger/index.html)

Для генерации документации выполните:

```sh
go install github.com/swaggo/swag/cmd/swag@latest
swag init
```
