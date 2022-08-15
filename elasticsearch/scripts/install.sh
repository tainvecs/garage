GARAGE_ELASTICSEARCH_ROOT="$(dirname $(cd $(dirname $0) >/dev/null 2>&1; pwd -P;))"


# if ES_VERSION not set, use 8.3.3
[[ -z ${ES_VERSION+x} ]] && ES_VERSION="8.3.3"


OS_NAME=`uname`

if [[ $OS_NAME = "Darwin" ]]; then

    curl -O "https://artifacts.elastic.co/downloads/elasticsearch/elasticsearch-$ES_VERSION-darwin-x86_64.tar.gz"
    curl "https://artifacts.elastic.co/downloads/elasticsearch/elasticsearch-$ES_VERSION-darwin-x86_64.tar.gz.sha512" | shasum -a 512 -c -
    tar -xzf "elasticsearch-$ES_VERSION-darwin-x86_64.tar.gz"
    cd "elasticsearch-$ES_VERSION/"

elif [[ $OS_NAME = "Linux" ]]; then

    wget "https://artifacts.elastic.co/downloads/elasticsearch/elasticsearch-$ES_VERSION-linux-x86_64.tar.gz"
    wget "https://artifacts.elastic.co/downloads/elasticsearch/elasticsearch-$ES_VERSION-linux-x86_64.tar.gz.sha512"
    shasum -a 512 -c "elasticsearch-$ES_VERSION-linux-x86_64.tar.gz.sha512"
    tar -xzf "elasticsearch-$ES_VERSION-linux-x86_64.tar.gz"
    cd "elasticsearch-$ES_VERSION/"

fi
