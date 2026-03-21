# 🚀 Estudos em Go (Golang)

Este repositório reúne minha trilha prática de aprendizado em **Go (Golang)**, organizada para consolidar os fundamentos da linguagem e evoluir gradualmente para tópicos mais avançados.

A proposta não é apenas guardar exemplos soltos, mas construir uma base de estudo progressiva, com foco em:

- entender os conceitos centrais da linguagem;
- praticar a sintaxe e os tipos nativos do Go;
- comparar a forma de pensar em Go com abordagens mais tradicionais de orientação a objetos;
- registrar exemplos pequenos, objetivos e fáceis de executar;
- preparar terreno para projetos mais completos envolvendo APIs, banco de dados, testes e concorrência.

## Objetivo do repositório

O objetivo deste repositório é servir como um **laboratório de aprendizado contínuo**, no qual cada pasta representa uma etapa da evolução em Go. A ideia é estudar de forma incremental: começar pelos fundamentos, exercitar a linguagem com exemplos simples e, aos poucos, avançar para aplicações mais realistas.

Além de apoiar os estudos do dia a dia, este material também funciona como referência para revisões futuras, permitindo revisitar conceitos, comparar soluções e acompanhar a própria evolução na linguagem.

## Estrutura
- `1-first-examples/` — primeiros exemplos para ganhar familiaridade com a sintaxe, execução de programas e conceitos iniciais.
- `2-simple-data-types/` — tipos simples, variáveis, constantes, booleanos, números, strings e erros.
- `3-pointers-and-values/` — diferenças entre passagem por valor, ponteiros e manipulação de dados.
- `4-aggregate-data-types/` — arrays, slices, maps e structs, com foco em modelagem e uso prático.
- `5-interfaces-and-errors/` — interfaces pequenas, composição e tratamento de erros idiomático.
- `6-testing-and-benchmarks/` — testes automatizados, table-driven tests e benchmarks.
- `7-http-apis/` — APIs HTTP com `net/http`, JSON, middleware e contexto.
- `8-database-sql/` — persistência relacional com `database/sql`, queries e transações.
- `9-concurrency/` — goroutines, channels, sincronização e cancelamento.
- `10-observability-and-production/` — logs, métricas, health checks, timeouts e operação em produção.
- `projects/` — projetos práticos incrementais para consolidar os conceitos estudados ao longo da trilha.

## Como executar
Cada pasta tem seu próprio `go.mod`. Entre na pasta desejada e rode:
```bash
go run .
```

## Tópicos estudados até agora
- Tipos simples
- Valores vs. ponteiros
- Tipos agregados (arrays, slices, maps e structs)

## Próximos tópicos (planejados)
- Interfaces, composição e tratamento de erros
- Testes e benchmarks
- HTTP e APIs
- Banco de dados com SQL
- Concorrência (goroutines, channels, sync)
- Observabilidade e operação em produção

## Projetos práticos
- Os projetos incrementais ficam em [`projects/`](./projects/), do nível introdutório até um projeto avançado de portfólio.

## Guia de estudos
- Veja o plano completo em [`GUIA-DE-ESTUDOS-GO.md`](./GUIA-DE-ESTUDOS-GO.md).
