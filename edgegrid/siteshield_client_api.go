package edgegrid

// Domains returns the Akamai GTM DomainSummary for each domain the SiteShieldClient
// is authorized to view and modify
func (c *SiteShieldClient) Maps() (MapSummary, error) {
	maps := &MapSummary{}
	err := resourceRequest(c, "GET", mapsEndpoint(c.GetCredentials()), nil, maps)
	if err != nil {
		return MapSummary{}, err
	}
	return *maps, err
}
