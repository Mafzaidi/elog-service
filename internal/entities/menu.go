package entities

type Menu struct {
	ID       string `json:"id"`
	Title    string `json:"title"`
	Url      string `json:"url"`
	Icon     string `json:"icon"`
	IsActive bool   `json:"is_active"`
	Group    string `json:"group"`
}
