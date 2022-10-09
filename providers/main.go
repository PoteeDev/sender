package providers

import "os"

type Provider struct {
	Name string
}

func NewProvider() *Provider {
	return &Provider{os.Getenv("PROVIDER")}
}

func (p *Provider) Send(recipient, text, mode string) {
	switch p.Name {
	case "telegram":
		NewTelegramMessage(recipient, text, mode).Send()
	}
}
