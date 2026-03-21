# Valores e Ponteiros

## Objetivo da etapa
Esta etapa existe para explicar como Go trabalha com passagem por valor, ponteiros e mutação de estado, evitando confusões comuns para quem vem de linguagens orientadas a objetos.

## O que você vai aprender
- por que Go passa valores por cópia;
- quando ponteiros são úteis;
- como diferenciar cópia de dados e compartilhamento de acesso ao mesmo valor;
- como structs se comportam quando copiadas ou referenciadas;
- como raciocinar sobre efeitos colaterais com mais clareza.

## Estrutura da pasta
- `pointers.go` — demonstra value semantics, pointer semantics e diferença entre copiar structs e apontá-las.

## Como estudar esta etapa
1. Leia os exemplos tentando prever o resultado antes de executar.
2. Compare cenários em que o dado é copiado com cenários em que o endereço é compartilhado.
3. Reescreva mentalmente cada caso perguntando: “quem pode alterar este estado?”.
4. Relacione os exemplos com seu uso futuro em métodos, structs e APIs.

## Como executar
```bash
go run .
```

## O que observar ao rodar
- quando uma alteração afeta apenas a variável atual;
- quando uma alteração reflete em mais de um ponto do código;
- como structs se comportam como tipos de valor;
- por que ponteiros devem ser usados com intenção clara, e não por padrão.

## Exercícios sugeridos
- criar um exemplo novo com uma struct `User` sendo passada por valor e por ponteiro;
- implementar uma função que atualize um campo de uma struct e comparar as duas abordagens;
- adicionar um pequeno teste manual impresso mostrando o estado antes e depois de cada chamada;
- escrever um resumo explicando a diferença entre referência em Java e ponteiro em Go.

## Próximo passo
Depois desta etapa, avance para `4-aggregate-data-types/` para combinar arrays, slices, maps e structs em estruturas mais realistas.
