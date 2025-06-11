# 🔐 Go Auth API (Estudo)

Este é um projeto de estudos que implementa uma API de autenticação usando **Go**, **Gin**, **JWT**, **Bcrypt** e **MongoDB**. O objetivo é compreender os fundamentos de autenticação, controle de acesso por roles e boas práticas de segurança em APIs REST.

---

## 🚀 Funcionalidades

- ✅ Cadastro de usuário (Signup)
- ✅ Login com geração de token JWT
- ✅ Hash de senha com Bcrypt
- ✅ Middleware de autenticação
- ✅ Middleware de autorização por role (ex: `ADMIN`)
- ✅ Rota para buscar todos os usuários (somente `ADMIN`)
- ✅ Rota para buscar usuário por ID (somente `ADMIN`)
- ✅ MongoDB como banco de dados principal

---

## 🛠️ Tecnologias usadas

- [Go](https://golang.org/) — linguagem de programação
- [Gin](https://github.com/gin-gonic/gin) — framework web
- [JWT](https://github.com/golang-jwt/jwt) — autenticação baseada em token
- [Bcrypt](https://pkg.go.dev/golang.org/x/crypto/bcrypt) — hash de senhas
- [MongoDB](https://www.mongodb.com/) — banco NoSQL

---

## 📚 Endpoints

### ✅ Auth

- POST /signup: Criação de usuário
- POST /login: Autenticação e retorno de token JWT

### 🔒 Rotas protegidas (requer token JWT)

- GET /users: Lista todos os usuários (requer role ADMIN)
- GET /users/:id: Busca usuário por ID (requer role ADMIN)

As rotas protegidas exigem o token JWT no header:
`Authorization: Bearer <token>`

## 🧪 Exemplos de payload

Para fins didáticos foi incluído ao corpo da requisição a `role` (ADMIN | USER).

```json
{
    "first_name": "John",
    "last_name": "Doe",
    "password": "password123",
    "email": "johndoe@email.com",
    "phone": "5581395243",
    "role": "ADMIN"
}
```

## 🧠 Aprendizados

- Implementação de autenticação com JWT
- Proteção de rotas com middlewares
- Uso de roles para controle de acesso
- Integração entre Go e MongoDB
- Boas práticas de segurança (hash de senhas, verificação de token, etc.)