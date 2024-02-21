package demo_mutation

import "graphql-api/module/demo"

// demo Mutation type
type MutationDemoType struct{}

// demo Mutation resolve
func (m *MutationDemoType) ChangeTitle(args struct{ Title string }) (*demo.DemoType, error) {

	var title = args.Title
	if len(title) == 0 {
		title = "default title."
	}
	return &demo.DemoType{
		Msg: title,
	}, nil
}
