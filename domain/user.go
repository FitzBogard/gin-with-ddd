package domain

type User struct {
	ID         int64
	NickName   string
	Sex        string
	BornAt     int64 // born timestamp
	RegisterAt int64
	Avatar     string // avatar image url
}
