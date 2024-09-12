package main

import (
  "fmt"
  "github.com/spf13/cobra"
)

func main(){
  rootCmd := &cobra.Command{
    Run: func(cmd *cobra.Command, args []string){
      fmt.Println(len(args))
    },
  }

  rootCmd.Execute()

  fmt.Println("HELLO, WOLRD!")
}
