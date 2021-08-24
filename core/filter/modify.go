package filter

func init() {
	Register("Modify", NewModify)
}

//replace, convert, round
func NewModify(data map[string]interface{}, renders []FilterConfig) (map[string]interface{}, bool) {
	for _, render := range renders {
		if render.status(data) {
			if row, ok := render.value(data); ok {
				data[render.Filed] = row
			} else {
				return data, false
			}
		}
	}
	return data, true
}
