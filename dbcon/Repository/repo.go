package Repository

import (
	"fmt"
	"log"
	models "sql-db-con/Models"
	"sql-db-con/Setting"
	"sql-db-con/files"
)

func GetDetails() {
	var con = Setting.Constr()
	users := []models.User{}
	rows, err := con.Query("select * from users")

	if err != nil {
		log.Fatal(err)
	}
	var id string
	var firstname string
	var lastname string
	var gender string
	var address string
	var dob string
	var country string
	var status string
	var emailid string
	var password string

	for rows.Next() {
		err := rows.Scan(&id, &firstname, &lastname, &gender, &address, &dob, &country, &status, &emailid, &password)

		if err != nil {
			log.Fatal(err)
		}
		s := models.User{
			UserID:    id,
			FirstName: firstname,
			LastName:  lastname,
			Gender:    gender,
			Address:   address,
			DOB:       dob,
			Country:   country,
			Status:    status,
			EmailId:   emailid,
			Password:  password,
		}
		users = append(users, s)
	}
	defer con.Close()
	files.WriteJson(users)
}

func InsertNewUser(u models.User) {
	var con = Setting.Constr()
	query := fmt.Sprint("insert into users values ('", u.FirstName, "','", u.LastName, "','", u.Gender, "','", u.Address, "','", u.DOB, "','", u.Country, "','", u.Status, "','", u.EmailId, "','", u.Password, "')")
	result, err := con.Exec(query)
	if err != nil {
		fmt.Print(err.Error())
		return
	}
	defer result.LastInsertId()
}

func Deleteuser(id int) {
	var con = Setting.Constr()
	query := fmt.Sprint("delete from users where UserID = ", id, ";")
	result, err := con.Exec(query)
	if err != nil {
		fmt.Print(err.Error())
		return
	}
	defer result.RowsAffected()
}

func UpdateUser(field string, value string, id string) {
	var con = Setting.Constr()
	query := fmt.Sprint("update users set ", field, " = '", value, "' where UserID = ", id, "")
	result, err := con.Exec(query)
	if err != nil {
		fmt.Print(err.Error())
		return
	}
	defer result.RowsAffected()
}
