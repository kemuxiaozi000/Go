package mock

// Retriver ...x
type Retriver struct {
	Contents string
}

// Get ...x
func (r Retriver) Get(url string) string {
	return r.Contents
}
