# Use uma imagem base do Go
FROM golang:1.23

# Instala o Air
RUN go install github.com/air-verse/air@latest

# Define o diretório de trabalho
WORKDIR /app

# Copia o go.mod e go.sum do diretório pai
COPY ../go.mod ../go.sum ./

# Baixa as dependências
RUN go mod download

# Copia o restante do código
COPY . .

# Exponha a porta que a aplicação vai rodar
EXPOSE 8080

# Comando para iniciar o Air
CMD ["air"]
