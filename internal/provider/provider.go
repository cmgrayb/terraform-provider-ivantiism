package provider

import (
	"context"
	"time"

	"github.com/hashicorp/go-hclog"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func init() {
	// Set descriptions to support markdown syntax, this will be used in document generation
	// and the language server.
	schema.DescriptionKind = schema.StringMarkdown

	// Customize the content of descriptions when output. For example you can add defaults on
	// to the exported descriptions if present.
	// schema.SchemaDescriptionBuilder = func(s *schema.Schema) string {
	// 	desc := s.Description
	// 	if s.Default != nil {
	// 		desc += fmt.Sprintf(" Defaults to `%v`.", s.Default)
	// 	}
	// 	return strings.TrimSpace(desc)
	// }
}

func New(version string) func() *schema.Provider {
	return func() *schema.Provider {
		p := &schema.Provider{
			DataSourcesMap: map[string]*schema.Resource{
				"ism_configuration_item_datasource": dataSourceConfigurationItem(),
			},
			ResourcesMap: map[string]*schema.Resource{
				"ism_configuration_item_resource": resourceConfigurationItem(),
			},
		}
	Schema:
		map[string]*schema.Schema{
			"username": {
				Type:        schema.TypeString,
				Required:    true,
				DefaultFunc: schema.EnvDefaultFunc("ISM_USERNAME", nil),
				Description: "Username passed to the Ivanti ISM API.  Can also be set through environment variable ISM_USERNAME",
			},
			"password": {
				Type:        schema.TypeString,
				Required:    true,
				DefaultFunc: schema.EnvDefaultFunc("ISM_PASSWORD", nil),
				Sensitive:   true,
				Description: "Password passed to the Ivanti ISM API.  Can also be set through environment variable ISM_PASSWORD",
			},
			"userrole": {
				Type:        schema.TypeString,
				Required:    true,
				DefaultFunc: schema.EnvDefaultFunc("ISM_USERROLE", nil),
				Description: "User role assigned to username in Ivanti Service Manager.  Can also be set through environment variable ISM_USERROLE",
			},
			"tenant": {
				Type:        schema.TypeString,
				Required:    true,
				DefaultFunc: schema.EnvDefaultFunc("ISM_TENANT", nil),
				Description: "Ivanti Service Manager tenant internal ID.  Can also be set through environment variable ISM_TENANT",
			},
		}

		p.ConfigureContextFunc = configure(version, p)

		return p
	}
}

type apiClient struct {
	// Add whatever fields, client or connection info, etc. here
	// you would need to setup to communicate with the upstream
	// API.
	BaseUrl   string
	Timeout   time.Duration
	UserAgent string
}

func configure(version string, p *schema.Provider) func(context.Context, *schema.ResourceData) (interface{}, diag.Diagnostics) {
	return func(context.Context, *schema.ResourceData) (interface{}, diag.Diagnostics) {
		log := hclog.Default()
		log.Trace("[TRACE] Configuring Ivanti Service Manager provider connection")
		// Setup a User-Agent for your API client (replace the provider name for yours):
		// userAgent := p.UserAgent("terraform-provider-ivantiism", version)
		// TODO: myClient.UserAgent = userAgent

		return &apiClient{}, nil
	}
}
