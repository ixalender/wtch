# wtch

Periodically executes specified system command.

## Why?!
It's a simpler variant of wtch(1) Linux util, but with needed milliseconds interval and numbered repeat feature.


## Usage

```
wtch [global options] help | "system command with args"

COMMANDS:
   help, h  Shows a list of commands or help for one command

GLOBAL OPTIONS:
   --interval value, -n value  interval in milliseconds (default: "1000")
   --repeat value, -r value    times to repeat (default: 0 - infinity) (default: "0")
   --help, -h                  show help (default: false)
   --version, -v               print the version (default: false)
```

## Example

```
$ wtch -n 500 "df -h"
2022/04/27 21:28:14 'df -h' every 500 millisecond
Filesystem       Size   Used  Avail Capacity iused     ifree %!i(MISSING)used  Mounted on
/dev/disk1s1s1  113Gi   23Gi   30Gi    44%! (MISSING)314738560    0%!/(MISSING)
devfs           194Ki  194Ki    0Bi   100%! (MISSING)        0  100%!/(MISSING)dev
/dev/disk1s4    113Gi  4.0Gi   30Gi    12%! (MISSING)314738560    0%!/(MISSING)System/Volumes/VM
/dev/disk1s2    113Gi  286Mi   30Gi     1%! (MISSING)314738560    0%!/(MISSING)System/Volumes/Preboot
/dev/disk1s6    113Gi  103Mi   30Gi     1%! (MISSING)314738560    0%!/(MISSING)System/Volumes/Update
/dev/disk1s5    113Gi   54Gi   30Gi    65%! (MISSING)314738560    0%!/(MISSING)System/Volumes/Data
map auto_home     0Bi    0Bi    0Bi   100%   /System/Volumes/Data/home
/dev/disk1s1    113Gi   23Gi   30Gi    44%! (MISSING)314738560    0%!/(MISSING)System/Volumes/Update/mnt1

```