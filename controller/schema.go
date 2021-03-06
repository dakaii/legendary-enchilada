package controller

import (
	"errors"
	"graphyy/model"

	"github.com/graphql-go/graphql"
	"github.com/graphql-go/graphql/gqlerrors"
)

var authType = graphql.NewObject(graphql.ObjectConfig{
	Name: "Auth",
	Fields: graphql.Fields{
		"tokenType": &graphql.Field{
			Type: graphql.String,
		},
		"token": &graphql.Field{
			Type: graphql.String,
		},
		"expiresIn": &graphql.Field{
			Type: graphql.Int,
		},
	},
})

func getRootMutation(contrs *Controllers) *graphql.Object {
	return graphql.NewObject(graphql.ObjectConfig{
		Name: "RootMutation",
		Fields: graphql.Fields{
			"signup": &graphql.Field{
				Type:        authType,
				Description: "Signup",
				Args: graphql.FieldConfigArgument{
					"username": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.String),
					},
					"password": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.String),
					},
				},
				Resolve: func(params graphql.ResolveParams) (interface{}, error) {
					username, _ := params.Args["username"].(string)
					password, _ := params.Args["password"].(string)
					res, err := contrs.authController.Signup(model.User{Username: username, Password: password})
					if err != nil {
						return nil, gqlerrors.FormatError(err)
					}
					return res, nil
				},
			},
			"login": &graphql.Field{
				Type:        authType,
				Description: "Login",
				Args: graphql.FieldConfigArgument{
					"username": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.String),
					},
					"password": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.String),
					},
				},
				Resolve: func(params graphql.ResolveParams) (interface{}, error) {
					username, _ := params.Args["username"].(string)
					password, _ := params.Args["password"].(string)
					res, err := contrs.authController.Login(model.User{Username: username, Password: password})
					if err != nil {
						return nil, gqlerrors.FormatError(err)
					}
					return res, nil
				},
			},
		},
	})
}

func getRootQuery(contrs *Controllers) *graphql.Object {
	return graphql.NewObject(graphql.ObjectConfig{
		Name: "RootQuery",
		Fields: graphql.Fields{
			"me": &graphql.Field{
				Type: graphql.NewObject(graphql.ObjectConfig{
					Name: "Me",
					Fields: graphql.Fields{
						"username": &graphql.Field{
							Type: graphql.String,
						},
						"createdAt": &graphql.Field{
							Type: graphql.DateTime,
						},
					},
				}),
				Description: "Get the logged-in user's info",
				Resolve: func(params graphql.ResolveParams) (interface{}, error) {
					rootValue := params.Info.RootValue.(map[string]interface{})
					user := rootValue["currentUser"].(model.User)
					if user.Username == "" {
						return nil, gqlerrors.FormatError(errors.New("user not found"))
					}
					return user, nil
				},
			},
		},
	})

}
