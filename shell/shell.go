package shell

import (
	"fmt"
	"io/ioutil"
	"log"
	"strings"

	"github.com/ak1ra24/tn/utils"
	"gopkg.in/yaml.v2"
)

// Tn tinet config
type Tn struct {
	PreInit     PreInit      `yaml:"preinit"`
	PostInit    PostInit     `yaml:"postinit"`
	PostFini    PostFini     `yaml:"postfini"`
	Nodes       []Node       `yaml:"nodes" mapstructure:"nodes"`
	Switches    []Switch     `yaml:"switches" mapstructure:"switches"`
	NodeConfigs []NodeConfig `yaml:"node_configs" mapstructure:"node_configs"`
	Test        []Test       `yaml:"test"`
}

// PreInit
type PreInit struct {
	Cmds []Cmd `yaml:"cmds" mapstructure:"cmds"`
}

// PostInit
type PostInit struct {
	Cmds []Cmd `yaml:"cmds" mapstructure:"cmds"`
}

// PostFini
type PostFini struct {
	Cmds []Cmd `yaml:"cmds" mapstructure:"cmds"`
}

// Node
type Node struct {
	Name       string      `yaml:"name"`
	Type       string      `yaml:"type"`
	NetBase    string      `yaml:"net_base"`
	Image      string      `yaml:"image"`
	Interfaces []Interface `yaml:"interfaces" mapstructure:"interfaces"`
}

// Interface
type Interface struct {
	Name string `yaml:"name"`
	Type string `yaml:"type"`
	Args string `yaml:"args"`
}

// Switch
type Switch struct {
	Name       string      `yaml:"name"`
	Interfaces []Interface `yaml:"interfaces" mapstructure:"interfaces"`
}

// NodeConfig
type NodeConfig struct {
	Name string `yaml:"name"`
	Cmds []Cmd  `yaml:"cmds" mapstructure:"cmds"`
}

// Cmd
type Cmd struct {
	Cmd string `yaml:"cmd"`
}

// Test
type Test struct {
	Name string
	Cmds []Cmd `yaml:"cmds" mapstructure:"cmds"`
}

// BuildCmd
func BuildCmd(nodes []Node) string {
	// log.Println("sorry not implement...")
	return "sorry not implement..."
}

// ExecConf Execute NodeConfig command
func ExecConf(nodeType string, nodeConfig NodeConfig) []string {
	var execConfCmds []string
	for _, nodeConfigCmd := range nodeConfig.Cmds {
		var execConfCmd string
		if nodeType == "docker" {
			execConfCmd = fmt.Sprintf("docker exec %s %s", nodeConfig.Name, nodeConfigCmd.Cmd)
		} else if nodeType == "netns" {
			execConfCmd = fmt.Sprintf("ip netns exec %s %s", nodeConfig.Name, nodeConfigCmd.Cmd)
		} else if nodeType == "" {
			execConfCmd = fmt.Sprintf("docker exec %s %s", nodeConfig.Name, nodeConfigCmd.Cmd)
		} else {
			// err := fmt.Errorf("not supported node type...")
			// log.Fatal(err)
			return []string{""}
		}
		// fmt.Println(execConfigCmd)
		execConfCmds = append(execConfCmds, execConfCmd)
	}

	return execConfCmds
}

// DeleteNode Delete docker and netns
func DeleteNode(node Node) string {
	var deleteCmd string
	if node.Type == "docker" {
		deleteCmd = fmt.Sprintf("docker stop %s", node.Name)
	} else if node.Type == "netns" {
		deleteCmd = fmt.Sprintf("ip netns del %s", node.Name)
	} else if node.Type == "" {
		deleteCmd = fmt.Sprintf("docker stop %s", node.Name)
	} else {
		// err := fmt.Errorf("not supported node type...")
		// log.Fatal(err)
		return ""
	}

	return deleteCmd
}

// DeleteSwitch Delete bridge
func DeleteSwitch(br Switch) string {
	deleteBrCmd := fmt.Sprintf("ip link delete %s", br.Name)
	return deleteBrCmd
}

// Exec Select Node exec command
func (tnconfig *Tn) Exec(nodeName string, Cmds []string) string {
	var execCommand string
	var selectedNode *Node

	for i, node := range tnconfig.Nodes {
		if node.Name == nodeName {
			selectedNode = &tnconfig.Nodes[i]
		}
	}

	if selectedNode == nil {
		err := fmt.Errorf("no such node...\n")
		log.Fatal(err)
	}

	if selectedNode.Type == "docker" {
		execCommand = fmt.Sprintf("docker exec %s", nodeName)
	} else if selectedNode.Type == "netns" {
		execCommand = fmt.Sprintf("ip netns exec %s", nodeName)
	} else if selectedNode.Type == "" {
		execCommand = fmt.Sprintf("docker exec %s", nodeName)
	} else {
		err := fmt.Errorf("no such node type...\n")
		log.Fatal(err)
	}

	var cmdStr string
	for _, cmd := range Cmds {
		cmdStr += fmt.Sprintf(" %s", cmd)
		execCommand += cmdStr
	}

	return execCommand
}

