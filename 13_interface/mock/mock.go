package mock

// Retriver ...x
type Retriver struct {
	Contents string
}

// Get ...x
func (r *Retriver) Get(url string) string {
	return r.Contents
}

// Post ...x
func (r *Retriver) Post(url string, form map[string]string) string {
	r.Contents = form["contents"]
	return "OK"
}
