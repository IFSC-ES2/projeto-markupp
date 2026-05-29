# Avaliação - Engenharia de Software II

| entrega | aluno                           | commit  | data     | correção | nota | peso |
| ------- | ------------------------------- | ------- | -------- | -------- | ---- | ---- |
| 1       | equipe                          | 0d60ee0 | 16/03/26 | 20/03/26 | 9,5  | 2    |
| 2       | equipe                          | 7242cb2 | 28/03/26 | 29/03/26 | 7,7  | 2    |
| 3       | equipe                          | 6e380d2 | 11/04/26 | 22/04/26 | 9,9  | 3    |
| 4       | equipe                          | 76b5943 | 18/04/26 | 09/05/26 | 9,9  | 3    |
| 5       | Gabriela Riedel                 | 29c6e1e | 07/05/26 | 24/05/26 | 7,2  | 10   |
| 5       | Luiz Renato Freitas de Almeida  | 29c6e1e | 07/05/26 | 24/05/26 | 8,5  | 10   |
| 5       | Nícolas Arthur Raulino Oliveira | 29c6e1e | 07/05/26 | 24/05/26 | 8,1  | 10   |
| 5       | Nicolas Pitz                    | 29c6e1e | 07/05/26 | 24/05/26 | 7,9  | 10   |
| 6       | Gabriela Riedel                 | 2653ad6 | 21/05/26 | 28/05/26 | 6,8  | 10   |
| 6       | Luiz Renato Freitas de Almeida  | 2653ad6 | 21/05/26 | 28/05/26 | 7,1  | 10   |
| 6       | Nícolas Arthur Raulino Oliveira | 2653ad6 | 21/05/26 | 28/05/26 | 7,8  | 10   |
| 6       | Nicolas Pitz                    | 2653ad6 | 21/05/26 | 28/05/26 | 5,8  | 10   |

## Nota parcial

| aluno                           | nota parcial |
| ------------------------------- | ------------ |
| Gabriela Riedel                 | 7,8          |
| Luiz Renato Freitas de Almeida  | 8,3          |
| Nícolas Arthur Raulino Oliveira | 8,4          |
| Nicolas Pitz                    | 7,7          |

## Comentários

### Entrega 1

1. Equipe formada: atendido.
2. Tema definido: atendido.
3. MVP: atendido.
   - A equipe não definiu o que ficará fora do escopo
4. Governança mínima: atendido.
   - As regras de governança devem ser implementadas no repositório.

### Entrega 2

1. Visão do produto: parcialmente atendido.
   - Na proposta de valor, a equipe apenas fez uma descrição do sistema a ser desenvolvido. Falta citar o que ele trará de benefícios e o que ele melhora (o que de valor ele pretende entregar)
2. Definição do MVP: parcialmente atendido.
   - A equipe não definiu o objetivo do MVP.
   - A equipe não informou porque o recorte das funcionalidades é viável para o semestre.
   - A equipe não informou os critérios usados para decidir o que entra e o que fica de fora.
3. Backlog inicial com critérios de aceitação: parcialmente atendida
   - A descrição das issues não está clara
   - O backlog não está priorizado
   - A equipe não definiu critérios de aceitação verificáveis.
4. Definition of Done (DoD): atendido.
5. ADRs iniciais: atendido.
6. Atualização do README: atendido.

- Não há regras de proteção da ramificação principal implementadas no GitHub.
- A equipe apenas definiu algumas regras de proteção da ramificação principal que ainda não estão em vigor.

### Entrega 3

1. Planejamento inicial e baseline: atendido
2. Registro da abordagem de estimativa: atendido
3. Capacidade planejada da equipe: atendido
4. Definição das métricas que serão acompanhadas: atendido
5. Ficha de cada métrica: parcialmente atendido
   - Separar as fichas de cada métrica em arquivos diferentes
     - Em cada uma adicionar data do acompanhamento e valor coletado
   - Numerar métricas para facilitar a referenciação

### Entrega 4

