# Python API Server
A simple API server created with Flask for `Meilisearch`.


## Environment
- macOS 12.4
- Python 3.10.4
  - Flask 2.2.1
  - meilisearch 0.19.1
- Meilisearch 0.27.2


## Installation
- Meilisearch
You can install and run `Meilisearch` application by referencing
[here](https://github.com/tainvecs/garage/tree/main/meilisearch).

- Python, Flask and meilisearch Packages
  ```bash
  cd python_api_server && ./scripts/install.sh
  ```

- Setup Environment Variables
Please Create the `.env` file by referencing `python_api_server/.env_template`.
To source the .env file, you can run the command or simply use command line
plugin `autoenv`.
  ```bash
  source python_api_server/.env
  ```

  - Environment Variables
    - `FLASK_APP`: [reference](https://flask.palletsprojects.com/en/2.2.x/cli/)
    - `FLASK_HOST`: where Flask API server is hosted
    - `MEILISEARCH_URL`: where `Meilisearch` server is hosted
    - `MEILISEARCH_DOCUMENT_INDEX`: `Meilisearch` doc index name


## API Server
Run the following command to start `Meilisearch` API server.
```bash
python api/meilisearch_api.py
```


## API

- POST `/v1/ms/docs`
  - Index Documents
  - Test Docs
    - **data/movies_test.json**
    - a toy dataset with 100 docs from the subset of [sample
      docs](https://docs.meilisearch.com/learn/getting_started/quick_start.html#add-documents)
      on Meilisearch official website.
  - Test with `curl`
    ```bash
    curl -XPOST "$FLASK_URL/ms/v1/docs" -H "Content-Type: application/json" --data-binary @data/movies_test.json
    ```

- GET `/v1/ms/docs?query=""`
  - Search Documents
  - Test with `curl` with arg query \"lord\"
    ```bash
    curl -XGET "$FLASK_URL/ms/v1/docs?query=lord"
    ```


## Reference
- [Meilisearch Documentation](https://docs.meilisearch.com/learn/getting_started/quick_start.html#setup-and-installation)
- [Creating Web APIs with Python and Flask](https://programminghistorian.org/en/lessons/creating-apis-with-python-and-flask)
- [Flask API](https://flask.palletsprojects.com/en/2.2.x/api/)
- [Using flask inside class](https://stackoverflow.com/questions/40460846/using-flask-inside-class)
