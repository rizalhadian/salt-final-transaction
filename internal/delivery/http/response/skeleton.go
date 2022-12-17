package http_response

type SkeletonSingleResponse struct {
	// Response_code string                 `json:"response_code"`
	Success bool   `json:"success"`
	Message string `json:"message"`
	Data    any    `json:"data"` //Pakai Mapstring Interface Agar Bisa Menerima Berbagai Macam Struct Yang Di Convert Menjadi Map Inteface
}

type SkeletonPaginateResponse struct {
	// Response_code string                 `json:"response_code"`
	Success    bool                   `json:"success"`
	Message    string                 `json:"message"`
	Data       map[string]interface{} `json:"data"`
	Total_data int32                  `json:"total_data"`
	Total_page int32                  `json:"total_page"`
	Limit      int32                  `json:"limit"`
}

type SkeletonError struct {
	// Response_code string                 `json:"response_code"`
	Success bool   `json:"success"`
	Message string `json:"message"`
	Errors  string `json:"errors,omitempty"`
	// Errors  []byte `json:"errors,omitempty"`
	// Errors string `json:"errors,omitempty"`
}
