# Projeto 01 — CLI de Tarefas

## Objetivo do projeto
Construir uma aplicação de linha de comando para gerenciar tarefas, consolidando structs, slices, arquivos, JSON e parsing de argumentos.

## O que este projeto exercita
- modelagem de dados com structs;
- persistência local em arquivo JSON;
- comandos simples de terminal;
- separação entre entrada, domínio e persistência;
- organização inicial de um projeto Go sem framework.

## Escopo inicial sugerido
- adicionar tarefa;
- listar tarefas;
- concluir tarefa;
- remover tarefa;
- persistir dados em arquivo local.

## Como desenvolver
1. Comece com um fluxo simples em memória.
2. Adicione persistência em JSON.
3. Extraia responsabilidades em pacotes apenas quando necessário.
4. Documente comandos e exemplos de uso.

## Critérios de aceite
- os comandos principais funcionam sem quebrar o fluxo;
- os dados persistem entre execuções;
- mensagens de erro são claras para argumentos inválidos;
- o projeto é fácil de rodar e entender.

## Próximo passo
Depois deste projeto, avance para `../02-parser-library/`.
