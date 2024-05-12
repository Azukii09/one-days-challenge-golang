package day_8

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSuccess(t *testing.T) {

	firstFlights := []Flight{
		Flight{Origin: "GLA", Destination: "CDG", Price: 1000},
		Flight{Origin: "GLA", Destination: "JFK", Price: 5000},
		Flight{Origin: "GLA", Destination: "SNG", Price: 3000},
	}
	secondFlights := []Flight{
		Flight{Origin: "GLA", Destination: "CDG", Price: 1000},
		Flight{Origin: "GLA", Destination: "JFK", Price: 5000},
		Flight{Origin: "GLA", Destination: "SNG", Price: 3000},
		Flight{Origin: "GLA", Destination: "AMS", Price: 500},
	}
	input := IsSubset(firstFlights, secondFlights)
	assert.Equal(t, true, input)
}

func TestFailed(t *testing.T) {

	firstFlights := []Flight{
		Flight{Origin: "GLA", Destination: "ABC", Price: 1000},
		Flight{Origin: "GLA", Destination: "DEF", Price: 5000},
		Flight{Origin: "GLA", Destination: "GHI", Price: 3000},
	}
	secondFlights := []Flight{
		Flight{Origin: "GLA", Destination: "CDG", Price: 1000},
		Flight{Origin: "GLA", Destination: "JFK", Price: 5000},
		Flight{Origin: "GLA", Destination: "SNG", Price: 3000},
		Flight{Origin: "GLA", Destination: "AMS", Price: 500},
	}
	input := IsSubset(firstFlights, secondFlights)
	assert.Equal(t, false, input)
}
