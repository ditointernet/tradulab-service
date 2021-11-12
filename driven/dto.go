package driven

type FileStatus string

const CREATED FileStatus = "CREATED"
const SUCCESS FileStatus = "SUCCESS"

type File struct {
	ID        string `gorm:"primaryKey"`
	ProjectID string
	Status    FileStatus
}

type Phrase struct {
	ID      string `gorm:"primaryKey"`
	FileID  string
	Key     string
	Content string
}
