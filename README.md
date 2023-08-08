# fedha-go-gen.grpc
This repository contains the generated Go code for the fedha project. The code is generated from the proto files in the db directory.

To include the generated code in your project, you can use the following command:

go mod edit -replace=github.com/MikeMwita/fedha-go-gen.grpc=/home/mike/Desktop/fedha-go-gen.grpc

The /home/mike/Desktop/fedha-go-gen.grpc part of the command is the path to the directory where you have cloned this repository. You will need to replace this path with the path to the directory on your computer.

Once you have run this command, the generated code will be added to your project's Go modules. You can then use the generated code in your project.

Here are some additional notes:


The generated code is compatible with Go 1.18 or later.
