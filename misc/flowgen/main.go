package main

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"math/rand"
	"net"
	"os"
	"time"

	"github.com/cloudflare/goflow/decoders/netflow"
	"github.com/spf13/cobra"
)

var config struct {
	arg1 string
	arg2 string
}

type IPFixMessage struct {
	Header   IPFixHeader
	FlowSets []IPFixFlowSet
}

type IPFixHeader struct {
	VersionNumber  uint16
	SysupTime      uint32
	UnixSecs       uint32
	SequenceNumber uint32
	SourceID       uint32
}

type IPFixFlowSet struct {
	FlowSetID uint16
	Template  IPFixFlowTemplate
	Flow      []IPFixFlow
}

type IPFixFlow struct {
	FlowEndMilliseconds      uint64
	FlowStartMilliseconds    uint64
	OctetDeltaCount          uint64
	PacketDeltaCount         uint64
	IpVersion                uint8
	IngressInterface         uint32
	EgressInterface          uint32
	FlowDirection            uint8
	SourceIPv4Address        uint32
	DestinationIPv4Address   uint32
	SourceTransportPort      uint16
	DestinationTransportPort uint16
	TcpControlBits           uint8
	ProtocolIdentifier       uint8
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
		Header: IPFixHeader{
			VersionNumber:  9,
			SysupTime:      0x00002250,
			UnixSecs:       0x62b7f72d,
			SequenceNumber: 1,
			SourceID:       0,
		},
		FlowSets: []IPFixFlowSet{
			{
				FlowSetID: 0,
				Template: IPFixFlowTemplate{
					TemplateID: 1024,
					FieldCount: 14,
					Fields: []IPFixFlowTemplateField{
						{
							FieldType:   netflow.IPFIX_FIELD_flowEndMilliseconds,
							FieldLength: 8,
						},
						{
							FieldType:   netflow.IPFIX_FIELD_flowStartMilliseconds,
							FieldLength: 8,
						},
						{
							FieldType:   netflow.IPFIX_FIELD_octetDeltaCount,
							FieldLength: 8,
						},
						{
							FieldType:   netflow.IPFIX_FIELD_packetDeltaCount,
							FieldLength: 8,
						},
						{
							FieldType:   netflow.IPFIX_FIELD_ipVersion,
							FieldLength: 1,
						},
						{
							FieldType:   netflow.IPFIX_FIELD_ingressInterface,
							FieldLength: 4,
						},
						{
							FieldType:   netflow.IPFIX_FIELD_egressInterface,
							FieldLength: 4,
						},
						{
							FieldType:   netflow.IPFIX_FIELD_flowDirection,
							FieldLength: 1,
						},
						{
							FieldType:   netflow.IPFIX_FIELD_sourceIPv4Address,
							FieldLength: 4,
						},
						{
							FieldType:   netflow.IPFIX_FIELD_destinationIPv4Address,
							FieldLength: 4,
						},
						{
							FieldType:   netflow.IPFIX_FIELD_sourceTransportPort,
							FieldLength: 2,
						},
						{
							FieldType:   netflow.IPFIX_FIELD_destinationTransportPort,
							FieldLength: 2,
						},
						{
							FieldType:   netflow.IPFIX_FIELD_tcpControlBits,
							FieldLength: 1,
						},
						{
							FieldType:   netflow.IPFIX_FIELD_protocolIdentifier,
							FieldLength: 1,
						},
					},
				},
			},
			{
				FlowSetID: 1024,
				Flow: []IPFixFlow{
					{
						FlowEndMilliseconds:      0x000001819e9d896b,
						FlowStartMilliseconds:    0x0000000000002118,
						OctetDeltaCount:          8482,
						PacketDeltaCount:         80,
						IpVersion:                4,
						IngressInterface:         0,
						EgressInterface:          0,
						FlowDirection:            0,
						SourceIPv4Address:        0x0a000005,
						DestinationIPv4Address:   0xcb000001,
						SourceTransportPort:      22,
						DestinationTransportPort: 63153,
						TcpControlBits:           0x18,
						ProtocolIdentifier:       6,
					},
					{
						FlowEndMilliseconds:      0x000001819e9d6565,
						FlowStartMilliseconds:    0x000001819e9d896b,
						OctetDeltaCount:          6732,
						PacketDeltaCount:         94,
						IpVersion:                4,
						IngressInterface:         0,
						EgressInterface:          0,
						FlowDirection:            0,
						SourceIPv4Address:        0xcb000001,
						DestinationIPv4Address:   0x0a000005,
						SourceTransportPort:      63153,
						DestinationTransportPort: 22,
						TcpControlBits:           0x18,
						ProtocolIdentifier:       6,
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
	// https://www.rfc-editor.org/rfc/rfc3954.html#section-5.1
	if err := binary.Write(buf, binary.BigEndian, &struct {
		VersionNumber  uint16
		Count          uint16
		SysupTime      uint32
		UnixSecs       uint32
		SequenceNumber uint32
		SourceID       uint32
	}{
		VersionNumber:  m.Header.VersionNumber,
		Count:          uint16(len(m.FlowSets)),
		SysupTime:      m.Header.SysupTime,
		UnixSecs:       m.Header.UnixSecs,
		SequenceNumber: m.Header.SequenceNumber,
		SourceID:       m.Header.SourceID,
	}); err != nil {
		return err
	}

	for _, flowset := range m.FlowSets {
		// https://www.rfc-editor.org/rfc/rfc3954.html#section-5.2
		flowsetlen := 0
		if flowset.FlowSetID == 0 { // template
			const flowsetHdrLen = 8
			flowsetlen = len(flowset.Template.Fields)*4 + flowsetHdrLen
		} else { // data
			const flowsetHdrLen = 4
			flowsetlen = len(flowset.Flow)*56 + flowsetHdrLen
		}

		if err := binary.Write(buf, binary.BigEndian, &struct {
			FlowSetID uint16
			Length    uint16
		}{
			flowset.FlowSetID,
			uint16(flowsetlen),
		}); err != nil {
			return err
		}

		if flowset.FlowSetID == 0 {
			if err := binary.Write(buf, binary.BigEndian, &struct {
				TemplateID uint16
				FieldCount uint16
			}{
				flowset.Template.TemplateID,
				flowset.Template.FieldCount,
			}); err != nil {
				return err
			}
			for _, field := range flowset.Template.Fields {
				if err := binary.Write(buf, binary.BigEndian, &field); err != nil {
					return err
				}
			}
		} else {
			if err := flowset.ToBuffer(buf); err != nil {
				return err
			}
		}
	}
	return nil
}

func (fs *IPFixFlowSet) ToBuffer(buf *bytes.Buffer) error {
	e := binary.BigEndian
	for _, flow := range fs.Flow {
		if err := binary.Write(buf, e, &flow.FlowEndMilliseconds); err != nil {
			return err
		}
		if err := binary.Write(buf, e, &flow.FlowStartMilliseconds); err != nil {
			return err
		}
		if err := binary.Write(buf, e, &flow.OctetDeltaCount); err != nil {
			return err
		}
		if err := binary.Write(buf, e, &flow.PacketDeltaCount); err != nil {
			return err
		}
		if err := binary.Write(buf, e, &flow.IpVersion); err != nil {
			return err
		}
		if err := binary.Write(buf, e, &flow.IngressInterface); err != nil {
			return err
		}
		if err := binary.Write(buf, e, &flow.EgressInterface); err != nil {
			return err
		}
		if err := binary.Write(buf, e, &flow.FlowDirection); err != nil {
			return err
		}
		if err := binary.Write(buf, e, &flow.SourceIPv4Address); err != nil {
			return err
		}
		if err := binary.Write(buf, e, &flow.DestinationIPv4Address); err != nil {
			return err
		}
		if err := binary.Write(buf, e, &flow.SourceTransportPort); err != nil {
			return err
		}
		if err := binary.Write(buf, e, &flow.DestinationTransportPort); err != nil {
			return err
		}
		if err := binary.Write(buf, e, &flow.TcpControlBits); err != nil {
			return err
		}
		if err := binary.Write(buf, e, &flow.ProtocolIdentifier); err != nil {
			return err
		}
	}
	return nil
}
