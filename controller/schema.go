package controller

import (
	"errors"
	"legendaryenchilada/model"

	"github.com/graphql-go/graphql"
	"github.com/graphql-go/graphql/gqlerrors"
)

var pointInterface = graphql.NewObject(graphql.ObjectConfig{
	Name: "Point",
	Fields: graphql.Fields{
		"title": &graphql.Field{
			Type: graphql.String,
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				if point, ok := p.Source.(model.Point); ok {
					return point.Title, nil
				}
				return nil, nil
			},
		},
		"lat": &graphql.Field{
			Type: graphql.NewNonNull(graphql.Float),
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				if point, ok := p.Source.(model.Point); ok {
					return point.Location.Coordinates[1], nil
				}
				return nil, nil
			},
		},
		"lng": &graphql.Field{
			Type: graphql.NewNonNull(graphql.Float),
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				if point, ok := p.Source.(model.Point); ok {
					return point.Location.Coordinates[0], nil
				}
				return nil, nil
			},
		},
	},
})

func getRootQuery(contrs *Controllers) *graphql.Object {
	return graphql.NewObject(graphql.ObjectConfig{
		Name: "RootQuery",
		Fields: graphql.Fields{
			"scooter": &graphql.Field{
				Args: graphql.FieldConfigArgument{
					"lat": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.Float),
					},
					"lng": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.Float),
					},
					"distance": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.Int),
					},
					"limit": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.Int),
					},
				},
				Type:        graphql.NewList(pointInterface),
				Description: "Get the scooters within the specified certain distance",
				Resolve: func(params graphql.ResolveParams) (interface{}, error) {
					lat, _ := params.Args["lat"].(float64)
					lng, _ := params.Args["lng"].(float64)
					distance, _ := params.Args["distance"].(int)
					limit, _ := params.Args["limit"].(int)

					// TODO find a better way to validate the inputs
					if lat < -90 || 90 < lat {
						return nil, gqlerrors.FormatError(errors.New("invalid value of lat"))
					}
					if lng < -180 || 180 < lng {
						return nil, gqlerrors.FormatError(errors.New("invalid value of lng"))
					}

					res, err := contrs.scooterController.GetNearbyScooters(lat, lng, int64(distance), int64(limit))
					if err != nil {
						return nil, gqlerrors.FormatError(err)
					}
					return res, nil
				},
			},
		},
	})

}
