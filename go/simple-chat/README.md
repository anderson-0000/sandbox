# run the program
go run main.go

# send message
curl -X POST http://localhost:8080/messages -d '{"text":"こんにちは"}'

# get message
curl http://localhost:8080/messages 
