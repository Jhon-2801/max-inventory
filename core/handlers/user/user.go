package handlers

import (
	"net/http"

	"github.com/Jhon-2801/max-inventory/core/user"
	"github.com/gin-gonic/gin"
)

type (
	Controller func(c *gin.Context)
	Endpoints  struct {
		RegisterUser Controller
		LoginUser    Controller
	}
	RegisterReq struct {
		Email    string `form:"email"`
		Name     string `form:"name"`
		Password string `form:"password"`
	}
	LoginReq struct {
		Email    string `form:"email"`
		Password string `form:"password"`
	}
	UserRes struct {
		Id    int    `form:"id"`
		Email string `form:"email"`
		Name  string `form:"name"`
	}
)

func MakeEndPoints(s user.Service) Endpoints {
	return Endpoints{
		RegisterUser: makeRegisterUser(s),
		LoginUser:    makeLoginUser(s),
	}
}

func makeRegisterUser(s user.Service) Controller {
	return func(c *gin.Context) {
		var req RegisterReq
		err := c.ShouldBind(&req)

		if err != nil {
			c.IndentedJSON(http.StatusBadRequest, gin.H{"status": 400, "message": err})
			return
		}
		if req.Email == "" {
			c.IndentedJSON(http.StatusBadRequest, gin.H{"status": 400, "message": "email is requerid"})
			return
		}
		if !s.IsValidMail(req.Email) {
			c.IndentedJSON(http.StatusBadRequest, gin.H{"status": 400, "message": "email is not valid"})
			return
		}
		_, err = s.GetUserByMail(req.Email)
		if err == nil {
			c.IndentedJSON(http.StatusBadRequest, gin.H{"status": 400, "message": "the email already exists"})
			return
		}
		if req.Name == "" {
			c.IndentedJSON(http.StatusBadRequest, gin.H{"status": 400, "message": "name is requerid"})
			return
		}
		if len(req.Password) < 8 {
			c.IndentedJSON(http.StatusBadRequest, gin.H{"status": 400, "message": "the password must be greater than 7 characters"})
			return
		}

		err = s.Register(req.Name, req.Email, req.Password)

		if err != nil {
			c.IndentedJSON(http.StatusInternalServerError, gin.H{"status": 500, "message": err})
			return
		}

		user, _ := s.GetUserByMail(req.Email)
		data := UserRes{
			Id:    user.Id,
			Name:  user.Name,
			Email: user.Email,
		}

		c.IndentedJSON(http.StatusCreated, gin.H{"status": 201, "user": data})
	}
}

func makeLoginUser(s user.Service) Controller {
	return func(c *gin.Context) {
		var req LoginReq

		err := c.ShouldBind(&req)
		if err != nil {
			c.IndentedJSON(http.StatusBadRequest, gin.H{"status": 400, "message": err})
			return
		}

		if req.Email == "" {
			c.IndentedJSON(http.StatusBadRequest, gin.H{"status": 400, "message": "email is requerid"})
		}

		if !s.IsValidMail(req.Email) {
			c.IndentedJSON(http.StatusBadRequest, gin.H{"status": 400, "message": "email is not valid"})
			return
		}
		user, err := s.GetUserByMail(req.Email)

		if err != nil {
			c.IndentedJSON(http.StatusBadRequest, gin.H{"status": 400, "message": "the email does not exist"})
			return
		}

		if len(req.Password) < 8 {
			c.IndentedJSON(http.StatusBadRequest, gin.H{"status": 400, "message": "the password must be greater than 7 characters"})
			return
		}

		err = s.ValidPassword(req.Email, req.Password)

		if err != nil {
			c.IndentedJSON(http.StatusBadRequest, gin.H{"status": 400, "message": "invalid password"})
			return
		}

		data := UserRes{
			Id:    user.Id,
			Name:  user.Name,
			Email: user.Email,
		}

		c.IndentedJSON(http.StatusAccepted, gin.H{"status": 200, "user": data})
	}
}
