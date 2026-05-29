# Arquitetura C4 do Markupp

Documento único com a visão C4 do projeto Markupp, considerando o plugin Obsidian e o servidor REST em Go.

## 1. Contexto

```mermaid
flowchart LR
    user[Usuário]
    obsidian[Obsidian Desktop]
    plugin[Plugin Markupp]
    server[Servidor Markupp]

    user --> obsidian
    obsidian --> plugin
    plugin -->|HTTP/JSON| server
```

## 2. Contêineres

```mermaid
flowchart LR
    obsidian[Obsidian Desktop]
    plugin[Plugin Markupp - TypeScript]
    vault[(Vault Obsidian - Markdown)]
    api[API REST - Go + chi]
    service[Serviço de Notas]
    repo[Repositório SQLite]
    db[(SQLite - markupp.db)]

    obsidian --> plugin
    plugin --> vault
    plugin -->|HTTP/JSON| api
    api --> service
    service --> repo
    repo --> db
```

## 3. Componentes do Servidor

```mermaid
flowchart LR
    subgraph server[Servidor Markupp]
        direction LR

        router[Router\ninternal/api/router.go]
        handler[notesHandler\ninternal/api/notes_handler.go]
        service[notes.Service\ninternal/notes/notes.go]
        repo[SqliteNotesRepository\ninternal/storage/notes.go]
        queries[sqlc Queries\ninternal/storage/gen]
        db[(SQLite)]

        router --> handler
        handler --> service
        service --> repo
        repo --> queries
        queries --> db
    end
```

## 4. Leitura Arquitetural

- O usuário interage com o sistema pelo Obsidian, onde o plugin Markupp executa comandos de sincronização, upload, download e importação.
- O plugin consome a API REST do servidor por HTTP/JSON, usando as rotas `/notes`.
- O servidor em Go expõe as rotas, valida as regras de negócio no serviço e persiste as notas em SQLite.
- O vault do Obsidian mantém os arquivos Markdown locais, enquanto o servidor mantém a fonte persistente das notas.

## 5. Decisões Estruturais

- Separação clara entre interface de usuário, API e persistência.
- Persistência única em SQLite, acessada via `sqlc`.
- O plugin permanece desacoplado do banco, falando apenas com a API.
