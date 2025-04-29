package response

type (
	LoginResponse struct {
		Token     string `json:"token"`
		ExpiresIn int    `json:"expires_in"` // seconds
	}
)
