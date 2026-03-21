# Projeto 03 — Motor de Regras de Desconto

## Objetivo do projeto
Implementar um módulo de cálculo de descontos para pedidos, praticando design testável, regras de negócio e benchmarks.

## O que este projeto exercita
- regras de negócio isoladas;
- testes orientados a tabela;
- benchmarks de cenários relevantes;
- separação entre cálculo, validação e contexto promocional;
- refatoração segura.

## Escopo inicial sugerido
- desconto por tipo de cliente;
- desconto por quantidade;
- cupons promocionais;
- cobertura de cenários principais e de borda.

## Como desenvolver
1. Liste os cenários de negócio em formato tabular.
2. Implemente primeiro uma função pura de cálculo.
3. Adicione variações de regra progressivamente.
4. Inclua benchmarks após estabilizar o comportamento.

## Critérios de aceite
- regras centrais possuem testes claros;
- cupons e limites possuem validação explícita;
- benchmarks ajudam a entender custo de execução;
- o módulo pode ser reutilizado em outro contexto.

## Próximo passo
Depois deste projeto, avance para `../04-product-catalog-api/`.
