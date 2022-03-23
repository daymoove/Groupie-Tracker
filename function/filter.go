package groupie

import (
	"strconv"
)

func Filter(loc []string, first []string, create []string, members []string) ArtistStruct {
	f2 := 0
	c2 := 0
	var ch2 []int
	l := ""
	if len(first) != 0 {
		f := first[0]
		f2, _ = strconv.Atoi(f)
	}
	if len(create) != 0 {
		c := create[0]
		c2, _ = strconv.Atoi(c)
	}
	if len(members) != 0 {
		var ch int
		for i := 0; i < len(members); i++ {
			ch, _ = strconv.Atoi(members[i])
			ch2 = append(ch2, ch)
		}
	}
	if len(loc) != 0 {
		l = loc[0]
	}
	var data ArtistStruct
	var path string
	var onedata ArtistStruct
	for i := 1; i <= 52; i++ {
		if locVerif(i, l) && firstVerif(i, f2) && createVerif(i, c2) && membersVerif(i, ch2) {
			path = "https://groupietrackers.herokuapp.com/api/artists/" + strconv.Itoa(i)
			onedata = APIRequest(path, "Artists")
			data.TabArtist = append(data.TabArtist, onedata.TabOneArtist)
		}
	}
	data.TabList = List()
	return data
}

func locVerif(i int, loc string) bool {
	if loc == "" {
		return true
	} else {
		path := "https://groupietrackers.herokuapp.com/api/locations/" + strconv.Itoa(i)
		data := APIRequest(path, "Locations")
		for j := 0; j < len(data.TabLocations.Locations); j++ {
			if data.TabLocations.Locations[j] == loc {
				return true
			}	
		}
	}
	return false
}

func firstVerif(i int, first int) bool {
	if first == 0 || first == 1962 {
		return true
	} else {
		path := "https://groupietrackers.herokuapp.com/api/artists/" + strconv.Itoa(i)	
		data := APIRequest(path, "Artists")
		date := ""
		for j := 0; j < len(data.TabOneArtist.FirstAlbum); j++ {
			if j > 5 {
				date = date + string(data.TabOneArtist.FirstAlbum[j])
			}
		}
		if date == strconv.Itoa(first) {
			return true
		}
	}
	return false
}
func createVerif(i int, create int) bool {	
	if create == 0 || create == 1957 {
		return true
	} else {
		path := "https://groupietrackers.herokuapp.com/api/artists/" + strconv.Itoa(i)
		data := APIRequest(path, "Artists")
		if data.TabOneArtist.CreationDate == create {
			return true
		}
	}
	return false
}

func membersVerif(i int, members []int) bool {
	if len(members) == 0 {
		return true
	} else {
		path := "https://groupietrackers.herokuapp.com/api/artists/" + strconv.Itoa(i)
		data := APIRequest(path, "Artists")
		for i := 0; i < len(members); i++ {
			if len(data.TabOneArtist.Member) == members[i] {
				return true
			}
		}
	}
	return false
}