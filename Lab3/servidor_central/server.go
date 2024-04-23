package servidor_central

import (
	"context"
	"fmt"
	"sync"
)

// server es nuestra implementación del servicio HelloWorldServer
type Server struct {
	UnimplementedServidorCentralServer
	AT    int
	MP    int
	Mutex sync.Mutex
}

func (s *Server) ActualizarMunicion(at int, mp int) {

	s.Mutex.Lock()
	defer s.Mutex.Unlock()
	at_max := 50
	mp_max := 20
	if s.AT+at >= at_max {
		s.AT = at_max
	} else {
		s.AT += at
	}
	if s.MP+mp >= mp_max {
		s.MP = mp_max
	} else {
		s.MP += mp
	}
}

func (s *Server) SolicitudMunicion(ctx context.Context, p *PedirMunicion) (*RespuestaMunicion, error) {

	fmt.Print("Recepción de solicitud desde equipo", p.Id, " con AT:", p.At, " y MP:", p.Mp)
	if (s.AT >= int(p.At)) && (s.MP >= int(p.Mp)) {
		s.Mutex.Lock()
		s.AT -= int(p.At)
		s.MP -= int(p.Mp)
		defer s.Mutex.Unlock()
		fmt.Print(" -- APROBADA -- ")
		fmt.Println("AT EN SISTEMA:", s.AT, "MP EN SISTEMA:", s.MP)
		return &RespuestaMunicion{Respuesta: 1}, nil
	} else {
		fmt.Print(" -- DENEGADA -- ")
		fmt.Println("AT EN SISTEMA:", s.AT, "MP EN SISTEMA:", s.MP)
		return &RespuestaMunicion{Respuesta: 0}, nil
	}
}
