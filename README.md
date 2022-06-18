# dir2map

procfs and sysfs have lots of small one-line files.
Checking them one by one with cat is annoying.
So I've implemented simple code to look those files inside dir at once.

call examples
```bash
go run . /proc/sys/kernel/
go run dir2map.go /sys/class/power_supply/BAT0/
# after building
./dir2map /sys/devices/system/cpu/cpu0/cache/index0
```

