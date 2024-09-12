package commands

import (
  "github.com/spf13/cobra"
  "github.com/Guilherme-Matosoli/FnWrapper/internal/handlers"
)


func howamiCommand() *cobra.Command{
  command := &cobra.Command{
    Use: "whoami",
    Run: handlers.HowamiHandler,
  }

  return command
}
