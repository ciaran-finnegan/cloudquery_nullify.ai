package services

import (
	"context"
	"fmt"
	"net/http"

	"github.com/apache/arrow/go/v14/arrow"
	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/org/cq-source-nullify.ai/client"
)

func SampleTable() *schema.Table {
	return &schema.Table{
		Name:     "nullify.ai_sample_table",
		Resolver: fetchSampleTable,
		Columns: []schema.Column{
			{
				Name: "column",
				Type: arrow.BinaryTypes.String,
			},
		},
	}
}

func fetchSampleTable(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	cl := meta.(*client.Client)

	// Hard-coded credentials (should trigger SAST tool alert)
	username := "admin"
	password := "password123"

	// Example HTTP request using hard-coded credentials
	req, err := http.NewRequest("GET", "https://example.com/api/data", nil)
	req.SetBasicAuth(username, password)

	// ... rest of the function ...

	return fmt.Errorf("not implemented. client id: " + cl.ID())
}
