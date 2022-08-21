# Golang API Server


## Environment
- macOS 12.0.1
- go 1.18.3
- elasticsearch 7.17.4


## TODO
- [x] Create New a Golang Project

- [ ] API
  - [ ] HTTP: `net/http` and `github.com/gin-gonic/gin`
  - [ ] elasticsearch handler

- [ ] Database: `github.com/go-gorm/gorm`
  - [ ] SQL database test dataset, schema and setup script
  - [ ] SQL database data access implementation
  - [ ] SQL database data access mock test
  - [ ] service db_dao
  - [ ] service db_dao test

- [ ] Elasticsearch: `github.com/elastic/go-elasticsearch/v8`
  - [x] elasticsearch test dataset, schema and setup script
  - [x] elasticsearch data access implementation
  - [ ] elasticsearch data access mock test
  - [x] service es_dao
  - [x] service es_dao test
