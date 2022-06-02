package files

import (
	"encoding/json"
	"io/ioutil"
	"log"
	models "sql-db-con/Models"
)

func WriteJson(users []models.User) {
	json_data2, err := json.Marshal(users) //convert to json format
	if err != nil {
		log.Fatal(err)
	}
	_ = ioutil.WriteFile("test.json", json_data2, 0644) //create json file and insert the data
}

func ReadJson() []models.User {
	///Json get Data ..........,.,.,.,.,.
	details := []models.User{}
	file, _ := ioutil.ReadFile("test.json")
	_ = json.Unmarshal([]byte(file), &details)
	return details
}
