package graphql

import (
	"errors"
	"github.com/graphql-go/graphql"
	"github.com/graphql-go/handler"
	"github.com/graphql-go/relay"
	"golang.org/x/net/context"
	"net/http"
	"sab.com/domain/country"
	"sab.com/domain/university"
)

var nodeDefinitions *relay.NodeDefinitions

var countrySchema *CountrySchema
var universitySchema *UniversitySchema

type GraphqlHandler struct {
	h *handler.Handler
}

func (graphqlHandler GraphqlHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	graphqlHandler.h.ServeHTTP(w, r)
}

func GetGraphqlHandler(countryService *country.CountryService, universityService *university.UniversityService) *GraphqlHandler {
	return &GraphqlHandler{handler.New(&handler.Config{
		Schema: getSabGraphqlSchema(countryService, universityService),
		Pretty: true,
	})}
}

func getSabGraphqlSchema(countryService *country.CountryService, universityService *university.UniversityService) *graphql.Schema {
	if schema, err := createSchema(countryService, universityService); err != nil {
		panic(err)
	} else {
		return schema
	}
}

func createSchema(countryService *country.CountryService, universityService *university.UniversityService) (*graphql.Schema, error) {
	nodeDefinitions = relay.NewNodeDefinitions(relay.NodeDefinitionsConfig{
		IDFetcher: func(id string, info graphql.ResolveInfo, ctx context.Context) (interface{}, error) {
			// resolve id from global id
			resolvedID := relay.FromGlobalID(id)

			// based on id and its type, return the object
			switch resolvedID.Type {
			case "Country":
				return countrySchema.GetCountryByGlobalId(resolvedID.ID)
			case "University":
				return universitySchema.GetUniversityByGlobalId(resolvedID.ID)
			default:
				return nil, errors.New("Unknown node type")
			}
		},

		TypeResolve: func(p graphql.ResolveTypeParams) *graphql.Object {
			// based on the type of the value, return GraphQLObjectType
			switch p.Value.(type) {
			case UniversityNode:
				return universitySchema.GetUniversityType()
			default:
				return countrySchema.GetCountryType()
			}
		},
	})

	countrySchema = NewCountrySchema(countryService, nodeDefinitions)
	universitySchema = NewUniversitySchema(universityService, nodeDefinitions)

	queryType := graphql.NewObject(graphql.ObjectConfig{
		Name: "SabQuery",
		Fields: graphql.Fields{
			"countries":    countrySchema.GetCountriesQuery(),
			"country":      countrySchema.GetCountryQuery(),
			"universities": universitySchema.GetUniversitiesQuery(),
			"university":   universitySchema.GetUniversityQuery(),
			"node":         nodeDefinitions.NodeField,
		},
	})

	mutationType := graphql.NewObject(graphql.ObjectConfig{
		Name:        "SabMutation",
		Description: "All the mutation available on Study abroad apy",
		Fields: graphql.Fields{
			"createCountry":    countrySchema.GetCreateCountryMutation(),
			"createUniversity": universitySchema.GetCreateUniversityMutation(),
			"updateUniversity": universitySchema.GetUpdateUniversityMutation(),
		},
	})

	aSchema, aErr := graphql.NewSchema(graphql.SchemaConfig{
		Query:    queryType,
		Mutation: mutationType,
	})

	return &aSchema, aErr
}
