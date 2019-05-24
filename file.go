package s420

// File describes a file saved on your storage system
type File struct {
	Path      string
	Name      string
	Size      int64
	MD5       string
	Type      ContentType
	Thumbnail []byte
}
