package wordnik

type WordOfTheDay struct {
	Id          int64
	ParentId    string
	Category    string
	CreatedBy   string
	CreatedAt   string
	HtmlExtra   string
	Word        string
	Definitions []SimpleDefinition
	Examples    []SimpleExample
	Note        string
	PublishDate string
}

type ContentProvider struct {
	Id   int64
	Name string
}

type SimpleDefinition struct {
	Test         string
	Source       string
	Note         string
	PartOfSpeech string
}

type SimpleExample struct {
	Id    int64
	Title string
	Text  string
	Url   string
}
