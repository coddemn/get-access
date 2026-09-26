package domain

import "time"

type User struct {
	ID    int
	Name  string
	Login string
	Pass  string
}

type Service struct {
	ID          int
	Name        string
	Description string
	Login       string
	Pass        string
	UserID      int
	AddedAt     time.Time
}

type Bind struct {
	ID   int
	Name string
}

type UserBind struct {
	UserID int
	BindID int
	Addr   string
}
