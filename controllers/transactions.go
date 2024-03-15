package controllers

import (
	"net/http"
	"financial-journey/database"
	"financial-journey/structs"
	"strconv"

	"financial-journey/repository"
	"github.com/gin-gonic/gin"
)

func GetAllTransactions(c *gin.Context) {
	id := c.GetUint("id")
	var (
		result gin.H
	) 

	transactions, err := repository.GetAllTransactions(database.DbConnection, id)

	if err != nil {
		result = gin.H{
			"result":  err,
		}
	} else {
		result = gin.H{
			"result":  transactions,
		}
	}

	c.JSON(http.StatusOK, result)
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
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		if goal != nil {
			goal.Amount += transactions.Amount
			err = repository.UpdateGoal(database.DbConnection, goal)
			if err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
				return
			}
		}
	}

		// Update jumlah (amount) tujuan dengan jumlah transaksi

	err = repository.InsertTransaction(database.DbConnection, transactions)
	if err != nil {
		panic(err)
	}
	c.JSON(http.StatusOK, gin.H{
		"result": "success insert transaction",
	})


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
	c.JSON(http.StatusOK, gin.H{
		"result": "success update transaction",
	})
}

func DeleteTransaction(c *gin.Context) {
	var transactions structs.Transactions
	id, _ := strconv.Atoi(c.Param("id"))
	transactions.ID = int(id)
	
	err := repository.DeleteTransaction(database.DbConnection, transactions)
	if err != nil {
		panic(err)
	}
	c.JSON(http.StatusOK, gin.H{
		"result": "success delete transaction",
	})
}