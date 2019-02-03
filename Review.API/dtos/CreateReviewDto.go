package dtos

// CreateReviewDto is a DTO for api requests
type CreateReviewDto struct {
	ProductID string `json:"productid"`
	UserID    string `json:"userid"`
	Text      string `json:"text"`
	Star      int8   `json:"star"`
}
