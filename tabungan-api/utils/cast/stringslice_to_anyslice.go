package cast

func StringsliceToAnyslice(stringslice []string) []interface{} {
	anyslice := make([]interface{}, len(stringslice))
	for i, v := range stringslice {
		anyslice[i] = v
	}
	return anyslice
}
