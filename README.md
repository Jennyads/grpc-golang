<h3>gRPC </H3>

- tem finalidade de facilitar a comunicação entre sistemas;
- ideal para microsserviços (latência entre os acessos de cada sistemas, tempo de resposta muito alto);
- geração das bibliotecas de forma automática - stubs (trechos de código que possibilita a comunicação de uma conta com outra que fala a mesma língua, cria uma camada de extração/abstração para não ter que lidar com complexidades)
- RPC: remote procedure call
```
cliente -> server
cliente[
server.soma(a,b)
]
server [
func soma(int a, intb){}
]
```
- .proto - contrato definido
-Protocol Buffers - linguagem e plataforma neutra, serializar estruturas de dados por meio de arquivos binários. 
arquivos binários < JSON, assim gasta menos recursos de rede.
- HTTP2: dados trafegados são binários e não texto(JSON) como no HTTP 1.1


-gerar stub/biblioteca (protoc compilador gere e utilize os arquivos): protoc --proto_path=proto proto/*.proto --go_out=plugins=grpc:pb
- go mod init
- evans -r -p 50055
- go run cmd/server/server.go
- go run cmd/client/client.go

