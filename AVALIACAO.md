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
| 7       | Gabriela Riedel                 | e4f7d2d | 29/05/26 | 03/06/26 | 5,6  | 10   |
| 7       | Luiz Renato Freitas de Almeida  | e4f7d2d | 29/05/26 | 03/06/26 | 5,4  | 10   |
| 7       | Nícolas Arthur Raulino Oliveira | e4f7d2d | 29/05/26 | 03/06/26 | 4,4  | 10   |
| 7       | Nicolas Pitz                    | e4f7d2d | 29/05/26 | 03/06/26 | 5,0  | 10   |
| 8       | Gabriela Riedel                 | af1e784 | 09/06/26 | 11/06/26 | 7,7  | 10   |
| 8       | Luiz Renato Freitas de Almeida  | af1e784 | 09/06/26 | 11/06/26 | 8,5  | 10   |
| 8       | Nícolas Arthur Raulino Oliveira | af1e784 | 09/06/26 | 11/06/26 | 8,4  | 10   |
| 8       | Nicolas Pitz                    | af1e784 | 09/06/26 | 11/06/26 | 8,1  | 10   |
| 9       |                                 |         |          |          |      | 10   |
| 10      |                                 |         |          |          |      | 10   |
| 11/12   |                                 |         |          |          |      | 30   |

## Nota parcial

| aluno                           | nota parcial |
| ------------------------------- | ------------ |
| Gabriela Riedel                 | 7,3          |
| Luiz Renato Freitas de Almeida  | 7,8          |
| Nícolas Arthur Raulino Oliveira | 7,6          |
| Nicolas Pitz                    | 7,3          |

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

### Entrega 7

1. Incremento funcional da Sprint 3: parcial.
   - Não há tag `v0.3.0` local nem remota e não há release da Sprint 3 no GitHub. A única release listada continua sendo `v0.2.0`.
   - Entrega avaliada em `e4f7d2d`, mas `v0.2.0..HEAD` em `main` contém apenas commits de avaliação, sem commits da equipe integrando a Sprint 3 na branch principal.
   - Há evidências de trabalho em PRs para `dev`: `#72` adiciona lint/formatação ao CI, `#73` renomeia `backend` para `markupp`, `#74` adiciona rota de busca `GET /notes/search` e `#76` atualiza baseline. Porém, isso não foi consolidado em `main` nem empacotado em release/tag do marco.
   - Issues da Sprint 3 existem, mas parte permanece aberta ou incompleta: `#68` e `#78` estão abertas; `#77` e `#79` estão abertas com mudanças solicitadas; `#75` de arquitetura também ficou aberto.
   - O backend da branch principal já possui CRUD e listagem de notas, e os testes passam, mas o incremento funcional específico da Sprint 3 não foi entregue de forma rastreável no marco obrigatório.
2. Documentação da arquitetura: não atendido.
   - Não há diagrama C4 ou equivalente integrado em `main`.
   - O documento `arquitetura-c4.md` existe apenas no PR `#75`, que ficou aberto e com `CHANGES_REQUESTED` antes de receber aprovação posterior em 01/06; portanto não foi entregue no marco.
3. ADRs consolidados: parcial.
   - Existem ADRs anteriores para linguagem, mas não encontrei ADR novo consolidado para as mudanças da Sprint 3.
   - Parte dos ajustes de nomenclatura em ADRs aparece no PR `#73`, mas a consolidação não chegou à branch principal/tag do marco.
4. Atualização das métricas: não atendido.
   - `metricas.md` continua apenas como definição das métricas; não registra valores observados ao final da Sprint 3, comparação com Sprint 2 nem análise dos fatores que influenciaram o resultado.
   - `baseline.md` em `main` ainda marca `#66`, `#67` e `#68` como “Não iniciado (Sprint 3)”. A atualização em `origin/dev` foi feita no PR `#76`, mas não traz análise completa de métricas antes/depois.
   - Não há arquivo específico de comparação de métricas antes/depois da manutenção/reengenharia.
