package main

import (
	"encoding/json"
	"log"
	"net"
	"net/http"
)

type handshakeResp struct {
	InterestedIn []string
	Name         string
	Author       string
	Org          string
	Website      string
}

type VolumeExtensionReq struct {
	HostPath    string
	ContainerID string
}

type VolumeExtensionResp struct {
	ModifiedHostPath string
}

func main() {
	http.HandleFunc("/v1/handshake", func(w http.ResponseWriter, r *http.Request) {
		err := json.NewEncoder(w).Encode(&handshakeResp{
			InterestedIn: []string{"volume"},
			Name:         "stub",
			Author:       "progrium",
			Org:          "no, wut?",
			Website:      "http://progrium.com",
		})
		if err != nil {
			log.Println("handshake encode:", err)
			http.Error(w, "encode error", http.StatusInternalServerError)
			return
		}
		log.Println("handshake success")
	})

	http.HandleFunc("/v1/volume/volumes", func(w http.ResponseWriter, r *http.Request) {
		defer r.Body.Close()
		var extReq VolumeExtensionReq
		err := json.NewDecoder(r.Body).Decode(&extReq)
		if err != nil {
			log.Println("bad request:", err)
			http.Error(w, "bad request", http.StatusBadRequest)
			return
		}

		log.Println("req:", extReq)

		err = json.NewEncoder(w).Encode(&VolumeExtensionResp{
			ModifiedHostPath: extReq.HostPath,
		})
		if err != nil {
			log.Println("resp encode:", err)
		}
	})

	sock := "/var/run/docker-plugin/p.s"
	l, err := net.Listen("unix", sock)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("listening on %s ...\n", sock)
	log.Fatal(http.Serve(l, nil))
}
