package entity

type Hello struct {
	Message	string
}

func NewHello(
	message	string,
) *Hello {
	return &Hello{
		Message:	message,
	}
}
