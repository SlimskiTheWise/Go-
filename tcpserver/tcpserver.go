package tcpserver

import (
	"fmt"
	"net"
)

func RequestHandler(c net.Conn) {
	data := make([]byte, 4096) // 4096크기의 바이트 슬라이스를 생성

	for {
		n, err := c.Read(data) // 클라이언트에서 받은 데이터를 읽음
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Println(string(data[:n])) //데이터 출력

		_, err = c.Write(data[:n]) // 클라이언트로 보냄
		if err != nil {
			fmt.Println(err)
			return
		}
	}
}
