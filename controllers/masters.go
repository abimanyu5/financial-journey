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

func GetAllMasters(c *gin.Context) {
	id := c.GetUint("id")

	masters, err := repository.GetAllMasters(database.DbConnection, id)

	if err != nil {
		helper.RespondWithError(c, http.StatusInternalServerError, "Failed to fetch masters", err)
		return
	}

	helper.RespondWithSuccess(c, http.StatusOK, "masters retrieved successfully", masters)
}

func InsertMaster(c *gin.Context) {
	var masters structs.Masters

	err := c.ShouldBindJSON(&masters)
	if err != nil {
		
		panic(err)
	}
	err = repository.InsertMaster(database.DbConnection, masters)
	if err != nil {
		helper.RespondWithError(c, http.StatusInternalServerError, "Failed to insert master data", err)
	}
	helper.RespondWithSuccess(c, http.StatusOK, "success insert master data", masters)


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
		helper.RespondWithError(c, http.StatusInternalServerError, "Failed to update master data", err)
	}
	helper.RespondWithSuccess(c, http.StatusOK, "success update master data", masters)
}

func DeleteMaster(c *gin.Context) {
	var masters structs.Masters
	id, _ := strconv.Atoi(c.Param("id"))
	masters.ID = int(id)
	
	err := repository.DeleteMaster(database.DbConnection, masters)
	if err != nil {
		helper.RespondWithError(c, http.StatusInternalServerError, "Failed to update master data", err)
	}
	helper.RespondWithSuccess(c, http.StatusOK, "success delete master data", masters)
}

func GetMasterById(c *gin.Context) {
	var masters structs.Masters
	id, _ := strconv.Atoi(c.Param("id"))

	masters.ID = int(id)

	master, err := repository.GetMasterById(database.DbConnection, masters)

	if err != nil {
		helper.RespondWithError(c, http.StatusInternalServerError, "Failed to fetch masters", err)
		return
	}

	helper.RespondWithSuccess(c, http.StatusOK, "masters retrieved by id successfully", master)
}
