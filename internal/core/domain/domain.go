package domain

type File struct {
	Id        string
	ProjectId string
	FilePath  string `json:",omitempty"`
	FileName  string `json:",omitempty"`
	Status    string
}

type Phrase struct {
	Id      string
	FileId  string
	Key     string
	Content string
}

type Suggestion struct {
	PhraseId string
	UserId   string
	Language string
	Approved bool
}

type Direction int

const UPVOTE Direction = 1
const DOWNVOTE Direction = -1

type Vote struct {
	SuggestionId string
	UserId       string
	Direction    Direction
}
