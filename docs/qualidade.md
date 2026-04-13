# Critérios de Qualidade do Projeto

**Projeto:** Markupp
**Versão:** 1.0
**Data:** 16/04/2026
**Equipe:** Renato Freitas (Arquiteto de Software), Nícolas Arthur (DevOps/Infra), Nicolas Pitz (Engenheiro de Qualidade), Gabriela Riedel (Scrum Master)

---

## 1. Definição Inicial de Critérios de Qualidade

### 1.1 Atributos de Qualidade Prioritários

Foram selecionados os quatro atributos de qualidade prioritários para este projeto, com base no contexto do MVP e nas necessidades dos stakeholders:

| # | Atributo | Categoria ISO/IEC 25010 |
|---|----------|------------------------|
| 1 | Usabilidade | Usabilidade |
| 2 | Compatibilidade | Compatibilidade |
| 3 | Confiabilidade | Confiabilidade |
| 4 | Desempenho | Eficiência de Desempenho |

> **Referência:** Os atributos acima são definidos conforme a norma ISO/IEC 25010, que estrutura a qualidade de produto de software em 8 características: Adequação Funcional, Eficiência de Desempenho, Compatibilidade, Usabilidade, Confiabilidade, Segurança, Manutenibilidade e Portabilidade.

---

### 1.2 Justificativa da Relevância para o MVP

**Atributo 1 — Usabilidade**
O Markupp tem como público-alvo usuários que já utilizam agentes de IA no dia a dia e times que precisam centralizar documentos. Para esse perfil, a interface de criação, edição e organização de documentos Markdown precisa ser intuitiva e eficiente, minimizando a curva de aprendizado. Uma experiência de uso fluida é condição direta para a adoção da plataforma e para que a colaboração entre usuário e IA ocorra sem fricção.

**Atributo 2 — Compatibilidade**
O sistema expõe uma API REST e precisa integrar-se a múltiplos clientes, agentes de IA e ferramentas externas. A compatibilidade garante que o Markupp consiga receber e enviar documentos corretamente em diferentes formatos de texto (incluindo variações de codificação como UTF-8 e ASCII), além de interoperar com os agentes e clientes que consumirão a API no MVP.

**Atributo 3 — Confiabilidade**
Como hub central de conhecimento, o Markupp precisa garantir que os documentos criados, editados e versionados sejam armazenados e recuperados com integridade. Falhas na confiabilidade — como perda de versões, corrompimento de dados ou indisponibilidade do serviço — comprometeriam diretamente a proposta de valor do produto, que é ser a fonte única e confiável de verdade para o conhecimento do usuário.

**Atributo 4 — Desempenho**
O MVP contempla processamento automático de documentos para busca semântica e operações em tempo real entre usuário e IA. Tempos de resposta elevados degradariam a experiência colaborativa e a percepção de valor da plataforma, especialmente em operações de busca e edição simultânea.

---

### 1.3 Como os Atributos Orientarão Decisões nas Próximas Etapas

| Atributo | Decisões que orientará |
|----------|----------------------|
| Usabilidade | Design da interface de edição Markdown, definição de critérios de aceite nas histórias de usuário relacionadas à navegação e organização hierárquica de documentos, planejamento de testes com usuários |
| Compatibilidade | Definição dos formatos e encodings suportados pela API REST, estratégia de testes de integração com clientes e agentes de IA externos, especificação dos contratos de API |
| Confiabilidade | Estratégia de controle de versões dos documentos, política de tratamento de erros e rollback, plano de testes de regressão e cobertura mínima de 80% definida no DoD |
| Desempenho | Escolha de arquitetura para busca semântica, definição de SLAs da API, estratégia de indexação e otimização de queries |

---

## 2. Relação entre Riscos e Qualidade

### 2.1 Mapeamento de Riscos × Atributos de Qualidade

