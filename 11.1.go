package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')
	input = strings.TrimSpace(input)

	numbers := strings.Split(input, " ")
	votes := make(map[int]int)
	totalVotes := -1 // -1 to not count the 0 at end
	validVotes := 0

	for _, num := range numbers {
		totalVotes++
		val, err := strconv.Atoi(num)
		if err != nil || val == 0 {
			continue
		}

		if val >= 1 && val <= 20 {
			votes[val]++
			validVotes++
		}
	}

	fmt.Printf("Suara masuk: %d\n", totalVotes)
	fmt.Printf("Suara sah: %d\n", validVotes)

	for i := 1; i <= 20; i++ {
		if votes[i] > 0 {
			fmt.Printf("%d: %d\n", i, votes[i])
		}
	}
}
