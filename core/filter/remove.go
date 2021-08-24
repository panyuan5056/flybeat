package filter

func init() {
	Register("Remove", NewRemove)
}

//has compare, remove field
func NewRemove(data map[string]interface{}, renders []FilterConfig) (map[string]interface{}, bool) {
	for _, render := range renders {
		if render.status(data) {
			delete(data, render.Filed)
		}
	}
	return data, true
}
