package handler

import (
	"github.com/labstack/echo"
	. "github.com/JermineHu/ait/consts"
	"strings"
	"fmt"
	"sync"
	"github.com/gorilla/websocket"
)

const MAX_CONNECTION int = 100
const JOIN_ROOM_FAILED int = -1
const Debug = true

type ChatRoom struct {
	sync.Mutex
	clients   map[int]*websocket.Conn
	currentId int
}

func (cr *ChatRoom)joinRoom(ws *websocket.Conn) int {
	cr.Lock()
	defer cr.Unlock()
	if len(cr.clients) >= MAX_CONNECTION {
		return JOIN_ROOM_FAILED
	}
	cr.currentId++
	cr.clients[cr.currentId] = ws
	return cr.currentId
}
func (cr *ChatRoom)leftRoom(id int) {
	delete(cr.clients, id)
}
func (cr *ChatRoom)sendMessage(msg string) {
	for _, ws := range cr.clients {
		if err := ws.WriteMessage(websocket.TextMessage, []byte(msg)); err != nil {
			log4Demo("发送失败，Err：" + err.Error())
			continue
		}
	}
}

var room ChatRoom
func init() {
	roomMap := make(map[int]*websocket.Conn, MAX_CONNECTION)
	room = ChatRoom{clients:roomMap, currentId:0}
}

func log4Demo(msg string) {
	if Debug {
		fmt.Println(msg)
	}
}

func WrapChatRoutes(c *echo.Group) {
	ChatHandler(c)
}

var (
	upgrader = websocket.Upgrader{}
)
 func ChatHandler(c *echo.Group){
	h:= func(c echo.Context) (err error){

		ws, err := upgrader.Upgrade(c.Response(), c.Request(), nil)//将当前http请求升级为websocket
		if err != nil {
			return
		}
		defer ws.Close()//跳出函数前关闭socket

		var id int
		if id = room.joinRoom(ws); id == JOIN_ROOM_FAILED { //将当前的socket对象加入room池子，用于后续的批量广播
			err=ws.WriteMessage(websocket.TextMessage,[]byte( "加入聊天室失败"))
			if err != nil {
				c.Logger().Error(err)
			}
			return
		}

		defer room.leftRoom(id) //离开房间就要从池子里删除
		ipAddress := strings.Split(ws.RemoteAddr().String(), ":")[0] + "："

		for {
			_, msg, err := ws.ReadMessage()//读取消息
			if err != nil {
				c.Logger().Error(err)
				return err
			}
			send_msg:= ipAddress + string(msg)
			room.sendMessage(send_msg)//将消息广播给进入该room的所有人，此处优化方案可以采用Kafka提高异步处理，提高并发效率
		}
		return
	}
	c.POST(ChatRoomRoute, h)

}
