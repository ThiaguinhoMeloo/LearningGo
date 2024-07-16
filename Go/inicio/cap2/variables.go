package cap2

import "fmt"

func VariablesOperatedValues() {
	x := 10
	y := "bom dia!"

	fmt.Printf("x: %v, %T\n", x, x)
	fmt.Printf("y: %v, %T\n", y, y)
	// operador de declaração ele vai declarar pelo menos uma variavel

	x, z := 20, 30
	fmt.Printf("x: %v, %T\n", x, x)
	fmt.Printf("z: %v, %T\n", z, z)

	// := utilizado para criar novas variáveis, dentro de code black
	// = para distribuir valores a variáveis já existentes
}
