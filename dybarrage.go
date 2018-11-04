package DouyuBarrage

import (
	"log"
	"net"
)

const (
	SERVER_PROTOCOL string = "tcp4"
	BARRAGE_SERVER  string = "openbarrage.douyutv.com"
	SERVER_PORT     string = "8601"
)

//initconn will do some initializing operations, including:
//	1. connect to barrage server(openbarrage.douyutv.com:8601) via TCP protocol
//	2. make log-in request to barrage server
//	3. if log-in succeeds, then will receive correspond message from barrage server.
//	4. send barrage group request to barrage server. (due to huge amounts of barrages, barrage grouping is necessary)
//	5. server will add client to specified barrage group after receiving barrage group request.
func initconn() {
	//step1.
	tcpaddr, err := net.ResolveTCPAddr(SERVER_PROTOCOL, BARRAGE_SERVER+":"+SERVER_PORT)
	if err != nil {
		log.Fatal(err)
	}
	tcpconn, err := net.DialTCP(SERVER_PROTOCOL, nil, tcpaddr)
	if err != nil {
		log.Fatal(err)
	}

}
