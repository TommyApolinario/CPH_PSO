package models

type Particle struct {
	NodeNumber int
	Position   Position
	Demand     int
}

type Position struct {
	X int
	Y int
}

// es un constructor, toma los datos del archivo y le pasa el numero del nodo, luego los dos siguientes valores, es la posicion de x y y la demanda
// NewParticle returns a new Particle
func NewParticle(nodeNumber int, position Position, demand int) Particle {
	return Particle{
		NodeNumber: nodeNumber,
		Position:   position,
		Demand:     demand,
	}
}

// es la conversion de servidor a cliente
// es una herencia, de una particula hereda la posicion y el numero de nodos tanto para el cliente como para el servidor
func (particle *Particle) ToServer(capacityServer int) *Server {
	return &Server{
		Particle:            particle,
		TotalDistance:       0,
		TotalClients:        0,
		AccumulatedCapacity: 0,
		Capacity:            capacityServer,
		Clients:             nil,
	}
}

// convierte una particula a un cliente
func (particle Particle) ToClient() Client {
	return Client{
		Particle: particle,
		Distance: 0,
	}
}
