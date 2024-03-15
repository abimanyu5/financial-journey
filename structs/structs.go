package structs

import (
	"github.com/golang-jwt/jwt"
	"time"
)

type Masters struct {
	ID          int       `json:"id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	UserId      int       `json:"user_id"`
	IncomeType  int       `json:"income_type"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

type Transactions struct {
	ID          int       `json:"id"`
	MasterId    int       `json:"master_id"`
	GoalId      int       `json:"goal_id"`
	UserId      int       `json:"user_id"`
	Amount      int       `json:"amount"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

type User struct {
	ID        uint      `json:"id"`
	Username  string    `json:"username"`
	Password  string    `json:"password"`
	Email     string    `json:"email"`
	Role      string    `json:"role"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// JWTClaims represents the claims for JWT token
type JWTClaims struct {
	Username string `json:"username"`
	Role     string `json:"role"`
	ID       uint   `json:"id"`
	jwt.StandardClaims
}

type LoginData struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type Goals struct {
	ID          int       `json:"id"`
	UserId      int       `json:"user_id"`
	Amount      int       `json:"amount"`
	AmountGoal  int       `json:"amount_goal"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

// type Reports struct {
// 	ID            int       `json:"id"`
// 	UserId        int       `json:"user_id"`
// 	TotalAmount   int       `json:"total_amount"`
// 	AmountIncome  int       `json:"amount_income"`
// 	AmountOutcome int       `json:"amount_outcome"`
// 	CreatedAt     time.Time `json:"created_at"`
// 	UpdatedAt     time.Time `json:"updated_at"`
// }
