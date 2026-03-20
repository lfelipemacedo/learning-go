# Primeiros Exemplos

## Objetivo
Exercícios iniciais para ganhar familiaridade com sintaxe, módulos e execução de programas Go.

## Conteúdo
- `hello-world/` — primeiro programa e execução com `go run`.
- `simple-cli/` — CLI simples que filtra linhas de um `log.txt` por nível.
- `simple-web-server/` — servidor HTTP básico com duas rotas.

## Como executar
Entre em cada pasta e rode:
```bash
go run .
```

### CLI (simple-cli)
Você pode filtrar o nível com `--level`, por exemplo:
```bash
go run . --level=ERROR
```

### Web server (simple-web-server)
O servidor sobe na porta `3000` e expõe:
- `/` — resposta simples.
- `/home` — serve o arquivo `home.html`.
