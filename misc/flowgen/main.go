package main

import (
	"fmt"
	"math/rand"
	"net"
	"os"
	"time"

	"github.com/spf13/cobra"
)

var config struct {
	arg1 string
	arg2 string
}

type IPFixMessage struct {
	Version      uint16
	Count        uint16
	SysUptime    uint32
	Timestamp    uint32
	FlowSequence uint32
	SourceID     uint32
	FlowSets     []IPFixFlowSet
}

type IPFixFlowSet struct {
	ID       uint16
	Length   uint16
	Template IPFixFlowTemplate
}

type IPFixFlowTemplate struct {
	ID     uint16
	Count  uint16
	Fields []IPFixFlowTemplateField
}

type IPFixFlowTemplateField struct {
	Type   uint16
	Length uint16
}

func main() {
	rand.Seed(time.Now().UnixNano())
	command := newCommand()
	err := command.Execute()
	if err != nil {
		os.Exit(1)
	}
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
	fmt.Printf("arg1=%s, arg2=%s\n", config.arg1, config.arg2)

	conn, err := net.Dial("udp", "10.146.0.6:2100")
	if err != nil {
		return err
	}
	defer conn.Close()

	if _, err = conn.Write([]byte("Ping")); err != nil {
		return err
	}

	return nil
}
