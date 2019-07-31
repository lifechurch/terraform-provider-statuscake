package main

import (
	"github.com/hashicorp/terraform/plugin"
	"github.com/lifechurch/terraform-provider-statuscake/statuscake"
)

func main() {
	plugin.Serve(&plugin.ServeOpts{
		ProviderFunc: statuscake.Provider})
}
