package swagger

import (
	"github.com/farukterzioglu/micGo-services/Review.API/Dtos"
)

// Request containing a review
// swagger:parameters createReviewReq
type swaggCreateReviewReq struct {
	// in:body
	// type: ReviewDto
	// required: true
	Body dtos.ReviewDto
}

// parameters:
// - name: reviewId
//   in: path
//   description: id of the review
//   type: string
//   required: true

// Request containing a review rate
// swagger:parameters rateReviewReq
type swaggerRateReviewReq struct {
	// in:path
	// description: id of the review
	// type: string
	// required: true
	ReviewID string
	// in:body
	// type: ReviewRatingDto
	// required: true
	Body dtos.ReviewRatingDto
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

// HTTP status code 200 and an array of review models in data
// swagger:response reviewsResp
type swaggReviewsResp struct {
	// in:body
	Body struct {
		// HTTP status code 200 - Status OK
		Code int `json:"code"`
		// Array of review models
		Data []dtos.ReviewDto `json:"data"`
	}
}
