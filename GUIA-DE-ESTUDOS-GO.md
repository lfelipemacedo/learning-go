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

#### Critérios de aceite
- os comandos principais (`add`, `list`, `done`, `remove`) funcionam sem quebrar o fluxo da aplicação;
- os dados persistidos em JSON sobrevivem a reinicializações da CLI;
- mensagens de erro orientam o usuário quando argumentos inválidos são enviados;
- a estrutura do projeto separa parsing de comando, regra de negócio e persistência de forma simples;
- existe ao menos um pequeno conjunto de testes para as regras de manipulação de tarefas.

#### Sugestões de desenvolvimento
- comece com uma versão em um único arquivo para validar o fluxo completo;
- depois extraia pacotes como `task`, `storage` e `cmd` apenas quando a duplicação aparecer;
- implemente primeiro persistência local síncrona e só depois pense em melhorias como filtros ou prioridades;
- documente exemplos de uso no `README` do projeto para facilitar revisões futuras.

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

#### Critérios de aceite
- a biblioteca expõe uma API pequena e clara para importar CSV e JSON;
- erros de validação informam linha, campo ou motivo da falha quando possível;
- formatos suportados compartilham comportamento comum sem duplicação excessiva;
- os testes cobrem entradas válidas, inválidas e casos de borda;
- pelo menos um projeto consumidor consegue reutilizar a biblioteca sem adaptações complexas.

#### Sugestões de desenvolvimento
- defina primeiro um tipo de domínio simples para servir de contrato de importação;
- crie uma camada de validação separada da leitura do arquivo;
- prefira interfaces só para dependências externas, como `io.Reader`;
- escreva testes table-driven desde cedo para facilitar a expansão de formatos e regras.

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

#### Critérios de aceite
- as regras de desconto estão cobertas por testes orientados a tabela com casos positivos e negativos;
- cupons inválidos, combinações incompatíveis e limites de desconto estão explicitamente testados;
- o módulo separa cálculo, validação e aplicação de contexto promocional de forma legível;
- existe pelo menos um benchmark comparando versões ou cenários de maior volume;
- a API do módulo permite reutilização em outro projeto sem depender de infraestrutura externa.

#### Sugestões de desenvolvimento
- escreva primeiro os cenários de negócio em formato de tabela antes da implementação;
- comece com uma função pura de cálculo e só depois introduza estruturas auxiliares;
- use nomes de testes que descrevam a regra de negócio, não apenas a função exercitada;
- ao benchmarkar, compare cenários realistas como grandes lotes de pedidos ou múltiplos cupons.

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

#### Critérios de aceite
- a API expõe CRUD funcional com contratos JSON consistentes e códigos HTTP corretos;
- paginação e filtros básicos funcionam com parâmetros claros e validados;
- erros de validação e falhas internas retornam respostas padronizadas;
- timeouts e cancelamentos via `context.Context` são propagados até a camada de persistência;
- existem testes cobrindo handlers principais e ao menos um fluxo de integração completo.

#### Sugestões de desenvolvimento
- comece com armazenamento em memória para validar contratos HTTP antes de adicionar banco;
- defina DTOs explícitos para request e response, evitando acoplar structs internas ao transporte;
- adicione middlewares de logging e recuperação antes de crescer o número de endpoints;
- documente exemplos de request/response para facilitar consumo e revisão da API.

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

#### Critérios de aceite
- o fluxo de criação e fechamento de pedidos usa transações de forma consistente;
- consultas de leitura retornam dados completos sem ambiguidade entre pedido, itens e cliente;
- falhas de banco são tratadas com mensagens úteis e rollback quando necessário;
- migrations permitem recriar o ambiente local sem intervenção manual complexa;
- testes de integração validam ao menos o caminho feliz e um cenário de falha transacional.

#### Sugestões de desenvolvimento
- modele primeiro o esquema relacional e os casos de uso antes de escrever os repositórios;
- mantenha SQL explícito em arquivos ou constantes bem nomeadas para facilitar revisão;
- use dados de teste realistas para validar joins, filtros e consistência de transações;
- introduza ferramentas como `sqlc` apenas depois de entender claramente o fluxo com `database/sql`.

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

#### Critérios de aceite
- o pipeline processa múltiplos arquivos em paralelo com resultado determinístico ou claramente documentado;
- cancelamento interrompe stages corretamente quando ocorre erro crítico ou timeout;
- não há vazamento de goroutines ao finalizar processamento ou interromper execução;
- métricas básicas permitem comparar execução sequencial versus concorrente;
- `go test -race` passa para os componentes principais do pipeline.

#### Sugestões de desenvolvimento
- implemente primeiro uma versão sequencial para ter baseline funcional e de performance;
- depois quebre o fluxo em stages pequenos com contratos claros entre canais;
- prefira encapsular ownership de fechamento de canais para evitar deadlocks;
- meça throughput e latência antes e depois de cada mudança de concorrência.

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

#### Critérios de aceite
- o endpoint de pagamento responde de forma idempotente para requisições repetidas com a mesma chave;
- timeouts e retries seguem política definida e não multiplicam efeitos colaterais indevidos;
- logs estruturados permitem rastrear uma requisição ponta a ponta;
- métricas básicas expõem volume, sucesso, falha e latência das integrações externas;
- o serviço realiza shutdown gracioso sem perder requisições em andamento de forma abrupta.

#### Sugestões de desenvolvimento
- comece simulando o provedor externo com um stub HTTP controlável;
- modele explicitamente estados do pagamento, como `pending`, `approved` e `failed`;
- implemente idempotência antes de adicionar retry para evitar duplicidade de efeitos;
- teste cenários de timeout, resposta lenta e erro intermitente como parte da definição de pronto.

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

#### Critérios de aceite
- o sistema entrega um fluxo fim a fim funcional, da ingestão até a consulta dos resultados;
- a arquitetura está documentada com responsabilidades claras entre entrada, processamento e persistência;
- há cobertura de testes para regras centrais e testes de integração para o fluxo principal;
- o projeto possui configuração reproduzível localmente, preferencialmente com Docker ou scripts simples;
- logs, métricas e tratamento de falhas demonstram preocupação real com operação em produção.

#### Sugestões de desenvolvimento
- escreva uma especificação curta do domínio e dos casos de uso antes de escolher bibliotecas;
- priorize uma primeira entrega vertical simples, mesmo sem todos os componentes distribuídos;
- mantenha observabilidade desde o início, ao menos com logs estruturados e health checks;
- encerre o projeto com uma retrospectiva técnica: o que mudaria, onde escalaria e quais trade-offs aceitou.

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
