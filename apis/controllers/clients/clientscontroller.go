package clientscontroller

import (
	"fmt"

	"github.com/gin-gonic/gin"
	cu "github.com/ravayak/copee_backend/apis/controllers/clients/utility"
	"github.com/ravayak/copee_backend/apis/domain/clients"
	data "github.com/ravayak/copee_backend/apis/domain/clients/data"
	cs "github.com/ravayak/copee_backend/apis/services/clients"
	gu "github.com/ravayak/copee_backend/app/utils/gin"
	"github.com/ravayak/copee_backend/datasources/mysql"
)

func Get(c *gin.Context) {
	clientID, exists := c.Get("objectId")
	cId := clientID.(int32)
	if exists {
		client, err := cs.ClientsService.GetClient(cId)
		if err != nil {
			c.JSON(400, gin.H{"error": err.Error()})
			return
		}
		c.JSON(200, client)
		return
	}
	c.JSON(404, gin.H{"error": fmt.Sprintf("error Client %d not found ", clientID)})
}

func Delete(c *gin.Context) {
	clientId, exists := c.Get("objectId")
	if exists {
		cId := clientId.(int32)
		cs.ClientsService.DeleteClient(cId)
		c.JSON(200, fmt.Sprintf("client %d deleted ", cId))
		return
	}
	c.JSON(404, gin.H{"error": fmt.Sprintf("error Client %d not found", clientId)})
}

func Update(c *gin.Context) {
	clientId, exists := c.Get("objectId")
	cId := clientId.(int32)
	if exists {
		updatedClient := clients.Client{ClientID: cId}
		clientParams, err := cu.SetClientParams(data.UPDATE_CLIENT_PARAMS, updatedClient, c)
		if err != nil {
			c.JSON(400, gin.H{"error": err.Error()})
			return
		}
		err = cs.ClientsService.UpdateClient(clientParams.(mysql.UpdateClientParams))
		if err != nil {
			c.JSON(400, gin.H{"error": err.Error()})
			return
		}

		c.JSON(200, fmt.Sprintf("client %d updated ", cId))
		return
	}
	c.JSON(404, gin.H{"error": fmt.Sprintf("error Client %d not found ", clientId)})
}

func Create(c *gin.Context) {
	userLogin, err := gu.GetUserLoginFromContext(c)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	newClient := clients.Client{}
	clientParams, err := cu.SetClientParams(data.CREATE_CLIENT_PARAMS, newClient, c)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	client, err := cs.ClientsService.InsertClient(clientParams.(mysql.CreateClientParams), userLogin.User.UserID)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	c.JSON(201, client)
}

func List(c *gin.Context) {
	userLogin, err := gu.GetUserLoginFromContext(c)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	clients, err := cs.ClientsService.ListClientsByUser(userLogin.User.UserID)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, clients)
}
