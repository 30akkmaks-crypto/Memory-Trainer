// memory_trainer.go
package main

import (
	"bufio"
	"flag"
	"fmt"
	"math/rand"
	"os"
	"os/exec"
	"runtime"
	"strconv"
	"strings"
	"time"
)

func clearScreen() {
	cmd := "clear"
	if runtime.GOOS == "windows" {
		cmd = "cls"
	}
	c := exec.Command(cmd)
	c.Stdout = os.Stdout
	c.Run()
}

func generateSequence(length int) []int {
	seq := make([]int, length)
	for i := range seq {
		seq[i] = rand.Intn(10)
	}
	return seq
}

func displaySequence(seq []int, showTime int) {
	fmt.Println("Remember these numbers:")
	for _, n := range seq {
		fmt.Printf("%d ", n)
	}
	fmt.Println()
	time.Sleep(time.Duration(showTime) * time.Second)
	clearScreen()
}

func getUserInput(length int) []int {
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Printf("Enter the %d numbers (space‑separated): ", length)
		line, _ := reader.ReadString('\n')
		line = strings.TrimSpace(line)
		parts := strings.Fields(line)
		if len(parts) != length {
			fmt.Printf("Please enter exactly %d numbers.\n", length)
			continue
		}
		seq := make([]int, length)
		ok := true
		for i, p := range parts {
			n, err := strconv.Atoi(p)
			if err != nil {
				fmt.Println("Please enter valid integers.")
				ok = false
				break
			}
			seq[i] = n
		}
		if ok {
			return seq
		}
	}
}

func playRound(length int, showTime int) bool {
	seq := generateSequence(length)
	displaySequence(seq, showTime)
	userSeq := getUserInput(length)
	for i := range seq {
		if seq[i] != userSeq[i] {
			return false
		}
	}
	return true
}

func main() {
	var (
		length     = flag.Int("l", 0, "Number of digits per sequence")
		difficulty = flag.String("d", "", "Preset: easy, medium, hard")
		rounds     = flag.Int("r", 1, "Number of rounds")
	)
	flag.Parse()

	var seqLen int
	if *length > 0 {
		seqLen = *length
	} else if *difficulty != "" {
		switch *difficulty {
		case "easy":
			seqLen = 4
		case "medium":
			seqLen = 6
		case "hard":
			seqLen = 8
		default:
			fmt.Println("Invalid difficulty. Use easy, medium, hard.")
			os.Exit(1)
		}
	} else {
		seqLen = 6
	}
	if seqLen < 1 {
		fmt.Println("Length must be at least 1.")
		os.Exit(1)
	}
	if *rounds < 1 {
		fmt.Println("Rounds must be at least 1.")
		os.Exit(1)
	}

	rand.Seed(time.Now().UnixNano())
	fmt.Printf("\n🧠 Memory Trainer\n")
	fmt.Printf("Difficulty: %s (%d digits)\n", func() string {
		if *difficulty != "" {
			return *difficulty
		}
		return "custom"
	}(), seqLen)
	fmt.Printf("Rounds: %d\n", *rounds)

	correct := 0
	showTime := 2
	for r := 1; r <= *rounds; r++ {
		fmt.Printf("\nRound %d/%d\n", r, *rounds)
		if playRound(seqLen, showTime) {
			fmt.Println("✅ Correct!")
			correct++
		} else {
			fmt.Println("❌ Wrong.")
		}
	}
	fmt.Printf("\nScore: %d/%d\n", correct, *rounds)
}
