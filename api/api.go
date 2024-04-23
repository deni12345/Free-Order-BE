package api

import "github/lambda-microservice/internal/logic"

type Server struct {
	logic logic.Logic
}

func NewServer() Server {
	l := logic.NewLogicImpl()
	return Server{
		logic: l,
	}
}
