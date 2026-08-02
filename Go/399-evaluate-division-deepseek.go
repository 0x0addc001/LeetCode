package main

import (
	"fmt"
)

func calcEquation(equations [][]string, values []float64, queries [][]string) []float64 {
	// 方法1：Floyd算法
	// 收集所有字符
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
	// 初始化图
	graph := make([][]float64, length)
	for i := range graph {
		graph[i] = make([]float64, length)
		for j := range graph[i] {
			graph[i][j] = -1.0
		}
		graph[i][i] = 1.0
	}

	// 更新图的初始值
	for i, eq := range equations {
		x := indexOf(chars, eq[0])
		y := indexOf(chars, eq[1])
		graph[x][y] = values[i]
		graph[y][x] = 1.0 / values[i]
	}

	// Floyd算法计算所有点对的最短路径
	for k := 0; k < length; k++ {
		for i := 0; i < length; i++ {
			for j := 0; j < length; j++ {
				if graph[i][k] != -1 && graph[k][j] != -1 && graph[i][j] == -1 {
					graph[i][j] = graph[i][k] * graph[k][j]
				}
			}
		}
	}

	// 处理查询
	res := make([]float64, len(queries))
	for i, q := range queries {
		x := indexOf(chars, q[0])
		y := indexOf(chars, q[1])
		if x == -1 || y == -1 {
			res[i] = -1.0
		} else {
			res[i] = graph[x][y]
		}
	}
	return res

	// 方法2：BFS
	/*
		graph := make(map[string][]Pair)
		for i, eq := range equations {
			a, b := eq[0], eq[1]
			val := values[i]
			graph[a] = append(graph[a], Pair{b, val})
			graph[b] = append(graph[b], Pair{a, 1.0 / val})
		}

		res := make([]float64, len(queries))
		for i, q := range queries {
			num, den := q[0], q[1]
			if _, ok := graph[num]; !ok {
				res[i] = -1.0
				continue
			}
			if _, ok := graph[den]; !ok {
				res[i] = -1.0
				continue
			}
			if num == den {
				res[i] = 1.0
				continue
			}

			queue := []Node{{num, 1.0}}
			visited := make(map[string]bool)
			visited[num] = true
			found := false

			for len(queue) > 0 && !found {
				node := queue[0]
				queue = queue[1:]

				for _, pair := range graph[node.name] {
					nb, val := pair.to, pair.val
					if nb == den {
						res[i] = node.val * val
						found = true
						break
					}
					if !visited[nb] {
						visited[nb] = true
						queue = append(queue, Node{nb, node.val * val})
					}
				}
			}

			if !found {
				res[i] = -1.0
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
			graph[b][a] = 1.0 / val
		}

		res := make([]float64, len(queries))
		for i, q := range queries {
			num, den := q[0], q[1]
			if _, ok := graph[num]; !ok {
				res[i] = -1.0
				continue
			}
			if _, ok := graph[den]; !ok {
				res[i] = -1.0
				continue
			}
			if num == den {
				res[i] = 1.0
				continue
			}

			visited := make(map[string]bool)
			res[i] = dfs(graph, num, den, 1.0, visited)
		}
		return res
	*/
}

// DFS递归实现
func dfs(graph map[string]map[string]float64, current, target string, product float64, visited map[string]bool) float64 {
	if current == target {
		return product
	}

	visited[current] = true

	for neighbor, val := range graph[current] {
		if !visited[neighbor] {
			result := dfs(graph, neighbor, target, product*val, visited)
			if result != -1.0 {
				return result
			}
		}
	}

	return -1.0
}

// 辅助结构体
type Pair struct {
	to  string
	val float64
}

type Node struct {
	name string
	val  float64
}

// 辅助函数：查找字符串在切片中的索引
func indexOf(chars []string, s string) int {
	for i, c := range chars {
		if c == s {
			return i
		}
	}
	return -1
}

func main2() {
	// 测试代码
	equations := [][]string{{"a", "b"}, {"b", "c"}}
	values := []float64{2.0, 3.0}
	queries := [][]string{{"a", "c"}, {"b", "a"}, {"a", "e"}, {"a", "a"}, {"x", "x"}}
	fmt.Println(calcEquation(equations, values, queries))
	// 预期输出: [6.0 0.5 -1.0 1.0 -1.0]
}
