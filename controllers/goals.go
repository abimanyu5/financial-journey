package controllers

import (
	"financial-journey/database"
	"financial-journey/helper"
	"financial-journey/structs"
	"net/http"
	"strconv"

	"financial-journey/repository"

	"github.com/gin-gonic/gin"
)

func GetAllGoals(c *gin.Context) {
	id := c.GetUint("id")

	masters, err := repository.GetAllGoals(database.DbConnection, id)

	if err != nil {
		helper.RespondWithError(c, http.StatusInternalServerError, "Failed to fetch Goals", err)
		return
	}

	helper.RespondWithSuccess(c, http.StatusOK, "Goals retrieved successfully", masters)
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
	helper.RespondWithSuccess(c, http.StatusOK, "success insert goals data", nil)

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
	helper.RespondWithSuccess(c, http.StatusOK, "success update goals data", nil)
}

func DeleteGoals(c *gin.Context) {
	var goals structs.Goals
	id, _ := strconv.Atoi(c.Param("id"))
	goals.ID = int(id)
	
	err := repository.DeleteGoals(database.DbConnection, goals)
	if err != nil {
		panic(err)
	}
	helper.RespondWithSuccess(c, http.StatusOK, "success delete goals data", nil)
}

func GetGoalsById(c *gin.Context) {
	var goals structs.Goals
	id, _ := strconv.Atoi(c.Param("id"))

	goals.ID = int(id)

	master, err := repository.GetGoalsById(database.DbConnection, goals)

	if err != nil {
		helper.RespondWithError(c, http.StatusInternalServerError, "Failed to fetch goals by id", err)
		return
	}

	helper.RespondWithSuccess(c, http.StatusOK, "Get goals by id successfully", master)
}
