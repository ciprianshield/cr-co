package mysql
 
import (
     "database/sql"
    _"github.com/go-sql-driver/mysql"
)

func Insert_user_table(email, password string){

	db, err := sql.Open("mysql", "adminc:adminc@tcp(127.0.0.1:3306)/test_db")
	
    // if there is an error opening the connection, handle it
    if err != nil {
        panic(err.Error())
	}

	insert, err := db.Query("INSERT INTO users VALUES ( null,"+email+","+password+")")

    if err != nil {
        panic(err.Error())
    }
	defer db.Close()
    defer insert.Close()
}

func Edit_user_password(email, oldpass, newpass string){

	db, err := sql.Open("mysql", "adminc:adminc@tcp(127.0.0.1:3306)/test_db")
	
    // if there is an error opening the connection, handle it
    if err != nil {
        panic(err.Error())
	}

	update, err := db.Query("UPDATE users SET pass = "+newpass+" WHERE pass = "+ oldpass+" AND email = "+ email)

	if err != nil{
		panic(err.Error())
	}
	defer db.Close()
	defer update.Close();
}

func Is_user_registered(email, password string) bool{

	db, err := sql.Open("mysql", "adminc:adminc@tcp(127.0.0.1:3306)/test_db")
	
    // if there is an error opening the connection, handle it
    if err != nil {
        panic(err.Error())
	}

	var exists bool = false
	err = db.QueryRow("SELECT exists (SELECT * FROM users WHERE email like '"+email+"' and pass like '"+password+"')").Scan(&exists)
	if err != nil{
		panic(err.Error())
	}
	defer db.Close()
	return exists
}