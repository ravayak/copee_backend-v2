package groupsuserscontroller

import (
	"fmt"

	"github.com/gin-gonic/gin"
	guu "github.com/ravayak/copee_backend/apis/controllers/groups/users/utility"
	gul "github.com/ravayak/copee_backend/apis/domain/groups/users/lnk"
	lnk "github.com/ravayak/copee_backend/apis/domain/groups/users/lnk"
	gus "github.com/ravayak/copee_backend/apis/services/groups/users"
)

func Get(c *gin.Context) {
	guIdParam, exists := c.Get("objectId")
	guId := guIdParam.(int32)
	if exists {
		gu, err := gus.GroupsUsersService.GetGroupUser(guId)
		if err != nil {
			c.JSON(400, gin.H{"error": err.Error()})
			return
		}
		c.JSON(200, gu)
		return
	}
	c.JSON(404, gin.H{"error": fmt.Sprintf("error group user %d not found ", guId)})
}

func List(c *gin.Context) {
	clients, err := gus.GroupsUsersService.ListGroupUsers()
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, clients)
}

func Create(c *gin.Context) {
	createGroupUserLnkParams := &lnk.CreateGroupUserLnkParams{}
	if err := c.ShouldBindJSON(createGroupUserLnkParams); err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	createGroupUserLnkParamsMysql := guu.SetParamsToMysqlGroupUserParams(createGroupUserLnkParams)
	var groupUser *gul.GroupsUsersLnk
	var err error

	if createGroupUserLnkParams.GuGroupID != 0 && createGroupUserLnkParams.GName != "" {
		c.JSON(400, gin.H{"error": "Both group id and group name cannot be provided"})
		return
	}
	if createGroupUserLnkParams.GuGroupID != 0 {
		groupUser, err = gus.GroupsUsersService.InsertGroupUser(createGroupUserLnkParamsMysql)
	} else if createGroupUserLnkParams.GName != "" {
		groupUser, err = gus.GroupsUsersService.InsertGroupUserByGroupName(createGroupUserLnkParamsMysql.GuUserID.Int32, createGroupUserLnkParams.GName)
	} else {
		c.JSON(400, gin.H{"error": "Either group id or group name must be provided"})
		return
	}
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, groupUser)
}

// func Update(c *gin.Context) {
// 	guIdParam, exists := c.Get("objectId")
// 	guId := guIdParam.(int32)
// 	if exists {
// 		updatedGu := &lnk.CreateGroupUserLnkParams{}
// 		guu.SetParamsToMysqlGroupUserParams()
// 		err = gus.GroupsUsersService.UpdateGroupUser(groupUserParams)
// 		if err != nil {
// 			c.JSON(400, gin.H{"error": err.Error()})
// 			return
// 		}

// 		c.JSON(200, fmt.Sprintf("client %d updated ", cId))
// 		return
// 	}
// 	c.JSON(404, gin.H{"error": fmt.Sprintf("error Client %d not found ", clientId)})
// }
