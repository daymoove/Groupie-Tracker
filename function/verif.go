package groupie

var data ArtistStruct

func SubApiVerif(apifile string, number string) ArtistStruct {
	path := ""
	switch apifile {
	case "Locations":
		path = "https://groupietrackers.herokuapp.com/api/locations/" + number 
	case "ConcertDates":
		path = "https://groupietrackers.herokuapp.com/api/dates/" + number
	case "Relations":
		path = "https://groupietrackers.herokuapp.com/api/relation/" + number
	}
	data = APIRequest(path, apifile)
	return data
}