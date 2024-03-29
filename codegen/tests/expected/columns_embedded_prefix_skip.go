package output

import (
	"context"

	"github.com/cloudquery/cq-provider-sdk/provider/schema"
)

func ColumnsEmbeddedPrefixSkips() *schema.Table {
	return &schema.Table{
		Name:     "test_columns_embedded_prefix_skip",
		Resolver: fetchColumnsEmbeddedPrefixSkips,
		Columns: []schema.Column{
			{
				Name: "int_value",
				Type: schema.TypeBigInt,
			},
			{
				Name: "bool_value",
				Type: schema.TypeBool,
			},
			{
				Name:     "rename_field",
				Type:     schema.TypeBigInt,
				Resolver: schema.PathResolver("Embedded.FieldA"),
			},
		},
	}
}

// ====================================================================================================================
//                                               Table Resolver Functions
// ====================================================================================================================

func fetchColumnsEmbeddedPrefixSkips(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	panic("not implemented")
}
