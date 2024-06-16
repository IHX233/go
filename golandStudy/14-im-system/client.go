package main
import (
	"fmt"
	"net"
	"flag"
	"io"
	"os"
)
type Client struct{
	ServerIp string
	ServerPort int
	Name string
	conn net.Conn
	mod int //代表当前客户端的模式
}
func (client *Client) menu() bool{
	var mod int
	fmt.Println("0:退出")
	fmt.Println("1:公聊模式")
	fmt.Println("2:私聊模式")
	fmt.Println("3:更新用户名")
	fmt.Scanln(&mod)
	if mod>=0&&mod<=3{
		client.mod = mod
		return true
	}else{
		fmt.Println("请输入合法范围内的数字")
		return false
	}
}

//修改用户名
func(client *Client) UpdateName() bool{
	fmt.Println("请输入新的用户名")
	fmt.Scanln(&client.Name)
	sendMsg := "rename|" + client.Name + "\n"
	_,err := client.conn.Write([]byte(sendMsg))
	if err != nil{
		fmt.Println("conn.Write error:",err)
		return false
	}
	return true
}

//公聊模式
func(client *Client) PublicChat(){
	var chatMsg string
	fmt.Println("请输入聊天内容，exit退出")
	fmt.Scanln(&chatMsg)
	for chatMsg != "exit"{
		if len(chatMsg) !=0 {
			sendMsg := chatMsg + "\n"
			_,err := client.conn.Write([]byte(sendMsg))
			if err != nil{
				fmt.Println("conn write error:",err)
				break
			}
		}
		chatMsg = ""
		fmt.Println("请输入聊天内容，exit退出")
		fmt.Scanln(&chatMsg)
	}
}

//查询当前在线用户

func(client *Client) SelectUsers(){
	sendMsg := "who\n"
	_,err := client.conn.Write([]byte(sendMsg))
	if err != nil{
		fmt.Println("conn write error:",err)
	}
}

//私聊模式
func(client *Client) PrivateChat(){
	var remoteName string
	var chatMsg string
	client.SelectUsers()
	fmt.Println("请输入聊天对象，exit退出")
	fmt.Scanln(&remoteName)
	for remoteName != "exit"{
		fmt.Println("请输入聊天内容，exit退出")
		fmt.Scanln(&chatMsg)
		for chatMsg != "exit"{
			if len(chatMsg) !=0 {
				sendMsg := "to|" + remoteName + "|" + chatMsg + "\n"
				_,err := client.conn.Write([]byte(sendMsg))
				if err != nil{
					fmt.Println("conn write error:",err)
					break
				}
			}
			chatMsg = ""
			fmt.Println("请输入聊天内容，exit退出")
			fmt.Scanln(&chatMsg)
		}
		client.SelectUsers()
		fmt.Println("请输入聊天对象，exit退出")
		fmt.Scanln(&remoteName)
	}
}
func (client *Client) Run(){
	for client.mod !=0{
		for client.menu() != true{}
		switch client.mod {
		case 1:
			client.PublicChat()
			break
		case 2:
			client.PrivateChat()
			break
		case 3:
			client.UpdateName()
			break
		}
	}
}

//处理server返回的消息
func (client *Client) DealResponse(){
	//一旦client.conn有数据，就直接copy到stdout上，永久阻塞
	io.Copy(os.Stdout,client.conn)
}
func NewClient(serverIp string,serverPort int) *Client{
	//创建用户对象
	client := &Client{
		ServerIp:serverIp,
		ServerPort:serverPort,
		mod:999,
	}
	//链接server
	conn,err := net.Dial("tcp",fmt.Sprintf("%s:%d",serverIp,serverPort))
	if err != nil{
		fmt.Println("net.Dial error:",err)
		return nil
	}
	client.conn = conn
	//返回对象
	return client
}
var serverIp string
var serverPort int
// ./client -ip 127.0.0.1 -port 8888
func init(){
	flag.StringVar(&serverIp,"ip","127.0.0.1","设置默认服务器IP地址（默认是127.0.0.1）")
	flag.IntVar(&serverPort,"port",8888,"设置默认服务器端口号（默认是8888）")
}
func main(){
	flag.Parse()
	client := NewClient(serverIp,serverPort)
	if client == nil{
		fmt.Println("链接服务器失败")
		return
	}
	go client.DealResponse()
	fmt.Println("链接服务器成功")
	client.Run()
}