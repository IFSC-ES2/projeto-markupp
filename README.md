# Editor de Texto Self-Hosted para centralização em MarkDowm

## Equipe:

 - Renato Freitas
 - Nícolas Arthur
 - Nicolas Pitz
 - Gabriela Riedel

## Descrição

Desenvolvimento de um editor de texto com arquitetura centralizada e self-hosted, permitindo que usuários criem, editem, organizem e versionem documentos markdown em uma estrutura hierárquica por projetos e categorias. O sistema deverá funcionar como um hub central de conhecimento, oferecendo operações completas de criação, leitura, atualização, exclusão e controle de versões, além de manter organização por caminhos lógicos semelhantes a diretórios.

## MVP

### O que fará
 - Prover um ambiente centralizado para criação, edição e organização de documentos Markdown
 - Estruturar o conhecimento em hierarquia lógica
 - Expor uma API REST que permita integração com clientes
 - Processar documentos automaticamente para busca semântica (chunking, embedding, etc…)

 ### O que pode ser feito: 
 - Usuários
 - Editor de texto markdown 
 - API para integração com agentes
 - Manter histórico de versões por documento
 - Desenvolver sistema de sync dos arquivos para múltiplos usuários editando

## Levantamento de Requisitos

### Requisitos Funcionais
 
- RF1: Permite salvar dados em markdown;
- RF2: Permite adicionar arquivos no frontend;
- RF3: Permite excluir arquivos no frontend;
- RF4: Permite renomear arquivos no frontend;
- RF5: Permite editar markdown;

### Requisitos Não Funcionais
- RNF1: Deve estar disponível 99% do tempo;
- RNF2: Teremos Login para Usuário;
- RNF3: Deve ter para Desktop;
- RNF4: Deve suportar milhares de requisições;
- RNF5: Deve ter problemas de segurança mínimos e não destrutivos ao final;

## Governança Mínima

Quem pode dar pull requests (Todos)
Quem pode aprovar pull requests (Qualquer um que não tenha aberto o pull request)
Ninguem pode dar merge diretamente na main (Somente HotFix)
