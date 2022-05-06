package transaction

import (
	"fita/Controller/DTO/request"
	"github.com/graphql-go/graphql"
)

type SchemaGqlItems interface {
	SchemaItems (nameSchema string) SchemaGqlItems
	Result() (*graphql.Result, error)
}

type schemaQuery struct {
	Query  string
	Object *graphql.Object
}

type schemaMutation struct {
	Mutation string
	Object   *graphql.Object
}

func RepositoryNew(reqBody request.ReqBody) SchemaGqlItems {
	if reqBody.Query != "" {
		return &schemaQuery{
			Query: reqBody.Query,
		}
	}

	return &schemaMutation{
		Mutation: reqBody.Mutation,
	}
}

const (
	kodePromo1  = "43N23P"
	kodePromo2  = "120P90"
	kodePromo3  = "A304SD"
	productFree = "Raspberry Pi B"
)
