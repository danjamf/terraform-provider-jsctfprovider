package main

import (
	"fmt"
	"os"
	"log"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/plugin"
	"jsctfprovider/endpoints/idp"
	"jsctfprovider/endpoints/blockpages"
	"jsctfprovider/endpoints/uemc"
	"jsctfprovider/endpoints/ztna"
	"jsctfprovider/endpoints/routes"
	"jsctfprovider/endpoints/activationprofiles"
	"jsctfprovider/endpoints/categories"
	"jsctfprovider/endpoints/groups"
	"jsctfprovider/internal/auth"
	
)

// Run "go generate" to format example terraform files and generate the docs for the registry/website

// If you do not have terraform installed, you can remove the formatting command, but its suggested to
// ensure the documentation is formatted properly.
//go:generate terraform fmt -recursive ./examples/

// Run the docs generation tool, check its repository for more information on how it works and how docs
// can be customized.
//go:generate go run github.com/hashicorp/terraform-plugin-docs/cmd/tfplugindocs generate --provider-name jsc

var (
	DomainName string
	Username   string
	Password   string
	Customerid string
)

func main() {

	// Perform authentication later
	//err := authenticate()
	//if err != nil {
	//	panic(fmt.Sprintf("failed to authenticate: %v", err))
	//} do not auth here - we need to get credentials first
	log.SetFlags(log.Flags() &^ (log.Ldate | log.Ltime))


	// Create a new plugin with a specific provider
	plugin.Serve(&plugin.ServeOpts{
		ProviderFunc: func() *schema.Provider {
			return &schema.Provider{
				Schema: map[string]*schema.Schema{
					"domain_name": {
						Type:        schema.TypeString,
						Optional:    true,
						Default:     "radar.wandera.com",
						Description: "The JSC domain.",
					},
					"username": {
						Type:        schema.TypeString,
						Required:    true,
						Description: "The JSC username used for authentication.",
					},
					"password": {
						Type:        schema.TypeString,
						Required:    true,
						Sensitive:   true,
						Description: "The JSC password used for authentication.",
					},
					"customerid": {
						Type:        schema.TypeString,
						Optional:    true,
						Default:     "empty",
						Description: "The optional customerID. If not provided, the provider will attempt to discover.",
					},
				},
				// Define the resources that this provider manages
				ResourcesMap: map[string]*schema.Resource{
					"jsc_oktaidp":   idp.ResourceOktaIdp(),
					"jsc_uemc":      uemc.ResourceUEMC(),
					"jsc_blockpage": blockpages.ResourceBlockPage(),
					"jsc_ztna": 	 ztna.Resourceztna(),
					"jsc_ap":		 activationprofiles.ResourceActivationProfile(),

				},
				// Define the datasources
				DataSourcesMap: map[string]*schema.Resource{
					"jsc_routes":   	routes.DataSourceRoutes(),
					"jsc_categories":   categories.DataSourceCategories(),
					"jsc_groups":       groups.DataSourceGroups(),
				},
				ConfigureFunc: providerConfigure,
			}

		},
	})

}

func providerConfigure(d *schema.ResourceData) (interface{}, error) {
	// Read the domain_name field from the configuration and assign it to domainName
	DomainName = d.Get("domain_name").(string)

	// Assign username and password from configuration
	Username = d.Get("username").(string)
	Password = d.Get("password").(string)
	Customerid = d.Get("customerid").(string)
	auth.Authenticate(DomainName, Username, Password, Customerid)


	return nil, nil
}

// GetClientPassword retrieves the 'password' value from the Terraform configuration.
// If it's not present in the configuration, it attempts to fetch it from the JSC_PASSWORD environment variable.
func GetClientPassword(d *schema.ResourceData) (string, error) {
	password := d.Get("password").(string)
	if password == "" {
		password = os.Getenv("JSC_PASSWORD")
		if password == "" {
			return "", fmt.Errorf("password must be provided either as an environment variable (JSC_PASSWORD) or in the Terraform configuration")
		}
	}
	return password, nil
}
