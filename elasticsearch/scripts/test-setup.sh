GARAGE_ELASTICSEARCH_ROOT="$(dirname $(cd $(dirname $0) >/dev/null 2>&1; pwd -P;))"


# if ES_HOST not set, use localhost
[[ -z ${ES_HOST+x} ]] && ES_HOST="localhost:9200"

ES_INDEX_TEST="news_category_test"


# create test index
curl -XPUT "$ES_HOST/$ES_INDEX_TEST" \
     -H 'Content-Type: application/json' \
     -d '@'"$GARAGE_ELASTICSEARCH_ROOT/resources/test-news_category.es_schema.json"

# index test data
curl -XPOST "$ES_HOST/$ES_INDEX_TEST/_bulk" \
     -H 'Content-Type: application/json' \
     --data-binary '@'"$GARAGE_ELASTICSEARCH_ROOT/data/test-news_category.es.json"

curl -XPOST "$ES_HOST/_refresh?pretty"
