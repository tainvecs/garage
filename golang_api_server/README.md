# Golang API Server
This repo is a golang implementation of API server.
Currently, it is still working in progress, and only support Elasticsearch services.
In future, the API server might also support other services.


## Environment
- macOS 12.0.1
- Go 1.18.3
- Elasticsearch 7.17.4


## Components
- API
  - Router
  - Handler

- Services
  - Elasticsearch Search and Index Services

- Data Access
  - Elasticsearch
  - PostgreSQL


## TODO
- [ ] Design and Documentation
  - [ ] Architecture Diagram
  - [ ] API Design
  - [x] Elasticsearch Schema Design
  - [ ] Database Schema Design

- [ ] API
  - [ ] HTTP: `net/http` and `github.com/gin-gonic/gin`
  - [ ] Elasticsearch Handler

- [ ] Database: `github.com/go-gorm/gorm`
  - [ ] SQL Database Test Data, Schema and Setup Script
  - [ ] SQL database Data Access Implementation
  - [ ] SQL database Data Access Mock Test
  - [ ] service db_dao
  - [ ] service db_dao test

- [ ] Elasticsearch: `github.com/elastic/go-elasticsearch/v8`
  - [x] Elasticsearch Test Dataset, Schema and Setup Script
  - [x] Elasticsearch Data Access Implementation
  - [ ] Elasticsearch Data Access Mock Test
  - [x] Service es_dao
  - [x] Service es_dao Test

- [ ] Deployment
  - [ ] Dockerfile
  - [ ] CI


## Reference

- Elasticsearch
  - [Language
    Analyzer](https://www.elastic.co/guide/en/elasticsearch/reference/current/analysis-lang-analyzer.html)
  - [Index
    API](https://www.elastic.co/guide/en/elasticsearch/reference/current/docs-index_.html)
  - [Go update document using
    elasticsearch](https://stackoverflow.com/questions/71048446/go-update-document-using-elastic-search)
