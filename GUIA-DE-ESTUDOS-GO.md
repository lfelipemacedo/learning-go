# Guia de Estudos de Golang

Este guia foi pensado para alguém que já é **desenvolvedor Java experiente** e já conhece o **básico de Go**. A proposta é acelerar sua evolução em Golang usando uma trilha que mistura:

- fundamentos da linguagem;
- diferenças de mentalidade entre Java e Go;
- prática incremental;
- projetos reais;
- foco em escrita de código simples, idiomático e sustentável.

> Objetivo final: sair de “sei o básico” para “consigo projetar, implementar, testar e operar aplicações Go com segurança”.

---

## 1. Estratégia de estudo

Como você já vem de Java, o maior risco não é entender sintaxe, e sim tentar escrever Go “com cabeça de Java”.

### O que você deve buscar
- Preferir **simplicidade** em vez de abstrações profundas.
- Usar **composição** em vez de hierarquias complexas.
- Modelar pacotes com responsabilidade clara.
- Tratar erros explicitamente.
- Aprender concorrência com responsabilidade.
- Escrever código legível antes de escrever código “genérico demais”.

### O que evitar no começo
- Reproduzir padrões de OO de Java sem necessidade.
- Criar camadas demais para aplicações pequenas.
- Usar interfaces cedo demais “por arquitetura”.
- Introduzir goroutines em tudo.
- Ignorar profiling, contexto e cancelamento em apps I/O-bound.

### Ritmo sugerido
- **6 a 10 horas por semana**.
- Para cada módulo:
  - 20% teoria;
  - 30% exercícios curtos;
  - 50% projeto prático.

---

## 2. Mapa de evolução: do básico ao avançado

A trilha abaixo vai do mais simples ao mais complexo. Cada etapa tem:

- **objetivo**;
- **tópicos**;
- **projeto prático**;
- **critérios de conclusão**.

---

## 3. Fase 1 — Fundamentos e fluência com a linguagem

### Objetivo
Ganhar velocidade com sintaxe, tipos, organização de pacotes e ferramentas básicas.

### Tópicos para dominar
- declaração de variáveis;
- tipos básicos;
- arrays, slices e maps;
- structs;
- ponteiros;
- funções e múltiplos retornos;
- visibilidade por maiúscula/minúscula;
- organização em packages;
- `go mod`, `go run`, `go build`, `go test`, `go fmt`.

### Comparação com Java
- Em Go, a simplicidade da estrutura de código é parte da linguagem.
- Não existe a mesma centralidade de classes e herança.
- Métodos existem, mas o design gira mais em torno de **tipos + funções + interfaces pequenas**.

### Exercícios curtos
1. Criar um programa que receba argumentos da linha de comando.
2. Implementar funções utilitárias com slices e maps.
3. Modelar um pequeno domínio com structs (`User`, `Order`, `Product`).
4. Reescrever algum exercício simples de Java em Go sem tentar manter a mesma arquitetura.

### Projeto prático 1 — CLI de tarefas
Construa um **gerenciador de tarefas em linha de comando**.

#### Escopo mínimo
- adicionar tarefa;
- listar tarefas;
- concluir tarefa;
- remover tarefa;
- persistir em JSON local.

#### O que você vai praticar
- structs;
- slices;
- leitura/escrita de arquivo;
- serialização com `encoding/json`;
- organização básica em pacotes;
- parsing de argumentos.

#### Evoluções opcionais
- filtros por status;
- prioridade;
- datas;
- exportação CSV.

#### Critério de conclusão
Você consegue rodar a CLI, persistir dados sem bugs básicos e manter o projeto organizado sem framework.

---

## 4. Fase 2 — Modelagem idiomática em Go

### Objetivo
Aprender a escrever código “go-like”, sem trazer acoplamentos desnecessários do ecossistema Java.

### Tópicos para dominar
- métodos em structs;
- interfaces pequenas;
- composição;
- `io.Reader` e `io.Writer` como exemplos de design idiomático;
- tratamento de erros com `error`;
- `errors.Is`, `errors.As`, `fmt.Errorf` com `%w`;
- boas práticas de naming;
- construção de pacotes coesos.

### Comparação com Java
- Em Java, interfaces muitas vezes são desenhadas antes da implementação.
- Em Go, frequentemente a interface aparece no **consumidor**, não no produtor.
- Em Java você pode criar muitas abstrações cedo; em Go, isso costuma piorar a legibilidade.

### Exercícios curtos
1. Refatorar o projeto CLI para separar domínio, storage e comando.
2. Introduzir interfaces só onde elas forem realmente úteis para teste ou desacoplamento.
3. Modelar erros de domínio e erros de infraestrutura.

