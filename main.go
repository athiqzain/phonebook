package main

import (
	"fmt"
	"net/http"

	"errors"

	"github.com/gin-gonic/gin"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type contact struct {
	Id     string `json:"id"`
	Name   string `json:"name"`
	Number string `json:"number"`
}

var contacts = []contact{
	{Id: "1", Name: "Athiq", Number: "789546123"},
	{Id: "2", Name: "Zain", Number: "852147963"},
	{Id: "3", Name: "Adam", Number: "879456321"},
}

func getContact(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, contacts)
}

func contactByName(c *gin.Context) {
	name := c.Param("name")
	contact, err := getContactByname(name)

	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "contact not found."})
		return
	}

	c.IndentedJSON(http.StatusOK, contact)
}

func updateContact(c *gin.Context) {
	name, ok := c.GetQuery("name")

	if !ok {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "Missing name query parameter."})
		return
	}

	contact, err := getContactByname(name)

	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "contact not found."})
		return
	}

	var newNumber string
	fmt.Println("Enter the new Number: ")
	fmt.Scan(&newNumber)
	contact.Number = newNumber
	c.IndentedJSON(http.StatusOK, contact)
}

func getContactByname(name string) (*contact, error) {
	for i, b := range contacts {
		if b.Name == name {
			return &contacts[i], nil
		}
	}

	return nil, errors.New("contact not found")
}

func createContact(c *gin.Context) {
	var newContact contact

	if err := c.BindJSON(&newContact); err != nil {
		return
	}

	contacts = append(contacts, newContact)
	c.IndentedJSON(http.StatusCreated, newContact)
}

func deleteContact(c *gin.Context) {
	id := c.Param("id")

	for i, a := range contacts {
		if a.Id == id {
			contacts = append(contacts[:i], contacts[i+1:]...)
			break
		}
	}

	c.Status(http.StatusNoContent)
}

func main() {
	server := gin.Default()
	server.GET("/contacts", getContact)
	server.GET("/contacts/:name", contactByName)
	server.POST("/contacts", createContact)
	server.PATCH("/update", updateContact)
	server.DELETE("/delete/:id", deleteContact)

	server.Run(":8080")
}
