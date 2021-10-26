package driven

type FileStatus string

const CREATED FileStatus = "CREATED"

type File struct {
	ID        string `gorm:"primaryKey"`
	ProjectID string
	Status    FileStatus
}
