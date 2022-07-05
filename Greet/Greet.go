package Greet

var icon string = "👜"

func Esperanto(name string) string {
	return "Saluton " + name + " " + icon
}

func English(name string) string {
	return "Hello " + name + " " + icon
}

func Spanish(name string) string {
	return "Hola " + name + " " + icon
}
