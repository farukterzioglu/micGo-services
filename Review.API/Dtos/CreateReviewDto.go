package dtos

// Review is a DTO for api requests
type Review struct {
	Text   string `json:"text"`
	Star   int8   `json:"star"`
	Status int8   `json:"status"`
}

// CreateReviewCommand is a DTO for create review command
type CreateReviewDto struct {
	Review Review `json:"review"`
}
