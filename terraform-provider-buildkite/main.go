package main

import (
	"github.com/hashicorp/terraform-plugin-sdk/plugin"
	"github.com/kushmansingh/terraform-buildkite/buildkite"
)

func main() {
	plugin.Serve(&plugin.ServeOpts{
		ProviderFunc: buildkite.Provider,
	})
}
