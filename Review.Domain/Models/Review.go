package models

import "context"

// Status defines the status of the review
type Status int

func (s Status) String() string {
	switch s {
	case Created:
		return "Created"
	case Approved:
		return "Approved"
	case Disapproved:
		return "Disapproved"
	case Removed:
		return "Removed"
	default:
		return "Unknown"
	}
}

const (
	// Created status means review created but not approved
	Created Status = iota + 1
	// Approved status means review created & approved
	Approved
	// Disapproved status means review created but disapproved
	Disapproved
	// Removed status means review created then disapproved
	Removed
)

// Review struct
type Review struct {
	Text   string
	Star   int8
	Status Status
}

// Storing review id in context ->
type key string

const reviewIDKey key = "reviewIDKey"

// NewContextWithReviewID returns a new Context that carries a provided review id value
func NewContextWithReviewID(ctx context.Context, reviewID string) context.Context {
	return context.WithValue(ctx, reviewIDKey, reviewID)
}

// ReviewIDFromContext extracts a review id from a Context
func ReviewIDFromContext(ctx context.Context) string {
	value := ctx.Value(reviewIDKey)

	if value != nil {
		return value.(string)
	}
	return ""
}
