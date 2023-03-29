package Entity

type User struct {
	Name string
	Id   string
	// 管道
	Msg chan string
}
