# garage
This is a testing space for different packages, frameworks, or libraries.
Each subdirectory is a sandbox for a small testing project.


## Implementations

### API Servers
| name                | description                                      | keywords                  |
|:--------------------|:-------------------------------------------------|---------------------------|
| [apisrv]            | Go API server implementation using gin and gorm. | `Elasticsearch`           |
| [Python API Server] | Python API server implementation using Flask.    | `Flask` and `Meilisearch` |

### Applications
| name       | description                                                                                                                                                    |
|:-----------|:---------------------------------------------------------------------------------------------------------------------------------------------------------------|
| [Fail2ban] | - shell scripts for banning IPs from a file with list of IPs or "denyhosts" artifact <br >- an Python script for parsing banned IP from a "denyhosts" artifact |


## Resources

### Datasets
| name            | description                                                                                                                     |
|:----------------|:--------------------------------------------------------------------------------------------------------------------------------|
| [News Categroy] | The dataset contains around 200k news headlines from the year 2012 to 2018 obtained from [HuffPost](https://www.huffpost.com/). |


## Notes and Testing

### Natural Language Processing
| name                 | description                                                         |
|:---------------------|:--------------------------------------------------------------------|
| [Keyword Extraction] | - Keyword extraction testing code using KeyBERT<br > - Test dataset |

### Search Engines
| name            | description                                                                                                                |
|:----------------|:---------------------------------------------------------------------------------------------------------------------------|
| [Elasticsearch] | - Elasticsearch installation and setup scripts<br > - Test dataset and schema                                              |
| [Meilisearch]   | - Meilisearch installation and setup scripts<br > - Test dataset<br > - Meilisearch Dockerfiles for Debian amd64 and arm64 |

### Databases
| name     | description                                                             |
|:---------|:------------------------------------------------------------------------|
| [Redis]  | Notes on Redis installation and command line.                           |
| [SQLite] | Notes on creating an SQLite database from a command line or SQL script. |


[apisrv]: https://github.com/tainvecs/garage/tree/main/apisrv
[Python API Server]: https://github.com/tainvecs/garage/tree/main/python_api_server

[Fail2ban]: https://github.com/tainvecs/garage/tree/main/fail2ban

[News Categroy]: https://github.com/tainvecs/garage/tree/main/datasets/news\_categroy

[Keyword Extraction]: https://github.com/tainvecs/garage/tree/main/keyword_extraction

[Elasticsearch]: https://github.com/tainvecs/garage/tree/main/elasticsearch
[Meilisearch]: https://github.com/tainvecs/garage/tree/main/meilisearch

[Redis]: https://github.com/tainvecs/garage/tree/main/redis
[SQLite]: https://github.com/tainvecs/garage/tree/main/sqlite
