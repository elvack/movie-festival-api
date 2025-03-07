package response

type (
	page struct {
		Current int   `json:"current"`
		Size    int   `json:"size"`
		Total   int64 `json:"total"`
	}

	res struct {
		Data    any    `json:"data"`
		Message string `json:"message"`
	}

	resWithPage struct {
		Data    any    `json:"data"`
		Message string `json:"message"`
		Page    page   `json:"page"`
	}
)
