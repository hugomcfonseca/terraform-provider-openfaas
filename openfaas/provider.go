package openfaas

import (
	"bytes"
	"fmt"
	"log"

	"github.com/hashicorp/terraform/helper/hashcode"
	"github.com/hashicorp/terraform/helper/mutexkv"
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/hashicorp/terraform/terraform"
	homedir "github.com/mitchellh/go-homedir"
)

// Provider returns a terraform.ResourceProvider.
func Provider() terraform.ResourceProvider {

	// The actual provider
	return &schema.Provider{
		Schema: map[string]*schema.Schema{
			"access_key": {
				Type:        schema.TypeString,
				Optional:    true,
				Default:     "",
				Description: "OpenFaaS access key",
			},

			"secret_key": {
				Type:        schema.TypeString,
				Optional:    true,
				Default:     "",
				Description: "OpenFaaS secret key",
			},

		},

		DataSourcesMap: map[string]*schema.Resource{
			// "openfaas_function":                    dataSourceOpenFaaSFunction(),
		},

		ResourcesMap: map[string]*schema.Resource{
			// "openfaas_function":                    resourceOpenFaaSFunction(),

		},
		ConfigureFunc: providerConfigure,
	}
}


func providerConfigure(d *schema.ResourceData) (interface{}, error) {
	config := Config{

	}

	return config.Client()
}