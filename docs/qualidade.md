# Critérios de Qualidade do Projeto

## Atributos Priorizados

- Manutenibilidade
- Compatibilidade
- Confiabilidade 
- Desempenho

### Justificativa

**Manutenibilidade**

O projeto tem um escopo que tende a crescer, aparecer novos requisitos e aumentar de escala. Evitar débitos técnicos em um projeto desse tipo é essencial.

**Compatibilidade**

O sistema tem a premissa de ser self host, isso faz ser importante alta compatibilidade entre diferentes ambiente com docker compose, além disso o editor de texto deve ser utilizável em diferentes ambientes

**Confiabilidade**

Trabalhando com armazenamento é importante ter confiabilidade para não acontecer problemas sérios como perda ou corrupção de dados, modificações perdidas, retrabalhos para os usuários

**Desempenho**

Para tornar a aplicação portável e usável em diferentes plataformas, tanto o servidor quanto a interface cliente devem ter bom desempenho evitando uso desnecessário de recursos

## Orientação para as Próximas Etapas

Cada atributo priorizado guia decisões nas sprints seguintes:

- Manutenibilidade: exige cobertura mínima de testes no DoD, revisão de código em PR e aplicação de boas práticas de código limpo.
- Compatibilidade: exige que a interface do cliente (editor de texto Markdown) funcione bem em diferentes ambientes e navegadores, e que o servidor self host rode sem problemas em diferentes sistemas operacionais.
- Confiabilidade: exige testes junto com as features, tratamento de erros no CRUD, integridade dos documentos e correção de bugs críticos antes de novas entregas.
- Desempenho: exige boas escolhas de recursos, atenção a memory leaks e cuidado para evitar más práticas que pesam sem necessidade no cliente e no servidor.

## Relação entre Riscos e Qualidade

### Mapeamento de Riscos × Atributos de Qualidade

Relação dos [riscos](docs/riscos.md) com os atributos afetados:

- RES-1 (complexidade da interface): afeta Manutenibilidade e Desempenho. Interface complexa gera código acoplado e componentes pesados.
- RPR-1 (backlog volumoso com poucas horas semanais): afeta Manutenibilidade e Confiabilidade. Pressão por entrega estimula atalhos e redução de testes.
- REQ-1 (turnover de membros): afeta Manutenibilidade e Confiabilidade. Saída de integrantes deixa áreas sem conhecimento documentado.
- REQ-2 (curva de aprendizado): afeta Manutenibilidade e Desempenho. Falta de domínio gera código pouco padronizado e soluções ineficientes.
- RTE-1 (falha de integração): afeta Compatibilidade e Confiabilidade. Dependências depreciadas quebram contratos da API e fluxos críticos.
- RQP-1 (bugs em homologação): afeta Confiabilidade e Manutenibilidade. Bugs frequentes indicam baixa cobertura de testes e código difícil de evoluir.

### Como as Mitigações Protegem os Atributos

Ações concretas que a equipe adota para proteger cada atributo:

- **Manutenibilidade**: linter e formatter rodando em CI/pre-commit; PR template com checklist obrigatório de testes e boas práticas.
- **Compatibilidade**: teste manual do editor em Chrome e Firefox antes de cada entrega.
- **Confiabilidade**: validação de entrada; testes de integração obrigatórios para os fluxos CRUD; confirmação no frontend antes de ações destrutivas.
- **Desempenho**: seguir boas práticas de código como evitar operações pesadas em eventos frequentes; evitar dependências pesadas sem necessidade; atenção a esses pontos durante o code review.

## Avaliação Preliminar da Qualidade

### Escopo Inicial

- Verificar se os fluxos CRUD de documentos funcionam sem perda de dados.
- Verificar se o editor de texto Markdown funciona bem em diferentes navegadores e sistemas operacionais.
- Verificar se o servidor self host roda sem problemas em diferentes ambientes.
- Observar se a aplicação se mantém leve no uso, sem consumo excessivo de memória ou CPU.
- Acompanhar a facilidade de adicionar novos requisitos ao longo das sprints.

### Métricas e Evidências Preliminares

**Manutenibilidade**

- Cobertura de testes unitários acima de 80%.
- 100% das PRs revisadas por pelo menos 1 pessoa.

**Compatibilidade**

- Editor funcionando nos principais navegadores.
- Servidor self host subindo em diferentes sistemas operacionais sem ajustes manuais.
- Build estável com versões fixadas das dependências.

**Confiabilidade**

- Cobertura de testes unitários acima de 80%.
- Acompanhar erros nos fluxos de CRUD.
- Zero bugs críticos abertos ao fim de cada sprint.
- Relatórios de bugs encontrados em formato de issues no repositório.

**Desempenho**

- Ausência de memory leaks perceptíveis.
- Uso de CPU e memória do servidor dentro de faixas razoáveis em operação normal.
- Escolha consciente de bibliotecas e recursos, evitando dependências pesadas sem necessidade.
- Interface respondendo rápido nas ações comuns do usuário.
