package main

import(
	"net/http"
	engine "./engine"
	)


func main(){

	http.HandleFunc("/", engine.Index)
	http.HandleFunc("/login_process", engine.Login_process)
	http.HandleFunc("/signup", engine.Signup)
	http.HandleFunc("/profile_page", engine.Profile_page)
	http.HandleFunc("/logout", engine.Logout)
	http.HandleFunc("/edit_profile", engine.Edit_profile)
	http.HandleFunc("/save_edit", engine.Save_edit)
	//http.HandleFunc("/show_profile", engine.Show_profile)
	
	http.ListenAndServe(":8080",nil)
}








