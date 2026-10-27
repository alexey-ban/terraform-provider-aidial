package main

import (
	"context"

	aidial "github.com/alexey-ban/terraform-provider-aidial/internal/provider"
	"github.com/hashicorp/terraform-plugin-framework/providerserver"
)

func main() {
	providerserver.Serve(
		context.Background(),
		aidial.New,
		providerserver.ServeOpts{Address: "registry.terraform.io/alexey-ban/aidial"},
	)
}
