#!/bin/bash


APISRV_ROOT="$(dirname $(cd $(dirname $0) >/dev/null 2>&1; pwd -P;))"


# create index
curl -XPUT "$ES_URL/$ES_UNIT_TEST_INDEX" \
     -H 'Content-Type: application/json' \
     -d "@$APISRV_ROOT/resources/es-unit_test.schema.json"


# add index aliases
curl -XPOST "$ES_URL/_aliases" -H "Content-Type: application/json" -d'
{
   "actions": [
       { "add": {
            "index": "'$ES_UNIT_TEST_INDEX'",
            "alias": "'$ES_UNIT_TEST_INDEX_ALIAS'"
       }},
       { "add": {
            "index": "'$ES_UNIT_TEST_INDEX'",
            "alias": "'$ES_UNIT_TEST_SEARCH_ALIAS'"
       }}
   ]
}
'
