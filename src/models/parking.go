package models

import (
    "sync"
)

type Parking struct {
    spaces         []bool        
    entranceMutex  *sync.Mutex   
    spacesChannel  chan int      
    totalSpaces    int           
}

func NewParking(totalSpaces int) *Parking {
    return &Parking{
        spaces:        make([]bool, totalSpaces),  
        entranceMutex: &sync.Mutex{},              
        spacesChannel: make(chan int, totalSpaces),
        totalSpaces:   totalSpaces,                
    }
}

// GetSpaces devuelve el canal que controla los espacios ocupados del estacionamiento.
func (p *Parking) GetSpaces() chan int {
    return p.spacesChannel
}

// GetEntrance devuelve el mutex de la entrada, utilizado para sincronización de entrada/salida.
func (p *Parking) GetEntrance() *sync.Mutex {
    return p.entranceMutex
}

// GetSpacesArray devuelve el slice de booleanos que representan los espacios ocupados/vacíos.
func (p *Parking) GetSpacesArray() []bool {
    return p.spaces
}

// FindAvailableSpace devuelve el índice del primer espacio disponible o -1 si no hay espacio.
func (p *Parking) FindAvailableSpace() int {
    for i, occupied := range p.spaces {
        if !occupied {
            return i
        }
    }
    return -1
}

// OccupySpace marca un espacio específico como ocupado.
func (p *Parking) OccupySpace(spaceIndex int) {
    if spaceIndex >= 0 && spaceIndex < len(p.spaces) {
        p.spaces[spaceIndex] = true
    }
}

// FreeSpace marca un espacio específico como libre.
func (p *Parking) FreeSpace(spaceIndex int) {
    if spaceIndex >= 0 && spaceIndex < len(p.spaces) {
        p.spaces[spaceIndex] = false
    }
}
