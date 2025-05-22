package menu

type FilterPayload struct {
	IsActive *bool  `query:"is_active"`
	Group    string `query:"group"`
}

type FilterResponse struct {
	Title string `json:"title"`
	Url   string `json:"url"`
	Icon  string `json:"icon"`
	Group string `json:"group"`
}
