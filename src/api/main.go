package main

import (
	"fmt"

	"./providers/locations_provider"
)

func main() {
	country, err := locations_provider.GetCountry("AR")

	fmt.Println(err)

	fmt.Println(country)
}
