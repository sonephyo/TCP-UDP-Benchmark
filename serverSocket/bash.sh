go build

env GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -o amd64ServerSocket

scp -r -p ./ pphyo@cs.oswego.edu:/home/pphyo/CSC445/assignment1 

ssh pphyo@rho.cs.oswego.edu

