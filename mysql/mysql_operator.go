package mysql
 
import (
     "database/sql"
    _"github.com/go-sql-driver/mysql"
	"fmt"
)

const USER, PASS string = "adminc", "adminc"
const IP, PORT string = "127.0.0.1", "3306"

func Insert_user_table(email, password,age,gender string){

	db, err := sql.Open("mysql", USER+":"+PASS+"@tcp("+IP+":"+PORT+")/test_db")
	
    // if there is an error opening the connection, handle it
    if err != nil {
        panic(err.Error())
	}
	if Is_user_registered(email,password){
		fmt.Println("User already registered!")
		return
	}
	query := "INSERT INTO users VALUES( null, '"+email+"','"+password+"','"+age+"','"+gender+"');"
	fmt.Println(query)
	insert, err := db.Query(query)

    if err != nil {
        panic(err.Error())
    }
	defer db.Close()
    defer insert.Close()
}

func Edit_user_password(email, oldpass, newpass string){

	db, err := sql.Open("mysql", USER+":"+PASS+"@tcp("+IP+":"+PORT+")/test_db")
	
    // if there is an error opening the connection, handle it
    if err != nil {
        panic(err.Error())
	}
	//to do
	//query := "UPDATE users SET pass = '"+newpass+"' WHERE + " '
	update, err := db.Query("UPDATE users SET pass = "+newpass+" WHERE pass = "+ oldpass+" AND email = "+ email)

	if err != nil{
		panic(err.Error())
	}
	defer db.Close()
	defer update.Close();
}

func Is_user_registered(email, password string) bool{

	db, err := sql.Open("mysql", USER+":"+PASS+"@tcp("+IP+":"+PORT+")/test_db")
	
    // if there is an error opening the connection, handle it
    if err != nil {
        panic(err.Error())
	}

	var exists bool = false
	
	//to do query
	err = db.QueryRow("SELECT exists (SELECT * FROM users WHERE email like '"+email+"' and pass like '"+password+"')").Scan(&exists)
	if err != nil{
		panic(err.Error())
	}
	defer db.Close()
	return exists
}

func Get_user_data(input_email string) (string ,string, string){
	db, err := sql.Open("mysql", USER+":"+PASS+"@tcp("+IP+":"+PORT+")/test_db")
	
    // if there is an error opening the connection, handle it
    if err != nil {
        panic(err.Error())
	}

	var (
		email string
		age string
		gender string
	)
	//fmt.Println("query email is***"+input_email)
	sqlStatement := "SELECT email, age, gender FROM users WHERE email like '" +input_email+"';"
	row := db.QueryRow(sqlStatement);

	switch err := row.Scan(&email,&age,&gender); err{
	case sql.ErrNoRows:
		fmt.Println("No rows returned!")
	case nil:
		return email, age, gender
	default:
		panic(err)
	}
	return input_email,age,gender
}