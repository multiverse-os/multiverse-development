package configuration

type Key struct {
	Id string
	Name string
}

type Value string







type Configuration struct {
	path string


	Values map[Key]Value
}