### Projeto prático 2 — Biblioteca de parser/importação
Crie uma **biblioteca reutilizável** que importe dados de arquivos CSV e JSON para structs de domínio.

#### Escopo mínimo
- leitura de CSV e JSON;
- validação de registros;
- retorno de erros descritivos;
- API simples para reuso em outros projetos.

#### O que você vai praticar
- design de pacotes;
- interfaces pequenas;
- testes unitários;
- tratamento de erro idiomático.

#### Critério de conclusão
Outro projeto do repositório consegue usar essa biblioteca com baixo acoplamento e boa legibilidade.

---

## 5. Fase 3 — Testes, qualidade e toolchain

### Objetivo
Dominar o ciclo de desenvolvimento profissional em Go.

### Tópicos para dominar
- testes com o pacote `testing`;
- testes orientados a tabela (table-driven tests);
- subtests;
- benchmarks;
- coverage;
- organização de `_test.go`;
- `go test ./...`;
- `go vet`;
- `golangci-lint` (se quiser adicionar depois);
- profiling básico com `pprof`.

### Exercícios curtos
1. Escrever table-driven tests para validação de regras.
2. Criar benchmarks para parsing ou transformação de dados.
3. Encontrar gargalos simples com benchmark + profiling.

### Projeto prático 3 — Motor de regras de desconto
Implemente um **módulo de cálculo de descontos** para pedidos.

#### Escopo mínimo
- regras por tipo de cliente;
- regras por quantidade;
- cupons;
- testes cobrindo cenários principais e de borda.

#### O que você vai praticar
- design testável;
- table-driven tests;
- benchmark de regras;
- refatoração segura.

#### Critério de conclusão
Cobertura razoável, testes claros, benchmark demonstrando entendimento de custo computacional.

---

## 6. Fase 4 — HTTP, APIs e contexto

### Objetivo
Construir serviços web idiomáticos com a biblioteca padrão antes de depender de frameworks.

### Tópicos para dominar
- `net/http`;
- handlers;
- middleware;
- roteamento simples;
- serialização JSON;
- validação de requests;
- status codes;
- `context.Context`;
- timeout e cancelamento;
- logging estruturado.

### Comparação com Java
- Em Java/Spring, o framework costuma esconder muito da infraestrutura HTTP.
- Em Go, entender a biblioteca padrão é uma vantagem enorme.
- Em vez de depender cedo de framework completo, vale dominar bem o básico com `net/http`.

### Exercícios curtos
1. Criar endpoints CRUD simples.
2. Adicionar middleware de logging e recuperação de pânico.
3. Propagar `context.Context` até a camada de acesso a dados.

### Projeto prático 4 — API REST de catálogo
Construa uma **API REST de catálogo de produtos**.

#### Escopo mínimo
- CRUD de produtos;
- paginação;
- filtros simples;
- validação;
- persistência inicial em memória e depois em banco.

#### Evolução recomendada
- adicionar PostgreSQL;
- versionamento de API;
- documentação OpenAPI/Swagger;
- health check e readiness.

#### O que você vai praticar
- design de handlers;
- DTOs;
- contexto e timeouts;
- separação entre transporte, domínio e persistência.

#### Critério de conclusão
API com rotas limpas, tratamento consistente de erro, testes de integração básicos e uso correto de contexto.

---

## 7. Fase 5 — Banco de dados e persistência real

### Objetivo
Aprender integração robusta com banco relacional sem perder simplicidade.

### Tópicos para dominar
- `database/sql`;
- drivers;
- transactions;
- prepared statements;
- mapeamento manual de rows;
- migrations;
- trade-offs entre SQL explícito e ORMs;
- `sqlc` ou query builders como evolução opcional.

### Comparação com Java
- Se você vem de JPA/Hibernate, prepare-se para escrever mais SQL.
- Em compensação, você ganha mais previsibilidade, menos magia e mais controle de performance.

### Exercícios curtos
1. Criar repositórios com `database/sql`.
2. Implementar transação de criação de pedido com itens.
3. Medir impacto de queries mal desenhadas.

### Projeto prático 5 — Sistema de pedidos
Construa um **backend de pedidos** com PostgreSQL.

#### Escopo mínimo
- clientes;
- produtos;
- pedidos;
- itens do pedido;
- transação de fechamento do pedido;
- consulta por status.

#### O que você vai praticar
- SQL explícito;
- transações;
- organização de camadas;
- testes de integração com banco.

#### Critério de conclusão
Você consegue modelar transações, lidar com erros de persistência e manter o domínio compreensível.

---

## 8. Fase 6 — Concorrência de verdade

### Objetivo
Sair do entendimento superficial de goroutines e channels para uso seguro e útil.

