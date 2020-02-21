package model

// Contains complete response struct derived from Elasticsearch JSON response
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

type QueryResponse struct {
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
