package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"regexp"
	"strings"

	"google.golang.org/grpc"
	pb "github.com/kaushik592k/equations/protobuf"

)

type server struct {
	pb.UnimplementedChemicalBalancerServer
}

func parseCompound(compound string) map[string]int {
	re := regexp.MustCompile(`([A-Z][a-z]*)(\d*)`)
	matches := re.FindAllStringSubmatch(compound, -1)
	elements := make(map[string]int)
	for _, match := range matches {
		element := match[1]
		count := 1
		if match[2] != "" {
			fmt.Sscanf(match[2], "%d", &count)
		}
		elements[element] += count
	}
	return elements
}

func parseSide(side string) []map[string]int {
	compounds := strings.Split(side, "+")
	var parsed []map[string]int
	for _, compound := range compounds {
		parsed = append(parsed, parseCompound(compound))
	}
	return parsed
}

func balanceEquation(equation string) (string, error) {
	parts := strings.Split(equation, "=")
	if len(parts) != 2 {
		return "", fmt.Errorf("invalid equation format")
	}
	reactants := parseSide(parts[0])
	products := parseSide(parts[1])

	// Collect unique elements
	elementSet := make(map[string]bool)
	for _, comp := range append(reactants, products...) {
		for elem := range comp {
			elementSet[elem] = true
		}
	}
	var elements []string
	for elem := range elementSet {
		elements = append(elements, elem)
	}

	// Create matrix and solve system
	// Placeholder: Implement a balancing algorithm (details omitted for brevity)
	// Example result (mocked):
	balanced := "2H2 + O2 = 2H2O"

	return balanced, nil
}

func (s *server) BalanceEquation(ctx context.Context, req *pb.BalanceRequest) (*pb.BalanceResponse, error) {
	equation := req.GetEquation()
	balanced, err := balanceEquation(equation)
	if err != nil {
		return &pb.BalanceResponse{
			BalancedEquation: "",
			ErrorMessage:     err.Error(),
		}, nil
	}
	return &pb.BalanceResponse{
		BalancedEquation: balanced,
		ErrorMessage:     "",
	}, nil
}

func main() {
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	grpcServer := grpc.NewServer()
	pb.RegisterChemicalBalancerServer(grpcServer, &server{})
	log.Println("Server is running on port 50051")
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