### Tópicos para dominar
- goroutines;
- channels buffered e unbuffered;
- `select`;
- `sync.WaitGroup`;
- `sync.Mutex` e `RWMutex`;
- `context` para cancelamento;
- pipelines;
- worker pools;
- prevenção de race conditions;
- uso de `go test -race`.

### Comparação com Java
- Você já conhece `ExecutorService`, threads, locks e futures.
- Em Go, o modelo muda para uma combinação de **goroutines leves + channels + sincronização explícita**.
- Nem todo problema precisa de channels; às vezes um mutex é mais simples e correto.

### Exercícios curtos
1. Implementar fan-out/fan-in.
2. Criar um worker pool para processamento paralelo.
3. Introduzir cancelamento por timeout.
4. Detectar race condition e corrigir.

### Projeto prático 6 — Processador concorrente de arquivos
Crie um **pipeline concorrente** que leia muitos arquivos, processe conteúdo e gere um relatório consolidado.

#### Escopo mínimo
- leitura paralela;
- stage de parsing;
- stage de agregação;
- cancelamento em caso de erro crítico;
- métricas simples de throughput.

#### O que você vai praticar
- pipelines;
- channels;
- coordenação de goroutines;
- tratamento correto de fechamento de canais;
- race detector.

#### Critério de conclusão
O sistema ganha desempenho real sem vazamentos de goroutines nem corrupção de estado.

---

## 9. Fase 7 — Arquitetura de serviços e observabilidade

### Objetivo
Evoluir de “aplicação que funciona” para “serviço que pode rodar em produção”.

### Tópicos para dominar
- configuração por ambiente;
- graceful shutdown;
- structured logging;
- métricas;
- tracing conceitual;
- health/readiness/liveness endpoints;
- retry com cautela;
- idempotência;
- limites e timeouts;
- princípios de observabilidade.

### Exercícios curtos
1. Adicionar shutdown gracioso a uma API.
2. Instrumentar logs estruturados com request ID.
3. Expor métricas de latência e contagem.

### Projeto prático 7 — Serviço de pagamentos fake/integrador
Monte um **serviço de integração** que recebe pedidos e conversa com um provedor externo simulado.

#### Escopo mínimo
- endpoint para processar pagamento;
- timeout por chamada externa;
- retry controlado;
- idempotência por chave;
- logs estruturados;
- métricas básicas.

#### O que você vai praticar
- resiliência;
- contratos HTTP;
- observabilidade;
- comportamento de produção.

#### Critério de conclusão
Seu serviço falha de forma previsível, observável e controlada.

---

## 10. Fase 8 — Projeto avançado / portfólio

### Objetivo
Consolidar o aprendizado em um projeto que demonstre domínio técnico de Go.

### Projeto prático 8 — Plataforma de ingestão e processamento
Escolha um projeto mais ambicioso, por exemplo:

- ingestão de eventos via HTTP;
- fila interna ou processamento assíncrono;
- persistência em PostgreSQL;
- workers concorrentes;
- endpoint de consulta de resultados;
- métricas e logs;
- testes unitários e de integração;
- Docker e ambiente local.

### Possíveis temas
- importador de transações financeiras;
- agregador de preços;
- pipeline de processamento de logs;
- sistema de notificações;
- mini plataforma de analytics.

### O que esse projeto deve demonstrar
- domínio da linguagem;
- capacidade de modelagem;
- uso correto de concorrência;
- qualidade de testes;
- clareza arquitetural;
- prontidão para produção.

---

## 11. Ordem sugerida dos projetos

Se quiser uma sequência objetiva, siga esta ordem:

1. **CLI de tarefas**;
2. **Biblioteca de parser/importação**;
3. **Motor de regras de desconto**;
4. **API REST de catálogo**;
5. **Sistema de pedidos com PostgreSQL**;
6. **Processador concorrente de arquivos**;
7. **Serviço de pagamentos fake/integrador**;
8. **Projeto avançado de portfólio**.

Essa ordem funciona bem porque cada projeto reaproveita competências do anterior.

---

## 12. Roadmap de 16 semanas

### Semanas 1–2
- revisar fundamentos;
- terminar exercícios de sintaxe e tipos;
- iniciar CLI de tarefas.

### Semanas 3–4
- refatorar CLI;
- estudar interfaces, composição e erros;
- criar biblioteca de parser/importação.

### Semanas 5–6
- foco em testes, benchmarks e profiling;
- implementar motor de regras de desconto.

### Semanas 7–9
- estudar `net/http`, middleware, JSON e contexto;
- construir API REST de catálogo.

### Semanas 10–11
- integrar PostgreSQL;
- implementar sistema de pedidos com transações.

