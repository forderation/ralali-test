package model

import (
	"errors"
	"strings"
)

type ApiGetCakesQuery struct {
	Page     int `form:"page"`
	PageSize int `form:"page_size" binding:"required"`
}

type ApiMutationCakePayload struct {
	Title       string  `json:"title" binding:"required"`
	Description *string `json:"description"`
	Rating      float32 `json:"rating" binding:"required"`
	Image       *string `json:"image"`
}

func (p *ApiMutationCakePayload) Validate() error {
	p.Title = strings.TrimSpace(p.Title)
	if len(p.Title) <= 0 {
		return errors.New("field 'title' cannot be empty")
	}
	if p.Description != nil {
		description := strings.TrimSpace(*p.Description)
		if len(description) <= 0 {
			return errors.New("field 'description' cannot be empty")
		}
		p.Description = &description
	}
	if p.Image != nil {
		image := strings.TrimSpace(*p.Image)
		if len(image) <= 0 {
			return errors.New("field 'image' cannot be empty")
		}
		p.Image = &image
	}
	if p.Rating < 0 || p.Rating > 5 {
		return errors.New("field 'rating' must be on range 0-5")
	}
	return nil
}
