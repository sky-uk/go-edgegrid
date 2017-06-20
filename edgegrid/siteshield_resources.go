package edgegrid

// DomainSummary is a representation of the Akamai GTM
// domain summary associated with each domain returned by
// the domains response at:
// http://apibase.com/config-gtm/v1/domains
type MapSummary struct {
	Map []interface{} `json:"siteShieldMaps"`
}
