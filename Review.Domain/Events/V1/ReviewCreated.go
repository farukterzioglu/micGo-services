package eventsV1

import (
	"github.com/farukterzioglu/micGo-services/Review.Domain/Models"
)

// ReviewCreated ...
type ReviewCreated struct {
	Review models.Review `json:"review"`
}
