package model

import (
	"encoding/json"
	"io"
	"io/ioutil"
)

// Contains complete response struct derived from Elasticsearch JSON response
// and helper functions to easily work with them

func BodyToIndexResponse(body io.ReadCloser) (*IndexResponse, error) {
	defer body.Close()
	bytes, err := ioutil.ReadAll(body)
	if err != nil {
		return nil, err
	}

	var res IndexResponse
	if err = json.Unmarshal(bytes, &res); err != nil {
		return nil, err
	}

	return &res, nil
}

type IndexResponse struct {
	Index         string `json:"_index"`
	Type          string `json:"_type"`
	ID            string `json:"_id"`
	Version       int    `json:"_version"`
	Result        string `json:"result"`
	ForcedRefresh bool   `json:"forced_refresh"`
	Shards        struct {
		Total      int `json:"total"`
		Successful int `json:"successful"`
		Failed     int `json:"failed"`
	} `json:"_shards"`
	SeqNo       int `json:"_seq_no"`
	PrimaryTerm int `json:"_primary_term"`
}

type QueryByNameResponse struct {
	Took     int  `json:"took"`
	TimedOut bool `json:"timed_out"`
	Shards   struct {
		Total      int `json:"total"`
		Successful int `json:"successful"`
		Skipped    int `json:"skipped"`
		Failed     int `json:"failed"`
	} `json:"_shards"`
	Hits struct {
		Total struct {
			Value    int    `json:"value"`
			Relation string `json:"relation"`
		} `json:"total"`
		MaxScore float64 `json:"max_score"`
		Hits     []struct {
			Index  string  `json:"_index"`
			Type   string  `json:"_type"`
			ID     string  `json:"_id"`
			Score  float64 `json:"_score"`
			Source struct {
				Name string `json:"name"`
				Desc string `json:"desc"`
			} `json:"_source"`
		} `json:"hits"`
	} `json:"hits"`
}

func BodyToQueryByNameResponse(body io.ReadCloser) (*QueryByNameResponse, error) {
	defer body.Close()
	bytes, err := ioutil.ReadAll(body)
	if err != nil {
		return nil, err
	}

	var res QueryByNameResponse
	if err = json.Unmarshal(bytes, &res); err != nil {
		return nil, err
	}

	return &res, nil
}
