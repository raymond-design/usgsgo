package water

// The client object for USGS Instantaneous Values Web Services
type Ivs struct {
}

// Initialize the client object
func InitIvsClient() *Ivs {
	return newIvsClient()
}

func newIvsClient() *Ivs {
	return &Ivs{}
}
