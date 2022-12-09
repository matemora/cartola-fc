package event

import (
	"context"
	"encoding/json"

	"github.com/confluentinc/confluent-kafka-go/kafka"
	"github.com/matemora/cartola-fc/internal/domain/entity"
	"github.com/matemora/cartola-fc/internal/usecase"
	"github.com/matemora/cartola-fc/pkg/uow"
)

type ProcessNewAction struct{}

func (p ProcessNewAction) Process(ctx context.Context, msg *kafka.Message, uow uow.UowInterface) error {
	var input usecase.ActionAddInput
	err := json.Unmarshal(msg.Value, &input)
	if err != nil {
		return err
	}
	actionTable := entity.ActionTable{}
	actionTable.Init()
	addNewActionUsecase := usecase.NewActionAddUseCase(uow, &actionTable)
	err = addNewActionUsecase.Execute(ctx, input)
	if err != nil {
		return err
	}
	return nil
}
