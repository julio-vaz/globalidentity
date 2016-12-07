# Global Identity
 Este é um pacote criado com o intuito de facilitar a utilização do Global Identity para autenticação de aplicações e usuários em seus projetos Go.

## Instalação

```go
go get github.com/julio-vaz/globalidentity
```

## Funcionalidades

 - **Autenticação de usuários**
   - AuthenticateUser(email string, password string, expirationInMinutes ...int) (string, error)

 - **Validação de tokens**
   - ValidateToken(token string) (bool, error)

 - **Validação de papeis de usuários**
   - IsUserInRoles(userKey string, roles ...string) (bool, error)

 - **Validação de aplicações**
  - ValidateApplication(applicationKey string, clientApplicationKey string, rawData string, encryptedData string) (bool, error)
  
 - **Renovação de tokens**
  - RenewToken(token string) (string, error)

