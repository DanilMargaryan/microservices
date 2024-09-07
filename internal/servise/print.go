package servise

import (
	"fmt"
	"github.com/DanilMargaryan/microservices/internal/config"
)

func PrintConfig(s *config.Config) {
	fmt.Println(s)
}
