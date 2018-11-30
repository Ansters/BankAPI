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
	Delete(id int) error
	CreateBank(b *BankAccount) error
	AllBankByID(id int) ([]BankAccount, error)
	DeleteBankByID(id int) error
	Withdraw(id int, amount Amount) error
	Deposit(id int, amount Amount) error
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
	users, err := h.userBankService.All()
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}
	c.JSON(http.StatusOK, users)
}

func (h *Handler) createUser(c *gin.Context) {
	var user User
	err := c.ShouldBindJSON(&user)
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}
	err = h.userBankService.CreateUser(&user)
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}
	c.JSON(http.StatusCreated, user)
}

func (h *Handler) updateUser(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}
	var user User
	user.ID = id
	err = c.ShouldBindJSON(&user)
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}
	err = h.userBankService.Update(&user)
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}
	c.JSON(http.StatusOK, user)
}

func (h *Handler) deleteUser(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}
	err = h.userBankService.Delete(id)
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}
	c.JSON(http.StatusOK, id)
}

func (h *Handler) createBank(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}
	var bankAccount BankAccount
	bankAccount.UserID = id
	bankAccount.Balance = 0.0
	err = c.ShouldBindJSON(&bankAccount)
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}
	err = h.userBankService.CreateBank(&bankAccount)
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}
	c.JSON(http.StatusCreated, bankAccount)
}

func (h *Handler) getAllBankByID(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}
	banks, err := h.userBankService.AllBankByID(id)
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}
	c.JSON(http.StatusOK, banks)
}

func (h *Handler) deleteBankAccount(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}
	err = h.userBankService.DeleteBankByID(id)
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}
	c.JSON(http.StatusOK, id)
}

func (h *Handler) withdrawMoney(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}
	var a Amount
	err = c.ShouldBindJSON(&a)
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}
	err = h.userBankService.Withdraw(id, a)
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}
	c.JSON(http.StatusOK, id)
}

func (h *Handler) depositMoney(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}
	var a Amount
	err = c.ShouldBindJSON(&a)
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}
	err = h.userBankService.Deposit(id, a)
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}
	c.JSON(http.StatusOK, id)
}

func (h *Handler) transferMoney(c *gin.Context) {
	var amount Amount
	err := c.ShouldBindJSON(&amount)
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}
	err = h.userBankService.Withdraw(amount.From, amount)
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}
	err = h.userBankService.Deposit(amount.To, amount)
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}
	c.JSON(http.StatusOK, "Success")
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
	r.POST("/users", h.createUser)
	r.PUT("/users/:id", h.updateUser)
	r.DELETE("/users/:id", h.deleteUser)
	r.POST("/users/:id/bankAccounts", h.createBank)
	r.GET("/users/:id/bankAccounts", h.getAllBankByID)
	r.DELETE("/bankAccounts/:id", h.deleteBankAccount)
	r.PUT("/bankAccounts/:id/withdraw", h.withdrawMoney)
	r.PUT("/bankAccounts/:id/deposit", h.depositMoney)
	r.POST("/transfers", h.transferMoney)

	return r.Run(addr)
}
