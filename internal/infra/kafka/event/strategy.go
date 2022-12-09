package event

import (
	"context"

	"github.com/confluentinc/confluent-kafka-go/kafka"
	"github.com/matemora/cartola-fc/pkg/uow"
)

type ProcessEventStrategy interface {
	Process(ctx context.Context, msg *kafka.Message, uow uow.UowInterface) error
}
