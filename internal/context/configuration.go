package context

import "os"

const (
	httpPort     = "HTTP_PORT"
	openAiApiKey = "OPEN_AI_API_KEY"
)

type Configuration struct {
	httpPort     string
	openAiApiKey string
}

func loadConfiguration() *Configuration {
	return &Configuration{
		httpPort:     os.Getenv(httpPort),
		openAiApiKey: os.Getenv(openAiApiKey),
	}
}

func (c *Configuration) HttpPort() string {
	return c.httpPort
}

func (c *Configuration) OpenAiApiKey() string {
	return c.openAiApiKey
}
