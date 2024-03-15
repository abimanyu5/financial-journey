package controllers

import (
	"net/http"
	"financial-journey/database"
	"financial-journey/structs"
	"strconv"

	"financial-journey/repository"

	"github.com/gin-gonic/gin"
)

func GetAllMasters(c *gin.Context) {
	id := c.GetUint("id")
	var (
		result gin.H
	) 

	masters, err := repository.GetAllMasters(database.DbConnection, id)

	if err != nil {
		result = gin.H{
			"result":  err,
		}
	} else {
		result = gin.H{
			"result":  masters,
		}
	}

	c.JSON(http.StatusOK, result)
}

func InsertMaster(c *gin.Context) {
	var masters structs.Masters

	err := c.ShouldBindJSON(&masters)
	if err != nil {
		
		panic(err)
	}
	err = repository.InsertMaster(database.DbConnection, masters)
	if err != nil {
		panic(err)
	}
	c.JSON(http.StatusOK, gin.H{
		"result": "success insert master data",
	})


}
func UpdateMaster(c *gin.Context) {
	var masters structs.Masters
	id, _ := strconv.Atoi(c.Param("id"))
	
	err := c.ShouldBindJSON((&masters))
	if err != nil {
		panic(err)
	}

	masters.ID = int(id)

	err = repository.UpdateMaster(database.DbConnection, masters)
	if err != nil {
		panic(err)
	}
	c.JSON(http.StatusOK, gin.H{
		"result": "success update master data",
	})
}

func DeleteMaster(c *gin.Context) {
	var masters structs.Masters
	id, _ := strconv.Atoi(c.Param("id"))
	masters.ID = int(id)
	
	err := repository.DeleteMaster(database.DbConnection, masters)
	if err != nil {
		panic(err)
	}
	c.JSON(http.StatusOK, gin.H{
		"result": "success delete master data",
	})
}

func GetMasterById(c *gin.Context) {
	var masters structs.Masters
	var (
		result gin.H
	)
	
	id, _ := strconv.Atoi(c.Param("id"))

	masters.ID = int(id)

	master, err := repository.GetMasterById(database.DbConnection, masters)

	if err != nil {
		result = gin.H{
			"result":  err,
		}
	} else {
		result = gin.H{
			"result":  master,
		}
	}

	c.JSON(http.StatusOK, result)
}
