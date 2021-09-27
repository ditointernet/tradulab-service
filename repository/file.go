package repository

type File struct {
	ID        string `gorm:"primaryKey"`
	ProjectID string
	FilePath  string
}
