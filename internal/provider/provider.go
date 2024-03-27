package provider

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/function"
	"github.com/hashicorp/terraform-plugin-framework/provider"
	"github.com/hashicorp/terraform-plugin-framework/resource"
)

var _ provider.Provider = (*sandboxProvider)(nil)
var _ provider.ProviderWithFunctions = (*sandboxProvider)(nil)

func New() func() provider.Provider {
	return func() provider.Provider {
		return &sandboxProvider{}
	}
}

func NewProvider() provider.Provider {
	return &sandboxProvider{}
}

type sandboxProvider struct{}

func (p *sandboxProvider) Schema(ctx context.Context, req provider.SchemaRequest, resp *provider.SchemaResponse) {

}

func (p *sandboxProvider) Configure(ctx context.Context, req provider.ConfigureRequest, resp *provider.ConfigureResponse) {

}

func (p *sandboxProvider) Metadata(ctx context.Context, req provider.MetadataRequest, resp *provider.MetadataResponse) {
	resp.TypeName = "examplecloud"
}

func (p *sandboxProvider) DataSources(ctx context.Context) []func() datasource.DataSource {
	return []func() datasource.DataSource{}
}

func (p *sandboxProvider) Resources(ctx context.Context) []func() resource.Resource {
	return []func() resource.Resource{
		NewThingResource,
	}
}

func (p *sandboxProvider) Functions(ctx context.Context) []func() function.Function {
	return []func() function.Function{}
}
