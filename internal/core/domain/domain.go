package domain

type File struct {
	ID        string
	ProjectID string
	FilePath  string `json:",omitempty"`
	FileName  string `json:",omitempty"`
	Status    string
}

type Phrase struct {
	ID      string
	FileID  string
	Key     string
	Content string
}

type Suggestion struct {
	PhraseID string
	UserID   string
	Language string
	Approved bool
}

type Direction int

const UPVOTE Direction = 1
const DOWNVOTE Direction = -1

type Vote struct {
	SuggestionID string
	UserID       string
	Direction    Direction
}
