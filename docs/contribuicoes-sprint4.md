# Contribuições individuais, Sprint 4

Janela: marco da Sprint 3 (29/05) ao marco da Sprint 4 (08/06). Cada membro confirma o próprio trecho.

## Gabriela Riedel

- Adição do arquivo DEPLOY.md
- Adição do Dockerfile para publicação no Dockerhub + uso na CI

## Luís Renato Freitas de Almeida

- a confirmar pelo próprio
- Reengenharia de isolamento da persistência na camada de storage (ADR-0010).
- Consolidação do marco v0.4.0 no `main` (#82).

## Nícolas Arthur Raulino Oliveira

- Redesenho do plugin no estilo git: Source Control View com fetch/pull/push/sync,
  substituindo os comandos avulsos de sincronização (#79).
- Reengenharia da camada de sincronização do plugin: consolidação em `core/operations.ts` +
  `core/status.ts`, extração do helper `recordSynced` e push/pull resilientes a conflito 409 e
  falha parcial. Decisão e métrica antes/depois em ADR-0011.
- CD do plugin: workflow que builda e anexa o artefato instalável à release
  (`.github/workflows/release.yml`).
- Documentação de instalação do plugin (seção do `DEPLOY.md` e README).
- Revisões de PRs.

## Nicolas Pitz

- Adicionada feature de sincronização do servidor, com flags e uma nova rota.
- Atualiza Testes para incluir nova rota e services.
- Atualiza baseline registrando os novos avanços do projeto.
- Métricas da Sprint 4 nas fichas de `docs/metricas/`.