// GenerateFile Generate tinet template config file
func (tnconfig *Tn) GenerateFile(cfgFile string) error {
	preinit := PreInit{
		Cmds: []Cmd{
			Cmd{
				Cmd: "",
			},
		},
	}
	postinit := PostInit{
		Cmds: []Cmd{
			Cmd{
				Cmd: "",
			},
		},
	}

	postfini := PostFini{
		Cmds: []Cmd{
			Cmd{
				Cmd: "",
			},
		},
	}

	nodes := Node{
		Name:    "",
		Image:   "",
		NetBase: "",
		Interfaces: []Interface{
			Interface{
				Name: "",
				Type: "",
				Args: "",
			},
		},
	}

	switches := Switch{
		Name: "",
		Interfaces: []Interface{
			Interface{
				Name: "",
				Type: "",
				Args: "",
			},
		},
	}

	nodeConfig := NodeConfig{
		Name: "",
		Cmds: []Cmd{
			Cmd{
				Cmd: "",
			},
		},
	}

	test := Test{
		Name: "",
		Cmds: []Cmd{
			Cmd{
				Cmd: "",
			},
		},
	}

	tnconfig = &Tn{
		PreInit:     preinit,
		PostInit:    postinit,
		PostFini:    postfini,
		Nodes:       []Node{nodes},
		Switches:    []Switch{switches},
		NodeConfigs: []NodeConfig{nodeConfig},
		Test:        []Test{test},
	}

	data, err := yaml.Marshal(tnconfig)
	if err != nil {
		return err
	}

	err = ioutil.WriteFile(cfgFile, data, 0644)
	if err != nil {
		return err
	}

	return nil
}

// DockerPs Show docker ps
func DockerPs(all bool) string {
	dockerPsCmd := "docker ps --format 'table {{.ID}}\\t{{.Names}}\\t{{.Status}}\\t{{.Networks}}'"
	if all {
		dockerPsCmd += " -a"
	}

	return dockerPsCmd
}

// NetnsPs Show ip netns list
func NetnsPs() string {
	netnsListCmd := "ip netns list"

	return netnsListCmd
}

// Pull Pull Docker Image
func Pull(nodes []Node) []string {
	var images []string
	for _, node := range nodes {
		images = append(images, node.Image)
	}
	images = utils.RemoveDuplicatesString(images)

	var pullCmds []string
	for _, image := range images {
		pullCmd := fmt.Sprintf("docker pull %s", image)
		pullCmds = append(pullCmds, pullCmd)
	}

	return pullCmds
}

// TnTestCmdExec Execute test cmds
func TnTestCmdExec(tests []Test) []string {
	var startMessage, doneMessage string
	var tnTestCmds []string
	for _, test := range tests {
		if test.Name != "" {
			startMessage = fmt.Sprintf("echo %s Test Start", test.Name)
			doneMessage = fmt.Sprintf("echo %s Test Done", test.Name)
		} else {
			startMessage = "echo Test Start"
			doneMessage = "echo Test Done"
		}

		tnTestCmds = append(tnTestCmds, startMessage)

		for _, testCmd := range test.Cmds {
			tnTestCmds = append(tnTestCmds, testCmd.Cmd)
		}

		tnTestCmds = append(tnTestCmds, doneMessage)
	}

	return tnTestCmds
}

// ExecCmd Execute cmds
func ExecCmd(cmds []Cmd) []string {
	var execCmds []string
	for _, cmd := range cmds {
		execCmds = append(execCmds, cmd.Cmd)
	}

	return execCmds
}

// CreateNode Create nodes set in config
func CreateNode(node Node) []string {
	var createNodeCmds []string

	var createNodeCmd string
	if node.NetBase == "" {
		node.NetBase = "none"
	}
	if node.Type == "docker" {
		createNodeCmd = fmt.Sprintf("docker run -td --hostname %s --net %s --name %s --rm --privileged %s", node.Name, node.NetBase, node.Name, node.Image)

	} else if node.Type == "netns" {
		createNodeCmd = fmt.Sprintf("ip netns add %s", node.Name)
	} else if node.Type == "" {
		createNodeCmd = fmt.Sprintf("docker run -td --hostname %s --net %s --name %s --rm --privileged %s", node.Name, node.NetBase, node.Name, node.Image)

	} else {
		// err := fmt.Errorf("unknown nodetype %s", node.Type)
		// log.Fatal(err)
		createNodeCmd = fmt.Sprintf("unknown nodetype %s", node.Type)
	}

	createNodeCmds = append(createNodeCmds, createNodeCmd)

	if node.Type == "netns" {
		infloUpCmd := fmt.Sprintf("ip netns exec %s ip link set lo up", node.Name)
		createNodeCmds = append(createNodeCmds, infloUpCmd)
	}

	return createNodeCmds
}

