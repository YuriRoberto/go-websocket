package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/websocket"
)

type Server struct {
	conns map[*websocket.Conn]bool
}

type Client struct {
	conn *websocket.Conn
	name string
}

type Message struct {
	Date string `json:"date"`
	Name string `json:"name"`
	Data string `json:"data"`
}

var temp = template.Must(template.ParseGlob("templates/*.html"))

func index(w http.ResponseWriter, r *http.Request) {
	temp.ExecuteTemplate(w, "Index", nil)
}

func NewServer() *Server {
	return &Server{
		conns: make(map[*websocket.Conn]bool),
	}
}

func (s *Server) handleWS(client *Client) {
	fmt.Println("new incoming connection from client:", client.conn.RemoteAddr())

	s.conns[client.conn] = true

	s.readLoop(client)
}

func (s *Server) readLoop(client *Client) {
	defer client.conn.Close()

	for {
		messageType, msg, err := client.conn.ReadMessage()
		if err != nil {
			fmt.Println("Erro ao ler a mensagem:", err)
			break
		}

		broadcastMsg := Message{
			Date: time.Now().Format("2006-01-02 15:04:05"),
			Name: client.name,
			Data: string(msg),
		}

		jsonMessage, err := json.Marshal(broadcastMsg)
		if err != nil {
			fmt.Println("Erro ao codificar a mensagem em JSON:", err)
			return
		}

		fmt.Println("Tipo de mensagem:", messageType)
		fmt.Printf("[%s] %s: %s\n", time.Now().Format("2006-01-02 15:04:05"), client.name, string(msg))

		//response := []byte("Obrigado pela mensagem!")
		//if err := client.conn.WriteMessage(websocket.TextMessage, response); err != nil {
		//	fmt.Println("Erro ao escrever a resposta:", err)
		//	break
		//}
		s.broadcast(jsonMessage, client)
	}
}

func handleWS(server *Server) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		upgrader := websocket.Upgrader{
			CheckOrigin: func(r *http.Request) bool {
				// Permitir todas as origens por enquanto
				return true
			},
		}

		// Atualiza a conexão HTTP para uma conexão WebSocket
		ws, err := upgrader.Upgrade(w, r, nil)
		if err != nil {
			log.Println("Erro ao atualizar a conexão para WebSocket:", err)
			return
		}

		defer ws.Close()

		if err := ws.WriteMessage(websocket.TextMessage, []byte("Enter your name: ")); err != nil {
			log.Println("Erro ao solicitar o nome do usuário:", err)
			return
		}

		// Lê o nome do usuário
		_, name, err := ws.ReadMessage()
		if err != nil {
			log.Println("Erro ao ler o nome do usuário:", err)
			return
		}

		client := &Client{
			conn: ws,
			name: string(name),
		}

		greeting := fmt.Sprintf("Hello %s! Bem vindo ao chat.", client.name)
		if err := ws.WriteMessage(websocket.TextMessage, []byte(greeting)); err != nil {
			log.Println("Erro ao enviar saudação:", err)
			return
		}

		// Chama o manipulador real para lidar com a conexão WebSocket
		server.handleWS(client)
	}
}

func (s *Server) broadcast(message []byte, sender *Client) {
	for client := range s.conns {
		go func(c *websocket.Conn) {
			if c != sender.conn {
				if err := c.WriteMessage(websocket.TextMessage, message); err != nil {
					fmt.Println("Erro ao enviar mensagem:", err)
				}
			}
			//if err := ws.WriteMessage(websocket.TextMessage, b); err != nil {
			//	fmt.Println("write error:", err)
			//}
		}(client)
	}
}

func main() {
	// Você pode usar a função handleWS como manipulador HTTP
	server := NewServer()

	http.HandleFunc("/ws", handleWS(server))
	http.HandleFunc("/", index)

	go func() {
		log.Println("Servidor iniciado na porta 3000")
		err := http.ListenAndServe(":3000", nil)
		if err != nil {
			log.Fatal("Erro ao iniciar o servidor:", err)
		}
	}()

	log.Println("Frontend iniciado na porta 8000")
	err := http.ListenAndServe(":8000", nil)
	if err != nil {
		log.Fatal("Erro ao iniciar o front:", err)
	}
}
