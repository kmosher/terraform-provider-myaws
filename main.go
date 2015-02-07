package main

import (
	"github.com/hashicorp/terraform/plugin"
	"github.com/kmosher/terraform-provider-myaws/myaws"
)

func main() {
    plugin.Serve(&plugin.ServeOpts{
		ProviderFunc: myaws.Provider,
	})
}
