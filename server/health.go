package main

type Health struct {
	Status string `json:"status"`
}

func NewHealth() Health {
	return Health{
		Status: "OK",
	}
}

func checkHealth() Health {
	return NewHealth()
}
