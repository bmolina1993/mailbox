package entities

type Properties struct {
	Folder  string
	Date    string
	From    string
	To      []string
	Subject string
	Body    string
}

type QueryMail struct {
	Index   string
	Content []Properties
}
