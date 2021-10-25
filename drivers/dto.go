package drivers

type Phrase struct {
	ID  string
	Key string
}

type File struct {
	ID        string `json:"id"`
	ProjectID string `json:"project_id"  binding:"required"`
	FilePath  string `json:"file_path"  binding:"required"`
	Status    string `json:"status"`
}
