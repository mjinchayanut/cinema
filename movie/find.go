package movie

func FindName(imdbId string) string {
	switch imdbId {
	case "1":
		return "Avenger"
	case "2":
		return "Superman"
	}
	return "not found"
}
