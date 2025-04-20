# Etapa 1: Build do binário Go
FROM golang:1.24-alpine3.21 AS builder

# Variáveis de ambiente para build seguro e reprodutível
ENV USER=appuser
ENV UID=10001

# Instala dependências necessárias para build
RUN apk update && apk add --no-cache git ca-certificates tzdata && update-ca-certificates

# Cria diretórios para o projeto
RUN mkdir -p /srv/bbb-convert /srv/bbb-convert/release
WORKDIR /srv/bbb-convert

# Copia os arquivos do projeto
COPY . .

# Baixa dependências
RUN go mod download

# Compila o binário de forma estática e enxuta
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags='-w -s -extldflags "-static"' -a -o /srv/bbb-convert/release/bbb-convert

# Etapa 2: Imagem final enxuta
FROM alpine:3.21

# Instala apenas as dependências de runtime necessárias
RUN apk add --no-cache chromium ffmpeg

# Cria usuário não-root para segurança
RUN adduser -D bigbluebutton bigbluebutton

# Copia o binário pronto da etapa de build
COPY --from=builder /srv/bbb-convert/release /srv/bbb-convert

USER bigbluebutton
WORKDIR /home/bigbluebutton

# Define o entrypoint para aceitar parâmetros dinâmicos
ENTRYPOINT ["/srv/bbb-convert/bbb-convert"]
