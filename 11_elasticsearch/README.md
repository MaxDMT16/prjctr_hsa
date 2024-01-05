# Elastic

To fit the requirement of **max 3 typos if word length is bigger than 7** I have used **ngram** tokenizer.
Here is a query to create index:
```json
PUT new_english_words
{
  "settings": {
    "max_ngram_diff": 7,
    "analysis": {
      "analyzer": {
        "ngram_analyzer": {
          "tokenizer": "ngram_tokenizer"
        }
      },
      "tokenizer": {
        "ngram_tokenizer": {
          "type": "ngram",
          "min_gram": 3,
          "max_gram": 8,
          "token_chars": ["letter", "digit"]
        }
      }
    }
  },
  "mappings": {
    "properties": {
      "word": {
        "type": "text",
        "analyzer": "ngram_analyzer"
      }
    }
  }
}
```

Data to the index was loaded using `app`.

To query documents I have used the following query:
```json
GET /new_english_words/_search
{
  "query": {
    "match": {
      "word": {
        "query": "trentepohbbb",
        "fuzziness": 3
      }
    }
  }
}
```

It gives me the following result:
```jsong
{
  "took" : 146,
  "timed_out" : false,
  "_shards" : {
    "total" : 1,
    "successful" : 1,
    "skipped" : 0,
    "failed" : 0
  },
  "hits" : {
    "total" : {
      "value" : 10000,
      "relation" : "gte"
    },
    "max_score" : 547.8214,
    "hits" : [
      {
        "_index" : "new_english_words",
        "_type" : "_doc",
        "_id" : "ue7l0IwBF1O5bWPFkFX-",
        "_score" : 547.8214,
        "_source" : {
          "word" : "trentepohlia"
        }
      },
      {
        "_index" : "new_english_words",
        "_type" : "_doc",
        "_id" : "PaSj2IwBaz9_msO-s_dL",
        "_score" : 547.8214,
        "_source" : {
          "word" : "trentepohlia"
        }
      },
      {
        "_index" : "new_english_words",
        "_type" : "_doc",
        "_id" : "Kzzl0IwB6J-hXPCBkeoD",
        "_score" : 429.87244,
        "_source" : {
          "word" : "trentepohliaceae"
        }
      },
      {
        "_index" : "new_english_words",
        "_type" : "_doc",
        "_id" : "p6Tl0IwBaz9_msO-kcEI",
        "_score" : 414.98114,
        "_source" : {
          "word" : "trentepohliaceous"
        }
      },
      {
        "_index" : "new_english_words",
        "_type" : "_doc",
        "_id" : "yO3g0IwBF1O5bWPFgPx2",
        "_score" : 183.60521,
        "_source" : {
          "word" : "rantepole"
        }
      },
      {
        "_index" : "new_english_words",
        "_type" : "_doc",
        "_id" : "XDvN0IwB6J-hXPCB5FCt",
        "_score" : 163.19458,
        "_source" : {
          "word" : "antepone"
        }
      },
      {
        "_index" : "new_english_words",
        "_type" : "_doc",
        "_id" : "6-zN0IwBF1O5bWPF5Lu6",
        "_score" : 163.19458,
        "_source" : {
          "word" : "anteport"
        }
      },
      {
        "_index" : "new_english_words",
        "_type" : "_doc",
        "_id" : "2KPN0IwBaz9_msO-5Ce0",
        "_score" : 149.29968,
        "_source" : {
          "word" : "anteporch"
        }
      },
      {
        "_index" : "new_english_words",
        "_type" : "_doc",
        "_id" : "9D3n0IwB6J-hXPCBBAQl",
        "_score" : 131.17262,
        "_source" : {
          "word" : "unrented"
        }
      },
      {
        "_index" : "new_english_words",
        "_type" : "_doc",
        "_id" : "WuzP0IwBF1O5bWPF_98A",
        "_score" : 130.2659,
        "_source" : {
          "word" : "brontephobia"
        }
      }
    ]
  }
}
```
