package model

type Auth struct {
	Username string `json:"user_name"`
	Password string `json:"password"`
}

type Token struct {
	Token string `json:"token"`
	Type  string `json:"type"`
}
