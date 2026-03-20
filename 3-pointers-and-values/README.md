# Valores e Ponteiros

## Conceito
Em Go, **tudo é passado por valor**. A diferença é que ponteiros carregam endereços, então apontam para o mesmo dado.

### Value Semantics
```go
a := 42
b := a
```
`b` recebe uma cópia de `a`. Se `a` mudar, `b` não muda.

### Pointer Semantics
```go
a := 42
b := &a
```
`b` aponta para `a`. Se `a` mudar, `*b` reflete a mudança.

## Exemplo no código
O arquivo `pointers.go` demonstra:
- Cópia de valores (`value semantics`)
- Compartilhamento via ponteiro (`pointer semantics`)
- Diferença entre copiar structs e apontá-las

## Como executar
```bash
go run .
```
