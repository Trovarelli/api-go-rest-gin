API Go REST com Gin + GORM

Resumo
- API simples para CRUD de alunos construída em Go usando o framework web Gin e o ORM GORM.
- Objetivo educacional: praticar e aprender conceitos essenciais de Gin (roteamento, middlewares, handlers) e GORM (mapeamento, migrações automáticas, consultas e atualizações).

Stack
- Go 1.22
- Gin v1.11
- GORM v1.x + driver Postgres
- Postgres (via Docker Compose)

Arquitetura
- `src/controllers`: Handlers HTTP (camada web)
- `src/repository`: Acesso a dados via GORM (camada de persistência)
- `src/models`: Modelos da aplicação
- `src/database`: Conexão e AutoMigrate
- `src/routes`: Definição das rotas

Modelo
- `Aluno` (campos):
  - `id` (int64, PK, auto-incremento)
  - `nome` (string)
  - `cpf` (string)
  - `rg` (string)

Rotas
- `GET    /alunos` → lista todos
- `GET    /alunos/:id` → busca por id
- `POST   /alunos` → cria (JSON: {"nome","cpf","rg"})
- `PUT    /alunos/:id` → atualiza (JSON: {"nome","cpf","rg"})
- `DELETE /alunos/:id` → remove

Como executar
1) Sobe o Postgres com Docker Compose
   - Pré‑requisito: Docker e Docker Compose
   - Comando: `docker compose up -d`
   - Serviços:
     - `postgres`: usuário, senha e banco `root` expostos na porta `5432`
     - `pgadmin-compose`: opcional, UI do PGAdmin em `http://localhost:54321`

2) Configurar a conexão da aplicação (opcional)
   - A aplicação lê a variável `DATABASE_URL` (DSN do Postgres). Exemplo:
     - `DATABASE_URL="host=localhost user=root password=root dbname=root port=5432 sslmode=disable"`
   - Caso não defina, o valor acima é usado por padrão.

3) Rodar a aplicação localmente
   - `go mod tidy`
   - `go run ./src`
   - Servidor em `http://localhost:8080`

Exemplos de requisições (curl)
- Criar
  - `curl -X POST http://localhost:8080/alunos -H "Content-Type: application/json" -d '{"nome":"João","cpf":"123","rg":"456"}'`
- Listar
  - `curl http://localhost:8080/alunos`
- Buscar por ID
  - `curl http://localhost:8080/alunos/1`
- Atualizar
  - `curl -X PUT http://localhost:8080/alunos/1 -H "Content-Type: application/json" -d '{"nome":"João da Silva","cpf":"123","rg":"456"}'`
- Excluir
  - `curl -X DELETE http://localhost:8080/alunos/1 -i`

Notas de implementação
- Migração automática: o `AutoMigrate` é executado ao iniciar a aplicação (veja `src/database/database.go`).
- O `docker-compose.yaml` referencia `migration/docker-database-initial.sql` apenas para extensões/SQL opcional; o schema da tabela de `alunos` é criado pelo GORM.
- Em `src/models/aluno.go`, foi removido o embed de `gorm.Model` para evitar conflito com o campo de chave primária `Id` (int64). Caso deseje timestamps, você pode adicionar campos `CreatedAt`, `UpdatedAt`, `DeletedAt` e/ou usar `gorm.Model` e alterar o código para usar `ID uint`.

Objetivo do aprendizado
- Gin
  - Roteamento, binding JSON e respostas HTTP
  - Organização em controllers e camadas
- GORM
  - Conexão com Postgres, AutoMigrate e CRUD
  - Filtros com `Where`, `First`, `Updates` e controle de `RowsAffected`
  - Boas práticas de repositório e contexto (`WithContext`)

Como contribuir / próximos passos
- Adicionar validações de entrada (ex.: validar formato de CPF)
- Introduzir testes de unidade/integrados (ex.: com banco em memória ou container)
- Parametrizar porta do servidor e CORS
- Melhorar logs e middlewares no Gin

