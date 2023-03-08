package gfpfuncs

import (
	"fmt"
	templates "gfpservice/merge-proto"
	"log"

	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"
)

// Service is service
type Service struct {
	Name string
	MergeClient templates.TemplateMakerClient
}

// NewService is new service
func NewService() (Service, error) {
	svc := Service {
		Name: "gfp_service",
	}

	mergeConn, err := grpc.Dial("thirdgrpcservice-project2:50105", grpc.WithInsecure())
	if err != nil {
		fmt.Println("generate grpc connect failed")
	}
	defer mergeConn.Close()

	svc.MergeClient = templates.NewTemplateMakerClient(mergeConn)

	gateway := gin.Default()
	gateway.GET("/gfp/:a/:b", svc.GuiFunc)

	if err := gateway.Run(":50098"); err != nil {
		log.Fatalf("Failed to run client service: %v", err)
	}
	
	return svc, nil
}