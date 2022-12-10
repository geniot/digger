package api

type ImageDescriptor struct {
	ImageName string
}

var (
	SomeImages = []ImageDescriptor{
		{
			ImageName: "ball.png",
		},
	}
)
