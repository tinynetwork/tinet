package main

import (
    "fmt"
    "math/rand"
    "os"
    "time"

    "github.com/spf13/cobra"
)

var config struct {
    arg1 string
    arg2 string
}

func newCommand() *cobra.Command {
    cmd := &cobra.Command{
        Use:  "cmd",
        RunE: appMain,
    }
    cmd.Flags().StringVar(&config.arg1, "arg1", "def1", "this is arg1")
    cmd.Flags().StringVar(&config.arg2, "arg2", "def2", "this is arg2")
    return cmd
}

func appMain(cmd *cobra.Command, args []string) error {
    fmt.Printf("arg1=%s, arg2=%s\\n", config.arg1, config.arg2)
    return nil
}

func main() {
    rand.Seed(time.Now().UnixNano())
    command := newCommand()
    err := command.Execute()
    if err != nil {
        os.Exit(1)
    }
}
