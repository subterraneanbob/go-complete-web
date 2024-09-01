package dict

type Finder interface {
	Find(prefix string, limit int) []string
}
