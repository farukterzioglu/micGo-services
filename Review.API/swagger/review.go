package swagger

import (
	"github.com/farukterzioglu/micGo-services/Review.API/Dtos"
)

// Repository model request
// swagger:parameters createReviewReq
type swaggCreateReviewReq struct {
	// in:body
	Body dtos.Review
}

// Rate review model request
// swagger:parameters rateReviewReq
type swaggerRateReviewReq struct {
	// in:body
	Body dtos.RateReviewDto
}

// HTTP status code 200
// swagger:response rateReviewResp
type swaggRateReviewResp struct {
	// in:body
	Body struct {
		// HTTP status code 200 - Status OK
		Code int `json:"code"`
	}
}
