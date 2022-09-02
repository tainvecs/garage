# News Category Dataset

The test data is forked from [News Category
Dataset](https://www.kaggle.com/datasets/rmisra/news-category-dataset).

It contains around 200k news headlines from the year 2012 to 2018 obtained from
[HuffPost](https://www.huffpost.com/).


## Raw Data
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


## Processed Data
- split authors string into list
- extract news_id from url
- update keys' names

```json
{
    "uuid":"8a34bbab-125a-4beb-b9b0-5b9918e83a78",
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

## Processed Data for Elasticsearch Indexing
```json
{
    "index":{
        "_id":"8a34bbab-125a-4beb-b9b0-5b9918e83a78"
    }
}

{
    "uuid":"8a34bbab-125a-4beb-b9b0-5b9918e83a78",
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


## Reference
- [News Category Dataset](https://www.kaggle.com/datasets/rmisra/news-category-dataset)
