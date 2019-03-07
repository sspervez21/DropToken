swagger generate server -f swagger.yml
go build .\cmd\x98point6-drop-token-server
.\x98point6-drop-token-server --port=8086

curl -X GET http://localhost:8086/drop_token
curl -X POST http://localhost:8086/drop_token -H "accept: application/json" -H "Content-Type: application/json" -d @input\createGame.json
curl -X POST http://localhost:8086/drop_token/Game0/player1 -H "accept: application/json" -H "Content-Type: application/json" -d @input\makeMove.json
