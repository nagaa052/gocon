package gocon

// ImgFormat is a convertible image format.
type ImgFormat string

const (
	JPEG ImgFormat = "jpeg"
	PNG  ImgFormat = "png"
)

// String is converted to string type and returned.
func (i *ImgFormat) String() string {
	return string(*i)
}

// GetExtentions returns the extension of the target format.
func (i *ImgFormat) GetExtentions() []string {
	if *i == JPEG {
		return []string{".jpeg", ".jpg"}
	}

	return []string{"." + string(*i)}
}

// Exist examines available formats.
func (i *ImgFormat) Exist() bool {
	return *i == JPEG || *i == PNG
}
