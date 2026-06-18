# Notas de release: v1.0.0-rc.1 (rascunho)

Rascunho das notas da Release Candidate do MVP. Serve de corpo para a release
`v1.0.0-rc.1` quando a equipe promover a `dev` para a `main` e criar a tag. Ainda
não publicado: não há tag nem imagem versionada gerada por estas notas.

## Descrição

Primeira Release Candidate do Markupp: o MVP completo e integrado entre o
servidor Go (API REST + SQLite) e o plugin do Obsidian, em condições de ser
demonstrado.

## Funcionalidades entregues

- CRUD de notas markdown via plugin do Obsidian e API REST: criar, ler, editar,
  renomear e excluir.
- Listagem (`GET /notes`) e busca por substring (`GET /notes/search`) no
  servidor.
- Sincronização via Source Control View no plugin (Fetch, Pull, Push, Sync), com
  status de notas novas, modificadas, deletadas e conflitantes.
- Detecção de conflito por optimistic locking (`409`) e resolução por
  sobrescrita forçada (force pull/push).

## Correções e melhorias

- Fixa a versão do golangci-lint no CI, tornando o lint reprodutível e alinhado
  ao ambiente local (PR #94).
- Corrige vulnerabilidades de dependências do plugin (PR #95).
- Adiciona testes de aceitação do MVP associados às issues (PR #96).
- Revisa a documentação para o RC: fechamento dos riscos, limitações conhecidas
  e escopo do MVP entregue no README (PR #97).
- Hardening anterior consolidado no marco `v0.4.0`: isolamento da persistência na
  camada de storage ([ADR-0010](adrs/ADR-0010-isolamento-persistencia.md)),
  Source Control View ([ADR-0011](adrs/ADR-0011-source-control-view.md)),
  normalização de timestamps para UTC e push/pull resilientes a conflito.

## Limitações conhecidas

Ver [limitações conhecidas](limitacoes-conhecidas.md). Em resumo: sem
versionamento/histórico real, sem autenticação (uso local), sem organização
hierárquica e sem busca semântica; a validação ponta a ponta no Obsidian é
manual.

## Ambiente de teste

- Servidor via Docker (ver [DEPLOY](DEPLOY.md)), em `http://localhost:8080`.
  Prefira a tag versionada da imagem (`riedelgab/ifsces2:v1.0.0-rc.1`) à `latest`
  para reprodutibilidade.
- Plugin instalável pelos artefatos anexados à release (`main.js`,
  `manifest.json`, `styles.css`).
- Sem credenciais: o servidor não tem autenticação
  ([ADR-0005](adrs/ADR-0005-seguranca-fora-do-mvp.md)).
