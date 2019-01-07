package commands

import "github.com/farukterzioglu/micGo-services/Review.CommandEngine/Models"

type CreateReviewCommand struct {
	Review models.Review `json:"review"`
}
