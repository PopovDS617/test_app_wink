package service

import "context"

type LBService interface {
	GetURL(ctx context.Context, inputURL string) (string, error)
}
