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
	FileID  string `gorm:"uniqueIndex:compositeindex"`
	Key     string `gorm:"uniqueIndex:compositeindex"`
	Content string
}

type Config struct {
	DataBaseUrl string `cfg:"database_url" cfgRequired:"true"`
	Page        int    `cfg:"page" cfgRequired:"true"`
}
