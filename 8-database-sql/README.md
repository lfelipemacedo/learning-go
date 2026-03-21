# Banco de Dados com SQL

## Objetivo da etapa
Esta etapa apresenta persistência relacional em Go com foco em SQL explícito, transações e integração com a biblioteca `database/sql`.

## O que você vai aprender
- como conectar uma aplicação Go a um banco relacional;
- como executar queries, scans e transações;
- como modelar repositórios sem depender de abstrações pesadas;
- como lidar com falhas de persistência de forma previsível;
- como pensar em desempenho e clareza ao escrever SQL.

## Estrutura da pasta
- `README.md` — guia de estudo da etapa.
- futuros exemplos podem incluir repositórios, migrations e fluxos transacionais.

## Como estudar esta etapa
1. Modele primeiro o domínio e as tabelas antes de escrever queries.
2. Comece com operações simples de leitura e escrita.
3. Em seguida, introduza transações em fluxos que exigem consistência.
4. Revise se a modelagem do banco e a modelagem Go permanecem compreensíveis juntas.

## Como executar
Quando exemplos forem adicionados nesta etapa, siga as instruções específicas de ambiente e então rode:

```bash
go run .
```

## O que observar ao estudar
- SQL explícito dá mais previsibilidade do que ORMs muito mágicos;
- transações devem ser usadas com intenção clara;
- erros de banco precisam carregar contexto suficiente;
- integração com banco em Go costuma premiar simplicidade e organização.

## Exercícios sugeridos
- escrever uma operação simples de inserção e leitura com `database/sql`;
- modelar um fluxo transacional de criação de pedido com itens;
- comparar duas queries e discutir legibilidade versus performance;
- adicionar um teste de integração para um caso de persistência.

## Próximo passo
Depois desta etapa, avance para `9-concurrency/` para explorar paralelismo, coordenação e cancelamento em Go.
