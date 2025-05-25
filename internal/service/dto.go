package service

type FilterHasAccountPayload struct {
	UserID   string `query:"user_id"`
	IsActive *bool  `query:"is_active"`
}

type FilterHasAccountResponse struct {
	ID          string `json:"id"`
	ServiceCode string `json:"service_code"`
	ServiceName string `json:"service_name"`
}
