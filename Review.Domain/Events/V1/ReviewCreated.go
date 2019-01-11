package eventsV1

import (
	"github.com/farukterzioglu/micGo-services/Review.Domain/Models"
)

type ReviewCreated struct {
	Review models.Review `json:"review"`
}
