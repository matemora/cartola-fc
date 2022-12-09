package event

import (
	"context"
	"encoding/json"

	"github.com/confluentinc/confluent-kafka-go/kafka"
	"github.com/matemora/cartola-fc/internal/usecase"
	"github.com/matemora/cartola-fc/pkg/uow"
)

type ProcessChooseTeam struct{}

func (p ProcessChooseTeam) Process(ctx context.Context, msg *kafka.Message, uow uow.UowInterface) error {
	var input usecase.MyTeamChoosePlayersInput
	err := json.Unmarshal(msg.Value, &input)
	if err != nil {
		return err
	}
	addNewMyTeamUsecase := usecase.NewMyTeamChoosePlayersUseCase(uow)
	err = addNewMyTeamUsecase.Execute(ctx, input)
	if err != nil {
		return err
	}
	return nil
}
