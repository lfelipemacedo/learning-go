# Tipos de Dados Simples

## Objetivo da etapa
Esta etapa serve para consolidar os tipos básicos da linguagem e entender como Go trata variáveis, constantes, zero values e erros como valores explícitos.

## O que você vai aprender
- como funcionam strings, números, booleanos e constantes em Go;
- a diferença entre tipagem explícita e inferida;
- como `error` participa do fluxo normal da linguagem;
- como zero values influenciam o comportamento das variáveis;
- como escrever exemplos pequenos para revisar conceitos essenciais.

## Estrutura da pasta
- `simple-data-types.go` — demonstra exemplos de strings, números, booleanos, `error`, variáveis e constantes.

## Como estudar esta etapa
1. Leia o arquivo observando cada grupo de tipos separadamente.
2. Rode o programa e compare a saída com sua expectativa antes da execução.
3. Revise quais valores são inicializados automaticamente e quais exigem definição explícita.
4. Observe onde Go favorece simplicidade em vez de conversões implícitas ou comportamento “mágico”.

## Como executar
```bash
go run .
```

## O que observar ao rodar
- diferenças entre strings interpretadas e raw strings;
- como números e conversões exigem mais clareza do que em linguagens com mais coerções implícitas;
- o papel de `nil` na ausência de erro;
- como variáveis e constantes ajudam a tornar a intenção do código mais clara.

## Exercícios sugeridos
- criar exemplos adicionais usando `rune` e `byte` para observar diferenças com strings;
- escrever uma função que retorne valor + erro para praticar o padrão idiomático;
- testar conversões entre `int`, `float64` e `string` com comentários sobre o que exige conversão explícita;
- separar o arquivo em blocos temáticos para deixar a revisão mais didática.

## Próximo passo
Depois desta etapa, avance para `3-pointers-and-values/` para entender melhor cópia, mutação e compartilhamento de estado em Go.
