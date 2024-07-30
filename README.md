# Help

Comando utilizado para iniciar o projeto
`go mod init estudo`

o padrão do go é criar um projeto desse jeito:
`go mod init github.com/<USUARIO-GITHUB>/<NOME-DO-REPOSITORIO>`

Este projeto deveria ter sico criado assim:
`go mod init github.com/NatalNW7/study-go-lang`

Executa o projeto
`go run main.go `

## Go Standards

apesar das varias pastas e padrões, 90% dos projetos já se resolvem com `cmd`, `internal` e `pkg`

nome das pastas e arquivos sempre no singular

Pasta `cmd` é onde fica o arquivo `main.go`. O ideia é criar uma outra pasta com o nome da aplicação e depois criar o aqruivo `main.go`

Pasta `internal`: Tudo q é interno a sua aplicação, e q fundamentalmente será utilizado apenas para a sua aplicação, fique dentro desta pasta

Pasta `pkg`: fica os arquivos/pacotes q podem ser utilizados por outros projetos

Pasta `build`: é onde fica os arquivos q são utilizado para fazer o build da aplicação

Pasta `config`: todos os arquivos de configuração ficam aqui

Pasta `docs`: é onde ficam as docs adicionais do projeto, não as geradas pelo `go docs`

Pasta `api`: aqui não ficam os arquivos de api, aqui ficam os arquivos para documentação ou geração de documentação para api's, um exemplo: aqui ficaria o arquivo do swagger

Pasta `scripts`: aqui fica os arquivos para automação, scripts auxiliares

Pasta `test`: é onde fica os arquivos de testes gerais

Pasta `exemples`: é onde ficam os exemplos de como implementar o codigo ou como usar a sua lib

Pasta `web`: é onde ficam os seus assets q sua aplicação vai utilizar

Pasta `website`: é onde ficam os arquivos q serão o website da sua aplicação

Link para Go Standards: https://github.com/golang-standards/project-layout
