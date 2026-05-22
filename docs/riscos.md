# Riscos
Este documento tem o objetivo de esclarecer possíveis problemas e riscos para mitigarmos ao logo do desenvolvimento do projeto Markupp.

Vamos separar os tópicos em:
- **1. Registro Inicial de Riscos do Projeto;**
- **2. Análise e Priorização dos Riscos;**
- **3. Plano de Resposta aos Riscos.**

## 1 Registro Inicial de Riscos do Projeto
### 1.1 Escopo
| ID | Descrição | Causa | Consequência/Impacto Esperado | Probabilidade | Impacto | Prioridade |  Estratégia de Mitigação |  Responsável pelo Acompanhamento |
| --- | --- | --- | --- | --- | --- | --- | --- | --- |
| RES-1 | Complexidade na criação/entrega da interface via plugin do Obsidian | Features mais profundas e complexas do que o esperado. | Grande volume de Tasks e possível atraso do prazo | Média | Médio | Média | Otimizar requisitos e criar mockup/protótipo para validação da interface gráfica. | Nicolas Pitz |

### 1.2 Prazo
| ID | Descrição | Causa | Consequência/Impacto Esperado | Probabilidade | Impacto | Prioridade |  Estratégia de Mitigação |  Responsável pelo Acompanhamento |
| --- | --- | --- | --- | --- | --- | --- | --- | --- |
| RPR-1 | Backlog volumoso e poucas horas de trabalho semanais. | O projeto tem tarefas abrangentemente complexas e os desenvolvedores responsáveis tem poucas horas de trabalho semanal em conjunto pois todos trabalham e estudam durante a semana. | Atraso na entrega de Tasks, possível atraso do prazo e qualidade afetada | Alta | Médio | Alta | Gerenciamento de Sprints por categoria de entrega (Ex: Sprint-1 - Foco na Modelagem) | Gabriela |

### 1.3 Equipe
| ID | Descrição | Causa | Consequência/Impacto Esperado | Probabilidade | Impacto | Prioridade |  Estratégia de Mitigação |  Responsável pelo Acompanhamento |
| --- | --- | --- | --- | --- | --- | --- | --- | --- |
| REQ-1 | Membro da Equipe ficar ausente do projeto (Turnover) | Situações que podem ocorrer com o discente: Trancar o curso, cancelar a disciplina, se ausentar por injurías/doenças  | Atraso na entrega de Tasks, possível atraso do prazo e qualidade afetada | Baixa | Alto | Média | Cada desenvolvedor documentar o andamento de suas tarefas | Renato e Gabriela |
| REQ-2 | Curva de Aprendizado das tecnologias escolhidas | Foram selecionadas tecnologias que necessitam que o desenvolvedor se adapte | Atraso na entrega de Tasks, possível atraso do prazo e qualidade afetada | Média | Alto | Alta | Agrupar tasks por tecnologia/usuário | Nicolas Arthur |

### 1.4 Tecnologia
| ID | Descrição | Causa | Consequência/Impacto Esperado | Probabilidade | Impacto | Prioridade |  Estratégia de Mitigação |  Responsável pelo Acompanhamento |
| --- | --- | --- | --- | --- | --- | --- | --- | --- |
| RTE-1 | Falha de Integração | Obslecência de tecnologias, funções deprecateds, sistemas sem conexão. | Atraso na entrega de Tasks, possível atraso do prazo e qualidade afetada | Baixa | Alto | Média | Ter alternativas para serviços sem conexão | Nicolas Arthur |

### 1.5 Qualidade/Processo
| ID | Descrição | Causa | Consequência/Impacto Esperado | Probabilidade | Impacto | Prioridade |  Estratégia de Mitigação |  Responsável pelo Acompanhamento |
| --- | --- | --- | --- | --- | --- | --- | --- | --- |
| RQP-1 | Alto indíce de bugs encontrados em fase de homologação | Testes mal-elaborados e fluxos sem testes | Qualidade do sistema | Alta | Alto | Alta | Manter a cobertura de testes acima de 80%, fazer relatórios de testes e priorizar bugs críticos | Nicolas Pitz |

## 2 Análise e Priorização dos Riscos

### 2.1 Critérios adotados

Para classificar cada risco, utilizamos uma escala qualitativa de três níveis para **probabilidade** e **impacto**, combinadas para definir a **prioridade**.

**Probabilidade** (chance do risco se concretizar ao longo do projeto):
- **Baixa**: evento pouco provável, sem indícios atuais de que possa ocorrer.
- **Média**: evento plausível, com algum indício ou histórico em projetos semelhantes.
- **Alta**: evento muito provável ou já observado em fases anteriores.

**Impacto** (efeito sobre prazo, escopo ou qualidade caso o risco ocorra):
- **Baixo**: pequeno atraso ou retrabalho pontual, sem afetar entregas.
- **Médio**: atraso de uma sprint ou necessidade de replanejar tarefas.
- **Alto**: compromete uma entrega importante, a qualidade do produto ou a continuidade do projeto.

