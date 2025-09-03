package user

import (
	"bufio"
	"fmt"
	"net"
	"strings"
)

func ReadMessage(conn net.Conn) {
	reader := bufio.NewReader(conn)
	for {
		msg, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("服务器连接断开:", err)
			return
		}
		fmt.Println(strings.TrimSpace(msg))
	}
}
