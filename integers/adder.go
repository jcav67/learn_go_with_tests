// Package integers proporciona funciones para realizar operaciones matemáticas
// básicas como la suma. Este paquete es útil para ejemplos simples de aritmética
// en aplicaciones Go.
//
// Ejemplo de uso:
//
//	import "github.com/jcav67/learn_go_with_tests/integers"
//
//	func main() {
//	    resultado := integers.Add(2, 3)
//	    fmt.Println(resultado)
//	}
package integers

// Add takes two integers and returns the sum of them.
func Add(a, b int) int {
	return a + b
}
