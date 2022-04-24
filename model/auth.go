package model

type Auth struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type Token struct {
	Token string `json:"token"`
	Type  string `json:"type"`
}

type TokenExtraction struct {
	Username string
	ExpTime  int64
}
