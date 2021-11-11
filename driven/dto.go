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
	FileID  string
	Key     string `gorm:"primaryKey"`
	Content string
}