5. Testes automatizados integrados ao pipeline: parcial.
   - Na branch principal, o workflow executa `go test ./... -race -count=1` no backend e `npm ci`, `npm run build` e `npm test --if-present` no plugin.
   - Execução local: `go test ./... -race -count=1` passou; `npm test` passou com 44 testes; `npm run build` passou; `npm run lint` passou.
   - O CI de `main` ainda não executa `npm run lint`, `gofmt`, `go vet` ou `golangci-lint`; essas melhorias aparecem no PR `#72`/`origin/dev`, mas não foram integradas na branch principal nem em release.
6. Integração contínua mínima: parcial.
   - Há CI em PRs e checks verdes em PRs como `#72`, `#73` e `#74`.
   - O PR `#77` tem check de backend falhando e está com mudanças solicitadas; `#79` também ficou aberto com mudanças solicitadas.
   - A branch principal não reflete o pipeline mais completo da Sprint 3, e não há evidência de check obrigatório associado a uma release/tag `v0.3.0`.
7. Release/tag do marco: não atendido.
   - O projeto ainda conta apenas com a tag `v0.2.0`.
   - O projeto ainda conta apenas com a release “Sprint 2”.
   - Ausência de tag/release `v0.3.0` compromete a rastreabilidade do marco e a definição do commit exato da entrega.
8. Registro de contribuição individual: parcial.
   - Não há relatório específico da Sprint 3 consolidando o que cada integrante implementou, revisou, testou ou documentou com links para commits/PRs.
   - Como `main` não recebeu commits da equipe depois de `v0.2.0`, as contribuições foram inferidas principalmente pelos PRs em `dev` e branches remotas.
    - Gabriela: contribuiu tecnicamente no PR `#74` com busca `GET /notes/search`, paginação, testes e correção do Makefile, além de commits de ajuste/review. A nota fica abaixo de 6 por estar fora de `main`, sem tag/release.
    - Luiz Renato: contribuiu no PR `#73` com refatoração de nomenclatura `backend` -> `markupp`, ajustes em CI, Docker, docs e plugin, além de revisão de PRs. A contribuição é relevante para reengenharia/manutenibilidade, mas não substitui a ausência de arquitetura/métricas/release consolidadas.
    - Nícolas Oliveira: aparece em merge tardio do PR `#76` e em PRs abertos de plugin/source control (`#79`) e sincronização, mas as principais evidências da Sprint 3 ficaram abertas, com mudanças solicitadas ou depois do prazo. A nota é limitada porque houve menos entrega integrada e aceita no marco.
    - Nicolas Pitz: contribuiu com PR `#72` de CI/lint/formatação e tentou documentar C4 no PR `#75`, além de atualizar baseline no PR `#76`. A nota reconhece contribuição em qualidade/documentação, mas penaliza a arquitetura não integrada, métricas insuficientes e fechamento posterior ao prazo.

### Entrega 8

1. Ambiente de staging ou equivalente acessível: atendido.
   - O projeto não expõe um SaaS público, mas a entrega define um equivalente self-hosted por imagem Docker publicada em Docker Hub: `riedelgab/ifsces2:latest`.
   - O smoke test documentado em `docs/DEPLOY.md` funcionou localmente via Docker: `POST /notes` criou nota, `GET /notes/{id}` retornou o conteúdo e `DELETE /notes/{id}` respondeu `204`.
   - Para o domínio do projeto, que é servidor local/self-hosted com plugin Obsidian, essa solução é compatível com o equivalente de staging, desde que a imagem versionada seja mantida.
2. Manutenção e atualização da integração contínua: parcial.
   - `.github/workflows/ci.yml` cobre backend com `gofmt`, `go vet`, `golangci-lint` e testes com `race`, além do plugin com `npm ci`, `npm run lint`, `npm run build` e `npm test`.
   - Localmente, `go test ./... -race -count=1` passou no backend e o plugin passou em `npm ci`, `npm run lint`, `npm run build` e `npm test` com 55 testes.
   - O comando do CI `go run github.com/golangci/golangci-lint/cmd/golangci-lint@latest run` falhou localmente porque o linter foi construído com Go 1.25 enquanto o módulo alvo declara `go 1.26.1`. A solução é fixar uma versão/binário de `golangci-lint` compatível com Go 1.26 ou alinhar o `go.mod` a uma versão suportada pelo linter.
