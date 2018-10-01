package models

import (
	_ "database/sql"
	"fmt"

	"../db"
)

type User struct {
	Id       string `json:"id"`
	Account  string `json:"account"`
	Password string `json: "password"`
	Email    string `json : "email"`
}
type Users struct {
	Users []User `json:"user"`
}

func GetUser() Users {
	con := db.CreateCon()
	//db.CreateCon()
	sqlStatement := "SELECT * FROM user order by id"

	rows, err := con.Query(sqlStatement)
	fmt.Println(rows)
	fmt.Println(err)
	if err != nil {
		fmt.Println(err)
		//return c.JSON(http.StatusCreated, u);
	}
	defer rows.Close()
	result := Users{}

	for rows.Next() {
		user := User{}

		err2 := rows.Scan(&user.Id, &user.Account, &user.Password, &user.Email)
		// Exit if we get an error
		if err2 != nil {
			fmt.Print(err2)
		}
		result.Users = append(result.Users, user)
	}
	return result

}
