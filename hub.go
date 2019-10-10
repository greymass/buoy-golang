package main

type Message struct {
  data []byte
  uuid string
}

type Subscription struct {
  conn *Connection
  uuid string
}

type Hub struct {
  broadcast chan Message
  register chan Subscription
  unregister chan Subscription
  uuids map[string]map[*Connection]bool
}

// func newHub() *Hub {
// 	return &Hub{
// 		broadcast:  make(chan []byte),
// 		register:   make(chan *Subscription),
// 		unregister: make(chan *Subscription),
// 		uuids:    make(map[*Client]bool),
// 	}
// }

// func newHub() *Hub {
// 	return &Hub{
//     broadcast:  make(chan Message),
//     register:   make(chan Subscription),
//     unregister: make(chan Subscription),
//     uuids:      make(map[string]map[*Connection]bool),
// 	}
// }

var h = Hub{
  broadcast:  make(chan Message),
  register:   make(chan Subscription),
  unregister: make(chan Subscription),
  uuids:      make(map[string]map[*Connection]bool),
}

func (h *Hub) run() {
    for {
        select {
        case s := <-h.register:
            connections := h.uuids[s.uuid]
            if connections == nil {
                connections = make(map[*Connection]bool)
                h.uuids[s.uuid] = connections
            }
            h.uuids[s.uuid][s.conn] = true
        case s := <-h.unregister:
            connections := h.uuids[s.uuid]
            if connections != nil {
                if _, ok := connections[s.conn]; ok {
                    delete(connections, s.conn)
                    close(s.conn.send)
                    if len(connections) == 0 {
                        delete(h.uuids, s.uuid)
                    }
                }
            }
        case m := <-h.broadcast:
            connections := h.uuids[m.uuid]
            for c := range connections {
                select {
                case c.send <- m.data:
                default:
                    close(c.send)
                    delete(connections, c)
                    if len(connections) == 0 {
                        delete(h.uuids, m.uuid)
                    }
                }
            }
        }
    }
}
