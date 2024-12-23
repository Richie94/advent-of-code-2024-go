package main

import (
	"advent-of-code-2024/util"
	"sort"
	"strings"
)

func Part1(fileName string) int {
	lines, _ := util.ReadFileAsArray(fileName)
	graph := buildGraph(lines)

	// Find and print triangles
	triangles := findTriangles(graph)
	// find triangles which contaain some string with prefix t
	count := 0
	for _, triangle := range triangles {
		for _, node := range triangle {
			if strings.HasPrefix(node, "t") {
				count++
				break
			}
		}
	}
	return count
}

func Part2(fileName string) string {
	lines, _ := util.ReadFileAsArray(fileName)
	graph := buildGraph(lines)
	largestClique := findLargestClique(graph)
	sort.Slice(largestClique, func(i, j int) bool {
		return largestClique[i] < largestClique[j]
	})
	return strings.Join(largestClique, ",")
}

// Find the largest clique (maximal fully connected subgraph)
func findLargestClique(graph map[string]map[string]bool) []string {
	var largestClique []string

	// Start the Bron-Kerbosch algorithm for all nodes
	for node := range graph {
		// Candidates: neighbors of the node
		candidates := make(map[string]bool)
		for neighbor := range graph[node] {
			candidates[neighbor] = true
		}

		// Excluded set starts empty
		excluded := make(map[string]bool)

		// Pivot set starts with the current node
		bronKerboschLargest([]string{node}, candidates, excluded, graph, &largestClique)
	}

	return largestClique
}

// Bron-Kerbosch recursive algorithm to find the largest clique
func bronKerboschLargest(pivotSet []string, candidates, excluded map[string]bool, graph map[string]map[string]bool, largestClique *[]string) {
	// If the candidate set is empty and the excluded set is empty, we found a maximal clique
	if len(candidates) == 0 && len(excluded) == 0 {
		// If this clique is the largest we've seen, update the largestClique result
		if len(pivotSet) > len(*largestClique) {
			*largestClique = append([]string{}, pivotSet...) // Copy pivotSet to largestClique
		}
		return
	}

	// Iterate over candidates
	for node := range candidates {
		// Add the current node to the pivot set (grow the clique)
		newPivotSet := append(pivotSet, node)

		// Filter neighbors of the current node for the next recursive steps
		newCandidates := intersect(candidates, graph[node])
		newExcluded := intersect(excluded, graph[node])

		// Recursive call
		bronKerboschLargest(newPivotSet, newCandidates, newExcluded, graph, largestClique)

		// Remove the node from candidates and add it to excluded
		delete(candidates, node)
		excluded[node] = true
	}
}

// Utility: Find the intersection of two sets
func intersect(set1, set2 map[string]bool) map[string]bool {
	result := make(map[string]bool)
	for key := range set1 {
		if set2[key] {
			result[key] = true
		}
	}
	return result
}

// Build a graph from edges
func buildGraph(connections []string) map[string]map[string]bool {
	graph := make(map[string]map[string]bool)

	for _, connection := range connections {
		// Split into two nodes
		parts := strings.Split(connection, "-")
		a, b := parts[0], parts[1]

		// Add edge a -> b
		if _, exists := graph[a]; !exists {
			graph[a] = make(map[string]bool)
		}
		graph[a][b] = true

		// Add edge b -> a
		if _, exists := graph[b]; !exists {
			graph[b] = make(map[string]bool)
		}
		graph[b][a] = true
	}

	return graph
}

// Find all triangles in the graph (optimized implementation)
func findTriangles(graph map[string]map[string]bool) [][]string {
	triangles := make([][]string, 0)
	seen := make(map[string]bool)

	// Iterate over each node in the graph
	for nodeA, neighborsA := range graph {
		// Iterate over each pair of neighbors of nodeA
		for nodeB := range neighborsA {
			if nodeB <= nodeA {
				continue // Avoid duplicate checks (guarantee order: nodeA < nodeB)
			}
			for nodeC := range graph[nodeB] {
				if nodeC <= nodeB || nodeC == nodeA {
					continue // Ensure nodeA < nodeB < nodeC
				}
				// Check if nodeC connects back to nodeA to form a triangle
				if graph[nodeC][nodeA] {
					// Create a sorted identifier for the triangle (e.g., "a,b,c")
					triangleKey := createTriangleKey(nodeA, nodeB, nodeC)
					if !seen[triangleKey] {
						triangles = append(triangles, []string{nodeA, nodeB, nodeC})
						seen[triangleKey] = true // Mark this triangle as seen
					}
				}
			}
		}
	}

	return triangles
}

// Create a unique key for a triangle by sorting the nodes alphabetically
func createTriangleKey(nodeA, nodeB, nodeC string) string {
	nodes := []string{nodeA, nodeB, nodeC}
	sort.Strings(nodes) // Ensure consistent order
	return strings.Join(nodes, ",")
}
