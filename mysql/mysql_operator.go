package mysql
 
import (
     "database/sql"
    _"github.com/go-sql-driver/mysql"
	"fmt"
)

const USER, PASS string = "adminc", "adminc"
const IP, PORT string = "127.0.0.1", "3306"

func DbConn() (db *sql.DB) {

    db, err := sql.Open("mysql", USER+":"+PASS+"@tcp("+IP+":"+PORT+")/test_db")
    if err != nil {
        panic(err.Error())
    }
    return db
}

func Insert_user_table(email, password,age,gender string){

	db := DbConn()
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

	db := DbConn()
	update, err := db.Query("UPDATE users SET pass = "+newpass+" WHERE pass = "+ oldpass+" AND email = "+ email)

	if err != nil{
		panic(err.Error())
	}
	defer db.Close()
	defer update.Close();
}

func Is_user_registered(email, password string) bool{

	db := DbConn()

	var exists bool = false
	
	//to do query
	err := db.QueryRow("SELECT exists (SELECT * FROM users WHERE email like '"+email+"' and pass like '"+password+"')").Scan(&exists)
	if err != nil{
		panic(err.Error())
	}
	defer db.Close()
	return exists
}

func Get_user_data(input_email string) (string ,string, string){

	db := DbConn()
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

func Update_user_table(email, password, age, gender string){

	db := DbConn()
	sqlStatement := "UPDATE users WHERE SET age="+age+", gender= "+gender+" WHERE email like '"+email+"' and pass like '"+password+"';"
    // if there is an error opening the connection, handle it

	update, err := db.Query(sqlStatement)

	if err != nil{
		panic(err.Error())
	}
	defer db.Close()
	defer update.Close();
}
