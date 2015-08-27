package degreespkg

import (
	"degreespkg/structs"
	//"encoding/json"
	"fmt"
	//"io/ioutil"
	//"net/http"
)

func FindMoreDegree(links int, MovieObj1 structs.Movies, MovieObj2 structs.Movies) {
	fmt.Println("\n\tThere is no possible 2 degree links between them...\n")
	fmt.Println("\n\tPlease be patience while searching more possibilities...\n")

	// for i := range MovieObj1.Cast {
	// 	var Actor1Obj structs.ActorMoviesList
	// 	var Actor2Obj structs.ActorMoviesList
	// 	for j := range MovieObj2.Cast {
	// 		data := FetchDetails(MovieObj1.Cast[i].Url)
	// 		err := json.Unmarshal(data, &Actor1Obj)
	// 		if err != nil {
	// 			fmt.Println(err.Error())
	// 			return
	// 		}

	// 	}
	// }

	fmt.Println("\n\t\t\tHave found", links, "links between them")
	fmt.Println("\n\t##########################################################")
}
