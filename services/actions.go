package services

import (
	"context"
)

type ActionFunc func(ctx context.Context) error

var PublisherAction = map[string]ActionFunc{
	"single": publishSingleMessage,
	"batch":  publishBatchMessage,
}

var ConsumerAction = map[string]ActionFunc{
	"single": consumeSingleMessage,
	"batch":  consumeBatchMesasge,
}
