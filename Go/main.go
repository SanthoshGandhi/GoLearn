package main

import (
	"database/sql"
	"fmt"
	"net/http"

	"log"
	//"fmt"
	//"errors"

	"github.com/gin-gonic/gin"
)

type User struct {
	userID    string `json:"userid"`
	firstName string `json:"firstname"`
	lastName  string `json:"lastname"`
	Gender    string `json:"gender"`
	Address   string `json:"address"`
	DOB       string `json:"dob"`
	Country   string `json:"country"`
	Status    string `json:"status"`
	emailId   string `json:"emailid"`
	Password  string `json:"password"`
}

var users = []string{}

//methods

func main() {
	router := gin.Default()
	router.GET("/albums", getAlbums)

	router.Run("localhost:8080")
}

func getAlbums(c *gin.Context) {
	GetUsers()
	c.IndentedJSON(http.StatusOK, users)
}

func GetUsers() {
	condb, errdb := sql.Open("mssql", "server=172.104.170.123;database=test;user id=psiber-sql-admin;password=ZhJhE6C&cx87W7@!;")

	if errdb != nil {

		fmt.Println(" Error open db:", errdb.Error())

	}

	var (
		firstName string
	)

	rows, err := condb.Query("select firstName from users")

	if err != nil {

		log.Fatal(err)

	}
	count := 0
	for rows.Next() {

		err := rows.Scan(&firstName)

		if err != nil {

			log.Fatal(err)

		}

		users[count] = firstName

		count++
	}

	defer condb.Close()
}
