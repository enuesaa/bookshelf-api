package plug

type ProviderInterface interface {
	Info() Info
	ListUiCards() []UiCard
	List(rid string) []Resource
	Describe(rid string) Resource
	// Create(resource Resource) error
	// Update(resource Resource) error
	// Delete(resource Resource) error
}

// rid stands for resource id like arn. example: `notes:<id>` or `notes:*` or `note:main`
type UiCard struct {
	Layout string
	Rid string // 
	CanList bool 
	CanCreate bool // ここが ok なら create form が表示されるみたいな
	CanUpdate bool
	CanDelete bool
}

type Info struct {
	Name string
	Description string
}

type Resource struct {
	Name string
	Values map[string]Value
}

type Value struct {
	Type string // string or markdown, int, bool
	StringValue string
	MarkdownValue string
	IntValue int
	BoolValue string // checkbox
	LinkValue string // like `tags:<id>`
	Readonly bool
}
