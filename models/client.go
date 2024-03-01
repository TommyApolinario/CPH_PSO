package models

import "math"

// Client un cliente de una red
type Client struct {
	Particle
	Distance float64
}

// calcular distancia al servidor
func (client *Client) CalculateDistance(server *Server) float64 {
	client.Distance = math.Sqrt(math.Pow(float64(server.Particle.Position.X-client.Particle.Position.X), 2) + math.Pow(float64(server.Particle.Position.Y-client.Particle.Position.Y), 2))
	return client.Distance
}
