package main
import (
	"net"
	"strings"
)
type User struct{
	Name string
	Addr string
	C chan string
	conn net.Conn
	server *Server
}

//创建一个用户api
func NewUser(conn net.Conn,server *Server) *User{
	userAddr := conn.RemoteAddr().String()
	user := &User{
		Name:userAddr,
		Addr:userAddr,
		C:make(chan string),
		conn:conn,
		server:server,
	}
	go user.ListenMessage()
	return user
}

//监听user channel 方法，一旦有消息，就直接发送给客户端
func (this *User) ListenMessage(){
	for{
		msg:=<-this.C
		this.conn.Write([]byte(msg + "\n"))
	}
}

//向当前用户发送消息
func (this *User) SendMsg(msg string){
	this.conn.Write([]byte(msg))
}
//用户上线业务
func (this *User) Online(){
	this.server.mapLock.Lock()
	this.server.OnlineMap[this.Name] = this
	this.server.mapLock.Unlock()
	//广播当前用户上线信息
	this.server.BroadCast(this,"已上线")
}
//用户下线业务
func (this *User) Offline(){
	this.server.mapLock.Lock()
	delete(this.server.OnlineMap,this.Name)
	this.server.mapLock.Unlock()
	//广播当前用户下线信息
	this.server.BroadCast(this,"已下线")
}
//用户处理信息业务
func (this *User) Domessage(msg string){
	if msg=="who"{
		//查询当前在线的客户有哪些
		this.server.mapLock.Lock()
		for _,user := range this.server.OnlineMap{
			onlineMessage := "[" + user.Addr + "]" + user.Name + ":" + "在线...\n"
			this.SendMsg(onlineMessage)
		}
		this.server.mapLock.Unlock()

	}else if len(msg)>7 && msg[:7]=="rename|"{
		newName := strings.Split(msg,"|")[1]
		//判断name是否存在
		_,ok := this.server.OnlineMap[newName]
		if ok{
			this.SendMsg("当前用户名被使用")
		}else{
			this.server.mapLock.Lock()
			delete(this.server.OnlineMap,this.Name)
			this.server.OnlineMap[newName] = this
			this.server.mapLock.Unlock()
			this.Name = newName
			this.SendMsg("您已经更新用户名为:"+this.Name+"\n")
		}
	}else if len(msg)>3 && msg[:3]=="to|"{
		remoteName := strings.Split(msg,"|")[1]
		if(remoteName==""){
			this.SendMsg("消息格式不正确，请使用\"to|sxj|hi\"格式")
			return
		}
		remoteUser,ok := this.server.OnlineMap[remoteName]
		if !ok {
			this.SendMsg("用户不存在")
			return
		}
		content := strings.Split(msg,"|")[2]
		if content == "" {
			this.SendMsg("消息为空，请重新发送")
			return
		}
		remoteUser.SendMsg(this.Name+"对您说："+content)

	}else{
		this.server.BroadCast(this,msg)
	}
}