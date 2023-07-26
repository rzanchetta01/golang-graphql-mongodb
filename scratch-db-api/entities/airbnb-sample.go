package entities

import (
	"github.com/graphql-go/graphql"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Model struct {
	Id                    string               `json:"id,omitempty" bson:"_id,omitempty"`
	Listing_url           string               `json:"listing_url,omitempty" bson:"listing_url,omitempty"`
	Name                  string               `json:"name,omitempty" bson:"name,omitempty"`
	Summary               string               `json:"summary,omitempty" bson:"summary,omitempty"`
	Space                 string               `json:"space,omitempty" bson:"space,omitempty"`
	Description           string               `json:"description,omitempty" bson:"description,omitempty"`
	Neighborhood_overview string               `json:"neighborhood_overview,omitempty" bson:"neighborhood_overview,omitempty"`
	Notes                 string               `json:"notes,omitempty" bson:"notes,omitempty"`
	Transit               string               `json:"transit,omitempty" bson:"transit,omitempty"`
	Acess                 string               `json:"acess,omitempty" bson:"acess,omitempty"`
	Interaction           string               `json:"interaction,omitempty" bson:"interaction,omitempty"`
	House_rules           string               `json:"house_rules,omitempty" bson:"house_rules,omitempty"`
	Property_type         string               `json:"property_type,omitempty" bson:"property_type,omitempty"`
	Room_type             string               `json:"room_type,omitempty" bson:"room_type,omitempty"`
	Bed_type              string               `json:"bed_type,omitempty" bson:"bed_type,omitempty"`
	Minimum_nights        string               `json:"minimum_nights,omitempty" bson:"minimum_nights,omitempty"`
	Maximum_nights        string               `json:"maximum_nights,omitempty" bson:"maximum_nights,omitempty"`
	Cancellation_policy   string               `json:"cancellation_policy,omitempty" bson:"cancellation_policy,omitempty"`
	Bedrooms              int                  `json:"bedrooms,omitempty" bson:"bedrooms,omitempty"`
	Beds                  int                  `json:"beds,omitempty" bson:"beds,omitempty"`
	Accommodates          int                  `json:"accommodates,omitempty" bson:"accommodates,omitempty"`
	Price                 primitive.Decimal128 `json:"price,omitempty" bson:"price,omitempty"`
}

var ModelType = graphql.NewObject(graphql.ObjectConfig{
	Name: "Model",
	Fields: graphql.Fields{
		"Id":                    &graphql.Field{Type: graphql.String},
		"Listing_url":           &graphql.Field{Type: graphql.String},
		"Name":                  &graphql.Field{Type: graphql.String},
		"Summary":               &graphql.Field{Type: graphql.String},
		"Space":                 &graphql.Field{Type: graphql.String},
		"Description":           &graphql.Field{Type: graphql.String},
		"Neighborhood_overview": &graphql.Field{Type: graphql.String},
		"Notes":                 &graphql.Field{Type: graphql.String},
		"Transit":               &graphql.Field{Type: graphql.String},
		"Acess":                 &graphql.Field{Type: graphql.String},
		"Interaction":           &graphql.Field{Type: graphql.String},
		"House_rules":           &graphql.Field{Type: graphql.String},
		"Property_type":         &graphql.Field{Type: graphql.String},
		"Room_type":             &graphql.Field{Type: graphql.String},
		"Bed_type":              &graphql.Field{Type: graphql.String},
		"Minimum_nights":        &graphql.Field{Type: graphql.String},
		"Maximum_nights":        &graphql.Field{Type: graphql.String},
		"Cancellation_policy":   &graphql.Field{Type: graphql.String},
		"Bedrooms":              &graphql.Field{Type: graphql.Int},
		"Beds":                  &graphql.Field{Type: graphql.Int},
		"Accommodates":          &graphql.Field{Type: graphql.Int},
		"Price":                 &graphql.Field{Type: graphql.String},
	},
})
