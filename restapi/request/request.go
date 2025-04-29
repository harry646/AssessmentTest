package request

type (
	LoginRequest struct {
		Username string `json:"username" binding:"required"`
		Password string `json:"password" binding:"required"`
	}

	AccountInfo struct {
		AccountID   string  `json:"accountId"`
		Balance     float64 `json:"balance"`
		Currency    string  `json:"currency"`
		AccountType string  `json:"accountType"`
	}
)
