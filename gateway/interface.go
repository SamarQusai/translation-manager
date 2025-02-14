package gateway

type Interface interface {
	Translate(texts []string, targetLanguage string) ([]string, error)
}
