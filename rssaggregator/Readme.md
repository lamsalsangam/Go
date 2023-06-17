# RSS Aggregator

- Initialize the go.mod file with `go mod init [pathname]` you can replace the pathname with your github repository name

- To add the other modules in your go programme you just need to do `go get [pathname]` in this case the github.com/joho/godotenv as the pathname

- Type `go mod tidy` it will clean up the modules

- You can make the copy locally by doing `go mod vendor`

- To spin up the server you need to install `go get [pathname]` in this case we use github.com/go-chi/cors
