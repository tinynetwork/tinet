package shell

import (
	"reflect"
	"testing"
)

func TestNodeConfig_ExecConf(t *testing.T) {
	type fields struct {
		Name string
		Cmds []Cmd
	}
	type args struct {
		nodeType string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   []string
	}{
		{
			name: "docker exec",
			fields: fields{
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
			args: args{
				nodeType: "",
			},
			want: []string{"docker exec R1 /usr/lib/frr/frr start", "docker exec R1 ip addr add 10.0.0.1/24 dev net0", "docker exec R1 vtysh -c 'conf t' -c 'router bgp 65001' -c 'bgp router-id 10.255.0.1' -c 'neighbor 10.0.0.2 remote-as 65002'"},
		},
		{
			name: "ip netns exec",
			fields: fields{
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
			args: args{
				nodeType: "netns",
			},
			want: []string{"ip netns exec H0 ip addr add 10.0.0.1/24 dev net0", "ip netns exec H0 ip addr add 10.1.0.1/24 dev net1"},
		},
		{
			name: "not support node",
			fields: fields{
				Name: "R0",
				Cmds: []Cmd{
					Cmd{
						Cmd: "ip addr add 10.0.0.1/24 dev net0",
					},
				},
			},
			args: args{
				nodeType: "hoge",
			},
			want: []string{""},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			nodeConfig := &NodeConfig{
				Name: tt.fields.Name,
				Cmds: tt.fields.Cmds,
			}
			if got := nodeConfig.ExecConf(tt.args.nodeType); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NodeConfig.ExecConf() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNode_DeleteNode(t *testing.T) {
	type fields struct {
		Name           string
		Type           string
		NetBase        string
		Image          string
		Interfaces     []Interface
		Sysctls        []Sysctl
		Mounts         []string
		DNS            []string
		DNSSearches    []string
		HostNameIgnore bool
		EntryPoint     string
		ExtraArgs      string
	}
	tests := []struct {
		name   string
		fields fields
		want   []string
	}{
		{
			name: "delete docker",
			fields: fields{
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
			want: []string{"docker rm -f R1 || true", "rm -rf /var/run/netns/R1"},
		},
		{
			name: "delete netns",
			fields: fields{
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
			want: []string{"ip netns del H1", "rm -rf /var/run/netns/H1"},
		},
		{
			name: "delete not support nodetype",
			fields: fields{
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
			want: []string{""},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			node := &Node{
				Name:           tt.fields.Name,
				Type:           tt.fields.Type,
				NetBase:        tt.fields.NetBase,
				Image:          tt.fields.Image,
				Interfaces:     tt.fields.Interfaces,
				Sysctls:        tt.fields.Sysctls,
				Mounts:         tt.fields.Mounts,
				DNS:            tt.fields.DNS,
				DNSSearches:    tt.fields.DNSSearches,
				HostNameIgnore: tt.fields.HostNameIgnore,
				EntryPoint:     tt.fields.EntryPoint,
				ExtraArgs:      tt.fields.ExtraArgs,
			}
			if got := node.DeleteNode(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Node.DeleteNode() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSwitch_DeleteSwitch(t *testing.T) {
	type fields struct {
		Name       string
		Interfaces []Interface
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{
			name: "delete bridge",
			fields: fields{
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
			want: "ovs-vsctl del-br SW",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			br := &Switch{
				Name:       tt.fields.Name,
				Interfaces: tt.fields.Interfaces,
			}
			if got := br.DeleteSwitch(); got != tt.want {
				t.Errorf("Switch.DeleteSwitch() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTn_Exec(t *testing.T) {
	type fields struct {
		PreCmd      []PreCmd
		PreInit     []PreInit
		PreConf     []PreConf
		PostInit    []PostInit
		PostFini    []PostFini
		Nodes       []Node
		Switches    []Switch
		NodeConfigs []NodeConfig
		Test        []Test
	}
	type args struct {
		nodeName string
		Cmds     []string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tnconfig := &Tn{
				PreCmd:      tt.fields.PreCmd,
				PreInit:     tt.fields.PreInit,
				PreConf:     tt.fields.PreConf,
				PostInit:    tt.fields.PostInit,
				PostFini:    tt.fields.PostFini,
				Nodes:       tt.fields.Nodes,
				Switches:    tt.fields.Switches,
				NodeConfigs: tt.fields.NodeConfigs,
				Test:        tt.fields.Test,
			}
			if got := tnconfig.Exec(tt.args.nodeName, tt.args.Cmds); got != tt.want {
				t.Errorf("Tn.Exec() = %v, want %v", got, tt.want)
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

func TestNode_CreateNode(t *testing.T) {
	type fields struct {
		Name           string
		Type           string
		NetBase        string
		VolumeBase     string
		Image          string
		Interfaces     []Interface
		Sysctls        []Sysctl
		Mounts         []string
		DNS            []string
		DNSSearches    []string
		HostNameIgnore bool
		EntryPoint     string
		ExtraArgs      string
	}
	tests := []struct {
		name   string
		fields fields
		want   []string
	}{
		{
			name: "create docker node net None",
			fields: fields{
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
			want: []string{"docker run -td --net none --name R1 --rm --privileged --hostname R1 -v /tmp/tinet:/tinet slankdev/frr"},
		},
		{
			name: "create docker node net None with sysctls",
			fields: fields{
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
			want: []string{"docker run -td --net none --name R1 --rm --privileged --hostname R1 --sysctl net.ipv4.ip_forward=1 --sysctl net.ipv6.conf.all.forwarding=1 -v /tmp/tinet:/tinet slankdev/frr"},
		},
		{
			name: "create docker node net bridge",
			fields: fields{
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
			want: []string{"docker run -td --net bridge --name R1 --rm --privileged --hostname R1 -v /tmp/tinet:/tinet slankdev/frr"},
		},
		{
			name: "create netns node",
			fields: fields{
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
			want: []string{"ip netns add C1", "ip netns exec C1 ip link set lo up"},
		},
		{
			name: "create netns node with sysctls",
			fields: fields{
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
			want: []string{"ip netns add C1", "ip netns exec C1 sysctl -w net.ipv4.ip_forward=1", "ip netns exec C1 sysctl -w net.ipv6.conf.all.forwarding=1", "ip netns exec C1 ip link set lo up"},
		},
		{
			name: "create not support node",
			fields: fields{
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
			want: []string{"unknown nodetype hoge"},
		},
		{
			name: "create node with mounts",
			fields: fields{
				Name:  "T1",
				Image: "slankdev/frr",
				Interfaces: []Interface{
					Interface{
						Name: "net0",
						Type: "direct",
						Args: "T2#net0",
					},
				},
				Mounts: []string{
					"`pwd`:/mnt/test",
					"/usr/share/vim:/mnt/vim",
				},
			},
			want: []string{"docker run -td --net none --name T1 --rm --privileged --hostname T1 -v /tmp/tinet:/tinet -v `pwd`:/mnt/test -v /usr/share/vim:/mnt/vim slankdev/frr"},
		},
		{
			name: "create node with DNS",
			fields: fields{
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
				DNS: []string{
					"8.8.8.8",
					"1.1.1.1",
				},
				DNSSearches: []string{
					"local",
					"corp",
				},
			},
			want: []string{"docker run -td --net bridge --name R1 --rm --privileged --hostname R1 -v /tmp/tinet:/tinet --dns=8.8.8.8 --dns=1.1.1.1 --dns-search=local --dns-search=corp slankdev/frr"},
		},
		{
			name: "create node with specify tinet volume",
			fields: fields{
				Name:       "T1",
				Image:      "slankdev/frr",
				VolumeBase: "/tmp/ak1ra24",
				Interfaces: []Interface{
					Interface{
						Name: "net0",
						Type: "direct",
						Args: "T2#net0",
					},
				},
			},
			want: []string{"docker run -td --net none --name T1 --rm --privileged --hostname T1 -v /tmp/ak1ra24:/tinet slankdev/frr"},
		},
		{
			name: "create node with hostname_ignore",
			fields: fields{
				Name:       "T1",
				Image:      "slankdev/frr",
				VolumeBase: "/tmp/ak1ra24",
				Interfaces: []Interface{
					Interface{
						Name: "net0",
						Type: "direct",
						Args: "T2#net0",
					},
				},
				HostNameIgnore: true,
			},
			want: []string{"docker run -td --net none --name T1 --rm --privileged -v /tmp/ak1ra24:/tinet slankdev/frr"},
		},
		{
			name: "create node with entrypoint",
			fields: fields{
				Name:       "T1",
				Image:      "slankdev/frr",
				VolumeBase: "/tmp/ak1ra24",
				Interfaces: []Interface{
					Interface{
						Name: "net0",
						Type: "direct",
						Args: "T2#net0",
					},
				},
				EntryPoint: "bash",
			},
			want: []string{"docker run -td --net none --name T1 --rm --privileged --hostname T1 --entrypoint bash -v /tmp/ak1ra24:/tinet slankdev/frr"},
		},
		{
			name: "create node with extra args",
			fields: fields{
				Name:  "T1",
				Image: "nginx",
				Interfaces: []Interface{
					Interface{
						Name: "net0",
						Type: "direct",
						Args: "T2#net0",
					},
				},
				ExtraArgs: "-p 8080:80",
			},
			want: []string{"docker run -td --net none --name T1 --rm --privileged --hostname T1 -v /tmp/tinet:/tinet -p 8080:80 nginx"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			node := &Node{
				Name:           tt.fields.Name,
				Type:           tt.fields.Type,
				NetBase:        tt.fields.NetBase,
				VolumeBase:     tt.fields.VolumeBase,
				Image:          tt.fields.Image,
				Interfaces:     tt.fields.Interfaces,
				Sysctls:        tt.fields.Sysctls,
				Mounts:         tt.fields.Mounts,
				DNS:            tt.fields.DNS,
				DNSSearches:    tt.fields.DNSSearches,
				HostNameIgnore: tt.fields.HostNameIgnore,
				EntryPoint:     tt.fields.EntryPoint,
				ExtraArgs:      tt.fields.ExtraArgs,
			}
			if got := node.CreateNode(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Node.CreateNode() = %v, want %v", got, tt.want)
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

func TestSwitch_CreateSwitch(t *testing.T) {
	type fields struct {
		Name       string
		Interfaces []Interface
	}
	tests := []struct {
		name   string
		fields fields
		want   []string
	}{
		{
			name: "create switch shell",
			fields: fields{
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
			want: []string{"ovs-vsctl add-br SW", "ip link set SW up"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			bridge := &Switch{
				Name:       tt.fields.Name,
				Interfaces: tt.fields.Interfaces,
			}
			if got := bridge.CreateSwitch(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Switch.CreateSwitch() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestInterface_N2nLink(t *testing.T) {
	type fields struct {
		Name string
		Type string
		Args string
		Addr string
	}
	type args struct {
		nodeName string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   []string
	}{
		{
			name: "peer between containers",
			fields: fields{
				Name: "net0",
				Type: "direct",
				Args: "R2#net0",
			},
			args: args{
				nodeName: "R1",
			},
			want: []string{"ip link add net0 netns R1 type veth peer name net0 netns R2", "ip netns exec R1 ip link set net0 up", "ip netns exec R2 ip link set net0 up"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			inf := &Interface{
				Name: tt.fields.Name,
				Type: tt.fields.Type,
				Args: tt.fields.Args,
				Addr: tt.fields.Addr,
			}
			if got := inf.N2nLink(tt.args.nodeName); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Interface.N2nLink() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestInterface_S2nLink(t *testing.T) {
	type fields struct {
		Name string
		Type string
		Args string
		Addr string
	}
	type args struct {
		nodeName string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   []string
	}{
		{
			name: "link connect between switch and container",
			fields: fields{
				Name: "net0",
				Type: "bridge",
				Args: "SW",
			},
			args: args{
				nodeName: "R1",
			},
			want: []string{"ip link add net0 netns R1 type veth peer name SW-R1", "ip netns exec R1 ip link set net0 up", "ip link set SW-R1 up", "ovs-vsctl add-port SW SW-R1"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			inf := &Interface{
				Name: tt.fields.Name,
				Type: tt.fields.Type,
				Args: tt.fields.Args,
				Addr: tt.fields.Addr,
			}
			if got := inf.S2nLink(tt.args.nodeName); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Interface.S2nLink() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestInterface_V2cLink(t *testing.T) {
	type fields struct {
		Name string
		Type string
		Args string
		Addr string
	}
	type args struct {
		nodeName string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   []string
	}{
		{
			name: "link connect between veth and container",
			fields: fields{
				Name: "port-0-6-0",
				Type: "veth",
				Args: "pp6",
			},
			args: args{
				nodeName: "R1",
			},
			want: []string{"ip link add port-0-6-0 type veth peer name pp6", "ip link set dev port-0-6-0 netns R1", "ip netns exec R1 ip link set port-0-6-0 up", "ip netns del R1", "ip link set pp6 up"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			inf := &Interface{
				Name: tt.fields.Name,
				Type: tt.fields.Type,
				Args: tt.fields.Args,
				Addr: tt.fields.Addr,
			}
			if got := inf.V2cLink(tt.args.nodeName); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Interface.V2cLink() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestInterface_P2cLink(t *testing.T) {
	type fields struct {
		Name string
		Type string
		Args string
		Addr string
	}
	type args struct {
		nodeName string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   []string
	}{
		{
			name: "Link connect between phys ethernet and Container",
			fields: fields{
				Name: "ens20f1",
				Type: "phys",
				Args: "",
			},
			args: args{
				nodeName: "R1",
			},
			want: []string{"ip link set dev ens20f1 netns R1", "ip netns exec R1 ip link set ens20f1 up", "ip netns del R1"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			inf := &Interface{
				Name: tt.fields.Name,
				Type: tt.fields.Type,
				Args: tt.fields.Args,
				Addr: tt.fields.Addr,
			}
			if got := inf.P2cLink(tt.args.nodeName); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Interface.P2cLink() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNode_Mount_docker_netns(t *testing.T) {
	type fields struct {
		Name           string
		Type           string
		NetBase        string
		Image          string
		Interfaces     []Interface
		Sysctls        []Sysctl
		Mounts         []string
		HostNameIgnore bool
		EntryPoint     string
		ExtraArgs      string
	}
	tests := []struct {
		name   string
		fields fields
		want   []string
	}{
		{
			name: "docker mount netns",
			fields: fields{
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
			want: []string{"mkdir -p /var/run/netns", "PID=`docker inspect R1 --format '{{.State.Pid}}'`", "ln -s /proc/$PID/ns/net /var/run/netns/R1"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			node := &Node{
				Name:           tt.fields.Name,
				Type:           tt.fields.Type,
				NetBase:        tt.fields.NetBase,
				Image:          tt.fields.Image,
				Interfaces:     tt.fields.Interfaces,
				Sysctls:        tt.fields.Sysctls,
				Mounts:         tt.fields.Mounts,
				HostNameIgnore: tt.fields.HostNameIgnore,
				EntryPoint:     tt.fields.EntryPoint,
				ExtraArgs:      tt.fields.ExtraArgs,
			}
			if got := node.Mount_docker_netns(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Node.Mount_docker_netns() = %v, want %v", got, tt.want)
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

func TestNode_DelNsCmd(t *testing.T) {
	type fields struct {
		Name           string
		Type           string
		NetBase        string
		Image          string
		Interfaces     []Interface
		Sysctls        []Sysctl
		Mounts         []string
		DNS            []string
		DNSSearches    []string
		HostNameIgnore bool
		EntryPoint     string
		ExtraArgs      string
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{
			name: "node netns del",
			fields: fields{
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
			want: "ip netns del R1",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			node := &Node{
				Name:           tt.fields.Name,
				Type:           tt.fields.Type,
				NetBase:        tt.fields.NetBase,
				Image:          tt.fields.Image,
				Interfaces:     tt.fields.Interfaces,
				Sysctls:        tt.fields.Sysctls,
				Mounts:         tt.fields.Mounts,
				DNS:            tt.fields.DNS,
				DNSSearches:    tt.fields.DNSSearches,
				HostNameIgnore: tt.fields.HostNameIgnore,
				EntryPoint:     tt.fields.EntryPoint,
				ExtraArgs:      tt.fields.ExtraArgs,
			}
			if got := node.DelNsCmd(); got != tt.want {
				t.Errorf("Node.DelNsCmd() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNode_BuildCmd(t *testing.T) {
	type fields struct {
		Name           string
		Type           string
		NetBase        string
		VolumeBase     string
		Image          string
		BuildFile      string
		BuildContext   string
		Interfaces     []Interface
		Sysctls        []Sysctl
		Mounts         []string
		DNS            []string
		DNSSearches    []string
		HostNameIgnore bool
		EntryPoint     string
		ExtraArgs      string
		Vars           map[string]interface{}
		Templates      []Template
	}
	tests := []struct {
		name         string
		fields       fields
		wantBuildCmd string
	}{
		{
			name: "docker build",
			fields: fields{
				Name:         "R1",
				Type:         "docker",
				Image:        "ak1ra24/testimage",
				BuildFile:    "`pwd`/dockerfile/Dockerfile",
				BuildContext: "",
				Interfaces: []Interface{
					Interface{
						Name: "net0",
						Type: "direct",
						Args: "R2#net0",
					},
				},
			},
			wantBuildCmd: "docker build -t ak1ra24/testimage -f `pwd`/dockerfile/Dockerfile .\n",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			node := &Node{
				Name:           tt.fields.Name,
				Type:           tt.fields.Type,
				NetBase:        tt.fields.NetBase,
				VolumeBase:     tt.fields.VolumeBase,
				Image:          tt.fields.Image,
				BuildFile:      tt.fields.BuildFile,
				BuildContext:   tt.fields.BuildContext,
				Interfaces:     tt.fields.Interfaces,
				Sysctls:        tt.fields.Sysctls,
				Mounts:         tt.fields.Mounts,
				DNS:            tt.fields.DNS,
				DNSSearches:    tt.fields.DNSSearches,
				HostNameIgnore: tt.fields.HostNameIgnore,
				EntryPoint:     tt.fields.EntryPoint,
				ExtraArgs:      tt.fields.ExtraArgs,
				Vars:           tt.fields.Vars,
				Templates:      tt.fields.Templates,
			}
			if gotBuildCmd := node.BuildCmd(); gotBuildCmd != tt.wantBuildCmd {
				t.Errorf("Node.BuildCmd() = %v, want %v", gotBuildCmd, tt.wantBuildCmd)
			}
		})
	}
}

func TestTest_TnTestCmdExec(t *testing.T) {
	type fields struct {
		Name string
		Cmds []Cmd
	}
	tests := []struct {
		name           string
		fields         fields
		wantTnTestCmds []string
	}{
		{
			name: "test name not set",
			fields: fields{
				Cmds: []Cmd{
					Cmd{
						Cmd: "echo slankdev",
					},
				},
			},
			wantTnTestCmds: []string{"echo slankdev"},
		},
		{
			name: "test name set",
			fields: fields{
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
			wantTnTestCmds: []string{"docker exec R1 echo hello", "docker exec R2 echo world"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tr := &Test{
				Name: tt.fields.Name,
				Cmds: tt.fields.Cmds,
			}
			if gotTnTestCmds := tr.TnTestCmdExec(); !reflect.DeepEqual(gotTnTestCmds, tt.wantTnTestCmds) {
				t.Errorf("Test.TnTestCmdExec() = %v, want %v", gotTnTestCmds, tt.wantTnTestCmds)
			}
		})
	}
}
