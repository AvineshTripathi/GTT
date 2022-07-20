package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"time"
)

func main() {
	var result Response
	resp, err := http.Get("https://api.github.com/users/AvineshTripathi/events")
	if err != nil {
		log.Fatalln(err)
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}

	if err := json.Unmarshal(body, &result); err != nil {
		fmt.Println("cannot unmarshal JSON")
	} 
	
	file := checkAndCreateFile()

	defer file.Close()

	_,err2 := file.WriteString(result[0].ID)
	if err != nil {
		log.Fatalln(err2)
	}

	fmt.Println("DONE!!!!!")

}


func checkAndCreateFile() *os.File {
	file, err := os.Create("tweets"+time.Now().Format("2006-02-01")+".txt")
	if err != nil {
		log.Fatalln(err)
	}

	return file 
}

func PrettyPrint(i interface{}) string {
    s, _ := json.MarshalIndent(i, "", "\t")
    return string(s)
}