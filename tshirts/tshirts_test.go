package tshirts

import (
	"testing"

	"github.com/Coffie/fortress/database"
	"github.com/Coffie/fortress/models"
	"github.com/stretchr/testify/assert"
)

func tshirtGroupSetup() (*TshirtService, models.TshirtGroup) {
	db := database.NewDB("localhost", "5432", "postgres", "postgres", "")
	database.DropAll(db)
	database.Migrate(db)

	tshirtService := New(db)
	db.Create(&models.Flag{1, "bla", "bla"})
	tshirtGroup := models.TshirtGroup{FlagID: 1, Name: "tg1"}
	return tshirtService, tshirtGroup
}

func TestTshirtGroups(t *testing.T) {
	t.Run("a group can be added", func(t *testing.T) {
		service, tshirtGroup := tshirtGroupSetup()
		res, _ := service.AddTshirtGroup(tshirtGroup)
		assert.Equal(t, uint(1), res.ID)
	})
	t.Run("retrieving group does not exist yields nil", func(t *testing.T) {
		service, _ := tshirtGroupSetup()
		res, err := service.GetTshirtGroup("notexists")
		assert.Equal(t, models.TshirtGroup{}, res)
		assert.NotNil(t, err)
	})
	t.Run("an added group can be retrieved", func(t *testing.T) {
		service, tshirtGroup := tshirtGroupSetup()
		expected, _ := service.AddTshirtGroup(tshirtGroup)
		actual, err := service.GetTshirtGroup(expected.Name)
		assert.Equal(t, expected, actual)
		assert.Nil(t, err)
	})
	t.Run("multiple groups can be retrieved", func(t *testing.T) {
		service, _ := tshirtGroupSetup()

		expected, _ := service.AddTshirtGroup(models.TshirtGroup{FlagID: 1, Name: "tg1"})
		expected2, _ := service.AddTshirtGroup(models.TshirtGroup{FlagID: 1, Name: "tg2"})

		actual, err := service.GetTshirtGroup(expected.Name)
		actual2, err2 := service.GetTshirtGroup(expected2.Name)

		assert.Equal(t, expected, actual)
		assert.Nil(t, err)
		assert.Equal(t, expected2, actual2)
		assert.Nil(t, err2)
	})
	t.Run("no groups yields empty list", func(t *testing.T) {
		service, _ := tshirtGroupSetup()
		res, err := service.ListTshirtGroups()
		assert.Equal(t, []models.TshirtGroup{}, res)
		assert.Nil(t, err)
	})
	t.Run("one group can be listed", func(t *testing.T) {
		service, tshirtGroup := tshirtGroupSetup()
		service.AddTshirtGroup(tshirtGroup)
		res, err := service.ListTshirtGroups()
		assert.Len(t, res, 1)
		assert.Equal(t, res[0].ID, uint(1))
		assert.Nil(t, err)
	})
}

func TestTshirts(t *testing.T) {
	t.Run("a tshirt can be added", func(t *testing.T) {
		server, tshirtGroup := tshirtGroupSetup()
		addedTshirtGroup, _ := server.AddTshirtGroup(tshirtGroup)
		expected := models.Tshirt{1, addedTshirtGroup.ID, "XL", "red"}
		actual, err := server.AddTshirt(expected)
		assert.Equal(t, expected, actual)
		assert.Nil(t, err)
	})

	t.Run("listing with no tshirt group returns error", func(t *testing.T) {
		service, _ := tshirtGroupSetup()
		actual, err := service.ListTshirts("notexist")
		assert.Equal(t, []models.Tshirt{}, actual)
		assert.Nil(t, err)
	})

	t.Run("listing no tshirts yields empty list", func(t *testing.T) {
		service, tshirtGroup := tshirtGroupSetup()
		service.AddTshirtGroup(tshirtGroup)
		actual, err := service.ListTshirts(tshirtGroup.Name)
		assert.Equal(t, []models.Tshirt{}, actual)
		assert.Nil(t, err)
	})
}
