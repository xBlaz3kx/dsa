package rate_limiting

import "github.com/google/uuid"

type Request interface {
	GetID() string
	GetContent() string
}

type SimpleRequest struct {
	ID      string
	Content string
}

func NewRequest(content string) SimpleRequest {
	return SimpleRequest{
		ID:      uuid.New().String(),
		Content: content,
	}
}

func (r SimpleRequest) GetID() string {
	return r.ID
}

func (r SimpleRequest) GetContent() string {
	return r.Content
}
