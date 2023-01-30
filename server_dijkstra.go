package main

import (
	"bufio"
	"fmt"
	"io"
	"net"
	"strconv"
	"strings"
)

type Node struct {
	start    int
	end      int
	distance int
	pred     int
}

func affichage_matrice(tab [][]int) {
	// print the [][]int into the terminal
	for i := 0; i < len(tab); i++ {
		for j := 0; j < len(tab[0]); j++ {
			fmt.Printf("%d ", tab[i][j])
		}
		fmt.Print("\n")
	}
}

func recep(conn net.Conn, matrix [][]int) [][]int {
	// convert string received into [][]int
	for {
		netData, err := bufio.NewReader(conn).ReadString('\n')
		if err != nil {
			fmt.Println(err)
			break
		} else if netData == "\n" {
			break
			//return err
		}

		var l []int
		split := strings.Split(netData, " ")
		for i := 0; i < len(split); i++ {
			val, _ := strconv.Atoi(split[i])
			l = append(l, val)
		}
		matrix = append(matrix, l)
	}
	affichage_matrice(matrix)
	return matrix
}

func Dijkstra(graph [][]int, start int, end int, ch chan Node) {
	// calculate Dijkstra from the start point to the end point and save the data in the channel ch
	n := len(graph)
	dist := make([]int, n)
	for i := 0; i < n; i++ {
		dist[i] = 9999
	}
	dist[start] = 0
	pred := start

	unvisited := make(map[int]bool)
	for i := 0; i < n; i++ {
		unvisited[i] = true
	}

	current := start
	for current != end {
		unvisited[current] = false

		for i, weight := range graph[current] {
			if weight != 0 {
				newDistance := dist[current] + weight
				if newDistance < dist[i] {
					dist[i] = newDistance
				}
			}
		}

		pred = current
		current = -1
		for i, visited := range unvisited {
			if visited && (current == -1 || dist[i] < dist[current]) {
				current = i
			}
		}
	}
	ch <- Node{start, end, dist[end], pred}
}

func Client(conn net.Conn) {
	// a go routine is created for each client to calculate Dijkstra simultaneously for each of them
	var matrix [][]int
	ch := make(chan Node, 10000)

	matrix = recep(conn, matrix)

	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix); j++ {
			go Dijkstra(matrix, i, j, ch)
		}
	}

	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix); j++ {
			node := <-ch

			fmt.Printf("%d to %d have distance %d with pred node %d\n", node.start, node.end, node.distance, node.pred)
			io.WriteString(conn, fmt.Sprintf("%d to %d have distance %d with pred node %d\n", node.start, node.end, node.distance, node.pred))
		}
		fmt.Print("\n")
		io.WriteString(conn, fmt.Sprintf("\n"))
	}
	fmt.Println("fin_server")
	io.WriteString(conn, fmt.Sprintf("fin\n"))

}

func main() {

	fmt.Println("Launching server...")

	// listen on all interfaces
	ln, _ := net.Listen("tcp", ":8081")

	for {
		// Listen for an incoming connection.
		conn, err := ln.Accept()
		if err != nil {
			fmt.Println("Error accepting: ", err.Error())
			continue
		}
		go Client(conn)

	}
}
