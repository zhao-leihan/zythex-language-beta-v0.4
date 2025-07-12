package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var vars = make(map[string]interface{})

func evalCondition(cond string) bool {
	cond = strings.TrimSpace(cond)
	if strings.Contains(cond, "==") {
		parts := strings.Split(cond, "==")
		left := strings.TrimSpace(parts[0])
		right := strings.TrimSpace(parts[1])
		leftVal := vars[left]
		if num, err := strconv.Atoi(right); err == nil {
			return fmt.Sprintf("%v", leftVal) == fmt.Sprintf("%d", num)
		}
		return fmt.Sprintf("%v", leftVal) == right
	}
	if cond == "true" {
		return true
	}
	if cond == "false" {
		return false
	}
	if num, err := strconv.Atoi(cond); err == nil {
		return num != 0
	}
	return false
}

func runLine(line string) {
	line = strings.TrimSpace(line)
	if line == "" {
		return
	}

	if strings.HasPrefix(line, "zylet ") {
		parts := strings.SplitN(line[6:], "=", 2)
		if len(parts) == 2 {
			name := strings.TrimSpace(parts[0])
			value := strings.TrimSpace(parts[1])
			if intVal, err := strconv.Atoi(value); err == nil {
				vars[name] = intVal
			} else {
				vars[name] = value
			}
		}
	} else if strings.HasPrefix(line, "zyth.print <") && strings.HasSuffix(line, ">") {
		content := strings.TrimSuffix(strings.TrimPrefix(line, "zyth.print <"), ">")
		if val, ok := vars[content]; ok {
			fmt.Println(val)
		} else {
			fmt.Println(content)
		}
	} else if strings.HasPrefix(line, "zuicide") {
		os.Exit(0)
	} else if strings.HasPrefix(line, "zythract ") {
		name := strings.TrimSuffix(strings.TrimPrefix(line, "zythract "), ":")
		fmt.Println("=== Contract:", name, "===")
	} else if strings.HasPrefix(line, "zyturn ") {
		out := strings.TrimSpace(strings.TrimPrefix(line, "zyturn "))
		fmt.Println("Return:", out)
		os.Exit(0)
	}
}

func main() {
	if len(os.Args) < 3 || os.Args[1] != "run" {
		fmt.Println("Usage: zythx run <file.zthx>")
		return
	}

	file := os.Args[2]
	f, err := os.Open(file)
	if err != nil {
		fmt.Println("Failed to read file:", err)
		return
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	var lines []string
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	executing := true

	for i := 0; i < len(lines); i++ {
		line := strings.TrimSpace(lines[i])
		if line == "" {
			continue
		}

		if !executing {
			// Skip lines until else block
			if strings.HasPrefix(line, "zelxe:") {
				executing = true
			} else if strings.HasPrefix(line, "zyif") || strings.HasPrefix(line, "zyif (") {
				// Skip nested ifs
				nested := 1
				for i+1 < len(lines) {
					i++
					next := strings.TrimSpace(lines[i])
					if strings.HasPrefix(next, "zyif") {
						nested++
					} else if next == "zelxe:" && nested == 1 {
						break
					} else if next == "zelxe:" {
						nested--
					}
				}
			}
			continue
		}

		if strings.HasPrefix(line, "zyif (") && strings.HasSuffix(line, "):") {
			cond := strings.TrimSuffix(strings.TrimPrefix(line, "zyif ("), "):")
			condResult := evalCondition(cond)
			if condResult {
				executing = true
			} else {
				executing = false
			}
		} else if line == "zelxe:" {
			executing = !executing
		} else {
			runLine(line)
		}
	}
}
