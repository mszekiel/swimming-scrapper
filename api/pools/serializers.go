package pools

import "github.com/gin-gonic/gin"

type LocationResponse struct {
	Lat int `json:"lat"`
	Lon int `json:"lon"`
}

type LocationSerializer struct {
	C *gin.Context
	Pool
}

func (s *LocationSerializer) Response() LocationResponse {
	response := LocationResponse{}

	response = LocationResponse{
		Lat: int(s.Location.Lat),
		Lon: int(s.Location.Lon),
	}

	return response
}

type PoolResponse struct {
	ID       uint             `json:"-"`
	Name     string           `json:"name"`
	Location LocationResponse `json:"location"`
}

type PoolSerializer struct {
	C    *gin.Context
	Pool Pool
}

func (s *PoolSerializer) Response() PoolResponse {
	locationSerializer := LocationSerializer{s.C, s.Pool}

	response := PoolResponse{
		ID:       s.Pool.ID,
		Name:     s.Pool.Name,
		Location: locationSerializer.Response(),
	}

	return response
}

type PoolsSerializer struct {
	C     *gin.Context
	Pools []Pool
}

func (s *PoolsSerializer) Response() []PoolResponse {
	response := []PoolResponse{}

	for _, pool := range s.Pools {
		locationSerializer := LocationSerializer{s.C, pool}

		response = append(response, PoolResponse{
			ID:       pool.ID,
			Name:     pool.Name,
			Location: locationSerializer.Response(),
		})
	}

	return response
}
