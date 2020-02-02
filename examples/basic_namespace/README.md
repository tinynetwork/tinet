
# Namespace Isolation

![](./topo.png)

```
cd path/to/here
tn -f spec.blue.yaml upconf | sudo sh
tn -f spec.green.yaml upconf | sudo sh
docker ps
CONTAINER ID        IMAGE                   COMMAND             CREATED             STATUS              PORTS               NAMES
6d5f9e4d6c51        slankdev/ubuntu:18.04   "/bin/bash"         2 seconds ago       Up 1 second                             green_R2
92210c5eea85        slankdev/ubuntu:18.04   "/bin/bash"         3 seconds ago       Up 2 seconds                            green_R1
48a50568c9c1        slankdev/ubuntu:18.04   "/bin/bash"         7 seconds ago       Up 6 seconds                            blue_R2
86c2c4c9fc52        slankdev/ubuntu:18.04   "/bin/bash"         7 seconds ago       Up 7 seconds                            blue_R1
```

