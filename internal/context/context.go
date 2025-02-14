package context

type ServiceContext interface {
	Configuration() *Configuration
}
type defaultServiceContext struct {
	configuration *Configuration
}

func NewTranslationServiceContext() ServiceContext {
	configuration := loadConfiguration()
	ctx := &defaultServiceContext{
		configuration: configuration,
	}
	return ctx
}

func (ctx *defaultServiceContext) Configuration() *Configuration {
	return ctx.configuration
}