### Semanas 12–13
- estudar concorrência seriamente;
- construir processador concorrente de arquivos.

### Semanas 14–15
- observabilidade, shutdown, timeouts e resiliência;
- criar serviço de pagamentos fake/integrador.

### Semana 16+
- projeto avançado de portfólio;
- revisão geral;
- documentação e polimento.

---

## 13. Checklist de domínio por assunto

### Linguagem
- [ ] Sei usar slices e maps com segurança.
- [ ] Entendo quando usar ponteiros.
- [ ] Sei modelar structs e métodos de forma idiomática.
- [ ] Sei organizar packages sem acoplamento excessivo.

### Erros e design
- [ ] Sei encapsular e propagar erros corretamente.
- [ ] Sei quando criar interface e quando não criar.
- [ ] Consigo refatorar para composição em vez de abstrações desnecessárias.

### Testes
- [ ] Sei escrever testes orientados a tabela.
- [ ] Sei criar benchmarks.
- [ ] Sei usar `go test ./...` com confiança.

### HTTP e APIs
- [ ] Sei criar handlers e middleware.
- [ ] Sei propagar `context.Context`.
- [ ] Sei padronizar respostas e erros.

### Persistência
- [ ] Sei trabalhar com `database/sql`.
- [ ] Sei modelar transações.
- [ ] Sei escrever queries claras e previsíveis.

### Concorrência
- [ ] Sei usar goroutines com propósito.
- [ ] Sei escolher entre channels e mutex.
- [ ] Sei detectar e corrigir race conditions.

### Produção
- [ ] Sei fazer graceful shutdown.
- [ ] Sei trabalhar com timeout, retry e idempotência.
- [ ] Sei adicionar logs e métricas úteis.

---

## 14. Boas práticas específicas para quem vem de Java

### 1. Não tente “encaixar Spring” mentalmente em Go
Go não depende do mesmo modelo de framework centralizador. Muitas vezes menos estrutura produz melhor resultado.

### 2. Prefira interfaces pequenas
Interfaces em Go brilham quando descrevem comportamento mínimo, não contratos gigantes.

### 3. Comece com a standard library
Antes de buscar framework web, ORM ou biblioteca de DI, domine o que a biblioteca padrão já entrega.

### 4. Escreva SQL conscientemente
Se você tem experiência com abstrações pesadas de ORM, Go pode ser uma ótima oportunidade para recuperar clareza na camada de dados.

### 5. Aprenda profiling cedo
Go é excelente para serviços eficientes, mas isso exige medir em vez de supor.

### 6. Use concorrência para resolver problemas reais
Não use goroutines como demonstração de linguagem. Use quando houver ganho claro de latência, throughput ou modelagem.

---

## 15. Como estudar cada projeto

Para cada projeto da trilha, siga este ciclo:

1. **Especificar**
   - escreva escopo, regras e limites;
   - defina o que é versão 1.

2. **Modelar**
   - crie structs principais;
   - defina pacotes;
   - pense no fluxo dos dados.

3. **Implementar o caminho feliz**
   - faça primeiro funcionar;
   - evite abstrações prematuras.

4. **Adicionar testes**
   - cubra regras centrais;
   - use table-driven tests sempre que fizer sentido.

5. **Refatorar**
   - reduza duplicação;
   - melhore nomes;
   - simplifique interfaces.

6. **Operacionalizar**
   - logs;
   - configuração;
   - timeouts;
   - documentação de execução.

7. **Revisar**
   - o código parece Go idiomático?
   - existe alguma “herança de Java” desnecessária?

---

## 16. Sugestão de estrutura para seus estudos no repositório

Você pode evoluir este repositório com uma estrutura parecida com esta:

```text
1-first-examples/
2-simple-data-types/
3-pointers-and-values/
4-aggregate-data-types/
5-interfaces-and-errors/
6-testing-and-benchmarks/
7-http-apis/
8-database-sql/
9-concurrency/
10-observability-and-production/
projects/
  01-cli-tarefas/
  02-parser-library/
  03-discount-engine/
  04-product-catalog-api/
  05-order-system/
  06-file-processing-pipeline/
  07-payment-integration/
  08-portfolio-project/
```

---

## 17. Meta final

Ao concluir essa trilha, você deve ser capaz de:

- construir CLIs, bibliotecas e APIs em Go;
- modelar aplicações sem excesso de abstração;
- testar, medir e evoluir código com segurança;
- trabalhar com banco de dados, concorrência e contexto;
- entregar serviços com características reais de produção.

Se quiser, o próximo passo natural é transformar este guia em um **plano executável dentro do próprio repositório**, criando as próximas pastas/projetos em ordem de aprendizado.
