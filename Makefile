GO=go
GO_CONF=build 

all:
	${GO} ${GO_CONF} -o termini src/main.go 

clean:
	rm termini

