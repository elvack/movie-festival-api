package response

type (
	page struct {
		Current int   `json:"current"`
		Size    int   `json:"size"`
		Total   int64 `json:"total"`
	}

	response struct {
		Data    any    `json:"data"`
		Message string `json:"message"`
	}

	responseWithPage struct {
		response
		Page    page   `json:"page"`
	}
)
