<script>
		function logout()
		{
		  sessionStorage.clear();
		}
		
		function checkIfLoggedIn(){
			if(sessionStorage.getItem('myUserEntity') == null){
			//Redirect to login page, no user entity available in sessionStorage
			//location.href='index.html';
			} else {
			//User already logged in
				var userEntity = {};
				userEntity = JSON.parse(sessionStorage.getItem('myUserEntity'));
				var profile = googleUser.getBasicProfile();
				console.log("ID Token: " + id_token);
				document.getElementById("signup_email").value = profile.getEmail();
				document.getElementById("signup_password").value = "Choose a password";
				document.getElementById("signup_password_confirm").value = "Repeat the chosen password";
				document.getElementById("Age").value = "How old are you?";
				document.getElementById("Gender").value = "-";
				location.href="profile_page.html";
			}
		}	
</script>