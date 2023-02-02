# Projet-GO

This code contains a Go implementation of a client and a server. The client opens a matrice file, reads its contents, and sends it to the server. The
server converts the received string into a 2D integer array, performs Dijkstra's shortest path algorithm on it, and sends the results back to the client. 
The client displays the results.

To start to use this program, need to do the following steps:
- Download the three files and put them in the same directory
- To run the CLI : first type "go run server_dijkstra.go" in a terminal, then type "go run client_dijkstra.go" in another terminal, both terminal opened in the directory of the files
