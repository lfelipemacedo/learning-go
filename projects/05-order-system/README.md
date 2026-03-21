# Projeto 05 — Sistema de Pedidos

## Objetivo do projeto
Construir um backend de pedidos com persistência em banco relacional, reforçando SQL explícito, transações e integração com domínio.

## O que este projeto exercita
- modelagem relacional;
- operações com `database/sql`;
- transações;
- leitura e escrita de dados com contexto;
- testes de integração com banco.

## Escopo inicial sugerido
- clientes;
- produtos;
- pedidos;
- itens do pedido;
- fechamento transacional do pedido.

## Como desenvolver
1. Modele tabelas e relacionamentos.
2. Implemente consultas simples.
3. Introduza fluxo transacional de criação ou fechamento.
4. Adicione testes de integração para o caminho feliz e falhas.

## Critérios de aceite
- o fluxo transacional é consistente;
- erros de persistência são tratados com clareza;
- o domínio continua compreensível apesar da infraestrutura;
- o ambiente local pode ser reproduzido com facilidade.

## Próximo passo
Depois deste projeto, avance para `../06-file-processing-pipeline/`.
