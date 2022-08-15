GARAGE_MEILISEARCH_ROOT="$(dirname $(cd $(dirname $0) >/dev/null 2>&1; pwd -P;))"


# If MEILISEARCH_URL or MEILISEARCH_INDEX not set, set default vaules.
[[ -z ${MEILISEARCH_URL+x} ]] && MEILISEARCH_URL="0.0.0.0:7700"
[[ -z ${MEILISEARCH_INDEX+x} ]] && MEILISEARCH_INDEX="movies"


# create indexes
curl -XPOST "$MEILISEARCH_URL/indexes" \
     -H 'Content-Type: application/json' \
     --data-binary '{
         "uid": "'$MEILISEARCH_INDEX'",
         "primaryKey": "id"
     }'


# update indexes settings
curl -XPATCH "$MEILISEARCH_URL/indexes/$MEILISEARCH_INDEX/settings" \
     -H 'Content-Type: application/json' \
     -d"@$GARAGE_MEILISEARCH_ROOT/resources/movies.settings.json"


# index test docs: movies-100.json
curl -XPOST "$MEILISEARCH_URL/indexes/$MEILISEARCH_INDEX/documents" \
     -H 'Content-Type: application/json' \
     --data-binary "@$GARAGE_MEILISEARCH_ROOT/data/movies-100.json"
