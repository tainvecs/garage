#!/bin/bash


APISRV_ROOT="$(dirname $(cd $(dirname $0) >/dev/null 2>&1; pwd -P;))"


# create index
curl -XPUT "$ES_URL/$ES_NEWS_DOC_INDEX" \
     -H 'Content-Type: application/json' \
     -d "@$APISRV_ROOT/resources/es-news_category.schema.json"


# add index aliases
curl -XPOST "$ES_URL/_aliases" -H "Content-Type: application/json" -d'
{
   "actions": [
       { "add": {
            "index": "'$ES_NEWS_DOC_INDEX'",
            "alias": "'$ES_NEWS_DOC_INDEX_ALIAS'"
       }},
       { "add": {
            "index": "'$ES_NEWS_DOC_INDEX'",
            "alias": "'$ES_NEWS_DOC_SEARCH_ALIAS'"
       }}
   ]
}
'


# # index test docs
# curl -XPOST "$ES_URL/$ES_NEWS_DOC_INDEX_ALIAS/_bulk" \
#      -H 'Content-Type: application/json' \
#      --data-binary "@$APISRV_ROOT/data/test-news_category.es.json"

# curl -XPOST "$ES_URL/_refresh?pretty"
