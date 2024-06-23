# Demo
- [box_Interface.go](box_Interface.go) interface and pointer receiver
- [error_is_not_pointer.go](error_is_not_pointer.go) error is not a pointer

# compile
```shell
# get importcfg source
go build -x -n -v *.go 2>&1 | sed  -n "/^# import config/,/EOF$/p" |grep -v EOF > importcfg
# this will generate xxx.o
go tool compile -N -l -importcfg importcfg xxx.go
# get the coomplie of xxx.go
go tool objdump xxxx.o >> xxx.txt
```

