# samplepb
This is a sample development environment for protocol buffers. Install golang and the protocol buffer development environment in a Docker container. By installing the vscode remote environment in the container, you can enter Docker and work with vscode.

## Usage
Download this project.
```
git clone https://github.com/tkhsnt0/samplepb.git
```
Create docker image and run.
```
docker-compose up -d
```
Run Visual Studio Code.
```
code .
```
After launching Visual Studio Code, attach to this running container.
Execute makeproto.sh
```
./makeproto.sh
```
By executing the shell script, protoc will be executed and the gRPC code will be automatically generated.