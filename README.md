Helper Bot Backend
=====

## Deployment & Run application
Copy and set env variables
```bash
$ cp .env.dist .env
```
Run application
```bash
$ go run cmd/server/main.go
```

## Env variables
1. `TELEGRAM_BOT_TOKEN` - Telegram API Token for connection with bot
2. `OPENAI_TOKEN` - OpenAI Token for API