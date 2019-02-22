package commands

import (
	"github.com/farukterzioglu/micGo-services/Review.Domain/Models"
)

type CreateReviewCommand struct {
	Review models.Review
}
