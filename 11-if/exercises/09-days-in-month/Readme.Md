# Titles

## Index Creating

> With 5 shard and 0 replicas according to our data node count.

<details>
  <summary>Create an Index</summary>

  ```json
PUT /dba-test-000001
{
  "settings": {
    "index": {
      "number_of_shards": 5,  
      "number_of_replicas": 0 
    }
  }
}


  ```
</details>


## Index Life Cycle Mamagement
### Index Template

> Here we specify an index template for the index dba-test when to rollovered with exact settings and not to loose our shard-replications as well.

<details>
  <summary>Apply Index Template</summary>

  ```json
PUT /_index_template/dba-test
{
  "index_patterns": [
    "dba-test-*"
  ],
    "template": {
      "settings": {
        "index": {
          "lifecycle": {
            "name": "dba-test",
            "rollover_alias": "dba-test"
      },
      "routing": {
        "allocation": {
          "require": {
            "data": "hot"
          }
        }
      },
      "mapping": {
        "ignore_malformed": "true"
      },
      "number_of_shards": "5",
      "number_of_replicas": "0"
    }
  }
}
}

  ```
</details>


### Index Alias Uygulama 

> We should create&apply our alias to manage our ILM and other settings for our index each by index names. This let us prevent confusing between large ES clusters.

<details>

  <summary>Create and Apply aliase to our index</summary>

  ```json
PUT dba-test-*/_settings
{
  "index.lifecycle": {
    "name": "dba-test",
    "rollover_alias": "dba-test"
  }
}

  ```
</details>

### Index Life Cycle Policy Apply

> In this section we tell through our policy when and in which conditions to be rollovered.
    
    | Hot age day : Specifices for how many day we would like to store data.
    | Delete age day: Index will be deleted after rollovered.
    | Hot size: Specifices after how much disk space allocated our index to be rollovered. 

<details>
  <summary>ILM Policy Apply</summary>

  ```json
PUT _ilm/policy/dba-test
{
    "policy": {
      "phases": {
        "hot": {
          "min_age": "0ms",
          "actions": {
            "rollover": {
              "max_age": "1d",
              "max_size": "30gb"
            },
            "set_priority": {
              "priority": 100
            }
          }
        },
        "warm": {
          "actions": {
            "allocate": {
              "number_of_replicas": 0,
              "include": {},
              "exclude": {},
              "require": {
                "data": "${ES_WARM_PHASE_NODE_SELECTOR}"
              }
            },
            "forcemerge": {
              "max_num_segments": 1
            },
            "set_priority": {
              "priority": 50
            }
          }
        },
        "cold": {
          "min_age": "0d",
          "actions": {
            "allocate": {
              "include": {},
              "exclude": {},
              "require": {
                "data": "hot"
              }
            },
            "set_priority": {
              "priority": 0
            }
          }
        },
        "delete": {
          "min_age": "7d",
          "actions": {
            "delete": {}
          }
        }
      }
    }
  }

  ```
</details>


## Index Search
### Match_all

<details>
  <summary>Example of match_all</summary>

  ```json
GET /dba-test/_search
{
  "query": {
    "match_all": {}
  }
}

  ```
</details>

### Multi_match

<details>
  <summary>Example of multi_match</summary>
  
  ```json
GET /dba-test/_search
{
  "query": {
    "multi_match": {
      "query": "burton",
      "fields": ["address","firstname"]
      }
    }
}

  ```
</details>

### Wildcard Match

> Example for the value matching with all possibilites.

<details>
  <summary>Example of wildcard match</summary>

  ```json
GET /dba-test/_search
{
  "query": {
    "wildcard": {
      "lastname": {
        "value": "o***ns"
        }
    }
  }
}

  ```
</details>

### Match Phrase

> Example for the string we looking for.

<details>
  <summary>Example of match phrase</summary>
  
  ```json
GET /dba-test/_search
{
  "query": {
    "match_phrase": {
      "employer": "eventex"
    }  
  }
}

  ```
</details>

### Compound Query

> A compound query is a combination of two or more queries that are joined together using a logical operator. There are three types of logical operators that can be used to combine queries in Elasticsearch.

<details>
  <summary>Example of compound query for bool</summary>
  
  ```json
GET /dba-test/_search
{
  "query": {
    "bool": {
      "must": [
        {
          "term": {
            "address": {
              "value": "street"
            }
          }
        }
      ]
      ,"must_not": [
        {
          "term": {
            "age": {
              "value": 27
            }
          }
        }
      ]
    }
  }
}

  ```
</details>

<details>
  <summary>Example of compound query for constant_score</summary>
  
  ```json
GET dba-test/_search
{
  "query": {
    "constant_score": {
      "filter": {
        "term": {
          "gender": "m"
        }
      },
      "boost": 1.2
    }
  }
}
  ```
</details>

    Boost scores can be applied to specific term or overall query as we used. This indicates relevance score and can be any positive or negative number.
    This can be seen by "max_score" : 1.0 if not used or 1.2 if used as the same way.

<details>
  <summary>Another example of boosting score query</summary>
  
  ```json
GET dba-test/_search
{
  "query": {
   "boosting": {
     "negative": {
       "term": {
         "gender": {
           "value": "m"
         }
       }
     },
     "positive": {
       "term": {
         "gender": {
           "value": "f"
         }
       }
     },
     "negative_boost": 0.5
   }
  }
}
  ```
</details>

## Query Costs
### Expensive Queries
Our wildcard search is very common example of slow query because of scanning millions of documents.

> search.allow_expensive_queries: false

    can be set (by default "enabled") to disallow about latencies and long time lapses.
### Fuzzy Search

    Uses "Levenshtein Distance"
    Returns documents that similar to search term.

<details>
  <summary>Fuzzy Search</summary>
  
  ```json
GET dba-test/_search
{
  "query": {
    "fuzzy": {
      "address": {
        "value": "strets"
      }
    }
  }
}
  ```
</details>


<details>
  <summary>Regexp Search</summary>
  
  ```json
GET dba-test/_search
{
  "query": {
    "regexp": {
      "city": {
        "value": ".*ksonbu.*"
      }
    }
  }
}
  ```
</details>

> "." Matches any character

> "+" Repeath the preciding character one or more times

> "?" Repeat the preceding character

> "|" Or operator to split left amd right side

### Shard
Documents within an index distributed across multiple nodes. 

Gain disk size heavines and search latency.

Collect result from each shard.

Primary shard is static setting which must be decided while creating index.

> Best Practise Teorically -> Nodes = primary * ( replica + 1)

> We dont use shard routing which assing some documents to specific shard. We use alias overall index to get search result but for the same index we can create more aliases with routing and serach via those aliases but keep in mind these queries will be read by the specific node only.
