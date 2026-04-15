package main

type client chan<- string //只能用来往通道里面发送消息的。

var (
	entering = make(chan client) //用来接收连接的管道。
	leaving  = make(chan client)
	messages = make(chan string) //所有客户端共享的管道
)

// 一个用户连接进来分为连接fd和消息fd，这点在epoll中有。
func Broadcaster() {

	clients := make(map[client]bool) //所有连接的客户端

	for {
		select {
		case msg := <-messages: //这等channel有消息可读取
			//广播消息给所有的客户端
			for cli := range clients {
				cli <- msg
			}

		case cli := <-entering:
			clients[cli] = true
		case cli := <-leaving:
			delete(clients, cli)
			close(cli)
		}
	}
}
