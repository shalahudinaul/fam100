package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/yulrizka/fam100"
)

var (
	input       = make(chan string)
	name        string
	roundPlayed int
	seed        int64
)

func main() {
	printHeader()
	seed = time.Now().UnixNano()

	scanner := bufio.NewScanner(os.Stdin)
	for name == "" {
		fmt.Print("Siapa nama Anda? ")
		scanner.Scan()
		name = strings.TrimSpace(scanner.Text())
	}

	fmt.Printf("\nHalo %s ketik '/keluar' kapanpun jika ingin berhenti!\n\n", name)
	go startGame()
	time.Sleep(300 * time.Millisecond)

	for scanner.Scan() {
		fmt.Print("> ")
		text := scanner.Text()
		if text == "/keluar" {
			os.Exit(0)
		}

		select {
		case input <- text:
		default:
		}
	}
}

func printHeader() {
	fmt.Println()
	fmt.Println("===============")
	fmt.Println("Famili 100 v0.1")
	fmt.Println("===============")
	fmt.Println()
}

func startGame() {
	fam100.LoadQuestion("fam100.db")
	for {
		fmt.Printf("Siap ? (y/n) ")
		ya := []string{"ya", "y", "yes"}
	READY:
		for i := range input {
			i = strings.TrimSpace(strings.ToLower(i))
			for _, v := range ya {
				if i == v {
					break READY
				}
			}
			fmt.Print("Siap ? (y/n) ")
		}
		fmt.Println()

		in := make(chan fam100.Message)
		out := make(chan fam100.Message)
		game, _ := fam100.NewGame("cli", in, out)
		game.Start()
		/*
			GAME:
					for {
						select {
						case m := <-out:
								switch m.Kind {
								case fam100.TextMessage:
									fmt.Println(m.Text)
									fmt.Println()
								case fam100.StateMessage:
									switch m.Text {
									case string(fam100.Finished):
										break GAME
									case string(fam100.RoundFinished):
										roundPlayed++
									}
								default:
									// for debuging
									//fmt.Printf("m = %+v\n", m)
								}
						case i := <-input:
							fmt.Println()
							msg := fam100.Message{
								Player: fam100.Player{ID: fam100.PlayerID(name), Name: name},
								Text:   i,
							}
							in <- msg
						}
					}
		*/
	}
}
