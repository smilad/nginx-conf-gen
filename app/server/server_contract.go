package server

import "context"

type IServer interface {
	SetUpServer(container DeliveryContainer)
	Shutdown(ctx context.Context) error
	Run() error
}
