package edgegrid

import (
	"fmt"
)

const gtmPath = "/config-gtm/v1/"
const gtmDiagnostics = "/diagnostic-tools/v2/gtm/"

func gtmBase(c *AuthCredentials) string {
	return concat([]string{
		c.APIHost,
		gtmPath,
	})
}

func gtmDiagnosticsBase(c *AuthCredentials) string {
	return concat([]string{
		c.APIHost,
		gtmDiagnostics,
	})
}

func domainsEndpoint(c *AuthCredentials) string {
	return concat([]string{
		gtmBase(c),
		"domains",
	})
}

func propertiesDiagnosticEndpoint(c *AuthCredentials) string {
	return concat([]string{
		gtmDiagnosticsBase(c),
		"gtm-properties",
	})
}

func propertiesDiagnosticIpsEndpoint(c *AuthCredentials, property, domain string) string {
	return concat([]string{
		gtmDiagnosticsBase(c),
		property, "/",
		domain, "/gtm-property-ips",
	})
}

func domainEndpoint(c *AuthCredentials, domain string) string {
	return concat([]string{
		domainsEndpoint(c),
		"/",
		domain,
	})
}

func domainStatusEndpoint(c *AuthCredentials, domain string) string {
	return concat([]string{
		domainsEndpoint(c),
		"/",
		domain,
		"/status/current",
	})
}

func dcEndpoint(c *AuthCredentials, domain string, id int) string {
	return fmt.Sprintf("%s/%d", dcsEndpoint(c, domain), id)
}

func dcsEndpoint(c *AuthCredentials, domain string) string {
	return concat([]string{
		domainEndpoint(c, domain),
		"/datacenters",
	})
}

func propertiesEndpoint(c *AuthCredentials, domain string) string {
	return concat([]string{
		domainEndpoint(c, domain),
		"/properties",
	})
}

func propertyEndpoint(c *AuthCredentials, domain, property string) string {
	return concat([]string{
		propertiesEndpoint(c, domain),
		"/",
		property,
	})
}

