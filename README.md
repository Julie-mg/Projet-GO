# Projet-GO

1. Introduction: Provide a brief overview of the project, what it does, and what problem it solves.

2. Installation: Detail the steps necessary to install the project, including any required dependencies and how to obtain them.

3. Usage: Explain how to use the project, including any command-line arguments or configuration files.

4. Configuration: If necessary, describe the configuration options and how to use them.

5. Examples: Provide examples of how to use the project, if possible.

6. Troubleshooting: List common problems and their solutions.

7. Contributing: Provide information about how to contribute to the project, such as submitting bug reports or making pull requests.

8. License: Include information about the license under which the project is distributed.

9. Contact Information: Provide contact information for the project, such as the author's email address.

10. Final Thoughts: Provide any final thoughts or comments about the project.
This code contains a Go implementation of a client and a server. The client opens a matrice file, reads its contents, and sends it to the server. The server converts the received string into a 2D integer array, performs Dijkstra's shortest path algorithm on it, and sends the results back to the client. The client displays the results.

The client uses the net package to create a connection to the server, and the bufio package to read data from the matrice file and to receive data from the server. The io package is used to write the data to the server. The os package is used to open the matrice file.

The server uses the net package to create a listening connection, and the bufio and io packages to receive data from the client and to write data to the client. The strconv package is used to convert strings to integers, and the strings package is used to split the received strings into individual values. The Dijkstra's algorithm implementation uses slices and maps to store the distances and predeccessors. The algorithm results are saved in a channel, which is passed as a parameter to the function.
