package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net"
	"net/http"
	"os"

	"github.com/costinm/tungate"
	"github.com/costinm/tungate/pkg/lwip"
	"github.com/costinm/ugate"
	"github.com/costinm/ugate/pkg/auth"
	"github.com/costinm/ugate/pkg/ugatesvc"
)

// Similar with the sample micro-gate, but adding a TUN capture.
// Used to experiment with TUN instead of iptables capture.

func main() {
	config := ugatesvc.NewConf(".")

	auth := auth.NewAuth(config, "", "h.webinf.info")

	cfg := &ugate.GateCfg{
		BasePort: 12100,
	}

	data, err := ioutil.ReadFile("h2gate.json")
	if err != nil {
		json.Unmarshal(data, cfg)
	}

	// By default, pass through using net.Dialer
	ug := ugatesvc.NewGate(&net.Dialer{}, auth, cfg)

	fd, err := tungate.OpenTun("dmesh1")
	if err != nil {
		log.Fatal("Failed to open tun", err)
	}

	tun := lwip.NewTUNFD(fd,ug, ug)
	ug.TUNUDPWriter = tun

	log.Println("TUN started ", tun)

	// direct TCP connect to local iperf3 and fortio (or HTTP on default port)
	ug.Add(&ugate.Listener{
		Address: ":12111",
		ForwardTo: "localhost:5201",
	})

	// VIP capture
	ug.Add(&ugate.Listener{
		Address: "10.13.0.1:5201",
		ForwardTo: "localhost:5201",
	})

	log.Println("Started debug on 12119, UID/GID", os.Getuid(), os.Getegid())
	err = http.ListenAndServe(":12119", nil)
	if err != nil {
		log.Fatal(err)
	}
}
