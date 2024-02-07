package handlers

import (
	"net/http"

	"github.com/Jhon-2801/max-inventory/core/roles"
	"github.com/gin-gonic/gin"
)

type (
	Controller func(c *gin.Context)
	Endpoints  struct {
		SaveUserRole   Controller
		RemoveUserRole Controller
	}

	UserRolReq struct {
		UserID string `form:"user_id"`
		RoleID string `form:"role_id"`
	}
)

func MakeEndPoints(s roles.Service) Endpoints {
	return Endpoints{
		SaveUserRole:   makeSaveUserRole(s),
		RemoveUserRole: makeRemoveUserRole(s),
	}
}

func makeSaveUserRole(s roles.Service) Controller {
	return func(c *gin.Context) {
		var req UserRolReq
		err := c.ShouldBind(&req)

		if err != nil {
			c.IndentedJSON(http.StatusBadRequest, gin.H{"status": 400, "message": err})
			return
		}
		if len(req.UserID) <= 0 {
			c.IndentedJSON(http.StatusBadRequest, gin.H{"status": 400, "message": "user_id is required"})
			return
		}
		err = s.UserByID(req.UserID)
		if err != nil {
			c.IndentedJSON(http.StatusBadRequest, gin.H{"status": 400, "message": "user_id not found"})
			return
		}
		if len(req.RoleID) <= 0 {
			c.IndentedJSON(http.StatusBadRequest, gin.H{"status": 400, "message": "role_id is required"})
			return
		}
		if req.RoleID >= "3" || req.RoleID == "0" {
			c.IndentedJSON(http.StatusBadRequest, gin.H{"status": 400, "message": "role_id not found"})
			return
		}
		userRole, err := s.SaveUserRole(req.UserID, req.RoleID)

		if err != nil {
			c.IndentedJSON(http.StatusInternalServerError, gin.H{"status": 500, "message": err})
			return
		}

		c.IndentedJSON(http.StatusCreated, gin.H{"status": 201, "data": userRole})
	}
}

func makeRemoveUserRole(s roles.Service) Controller {
	return func(c *gin.Context) {

	}
}
