package shell

import (
	"reflect"
	"testing"
)

func TestBuildCmd(t *testing.T) {
	type args struct {
		nodes []Node
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := BuildCmd(tt.args.nodes); got != tt.want {
				t.Errorf("BuildCmd() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestExecConf(t *testing.T) {
	type args struct {
		nodeType   string
		nodeConfig NodeConfig
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{
			name: "docker exec",
			args: args{
				// nodeType: "docker",
				nodeType: "",
				nodeConfig: NodeConfig{
					Name: "R1",
					Cmds: []Cmd{
						Cmd{
							Cmd: "/usr/lib/frr/frr start",
						},
						Cmd{
							Cmd: "ip addr add 10.0.0.1/24 dev net0",
						},
						Cmd{
							Cmd: "vtysh -c 'conf t' -c 'router bgp 65001' -c 'bgp router-id 10.255.0.1' -c 'neighbor 10.0.0.2 remote-as 65002'",
						},
					},
				},
			},
			want: []string{"docker exec R1 /usr/lib/frr/frr start", "docker exec R1 ip addr add 10.0.0.1/24 dev net0", "docker exec R1 vtysh -c 'conf t' -c 'router bgp 65001' -c 'bgp router-id 10.255.0.1' -c 'neighbor 10.0.0.2 remote-as 65002'"},
		},
		{
			name: "ip netns exec",
			args: args{
				nodeType: "netns",
				nodeConfig: NodeConfig{
					Name: "H0",
					Cmds: []Cmd{
						Cmd{
							Cmd: "ip addr add 10.0.0.1/24 dev net0",
						},
						Cmd{
							Cmd: "ip addr add 10.1.0.1/24 dev net1",
						},
					},
				},
			},
			want: []string{"ip netns exec H0 ip addr add 10.0.0.1/24 dev net0", "ip netns exec H0 ip addr add 10.1.0.1/24 dev net1"},
		},
		{
			name: "not support node",
			args: args{
				nodeType: "hoge",
				nodeConfig: NodeConfig{
					Name: "R0",
					Cmds: []Cmd{
						Cmd{
							Cmd: "ip addr add 10.0.0.1/24 dev net0",
						},
					},
				},
			},
			want: []string{""},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ExecConf(tt.args.nodeType, tt.args.nodeConfig); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ExecConf() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDeleteNode(t *testing.T) {
	type args struct {
		node Node
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{
			name: "delete docker",
			args: args{
				node: Node{
					Name:  "R1",
					Image: "slankdev/frr",
					Type:  "docker",
					Interfaces: []Interface{
						Interface{
							Name: "net0",
							Type: "direct",
							Args: "R2#net0",
						},
					},
				},
			},
			want: []string{"docker stop R1", "rm -rf /var/run/netns/R1"},
		},
		{
			name: "delete netns",
			args: args{
				node: Node{
					Name: "H1",
					Type: "netns",
					Interfaces: []Interface{
						Interface{
							Name: "net0",
							Type: "direct",
							Args: "H2#net0",
						},
					},
				},
			},
			want: []string{"ip netns del H1", "rm -rf /var/run/netns/H1"},
		},
		{
			name: "delete not support nodetype",
			args: args{
				node: Node{
					Name: "U1",
					Type: "hoge",
					Interfaces: []Interface{
						Interface{
							Name: "net0",
							Type: "direct",
							Args: "U2#net0",
						},
					},
				},
			},
			want: []string{""},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := DeleteNode(tt.args.node); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("DeleteNode() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDeleteSwitch(t *testing.T) {
	type args struct {
		br Switch
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "delete bridge",
			args: args{
				br: Switch{
					Name: "SW",
					Interfaces: []Interface{
						Interface{
							Name: "net2",
							Type: "docker",
							Args: "R1",
						},
						Interface{
							Name: "net0",
							Type: "docker",
							Args: "R2",
						},
						Interface{
							Name: "net0",
							Type: "docker",
							Args: "R3",
						},
					},
				},
			},
			want: "ip link delete SW",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := DeleteSwitch(tt.args.br); got != tt.want {
				t.Errorf("DeleteSwitch() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTn_Exec(t *testing.T) {
	type args struct {
		nodeName string
		Cmds     []string
	}
	tests := []struct {
		name     string
		tnconfig *Tn
		args     args
		want     string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.tnconfig.Exec(tt.args.nodeName, tt.args.Cmds); got != tt.want {
				t.Errorf("Tn.Exec() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDockerPs(t *testing.T) {
	type args struct {
		all bool
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "docker ps",
			args: args{
				all: false,
			},
			want: "docker ps --format 'table {{.ID}}\\t{{.Names}}\\t{{.Status}}\\t{{.Networks}}'",
		},
		{
			name: "docker ps -a",
			args: args{
				all: true,
			},
			want: "docker ps --format 'table {{.ID}}\\t{{.Names}}\\t{{.Status}}\\t{{.Networks}}' -a",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := DockerPs(tt.args.all); got != tt.want {
				t.Errorf("DockerPs() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNetnsPs(t *testing.T) {
	tests := []struct {
		name string
		want string
	}{
		{
			name: "ip netns list",
			want: "ip netns list",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NetnsPs(); got != tt.want {
				t.Errorf("NetnsPs() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPull(t *testing.T) {
	type args struct {
		nodes []Node
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{
			name: "docker pull",
			args: args{
				nodes: []Node{
					Node{
						Name:  "R1",
						Image: "slankdev/frr",
					},
					Node{
						Name:  "R2",
						Image: "slankdev/frr",
					},
				},
			},
			want: []string{"docker pull slankdev/frr"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Pull(tt.args.nodes); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Pull() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTnTestCmdExec(t *testing.T) {
	type args struct {
		tests []Test
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{
			name: "test name not set",
			args: args{
				tests: []Test{
					Test{
						Cmds: []Cmd{
							Cmd{
								Cmd: "echo slankdev",
							},
						},
					},
				},
			},
			want: []string{"echo slankdev"},
		},
		{
			name: "test name set",
			args: args{
				tests: []Test{
					Test{
						Name: "p2p",
						Cmds: []Cmd{
							Cmd{
								Cmd: "docker exec R1 echo hello",
							},
							Cmd{
								Cmd: "docker exec R2 echo world",
							},
						},
					},
				},
			},
			want: []string{"docker exec R1 echo hello", "docker exec R2 echo world"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := TnTestCmdExec(tt.args.tests); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("TnTestCmdExec() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestExecCmd(t *testing.T) {
	type args struct {
		cmds []Cmd
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{
			name: "Exec Cmd",
			args: args{
				cmds: []Cmd{
					Cmd{
						Cmd: "ovs-vsctl add-port ovs0 C0net0 tag=10",
					},
					Cmd{
						Cmd: "ovs-vsctl del-port ovs0 C0net0",
					},
				},
			},
			want: []string{"ovs-vsctl add-port ovs0 C0net0 tag=10", "ovs-vsctl del-port ovs0 C0net0"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ExecCmd(tt.args.cmds); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ExecCmd() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCreateNode(t *testing.T) {
	type args struct {
		node Node
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{
			name: "create docker node net None",
			args: args{
				node: Node{
					Name:  "R1",
					Image: "slankdev/frr",
					Interfaces: []Interface{
						Interface{
							Name: "net0",
							Type: "direct",
							Args: "C1#net0",
						},
					},
				},
			},
			want: []string{"docker run -td --hostname R1 --net none --name R1 --rm --privileged slankdev/frr"},
		},
		{
			name: "create docker node net None with sysctls",
			args: args{
				node: Node{
					Name:  "R1",
					Image: "slankdev/frr",
					Interfaces: []Interface{
						Interface{
							Name: "net0",
							Type: "direct",
							Args: "C1#net0",
						},
					},
					Sysctls: []Sysctl{
						Sysctl{
							Sysctl: "net.ipv4.ip_forward=1",
						},
						Sysctl{
							Sysctl: "net.ipv6.conf.all.forwarding=1",
						},
					},
				},
			},
			want: []string{"docker run -td --hostname R1 --net none --name R1 --rm --privileged --sysctl net.ipv4.ip_forward=1 --sysctl net.ipv6.conf.all.forwarding=1 slankdev/frr"},
		},
		{
			name: "create docker node net bridge",
			args: args{
				node: Node{
					Name:    "R1",
					Image:   "slankdev/frr",
					NetBase: "bridge",
					Interfaces: []Interface{
						Interface{
							Name: "net0",
							Type: "direct",
							Args: "C1#net0",
						},
					},
				},
			},
			want: []string{"docker run -td --hostname R1 --net bridge --name R1 --rm --privileged slankdev/frr"},
		},
		{
			name: "create netns node",
			args: args{
				node: Node{
					Name:  "C1",
					Type:  "netns",
					Image: "",
					Interfaces: []Interface{
						Interface{
							Name: "net0",
							Type: "direct",
							Args: "R1#net0",
						},
					},
				},
			},
			want: []string{"ip netns add C1", "ip netns exec C1 ip link set lo up"},
		},
		{
			name: "create netns node with sysctls",
			args: args{
				node: Node{
					Name:  "C1",
					Type:  "netns",
					Image: "",
					Interfaces: []Interface{
						Interface{
							Name: "net0",
							Type: "direct",
							Args: "R1#net0",
						},
					},
					Sysctls: []Sysctl{
						Sysctl{
							Sysctl: "net.ipv4.ip_forward=1",
						},
						Sysctl{
							Sysctl: "net.ipv6.conf.all.forwarding=1",
						},
					},
				},
			},
			want: []string{"ip netns add C1", "ip netns exec C1 sysctl -w net.ipv4.ip_forward=1", "ip netns exec C1 sysctl -w net.ipv6.conf.all.forwarding=1", "ip netns exec C1 ip link set lo up"},
		},
		{
			name: "create not support node",
			args: args{
				node: Node{
					Name:  "U1",
					Type:  "hoge",
					Image: "",
					Interfaces: []Interface{
						Interface{
							Name: "net0",
							Type: "direct",
							Args: "U2#net0",
						},
					},
				},
			},
			want: []string{"unknown nodetype hoge"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CreateNode(tt.args.node); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CreateNode() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestHostLinkUp(t *testing.T) {
	type args struct {
		linkName string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "netns link up",
			args: args{
				linkName: "net0",
			},
			want: "ip link set net0 up",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := HostLinkUp(tt.args.linkName); got != tt.want {
				t.Errorf("HostLinkUp() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNetnsLinkUp(t *testing.T) {
	type args struct {
		netnsName string
		linkName  string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "netns link up",
			args: args{
				netnsName: "R1",
				linkName:  "net0",
			},
			want: "ip netns exec R1 ip link set net0 up",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NetnsLinkUp(tt.args.netnsName, tt.args.linkName); got != tt.want {
				t.Errorf("NetnsLinkUp() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCreateSwitch(t *testing.T) {
	type args struct {
		bridge Switch
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{
			name: "create switch shell",
			args: args{
				bridge: Switch{
					Name: "SW",
					Interfaces: []Interface{
						Interface{
							Name: "net0",
							Type: "docker",
							Args: "R1",
						},
						Interface{
							Name: "net0",
							Type: "docker",
							Args: "R2",
						},
					},
				},
			},
			want: []string{"ip link add SW type bridge", "ip link set SW up"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CreateSwitch(tt.args.bridge); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CreateSwitch() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestN2nLink(t *testing.T) {
	type args struct {
		nodeName string
		inf      Interface
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{
			name: "peer between containers",
			args: args{
				nodeName: "R1",
				inf: Interface{
					Name: "net0",
					Type: "direct",
					Args: "R2#net0",
				},
			},
			want: []string{"ip link add net0 netns R1 type veth peer name net0 netns R2", "ip netns exec R1 ip link set net0 up", "ip netns exec R2 ip link set net0 up"},
		},
		{
			name: "peer between containers and set macaddress",
			args: args{
				nodeName: "R1",
				inf: Interface{
					Name: "net0",
					Type: "direct",
					Args: "R2#net0",
					Addr: "52:54:00:11:11:11",
				},
			},
			want: []string{"ip link add net0 netns R1 type veth peer name net0 netns R2", "ip netns exec R1 ip link set net0 up", "ip netns exec R2 ip link set net0 up", "ip netns exec R1 ip link set net0 address 52:54:00:11:11:11"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := N2nLink(tt.args.nodeName, tt.args.inf); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("N2nLink() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestS2nLink(t *testing.T) {
	type args struct {
		nodeName string
		inf      Interface
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{
			name: "link connect between switch and container",
			args: args{
				nodeName: "R1",
				inf: Interface{
					Name: "net0",
					Type: "bridge",
					Args: "SW",
				},
			},
			want: []string{"ip link add net0 netns R1 type veth peer name SW-R1", "ip netns exec R1 ip link set net0 up", "ip link set SW-R1 up", "ip link set dev SW-R1 master SW"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := S2nLink(tt.args.nodeName, tt.args.inf); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("S2nLink() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestV2cLink(t *testing.T) {
	type args struct {
		nodeName string
		inf      Interface
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{
			name: "link connect between veth and container",
			args: args{
				nodeName: "R1",
				inf: Interface{
					Name: "port-0-6-0",
					Type: "veth",
					Args: "pp6",
				},
			},
			want: []string{"ip link add port-0-6-0 type veth peer name pp6", "ip link set dev port-0-6-0 netns R1", "ip netns exec R1 ip link set port-0-6-0 up", "ip netns del R1", "ip link set pp6 up"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := V2cLink(tt.args.nodeName, tt.args.inf); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("V2cLink() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestP2cLink(t *testing.T) {
	type args struct {
		nodeName string
		inf      Interface
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{
			name: "Link connect between phys ethernet and Container",
			args: args{
				nodeName: "R1",
				inf: Interface{
					Name: "ens20f1",
					Type: "phys",
					Args: "",
				},
			},
			want: []string{"ip link set dev ens20f1 netns R1", "ip netns exec R1 ip link set ens20f1 up", "ip netns del R1"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := P2cLink(tt.args.nodeName, tt.args.inf); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("P2cLink() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMount_docker_netns(t *testing.T) {
	type args struct {
		node Node
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{
			name: "docker mount netns",
			args: args{
				node: Node{
					Name:  "R1",
					Image: "slankdev/frr",
					Type:  "docker",
					Interfaces: []Interface{
						Interface{
							Name: "net0",
							Type: "direct",
							Args: "R2#net0",
						},
					},
				},
			},
			want: []string{"mkdir -p /var/run/netns", "PID=`docker inspect R1 --format '{{.State.Pid}}'`", "ln -s /proc/$PID/ns/net /var/run/netns/R1"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Mount_docker_netns(tt.args.node); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Mount_docker_netns() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetContainerPid(t *testing.T) {
	type args struct {
		nodename string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "d",
			args: args{
				nodename: "R1",
			},
			want: "PID=`docker inspect R1 --format '{{.State.Pid}}'`",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetContainerPid(tt.args.nodename); got != tt.want {
				t.Errorf("GetContainerPid() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGenerateFile(t *testing.T) {
	tests := []struct {
		name    string
		want    string
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := GenerateFile()
			if (err != nil) != tt.wantErr {
				t.Errorf("GenerateFile() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("GenerateFile() = %v, want %v", got, tt.want)
			}
		})
	}
}
