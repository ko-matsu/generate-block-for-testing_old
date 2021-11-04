package main

import (
	"github.com/ko-matsu/generate-block-for-testing/internal/domain/service"
	"github.com/ko-matsu/generate-block-for-testing/internal/handler"
	"github.com/ko-matsu/generate-block-for-testing/internal/infrastructure/repository"
	"github.com/ko-matsu/generate-block-for-testing/internal/usecase"
)

func NewHandler(
	argObj *argument,
) handler.Handler {
	blockchainConfig := repository.NewBlockchainRpcConfig(
		argObj.Host, argObj.RpcUserID, argObj.RpcPassword)
	blockchainRepo := repository.NewBlockchainRpc(blockchainConfig)

	genBlockService := service.NewGenerateBlock(blockchainRepo)
	genBlockUsecase := usecase.NewGenerateBlock(genBlockService)

	return handler.NewHandler(genBlockUsecase)
}
