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

// https://www.rfc-editor.org/rfc/rfc3954.html#section-5.1
type IPFixMessage struct {
	VersionNumber  uint16
	Count          uint16
	SysupTime      uint32
	UnixSecs       uint32
	SequenceNumber uint32
	SourceID       uint32
	FlowSets       []IPFixFlowSet
}

// https://www.rfc-editor.org/rfc/rfc3954.html#section-5.2
type IPFixFlowSet struct {
	FlowSetID uint16
	Length    uint16
	Template  IPFixFlowTemplate
}

// https://www.rfc-editor.org/rfc/rfc3954.html#section-5.2
type IPFixFlowTemplate struct {
	TemplateID uint16
	FieldCount uint16
	Fields     []IPFixFlowTemplateField
}

// https://www.rfc-editor.org/rfc/rfc3954.html#section-5.2
type IPFixFlowTemplateField struct {
	FieldType   uint16
	FieldLength uint16
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
