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
	Records []Properties
}

type GetDocBody struct {
	Search_type string
	From        int
	Size        int
}

type GetDocProps struct {
	Hits HitsProps
}
type HitsProps struct {
	Hits []HitsData
}

type HitsData struct {
	Index  string `json:"_index"`
	Source Data   `json:"_source"`
}

type Data struct {
	Body    string
	Date    string
	Folder  string
	From    string
	Subject string
	To      []string
}
