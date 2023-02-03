# GO server-client Dijkstra

This is a Go (Golang) program that implements a TCP server. The server listens for incoming connections on port 8081 and creates a go routine for each client to handle the communication. The program receives a matrix of integers from the client, and for each pair of nodes in the graph represented by the matrix, it calculates the shortest path between them using Dijkstra's algorithm and sends the result back to the client. The result includes the start and end nodes, the distance between them, and the previous node in the shortest path.

## Requirements

* GO for windows or linux (tested on `go version go1.19.4 windows/amd64` and `go version go1.15.9 linux/amd64`)

## Running the Application

To run the application, install Golang, clone this repository or download the files in the same repository on your computer, and use the following commands in two different terminals open in the same repository (where the files of this application are saved):

`go run server_dijkstra.go`
`go run client_dijkstra.go`

## Functionality

* Multiple clients can ask simultaneously to calculate the Dijkstra on different matrix (one at a time per client)

The client opens a matrice file, reads its contents, and sends it to the server. The server converts the received string into a 2D integer array, performs Dijkstra's shortest path algorithm on it, and sends the results back to the client. 
The client displays the results.

The program uses the bufio, fmt, io, net, strconv, and strings packages.

