# Meilisearch


## Environment
- macOS 12.4
- Meilisearch 0.27.2


## Install
Run the script `garage/meilisearch/scripts/install.sh` will download
`meilisearch` binary file into `meilisearch/bin/meilisearch`.
```bash
cd meilisearch && ./scripts/install.sh
```


## Start Meilisearch
Run the binary file `meilisearch/bin/meilisearch` to start the search engine at
localhost (`http://127.0.0.1:7700`).
```bash
./bin/meilisearch
```


## Index
This is an example shows how to index a [movie json
file](https://docs.meilisearch.com/movies.json) to a new `movies` index at
`localhost`.

- curl
  ```sh
  curl \
    -X POST 'http://localhost:7700/indexes/movies/documents' \
    -H 'Content-Type: application/json' \
    --data-binary @movies.json
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
Search query `botman` on example index `movies`.
- curl
  ```sh
  curl \
    -X POST 'http://localhost:7700/indexes/movies/search' \
    -H 'Content-Type: application/json' \
    --data-binary '{ "q": "botman" }'
  ```

- python
  ```python
  client.index('movies').search('botman')
  ```


## Reference
- [Meilisearch Documentation](https://docs.meilisearch.com/)
