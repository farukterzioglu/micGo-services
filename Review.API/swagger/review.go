package swagger

import (
	"github.com/farukterzioglu/micGo-services/Review.API/Dtos"
)

// Rate review model request
// swagger:parameters rateReviewReq
type swaggerRateReviewReq struct {
	// in:body
	Body dtos.RateReviewDto
}

// Success response
// swagger:response ok
type swaggScsResp struct {
	// in:body
	Body struct {
		// HTTP status code 200 - OK
		Code int `json:"code"`
	}
}
