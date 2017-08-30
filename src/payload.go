package conductor

type Payload struct {
	Site            Site
	Title           string
	MetaTitle       string
	MetaDescription string
	Errors          []string
	Page            interface{}
	User            User
	Breadcrumbs     []Breadcrumb
}

type Breadcrumb struct {
	Active bool
	Label  string
	Link   string
}
