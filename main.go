package main

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/plugin"

	"github.com/josep-pla-jt/terraform-provider-redshift"
)

//go:generate go run github.com/hashicorp/terraform-plugin-docs/cmd/tfplugindocs

func main() {
	plugin.Serve(&plugin.ServeOpts{
		Debug:        debug,
		ProviderAddr: "jt.dev/tf/redshifttf",
		ProviderFunc: func() *schema.Provider {
			return redshift.Provider()
		},
	})
}