// HostLinkUp Link up link of host
func HostLinkUp(linkName string) string {
	linkUpCmd := fmt.Sprintf("ip link set %s up", linkName)
	return linkUpCmd
}

// NetnsLinkUp Link up link of netns
func NetnsLinkUp(netnsName string, linkName string) string {
	netnsLinkUpCmd := fmt.Sprintf("ip netns exec %s ip link set %s up", netnsName, linkName)
	return netnsLinkUpCmd
}

// CreateSwitch Create bridge set in config
func CreateSwitch(bridge Switch) []string {
	var createSwitchCmds []string

	addSwitchCmd := fmt.Sprintf("ip link add %s type bridge", bridge.Name)
	createSwitchCmds = append(createSwitchCmds, addSwitchCmd)

	bridgeUpCmd := HostLinkUp(bridge.Name)
	createSwitchCmds = append(createSwitchCmds, bridgeUpCmd)

	return createSwitchCmds
}

// N2nLink Connect links between nodes
func N2nLink(nodeName string, inf Interface) []string {
	var n2nLinkCmds []string

	nodeinf := inf.Name
	peerinfo := strings.Split(inf.Args, "#")
	peerNode := peerinfo[0]
	peerinf := peerinfo[1]
	n2nlinkCmd := fmt.Sprintf("ip link add %s netns %s type veth peer name %s netns %s", nodeinf, nodeName, peerinf, peerNode)
	n2nLinkCmds = append(n2nLinkCmds, n2nlinkCmd)
	n2nLinkCmds = append(n2nLinkCmds, NetnsLinkUp(nodeName, nodeinf))
	n2nLinkCmds = append(n2nLinkCmds, NetnsLinkUp(peerNode, peerinf))

	return n2nLinkCmds
}

// S2nLink Connect links between nodes and switches
func S2nLink(nodeName string, inf Interface) string {
	nodeinf := inf.Name
	peerBr := inf.Args
	s2nLinkCmd := fmt.Sprintf("ip link add %s netns %s type veth peer name %s-%s", nodeinf, nodeName, peerBr, nodeName)
	return s2nLinkCmd
}

// V2cLink Connect links between veth and container
func V2cLink(nodeName string, inf Interface) []string {
	var v2cLinkCmds []string
	nodeinf := inf.Name
	peerName := inf.Args
	v2cLinkCmd := fmt.Sprintf("ip link add %s type veth peer name %s", nodeinf, peerName)
	v2cLinkCmds = append(v2cLinkCmds, v2cLinkCmd)

	v2cLinkCmds = append(v2cLinkCmds, P2cLink(nodeName, inf)...)
	v2cLinkCmds = append(v2cLinkCmds, HostLinkUp(peerName))

	return v2cLinkCmds
}

// P2cLink Connect links between phys-eth and container
func P2cLink(nodeName string, inf Interface) []string {
	var p2cLinkCmds []string

	physInf := inf.Name
	setNsCmd := fmt.Sprintf("ip link set dev %s netns %s", physInf, nodeName)
	p2cLinkCmds = append(p2cLinkCmds, setNsCmd)
	execNsCmd := fmt.Sprintf("ip netns exec %s ip link set %s up", nodeName, physInf)
	p2cLinkCmds = append(p2cLinkCmds, execNsCmd)
	delNsCmd := fmt.Sprintf("ip netns del %s", nodeName)
	p2cLinkCmds = append(p2cLinkCmds, delNsCmd)

	return p2cLinkCmds
}

// Mount_docker_netns Mount docker netns to ip netns
func Mount_docker_netns(node Node) []string {
	var mountDockerNetnsCmds []string

	netnsDir := "/var/run/netns"
	mkdirCmd := fmt.Sprintf("mkdir -p %s", netnsDir)
	mountDockerNetnsCmds = append(mountDockerNetnsCmds, mkdirCmd)
	dockerPid := GetContainerPid(node.Name)
	mountDockerNetnsCmds = append(mountDockerNetnsCmds, dockerPid)
	mountDockerNetnsCmd := fmt.Sprintf("ln -s /proc/$PID/ns/net /var/run/netns/%s", node.Name)
	mountDockerNetnsCmds = append(mountDockerNetnsCmds, mountDockerNetnsCmd)

	return mountDockerNetnsCmds
}

// GetContainerPid func is Output get Docker PID Command
func GetContainerPid(nodename string) string {
	getpidcmd := fmt.Sprintf("PID=`docker inspect %s --format '{{.State.Pid}}'`", nodename)

	return getpidcmd
}
