//show Login OMIT
func (l *loginService) Login(userName, password string) {
	if l.userRepository.IsValid(userName, password) {
		redirect("homepage")
	} else {
		addFlash("error", "Bad credentials")
		redirect("login")
	}
}
//end show Login OMIT

//show Login2 OMIT
func (l *loginService) Login(userName, password string) {
	if l.userRepository.IsValid(userName, password) {
		redirect("homepage")
		return
	}
	addFlash("error", "Bad credentials")
	redirect("login")
}
//end show Login2 OMIT