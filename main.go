package main

import (
	"fmt"
	"os"

	"github.com/theMillenniumFalcon/pexels/utils"
)

func main() {
	os.Setenv("PEXELS_TOKEN", "563492ad6f917000010000012e400882eb1c444aa5b97c57ad270189")
	var TOKEN = os.Getenv("PEXELS_TOKEN")
	var c = utils.NewClient(TOKEN)

	// result, err := c.SearchPhotos("waves", 15, 1)
	// result, err := c.SearchVideo("waves", 15, 1)
	// result, err := c.GetRandomPhoto()
	result, err := c.GetRandomVideo()
	if err != nil {
		fmt.Printf("Search ertror:%v", err)
	}

	// if result.Page == 0 {
	// 	fmt.Printf("search result wrong")
	// }
	fmt.Println(result)
}
