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
    <img src="https://img.shields.io/badge/TAG-APISRV%2Fv0.0.1-blue?style=for-the-badge" alt="GitHub tag (hardcoded)">
  </a>
  <a href="https://github.com/tainvecs/garage/blob/main/apisrv/go.mod" >
    <img src="https://img.shields.io/github/go-mod/go-version/tainvecs/garage?filename=apisrv%2Fgo.mod&style=for-the-badge&logo=go" alt="GitHub go.mod Go version">
>>>>>>> apisrv/add-news-service
  </a>
</div>
<p></p>

This repo is a golang implementation of an API server.
I connect different applications and provide services through APIs. </br >

I am working on this project for testing and practice purposes.
Thus, most of the datasets will be a small test set.
Nevertheless, hoping that it can demonstrate how an API server can be implemented with golang and provide some references.


## Features
- (WIP) Elasticsearch:
    - search engine provides search services
    - `elastic/go-elasticsearch/v8`
- (WIP) PostgreSQL
    - SQL database for data storage
    - `go-gorm/gorm`
- (WIP) API Server
    - provides an access interface to backend services
    - `net/http`
    - `gin-gonic/gin`


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
