# Go API Server

<div>
  <a href="https://codecov.io/github/tainvecs/garage" >
    <img src="https://img.shields.io/codecov/c/github/tainvecs/garage?flag=apisrv&token=A508HNNW6R&style=for-the-badge&logo=codecov" alt="Codecov" >
  </a>
  <a href="https://github.com/tainvecs/garage/actions/workflows/apisrv_test.yaml" >
    <img src="https://img.shields.io/github/workflow/status/tainvecs/garage/APISRV%20Unit%20Tests?label=Unit%20Tests&style=for-the-badge&logo=github" alt="GitHub Test Workflow Status">
  <a>
  <a href="https://goreportcard.com/report/github.com/tainvecs/garage/apisrv" >
    <img src="https://goreportcard.com/badge/github.com/tainvecs/garage/apisrv?style=for-the-badge" alt="Go Report Card" >
  </a>
  <a href="https://github.com/tainvecs/garage/tags" >
    <img src="https://img.shields.io/github/v/tag/tainvecs/garage?style=for-the-badge" alt="GitHub tag (latest by date)">
  </a>
  <a href="https://github.com/tainvecs/garage/tree/main/apisrv">
    <img src="https://img.shields.io/github/repo-size/tainvecs/garage?style=for-the-badge" alt="Repo Size">
  </a>
</div>
<p></p>

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
