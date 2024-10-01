# mailer-microservice

## Início Rápido

Siga os passos abaixo para iniciar a aplicação:

1. **Clone o repositório**:

   ```bash
   git clone https://github.com/ijoaog/mailer-microservice.git
   cd mailer-microservice
   ```

2. **Instale as dependências**:

   ```bash
   go mod tidy
   ```

3. **Coloque o arquivo .env na raiz**:

   ```bash
   Seguir o arquivo .env.example
   ```

4. **Para carregar .env**:

   ```bash
   no arquivo main.go:
   descomentar a linha 18: 
   ## //config.Init()
   ```

5. **Inicie a aplicação em modo de desenvolvimento**:

   ```bash
   Na pasta raiz executar o seguinte comando:
   go run main.go
   ```

