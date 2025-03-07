package movie

type (
	CreateReqBody struct {
		Description string
		Duration    int16
		Title       string `binding:"required"`
		WatchUrl    string `json:"watch_url"`
	}
)
