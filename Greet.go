package hello

var Icon string = "👜"

func Esperanto(name string) string {
	return "Saluton " + name + " " + Icon
}

func English(name string) string {
	return "Hello " + name + " " + Icon
}

func Spanish(name string) string {
	return "Hola " + name + " " + Icon
}
