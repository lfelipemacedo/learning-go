# Interfaces e Tratamento de Erros

## Objetivo da etapa
Esta etapa aprofunda dois pilares do Go idiomático: o uso de interfaces pequenas e o tratamento explícito de erros como parte do fluxo normal da aplicação.

## O que você vai aprender
- quando faz sentido criar interfaces em Go;
- por que interfaces costumam nascer no consumidor, e não na implementação;
- como modelar erros de domínio e erros de infraestrutura;
- como usar `errors.Is`, `errors.As` e `fmt.Errorf` com `%w`;
- como evitar abstrações excessivas vindas de outras linguagens.

## Estrutura da pasta
- `README.md` — guia de estudo da etapa.
- futuros exemplos desta pasta podem incluir contratos simples, serviços com composição e cenários de propagação de erro.

## Como estudar esta etapa
1. Revise mentalmente onde você criaria uma interface em Java e questione se ela realmente é necessária em Go.
2. Pratique começando pela implementação concreta e só extraia interfaces quando houver um consumidor real.
3. Modele erros com contexto suficiente para debugging, sem esconder a origem da falha.
4. Releia código anterior do repositório e identifique pontos em que interfaces ou erros poderiam ser melhorados.

## Como executar
Quando exemplos de código forem adicionados nesta etapa, entre na pasta desejada e rode:

```bash
go run .
```

## O que observar ao estudar
- interfaces pequenas tendem a ser mais reutilizáveis e legíveis;
- erros em Go fazem parte da modelagem e não devem ser tratados como exceções invisíveis;
- composição costuma produzir código mais simples do que hierarquias artificiais;
- mensagens de erro boas ajudam tanto no desenvolvimento quanto na operação.

## Exercícios sugeridos
- criar uma interface mínima para leitura de dados baseada em um consumidor real;
- implementar um fluxo com erro encapsulado usando `fmt.Errorf("...: %w", err)`;
- escrever uma função que distinga dois tipos de falha usando `errors.Is` ou `errors.As`;
- revisar um exemplo anterior do repositório e simplificar abstrações desnecessárias.

## Próximo passo
Depois desta etapa, avance para `6-testing-and-benchmarks/` para transformar essas abstrações e regras em código validado por testes.
