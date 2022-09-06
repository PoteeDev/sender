# Sender
App for send message to social media
## How to use
### Build
here is avaliable two modes `http` and `amqp` (default)

to use http mode you need to uncomment `http` examle 
```
cp .env.example .env
# write BOT_TOKEN 
docker-compsoe up -d
```

### Send Message
#### http mode
```
curl -X POST http://localhost:8080/send/telegram 
   -H 'Content-Type: application/json'
   -d '{"chat_id":"<telegram_chat_id>","message":"hello, http!"}'
```
