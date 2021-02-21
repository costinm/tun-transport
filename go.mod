module github.com/costinm/tungate

go 1.16

//replace github.com/google/netstack => github.com/costinm/netstack v0.0.0-20190601172006-f6e50d4d2856

//replace github.com/google/netstack => ../netstack

//replace gvisor.dev/gvisor => ../gvisor

//replace github.com/costinm/ugate => ../ugate

require (
	github.com/costinm/ugate v0.0.0-20210106052904-4da1a58a92e6
	github.com/eycorsican/go-tun2socks v1.16.11
	github.com/songgao/water v0.0.0-20200317203138-2b4b6d7c09d
)
