package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

type goat struct {
	grazed int
}

func (g *goat) get_page_info(url string) string {
	// Make HHTP GET request
	resp, err := http.Get(url)
	check(err)

	defer resp.Body.Close()

	// convert body to string and return
	if resp.StatusCode == http.StatusOK {
		bodyBytes, err2 := ioutil.ReadAll(resp.Body)
		check(err2)
		bodyString := string(bodyBytes)
		fmt.Println(bodyString)

		g.grazed += 1
		return bodyString
	}
	return "none"
}

func main() {
	g := goat{grazed: 0}
	g.get_page_info("http://www.doglost.co.uk/dog-search.php?status=Lost")
	g.get_page_info("http://www.doglost.co.uk/dog-search.php?status=Lost&page=2#results")

	fmt.Println("Scaped:", g.grazed)
}
