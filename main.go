package main

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"regexp"
	"strconv"
	"strings"

	//"strings"
	"time"
)

func main() {
	var result Response
	resp, err := http.Get("https://github.com/AvineshTripathi.atom")
	if err != nil {
		log.Fatalln(err)
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}

	if err := xml.Unmarshal(body, &result); err != nil {
		fmt.Println("cannot unmarshal JSON")
	}

	file := checkAndCreateFile()

	defer file.Close()

	var data Metadata
	data.GetTemplate(&result)

	fmt.Println(data)
	
	fmt.Println("DONE!!!!!")

}

func (data *Metadata) editFile(file *os.File, result *Response) {
	data.GetTemplate(result)

	
}

func (data *Metadata) GetTemplate(result *Response) {
	var templates []Template
	for i, newActivity := range result.Entry {
		if i == 5 {
			break
		}
		var s = strings.Split(newActivity.ID,"/")[1]
		var id, err = strconv.Atoi(s)
		 if err != nil {
			fmt.Println("Error: %v", err)
		 }
		event := GetEvent(newActivity.ID)
		templates = append(templates, Template{id, event, newActivity.Title.Text, newActivity.Link.Href, newActivity.Updated }) 
	}

	data.TweetDate = "today"
	data.CurrentTopId = templates[0].ID
	data.Templates = templates
}


func GetEvent(s string) string {
	var rgx = regexp.MustCompile(`:(.*?)/`)
	rs := rgx.FindStringSubmatch(s)
	return strings.Split(rs[1], ":")[1]
}

func checkAndCreateFile() *os.File {
	file, err := os.Create("tweets" + time.Now().Format("2006-02-01") + ".txt")
	if err != nil {
		log.Fatalln(err)
	}

	return file
}

func PrettyPrint(i interface{}) string {
	s, _ := json.MarshalIndent(i, "", "\t")
	return string(s)
}


