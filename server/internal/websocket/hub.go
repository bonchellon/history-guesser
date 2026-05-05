package websocket

import (
	"context"
	"encoding/json"
	"net/http"
	"sync"
	"time"

	"history-guesser/server/internal/rounds"
	"history-guesser/server/internal/scoring"

	"github.com/go-chi/chi/v5"
	gws "github.com/gorilla/websocket"
)

type PlayerGuess struct { Lat, Lon float64 `json:"lat"`; Year int `json:"year"`; Submitted bool `json:"submitted"` }
type Player struct { ID, Name string `json:"id"`; Conn *gws.Conn `json:"-"`; Guess PlayerGuess `json:"guess"`; Ready bool `json:"ready"` }
type Room struct { Code, HostID string `json:"code"`; Players map[string]*Player `json:"players"`; Round rounds.Round `json:"round"`; EndsAt time.Time `json:"endsAt"` }
type Hub struct { rooms map[string]*Room; mu sync.Mutex; rounds rounds.Repository }

func NewHub(repo rounds.Repository) *Hub { return &Hub{rooms: map[string]*Room{}, rounds: repo} }
func Router(h *Hub) http.Handler { r:=chi.NewRouter(); r.Get("/health", func(w http.ResponseWriter,r *http.Request){w.Write([]byte("ok"))}); r.Get("/ws", h.HandleWS); return r }

var upgrader = gws.Upgrader{CheckOrigin: func(r *http.Request) bool { return true }}

func (h *Hub) HandleWS(w http.ResponseWriter, r *http.Request) {
	c, err := upgrader.Upgrade(w, r, nil); if err != nil { return }
	defer c.Close()
	playerID := r.URL.Query().Get("playerId"); if playerID=="" { playerID = time.Now().Format("150405.000") }
	for { _, data, err := c.ReadMessage(); if err != nil { h.disconnect(playerID); return }
		var m map[string]any; _ = json.Unmarshal(data,&m)
		h.handleMessage(context.Background(), c, playerID, m)
	}
}

func (h *Hub) handleMessage(ctx context.Context, c *gws.Conn, pid string, m map[string]any) {
	t := m["type"].(string)
	switch t {
	case "room.create":
		code := "ROOM" + pid[len(pid)-3:]
		h.mu.Lock(); h.rooms[code]=&Room{Code:code,HostID:pid,Players:map[string]*Player{pid:{ID:pid,Name:"Player "+pid,Conn:c}}}; room:=h.rooms[code]; h.mu.Unlock();
		h.broadcast(room, "room.state", room)
	case "room.join":
		code := m["code"].(string); h.mu.Lock(); room:=h.rooms[code]; if room!=nil { room.Players[pid]=&Player{ID:pid,Name:"Player "+pid,Conn:c} }; h.mu.Unlock(); if room!=nil { h.broadcast(room,"room.state",room) }
	case "match.start":
		code := m["code"].(string); h.mu.Lock(); room:=h.rooms[code]; h.mu.Unlock(); if room==nil { return }
		round, _ := h.rounds.Random(ctx); room.Round = round; room.EndsAt = time.Now().Add(60*time.Second)
		h.broadcast(room,"round.started", map[string]any{"round": round, "endsAt": room.EndsAt})
		go h.timer(code, room.EndsAt)
	case "guess.submit":
		code := m["code"].(string); h.mu.Lock(); room:=h.rooms[code]; if room==nil { h.mu.Unlock(); return }
		p:=room.Players[pid]; if p==nil || p.Guess.Submitted || time.Now().After(room.EndsAt) { h.mu.Unlock(); return }
		lat:=m["lat"].(float64); lon:=m["lon"].(float64); year:=int(m["year"].(float64));
		if lat < -90 || lat > 90 || lon < -180 || lon > 180 || year < room.Round.MinYear || year > room.Round.MaxYear { h.mu.Unlock(); return }
		p.Guess=PlayerGuess{Lat:lat,Lon:lon,Year:year,Submitted:true}
		h.mu.Unlock(); h.broadcast(room,"guess.accepted", map[string]any{"playerId":pid})
	}
}

func (h *Hub) timer(code string, end time.Time) { <-time.After(time.Until(end)); h.endRound(code) }
func (h *Hub) endRound(code string) {
	h.mu.Lock(); room:=h.rooms[code]; if room==nil { h.mu.Unlock(); return }
	results := map[string]int{}
	for id,p := range room.Players { d:=scoring.HaversineKm(p.Guess.Lat,p.Guess.Lon,room.Round.CorrectLatitude,room.Round.CorrectLongitude); y:=float64(abs(p.Guess.Year-room.Round.CorrectYear)); results[id]=scoring.LocationScore(d,3000)+scoring.TimeScore(y,2000) }
	h.mu.Unlock(); h.broadcast(room,"round.ended",map[string]any{"results":results,"correct":room.Round})
}
func (h *Hub) broadcast(room *Room, t string, payload any) { b,_:=json.Marshal(map[string]any{"type":t,"payload":payload}); for _,p:= range room.Players { _=p.Conn.WriteMessage(gws.TextMessage,b) } }
func (h *Hub) disconnect(pid string) { h.mu.Lock(); defer h.mu.Unlock(); for _,room:= range h.rooms { delete(room.Players,pid) } }
func abs(i int) int { if i < 0 { return -i }; return i }
