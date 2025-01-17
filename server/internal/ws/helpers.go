package ws

import "fmt"

func listRooms(h *Hub) {
	fmt.Println()
	for _, v := range h.rooms {
		fmt.Printf("Room: '%s' ", v.id)
	}
	fmt.Println()
}
