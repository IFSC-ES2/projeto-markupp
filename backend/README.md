# Markupp backend

Servidor REST do Markupp em Go.

## Build

```sh
go build ./cmd/markupp
```

## Execução

```sh
./markupp
```

## Configuração

O servidor lê a configuração de um arquivo JSON. O caminho é resolvido nesta ordem:

1. Variável de ambiente `MARKUPP_CONFIG_PATH`
2. Fallback `./config.json` no diretório de trabalho

Se o arquivo não existir, os defaults são aplicados silenciosamente. Chaves ausentes no arquivo preservam o default correspondente.

Exemplo em `config.example.json`.
