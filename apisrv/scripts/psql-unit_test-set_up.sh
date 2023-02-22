#!/bin/bash

# create db: news
psql -h $POSTGRES_HOST -U $POSTGRES_USER -p $POSTGRES_PORT -d postgres -c "CREATE DATABASE $POSTGRES_NEWS_DB;"

# create tables for db: news
psql -h $POSTGRES_HOST -U $POSTGRES_USER -p $POSTGRES_PORT -d $POSTGRES_NEWS_DB -f data/test-news_category.psql.table_create.sql

# insert data to db: news
psql -h $POSTGRES_HOST -U $POSTGRES_USER -p $POSTGRES_PORT -d $POSTGRES_NEWS_DB -f data/test-news_category.psql.data_insert.sql
