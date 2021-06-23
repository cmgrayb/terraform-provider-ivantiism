package provider

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
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
				"configuration_item": dataSourceConfigurationItem(),
			},
			ResourcesMap: map[string]*schema.Resource{
				"configuration_item": resourceConfigurationItem(),
			},
			Schema: map[string]*schema.Schema{
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
					Description: "Password passed to the Ivanti ISM API.  Can also be set through environment variable ISM_PASSWORD, though not recommended",
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
	Tenant    string
	Username  string
	Password  string
	Role      string
	Timeout   time.Duration
	UserAgent string
}

func configure(version string, p *schema.Provider) func(context.Context, *schema.ResourceData) (interface{}, diag.Diagnostics) {
	return func(ctx context.Context, d *schema.ResourceData) (interface{}, diag.Diagnostics) {
		log := hclog.Default()
		log.Trace("[TRACE] Configuring Ivanti Service Manager provider connection")
		username := url.QueryEscape(d.Get("username").(string))
		password := url.QueryEscape(d.Get("password").(string))
		tenant := d.Get("tenant").(string)
		baseurl := d.Get("baseurl").(string)
		userrole := d.Get("role").(string)
		timeout, err := time.ParseDuration(d.Get("timeout").(string))

		// Setup a User-Agent for your API client (replace the provider name for yours):
		userAgent := p.UserAgent("terraform-provider-ivantiism", version)
		// TODO: myClient.UserAgent = userAgent

		config := &apiClient{
			BaseUrl:   baseurl,
			Timeout:   timeout,
			UserAgent: userAgent,
		}

		if err != nil {
			return config, diag.FromErr(err)
		}

		client := &http.Client{
			Timeout: config.Timeout,
		}

		var diags diag.Diagnostics

		url := baseurl + "/api/rest/authentication/login"
		method := "POST"
		reqbody := fmt.Sprintf(
			"{,`tenant : %s,`username : %s,`password : %s,`role : %s,`}",
			tenant, username, password, userrole,
		)

		log.Trace("Sending authentication request", "URL", url, "Body", reqbody, "Timeout", config.Timeout)
		req, err := http.NewRequest(method, url, strings.NewReader(reqbody))
		if err != nil {
			return config, diag.FromErr(err)
		}
		req.Header.Add("Content-Type", "application/json")

		resp, err := client.Do(req)
		if err != nil {
			return config, diag.FromErr(err)
		}
		defer resp.Body.Close()

		log.Trace("Received OAuth response", "StatusCode", resp.StatusCode, "ContentLength", resp.ContentLength)
		if resp.StatusCode != 200 {
			respbody, _ := io.ReadAll(resp.Body)
			return config, diag.Errorf("Oauth token response: (%d) %s", resp.StatusCode, respbody)
		}
		return config, diags
	}
}
