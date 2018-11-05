package DouyuBarrage

import (
	"bytes"
	"encoding/binary"
	"flag"
	"io"
	"log"
	"net"
	"strconv"
	"time"
)

const (
	SERVER_PROTOCOL string = "tcp4"
	BARRAGE_SERVER  string = "openbarrage.douyutv.com"
	SERVER_PORT     string = "8601"
)

var (
	tcpconn *net.TCPConn
	roomid  = flag.String("rid", "90016", "room id.")
)

//initconn will do some initializing operations, including:
//	1. connect to barrage server(openbarrage.douyutv.com:8601) via TCP protocol
//	2. make log-in request to barrage server
//	3. if log-in succeeds, then will receive correspond message from barrage server.
//	4. send barrage group request to barrage server. (due to huge amounts of barrages, barrage grouping is necessary)
//	5. server will add client to specified barrage group after receiving barrage group request.
func initconn() {
	//step1. connect to server
	tcpaddr, err := net.ResolveTCPAddr(SERVER_PROTOCOL, BARRAGE_SERVER+":"+SERVER_PORT)
	checkErr(err)

	tcpconn, err := net.DialTCP(SERVER_PROTOCOL, nil, tcpaddr)
	checkErr(err)
	log.Println("TCP connect ok.")

	//step2. seng log-in request
	loginreq := []byte("type@=loginreq" + "/roomid@=" + *roomid + "/")
	sendmsg(tcpconn, loginreq)

	//step3. check the server response
	srvres, err := readall(tcpconn)
	log.Printf("log-in response:%s\n", srvres)
	checkErr(err)

	//step4. send barrage group request to server
	groupreq := []byte("type@=joingroup/rid@=" + *roomid + "/gid@=-9999/")
	sendmsg(tcpconn, groupreq)

	//step5. see what response will be returned
	srvres2, err := readall(tcpconn)
	log.Printf("groupreq response:%s\n", srvres2)
	checkErr(err)

}

func sendmsg(tc *net.TCPConn, b []byte) {
	msglen := len(b) + 8 + 1
	msgtype := 689
	var (
		msglenbuf  [4]byte
		msgtypebuf [4]byte
	)
	binary.LittleEndian.PutUint32(msglenbuf[:], uint32(msglen))
	binary.LittleEndian.PutUint32(msgtypebuf[:], uint32(msgtype))

	msghead := append(append(msglenbuf[:], msglenbuf[:]...), msgtypebuf[:]...)
	msgall := append(append(msghead, b...), 0x00)
	_, err := tc.Write(msgall)
	log.Printf("sent->%s\n", msgall)
	checkErr(err)

	/*
		for sent := 0; sent < len(b); {
			sgst, err := tc.Write(b)
			log.Printf("sent->%s\n", b)
			checkErr(err)
			sent = sent + sgst
		}*/

}

func readall(tc *net.TCPConn) ([]byte, error) {
	result := bytes.NewBuffer(nil)
	var buf [512]byte

	for {
		n, err := tc.Read(buf[0:])
		result.Write(buf[0:n])
		if err != nil {
			if err == io.EOF {
				break
			}
			return nil, err
		}
	}
	log.Printf("response->%s\n", result.Bytes())
	return result.Bytes(), nil
}

func keeplive() {
	//msg format: type@=keeplive/tick@=1541401463/
	keephead := []byte("type@=keeplive/tick@=")
	for {
		tn := strconv.AppendInt(keephead, time.Now().Unix(), 10) // get string type of unix timestamp
		livemsg := append(tn, []byte("/\\0")...)
		sendmsg(tcpconn, livemsg)
		_, err := readall(tcpconn)
		checkErr(err)
		//log.Printf("keeplive->response:%s\n", res)
		time.Sleep(45 * time.Second)
	}
}

func checkErr(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
