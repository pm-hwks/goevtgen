# goevtgen
Large scale Windows eventlog generation in golang

# Install / Download
Go to the releases and download the latest version of *"go_build_goevtgen_go.exe"*

# Sample usage

```
C:\Users\Administrator\go\src\goevtgen\output>go_build_goevtgen_go.exe --eps 1000 --duration 3
Initiating goevtgen with the following parameters !
EPS : 1000
Duration : 3
Event Source : samplesource
Event ID : 123
Event Type : Info
Message : This is a sample log


Tick at 2019-10-08 07:05:59.7321393 +0000 GMT m=+1.008584401
Sent 1000 messages to evt log in 56.7694ms
Tick at 2019-10-08 07:06:00.7357293 +0000 GMT m=+2.012174401
Sent 1000 messages to evt log in 37.2347ms
Ticker stopped
```

# Usage help
```
C:\Users\Administrator\go\src\goevtgen\output>go_build_goevtgen_go.exe --help
NAME:
   goevtgen - Generate windows eventlogs

USAGE:
   go_build_goevtgen_go.exe [global options] command [command options] [arguments...]

VERSION:
   0.2

COMMANDS:
   help, h  Shows a list of commands or help for one command

GLOBAL OPTIONS:
   --eps value          no of events per second to be generated (tested with 20K eps) (default: 100)
   --duration value     No of seconds to run the generation program (default: 3)
   --EventSource value  Event source name (default: "samplesource")
   --EventID value      Event ID (default: 123)
   --EventType value    Event Type (Info, Warning, Error) (default: "Info")
   --Message value      Event payload - message (default: "This is a sample log")
   --help, -h           show help
   --version, -v        print the version

C:\Users\Administrator\go\src\goevtgen\output>
```
