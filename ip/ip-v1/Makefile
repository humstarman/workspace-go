all: compile scp

compile:
	@CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o ip ./ip.go

scp:
	@scp ./ip root@192.168.100.162:/tmp/ip
