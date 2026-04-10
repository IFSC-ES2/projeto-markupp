# Métricas

## Produto

### Cobertura de testes

- Classificação: produto
- Objetivo: quanto do código está coberto por testes
- Definição: funções chamadas por testes / total de funções (%)
- Fonte: ferramenta de cobertura no CI
- Frequência: a cada PR mergeado
- Responsável: Nicolas Pitz
- Interpretação: subindo é bom, caindo indica código novo sem teste

### Funcionalidades do MVP entregues vs planejadas

- Classificação: produto
- Objetivo: progresso real do produto
- Definição: funcionalidades concluídas / funcionalidades planejadas (%)
- Fonte: issues no GitHub
- Frequência: a cada marco
- Responsável: Nicolas Pitz
- Interpretação: abaixo de 70% indica necessidade de revisar escopo ou capacidade

### Critérios de aceitação atendidos

- Classificação: produto
- Objetivo: se o que foi entregue atende o que foi especificado
- Definição: critérios atendidos / total de critérios das issues concluídas (%)
- Fonte: issues no GitHub
- Frequência: a cada issue concluída
- Responsável: Nicolas Pitz
- Interpretação: abaixo de 100% indica entregas incompletas

## Processo

### Participação nos PRs

- Classificação: processo
- Objetivo: se todos estão colaborando nas revisões
- Definição: quantidade de reviews feitas por membro no período
- Fonte: pull requests no GitHub
- Frequência: semanal
- Responsável: Gabriela Riedel
- Interpretação: deve ser analisada considerando o papel e a disponibilidade de cada membro, concentração em uma pessoa indica gargalo

### PRs com revisão antes do merge

- Classificação: processo
- Objetivo: se o DoD está sendo seguido na parte de revisão
- Definição: PRs com ao menos uma aprovação antes do merge / total de PRs (%)
- Fonte: pull requests no GitHub
- Frequência: semanal
- Responsável: Gabriela Riedel
- Interpretação: ideal é 100%, abaixo disso indica PRs entrando sem revisão

### Distribuição de commits por membro

- Classificação: processo
- Objetivo: equilíbrio de contribuição na equipe
- Definição: commits por membro no período
- Fonte: histórico do git
- Frequência: semanal
- Responsável: Gabriela Riedel
- Interpretação: deve ser proporcional à disponibilidade e papel de cada membro, desequilíbrio desproporcional pode indicar sobrecarga ou ociosidade

## Projeto

### Burndown do backlog

- Classificação: projeto
- Objetivo: progresso da equipe ao longo do tempo
- Definição: itens restantes do backlog por semana (gráfico)
- Fonte: issues no GitHub
- Frequência: semanal
- Responsável: Renato Freitas e Gabriela Riedel
- Interpretação: linha descendo é ritmo saudável, estagnada ou subindo indica atraso ou aumento de escopo

### Marcos atingidos no prazo

- Classificação: projeto
- Objetivo: se a equipe está cumprindo as datas planejadas
- Definição: marcos entregues no prazo / total de marcos (%)
- Fonte: milestones no GitHub
- Frequência: a cada marco
- Responsável: Gabriela Riedel
- Interpretação: 100% indica planejamento realista, abaixo disso indica necessidade de ajustar escopo ou capacidade
