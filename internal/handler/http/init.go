package http

import (
	"translationManager/internal/context"
	"translationManager/internal/service"
)

type Handler struct {
	translatorService service.Service
	serviceContext    context.ServiceContext
}

func NewHttpRequestHandler() *Handler {
	serviceContext := context.NewTranslationServiceContext()
	translationService := service.NewTranslationService(serviceContext)
	return &Handler{
		serviceContext:    serviceContext,
		translatorService: translationService,
	}
}
