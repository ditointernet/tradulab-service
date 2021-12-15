package drivers

type Phrase struct {
	Id  string
	Key string
}

type File struct {
	ID        string `json:"id"`
	ProjectID string `json:"project_id"  binding:"required"`
	FilePath  string `json:"file_path"`
	FileName  string `json:"file_name"  binding:"required"`
	Status    string `json:"status"`
}
