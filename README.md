API Go REST com Gin + GORM

Resumo
- API para CRUD de alunos construída em Go usando Gin (web) e GORM (ORM).
- Atualizações recentes: busca por CPF, normalização de CPF no backend, validações de entrada, timestamps e soft delete, padrão de repositório com contexto.

Stack
- Go 1.22
- Gin v1.11
- GORM v1.x + driver Postgres
- Postgres (via Docker Compose)

Arquitetura
- `src/controllers`: Handlers HTTP (camada web)
- `src/repository`: Acesso a dados via GORM (camada de persistência)
- `src/models`: Modelos da aplicação + validações
- `src/database`: Conexão e AutoMigrate
- `src/routes`: Definição das rotas
- `src/helpers`: Utilitários (ex.: normalização de strings)

Modelo
- `Aluno` (campos):
  - `id` (int64, PK, auto-incremento)
  - `nome` (string)
  - `cpf` (string; normalizado no backend removendo pontuação/caracteres especiais)
  - `rg` (string)
  - `created_at`, `updated_at` (gerenciados pelo GORM)
  - `deleted_at` (soft delete; não exposto no JSON)

Validações (validator.v2)
- `nome`: obrigatório (`nonzero`)
- `cpf`: tamanho 9 (`len=9`)
- `rg`: tamanho 11 (`len=11`)

Rotas
- `GET    /alunos` — lista todos
- `GET    /alunos/:id` — busca por id
- `GET    /alunos/cpf/:cpf` — busca por CPF
- `POST   /alunos` — cria (JSON: {"nome","cpf","rg"})
- `PUT    /alunos/:id` — atualiza (JSON: {"nome","cpf","rg"})
- `DELETE /alunos/:id` — remove (soft delete)

Status e erros
- 200 OK (consultas e update)
- 201 Created (criação)
- 204 No Content (remoção)
- 400 Bad Request (ID inválido ou payload inválido)
- 404 Not Found (registro inexistente)
- 500 Internal Server Error (erros inesperados)

Como executar
1) Subir o Postgres com Docker Compose
   - Pré‑requisitos: Docker e Docker Compose
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
- Criar (CPF pode conter pontuação; será normalizado)
  - `curl -X POST http://localhost:8080/alunos -H "Content-Type: application/json" -d '{"nome":"João","cpf":"123.456.789","rg":"12345678901"}'`
- Listar
  - `curl http://localhost:8080/alunos`
- Buscar por ID
  - `curl http://localhost:8080/alunos/1`
- Buscar por CPF
  - `curl http://localhost:8080/alunos/cpf/123456789`
- Atualizar
  - `curl -X PUT http://localhost:8080/alunos/1 -H "Content-Type: application/json" -d '{"nome":"João da Silva","cpf":"123.456.789","rg":"12345678901"}'`
- Excluir
  - `curl -X DELETE http://localhost:8080/alunos/1 -i`

Notas de implementação
- Migração automática: `AutoMigrate` executa ao iniciar a aplicação (`src/database/database.go`).
- Repositório com contexto: operações usam `db.WithContext(ctx)` para melhor controle e cancelamento.
- Normalização de CPF: `helpers.NormalizeString` remove caracteres especiais antes de persistir.
- Soft delete: o campo `DeletedAt` permite restauração futura e evita remoção física.
- JSON: timestamps são expostos como `created_at`/`updated_at`; `deleted_at` não é serializado.

Objetivo do aprendizado
- Gin
  - Roteamento, binding JSON e respostas HTTP
  - Organização em controllers e camadas
- GORM
  - Conexão com Postgres, AutoMigrate e CRUD
  - Filtros com `Where`, `First`, `Updates` e controle de `RowsAffected`
  - Boas práticas de repositório e contexto (`WithContext`)

Como contribuir / próximos passos
- Regras de validação adicionais (ex.: validar formato real de CPF)
- Testes de unidade e integração (ex.: com container de banco)
- Parametrizar porta do servidor e CORS
- Melhorar logs e middlewares no Gin

