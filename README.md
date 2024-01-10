# dagger-pipeline-example
An example of using Dagger.io to test, build and publish an application

## How to run the example locally

### Requirements
- [Docker](https://docs.docker.com/engine/install/) or any other OCI runtime installed

### Approach 1: Use DevContainer
This approach uses the [DevContainer](https://marketplace.visualstudio.com/items?itemName=ms-vscode-remote.remote-containers) extension to spin up a development environment with [Go](https://go.dev/) and [Dagger.io CLI](https://docs.dagger.io/cli) pre-installed.

1. Clone the repository
```bash
git clone https://github.com/techchapter/dagger-pipeline-example.git
```

2. Open the repository in Visual Studio Code
```bash
code dagger-pipeline-example
```

3. Run command "Reopen in Container" inside Visual Studio Code

4. Run the [Dagger.io CLI](https://docs.dagger.io/cli) to test and build the pipeline locally
```bash
dagger run --silent go run ./ci/ci.go
```

5. The output should look something like this
```bash
DEV environment: exported file to ./ping-socket.tar 
To run use docker load -i ./ping-socket.tar
```

### Approach 2: Install tools locally
This approach expect you to install [Go](https://go.dev/) and [Dagger.io CLI](https://docs.dagger.io/cli) on your development machine.

1. Install [Go](https://go.dev/) and [Dagger.io CLI](https://docs.dagger.io/cli)
2. Clone the repository
```bash
git clone https://github.com/techchapter/dagger-pipeline-example.git
```

3. Run the [Dagger.io CLI](https://docs.dagger.io/cli) to test and build the pipeline locally
```bash
dagger run --silent go run ./ci/ci.go
```

4. The output should look something like this
```bash
DEV environment: exported file to ./ping-socket.tar 
To run use docker load -i ./ping-socket.tar
```

## How to run the exmaple in GitHub actions

1. Fork the GitHub repository
2. Clone the forked repository
2. Create a simple change like the following
```bash
echo "RUN" >> changed-file.txt
```
3. Commit and Push the change to the forked repo
4. See the GitHub action running the CI pipeline, but this time actually publish the image to ttl.sh
5. The output should look something like this
```bash
Published image to: ttl.sh/ping-socket-2021992@sha256:876f28eb26c9469df381f03ad9c82d4648a56349c984d37b186379228ba452b1
```
