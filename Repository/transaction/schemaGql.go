package transaction

import (
	"fita/Controller/DTO/response"
	"github.com/graphql-go/graphql"
)

func (obj *schemaQuery) SchemaItems(nameSchema string) SchemaGqlItems {
	objItems := graphql.NewObject(graphql.ObjectConfig{
		Name:        "Items",
		Fields: graphql.Fields{
			"sku": &graphql.Field{
				Type:        graphql.String,
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					if item, ok := p.Source.(*response.CheckingOutDTO); ok {
						return item.Sku, nil
					}

					return nil, nil
				},
			},
			"name": &graphql.Field{
				Type:        graphql.String,
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					if item, ok := p.Source.(*response.CheckingOutDTO); ok {
						return item.Name, nil
					}

					return nil, nil
				},
			},
			"total_price": &graphql.Field{
				Type:        graphql.String,
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					if item, ok := p.Source.(*response.CheckingOutDTO); ok {
						return item.TotalPrice, nil
					}

					return nil, nil
				},
			},
			"total_qty": &graphql.Field{
				Type:        graphql.Int,
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					if item, ok := p.Source.(*response.CheckingOutDTO); ok {
						return item.TotalQty, nil
					}

					return nil, nil
				},
			},
		},
	})

	if nameSchema == objItems.Name() {
		obj.Object = objItems
	}

	return obj
}

func (obj *schemaMutation) SchemaItems(nameSchema string) SchemaGqlItems {
	var objItems = graphql.NewObject(graphql.ObjectConfig{
		Name:        "Items",
		Fields: graphql.Fields{
			"sku": &graphql.Field{
				Type:        graphql.NewNonNull(graphql.String),
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					if item, ok := p.Source.(*response.ItemsCartDTO); ok {
						return item.Sku, nil
					}

					return nil, nil
				},
			},
			"name": &graphql.Field{
				Type:        graphql.NewNonNull(graphql.String),
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					if item, ok := p.Source.(*response.ItemsCartDTO); ok {
						return item.Name, nil
					}

					return nil, nil
				},
			},
			"qty": &graphql.Field{
				Type:        graphql.NewNonNull(graphql.Int),
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					if item, ok := p.Source.(*response.ItemsCartDTO); ok {
						return item.Qty, nil
					}

					return nil, nil
				},
			},
		},
	})

	if nameSchema == objItems.Name() {
		obj.Object = objItems
	}

	return obj
}

func (obj *schemaQuery) Result() (*graphql.Result, error) {
	var inputType = graphql.NewInputObject(
		graphql.InputObjectConfig{
			Name: "InputType",
			Fields: graphql.InputObjectConfigFieldMap{
				"sku": &graphql.InputObjectFieldConfig{
					Type: graphql.String,
				},
				"qty": &graphql.InputObjectFieldConfig{
					Type: graphql.Int,
				},
			},
		},
	)

	rootQuery := graphql.NewObject(graphql.ObjectConfig{
		Name: "RootQuery",
		Fields: graphql.Fields{
			"items": &graphql.Field{
				Type: graphql.NewList(obj.Object),
				Args: graphql.FieldConfigArgument{
					"params": &graphql.ArgumentConfig{
						Type: graphql.NewList(inputType),
					},
				},
				Resolve: checkingOut,
			},
		},
	})

	schema, _ := graphql.NewSchema(graphql.SchemaConfig{
		Query:    rootQuery,
	})

	params := graphql.Params{Schema: schema, RequestString: obj.Query}
	exec := graphql.Do(params)
	if len(exec.Errors) > 0 {
		return nil, exec.Errors[0].OriginalError()
	}

	return exec, nil
}

func (obj *schemaMutation) Result() (*graphql.Result, error) {
	mutationQuery := graphql.NewObject(graphql.ObjectConfig{
		Name: "MutationQuery",
		Fields: graphql.Fields{
			"addToCart": &graphql.Field{
				Type: obj.Object,
				Args: graphql.FieldConfigArgument{
					"sku": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.String),
					},
					"name": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.String),
					},
					"qty": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.Int),
					},
				},
				Resolve: addToChart,
			},
		},
	})

	schema, _ := graphql.NewSchema(graphql.SchemaConfig{
		Query: mutationQuery,
		Mutation:    mutationQuery,
	})

	params := graphql.Params{Schema: schema, RequestString: obj.Mutation}
	exec := graphql.Do(params)
	if len(exec.Errors) > 0 {
		return nil, exec.Errors[0].OriginalError()
	}

	return exec, nil
}

