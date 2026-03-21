# Testes e Benchmarks

## Objetivo da etapa
Esta etapa introduz o ciclo de qualidade em Go, mostrando como validar comportamento com testes automatizados e como medir desempenho com benchmarks simples.

## O que você vai aprender
- como escrever testes com o pacote `testing`;
- como estruturar testes orientados a tabela;
- como usar subtests para cenários relacionados;
- como criar benchmarks para comparar implementações;
- como validar cobertura e revisar comportamento antes de refatorar.

## Estrutura da pasta
- `README.md` — guia de estudo da etapa.
- futuros exemplos podem incluir `_test.go`, benchmarks e pequenas funções de domínio para validação.

## Como estudar esta etapa
1. Escolha uma função pequena e escreva primeiro os casos de teste antes da refatoração.
2. Organize entradas e saídas esperadas em tabelas para reduzir duplicação.
3. Adicione benchmarks apenas depois de entender claramente o comportamento funcional.
4. Compare resultados antes e depois de mudanças para desenvolver uma cultura de medição.

## Como executar
Quando exemplos forem adicionados nesta etapa, use os comandos:

```bash
go test ./...
```

```bash
go test -bench=.
```

## O que observar ao estudar
- testes bons documentam comportamento, não apenas implementação;
- benchmarks sem contexto podem induzir conclusões erradas;
- pequenas funções puras são mais fáceis de testar;
- qualidade em Go depende mais de clareza e disciplina do que de frameworks pesados.

## Exercícios sugeridos
- transformar uma regra simples em table-driven tests;
- escrever subtests para cenários válidos, inválidos e de borda;
- benchmarkar duas abordagens de manipulação de slice ou string;
- revisar testes para melhorar nomes, legibilidade e intenção.

## Próximo passo
Depois desta etapa, avance para `7-http-apis/` para aplicar testes e modelagem em serviços web reais.
