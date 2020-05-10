package definitions

type ClassDefinition struct {
	ClassName        string
	Template         string
	WorkingDirectory string
	Location         string
	Property         []Property
}
