package demo

type (
	DemoTitle struct {
		Title string
	}

	DemoType struct {
		Msg string
	}

	ResponseType struct {
		*DemoType
	}
)

func (d *DemoType) Title() string {
	return d.Msg
}
