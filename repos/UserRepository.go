package repos
 // temporary userbase test package
 //remove after implementing db
func UserIsValid(uName, pwd string) bool {
    // DB simulation
    _uName, _pwd, _isValid := "cristian", "1234!*.", false
 
    if uName == _uName && pwd == _pwd {
        _isValid = true
    } else {
        _isValid = false
    }
 
    return _isValid
}