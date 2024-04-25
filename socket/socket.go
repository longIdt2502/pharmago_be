package socket

import (
	"encoding/json"
	"net/http"
	"strconv"
	"sync"

	"github.com/gorilla/mux"
	"github.com/gorilla/websocket"
	db "github.com/longIdt2502/pharmago_be/db/sqlc"
	"github.com/rs/zerolog/log"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

// WebSocketClient đại diện cho một client WebSocket
type WebSocketClient struct {
	conn *websocket.Conn
}

var (
	clients   = make(map[*WebSocketClient]bool) // Danh sách các client đã kết nối
	clientsMu sync.Mutex
)

// handleMessage xử lý thông điệp từ client
func handleMessage(msg []byte) {
	// Xử lý các hành động từ thông điệp, ví dụ như gửi lại thông điệp tới tất cả các client
	broadcastMessage(msg)
}

// broadcastMessage gửi thông điệp tới tất cả các client
func broadcastMessage(msg []byte) {
	clientsMu.Lock()
	defer clientsMu.Unlock()
	for client := range clients {
		err := client.conn.WriteMessage(websocket.TextMessage, msg)
		if err != nil {
			client.conn.Close()
			delete(clients, client)
		}
	}
}

// WebsocketHandler xử lý kết nối WebSocket
func WebsocketHandler(w http.ResponseWriter, r *http.Request) {
	// Lấy ID từ URL path
	params := mux.Vars(r)
	id, _ := strconv.Atoi(params["company"])

	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Err(err).Msg("WS")
		return
	}
	defer conn.Close()

	client := &WebSocketClient{conn: conn}

	// Thêm client vào danh sách
	clientsMu.Lock()
	clients[client] = true
	clientsMu.Unlock()

	// Lắng nghe và xử lý thông điệp từ client
	for {
		_, message, err := conn.ReadMessage()
		if err != nil {
			log.Err(err)
			break
		}
		var payload db.CreateNotificationParams
		err = json.Unmarshal(message, &payload)
		if err != nil {
			log.Err(err)
		}
		if payload.Company.Int32 == int32(id) {
			// Xử lý thông điệp
			handleMessage(message)
		}
	}

	// Xóa client ra khỏi danh sách khi kết nối bị đóng
	clientsMu.Lock()
	delete(clients, client)
	clientsMu.Unlock()
}

func StartSocket(store db.Store) {
	router := mux.NewRouter()
	router.HandleFunc("/websocket/{company}", WebsocketHandler)
	http.ListenAndServe(":8000", router)
}
