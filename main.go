package main

//go:generate ./gen.sh

import (
	"github.com/hashicorp/terraform/plugin"
	"github.com/hashicorp/terraform/terraform"

	"github.com/styk-tv/terraform-provider-contrail/contrail"
)

func main() {
	plugin.Serve(&plugin.ServeOpts{
		ProviderFunc: func() terraform.ResourceProvider {
			return contrail.Provider()
		},
	})
}
