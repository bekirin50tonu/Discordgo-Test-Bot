package helpers

func Contains(list []string, include string) bool {
	for _, v := range list {
		if v == include {
			return true
		}
	}
	return false
}
