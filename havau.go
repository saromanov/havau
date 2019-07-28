package havau

import (
	"github.com/hashicorp/vault/api"
	"golang.org/x/xerrors"
)

var errNoToken = xerrors.New("token is not defined")

// Havau defines main struct for the project
type Havau struct {
	client *api.Client
}

// New provides initialization of the package
func New(c *api.Config, token string) (*Havau, error) {
	if token == "" {
		return nil, errNoToken
	}
	client, err := api.NewClient(c)
	if err != nil {
		return nil, xerrors.Errorf("unable to initialize Vault client: %v", err)
	}
	client.SetToken(token)
	return &Havau{
		client: client,
	}, nil
}

// Write provides writing of the kv to the attached Vault store
func (h *Havau) Write(path string, kv map[string]interface{}) error {
	c := h.client.Logical()
	_, err := c.Write(path, kv)
	if err != nil {
		return err
	}
	return nil
}

// Read provides reading from kv store
func (h *Havau) Read(path string) (map[string]interface{}, error) {
	c := h.client.Logical()
	s, err := c.Read(path)
	if err != nil {
		return nil, err
	}
	return s.Data, nil
}
