# Python API Server
A simple API server created with Flask for Meilisearch.


## Environment
- macOS 12.4
- Python 3.10.4
  - Flask 2.2.1
  - meilisearch 0.19.1
- Meilisearch 0.27.2


## Installation
- Meilisearch
You can install and run `Meilisearch` application by referencing [here](https://github.com/tainvecs/garage/tree/main/meilisearch).

- Python, Flask and meilisearch Packages
  ```bash
  cd python_api_server && ./scripts/install.sh
  ```

- Setup Environment Variables
Please Create the `.env` file by referencing `python_api_server/.env_template`.
To source the .env file, you can run the command or simply use command line plugin `autoenv`.
  ```bash
  source python_api_server/.env
  ```


## Reference
- [Meilisearch Documentation](https://docs.meilisearch.com/learn/getting_started/quick_start.html#setup-and-installation)
- [Creating Web APIs with Python and Flask](https://programminghistorian.org/en/lessons/creating-apis-with-python-and-flask)
- [Flask API](https://flask.palletsprojects.com/en/2.2.x/api/)
- [Using flask inside class](https://stackoverflow.com/questions/40460846/using-flask-inside-class)
