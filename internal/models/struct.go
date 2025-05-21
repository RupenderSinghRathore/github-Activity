package models

type Activity struct {
	Type string `json:"type"`
	Repo struct {
		Name string `json:"name"`
	} `json:"repo"`
	Payload struct {
		Ref_type string `json:"ref_type"`
		Commits  []struct {
			Message string `json:"message"`
		} `json:"commits"`
	} `json:"payload"`
	Created_at string `json:"Created_at"`
}
