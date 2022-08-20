#!/bin/zsh


# ------------------------------------------------------------------------------
# args
#
# DEBIAN_ARCHITECTURE: amd64, arm32/v5, arm32/v7, arm64/v8, i386, mips64le, ppc64le, riscv64, s390x
# MEILISEARCH_ARCHITECTURE: amd64, aarch64
# MEILISEARCH_ENV: production, development
# ------------------------------------------------------------------------------


GARAGE_MS_ROOT="$(dirname $(cd $(dirname $0) >/dev/null 2>&1; pwd -P;))"


declare -A args_arr
args_arr[MEILISEARCH_VERSION]="0.28.1"
args_arr[MEILISEARCH_ENV]="development"
args_arr[MEILISEARCH_HTTP_ADDR]="0.0.0.0:7700"
args_arr[MEILISEARCH_DB_PATH]="/data.ms"
args_arr[PLATFORM_OS]="linux"
args_arr[DEBIAN_ARCHITECTURE]="amd64"
args_arr[MEILISEARCH_ARCHITECTURE]="amd64"


VCS_REF="<VCS_REF>"
BUILD_VERSION=$MEILISEARCH_VERSION
BUILD_DATE=$(date +'%Y-%m-%dT%H:%M:%S_%z')


# ------------------------------------------------------------------------------
# Dockerfile
# ------------------------------------------------------------------------------


# generate vars replace strings
declare -a replace_args_str_arr=()
for key val in ${(@kv)args_arr}; do
    replace_var_name_arr+="s|%%$key%%|$val|g; "
done


# join replace strings
function join_by { local IFS="$1"; shift; echo "$*"; }
dockerfile_var_replace_str=$(join_by "" "${replace_var_name_arr[@]}")


# generate Dockerfile
dockerfile_tag="${args_arr[MEILISEARCH_VERSION]}-${args_arr[MEILISEARCH_ARCHITECTURE]}"
dockerfile_template_path="$GARAGE_MS_ROOT/deployment/Dockerfile.template"
dockerfile_path="$GARAGE_MS_ROOT/deployment/Dockerfile.$dockerfile_tag"

sed $dockerfile_var_replace_str $dockerfile_template_path > $dockerfile_path


# ------------------------------------------------------------------------------
# build
# ------------------------------------------------------------------------------


# docker image
docker build \
       -f $dockerfile_path \
       -t "meilisearch-debian:$dockerfile_tag" \
       --build-arg BUILD_VERSION="$BUILD_VERSION" \
       --build-arg BUILD_DATE="$BUILD_DATE" \
       --build-arg VCS_REF="$VCS_REF" \
       .
