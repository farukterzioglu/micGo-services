package dtos

// ReviewDto is a DTO for api requests
type ReviewDto struct {
	ID   string `json:"id"`
	Text string `json:"text"`
	Star int8   `json:"star"`
}
