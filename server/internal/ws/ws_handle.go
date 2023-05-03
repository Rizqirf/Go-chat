package ws

import (
	"fmt"
	"net/http"
	"strconv"
	"sync"

	"github.com/gorilla/websocket"
	"github.com/labstack/echo"
)


var (
	lock = sync.Mutex{}
	roomIdSeq = 0
)

type Handler struct {
	hub *Hub
}

func NewHandler(h *Hub) *Handler {
	return &Handler{hub: h}
}

type CreateRoomReq struct {
	Id string `json:"id"`
	Name string `json:"name"`
}

type JoinRoomReq struct {
	Id string `json:"id"`
	Name string `json:"name"`
}

type GetRoomsRes struct {
	Id string `json:"id"`
	Name string `json:"name"`
}

type GetClientsRes struct {
	Id string `json:"id"`
	Username string `json:"username"`
}

func(h *Handler) CreateRoom (c echo.Context) error {
	lock.Lock()
	defer lock.Unlock()

	var req CreateRoomReq

	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	h.hub.Rooms[strconv.Itoa(roomIdSeq)] = &Room{
		Id: strconv.Itoa(roomIdSeq),
		Name: req.Name,
		Clients: make(map[string]*Client),
	}

	res := map[string]string{}

	res["Id"]=strconv.Itoa(roomIdSeq)
	res["Name"]=req.Name
	

	roomIdSeq++

	return c.JSON(http.StatusOK,res)
}

var upgrader = websocket.Upgrader{
	ReadBufferSize: 1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(r *http.Request) bool {
		// origin := r.Header.Get("origin")
		// return origin == "http://localhost:3000"

		return true
	},
}

func(h *Handler) JoinRoom (c echo.Context) error {
	lock.Lock()
	defer lock.Unlock()

	conn,err := upgrader.Upgrade(c.Response().Writer,c.Request(),nil)
	if err != nil {
		return c.JSON(http.StatusBadRequest,err)
	}

	roomId := c.Param("roomId")
	clientId := c.QueryParams().Get("userId")
	username := c.QueryParams().Get("username")

	cl := &Client{
		Conn : conn,
		Message: make(chan *Message,10),
		Id : clientId,
		RoomId: roomId,
		Username: username,
	}

	m := &Message{
		Content: fmt.Sprintf("%s has joined the room",username),
		RoomId: roomId,
		Username: username,
	}

	h.hub.Register <- cl
	h.hub.Broadcast <- m

	go cl.WriteMessage()
	cl.ReadMessage(h.hub)

	return c.JSON(http.StatusOK,nil)
}

func (h *Handler) GetRooms (c echo.Context) error {
	lock.Lock()
	defer lock.Unlock()

	rooms := make([]GetRoomsRes,0)

	for _,r := range h.hub.Rooms {
		rooms = append(rooms, GetRoomsRes{
			Id:r.Id,
			Name: r.Name,
		})
	}

	return c.JSON(http.StatusOK,rooms)
}

func (h *Handler) GetClients (c echo.Context) error {
	lock.Lock()
	defer lock.Unlock()
	
	var clients []GetClientsRes
	roomId := c.Param("roomId")


	if _,ok := h.hub.Rooms[roomId];!ok {
		clients = make([]GetClientsRes,0)
		return c.JSON(http.StatusOK,clients)
	}

	for _,c := range h.hub.Rooms[roomId].Clients {
		clients = append(clients, GetClientsRes{
			Id:c.Id,
			Username: c.Username,
		})
	}

	return c.JSON(http.StatusOK,clients)
}