3. Documentação de deploy: atendido com ressalvas.
   - `docs/DEPLOY.md` descreve pré-requisitos, execução via Docker, instalação do plugin, validação por `curl`, build local e aviso de segurança por ausência de autenticação.
   - O README referencia o artefato da release `v0.4.0` e o fluxo de instalação do plugin no Obsidian.
   - A documentação poderia usar imagem/tag versionada no comando principal, em vez de `latest`, para melhorar reprodutibilidade do marco avaliado.
4. Atualização das métricas do projeto: parcial.
   - As fichas foram separadas em `docs/metricas/` e há atualização visual para Sprint 4 em cobertura, MVP, reviews, PRs revisados e distribuição de commits.
   - A evidência está majoritariamente em imagens, com pouca explicação textual, valores numéricos, data de coleta, metodologia e análise de tendência/decisão.
   - O baseline foi atualizado para marcar itens como feitos, mas ainda mistura linha de base original com acompanhamento posterior e mantém trechos históricos inconsistentes, como período coberto até a Entrega 3 ao mesmo tempo em que a versão é 0.4.
5. Manutenção/reengenharia: atendido.
   - `ADR-0010` registra isolamento de detalhes de persistência na camada `storage`, removendo vazamentos de `database/sql` e `sql.ErrNoRows` para camadas superiores.
   - `ADR-0011` registra a substituição de comandos avulsos pela Source Control View e pela camada `core/operations.ts`/`core/status.ts`.
   - A reengenharia é coerente com problemas reais de duplicação, coesão e visibilidade de estado para o usuário.
   - Como evidência funcional associada, a Sprint 4 consolida o produto com Source Control View no plugin, comandos `fetch`, `pull`, `push` e `sync`, fluxo mais próximo de controle de versão, suporte a conflitos/força de sincronização e correções em rotas/serviços do backend.
   - A validação completa dentro do Obsidian continua dependente de execução manual no aplicativo, mas há testes automatizados expressivos do plugin.
6. Comparação de métrica antes/depois: atendido.
   - `ADR-0011` apresenta comparação objetiva entre antes e depois: 5 comandos com 413 LOC de lógica duplicada removidos, persistência de metadados reduzida de 5 pontos para 1 helper e `forcePush` deixando de duplicar CRUD.
   - A métrica não é apenas contagem de linhas; ela conecta a alteração a manutenibilidade, coesão e visibilidade do fluxo de sincronização.
7. Release/tag do marco: parcial.
   - A tag `v0.4.0` existe e aponta para `af1e784`, que também é o `HEAD` de `main`.
   - A release `Sprint 4` foi publicada para `v0.4.0`, mas em 09/06/2026, após o marco previsto de 08/06/2026.
   - A entrega recupera parte importante da rastreabilidade perdida na Sprint 3: há também tag/release `v0.3.0` e documento de recuperação da Entrega 7.
8. Registro de contribuição individual: atendido com ressalvas.
   - `docs/contribuicoes-sprint4.md` registra contribuições por membro.
   - O histórico `v0.3.0..v0.4.0` confirma participação dos quatro integrantes, embora com aliases diferentes.
   - O documento é curto e não vincula todas as contribuições a PRs/commits, mas é suficiente para diferenciar responsabilidades principais.
   - Contribuições individuais:
      - Gabriela: contribuição relevante em deploy, Dockerfile e documentação operacional. A nota é limitada pelo menor volume relativo de evidências técnicas no incremento principal e pela release após o prazo.
      - Luiz Renato: contribuição forte em reengenharia de persistência, ADR-0010, testes, consolidação do marco e revisões. Recebe a maior nota por protagonismo técnico e integração da entrega, com ressalva pela falha do lint local e atraso da release.
      - Nícolas Arthur: contribuição forte no incremento principal do plugin, Source Control View, reorganização de sincronização, ADR-0011, CD do plugin e documentação de instalação. Nota alta, limitada por dependência de validação manual no Obsidian e pela entrega tardia do marco.
      - Nicolas Pitz: contribuição relevante em sincronização do servidor, testes, baseline e métricas. A nota reconhece sustentação de qualidade, mas é limitada porque as métricas da Sprint 4 ficaram pouco analíticas e baseadas quase só em imagens.
