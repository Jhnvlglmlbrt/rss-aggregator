<!-- <p align="center">
<img src="https://pepy.tech/badge/rss-aggregator" alt="https://pepy.tech/project/rss-aggregator">
<img src="https://pepy.tech/badge/rss-aggregator/month" alt="https://pepy.tech/project/rss-aggregator">
<img src="https://img.shields.io/github/license/Jhnvlglmlbrt/rss-aggregator.svg" alt="https://github.com/Jhnvlglmlbrt/rss-aggregator/blob/master/LICENSE"> -->

# ⚙️ Rss feed aggregator

#### # Создание RSS feed агрегатора, который позволяет клиенту:

- Добавлять RSS-каналы для сбора постов
- Подписываться и отписываться от RSS-каналов, добавленных другими пользователями
- Получать все последние посты из RSS-каналов, на которые подписан пользователь

## ❗ Requirements

- .env 
    ```
    PORT=8080
    DB_URL=postgres://postgres:Username@localhost:5432/dbname?sslmode=disable
    ```

## 💿 Installation

```
go get 
```

<!-- ## 💻 Example -->

## 🪛 How to use?

```
cd sql/schema && goose postgres postgres://postgres:Username@localhost:5432/dbname up
cd ../../ && make run
```


<!-- - **request_method** -  is a callable (like app.get, app.post, foo_router.patch and so on.).  
- **service_url** - the path to the endpoint on another service (like "https://microservice1.example.com").  
- **service_path** - the path to the method in microservice (like "/v1/microservice/users").  
- **gateway_path** - is the path to bind gateway.  
For example, your gateway api is located here - *https://gateway.example.com* and the path to endpoint (**gateway_path**) - "/users" then the full way to this method will be - *https://gateway.example.com/users*
- **override_headers** - Boolean value allows you to return all the headlines that were created by microservice in gateway.  
- **query_params** - used to extract query parameters from endpoint and transmission to microservice
- **form_params** -  used to extract form model parameters from endpoint and transmission to microservice
- **param body_params** - used to extract body model from endpoint and transmission to microservice -->

<!-- ⚠️ - **Be sure to transfer the name of the argument to the router, which is in the endpoint func!**   --> -->
