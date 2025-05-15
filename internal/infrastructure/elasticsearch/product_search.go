package elasticsearch

import (
	"context"
	"encoding/json"
	"fmt"
	"strings"

	"github.com/elastic/go-elasticsearch/v8"
)

type ProductSearchRepository struct {
	client *elasticsearch.Client
}

func NewProductSearchRepository(client *elasticsearch.Client) *ProductSearchRepository {
	return &ProductSearchRepository{client: client}
}

type esHit struct {
	Id string `json:"_id"`
}
type esHits struct {
	Hits struct {
		Hits []esHit `json:"hits"`
	} `json:"hits"`
}

func (r *ProductSearchRepository) Search(ctx context.Context, query string) ([]string, error) {
	res, err := r.client.Search(
		r.client.Search.WithContext(ctx),
		r.client.Search.WithIndex("products"),
		r.client.Search.WithBody(strings.NewReader(fmt.Sprintf(`{"query":{"match":{"name":"%s"}}}`, query))),
	)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()
	var hits esHits
	if err := json.NewDecoder(res.Body).Decode(&hits); err != nil {
		return nil, err
	}
	ids := make([]string, 0, len(hits.Hits.Hits))
	for _, h := range hits.Hits.Hits {
		ids = append(ids, h.Id)
	}
	return ids, nil
}
