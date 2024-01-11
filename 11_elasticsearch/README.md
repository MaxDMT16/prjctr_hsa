# Elastic

To fit the requirement of **max 3 typos if word length is bigger than 7** I have used **ngram** tokenizer.
Here is a request to create index:
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
          "min_gram": 2,
          "max_gram": 3,
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

As a word for testing I have used `trental`. It has length of 7 and it is in the loaded  dictionary.
Analyzer of this index will split this words to many tokens with min length of 2 and max length of 3.

If you try to search this word with the following request:
```json
GET new_english_words/_search
{
  "size": 10000, 
  "query": {
    "match": {
      "word": {
        "query": "txextxl", 
        "fuzziness": "AUTO"
      }
    }
  }
}
```
The doc with this word will be returned. However it's score is quire low.
```json
{
  "_index" : "new_english_words",
  "_type" : "_doc",
  "_id" : "pqTl0IwBaz9_msO-kMH6",
  "_score" : 10.513338,
  "_source" : {
    "word" : "trental"
  }
}
```

So, the requirement is met.