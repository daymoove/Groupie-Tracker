package groupie

import (
	"os"
	"log"
	"bufio"
)

func SeparateString(api string) (string, string) {
	other := false
	apitype := ""
	number := ""
	for i:= 0; i < len(api); i++ {
		if api[i] == ' '  {
			other = true
		} else if other {
			number = number + string(api[i])
		} else {
			apitype = apitype + string(api[i])
		}
	}
	return number, apitype
}

func List() []string {
	var wordtab []string
	f, err := os.Open("../static/location.txt")
    if err != nil {
        log.Fatal(err)
    }
	defer f.Close()
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		wordtab = append(wordtab, scanner.Text())
    }
	return wordtab
}
	
// func AllRelation() {
// 	path := ""
// 	l := "Locations"
// 	var locationTab []string
// 	in := false
// 	for i := 1; i <= 52; i++ {
// 		path = "https://groupietrackers.herokuapp.com/api/locations/" + strconv.Itoa(i)
// 		data := APIRequest(path, l)
// 		for k := 0; k < len(data.TabLocations.Locations); k++ {
// 			if (len(locationTab) != 0) {
// 				for j := 0; j < len(locationTab); j++ {
// 					if (locationTab[j] == data.TabLocations.Locations[k]) {
// 						in = true
// 					}
// 				}
// 				if (!in) {
// 					in = false
// 				} else {
// 					locationTab = append(locationTab, data.TabLocations.Locations[k])
// 				}
// 			} else {
// 				locationTab = append(locationTab, data.TabLocations.Locations[k])
// 			}
// 		}
// 	}		
//     f, err := os.Create("../static/location.txt")

//     if err != nil {
//         log.Fatal(err)
//     }

//     defer f.Close()

// 	for i := 0; i < len(locationTab); i++ {
// 		_, err2 := f.WriteString(locationTab[i] + "\n")


// 		if err2 != nil {
// 			log.Fatal(err2)
// 		}
// 	}

		
		
		
		
		