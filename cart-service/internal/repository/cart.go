package repository

import "context"

type CartStorage interface {
	CreateCart(ctx context.Context)
}
