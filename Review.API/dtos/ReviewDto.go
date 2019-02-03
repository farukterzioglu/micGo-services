package dtos

// ReviewDto is a DTO for api requests
type ReviewDto struct {
	ID        string `json:"id"`
	ProductID string `json:"productid"`
	UserID    string `json:"userid"`
	Text      string `json:"text"`
	Star      int8   `json:"star"`
}
