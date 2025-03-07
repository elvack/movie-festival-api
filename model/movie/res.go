package movie

type ListResData struct {
	CreatedAt   string `json:"created_at"`
	Description string `json:"description"`
	Duration    int16  `json:"duration"`
	Id          uint32 `json:"id"`
	Title       string `json:"title"`
	UpdatedAt   string `json:"updated_at"`
	WatchUrl    string `json:"watch_url"`
}
