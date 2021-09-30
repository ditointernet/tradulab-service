package driven

type File struct {
	ID        string `gorm:"primaryKey"`
	ProjectID string
	FilePath  string
}
