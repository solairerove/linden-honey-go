# Linden Honey

> REST API for the lyrics of Yegor Letov - GoLang Edition

## Technologies

* [gorilla/mux](https://github.com/gorilla/mux)
* [Colly](https://github.com/gocolly/colly)

## Usage

### Local

Start application:
```
go run cmd/main.go
```

Start scrapper:
```
go run cmd/scrapper/scrapper.go
```

### Docker

Bootstrap db using docker-compose:
```
docker-compose up
```

Stop and remove containers, networks, images, and volumes:
```
docker-compose down
```
