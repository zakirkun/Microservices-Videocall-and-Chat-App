package handlers

import (
	"net/http"

	"github.com/Slimo300/Microservices-Videocall-and-Chat-App/backend/lib/apperrors"
	"github.com/Slimo300/Microservices-Videocall-and-Chat-App/backend/lib/events"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func (s *Server) GetUserGroups(c *gin.Context) {
	userID := c.GetString("userID")
	userUID, err := uuid.Parse(userID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"err": "invalid ID"})
	}

	groups, err := s.DB.GetUserGroups(userUID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"err": err.Error()})
		return
	}
	if len(groups) == 0 {
		c.Status(http.StatusNoContent)
		return
	}

	c.JSON(http.StatusOK, groups)

}

func (s *Server) CreateGroup(c *gin.Context) {
	userID := c.GetString("userID")
	userUID, err := uuid.Parse(userID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"err": "invalid ID"})
		return
	}

	payload := struct {
		Name string `json:"name"`
	}{}

	err = c.ShouldBindJSON(&payload)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"err": err.Error()})
		return
	}
	if payload.Name == "" {
		c.JSON(http.StatusBadRequest, gin.H{"err": "name not specified"})
		return
	}

	group, err := s.DB.CreateGroup(userUID, payload.Name)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"err": err.Error()})
		return
	}

	user, err := s.DB.GetUser(userUID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"err": err.Error()})
		return
	}

	if err = s.Emitter.Emit(events.MemberCreatedEvent{
		ID:      group.Members[0].ID,
		GroupID: group.ID,
		UserID:  userUID,
		User: events.User{
			UserName: user.UserName,
			Picture:  user.Picture,
		},
		Creator: true,
	}); err != nil {
		panic(err)
	}

	c.JSON(http.StatusCreated, group)
}

func (s *Server) DeleteGroup(c *gin.Context) {
	userID := c.GetString("userID")
	userUUID, err := uuid.Parse(userID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"err": "invalid ID"})
		return
	}

	groupID := c.Param("groupID")
	groupUUID, err := uuid.Parse(groupID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"err": "invalid group ID"})
		return
	}
	group, err := s.DB.DeleteGroup(userUUID, groupUUID)
	if err != nil {
		c.JSON(apperrors.Status(err), gin.H{"err": err.Error()})
		return
	}
	if group.Picture != "" {
		if err := s.Storage.DeleteFile(group.Picture); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"err": err.Error()})
			return
		}
	}

	if err = s.Emitter.Emit(events.GroupDeletedEvent{
		ID: group.ID,
	}); err != nil {
		panic(err)
	}

	c.JSON(http.StatusOK, gin.H{"message": "group deleted"})

}
