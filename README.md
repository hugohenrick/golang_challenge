# gRPC Go Challenge

## Notes

### `Windows`

- I recommend you use powershell (try to update: [see](https://github.com/PowerShell/PowerShell/releases)) for following this course, you might have unexepected behavior if you use Git bash or other (especially with OpenSSL)
- I recommend you use [Chocolatey](https://chocolatey.org/) as package installer (see [Install](https://chocolatey.org/install))

### Build

#### `Linux/MacOS`

```shell
make all
```

#### `Windows - Chocolatey`
```shell
choco install make
make all
```

#### `Windows - Without Chocolatey`

```shell

protoc -Icoin/proto --go_opt=module=github.com/hugohenrick/golang_challenge --go_out=. --go-grpc_opt=module=github.com/hugohenrick/golang_challenge --go-grpc_out=. coin/proto/*.proto

Start MongoDB:
acces folder coin and run: docker compose up
go build -o bin/coin/server.exe ./coin/server
go build -o bin/coin/client.exe ./coin/client
```

References: https://www.udemy.com/course/grpc-golang/
