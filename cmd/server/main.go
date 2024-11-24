package main

import (
	"github.com/DoCongThanhPhuong/go-backend/internal/routers"
)

func main() {
  r := routers.NewRouter()
  r.Run()
}

