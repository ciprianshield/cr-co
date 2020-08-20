package main

import(
	"net/http"
	"html/template"
	"fmt"
	mysql "./mysql"
    //"github.com/gorilla/securecookie"
	helpers "./helpers"
	"github.com/gorilla/sessions"
	)

	var tpl *template.Template
	
	func init(){
		tpl = template.Must(template.ParseGlob("templates/*.html"))
	}
	var (
	current_email_logged_in string
	store = sessions.NewCookieStore([]byte("mysession"))
	)



func main(){
	http.HandleFunc("/", index)
	http.HandleFunc("/login_process", login_process)
	http.HandleFunc("/signup", signup)
	http.HandleFunc("/profile_page", profile_page)
	http.HandleFunc("/logout", Logout)
	
	http.ListenAndServe(":8080",nil)
}

func index(w http.ResponseWriter, r *http.Request){
	tpl.ExecuteTemplate(w,"index.html", nil)
}

func Logout(w http.ResponseWriter, r *http.Request){
	//session, _ := store.Get(r,"mysession")
	//session.Options.MaxAge = -1
	//session.Save(r,w)
	clearSession(w)
	http.Redirect(w, r, "/", 302)
}

func clearSession(w http.ResponseWriter){
	cookie := &http.Cookie{
	Name: "mysession",
	Value: "",
	Path: "/",
	MaxAge: -1,
	}
	http.SetCookie(w, cookie)
}

func login_process(w http.ResponseWriter, r *http.Request){
	if r.Method != "POST"{
		http.Redirect(w,r,"/", http.StatusSeeOther)
		return
	}
	
	email := r.FormValue("email")
	password := r.FormValue("password")
	fmt.Println(email+ " " + password)
	
	if mysql.Is_user_registered(email,password){
		//logged_in = true
		current_email_logged_in = email
		//fmt.Fprintln(w, "Logged in!")
		//http.Redirect(w,r,"/user_profile", http.StatusSeeOther)
		
		//
		session, _ := store.Get(r,"mysession")
		session.Values["email"] = email
		session.Save(r,w)
		http.Redirect(w,r,"/profile_page", http.StatusSeeOther)
		//
		
	} else {
		//fmt.Fprintln(w, "Register first!")
		data := map[string]interface{}{
			"err": "Invalid",
		}
		tpl, _ = template.ParseFiles("templates/index.html")
		tpl.Execute(w,data)
	}
}
func profile_page(w http.ResponseWriter, r *http.Request){
	//tpl.ExecuteTemplate(w,"profile_page.html", nil)
	
	//email, age, gender := mysql.Get_user_data(current_email_logged_in)
	session, _ := store.Get(r,"mysession")
	email := session.Values["email"]
	_ , age, gender := mysql.Get_user_data(current_email_logged_in)
	data := map[string]interface{}{
			"email": email,
			"age": "Your age: "+ age,
			"gender": "Your gender: "+ gender,
		}
	tmp, _ := template.ParseFiles("templates/profile_page.html")
	tmp.Execute(w, data)
	
}

func signup(w http.ResponseWriter, r *http.Request){
	if r.Method != "POST"{
		http.Redirect(w,r,"/", http.StatusSeeOther)
		return
	}
	r.ParseForm()
	
	email := r.FormValue("signup_email")
	password := r.FormValue("signup_password")
	confirPw := r.FormValue("signup_password_confirm")
	age := r.FormValue("signup_age")
	gender := r.FormValue("signup_gender")
	
	if password != confirPw {
		fmt.Fprintln(w, "Passwords do not match!")
		return
	}
	
	//insert in db to be implemented
	_email, _pass, _confirmPass, _age, _gender := false, false, false, false, false
	_email = !helpers.IsEmpty(email)
	_pass = !helpers.IsEmpty(password)
	_confirmPass = !helpers.IsEmpty(confirPw)
	_age = !helpers.IsEmpty(age)
	_gender = !helpers.IsEmpty(gender)
	
	if _email &&  _pass && _confirmPass && _age && _gender {
		
		mysql.Insert_user_table(email,password,age,gender)
		fmt.Fprintln(w, "Signed up successfully!")
	} else {
		fmt.Fprintln(w, "Fields can not be blank!")
	}
	http.Redirect(w,r,"/", http.StatusSeeOther)
}







