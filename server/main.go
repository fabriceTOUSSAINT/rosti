package main

import (
	"fmt"

	integers "github.com/fabriceTOUSSAINT/rosti/server/adder"
	"github.com/fabriceTOUSSAINT/rosti/server/hello"
	"github.com/spf13/cobra"
)

func main() {
	// fmt.Println("Hello, Modules!")
	// fmt.Println(hello.Hello("Fab", "French"))
	// fmt.Println(integers.Add(2,2))

	cmd := &cobra.Command{
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println(hello.Hello("Fabrice", "French"))
			fmt.Println(integers.Add(2,2))
		},
	}

	fmt.Println("Calling cmd.Execute()!")
	cmd.Execute()
}