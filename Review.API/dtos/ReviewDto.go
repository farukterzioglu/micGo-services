package dtos

// Review is a DTO for api requests
type ReviewDto struct {
	Text string `json:"text"`
	Star int8   `json:"star"`
}
