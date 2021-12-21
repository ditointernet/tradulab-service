package driven

type FileStatus string

const CREATED FileStatus = "CREATED"
const SUCCESS FileStatus = "SUCCESS"

type File struct {
	Id        string `gorm:"primaryKey"`
	ProjectId string
	Status    FileStatus
}

type Phrase struct {
	Id      string `gorm:"primaryKey"`
	FileId  string `gorm:"uniqueIndex:compositeindex"`
	Key     string `gorm:"uniqueIndex:compositeindex"`
	Content string
}
