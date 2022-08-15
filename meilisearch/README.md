# Meilisearch


## Environment
- macOS 12.4
- Meilisearch 0.27.2


## Install
Run the script `garage/meilisearch/scripts/install.sh` will download
`meilisearch` binary file into `meilisearch/bin/meilisearch`.
```bash
cd meilisearch
source .env
./scripts/install.sh
```


## Start Meilisearch
Run the binary file `meilisearch/bin/meilisearch` to start the search engine at
localhost (`http://127.0.0.1:7700`).
```bash
./bin/meilisearch
```


## Index
The test data is available at [movie json
file](https://docs.meilisearch.com/movies.json).
`data/movies-100.json` sampled the first 100 docs for testing.
- test index: `movies`
- test index settings: `resources/movies.settings.json`
- test data: `data/movies-100.json`

Run the `scripts/meilisearch_test_setup.sh` to create the test index, update the
settings, and index the test data.

- create test index
  ```sh
  curl -XPOST "$MEILISEARCH_URL/indexes" \
       -H 'Content-Type: application/json' \
       --data-binary '{
           "uid": "'$MEILISEARCH_INDEX'",
           "primaryKey": "id"
       }'
  ```

- update indexes settings
  ```sh
  curl -XPATCH "$MEILISEARCH_URL/indexes/$MEILISEARCH_INDEX/settings" \
       -H 'Content-Type: application/json' \
       -d"@$GARAGE_MEILISEARCH_ROOT/resources/movies.settings.json"
  ```

- index test docs
  ```sh
  curl -XPOST "$MEILISEARCH_URL/indexes/$MEILISEARCH_INDEX/documents" \
       -H 'Content-Type: application/json' \
       --data-binary "@$GARAGE_MEILISEARCH_ROOT/data/movies-100.json"
  ```

Alternatively, index test docs at `localhost` without changing the index settings.

- curl
  ```sh
  curl -X POST 'http://localhost:7700/indexes/movies/documents' \
       -H 'Content-Type: application/json' \
       --data-binary @data/movies-100.json
  ```

- python
  - install python meilisearch package by pip3
  ```sh
  pip3 install meilisearch
  ```
  - load json and index it into new index `movies`
  ```python
  import meilisearch
  import json

  json_file = open('movies.json')
  movies = json.load(json_file)

  client = meilisearch.Client('http://127.0.0.1:7700')
  client.index('movies').add_documents(movies)
  ```


## Search
Search query `city` on example index `movies`.

- curl
  ```sh
  curl \
    -X POST 'http://localhost:7700/indexes/movies/search' \
    -H 'Content-Type: application/json' \
    --data-binary '{ "q": "city" }'
  ```

- python
  ```python
  client.index('movies').search(city')
  ```


## Reference
- [Meilisearch Documentation](https://docs.meilisearch.com/)
