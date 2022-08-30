package storage

type FileInfo struct {
	FileName     string `json:"filename"`
	FileFullPath string `json:"filefullpath"`
	Size         string `json:"size"`
	ModTime      string `json:"modtime"`
	Type         string `json:"type"`
}
