package session

type State struct {
	Form Form
}

type Form struct {
	Username string
	Error    string
}
