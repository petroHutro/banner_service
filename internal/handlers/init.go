package handlers

type Handler struct {
	// tokenSecretKey string
	// tokenEXP       time.Duration
}

func Init() *Handler {
	return &Handler{}
}
