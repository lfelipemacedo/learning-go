# HTTP e APIs

## Objetivo da etapa
Esta etapa foca na construção de APIs HTTP usando a biblioteca padrão, com atenção a handlers, JSON, middleware, contexto e contratos claros.

## O que você vai aprender
- como criar handlers com `net/http`;
- como serializar e desserializar JSON;
- como estruturar rotas e middleware simples;
- como usar `context.Context` em fluxos HTTP;
- como tratar erros e respostas de forma consistente.

## Estrutura da pasta
- `README.md` — guia de estudo da etapa.
- futuros exemplos podem incluir APIs CRUD, middlewares de logging e rotas de health check.

## Como estudar esta etapa
1. Comece criando um endpoint simples com resposta JSON.
2. Em seguida, adicione parsing de request, validação e status codes corretos.
3. Introduza middleware só depois de entender claramente o fluxo básico do request.
4. Propague `context.Context` até dependências internas sempre que fizer sentido.

## Como executar
Quando exemplos forem adicionados nesta etapa, rode:

```bash
go run .
```

## O que observar ao estudar
- a standard library já resolve muito do que um framework costuma abstrair;
- contratos HTTP claros facilitam testes e manutenção;
- `context` é importante para timeout, cancelamento e rastreamento de operações;
- consistência nas respostas melhora a experiência de quem consome a API.

## Exercícios sugeridos
- criar um CRUD básico em memória para um recurso simples;
- adicionar middleware de logging e recuperação de pânico;
- padronizar respostas de erro com JSON;
- escrever um teste de integração simples para um handler.

## Próximo passo
Depois desta etapa, avance para `8-database-sql/` para persistir dados de forma explícita e previsível.
