package main

import (
	//"bytes"
	"flag"
	"fmt"
	"io/ioutil"
	"net/http"
)

var apiKey string

func main() {
	flag.StringVar(&apiKey, "key", "", "GW2 API Key")
	flag.Parse()

	url := "https://api.guildwars2.com/v2/account"
	var bearer = "Bearer " + apiKey

	request, err := http.NewRequest("POST", url, nil)
	//resp, err := http.Get(url)
	if err != nil {
		fmt.Println(err)
	}

	request.Header.Add("Authorization", bearer)
	client := &http.Client{}

	resp, err := client.Do(request)
	if err != nil {
		fmt.Println(err)
	}

	defer resp.Body.Close()

	fmt.Println("response Status:", resp.Status)
	fmt.Println("response Headers:", resp.Header)
	body, _ := ioutil.ReadAll(resp.Body)
	fmt.Println("response Body: ", string(body))

}
