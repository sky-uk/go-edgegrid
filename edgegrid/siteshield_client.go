package edgegrid

import "net/http"

// SiteShieldClient is an Akamai GTM API client.
// https://developer.akamai.com/api/luna/config-gtm/overview.html
type SiteShieldClient struct {
	Credentials *AuthCredentials
	HTTPClient  *http.Client
}

// GetCredentials takes a SiteShieldClient and returns its Credentials.
func (c SiteShieldClient) GetCredentials() *AuthCredentials {
	return c.Credentials
}

// GetHTTPClient takes a SiteShieldClient and returns its HTTPClient.
func (c SiteShieldClient) GetHTTPClient() *http.Client {
	return c.HTTPClient
}
