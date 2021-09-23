package domain

type File struct {
	ID        string
	ProjectID string
	FilePath  string
}

type Phrase struct {
	FileID string
	Key    string
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
