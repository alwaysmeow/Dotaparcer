package request

func DotabuffUrl(path string) string {
	domens := []string{
		"https://dotabuff.com",
		"https://it.dotabuff.com",
		"https://ka.dotabuff.com",
		"https://de.dotabuff.com",
		"https://fr.dotabuff.com",
	}

	url := domens[0] + path

	return url
}
