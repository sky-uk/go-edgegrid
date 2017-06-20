package edgegrid

const siteshieldPath = "/siteshield/v1/"

func siteshieldBase(c *AuthCredentials) string {
	return concat([]string{
		c.APIHost,
		siteshieldPath,
	})
}

func mapsEndpoint(c *AuthCredentials) string {
	return concat([]string{
		siteshieldBase(c),
		"maps",
	})
}
