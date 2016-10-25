# Global Identity
 Este é um pacote criado com o intuito de facilitar a utilização do Global Identity para autenticação de aplicações e usuários em seus projetos Go.

## Instalação

```go
go get github.com/julio-vaz/globalidentity
```

## Funcionalidades

 - **Autenticação de usuários**
   - AuthenticateUser(applicationKey string, email string, password string, expirationInMinutes int) map[string]interface{}

 - **Validação de tokens**
   - ValidateToken(applicationKey string, token string) map[string]interface{}

 - **Validação de papeis de usuários**
   - HasRoles(applicationKey string, userKey string, roles []string) map[string]interface{}

 - **Validação de aplicações**
  - ValidateApplication(applicationKey string, clientApplicationKey string, rawData string, encryptedData string) map[string]interface{}

