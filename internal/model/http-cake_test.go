package model

import (
	"testing"

	"gopkg.in/guregu/null.v4"
)

func TestApiMutationCakePayload_Validate(t *testing.T) {
	type fields struct {
		Title       string
		Description *string
		Rating      float32
		Image       *string
	}
	tests := []struct {
		name    string
		fields  fields
		wantErr bool
	}{
		{
			name: "basic mapping",
			fields: fields{
				Title:       "title",
				Description: null.StringFrom("desctiption").Ptr(),
				Rating:      4.50,
				Image:       null.StringFrom("image").Ptr(),
			},
			wantErr: false,
		},
		{
			name: "basic mapping contain nil field",
			fields: fields{
				Title:       "title",
				Description: nil,
				Rating:      4.50,
				Image:       nil,
			},
			wantErr: false,
		},
		{
			name: "rating out of range",
			fields: fields{
				Title:       "title",
				Description: nil,
				Rating:      5.50,
				Image:       nil,
			},
			wantErr: true,
		},
		{
			name: "title empty value",
			fields: fields{
				Title:       "",
				Description: nil,
				Rating:      4.50,
				Image:       nil,
			},
			wantErr: true,
		},
		{
			name: "description empty value",
			fields: fields{
				Title:       "title",
				Description: null.StringFrom("").Ptr(),
				Rating:      4.50,
				Image:       nil,
			},
			wantErr: true,
		},
		{
			name: "image empty value",
			fields: fields{
				Title:       "title",
				Description: nil,
				Rating:      4.50,
				Image:       null.StringFrom("").Ptr(),
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &ApiMutationCakePayload{
				Title:       tt.fields.Title,
				Description: tt.fields.Description,
				Rating:      tt.fields.Rating,
				Image:       tt.fields.Image,
			}
			if err := p.Validate(); (err != nil) != tt.wantErr {
				t.Errorf("ApiMutationCakePayload.Validate() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
