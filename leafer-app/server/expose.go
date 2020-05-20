package server

import (
	"log"

	"gitlab.com/NebulousLabs/go-upnp"
)

// Expose server through UPNP Router (proto to test it)
func Expose() {
	// connect to router
	d, err := upnp.Discover()
	if err != nil {
		log.Fatal(err)
	}

	// discover external IP
	ip, err := d.ExternalIP()
	if err != nil {
		log.Fatal(err)
	}

	// forward a port
	err = d.Forward(8888, "Leafer Media Server")
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Leafer Media Server exposed to: http://%s:%v", ip, 8888)
}

// Unexpose server through UPNP Router (proto to test it)
func Unexpose() {
	// connect to router
	d, err := upnp.Discover()
	if err != nil {
		log.Fatal(err)
	}

	// un-forward a port
	err = d.Clear(8888)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Leafer Media Server not exposed anymore")
}
