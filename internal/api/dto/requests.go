package dto

type LogIn struct {
	Login string `json:"login"`
	Pass  string `json:"pass"`
}

type UserRegictrate struct {
	Name     string `json:"name"`
	Login    string `json:"login"`
	Password string `json:"pass"`
}

type ServiceReq struct {
	UserId   int    `json:"user_id"`
	ServName string `json:"name"`
}

type ServiceAdd struct {
	Name        string `json:"name"`
	Description string `json:"descript"`
	Login       string `json:"login"`
	Pass        string `json:"pass"`
	UserID      int    `json:"user_id"`
}
