package components

func SetProps(original map[string]string, new map[string]string) map[string]string {
	for prop, value := range new {
		original[prop] = value
	}

	return original
}
