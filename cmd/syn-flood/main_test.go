package main

import (
	"fmt"
	"github.com/bilalcaliskan/syn-flood/internal/options"
	"testing"
)

func TestMainProgram(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main function", r)
		}
	}()

	sfo := options.GetSynFloodOptions()
	sfo.FloodDurationSeconds = 1
	main()
}
