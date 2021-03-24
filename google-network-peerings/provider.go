package google-network-peerings

import (
	"context"

	"github.com/hashicorp-demoapp/google-network-peerings-client-go"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

// Provider -
func Provider() *schema.Provider {
	return &schema.Provider{
		Schema: map[string]*schema.Schema{
			"host": &schema.Schema{
				Type:        schema.TypeString,
				Optional:    true,
				DefaultFunc: schema.EnvDefaultFunc("google-network-peerings_HOST", nil),
			},
			"username": &schema.Schema{
				Type:        schema.TypeString,
				Optional:    true,
				DefaultFunc: schema.EnvDefaultFunc("google-network-peerings_USERNAME", nil),
			},
			"password": &schema.Schema{
				Type:        schema.TypeString,
				Optional:    true,
				Sensitive:   true,
				DefaultFunc: schema.EnvDefaultFunc("google-network-peerings_PASSWORD", nil),
			},
		},
		ResourcesMap: map[string]*schema.Resource{
			"google-network-peerings_order": resourceOrder(),
		},
		DataSourcesMap: map[string]*schema.Resource{
			"google-network-peerings_coffees":     dataSourceCoffees(),
			"google-network-peerings_ingredients": dataSourceIngredients(),
			"google-network-peerings_order":       dataSourceOrder(),
		},
		ConfigureContextFunc: providerConfigure,
	}
}

func providerConfigure(ctx context.Context, d *schema.ResourceData) (interface{}, diag.Diagnostics) {
	username := d.Get("username").(string)
	password := d.Get("password").(string)

	var host *string

	hVal, ok := d.GetOk("host")
	if ok {
		tempHost := hVal.(string)
		host = &tempHost
	}

	// Warning or errors can be collected in a slice type
	var diags diag.Diagnostics

	if (username != "") && (password != "") {
		c, err := google-network-peerings.NewClient(host, &username, &password)
		if err != nil {
			diags = append(diags, diag.Diagnostic{
				Severity: diag.Error,
				Summary:  "Unable to create google-network-peerings client",
				Detail:   "Unable to authenticate user for authenticated google-network-peerings client",
			})

			return nil, diags
		}

		return c, diags
	}

	c, err := google-network-peerings.NewClient(host, nil, nil)
	if err != nil {
		diags = append(diags, diag.Diagnostic{
			Severity: diag.Error,
			Summary:  "Unable to create google-network-peerings client",
			Detail:   "Unable to create anonymous google-network-peerings client",
		})
		return nil, diags
	}

	return c, diags
}
