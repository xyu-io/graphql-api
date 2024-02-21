package demo_query

import (
	"graphql-api/module/demo"
)

// demo Query type
type QueryDemoType struct{}

// demo Query resolve
func (t *QueryDemoType) GetTitle() (*demo.DemoType, error) {

	var title = "welcome to use graphql api with demo app."
	return &demo.DemoType{
		Msg: title,
	}, nil
}
