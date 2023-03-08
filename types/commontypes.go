package types

type ProductVersion struct {
	Node    string
	Vue     string
	Python  string
	Django  string
	React   string
	Angular string
}

type Version interface {
	GetVersion() ProductVersion
}
