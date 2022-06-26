package main

import (
	"bytes"
	"encoding/binary"
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
	Header   IPFixMessageHeader
	FlowSets []IPFixFlowSet
}

type IPFixMessageHeader struct {
	VersionNumber  uint16
	Count          uint16
	SysupTime      uint32
	UnixSecs       uint32
	SequenceNumber uint32
	SourceID       uint32
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
	if err := newCommand().Execute(); err != nil {
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

	msg := IPFixMessage{
		Header: IPFixMessageHeader{
			VersionNumber:  9,
			Count:          1,
			SysupTime:      0x00002250,
			UnixSecs:       0x62b7f72d,
			SequenceNumber: 1,
			SourceID:       0,
		},
		FlowSets: []IPFixFlowSet{
			{
				FlowSetID: 0,
				Length:    64,
				Template: IPFixFlowTemplate{
					Fields: []IPFixFlowTemplateField{
						{
							FieldType:   153,
							FieldLength: 8,
						},
						{
							FieldType:   152,
							FieldLength: 8,
						},
						{
							FieldType:   1,
							FieldLength: 8,
						},
						{
							FieldType:   2,
							FieldLength: 8,
						},
						{
							FieldType:   60,
							FieldLength: 1,
						},
						{
							FieldType:   10,
							FieldLength: 4,
						},
						{
							FieldType:   14,
							FieldLength: 4,
						},
						{
							FieldType:   61,
							FieldLength: 1,
						},
						{
							FieldType:   8,
							FieldLength: 4,
						},
						{
							FieldType:   12,
							FieldLength: 4,
						},
						{
							FieldType:   7,
							FieldLength: 2,
						},
						{
							FieldType:   11,
							FieldLength: 2,
						},
						{
							FieldType:   6,
							FieldLength: 1,
						},
						{
							FieldType:   4,
							FieldLength: 1,
						},
					},
				},
			},
		},
	}

	buf := &bytes.Buffer{}
	if err := msg.ToBuffer(buf); err != nil {
		return err
	}
	if err := udptransmit("10.146.0.6:2100", buf); err != nil {
		return err
	}
	return nil
}

func udptransmit(dst string, buf *bytes.Buffer) error {
	conn, err := net.Dial("udp", dst)
	if err != nil {
		return err
	}
	defer conn.Close()
	if _, err = conn.Write(buf.Bytes()); err != nil {
		return err
	}
	return nil
}

func (m *IPFixMessage) ToBuffer(buf *bytes.Buffer) error {
	if err := binary.Write(buf, binary.BigEndian, &m.Header); err != nil {
		return err
	}
	return nil
}
