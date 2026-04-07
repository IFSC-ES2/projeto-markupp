# Git Flow

## Branches

- main: branch de produção, só recebe merges da dev
- dev: branch de desenvolvimento, todas as features são mergeadas aqui primeiro
- docs/, feature/, fix/, etc: branches de trabalho criadas a partir da dev

## Fluxo

1. Criar branch a partir da dev
2. Desenvolver na branch
3. Abrir PR para a dev
4. Revisar e aprovar por ao menos um par
5. Mergear na dev
6. Quando a dev estiver estável, abrir PR da dev para a main

## Regras

- Nunca abrir PR direto para a main a partir de branches de trabalho
- A main só recebe merge da dev
- Todo PR precisa de ao menos uma aprovação
- Commits seguem o padrão conventional commits em português
