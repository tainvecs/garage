# Meilisearch Debian
This repo provides Meilisearch Dockerfiles for Debian `amd64` and `arm64` (`aarch64`).

The docker images are available at the [DockerHub
repo](https://hub.docker.com/r/tainvecs/meilisearch-debian).

Alternatively, you can generate Dockerfiles and build docker images locally by
running the script `docker_build_local.zsh`.


## Environment
- macOS 12.4 M1 and macOS 12.0.1 Intel
- zsh
- Meilisearch 0.28.1


## Images on DockerHub
- `0.28.1-amd64`
  - Built and tested on my Intel MacBook Pro with macOS 12.0.1.
  - [Dockerfile.0.28.1-amd64](https://github.com/tainvecs/garage/blob/main/meilisearch/deployment/Dockerfile.0.28.1-amd64)

- `0.28.1-aarch64`
  - Built and tested on my M1 MacBook Pro with macOS 12.4.
  - [Dockerfile.0.28.1-aarch64](https://github.com/tainvecs/garage/blob/main/meilisearch/deployment/Dockerfile.0.28.1-aarch64)


## Build Your Own Docker Images
To build your own docker image, update the arguments (if needed) and run the
script `docker_build_local.zsh`.

- Build Docker Image
  - Run the script `docker_build_local.zsh` to generate a new Dockerfile based on the
    arguments, and build a new Docker image.
    ```sh
    zsh docker_build_local.zsh
    ```

- (Optional) Update Arguments in `docker_build_local.zsh`

  - `MEILISEARCH_VERSION`
    - the version number of [Meilisearch Release](https://github.com/meilisearch/meilisearch/releases)
    - default: `0.28.1`

  - `MEILISEARCH_ENV`
    - available options: `production` and `development`
    - default: `development`
    - ([more info](https://docs.meilisearch.com/learn/configuration/instance_options.html#environment))

  - `MEILISEARCH_HTTP_ADDR`
    - the HTTP address and port Meilisearch will use
    - default: `0.0.0.0:7700`

  - `MEILISEARCH_DB_PATH`
    - the location where database files will be created and retrieved
    - default: `/data.ms`
    - ([more info](https://docs.meilisearch.com/learn/configuration/instance_options.html#database-path))

  - `PLATFORM_OS`
    - currently only support `linux`

  - `DEBIAN_ARCHITECTURE`
    - available options: `amd64` and `arm64`
    - default: `arm64`
    - ([more info](https://github.com/docker-library/official-images#architectures-other-than-amd64))

  - `MEILISEARCH_ARCHITECTURE`
    - the architecture name of [Meilisearch Release](https://github.com/meilisearch/meilisearch/releases)
    - available options: `amd64` and `aarch64`
    - default: `aarch64`


## Reference
- [MeiliSearch Arm64 Docker](https://github.com/mukul-mehta/MeiliSearch-Arm64-Docker/blob/main/Dockerfile)
- [MeiliSearch Dockerfile](https://github.com/meilisearch/meilisearch/blob/main/Dockerfile)
- [MeiliSearch Env](https://docs.meilisearch.com/learn/configuration/instance_options.html#all-instance-options)
