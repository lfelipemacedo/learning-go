# Tipos Agregados

## Objetivo
Revisar os principais tipos agregados de Go e como eles se comportam em cópia/compartilhamento.

## Arrays
```go
arr := [3]int{1, 2, 3}
arr2 := arr            // arrays são copiados por valor
arr[0] = 99
fmt.Println(arr)       // [99 2 3]
fmt.Println(arr2)      // [1 2 3]
```

## Slices
```go
s := []string{"foo", "bar", "baz"}
s2 := s                // slices compartilham o mesmo array base
s[0] = "qux"
fmt.Println(s, s2)     // [qux bar baz] [qux bar baz]
```
Para copiar slices, use `slices.Clone` (Go 1.20+).

## Maps
```go
m := map[string]int{"foo": 1, "bar": 2}
m2 := m                 // maps compartilham a mesma estrutura interna
m["bar"] = 99
delete(m, "foo")
fmt.Println(m, m2)      // map[bar:99] map[bar:99]

v, ok := m["foo"]       // 0, false
fmt.Println(v, ok)
```
- Maps não são ordenados e não são comparáveis.

## Structs
```go
type person struct {
	name string
	id   int
}

p1 := person{name: "Arthur", id: 42}
p2 := p1               // structs são copiados por valor
p2.name = "Tricia"
fmt.Println(p1, p2)
```
- Structs são tipos de valor, não de referência.

## Exemplo no código
O arquivo `aggregate-data-types.go` imprime exemplos de arrays, slices, maps e structs.

## Como executar
```bash
go run .
```
