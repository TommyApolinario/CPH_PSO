package main

import (
	"CPH_PSO/models"
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"
)

var (
	totalHubs       = 0
	quantityServers = 0
	capacityServers = 0

	Particles []models.Particle
)

func main() {
	iteraciones := 100000
	// Leer el archivo
	file, err := os.Open("data/phub_50_5_1.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	// Leemos linea a linea
	for scanner.Scan() {
		line := scanner.Text()
		if len(line) == 0 {
			continue
		}
		// Convertimos la linea a un arreglo
		nums, err := ConvertToArray(line)
		if err != nil {
			panic(err)
		}

		if len(nums) == 3 {
			totalHubs = nums[0]
			quantityServers = nums[1]
			capacityServers = nums[2]
		} else {

			// Creamos la particula
			hub := models.NewParticle(nums[0], models.Position{X: nums[1], Y: nums[2]}, nums[3])
			// Agregamos la particula
			Particles = append(Particles, hub)

		}

	}
	//el algoritmo paso no hay pensarlo como una libreria, si no como un algoritmo conceptual el cual te implementara particularmente a un problema,
	//en este caso el problema de p-hub que es encontrar una solución optima para las conexiones entre clientes y servidores
	particlecopy := make([]models.Particle, len(Particles))
	copy(particlecopy, Particles)
	bestSolution := getOneSolution(particlecopy, quantityServers, capacityServers)

	for i := 0; i < iteraciones; i++ {
		particlecopy = make([]models.Particle, len(Particles))
		copy(particlecopy, Particles)

		solution := getOneSolution(particlecopy, quantityServers, capacityServers)
		fmt.Println(solution.Fitness)
		// if f(pi) < f(g) then pi es una solucion - g es otra
		// f es la que calcula el fitness o que tan eficiente es
		if solution.Fitness < bestSolution.Fitness {
			bestSolution = solution
		}
	}

	fmt.Println("Best solution: ", bestSolution.Fitness)
}

// convierte un string en un array de enteros
// se lo usa para convertir para los datos del archivo a un array de enteros para hacer los calculos
func ConvertToArray(s string) ([]int, error) {
	s = strings.TrimSpace(s)
	numbersString := strings.Split(s, " ")
	var nums []int
	for _, num := range numbersString {
		n, err := strconv.Atoi(num)
		if err != nil {
			return nums, err
		}
		nums = append(nums, n)
	}
	return nums, nil
}

// para obtener una solucion
// esta funcion genera una unica solucion, para que el algoritmo pso tome esta solucion y las compara
func getOneSolution(particles []models.Particle, quantityServer int, capacityServer int) models.Solution {
	// Generamos servidores aleatorios
	servers := randomServers(particles, quantityServer, capacityServer)

	// Agregamos clientes a servidores aleatorios
	for len(particles) > 0 {
		rndIndexServer := rand.Intn(len(servers))
		rndIndexClient := rand.Intn(len(particles))

		server := servers[rndIndexServer]
		client := particles[rndIndexClient].ToClient()
		if server.AddClient(client) {
			particles = append(particles[:rndIndexClient], particles[rndIndexClient+1:]...)
		}
	}

	var solution models.Solution
	for _, server := range servers {
		solution.Fitness += server.TotalDistance
	}

	solution.Assignments = servers

	return solution
}

// genera servidores randoms
func randomServers(hubs []models.Particle, quantityServer int, capacityServer int) []*models.Server {
	rand.Seed(time.Now().UnixNano())

	// Crear un mapa para evitar servidores duplicados
	usedHubs := make(map[int]bool)

	// Lista de servidores generados aleatoriamente
	servers := make([]*models.Server, 0)

	// Mientras haya servidores que necesitemos generar
	for i := 0; i < quantityServer; {
		// Elegir un hub aleatorio
		randomIndex := rand.Intn(len(hubs))
		hub := hubs[randomIndex]

		// Verificar si el hub ya se ha utilizado como servidor
		if !usedHubs[hub.NodeNumber] {
			// Convertir el hub en un servidor con la capacidad especificada
			server := hub.ToServer(capacityServer)
			// Agregar el servidor a la lista de servidores
			servers = append(servers, server)
			// Marcar el hub como utilizado
			usedHubs[hub.NodeNumber] = true
			// Incrementar el contador de servidores generados
			i++
		}
	}

	return servers
}
