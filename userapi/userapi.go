package userapi

import (
	"database/sql"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type UserBankService interface {
	FindByID(id int) (*User, error)
	All() ([]User, error)
	CreateUser(user *User) error
	Update(user *User) error
	Delete(user *User) error
	CreateBank(id int, b BankAccount) error
	AllBankByID(id int) ([]BankAccount, error)
	DeleteBankByID(id int) error
	Withdraw(id int, amount float64) error
	Deposit(id int, amount float64) error
	Transfer(fromID int, toID int, amount float64) error
}

type Handler struct {
	userBankService UserBankService
}

func (h *Handler) getUser(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	user, err := h.userBankService.FindByID(id)
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusOK, user)
}

func (h *Handler) getAllUser(c *gin.Context) {

}

func (h *Handler) createUser(c *gin.Context) {

}

func (h *Handler) updateUser(c *gin.Context) {

}

func (h *Handler) deleteUser(c *gin.Context) {

}

func (h *Handler) createBank(c *gin.Context) {

}

func (h *Handler) getAllBankByID(c *gin.Context) {

}

func (h *Handler) deleteBankAccount(c *gin.Context) {

}

func (h *Handler) withdrawMoney(c *gin.Context) {

}

func (h *Handler) depositMoney(c *gin.Context) {

}

func (h *Handler) transferMoney(c *gin.Context) {

}

func StartServer(addr string, db *sql.DB) error {
	r := gin.Default()
	h := &Handler{
		userBankService: &Service{
			DB: db,
		},
	}

	r.GET("/users/:id", h.getUser)
	r.GET("/users", h.getAllUser)
	r.POST("/users/:id", h.createUser)
	r.PUT("/users/:id", h.updateUser)
	r.DELETE("/users/:id", h.deleteUser)
	r.POST("/users/:id/bankAccounts", h.createBank)
	r.GET("/users/:id/bankAccounts", h.getAllBankByID)
	r.DELETE("/bankAccounts", h.deleteBankAccount)
	r.PUT("/bankAccounts/:id/withdraw", h.withdrawMoney)
	r.PUT("/bankAccounts/:id/deposit", h.depositMoney)
	r.POST("/transfers", h.transferMoney)

	return r.Run(addr)
}
