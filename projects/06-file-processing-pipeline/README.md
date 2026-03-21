# Projeto 06 — Pipeline Concorrente de Arquivos

## Objetivo do projeto
Construir um pipeline concorrente para leitura, parsing e agregação de arquivos, praticando goroutines, channels, cancelamento e métricas simples.

## O que este projeto exercita
- pipeline por stages;
- coordenação com goroutines;
- channels buffered e unbuffered;
- cancelamento por erro ou timeout;
- verificação com race detector.

## Escopo inicial sugerido
- leitura paralela de arquivos;
- parsing em uma segunda etapa;
- agregação em uma etapa final;
- cancelamento em caso de erro crítico.

## Como desenvolver
1. Implemente primeiro uma versão sequencial.
2. Separe o fluxo em stages claros.
3. Introduza canais e sincronização progressivamente.
4. Compare desempenho e comportamento com e sem concorrência.

## Critérios de aceite
- o pipeline processa múltiplos arquivos corretamente;
- o cancelamento funciona sem vazamento de goroutines;
- `go test -race` passa quando houver testes;
- a solução concorrente é justificável e explicável.

## Próximo passo
Depois deste projeto, avance para `../07-payment-integration/`.
