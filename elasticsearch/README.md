# Elasticsearch


## Environment
- macOS 12.4
- Elasticsearch 8.3.3


## Install
Run the following script to download and unarchive Elasticsearch package.
Update variable `ES_VERSION` in the `.env` and source it to install other
versions.
  ```bash
  cd elasticsearch
  source .env
  ./scripts/elasticsearch_install.sh
  ```


## Start Elasticsearch
Run the binary file in `elasticsearch-<es-version>/bin/elasticsearch` to start
elasticsearch on `127.0.0.1:9200` as default.
  ```bash
  ./elasticsearch-<es-version>/bin/elasticsearch
  ```

The first time starting Elasticsearch will generate password for the `elastic`
built-in superuser. Following the [official
website](https://www.elastic.co/guide/en/elasticsearch/reference/current/targz.html#targz-running)
to setup the authentication and authorization.

Alternatively, for testing locally, you can disable the authentication and
authorization by updating
`./elasticsearch-<es-version>/config/elasticsearch.yml`, setting
`xpack.security.enabled` to false, and restart Elasticsearch.

## Test Elasticsearch
The test data is forked from [News Category
Dataset](https://www.kaggle.com/datasets/rmisra/news-category-dataset) and
sampled 100 docs.

- test index: `news_category_test`
- test schema: `resources/test-news_category.es_schema.json`
- test data: `data/test-news_category.es.json`

### Create an Test Index
Run the setup script to create test index and index test data
`./scripts/elasticsearch_test_setup.sh`.

- Create a Test Index
  - `curl`
    ```bash
    curl -XPUT "$ES_HOST/$ES_INDEX_TEST" \
    -H 'Content-Type: application/json' \
    -d '@'"$GARAGE_ELASTICSEARCH_ROOT/resources/test-news_category.es_schema.json"
    ```
  - Elasticsearch Schema
    - settings
      ```json
      {
         "settings":{
            "index":{
               "number_of_shards":1,
               "number_of_replicas":1
            }
         }
      }
      ```
    - mappings
      ```json
      "mappings":{
          "properties":{
              "id":{
                  "type":"keyword"
              },
              "link":{
                  "type":"keyword"
              },
              "title":{
                  "type":"text",
                  "analyzer":"english"
              },
              "description":{
                  "type":"text",
                  "analyzer":"english"
              },
              "authors":{
                  "type":"keyword"
              },
              "category":{
                  "type":"keyword"
              },
              "created_at":{
                  "type":"date"
              },
              "updated_at":{
                  "type":"date"
              },
              "deleted_at":{
                  "type":"date"
              }
          }
      }
      ```

- Index Test Data
  - `curl`
    ```bash
    curl -XPOST "$ES_HOST/$ES_INDEX_TEST/_bulk" \
         -H 'Content-Type: application/json' \
         --data-binary '@'"$GARAGE_ELASTICSEARCH_ROOT/data/test-news_category.es.json"
    ```
  - Index Test Data Using `_bulk` API
    ```json
    {"index": {"_id": "1bfe7f10-e8da-4cde-a8e7-f84c1802a9c7"}}
    {
        "uuid": "1bfe7f10-e8da-4cde-a8e7-f84c1802a9c7",
        "link": "https://www.huffingtonpost.com/entry/texas-amanda-painter-mass-shooting_us_5b081ab4e4b0802d69caad89",
        "title": "There Were 2 Mass Shootings In Texas Last Week, But Only 1 On TV",
        "description": "She left her husband. He killed their children. Just another day in America.",
        "date": "2018-05-26",
        "authors": [
            "Melissa Jeltsen"
        ],
        "category": "CRIME"
    }
    ```


### Run Test Query on the Test Index

  - Get All docs
    ```bash
    curl -XGET "$ES_HOST/news_category_test/_search/?pretty" \
         -H "Content-Type: application/json" \
         -d '{
                 "from":0,
                 "size":'10',
                 "query": {"match_all": {}}
             }'
    ```

  - A Simple Keyword Search at `description` field
    ```bash
    curl -XGET "$ES_HOST/news_category_test/_search/?pretty" \
         -H "Content-Type: application/json" \
         -d '{
                 "from":0,
                 "size":'10',
                 "query": {
                     "match": {
                         "description": {
                             "query": "film"
                         }
                     }
                 }
             }'
    ```


- Get All Index
  ```bash
  curl -XGET "$ES_HOST/_cat/indices?v&pretty"
  ```


# Reference
- [News Category Dataset](https://www.kaggle.com/datasets/rmisra/news-category-dataset)
- [Install Elasticsearch from archive on Linux or MacOS](https://www.elastic.co/guide/en/elasticsearch/reference/current/targz.html#targz-running)
