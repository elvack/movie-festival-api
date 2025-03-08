package movie

type (
	CreateReqBody struct {
		Description string
		Duration    int16
		Title       string `binding:"required"`
		WatchUrl    string `json:"watch_url"`
	}

	ListReqQuery struct {
		Limit  int `form:"limit"`
		Offset int
		Page   int `form:"page"`
	}

	UpdateReq struct {
		Body struct {
			Description string
			Duration    int16
			Title       string `binding:"required"`
			WatchUrl    string `json:"watch_url"`
		}
		Path struct {
			Id uint32 `binding:"required" uri:"id"`
		}
	}
)
