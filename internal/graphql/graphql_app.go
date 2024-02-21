package graphql

import (
	"graphql-api/module/demo/graphql/demo_mutation"
	"graphql-api/module/demo/graphql/demo_query"
)

type Resolver struct{}

func RegisterResolver() *Resolver {
	//注册
	return &Resolver{}
}

// Demo-App
type AppObjectType struct {
	*demo_query.QueryDemoType

	*demo_mutation.MutationDemoType
}

func (r *Resolver) Demo() (*AppObjectType, error) {
	return &AppObjectType{}, nil
}
