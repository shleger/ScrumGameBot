package domain

//test task struct
type Task struct {
	Description string
}

//Token store
type Props struct {
	Key string
	Val string
}

//echo resp
type EchoResponce struct {
	ID      uint32  `json:"update_id"`
	Message Message `json:"message"`
}

//telegram message
type Message struct {
	ID   uint32 `json:"message_id"`
	From User   `json:"from"`
	Date uint32 `json:"date"`
	Chat Chat   `json:"chat"`
	Text string `json:"text"`
}

//teleram User struct
type User struct {
	ID        uint32 `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	UserName  string `json:"user_name"`
	LangCode  string `json:"language_code"`
}

//telegrm chat struct
type Chat struct {
	ID       uint32 `json:"id"`
	Type     string `json:"type"`
	Title    string `json:"title"`
	LastName string `json:"last_name"`
	UserName string `json:"user_name"`
}
