package havau

import (
	"github.com/hashicorp/vault/api"
	"github.com/pkg/errors"
)

// Havau defines main struct for the project
type Havau struct {
	client *api.Client
}

// New provides initialization of the package
func New(c *api.Config, token string) (*Havau, error) {
	client, err := api.NewClient(c)
	if err != nil {
		return nil, errors.Wrap(err, "unable to initialize Vault client")
	}
	client.SetToken(token)
	return &Havau{
		client: client,
	}, nil
}
