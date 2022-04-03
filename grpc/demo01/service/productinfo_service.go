package main

import (
	"context"
	pb "productinfo/service/encommerce"
)

type server struct {
	productMap map[string]pb.Product
}

func (s *server) AddProduct(ctx context.Context, in *pb.Product) (*pb.ProductID, error) {

}

func (s *server) GetProduct(ctx context.Context, in *pb.ProductID) (*pb.Product, error) {

}

func main() {

}
