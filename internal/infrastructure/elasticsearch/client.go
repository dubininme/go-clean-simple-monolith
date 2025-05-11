package elasticsearch

import (
	"github.com/elastic/go-elasticsearch/v8"
)

func NewClient(address string) (*elasticsearch.Client, error) {
	cfg := elasticsearch.Config{
		Addresses: []string{address},
	}
	return elasticsearch.NewClient(cfg)
}
