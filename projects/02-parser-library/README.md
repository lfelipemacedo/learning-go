# Projeto 02 — Biblioteca de Parser/Importação

## Objetivo do projeto
Criar uma biblioteca reutilizável para importar dados de CSV e JSON em structs de domínio, reforçando interfaces pequenas, validação e tratamento de erro.

## O que este projeto exercita
- desenho de API reutilizável;
- leitura de arquivos e streams;
- validação de entrada;
- erros com contexto;
- testes unitários orientados a tabela.

## Escopo inicial sugerido
- importar CSV e JSON;
- validar registros;
- retornar erros descritivos;
- disponibilizar uma API simples para outros projetos.

## Como desenvolver
1. Defina primeiro o contrato de importação.
2. Implemente parsing para um formato.
3. Extraia validação para uma camada própria.
4. Adicione suporte ao segundo formato e testes.

## Critérios de aceite
- a API da biblioteca é pequena e clara;
- erros ajudam a localizar a falha;
- há cobertura para casos válidos e inválidos;
- outro projeto consegue reutilizar a biblioteca com pouco acoplamento.

## Próximo passo
Depois deste projeto, avance para `../03-discount-engine/`.
