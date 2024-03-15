package controllers

import (
	"net/http"
	"financial-journey/database"
	"financial-journey/structs"
	"strconv"

	"financial-journey/repository"

	"github.com/gin-gonic/gin"
)

func GetAllGoals(c *gin.Context) {
	id := c.GetUint("id")
	var (
		result gin.H
	) 

	masters, err := repository.GetAllGoals(database.DbConnection, id)

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

func InsertGoals(c *gin.Context) {
	var goals structs.Goals

	err := c.ShouldBindJSON(&goals)
	if err != nil {
		
		panic(err)
	}
	err = repository.InsertGoals(database.DbConnection, goals)
	if err != nil {
		panic(err)
	}
	c.JSON(http.StatusOK, gin.H{
		"result": "success insert goals data",
	})


}
func UpdateGoals(c *gin.Context) {
	var goals structs.Goals
	id, _ := strconv.Atoi(c.Param("id"))
	
	err := c.ShouldBindJSON((&goals))
	if err != nil {
		panic(err)
	}

	goals.ID = int(id)

	err = repository.UpdateGoals(database.DbConnection, goals)
	if err != nil {
		panic(err)
	}
	c.JSON(http.StatusOK, gin.H{
		"result": "success update master data",
	})
}

func DeleteGoals(c *gin.Context) {
	var goals structs.Goals
	id, _ := strconv.Atoi(c.Param("id"))
	goals.ID = int(id)
	
	err := repository.DeleteGoals(database.DbConnection, goals)
	if err != nil {
		panic(err)
	}
	c.JSON(http.StatusOK, gin.H{
		"result": "success delete master data",
	})
}

func GetGoalsById(c *gin.Context) {
	var goals structs.Goals
	var (
		result gin.H
	)
	
	id, _ := strconv.Atoi(c.Param("id"))

	goals.ID = int(id)

	master, err := repository.GetGoalsById(database.DbConnection, goals)

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
