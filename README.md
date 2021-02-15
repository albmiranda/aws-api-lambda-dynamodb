# Operação Fogo de Quaser

Este projeto é estudo de utilização de API GATEWAY, LAMBDA e DYNAMODB da AWS.

A cloud computing adotada foi a Amazon AWS devido prévia familiaridade do desenvolvedor com a plataforma.
Para gerenciar o recebimento das chamadas de API foi utilizado o serviço **API GATEWAY** da AWS que, além do gerencimanento das chamadas, oferece maior segurança em eventual controle de autorização, etc.
A aplicação backend é executada em plataforma serverless **Lambda** e os dados persistentes armazenados em **DynamoDB**.

## Dependências

Para seu desenvolvimento foram necessárias apenas pacotes nativos do Go e outros *3rd party*. Para sua resolução faça:

```go
    go get -u github.com/aws/aws-lambda-go/events
    go get -u github.com/aws/aws-lambda-go/lambda
    go get -u github.com/aws/aws-sdk-go/aws
    go get -u github.com/aws/aws-sdk-go/aws/session
    go get -u github.com/aws/aws-sdk-go/service/dynamodb
    go get -u github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute
    go get -u github.com/gookit/validate
    go get -u golang.org/x/lint/golint
```
## Compilação

A compilação da aplicação pode seguir o padrão de qualquer projeto em GO. No entanto, para facilitar o trabalho do desenvolvedor, foi criado um Makefile com targets padrões.

Para compilar faça:

```shell
    make build
```

A fim de possibilitar o empacotamento do binário em pacote compatível com AWS Lambda foi criado um target específico chamado "package". Para gerar o artefato (.zip) faça:
```shell
    make package
```

Para mais ajuda com o Makefile digite "make help":
```
andre@laptop:~/go/src/aws-api-lambda-dynamodb$ make help
build         Build main application 
test          Run unit test and code coverage
clean         Clean generated artifacts
package       Generate binary application as .zip artifact to be used on AWS Lambda
lint          Run golint to check style mistakes
fmt           List format suggestions to the code
vet           Examines the code and reports suspicious constructs
help          Show this help message
```

## Inspeção estática

Neste projeto foram adotadas ferramentas para inspecionar estaticamente o código fonte em busca de falhas de padronização e, também, construções suspeitas.

As ferramentas foram disponibilizadas em Makefile para facilitar suas chamadas, uma vez que são repetidademente invocadas pelo desenvolvedor e demandam entradas, como os pacotes do projeto ou ainda os fontes existentes.

#### golint

Um linter para projetos em Go [<link>](https://github.com/golang/lint).
Para execução no projeto pode-se utilizar o Makefile:

```shell
    make lint
```

#### gofmt
Auxilia na formatação e consequente padronização do código.
Apesar do gofmt permitir auto-correção do código fonte, a chamada via Makefile apenas apontará eventuais defeitos encontrados pela ferramenta, cabendo ao desenvolvedor corrigí-las ou, se preferir, exececutar gofmt manualmente com parâmetro "-w".

```shell
    make fmt
```

#### go vet

Para avaliação da construção de estruturas, entre outras vantagens, foi utilizado o go vet. O mesmo também esta disponível no Makefile.

```shell
    make vet
```

## Testes

Para maior confiabilidade e identificação precoce de eventuais falhas do código, foram elaborados testes unitários no core da aplicação, i.e., no pacote responsável por encontrar a posição da nave e, também, que decrifra a mensagem recebida.

Para execução dos testes e conseguinte análise de cobertura do código faça:
```shell
make test
```

Exemplo de saída:
```
andre@laptop:~/go/src/aws-api-lambda-dynamodb$ make test

Test and coverage
go test -coverpkg=./... -coverprofile=profile.cov ./...
?   	aws-api-lambda-dynamodb/cmd/main	[no test files]
?   	aws-api-lambda-dynamodb/internal/db	[no test files]
?   	aws-api-lambda-dynamodb/internal/handler	[no test files]
?   	aws-api-lambda-dynamodb/internal/http	[no test files]
ok  	aws-api-lambda-dynamodb/internal/satellite	0.015s	

coverage: 19.1% of statements in ./...
?   	aws-api-lambda-dynamodb/pkg/dynamodb	[no test files]
go tool cover -func profile.cov
aws-api-lambda-dynamodb/cmd/main/main.go:12:				handlers				0.0%
aws-api-lambda-dynamodb/cmd/main/main.go:38:				main					0.0%
aws-api-lambda-dynamodb/internal/db/db.go:10:				GetAllSatellites			0.0%
aws-api-lambda-dynamodb/internal/db/db.go:26:				UpdateSingleSatellite			0.0%
aws-api-lambda-dynamodb/internal/db/db.go:38:				UpdateMultipleSatellites		0.0%
aws-api-lambda-dynamodb/internal/handler/GetShipData.go:16:		GetShipData				0.0%
aws-api-lambda-dynamodb/internal/handler/PostMultipleSatellites.go:17:	PostMultipleSatellites			0.0%
aws-api-lambda-dynamodb/internal/handler/PostSingleSatellite.go:17:	PostSingleSatellite			0.0%
aws-api-lambda-dynamodb/internal/http/http.go:19:			ClientError				0.0%
aws-api-lambda-dynamodb/internal/satellite/findShip.go:9:		FindShip				0.0%
aws-api-lambda-dynamodb/internal/satellite/location.go:19:		round					100.0%
aws-api-lambda-dynamodb/internal/satellite/location.go:23:		toFixed					100.0%
aws-api-lambda-dynamodb/internal/satellite/location.go:28:		calculateThreeCircleIntersection	75.8%
aws-api-lambda-dynamodb/internal/satellite/location.go:108:		GetLocation				0.0%
aws-api-lambda-dynamodb/internal/satellite/message.go:8:		deleteElementFromArray			100.0%
aws-api-lambda-dynamodb/internal/satellite/message.go:16:		deleteEmptyWord				100.0%
aws-api-lambda-dynamodb/internal/satellite/message.go:31:		deletePreviousWord			0.0%
aws-api-lambda-dynamodb/internal/satellite/message.go:50:		deleteDuplicatedWord			0.0%
aws-api-lambda-dynamodb/internal/satellite/message.go:63:		GetMessage				0.0%
aws-api-lambda-dynamodb/pkg/dynamodb/dynamodb.go:19:			Scan					0.0%
aws-api-lambda-dynamodb/pkg/dynamodb/dynamodb.go:29:			GetItemSatellite			0.0%
aws-api-lambda-dynamodb/pkg/dynamodb/dynamodb.go:41:			NewItem					0.0%
aws-api-lambda-dynamodb/pkg/dynamodb/dynamodb.go:56:			PutItem					0.0%
total:							(statements)				19.1%
```

## Pipeline 

O projeto também contempla pipeline de integração contínua. Para ter acesso basta qualquer commit em branches develop, main e feature/.

## Backlog de evolução

Apesar do projeto atender o objetivo proposto há diversas possibilidades de melhoria para serem feitas com o investimento de mais horas de desenvolvimento.

Destas melhorias, particularmente, eu, atuaria em frentes para garantir melhor resultado na entrega, tais como:
* Aumento da cobertura de testes unitários
* Testes regressivos com mocks
* Pipeline de CD



# obrigado


