# Projeto 07 — Serviço de Integração de Pagamentos

## Objetivo do projeto
Criar um serviço que processa pagamentos com integração externa simulada, reforçando resiliência, observabilidade, timeout e idempotência.

## O que este projeto exercita
- contratos HTTP entre serviços;
- integração externa com falhas controladas;
- timeout e retry;
- logs estruturados e métricas;
- shutdown gracioso.

## Escopo inicial sugerido
- endpoint de processamento de pagamento;
- provedor externo simulado;
- chave de idempotência;
- timeout por chamada externa;
- métricas básicas.

## Como desenvolver
1. Crie primeiro um stub do provedor externo.
2. Modele estados do pagamento de forma explícita.
3. Implemente idempotência antes de retry.
4. Adicione logs, métricas e shutdown gracioso.

## Critérios de aceite
- chamadas repetidas com a mesma chave não geram efeitos duplicados;
- o sistema reage de forma previsível a timeout e falhas externas;
- há visibilidade por logs e métricas;
- o serviço falha de forma controlada.

## Próximo passo
Depois deste projeto, avance para `../08-portfolio-project/`.
