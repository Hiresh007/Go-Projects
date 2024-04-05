package leads

import (
	"go-fibre/database"

	"github.com/gofiber/fiber"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

type Lead struct {
	gorm.Model
	Name    string `json:"name"`
	Company string `json:"company"`
	Email   string `json:"email"`
	Phone   string `json:"phone"`
}

func GetLeads(c *fiber.Ctx) {
	db := database.DBCon
	var leads []Lead
	db.Find(&leads)
	c.JSON(leads)
}

func GetLead(c *fiber.Ctx) {
	id := c.Params("id")
	db := database.DBCon
	var lead Lead
	db.Find(&lead, id)
	c.JSON(lead)
}

func DeleteLead(c *fiber.Ctx) {
	db := database.DBCon
	id := c.Params("id")
	var lead Lead
	db.First(&lead, id)
	if lead.Name == "" {
		c.Status(500).Send([]byte("Not found lead"))
		return
	}
	db.Delete(&lead, id)
	c.Send([]byte("Deleted Successfullys"))
}

func NewLead(c *fiber.Ctx) {
	db := database.DBCon
	lead := new(Lead)
	if err := c.BodyParser(lead); err != nil {
		c.Status(503).Send([]byte(err.Error()))
		return
	}
	db.Create(&lead)
	c.JSON(lead)
}
