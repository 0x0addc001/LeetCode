package main

import "fmt"

func myCalcEquation(equations [][]string, values []float64, queries [][]string) []float64 {
	// 方法1：Floyd算法
	chars := make([]string, 0)
	charSet := make(map[string]bool)
	for _, eq := range equations {
		if !charSet[eq[0]] {
			charSet[eq[0]] = true
			chars = append(chars, eq[0])
		}
		if !charSet[eq[1]] {
			charSet[eq[1]] = true
			chars = append(chars, eq[1])
		}
	}
	length := len(chars)
	graph := make([][]float64, length)
	for i := range graph {
		graph[i] = make([]float64, length) //  虽然外层切片已经初始化，但其内部的每个 graph[i] 还需被分配内存空间
		for j := range graph[i] {
			graph[i][j] = -1.0
		}
		graph[i][i] = 1.0
	}

	for i, eq := range equations {
		x := myIndexOf(chars, eq[0])
		y := myIndexOf(chars, eq[1])
		graph[x][y] = values[i]
		graph[y][x] = 1.0 / values[i]
	}

	for k := 0; k < length; k++ {
		for i := 0; i < length; i++ {
			for j := 0; j < length; j++ {
				if graph[i][j] == -1.0 && graph[k][i] != -1.0 && graph[k][j] != -1.0 {
					graph[i][j] = graph[i][k] * graph[k][j]
				}
			}
		}
	}

	res := make([]float64, len(queries))
	for i, q := range queries {
		x := myIndexOf(chars, q[0])
		y := myIndexOf(chars, q[1])
		if x == -1 || y == -1 {
			res[i] = -1.0
		} else {
			res[i] = graph[x][y]
		}
	}
	return res

	// 方法2：BFS
	/*
		graph := make(map[string][]MyPair)
		for i, eq := range equations {
			a, b := eq[0], eq[1]
			val := values[i]
			graph[a] = append(graph[a], MyPair{b, val})
			graph[b] = append(graph[b], MyPair{a, 1.0 / val})
		}

		res := make([]float64, len(queries))
		for i, q := range queries {
			num, den := q[0], q[1]
			if _, ok := graph[num]; !ok {
				res[i] = -1.0
			} else if _, ok := graph[den]; !ok {
				res[i] = -1.0
			} else if num == den {
				res[i] = 1.0
			} else {
				res[i] = mybfs(num, den, graph)
			}
		}

		return res
	*/

	// 方法3：DFS
	/*
		graph := make(map[string]map[string]float64)
		for i, eq := range equations {
			a, b := eq[0], eq[1]
			val := values[i]
			if _, ok := graph[a]; !ok {
				graph[a] = make(map[string]float64)
			}
			if _, ok := graph[b]; !ok {
				graph[b] = make(map[string]float64)
			}
			graph[a][b] = val
			graph[b][a] = 1 / val
		}

		res := make([]float64, len(queries))
		for i, query := range queries {
			num, den := query[0], query[1]
			if _, ok := graph[num]; !ok {
				res[i] = -1.0
			} else if _, ok := graph[den]; !ok {
				res[i] = -1.0
			} else if num == den {
				res[i] = 1.0
			} else {
				visited := make(map[string]bool)
				res[i] = mydfs(num, den, graph, 1.0, visited)
			}
		}
		return res
	*/
}

func myIndexOf(chars []string, s string) int {
	for i, c := range chars {
		if c == s {
			return i
		}
	}
	return -1
}

func mybfs(num string, den string, graph map[string][]MyPair) float64 {
	queue := []MyNode{{num, 1.0}}
	visited := make(map[string]bool)
	visited[num] = true
	for len(queue) > 0 {
		node := queue[0]
		queue = queue[1:]
		for _, pair := range graph[node.cur] { // 注意'_' 不是 for pair :=range graph[node.cur]
			if pair.to == den {
				return node.val * pair.val
			}
			if !visited[pair.to] {
				visited[pair.to] = true // 注意将节点加入队列时，需要立即标记为已访问，避免重复处理
				queue = append(queue, MyNode{pair.to, node.val * pair.val})
			}
		}
	}
	return -1.0
}

func mydfs(num string, den string, graph map[string]map[string]float64, product float64, visited map[string]bool) float64 {
	if num == den {
		return product
	}
	visited[num] = true
	for neighbor, val := range graph[num] {
		if !visited[neighbor] {
			result := mydfs(neighbor, den, graph, product*val, visited)
			if result != -1.0 {
				return result
			}
		}
	}
	return -1.0
}

type MyPair struct {
	to  string
	val float64
}
type MyNode struct {
	cur string
	val float64
}

func main() {
	// 测试代码
	equations := [][]string{{"a", "b"}, {"b", "c"}}
	values := []float64{2.0, 3.0}
	queries := [][]string{{"a", "c"}, {"b", "a"}, {"a", "e"}, {"a", "a"}, {"x", "x"}}
	fmt.Println(myCalcEquation(equations, values, queries))
	// 预期输出: [6.0 0.5 -1.0 1.0 -1.0]
}
