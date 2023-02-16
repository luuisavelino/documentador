# Documentador

Ferramenta que gera automaticamente uma documentação utilizando a API OpenAI.

## Requerimentos

- Cadastro na API OpenAI;
- Chave de API;
- Arquivo de código fonte.

## Instrução de uso

1. Cadastre-se na API OpenAI e obtenha a sua chave de API;
2. Salve a chave de API em uma variável de ambiente chamada OPENAI_KEY;
3. Execute a ferramenta e informe o caminho do arquivo de código fonte.

## Descrição do código

O código se conecta à API OpenAI utilizando a chave de API fornecida pelo usuário e envia o código fonte para análise. A API OpenAI processa o código fonte e gera a documentação, que é salva em um arquivo na mesma pasta do código fonte.
