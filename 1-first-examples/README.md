# Primeiros Exemplos

## Objetivo da etapa
Esta etapa existe para criar familiaridade com a estrutura mínima de programas em Go, com a execução via `go run` e com os primeiros formatos de aplicação mais comuns: terminal e HTTP.

## O que você vai aprender
- como um programa Go é organizado a partir de `package main` e `func main`;
- como executar projetos com `go run .`;
- como ler argumentos simples de linha de comando;
- como subir um servidor HTTP básico usando a standard library;
- como dividir os estudos em exemplos pequenos e fáceis de revisar.

## Estrutura da pasta
- `hello-world/` — primeiro programa para validar ambiente, sintaxe e execução.
- `simple-cli/` — exemplo de CLI simples para filtrar linhas de `log.txt` por nível.
- `simple-web-server/` — servidor HTTP básico com rotas simples e resposta HTML.

## Como estudar esta etapa
1. Comece pelo `hello-world/` para validar a estrutura mínima de um programa Go.
2. Siga para `simple-cli/` para observar como entrada via terminal pode ser tratada.
3. Finalize com `simple-web-server/` para entender como Go lida com HTTP sem framework.
4. Compare os três exemplos e identifique o que muda e o que permanece igual entre eles.

## Como executar
Entre na pasta do exemplo desejado e rode:

```bash
go run .
```

### CLI (`simple-cli`)
Você pode filtrar o nível com `--level`, por exemplo:

```bash
go run . --level=ERROR
```

### Web server (`simple-web-server`)
O servidor sobe na porta `3000` e expõe:
- `/` — resposta simples.
- `/home` — serve o arquivo `home.html`.

## O que observar ao rodar
- como o `main` funciona como ponto de entrada da aplicação;
- como argumentos mudam o comportamento da CLI;
- como handlers HTTP podem ser definidos de forma direta;
- como exemplos pequenos ajudam a entender a linguagem sem excesso de abstração.

## Exercícios sugeridos
- alterar o `hello-world` para imprimir informações recebidas por argumentos;
- adicionar um novo filtro no `simple-cli`, como busca por palavra-chave;
- criar uma nova rota no `simple-web-server` retornando JSON;
- escrever um pequeno resumo do que cada exemplo ensina sobre o estilo da linguagem.

## Próximo passo
Depois desta etapa, avance para `2-simple-data-types/` para consolidar os tipos básicos e o comportamento das variáveis em Go.
