# Elasticsearch


## Environment
- macOS 12.4
- Elasticsearch 8.3.3


## Install
Run the following script to download and unarchive Elasticsearch package.
Update variable `ES_VERSION` in the script to install other versions.
  ```bash
  cd elasticsearch && ./scripts/install.sh
  ```


## Start Elasticsearch
Run the binary file in `elasticsearch-<es-version>/bin/elasticsearch` to start
elasticsearch on `127.0.0.1:9200` as default.
  ```bash
  ./elasticsearch-<es-version>/bin/elasticsearch
  ```

## Testing Elasticsearch
The first time starting Elasticsearch will generate password for the `elastic`
built-in superuser. Following the [official
website](https://www.elastic.co/guide/en/elasticsearch/reference/current/targz.html#targz-running)
to setup the authentication and authorization.

Alternatively, for testing locally, you can disable the authentication and
authorization by updating
`./elasticsearch-<es-version>/config/elasticsearch.yml`, setting
`xpack.security.enabled` to false, and restart Elasticsearch.
- get all index by `curl`
  ```bash
  curl -XGET "https://127.0.0.1:9200/_cat/indices?v&pretty"
  ```
