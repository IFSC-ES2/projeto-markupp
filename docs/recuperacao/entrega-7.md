# Recuperação da Entrega 7 (Sprint 3, marco `v0.3.0`)

Este documento registra a recuperação da **Entrega 7 (Sprint 3)**, conforme as
regras de recuperação da disciplina (Projeto Prático a2, 25/03/2026).

## Contexto

Na correção original da Entrega 7, toda a equipe ficou abaixo da nota mínima
(maior ou igual a 6): **5,6 / 5,4 / 4,4 / 5,0**. O principal motivo apontado foi
que o incremento da Sprint 3 **não havia sido consolidado na branch principal**
nem empacotado em uma release/tag do marco; ainda havia trabalho em PRs sobre
`dev`, mas não integrado em `main`.

- Branch de recuperação: `rec-7` (regra 3).
- Forma de entrega: pull request de `rec-7` para `main` (regra 4).
- Escopo: apenas itens dos requisitos da Sprint 3 (regra 6).

> Observação de procedimento: o conteúdo técnico da Sprint 3 já foi integrado à
> `main` e à release `v0.3.0` antes da abertura deste PR. Como as regras 11 e 12
> vedam a reescrita de histórico, esta recuperação é apresentada por meio deste
> relatório, que mapeia cada item insuficiente às evidências já rastreáveis no
> repositório (commits, PRs, tag e release). Nenhuma autoria foi atribuída
> retroativamente.

## Itens recuperados (regra 5)

Para cada item: **(a)** o que estava insuficiente/ausente na entrega original,
**(b)** o que foi feito para corrigir e **(c)** onde está a evidência.

### 1. Incremento funcional da Sprint 3

- **(a) Insuficiente:** não havia tag/release `v0.3.0` e a `main` não recebia os
  commits da equipe integrando a Sprint 3; o incremento ficou apenas em PRs
  sobre `dev`.
- **(b) Correção:** a Sprint 3 foi integrada à branch principal e empacotada no
  marco `v0.3.0`. Inclui a rota de busca `GET /notes/search` com paginação
  (PR #74) e a reengenharia de nomenclatura de `backend` para `markupp` e de
  `backendUrl` para `serverUrl` (PR #73).
- **(c) Evidência:** merge de integração `02416f3` ("integra a Sprint 3 (dev)
  para o marco v0.3.0"); PRs #73 e #74.

### 2. Documentação da arquitetura (C4)

- **(a) Insuficiente:** o diagrama/documento C4 existia apenas no PR #75, que
  ficou aberto e com mudanças solicitadas; não estava integrado em `main`.
- **(b) Correção:** PR #75 revisado, aprovado e mergeado na branch principal.
- **(c) Evidência:** `docs/arquitetura-c4.md` em `main`; merge `d3ad9b3`
  (PR #75 `docs/c4Arquitetura`).

### 3. ADRs consolidados

- **(a) Insuficiente:** não havia ADR novo consolidando as mudanças da Sprint 3.
- **(b) Correção:** consolidação dos ADRs do marco, incluindo a decisão de
  licença AGPL (ADR-0008), o rename de `backend` para `markupp` (ADR-0009) e o
  isolamento da persistência (ADR-0010).
- **(c) Evidência:** `docs/adrs/ADR-0008-licenca-agpl.md`,
  `docs/adrs/ADR-0009-rename-backend-para-markupp.md`,
  `docs/adrs/ADR-0010-isolamento-persistencia.md`; commit `2f47e15`
  ("consolida adrs da sprint 3 e licenca agpl").

### 4. Atualização das métricas

- **(a) Insuficiente:** `metricas.md` era apenas a definição; não havia valores
  observados ao fim da Sprint 3, data de coleta nem análise/comparação com a
  Sprint 2.
- **(b) Correção:** as fichas foram separadas em arquivos numerados, cada uma com
  gráfico por sprint (Sprint 1, 2 e 3), permitindo comparação visual entre os
  ciclos.
- **(c) Evidência:** as fichas de `docs/metricas/01-cobertura-de-testes.md` até
  `docs/metricas/06-burndown-do-backlog.md` e imagens
  `docs/metricas/img/*-sprint3.png`; commit `78e0feb` ("separa fichas de
  metricas com graficos por sprint").
- **Pendência conhecida:** as fichas trazem os gráficos por sprint, mas ainda
  carecem de texto com valores numéricos observados, data de coleta e análise de
  tendência/comparação da Sprint 2 para a Sprint 3.

### 5. Testes automatizados integrados ao pipeline

- **(a) Insuficiente:** o CI da `main` não executava `npm run lint`, `gofmt`,
  `go vet` ou `golangci-lint`.
- **(b) Correção:** o workflow de CI passou a rodar, na `main` e em PRs:
  verificação `gofmt`, `go vet ./...`, `golangci-lint`, `go test ./... -race
  -count=1` (servidor) e `npm ci` / `npm run lint` / `npm run build` /
  `npm test --if-present` (plugin).
- **(c) Evidência:** `.github/workflows/ci.yml` em `main`; origem no PR #72
  (CI/lint/formatação).

### 6. Integração contínua mínima

- **(a) Insuficiente:** a `main` não refletia o pipeline mais completo da
  Sprint 3 e não havia check associado ao marco.
- **(b) Correção:** o pipeline completo passou a rodar na branch principal a cada
  push/PR (mesmo workflow do item 5).
- **(c) Evidência:** `.github/workflows/ci.yml` (gatilho em `push` para `main` e
  em `pull_request`).

### 7. Release/tag do marco

- **(a) Não atendido:** o projeto contava apenas com a tag/release `v0.2.0`.
- **(b) Correção:** criação da tag `v0.3.0` e publicação da release do marco da
  Sprint 3.
- **(c) Evidência:** tag `v0.3.0` em `04b54fd` (PR #80 `release/v0.3.0`); release
  `v0.3.0` publicada no GitHub.

### 8. Registro das contribuições individuais

- **(a) Insuficiente:** não havia relatório específico da Sprint 3 consolidando o
  que cada integrante implementou/revisou/documentou.
- **(b) Correção:** criação do relatório de contribuições individuais da Sprint 3,
  com vínculo a PRs/commits por integrante.
- **(c) Evidência:** `docs/contribuicoes-sprint3.md`; commit `6da177a`
  ("adiciona relatorio de contribuicao da sprint 3").

## Resumo

| # | Item da Sprint 3 | Situação na recuperação |
| - | ---------------- | ----------------------- |
| 1 | Incremento funcional | Atendido |
| 2 | Documentação da arquitetura (C4) | Atendido |
| 3 | ADRs consolidados | Atendido |
| 4 | Atualização das métricas | Parcial (gráficos por sprint; falta análise textual) |
| 5 | Testes integrados ao pipeline | Atendido |
| 6 | Integração contínua mínima | Atendido |
| 7 | Release/tag do marco | Atendido |
| 8 | Contribuições individuais | Atendido |
