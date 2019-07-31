package auth

type User struct{
	Username string
	Nickname string
	Email    string
	Password string
	ImgURI   string
}

type TokenAuth struct{
	User   User
	Device string
	Token  string
}