# News Category Dataset

The test data is derived from [News Category Dataset].<br />
This dataset contains approximately 200,000 news headlines collected between 2012 and 2018 from [HuffPost].


## Preprocessing
- `resources/datasets/news_category/preprocess`

- Check `process_dataset.ipynb` for data preprocessing detail.
  - split authors string into list
  - extract news_id from url
  - update keys' names

- Raw Data
```json
{
    "link":"https://www.huffingtonpost.com/entry/amazon-prime-what-to-watch_us_5b044625e4b0c0b8b23ec14f",
    "headline":"What To Watch On Amazon Prime That\u2019s New This Week",
    "short_description":"There's a great mini-series joining this week.",
    "date":"2018-05-26"
    "authors":"Todd Van Luling",
    "category":"ENTERTAINMENT",
}
```

- Processed Data
```json
{
    "uuid":"8a34bbab-125a-4beb-b9b0-5b9918e83a78",
    "news_id":"5b044625e4b0c0b8b23ec14f",
    "link":"https://www.huffingtonpost.com/entry/amazon-prime-what-to-watch_us_5b044625e4b0c0b8b23ec14f",
    "title":"What To Watch On Amazon Prime That\u2019s New This Week",
    "description":"There's a great mini-series joining this week.",
    "date":"2018-05-26",
    "authors":[
        "Todd Van Luling"
    ],
    "category":"ENTERTAINMENT"
}
```


## Elasticsearch
- `resources/datasets/news_category/es`

- Processed Data for Indexing
```json
{
    "index":{
        "_id":"8a34bbab-125a-4beb-b9b0-5b9918e83a78"
    }
}
{
    "uuid":"8a34bbab-125a-4beb-b9b0-5b9918e83a78",
    "news_id":"5b044625e4b0c0b8b23ec14f",
    "link":"https://www.huffingtonpost.com/entry/amazon-prime-what-to-watch_us_5b044625e4b0c0b8b23ec14f",
    "title":"What To Watch On Amazon Prime That\u2019s New This Week",
    "description":"There's a great mini-series joining this week.",
    "date":"2018-05-26",
    "authors":[
        "Todd Van Luling"
    ],
    "category":"ENTERTAINMENT"
}
```


## PostgreSQL
- `resources/datasets/news_category/psql`

- Check schema file `schema-news_category.psql.sql` for more detail.

- Processed Data for Indexing
```sql
INSERT INTO "news" ("id", "uuid", "link", "title", "description", "created_at", "category") VALUES
(8, '8a34bbab-125a-4beb-b9b0-5b9918e83a78', 'https://www.huffingtonpost.com/entry/amazon-prime-what-to-watch_us_5b044625e4b0c0b8b23ec14f', 'What To Watch On Amazon Prime That’s New This Week', 'There''s a great mini-series joining this week.', '2018-05-26', 'ENTERTAINMENT');

INSERT INTO "authors" ("id", "name") VALUES
(8, 'Todd Van Luling');

INSERT INTO "news_authors" ("news_id", "authors_id") VALUES
(8, 8);
```


## Reference
- [News Category Dataset]


[HuffPost]: https://www.huffpost.com/
[News Category Dataset]: https://www.kaggle.com/datasets/rmisra/news-category-dataset
