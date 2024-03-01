package models

// Server un servidor de una red
type Server struct {
	Particle            *Particle
	TotalDistance       float64
	TotalClients        int
	AccumulatedCapacity int
	Capacity            int
	Clients             []Client
}

// esta funcion agrega un cliente a un servidor
func (server *Server) AddClient(client Client) bool {
	if server.AccumulatedCapacity+client.Demand <= server.Capacity {
		server.Clients = append(server.Clients, client)
		server.AccumulatedCapacity += client.Demand
		server.TotalDistance += client.CalculateDistance(server)
		server.TotalClients++
		return true
	}
	return false
}