1. Registro inicial de riscos do projeto: parcial
   - A identificação dos risco separada por siglas é difícil de acompanhar; mudem para uma numeração sequencial;
   - A tabela horizontal dificulta a visualização;
2. Análise e priorização dos riscos: atendido
3. Plano de resposta aos riscos: atendido
4. Consolidação do fluxo de trabalho no repositório: atendido
5. Definição inicial de critérios de qualidade do projeto: atendido
6. Relação entre riscos e qualidade: atendido
7. Definição preliminar de avaliação da qualidade: atendido
8. Atualização da documentação do projeto: atendido

### Entrega 5

1. Primeiro incremento funcional do sistema: parcial.
   - O incremento implementa parte relevante do vertical slice do MVP: backend Go com domínio de notas, repositório SQLite, migrations, handler HTTP para criação/resgate de notas e plugin do Obsidian com comando para enviar a nota ativa ao backend via `POST /notes`.
   - O slice atravessa plugin, API e persistência, mas ainda é inicial: no commit avaliado, o plugin cobre principalmente upload da nota ativa; CRUD completo via plugin, sincronização e operações de update/delete ficaram para commits posteriores/Sprint 2.
   - A documentação da raiz não possui instruções claras de execução da entrega; instruções melhores aparecem apenas no `obsidian-plugin/README.md`.
2. Testes de unidade automatizados: atendido.
   - O backend possui testes versionados para domínio, repositório SQLite, handlers HTTP e integração do fluxo de criação/resgate de notas.
3. Escopo da Sprint 1 explicitado e justificado: parcial.
   - O MVP prioriza CRUD de notas Markdown via plugin do Obsidian e API REST, e os commits/PRs mostram foco na fundação desse slice.
   - Não há documento específico de Sprint 1 consolidando issues planejadas, concluídas, parciais e replanejadas, nem justificativa explícita do recorte do vertical slice.
4. Backlog e board atualizados: parcial.
   - Há issues/PRs associados a backend, compose, plugin e CI (`#45`, `#47`, `#52`, `#53`), com revisão por pares.
   - A vinculação completa entre issues, commits e PRs não está consolidada em documentação da sprint.
5. Fluxo de trabalho evidenciado no repositório: atendido.
   - Há desenvolvimento por branches e PRs para `dev`, com aprovações de pares antes dos merges relevantes.
   - PRs principais avaliados: `#45` backend de criação de notas, `#47` compose de desenvolvimento, `#52` estrutura inicial do plugin e `#53` resgate de notas por ID.
   - O workflow foi adicionado na Sprint 1 e contempla backend e plugin.
6. Registro das contribuições individuais: parcial.
   - Não há relatório específico de contribuições individuais da Sprint 1.
   - A autoria é rastreável pelos commits e PRs.
   - Contribuições individuais:
     - Gabriela: era Scrum Master, mas não há contribuição rastreável relevante no commit usado para avaliação da Sprint 1.
     - Luiz Renato: principal responsável pela base backend em Go, incluindo domínio de notas, persistência SQLite, migrations, handlers HTTP, integração e testes. Contribuição central no vertical slice entregue.
     - Nícolas Arthur: responsável por infraestrutura/compose, CI e estrutura inicial do plugin do Obsidian, incluindo cliente, comando de upload e testes. A nota é limitada pela entrega tardia, ausência de release e documentação insuficiente na raiz.
     - Nicolas Pitz: contribuiu com testes, correções e endpoint de resgate de notas por ID, além de reviews relevantes. Menor protagonismo no slice central, com parte do trabalho concentrado em complemento do backend.
7. Documentação atualizada: parcial.
   - Documentos de arquitetura/ADRs foram atualizados para refletir plugin do Obsidian, SQLite e API REST.
   - `obsidian-plugin/README.md` descreve build, testes e configuração do plugin.
   - O README raiz não informa como executar backend, testes, Docker Compose ou plugin, apesar de ser o ponto principal de entrada do projeto.
