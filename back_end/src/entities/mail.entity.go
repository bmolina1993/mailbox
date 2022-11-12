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

type HitsData struct {
	Id string `json:"_id"`
}

type HitsProps struct {
	Hits []HitsData
}

type GetDocProps struct {
	Hits HitsProps
}

type GetDocBody struct {
	Search_type string
	From        int
	Max_results int
}
