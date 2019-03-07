### This project requires the use of OpenAPI2.0(swagger) and Golang commandline tools
### We also use 'dep' to manage dependencies which will need to be downloaded before the project can be built

### generate swagger files
swagger generate server -f swagger.yml

### download dependences
dep init
dep ensure

### build the server
go build .\cmd\x98point6-drop-token-server

### start the server
.\x98point6-drop-token-server --port=8086


### sample usage
curl -X GET http://localhost:8086/drop_token
curl -X POST http://localhost:8086/drop_token -H "accept: application/json" -H "Content-Type: application/json" -d @input\createGame.json

### In Game0 player1 makes a move in column 0
curl -X POST http://localhost:8086/drop_token/Game0/player1 -H "accept: application/json" -H "Content-Type: application/json" -d @input\makeMove0.json

### Known issues

- When the user asks for a list of moves the column number is omitted from any moves made to column 0. e.g.

curl -X GET http://localhost:8086/drop_token/Game1/moves
{"moves":[{"player":"player1","type":"MOVE"},{"column":1,"player":"player2","type":"MOVE"},{"player":"player2","type":"QUIT"}]}

The reason for this error is that "column" is an optional parameter to Move and is omitted when empty. A value of 0 is taken as the empty value and this field is removed.
This should be fixed by porting the swagger spec to OpenAPI 3.0 which allows for multiple object types in an array.

- Missing functional tests

- The project only supports 2 players on a 4x4 board. Additional tests are needed before increasing either of these dimensions.