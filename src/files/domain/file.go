package fileDomain

type File struct {
	Path string `json:"path"`
	Name string `json:"name"`
	Metadata any `json:"metadata"`
}

func New() File {
	return File{}
}