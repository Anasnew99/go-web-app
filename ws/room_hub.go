package ws

import "time"

type RoomHub struct {
	roomHubMap map[string]*Hub
	stopChans  map[string]chan bool
}

func NewRoomHub() *RoomHub {
	return &RoomHub{
		roomHubMap: make(map[string]*Hub),
		stopChans:  make(map[string]chan bool),
	}
}

func (r *RoomHub) CreateRoomHub(roomId string) *Hub {
	// if roomId does not exist in roomHubMap, then
	// create a new hub and run it
	if _, ok := r.roomHubMap[roomId]; !ok {
		stop := make(chan bool)
		r.stopChans[roomId] = stop
		r.roomHubMap[roomId] = NewHub(stop)
		go r.roomHubMap[roomId].Run()
	}
	return r.roomHubMap[roomId]
}

func (r *RoomHub) deleteRoomHub(roomId string) {
	if r.roomHubMap == nil || r.roomHubMap[roomId] == nil {
		return
	}
	r.stopChans[roomId] <- true
	close(r.stopChans[roomId])
	delete(r.stopChans, roomId)
	delete(r.roomHubMap, roomId)
}

var inactiveRooms = make(map[string]bool)

func (r *RoomHub) deleteRoomAfterInactivity(roomId string) {
	if inactiveRooms[roomId] {
		inactiveRooms[roomId] = true
		time.Sleep(5 * time.Minute)
		if len(r.roomHubMap[roomId].clients) == 0 {
			r.deleteRoomHub(roomId)
		}
	}

}

func (r *RoomHub) Run() {
	for {
		for roomId, hub := range r.roomHubMap {
			if len(hub.clients) == 0 {
				go r.deleteRoomAfterInactivity(roomId)
			}
		}
	}
}
