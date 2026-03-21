# Concorrência

## Objetivo da etapa
Esta etapa desenvolve o raciocínio necessário para usar goroutines, channels e primitivas de sincronização com segurança e propósito.

## O que você vai aprender
- como lançar goroutines de forma consciente;
- quando usar channels e quando usar mutexes;
- como coordenar múltiplas tarefas com `sync.WaitGroup` e `select`;
- como cancelar processamento com `context.Context`;
- como evitar vazamentos, deadlocks e race conditions.

## Estrutura da pasta
- `README.md` — guia de estudo da etapa.
- futuros exemplos podem incluir worker pools, pipelines e processamento concorrente de arquivos ou tarefas.

## Como estudar esta etapa
1. Comece resolvendo o problema de forma sequencial para entender o baseline.
2. Depois introduza concorrência apenas onde houver ganho claro ou necessidade de modelagem.
3. Observe quem cria, consome e fecha cada canal.
4. Valide sempre se a solução concorrente continua simples de explicar.

## Como executar
Quando exemplos forem adicionados nesta etapa, use:

```bash
go run .
```

E, quando houver testes concorrentes:

```bash
go test -race ./...
```

## O que observar ao estudar
- concorrência útil melhora throughput ou organização do fluxo, não apenas “mostra recurso da linguagem”;
- channels não substituem toda forma de sincronização;
- cancelamento precisa ser parte do desenho, não remendo posterior;
- simplicidade continua sendo um critério importante mesmo em código paralelo.

## Exercícios sugeridos
- implementar um worker pool simples;
- montar um pipeline com duas ou três stages;
- simular cancelamento por timeout em uma operação mais lenta;
- comparar uma solução com channel e outra com mutex para o mesmo problema.

## Próximo passo
Depois desta etapa, avance para `10-observability-and-production/` para transformar aplicações funcionais em serviços mais preparados para produção.
