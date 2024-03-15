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

func GetAllTransactions(c *gin.Context) {
	id := c.GetUint("id")


	transactions, err := repository.GetAllTransactions(database.DbConnection, id)

	if err != nil {
		helper.RespondWithError(c, http.StatusInternalServerError, "Failed to fetch transactions", err)
		return
	}

	helper.RespondWithSuccess(c, http.StatusOK, "Transactions retrieved successfully", transactions)
}

func InsertTransaction(c *gin.Context) {
	var transactions structs.Transactions

	err := c.ShouldBindJSON(&transactions)
	if err != nil {
		
		panic(err)
	}
	if transactions.GoalId != 0 {
		// Periksa apakah tujuan (goals) ditemukan
		goal, err := repository.CheckGoals(database.DbConnection, transactions.GoalId)
		if err != nil {
			helper.RespondWithError(c, http.StatusInternalServerError, "Failed to add transactions", err)
			return
		}
		if goal != nil {
			goal.Amount += transactions.Amount
			err = repository.UpdateGoal(database.DbConnection, goal)
			if err != nil {
				helper.RespondWithError(c, http.StatusInternalServerError, "Failed to add transactions", err)
				return
			}
		}
	}

		// Update jumlah (amount) tujuan dengan jumlah transaksi

	err = repository.InsertTransaction(database.DbConnection, transactions)
	if err != nil {
		panic(err)
	}
	helper.RespondWithSuccess(c, http.StatusOK, "Transactions add successfully", nil)


}
func UpdateTransaction(c *gin.Context) {
	var transactions structs.Transactions
	id, _ := strconv.Atoi(c.Param("id"))
	
	err := c.ShouldBindJSON((&transactions))
	if err != nil {
		panic(err)
	}

	transactions.ID = int(id)

	err = repository.UpdateTransaction(database.DbConnection, transactions)
	if err != nil {
		panic(err)
	}
	helper.RespondWithSuccess(c, http.StatusOK, "success update transaction", nil)
}

func DeleteTransaction(c *gin.Context) {
	var transactions structs.Transactions
	id, _ := strconv.Atoi(c.Param("id"))
	transactions.ID = int(id)
	
	err := repository.DeleteTransaction(database.DbConnection, transactions)
	if err != nil {
		panic(err)
	}
	helper.RespondWithSuccess(c, http.StatusOK, "success delete transaction", nil)
}