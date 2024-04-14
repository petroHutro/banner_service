package authorization

type Authorization struct {
	secretKey string
}

func Init(secretKey string) *Authorization {
	return &Authorization{
		secretKey: secretKey,
	}
}