| Risco Relevante | Probabilidade / Impacto | Atributo(s) Afetado(s) | Impacto Potencial na Qualidade | Ação de Mitigação | Como a Mitigação Protege o Atributo |
|----------------|------------------------|----------------------|-------------------------------|-------------------|-------------------------------------|
| Complexidade do processamento semântico sobrecarregar a API em tempo real | Média / Alto | Desempenho, Confiabilidade | Latência elevada nas operações de busca e edição, comprometendo a colaboração usuário-IA em tempo real | Processar indexação semântica de forma assíncrona, separando o fluxo crítico de edição do pipeline de busca | Garante que operações de escrita e leitura permaneçam responsivas mesmo durante reindexação de documentos |
| Integração inconsistente entre clientes externos e a API REST | Alta / Médio | Compatibilidade, Confiabilidade | Falhas silenciosas na troca de documentos entre o Markupp e agentes de IA ou ferramentas terceiras | Definir e versionar contratos de API com especificação OpenAPI, implementar testes de integração automatizados | Assegura que mudanças na API não quebrem integrações existentes e que os formatos aceitos sejam validados de forma sistemática |
| Equipe reduzida com acúmulo de funções | Média / Alto | Confiabilidade, Desempenho | Dívida técnica acumulada, cobertura de testes abaixo da meta de 80% definida no DoD, bugs em produção | Enforçar o DoD (cobertura ≥ 80%, aprovação por par via PR) e adotar conventional commits para rastreabilidade | Mantém o padrão de qualidade independentemente da pressão de prazo, garantindo que entregas com risco elevado não sejam mescladas sem revisão |
| Perda ou corrompimento de versões de documentos | Baixa / Alto | Confiabilidade | Impossibilidade de recuperar histórico de documentos, violando a proposta central do produto como hub de conhecimento | Implementar testes automatizados de integridade do sistema de versionamento e política de backup dos dados | Garante que o controle de versões — funcionalidade core do Markupp — opere sem falhas de integridade |
| Curva de aprendizado elevada para novos usuários na edição Markdown | Alta / Médio | Usabilidade | Baixa adoção do produto pelo público-alvo, especialmente usuários menos familiarizados com Markdown | Realizar ao menos uma rodada de testes com usuários antes da entrega do MVP e iterar sobre o fluxo de edição | Alinha a interface às expectativas reais dos usuários, reduzindo abandonos e aumentando a percepção de valor |

> **Escala de referência:** Probabilidade — Baixa / Média / Alta. Impacto — Baixo / Médio / Alto.

---

### 2.2 Observações

- Os riscos de integração com clientes externos e de corrompimento de versões afetam diretamente a Compatibilidade e a Confiabilidade simultaneamente, tornando esses dois atributos os de maior exposição no MVP.
- O DoD já definido pela equipe (cobertura ≥ 80% e aprovação por par) atua como mecanismo de mitigação transversal, protegendo especialmente a Confiabilidade e o Desempenho ao impedir que código de baixa qualidade chegue à main.

---

## 3. Definição Preliminar de Avaliação da Qualidade

### 3.1 Escopo Inicial da Avaliação

A avaliação de qualidade nesta fase inicial abrangerá:

- **Funcionalidades core do MVP:** criação, edição, organização hierárquica e versionamento de documentos Markdown; operações CRUD via API REST; busca semântica básica;
- **Atributos selecionados:** avaliação focada nos 4 atributos definidos na Seção 5;
- **Ambientes cobertos:** aplicação self-hosted em ambiente desktop, consumida via API REST por clientes e agentes de IA;
- **O que está fora do escopo nesta fase:** avaliação de acessibilidade avançada, portabilidade para ambientes mobile, compatibilidade com navegadores legados.

A avaliação poderá ser expandida em versões posteriores ao MVP, incorporando características adicionais da ISO/IEC 25010, como Segurança e Manutenibilidade.

---

### 3.2 Modelo de Qualidade Adotado: ISO/IEC 25010

Este projeto adota a **ISO/IEC 25010** como modelo de referência para definição e avaliação da qualidade do produto de software.

A norma organiza a qualidade do produto em 8 características principais, das quais este projeto prioriza as seguintes:

| Característica ISO 25010 | Relevante para este projeto? | Atributo mapeado |
|--------------------------|-----------------------------|--------------------|
| Adequação Funcional | Não | — |
| Eficiência de Desempenho | Sim | Desempenho |
| Compatibilidade | Sim | Compatibilidade |
| Usabilidade | Sim | Usabilidade |
| Confiabilidade | Sim | Confiabilidade |
| Segurança | Não | — |
| Manutenibilidade | Não | — |
| Portabilidade | Não | — |

---

### 3.3 Métricas e Evidências Preliminares

| Atributo | Característica ISO 25010 | O que será avaliado | Métrica / Evidência Preliminar | Meta Inicial |
|----------|--------------------------|--------------------|---------------------------------|--------------|
| Usabilidade | Usabilidade | Facilidade de uso dos fluxos de criação, edição e organização de documentos | Score no System Usability Scale (SUS) em sessões de teste com usuários; taxa de conclusão de tarefas sem auxílio | SUS ≥ 70; taxa de conclusão ≥ 80% |
| Compatibilidade | Compatibilidade | Interoperabilidade da API REST com clientes externos e agentes de IA; suporte a encodings de texto | Resultados de testes de integração automatizados com diferentes clientes; validação de formatos (UTF-8, ASCII) | 100% dos contratos de API validados; zero falhas de encoding em testes de integração |
| Confiabilidade | Confiabilidade | Integridade do versionamento de documentos; estabilidade dos fluxos críticos (CRUD) | Cobertura de testes unitários (meta do DoD); taxa de erro em fluxos críticos; testes de integridade do versionamento | Cobertura ≥ 80% (conforme DoD); taxa de erro < 1% nos fluxos de CRUD |
| Desempenho | Eficiência de Desempenho | Tempo de resposta da API nas operações de leitura, escrita e busca semântica | Tempo de resposta medido via testes de carga (ex: k6); latência de operações de busca semântica | p95 < 2s para operações CRUD; p95 < 5s para buscas semânticas |
