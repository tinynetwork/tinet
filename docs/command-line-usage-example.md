# Command Line Usage Example

## tn build
That hasn't been implemented yet.

## tn check
```
## Check link node to node
tn check -c spec.yaml
```

## tn conf
```
tn conf -c spec.yaml

## docker and netns exec config
tn conf -c spec.yaml | sudo sh -x
```

## tn down
```
tn down -c spec.yaml

## Remove docker container and netns
tn conf -c spec.yaml | sudo sh -x
```

## tn exec
That hasn't been implemented yet.

## tn help
```
tn help
tn -h
```

## tn img
```
## Output dot
tn img -c spec.yaml

## Generate img file
tn img -c spec.yaml | dot -Tpng > spec.png
```

## tn init
```
## Output tinet config template
tn init

## Generate Tinet config file
tn init > spec.yaml
```

## tn print
```
tn print -c spec.yaml
```

## tn ps
```
## Output docker and netns info cmd
tn ps -c spec.yaml

## Output docker and netns info
tn ps -c spec.yaml | sudo sh -x
```

## tn pull
```
tn pull -c spec.yaml

## Execute docker pull
tn pull -c spec.yaml | sudo sh -x
```

## tn reconf
```
tn reconf -c spec.yaml

## down, up, conf
tn reconf -c spec.yaml | sudo sh -x
```

## tn reup
```
tn reup -c spec.yaml

## down, up
tn reup -c spec.yaml | sudo sh -x
```

## tn up
```
tn up -c spec.yaml

## up
tn up -c spec.yaml | sudo sh -x
```

## tn upconf
```
tn upconf -c spec.yaml

## up, conf
tn upconf -c spec.yaml | sudo sh -x
```

## tn version
```
tn version
```
