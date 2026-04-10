# Planejamento Inicial e Baseline

> **Projeto:** Markupp
> **Versão:** 1.0
> **Data de Registro da Linha de Base:** 09/04/2026

---

## (a) Recorte do Backlog

Lista completa de issues do repositório que servem de base para este planejamento.

| ID  | Item                                                               | Tipo    | Milestone |
|-----|--------------------------------------------------------------------|---------|-----------|
| #16 | Implementar Endpoint GET /files/{id} para Recuperação de Conteúdo | Task    | RF1       |
| #15 | Implementar Endpoint POST /files para Upload de Markdown           | Task    | RF1       |
| #14 | Configurar Provedor de Blob Storage (S3-compatible/Docker)         | Task    | RF1       |
| #13 | Definir rotas da API do backend                                    | Feature | —         |
| #11 | [INIT][3] Ferramentas Iniciais                                     | Feature | Init      |
| #10 | [INIT][2] Arquitetura Frontend                                     | Feature | Init      |
| #9  | [INIT][1] Arquitetura Backend                                      | Feature | Init      |
| #8  | [INIT] Versão inicial do projeto (Scaffolding)                     | Feature | Init      |
| #7  | [RF5] Edição de Conteúdo Markdown                                  | Feature | RF5       |
| #6  | [RF4] Renomear Arquivos via Frontend                               | Feature | RF4       |
| #5  | [RF3] Exclusão de Arquivos via Frontend                            | Feature | RF3       |
| #2  | [RF1] Implementação de Armazenamento de Arquivos Markdown          | Feature | RF1       |
| #1  | [RF2] Criação de Arquivos via Frontend                             | Feature | RF2       |

---

## (b) Priorização dos Itens do MVP

Priorizamos os Requisitos Funcionais (RF1 a RF5) como itens do MVP, por representarem o núcleo de valor entregável ao usuário final.

| Prioridade | ID  | Item                                                      | Justificativa                                          |
|------------|-----|-----------------------------------------------------------|--------------------------------------------------------|
| Alta       | #2  | [RF1] Implementação de Armazenamento de Arquivos Markdown | Base para todas as demais funcionalidades do sistema   |
| Alta       | #1  | [RF2] Criação de Arquivos via Frontend                    | Funcionalidade central da experiência do usuário       |
| Alta       | #5  | [RF3] Exclusão de Arquivos via Frontend                   | Operação essencial para gerenciamento do ciclo de vida |
| Média      | #6  | [RF4] Renomear Arquivos via Frontend                      | Melhora a usabilidade e organização dos arquivos       |
| Média      | #7  | [RF5] Edição de Conteúdo Markdown                         | Completa o fluxo CRUD e agrega valor direto ao usuário |

> **Critérios de priorização:** valor de negócio, dependências técnicas entre RFs e viabilidade de entrega no horizonte atual do projeto.

---

## (c) Estimativas dos Itens Priorizados

Vamos definir as estimativas via T-Shirt Sizing em uma sessão de refinamento antes do início da Sprint 1.

| ID  | Item                                                      | Tamanho (T-Shirt) |
|-----|-----------------------------------------------------------|-------------------|
| #2  | [RF1] Implementação de Armazenamento de Arquivos Markdown | A preencher       |
| #1  | [RF2] Criação de Arquivos via Frontend                    | A preencher       |
| #5  | [RF3] Exclusão de Arquivos via Frontend                   | A preencher       |
| #6  | [RF4] Renomear Arquivos via Frontend                      | A preencher       |
| #7  | [RF5] Edição de Conteúdo Markdown                         | A preencher       |

> Após a sessão de refinamento descrita na seção (d), atualizaremos este documento com os tamanhos definidos, gerando uma nova versão da linha de base.

---

## (d) Técnica de Estimativa Adotada

**Técnica:** T-Shirt Sizing

**Descrição:** Classificamos cada item do backlog em tamanhos relativos — XS, S, M, L, XL — com base na complexidade percebida e no esforço estimado. Cada integrante vota individualmente e, em caso de divergência, discutimos até chegar a um consenso.

| Tamanho | Descrição orientativa          |
|---------|-------------------------------|
| XS      | Tarefa trivial, poucas horas  |
| S       | Tarefa simples, menos de 1 dia |
| M       | Esforço moderado, 1 a 2 dias  |
| L       | Tarefa complexa, vários dias  |
| XL      | Épico, requer decomposição    |

---

## (e) Hipóteses Assumidas

Premissas que assumimos ao construir este planejamento. Caso se mostrem inválidas, as estimativas e a previsão podem precisar de revisão.

