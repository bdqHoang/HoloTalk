package handlers

import (
	"api-gateway/config"
	"api-gateway/dto"
	"api-gateway/proto"

	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"
)

var env = config.LoadEnv()
func RegisterHandler(c *gin.Context) {
	conn, err := grpc.Dial("localhost:"+ env.PORT_AUTH, grpc.WithInsecure())
	if err != nil {
		c.JSON(500, gin.H{
			"message": "Error connecting to auth service",
		})
		return
	}
	defer conn.Close()

	client := proto.NewAuthServiceClient(conn)
	var req dto.RegisterRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(400, gin.H{
			"message": "Invalid request",
		})
		return
	}

	res, err := client.Register(c, &proto.RegisterRequest{
		Email:           req.Email,
		Phone:           req.Phone,
		Name:            req.Name,
		DateOfBirth:     req.DateOfBirth,
		Password:        req.Password,
		ConfirmPassword: req.ConfirmPassword,
		Role:            uint64(req.Role),
	})
	if err != nil {
		c.JSON(400, gin.H{
			"message": "Error registering user",
		})
		return
	}

	c.JSON(200, res.Message)
}

func LoginHandler(c *gin.Context) {
	conn, err := grpc.Dial("localhost:"+ env.PORT_AUTH, grpc.WithInsecure())
	if err != nil {
		c.JSON(500, gin.H{
			"message": "Error connecting to auth service",
		})
		return
	}
	defer conn.Close()

	client := proto.NewAuthServiceClient(conn)
	var req dto.LoginRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(400, gin.H{
			"message": "Invalid request",
		})
		return
	}

	res, err := client.Login(c, &proto.LoginRequest{		
		Username: req.Username,
		Password: req.Password,		
	})
	if err != nil {
		c.JSON(400, gin.H{
			"message": "Error logging in user",
		})
		return
	}

	c.JSON(200, res)
}

func RefreshTokenHandler(c *gin.Context) {
	conn, err := grpc.Dial("localhost:"+ env.PORT_AUTH, grpc.WithInsecure())
	if err != nil {
		c.JSON(500, gin.H{
			"message": "Error connecting to auth service",
		})
		return
	}
	defer conn.Close()

	client := proto.NewAuthServiceClient(conn)
	var req dto.RefreshTokenRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(400, gin.H{
			"message": "Invalid request",
		})
		return
	}

	res, err := client.RefreshToken(c, &proto.RefreshTokenRequest{
		RefreshToken: req.RefreshToken,
	})
	if err != nil {
		c.JSON(400, gin.H{
			"message": "Error refreshing token",
		})
		return
	}

	c.JSON(200, res)	
}