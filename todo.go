package project

type TodoList struct {
	Id          int    `json:"id"`
	Tite        string `json:"title"`
	Description string `json:"description"`
}

type UserList struct {
	Id     int
	UserId int
	ListId int
}

type TodoItem struct {
	Id          int    `json:"id"`
	Titel       string `json:"title"`
	Description string `json:"description"`
	Done        bool   `json:"done"`
}

type ListItem struct {
	Id     int
	ListId int
	ItemId int
}
