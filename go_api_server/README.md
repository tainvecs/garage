# Go API Server

This repo is a golang implementation of API server.
Currently, it is still working in progress, and only support Elasticsearch services.
In future, the API server might also support other services.


## Environment
- macOS 13.0.1
- Go 1.19.4
- Elasticsearch 7.17.8


## Components
- API
  - Server
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
  - [x] Elasticsearch Data Access Test
  - [ ] Service esdao
  - [ ] Service esdao Test

- [ ] Deployment
  - [x] Test CI
  - [ ] Build CI
  - [ ] Dockerfile


## Reference
- Elasticsearch
  - [Language Analyzer](https://www.elastic.co/guide/en/elasticsearch/reference/current/analysis-lang-analyzer.html)
  - [Index API](https://www.elastic.co/guide/en/elasticsearch/reference/current/docs-index_.html)
  - [Go update document using elasticsearch](https://stackoverflow.com/questions/71048446/go-update-document-using-elastic-search)

- GitHub Actions
  - [Service elasticsearch is not visible when run tests](https://stackoverflow.com/questions/64204333/service-elasticsearch-is-not-visible-when-run-tests)
  - [golangci-lint-action Error using golang 1.18](https://github.com/golangci/golangci-lint-action/issues/442)
