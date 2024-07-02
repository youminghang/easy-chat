package logic

type User struct {
	Id    string
	Name  string
	Phone string
}

var users = map[string]*User{
	"1": {
		Id:    "1",
		Name:  "keji",
		Phone: "12131212312",
	},
	"2": {
		Id:    "2",
		Name:  "ryfuee",
		Phone: "124124123412",
	},
}
