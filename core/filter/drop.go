package filter

func init() {
	Register("Drop", NewDrop)
}

//if value drop row
func NewDrop(data map[string]interface{}, renders []FilterConfig) (map[string]interface{}, bool) {
	for _, render := range renders {
		if render.status(data) {
			//delete(data, render.Filed)
			return nil, false
		}
	}
	return data, true
}
