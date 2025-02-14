package service

import (
	"translationManager/gateway"
	"translationManager/internal/context"
)

type Service struct {
	translatorGateway gateway.Interface
	serviceContext    context.ServiceContext
}

func NewTranslationService(serviceContext context.ServiceContext) Service {
	return Service{
		serviceContext:    serviceContext,
		translatorGateway: gateway.NewTranslationGateway(serviceContext.Configuration(), gateway.OpenAI),
	}
}
