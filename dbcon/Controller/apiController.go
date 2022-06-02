package Controller

import (
	"fmt"
	"net/http"
	"os"
	"sql-db-con/Models"
	"sql-db-con/Repository"
	"sql-db-con/files"
	"strconv"

	"github.com/gin-gonic/gin"
)

func NewAPIHandler(router *gin.Engine) {

	Router := router.Group("/api/")
	{
		Router.GET("/userdetails", GetDetails)
		Router.POST("/addnewuser", InsertDetails)
		Router.GET("/deleteuser/:deleteid", DeleteUser)
		Router.GET("/updateuser/:updatefield/:updatevalue/:updateID", UpdateUser)
	}
}

func GetDetails(c *gin.Context) {
	Repository.GetDetails()
	data := files.ReadJson()
	c.IndentedJSON(http.StatusOK, data)
}

func InsertDetails(g *gin.Context) {
	var param Models.User
	g.BindJSON(&param)
	fmt.Print(param.UserID)
	Repository.InsertNewUser(param)
	g.JSON(http.StatusOK, gin.H{
		"status":  http.StatusOK,
		"message": "success",
		"data":    param.UserID,
	})

}

func DeleteUser(g *gin.Context) {
	v := g.Param("deleteid")
	i, err := strconv.Atoi(v)
	if err != nil {
		fmt.Println(err)
		os.Exit(2)
	}
	Repository.Deleteuser(i)
	g.JSON(http.StatusOK, gin.H{
		"status":  http.StatusOK,
		"message": "success",
		"data":    v,
	})
}

func UpdateUser(g *gin.Context) {
	field := g.Param("updatefield")
	value := g.Param("updatevalue")
	Id := g.Param("updateID")
	Repository.UpdateUser(field, value, Id)
	g.JSON(http.StatusOK, gin.H{
		"status":  http.StatusOK,
		"message": "success",
		"data":    Id,
	})
}
