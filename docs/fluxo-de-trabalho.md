# Fluxo de Trabalho

Este documento consolida as regras de colaboração no repositório do Markupp. Ele reúne, em um só lugar, como o código entra no projeto: branches, pull requests, revisão e integração na `main`.

## Uso obrigatório de branches

Todo desenvolvimento acontece em branches criadas a partir da `dev`. Não é permitido trabalhar diretamente na `main`.

Commits direto na `dev` devem ser justificados, de tamanho atômico e pontuais, qualquer trabalho que exija mais de um commit deve seguir em uma branch separada.

- `main`: branch de produção, só recebe merges da `dev`
- `dev`: branch de integração, onde as features convergem antes de ir para produção
- Branches de trabalho: criadas a partir da `dev`, seguindo o padrão [conventional branches](https://conventional-branch.github.io/).

Detalhes em [git-flow.md](./git-flow.md).

## Integração de mudanças via pull request

Nenhum código entra na `main` sem passar por pull request. Commits seguem o padrão conventional commits em português.

Ciclo de vida de uma mudança:

1. Criar branch de trabalho a partir da `dev`
2. Desenvolver e commitar na branch
3. Abrir PR da branch de trabalho para a `dev`
4. Obter revisão e aprovação de um par
5. Mergear na `dev`
6. Quando a `dev` estiver estável, abrir PR da `dev` para a `main`

## Quem revisa e aprova PRs

| Ação | Quem pode |
|------|-----------|
| Abrir PR | Qualquer integrante da equipe |
| Revisar e aprovar PR | Qualquer par (integrante diferente do autor) |
| Mergear na `main` | Somente via PR aprovado vindo da `dev` |

Ver seção Governança no [README.md](../README.md).

## Aprovação mínima antes da integração

Todo PR exige no mínimo 1 aprovação de um par antes de ser mergeado. O autor do PR não pode aprovar o próprio PR.

## Política de push direto na branch principal

- Não se deve fazer push direto para `main`.
- Não se deve abrir PR de branches de trabalho diretamente para a `main`.
- A `main` só aceita merge vindo da `dev` via PR aprovado por um par.

## Template de PR com checklist mínimo

Todo PR aberto no repositório utiliza o template em [`.github/PULL_REQUEST_TEMPLATE.md`](../.github/PULL_REQUEST_TEMPLATE.md)
