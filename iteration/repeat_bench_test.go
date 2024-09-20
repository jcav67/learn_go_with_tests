package iteration

import "testing"

// go test -bench="." para correr todos los bench
// go test -bench=BenchmarkRepeat test especifico
/*
Los benchmark ejecutan una funcion b.N veces, este numero de ciclos los decide el framework, dando como resultado
una salida parecida a la siguiente:
------------------------------------------------------------------------
goos: windows
goarch: amd64
pkg: github.com/jcav67/learn_go_with_tests/iteration
cpu: AMD Ryzen 9 5900HX with Radeon Graphics
BenchmarkRepeat-16      14561374                82.70 ns/op
PASS
ok      github.com/jcav67/testing/iteration     1.731s
-----------------------------------------------------------------------
de lo anterior podemos sacar informacion importante como:
- BenchmarkRepeat = Nombre del benchmark
- 16 Numero de CPUs utilizados en los tests
- 14561374 Numero de veces que el framework determino que debía ejecutar el test para obtener informacion precisa
- 82.70 ns/op Tiempo promedio de cada test 82 nanosegundos
*/
func BenchmarkRepeat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Repeat("a", 5)
	}
}
