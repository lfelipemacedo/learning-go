# Observabilidade e Produção

## Objetivo da etapa
Esta etapa conecta o estudo da linguagem com preocupações reais de operação: logging, métricas, shutdown gracioso, configuração e comportamento previsível em produção.

## O que você vai aprender
- como estruturar logs úteis e consistentes;
- como expor métricas e sinais de saúde da aplicação;
- como aplicar timeouts e shutdown gracioso;
- como pensar em resiliência, idempotência e retry com cuidado;
- como aproximar seus estudos de cenários reais de produção.

## Estrutura da pasta
- `README.md` — guia de estudo da etapa.
- futuros exemplos podem incluir serviços instrumentados, endpoints de health check e fluxos com integração externa.

## Como estudar esta etapa
1. Releia aplicações anteriores pensando em operação, não só em funcionalidade.
2. Adicione logs, configuração e pontos de observabilidade progressivamente.
3. Simule falhas para entender como o sistema se comporta em cenários degradados.
4. Revise se os sinais produzidos pela aplicação realmente ajudam a investigar problemas.

## Como executar
Quando exemplos forem adicionados nesta etapa, rode:

```bash
go run .
```

## O que observar ao estudar
- aplicações em produção precisam ser observáveis, não apenas corretas no caminho feliz;
- logs demais sem contexto atrapalham tanto quanto logs de menos;
- timeout, retry e idempotência exigem decisões conscientes;
- shutdown gracioso e health checks ajudam a operar serviços com segurança.

## Exercícios sugeridos
- adicionar request ID e logging estruturado em uma API simples;
- criar endpoints de health e readiness;
- simular timeout em integração externa e revisar a resposta do sistema;
- documentar quais métricas seriam mais úteis para acompanhar o comportamento da aplicação.

## Próximo passo
Depois desta etapa, o caminho natural é transformar o estudo em projetos completos de portfólio usando os módulos anteriores em conjunto.
