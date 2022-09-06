# Sender
App for send message to social media
## How to use
### Build
Here is avaliable two modes `http` and `amqp` (default)

#### AMQP example
```
cp .env.example .env
# write BOT_TOKEN 
docker-compsoe up -d
```
#### Http example
```
docker build -t sender .
docker run --rm -p 8080:8080 -e BOT_TOKEN=<telegram_bot_token> sender http
```

### Send Message
#### http mode
```
curl -X POST http://localhost:8080/send/telegram \
   -H 'Content-Type: application/json' \
   -d '{"chat_id":"<telegram_chat_id>","message":"hello, http!"}'
```