8. Release do marco: não atendido.
    - Não existe tag `v0.1.0` localmente nem release `v0.1.0` no GitHub.

### Entrega 6

1. Incremento funcional do MVP: parcial.
   - Funcionalidade declaradas para entrega na Sprint 2:
      - #20 - Frontend: UI de Exclusão e Sincronização de Estado
      - #48 - Documentação das rotas já existentes
      - #44 - Configuração do servidor via arquivo JSON
   - A issue `#7` foi fechada, mas um critério de aceitação importante não está implementado: não há versionamento, histórico ou log de alteração antes de sobrescrever conteúdo; a tabela `notes` mantém apenas o estado atual.
   - A sincronização detecta conflito quando servidor e arquivo local mudaram desde a última sincronização, mas não há resolução guiada além de o usuário escolher manualmente subir ou baixar.
   - O plugin depende do ambiente real do Obsidian para demonstração completa; os testes cobrem os comandos por mocks, mas não foi verificada uma execução ponta a ponta dentro do Obsidian.
2. Testes automatizados: atendido.
3. Integração contínua mínima: parcial.
   - O CI não executa `npm run lint`, embora o lint exista e passe localmente.
4. Pull requests com revisão: atendido.
   - O PR consolidador `#69` é grande e incorpora muitos commits antigos e PRs já existentes, o que reduz um pouco a clareza da rastreabilidade fina da sprint.
5. Aplicação justificada de padrões OO: não atendido.
   - Não há indicação ou justificativa específica de padrão de projeto para a Sprint 2 além das ADRs e da estrutura em camadas já usada.
6. Atualização das métricas: não atendido.
   - Não há valores observados da Sprint 2, data de coleta, análise de tendência ou comparação entre planejado e realizado.
7. Atualização dos riscos: não atendido.
   - A atualização permanece genérica; não registra acompanhamento concreto do fim da Sprint 2, riscos materializados, riscos encerrados ou mudanças de probabilidade/impacto baseadas na execução.
8. Release do marco: parcial.
   - A descrição da release é uma lista de PRs, sem relatório claro da Sprint 2 com escopo planejado, concluído, pendente, métricas e justificativas.
9. Registro das contribuições individuais: parcial.
    - Não há relatório específico de contribuições individuais da Sprint 2; a distribuição foi inferida por PRs, commits e autoria.
    - Contribuições individuais:
      - Gabriela: implementou `PUT /notes/{id}` e `DELETE /notes/{id}` no PR `#54`, adicionou OpenAPI no `#60` e trabalhou em Docker/Makefile no `#61`. A contribuição técnica é relevante, mas a nota é limitada porque o Makefile entregue quebra na tag e parte da documentação/reprodutibilidade ficou incompleta.
      - Luiz Renato: contribuiu com configuração JSON (`#62`), ajustes estruturais, documentação/ADRs, revisão de PRs e consolidação da Sprint 2 no `#69`. A participação foi importante para integração e arquitetura, mas menos central que a implementação do plugin nesta entrega e a release consolidada saiu atrasada.
      - Nícolas Arthur: foi o principal responsável pelo incremento do plugin e da sincronização no PR `#58`, incluindo comandos, metadados locais, importação, sincronização em lote, tratamento de conflitos e testes. Recebe a maior nota individual, limitada por pendências do produto, pela ausência de versionamento/log e pela dependência de validação manual no Obsidian.
      - Nicolas Pitz: atualizou o baseline no `#63` e realizou contribuições pontuais/anteriores em backend e testes, mas a documentação da Sprint 2 ficou incompleta e as métricas sob sua responsabilidade não tiveram coleta de valores observados. A nota individual é menor pela menor entrega rastreável no incremento funcional avaliado.
10. Documentação e reprodutibilidade: parcial.
   - O README da raiz não concentra instruções completas de execução/teste da entrega.
   - O comando automatizado `make all` falha na tag avaliada com `Makefile:20: *** faltando o separador.  Pare.`, o que compromete a reprodutibilidade, embora os comandos diretos equivalentes funcionem.
