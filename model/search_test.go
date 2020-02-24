package model

import (
	"io/ioutil"
	"strings"
	"testing"
)

var indexExample = `
{
  "_index" : "pets",
  "_type" : "_doc",
  "_id" : "8e8e21be-3e4a-4ddb-a3ed-bc507f88c8b7",
  "_version" : 1,
  "result" : "created",
  "forced_refresh" : true,
  "_shards" : {
    "total" : 2,
    "successful" : 1,
    "failed" : 0
  },
  "_seq_no" : 0,
  "_primary_term" : 1
}
`

var queryExample = `
{
  "took" : 182,
  "timed_out" : false,
  "_shards" : {
    "total" : 1,
    "successful" : 1,
    "skipped" : 0,
    "failed" : 0
  },
  "hits" : {
    "total" : {
      "value" : 1,
      "relation" : "eq"
    },
    "max_score" : 0.2876821,
    "hits" : [
      {
        "_index" : "pets",
        "_type" : "_doc",
        "_id" : "8e8e21be-3e4a-4ddb-a3ed-bc507f88c8b7",
        "_score" : 0.2876821,
        "_source" : {
          "name" : "sushi",
          "desc" : "He is a good boy."
        }
      }
    ]
  }
}
`

func TestBodyToIndexResponse(t *testing.T) {
	body := ioutil.NopCloser(strings.NewReader(indexExample))
	res, err := BodyToIndexResponse(body)
	if err != nil {
		t.Fatal(err)
	}

	if res.Index != "pets" {
		t.Fatalf("expected [%s], got [%s]", "pets", res.Index)
	}
}

func TestBodyToQueryResponse(t *testing.T) {
	body := ioutil.NopCloser(strings.NewReader(queryExample))
	res, err := BodyToQueryResponse(body)
	if err != nil {
		t.Fatal(err)
	}

	if len(res.Hits.Hits) != 1 {
		t.Fatalf("expected [%v], got [%v]", 1, len(res.Hits.Hits))
	}

	if res.Hits.Hits[0].Source.Name != "sushi" {
		t.Fatalf("expected [%v], got [%v]", "sushi", res.Hits.Hits[0].Source.Name)
	}
}
