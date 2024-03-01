package models

// esta es la solucion va a tener assignments(asignaciones de servidores a clientes) y se conectan
// fitness es la suma de todas las distancias
type Solution struct {
	Assignments []*Server
	Fitness     float64
}
