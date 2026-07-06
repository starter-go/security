package permissions

// permissions.URI
// format: '[method]:[path]'
// like: 'post:/a/b/c/d'
type URI string

func (u URI) String() string {
	return string(u)
}
