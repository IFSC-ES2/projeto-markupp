# Critérios de Qualidade do Projeto

**Projeto:** Markupp
**Versão:** 1.1
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

A tabela abaixo referencia os riscos formalmente registrados no documento de gestão de riscos do projeto, relacionando cada um aos atributos de qualidade que podem ser afetados.

| ID | Risco | Prob. / Impacto | Prioridade | Atributo(s) Afetado(s) | Impacto Potencial na Qualidade | Ação de Mitigação | Como a Mitigação Protege o Atributo |
|----|-------|----------------|------------|----------------------|-------------------------------|-------------------|-------------------------------------|
| RQP-1 | Alto índice de bugs encontrados em fase de homologação | Alta / Alto | **Alta** | Confiabilidade, Desempenho | Fluxos críticos instáveis e cobertura insuficiente comprometem a integridade dos documentos e a performance do sistema em produção | Manter cobertura ≥ 80% (DoD), fazer relatórios de testes e priorizar bugs críticos; abrir sprint de estabilização se necessário | A cobertura garantida pelo DoD e a revisão obrigatória por par impedem que código de baixa qualidade chegue à main, preservando a estabilidade dos fluxos core |
| RPR-1 | Backlog volumoso e poucas horas de trabalho semanais | Alta / Médio | **Alta** | Confiabilidade, Usabilidade | Pressão de prazo pode levar a testes negligenciados e funcionalidades de interface mal refinadas, degradando a experiência do usuário e a estabilidade do sistema | Sprints temáticas por categoria de entrega; repriorizar backlog no mínimo viável; renegociar escopo quando necessário | Sprints bem dimensionadas evitam que a pressão de tempo comprometa os atributos mais críticos do MVP, mantendo entregas focadas e testadas |
| REQ-2 | Curva de aprendizado das tecnologias escolhidas | Média / Alto | **Alta** | Desempenho, Compatibilidade | Implementações subótimas de integração e do pipeline semântico podem gerar latência elevada na API e inconsistências na troca de documentos com clientes externos | Agrupar tasks por tecnologia e usuário; estudos dirigidos; pareamento entre membros mais e menos experientes | Reduz a probabilidade de decisões técnicas inadequadas que impactam diretamente o desempenho da API e a interoperabilidade com clientes |
| RES-1 | Complexidade na criação/entrega da interface gráfica web | Média / Médio | **Média** | Usabilidade | Features mais complexas que o esperado podem resultar em fluxos de edição mal acabados, prejudicando a experiência central do produto | Validar mockups/protótipos antes da implementação; revisar requisitos ao fim de cada sprint; reduzir features secundárias se necessário | A validação antecipada de protótipos alinha a interface às expectativas dos usuários, evitando retrabalho e preservando a usabilidade nas entregas |
| RTE-1 | Falha de integração (obsolescência, funções depreciadas, sistemas sem conexão) | Baixa / Alto | **Média** | Compatibilidade, Confiabilidade | Quebra de integrações com clientes e agentes de IA, falhas silenciosas na troca de documentos e possível perda de dados em operações em andamento | Mapear dependências externas; fixar versões estáveis; prever alternativas para serviços críticos; isolar módulo afetado em caso de falha | A adoção de versões fixas e alternativas mapeadas mantém o contrato de API estável e evita que falhas externas corrompam o fluxo de documentos |
| REQ-1 | Membro da equipe ficar ausente (turnover) | Baixa / Alto | **Média** | Confiabilidade | Perda de conhecimento sobre partes críticas do sistema pode gerar inconsistências no versionamento e na lógica de negócio por redistribuição não planejada de tarefas | Cada membro documenta o andamento das suas tarefas; manter dois integrantes familiarizados com cada área; redistribuir e revisar prazos se necessário | A documentação contínua e o conhecimento compartilhado reduzem o impacto de ausências na estabilidade e integridade das funcionalidades core |

> **Escala de referência:** Probabilidade — Baixa / Média / Alta. Impacto — Baixo / Médio / Alto. IDs conforme riscos.md do projeto.

---

### 2.2 Observações

- **RQP-1** é o risco de maior exposição à qualidade do produto: combina alta probabilidade, alto impacto e afeta dois atributos prioritários simultaneamente (Confiabilidade e Desempenho). O DoD da equipe — cobertura ≥ 80% e aprovação por par — é a principal salvaguarda e deve ser tratado como inegociável.
- **RPR-1** e **REQ-2** têm caráter estrutural: são condições permanentes do projeto (restrição de horas semanais e aprendizado de novas tecnologias) que pressionam os atributos de forma contínua, exigindo monitoramento ativo a cada sprint.
- Os riscos **RTE-1** e **REQ-1**, embora de baixa probabilidade, têm impacto alto e afetam atributos centrais (Compatibilidade e Confiabilidade), justificando a manutenção das ações preventivas mesmo sem sinais imediatos de ocorrência.

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
| Usabilidade | Usabilidade | Facilidade de uso nos fluxos de criação, edição e organização de documentos; validação dos protótipos da interface (RES-1) | Score no System Usability Scale (SUS) em sessões de teste com usuários; taxa de conclusão de tarefas sem auxílio | SUS ≥ 70; taxa de conclusão ≥ 80% |
| Compatibilidade | Compatibilidade | Interoperabilidade da API REST com clientes externos e agentes de IA; suporte a encodings; estabilidade frente a atualizações de dependências (RTE-1) | Resultados de testes de integração automatizados; validação de formatos (UTF-8, ASCII); build estável com versões fixadas | 100% dos contratos de API validados; zero falhas de encoding; zero quebras por dependência depreciada |
| Confiabilidade | Confiabilidade | Integridade do versionamento de documentos; estabilidade dos fluxos CRUD; prevenção de regressões (RQP-1) | Cobertura de testes unitários (meta do DoD); taxa de erro em fluxos críticos; relatórios de bugs em homologação | Cobertura ≥ 80% (DoD); taxa de erro < 1% no CRUD; zero bugs críticos abertos ao fim de cada sprint |
| Desempenho | Eficiência de Desempenho | Tempo de resposta da API nas operações de leitura, escrita e busca semântica; impacto de implementações da curva de aprendizado (REQ-2) | Tempo de resposta via testes de carga (ex: k6); latência das operações de busca semântica; revisão de código com foco em otimização nas PRs | p95 < 2s para operações CRUD; p95 < 5s para buscas semânticas |
