package swagger

import (
	"github.com/farukterzioglu/micGo-services/Review.API/dtos"
)

// Request containing review id
// swagger:parameters getReviewReq
type swaggGetReviewReq struct {
	// in:path
	// description: id of the review
	// type: string
	// required: true
	ReviewID string
}

// Request containing a review
// swagger:parameters createReviewReq
type swaggCreateReviewReq struct {
	// in:body
	// type: CreateReviewDto
	// required: true
	Body dtos.CreateReviewDto
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
	// Array of review models
	// in:body
	Body []dtos.ReviewDto
}

// HTTP status code 200 and a review model in data
// swagger:response reviewResp
type swaggReviewResp struct {
	// A review models
	// in:body
	Body dtos.ReviewDto
}