- [ ] Todos os integrantes da equipe permanecerão ativos até o próximo marco
- [ ] O escopo do MVP (RF1 a RF5) não sofrerá mudanças significativas durante o período
- [ ] A stack definida na etapa anterior será mantida
- [ ] O backlog atual está suficientemente refinado para permitir estimativas iniciais
- [ ] O ambiente de desenvolvimento (local e/ou Docker) estará configurado no início da sprint
- [ ] A infraestrutura de Blob Storage (S3-compatible) estará disponível para testes
- [ ] As dependências entre RFs seguem a ordem RF1 → RF2 → RF3 → RF4 → RF5
- [ ] Não haverá dependências externas críticas bloqueando o avanço do projeto

---

## (f) Capacidade Planejada da Equipe

Nossa disponibilidade até o próximo marco do projeto — **Entrega 3, em 09/04/2026**.

| Membro          | Papel                   | Disponibilidade             |
|-----------------|-------------------------|-----------------------------|
| Renato Freitas  | Arquiteto de Software   | Horas de aula + 2h semanais |
| Nícolas Arthur  | DevOps / Infra          | Horas de aula + finais de semana |
| Gabriela Riedel | Scrum Master            | Horas de aula + 2h semanais |
| Nícolas Pitz    | Engenheiro de Qualidade | Horas de aula + finais de semana |

> **Período coberto por esta linha de base:** início do semestre até 09/04/2026
> **Próximo marco:** Entrega 3 — Estimativas e Métricas (Baseline)

---

## (g) Previsão Inicial do que se Espera Concluir no Período

No horizonte coberto por esta linha de base da abertura do projeto até a Entrega 3 (09/04/2026), ainda nos encontramos na **fase de planejamento e documentação**. Nenhum item de implementação foi concluído até o momento, uma vez que as issues de inicialização (#8 a #11) ainda estão em aberto.

A tabela abaixo registra nossa previsão inicial para este período, considerando a capacidade limitada e o volume total do backlog (13 issues). O conjunto completo evidencia que **não é viável concluir todos os itens no período atual**, sendo o foco desta entrega o estabelecimento da linha de base para as sprints seguintes.

| ID  | Item                                                               | Tipo    | Previsão para o Período      |
|-----|--------------------------------------------------------------------|---------|------------------------------|
| #8  | [INIT] Versão inicial do projeto (Scaffolding)                     | Feature | ⚠️ Em andamento              |
| #9  | [INIT][1] Arquitetura Backend                                      | Feature | ⚠️ Em andamento              |
| #10 | [INIT][2] Arquitetura Frontend                                     | Feature | ⚠️ Em andamento              |
| #11 | [INIT][3] Ferramentas Iniciais                                     | Feature | ⚠️ Em andamento              |
| #13 | Definir rotas da API do backend                                    | Feature | ⚠️ Em andamento              |
| #14 | Configurar Provedor de Blob Storage (S3-compatible/Docker)         | Task    | ❌ Não iniciado               |
| #2  | [RF1] Implementação de Armazenamento de Arquivos Markdown          | Feature | ❌ Não iniciado               |
| #15 | Implementar Endpoint POST /files para Upload de Markdown           | Task    | ❌ Não iniciado               |
| #16 | Implementar Endpoint GET /files/{id} para Recuperação de Conteúdo  | Task    | ❌ Não iniciado               |
| #1  | [RF2] Criação de Arquivos via Frontend                             | Feature | ❌ Não iniciado               |
| #5  | [RF3] Exclusão de Arquivos via Frontend                            | Feature | ❌ Não iniciado               |
| #6  | [RF4] Renomear Arquivos via Frontend                               | Feature | ❌ Não iniciado               |
| #7  | [RF5] Edição de Conteúdo Markdown                                  | Feature | ❌ Não iniciado               |

**Legenda:**
- ⚠️ **Em andamento** — iniciado no período, conclusão prevista para a Sprint 1
- ❌ **Não iniciado** — fora do escopo deste período, previsto para sprints futuras

> **Observação:** revisaremos este baseline ao final da Sprint 1, quando as estimativas da seção (c) estiverem definidas e o progresso real puder ser comparado com esta previsão inicial.

---

## (h) Data de Registro da Linha de Base

| Campo                         | Valor                    |
|-------------------------------|--------------------------|
| **Data de registro**          | 09/04/2026               |
| **Responsável pelo registro** | Gabriela Riedel          |
| **Versão do documento**       | 1.0                      |
| **Aprovado por**              | Adriano Lima / Professor |

---

*Este documento representa a linha de base do planejamento do projeto Markupp, registrada na Entrega 3. Usaremos ele como referência para acompanhamento, controle e comparação com o progresso real ao longo das sprints seguintes.*
