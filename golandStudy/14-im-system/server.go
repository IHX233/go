package main
import (
	"fmt"
	"net"
	"sync"
	"io"
	"time"
)
type Server struct{
	Ip string
	Port int
	OnlineMap map[string]*User
	mapLock sync.RWMutex
	Message chan string
}

func NewServer(ip string,port int) *Server{
	server := &Server{
		Ip:ip,
		Port:port,
		OnlineMap:make(map[string]*User),
		Message:make(chan string),
	}
	return server;
}
//监听Message广播消息channel的goroutine，一旦有消息就会发送给全部在线的User
func (this *Server) ListenMessage(){
	for{
		msg := <-this.Message
		this.mapLock.Lock()
		for _,cli := range this.OnlineMap{
			cli.C <- msg
		}
		this.mapLock.Unlock()
	}
}
//广播消息的方法
func (this *Server) BroadCast(user *User,msg string){
	sendMsg := "[" + user.Addr + "]" + user.Name + ":" + msg
	this.Message <- sendMsg
}
func(this *Server) Handler(conn net.Conn){
	user := NewUser(conn,this)
	user.Online()
	//监听用户是否活跃的channel
	isLive := make(chan bool)
	//接收客户端发来的消息
	go func(){
		buf := make([]byte,4096)
		for{
			n,err := conn.Read(buf)
			if n == 0 {
				user.Offline()
				return
			}
			if err!=nil && err!=io.EOF{
				fmt.Println("Conn Read err:",err)
			}
			//提取用户的消息（去除“\n”）
			msg := string(buf[:n-1])
			user.Domessage(msg)
			isLive <- true
		}
	}()
	for{
		select{
		case<-isLive:
			//当前用户是活跃的，应该重置定时器
			//什么都不做，为了激活select,更新下面的定时器
		case<-time.After(time.Minute*10):
			//已经超时
			//将当前User强制关闭
			user.SendMsg("你被踢了")
			close(user.C)
			conn.Close()
			//退出当前handle
			return
		}
	}
	
}

func (this *Server) Start(){
	listener,err := net.Listen("tcp",fmt.Sprintf("%s:%d",this.Ip,this.Port))
	if err!=nil{
		fmt.Println("net.Listen err:",err)
		return
	}
	defer listener.Close()

	go this.ListenMessage()

	for{
		conn,err:=listener.Accept()
		if err!=nil{
			fmt.Println("net.Accept err:",err)
			continue
		}
		go this.Handler(conn)
	}
}