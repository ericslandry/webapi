# Go OpenAPI-generated Web Service in Dev Container

This repository contains a Go web service example project that is meant to be ran in a [dev container](https://code.visualstudio.com/docs/devcontainers/containers).

## Prerequisites

- [Install Docker](https://docs.docker.com/get-docker/)
- [Install VSCode](https://code.visualstudio.com/) with [Dev Containers extension](https://marketplace.visualstudio.com/items?itemName=ms-vscode-remote.remote-containers)

## Getting Started

1. Clone this repository:
   ```shell
   git clone https://github.com/ericslandry/webapi.git
   ```
2. Open the repository in VSCode:
   ```shell
   code grpc
   ```
3. Reopen the repository in a container:
      - Click on the green icon in the bottom left corner of the window and select `Reopen in Container`.
      - Alternatively, you can open the command palette (Ctrl+Shift+P) and type `Remote-Containers: Reopen in Container`.

   After a reload, VSCode should display "Dev Container: Go" at the bottom left:
   ![Dev Container: Go](./docs/devContainer.png)

4. Run the `default` [Taskfile](https://taskfile.dev/) target:
   ```shell
   task
   ```
   You should see the following output:
   ```
   vscode ➜ /workspaces/grpc (master) $ task
   task: [build:proto] protoc *.proto --proto_path=. --go_out=. --go_opt=module=github.com/ericslandry/grpc/pb/greeter --go-grpc_out=. --go-grpc_opt=module=github.com/ericslandry/grpc/pb/greeter
   task: [build:setup] mkdir -p ./bin
   task: [build:client] go build -o ./bin/grpc-client ./cmd/client
   go: downloading github.com/urfave/cli/v2 v2.27.1
   go: downloading google.golang.org/grpc v1.61.0
   go: downloading golang.org/x/net v0.21.0
   go: downloading google.golang.org/genproto/googleapis/rpc v0.0.0-20240205150955-31a09d347014
   go: downloading golang.org/x/sys v0.17.0
   go: downloading github.com/xrash/smetrics v0.0.0-20201216005158-039620a65673
   go: downloading github.com/cpuguy83/go-md2man/v2 v2.0.2
   go: downloading github.com/russross/blackfriday/v2 v2.1.0
   go: downloading golang.org/x/text v0.14.0
   task: [build:setup] mkdir -p ./bin
   task: [build:server] go build -o ./bin/grpc-server ./cmd/server
   task: [run:server] ./bin/grpc-server &
   task: [run:server] until grpcurl -proto ./pb/greeter.proto -plaintext localhost:8080 list > /dev/null; do printf '.'; sleep 1; done
   task: [run:client] ./bin/grpc-client --name=Mike
   2024/02/16 13:50:04 server is listening at [::]:8080
   2024/02/16 13:50:04 Received: Mike
   2024/02/16 13:50:04 Greeting: Hello, Mike
   2024/02/16 13:50:04 server is being stopped
   ```

   See `Taskfile.yml` for more details on the various targets or run `task --list`.