### 2.2 Matriz de Riscos (Probabilidade x Impacto)

|                    | Impacto Baixo | Impacto Médio         | Impacto Alto              |
| ------------------ | ------------- | --------------------- | ------------------------- |
| **Probabilidade Alta**  | -             | RPR-1                 | RQP-1                     |
| **Probabilidade Média** | -             | RES-1                 | REQ-2                     |
| **Probabilidade Baixa** | -             | -                     | REQ-1, RTE-1              |

### 2.3 Tabela de Priorização

| Prioridade | Riscos                  | Justificativa |
| ---------- | ----------------------- | ------------- |
| **Alta**   | RQP-1                   | Combina alta probabilidade e alto impacto, sendo o risco mais crítico: afeta diretamente a qualidade entregue. |
| **Alta**   | RPR-1                   | Condição estrutural do time (todos trabalham e estudam) torna a probabilidade alta; mesmo com impacto médio, a recorrência justifica alta prioridade. |
| **Alta**   | REQ-2                   | Probabilidade média e impacto alto: afeta a velocidade de todo o time e já aparece nas primeiras sprints. |
| **Média**  | REQ-1, RTE-1            | Probabilidade baixa, porém impacto alto: se ocorrerem comprometem entregas inteiras, mas a baixa frequência reduz a prioridade geral. |
| **Média**  | RES-1                   | Probabilidade e impacto médios: afeta o ritmo do projeto, mas pode ser absorvido com replanejamento. |

### 2.4 Riscos mais críticos no momento

- **RQP-1: Alto índice de bugs em homologação**: é o único risco com probabilidade e impacto altos; compromete diretamente a qualidade entregue ao cliente.
- **RPR-1: Backlog volumoso e poucas horas semanais**: condição permanente da equipe, com alta probabilidade de recorrência a cada sprint.
- **REQ-2: Curva de aprendizado das tecnologias**: já observado nas primeiras sprints, afeta o ritmo de toda a equipe.

## 3 Plano de Resposta aos Riscos

### 3.1 Ações preventivas

Ações contínuas para reduzir a probabilidade ou o impacto antes que o risco se concretize:

| ID    | Ações Preventivas |
| ----- | ----------------- |
| RES-1 | Validar mockups/protótipos com o cliente antes de iniciar a implementação; revisar requisitos ao fim de cada sprint. |
| RPR-1 | Planejar sprints curtas e temáticas; priorizar o backlog por valor; revisar capacidade real da equipe a cada sprint. |
| REQ-1 | Documentar o andamento de cada tarefa no board; versionar código frequentemente; manter pelo menos dois integrantes familiarizados com cada área do sistema. |
| REQ-2 | Estudos dirigidos nas tecnologias escolhidas; pareamento entre quem tem mais e menos experiência; agrupar tasks por tecnologia. |
| RTE-1 | Mapear dependências externas; fixar versões estáveis; prever alternativas para serviços críticos. |
| RQP-1 | Escrever testes junto com a feature; manter cobertura acima de 80%; revisar PRs com checklist de testes. |

### 3.2 Ações de contingência (caso o risco se concretize)

| ID    | Ações de Contingência |
| ----- | --------------------- |
| RES-1 | Renegociar escopo da sprint com o cliente; reduzir features secundárias da interface. |
| RPR-1 | Repriorizar backlog focando no mínimo viável; remanejar tasks menos críticas para sprints futuras. |
| REQ-1 | Redistribuir tarefas do ausente entre o restante do time; acionar backup documentado; revisar prazos impactados. |
| REQ-2 | Promover sessão de mentoria interna; substituir biblioteca/tecnologia por alternativa mais simples se necessário. |
| RTE-1 | Ativar alternativa mapeada; isolar o módulo afetado; postergar integração até estabilização. |
| RQP-1 | Congelar novas features e abrir sprint de estabilização focada em correção de bugs críticos. |

### 3.3 Acompanhamento ao longo do projeto

- **Revisão quinzenal dos riscos**: ao final de cada sprint, o responsável de cada risco revisa status, probabilidade e impacto, registrando alterações neste documento.
- **Reunião de retrospectiva**: novos riscos identificados durante a sprint são adicionados ao registro, e riscos extintos são marcados como encerrados.
- **Indicadores monitorados**:
  - Velocidade da equipe (burn-down): sinaliza RPR-1 e REQ-2.
  - Cobertura de testes e número de bugs em homologação: sinaliza RQP-1.
  - Presença e atividade dos membros no board: sinaliza REQ-1.
  - Falhas em builds/integração: sinaliza RTE-1.
- **Escalonamento**: riscos que migrem para a faixa Alta/Alta na matriz passam a ser discutidos em toda reunião semanal, até retornarem a um patamar aceitável.
