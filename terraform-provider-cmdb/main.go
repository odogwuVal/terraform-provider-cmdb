package main

import (
	"github.com/hashicorp/terraform-plugin-sdk/plugin"
	"github.com/odogwuVal/terraform-provider-cmdb/cmdb"
)

func main() {
	plugin.Serve(&plugin.ServeOpts{
		ProviderFunc: cmdb.Provider})
}
