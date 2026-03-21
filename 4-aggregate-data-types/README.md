# Tipos Agregados

## Objetivo da etapa
Esta etapa aprofunda os principais tipos agregados de Go e mostra como eles se comportam em cenários de cópia, compartilhamento e modelagem de dados.

## O que você vai aprender
- a diferença entre arrays e slices;
- como maps funcionam e quais cuidados exigem;
- como structs ajudam a modelar domínio de forma simples;
- como esses tipos se combinam em aplicações reais;
- como identificar quando há cópia de valor e quando há compartilhamento de dados internos.

## Estrutura da pasta
- `aggregate-data-types.go` — imprime exemplos de arrays, slices, maps e structs.

## Como estudar esta etapa
1. Leia cada bloco separadamente: arrays, slices, maps e structs.
2. Rode o programa observando quais tipos são copiados e quais compartilham dados internamente.
3. Tente explicar por que a saída muda em cada caso.
4. Relacione esses comportamentos com futuras estruturas de domínio, coleções e APIs.

## Como executar
```bash
go run .
```

## O que observar ao rodar
- arrays são copiados por valor;
- slices compartilham o mesmo array base, salvo cópia explícita;
- maps compartilham a mesma estrutura interna e não são ordenados;
- structs são tipos de valor e ajudam a organizar dados relacionados.

## Exercícios sugeridos
- criar um exemplo adicional mostrando como `append` pode alterar ou não o backing array de um slice;
- modelar uma lista de produtos usando `[]struct` e um índice usando `map[string]int`;
- implementar funções simples de busca, filtro ou agregação sobre slices;
- combinar arrays, slices, maps e structs em um mini exemplo de catálogo ou agenda.

## Próximo passo
Depois desta etapa, o caminho natural é avançar para tópicos como interfaces, tratamento de erros, testes e composição de tipos.
