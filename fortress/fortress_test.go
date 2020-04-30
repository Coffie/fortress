package fortress

import (
	"testing"

	"github.com/Coffie/fortress/database"
	"github.com/Coffie/fortress/models"
	"github.com/jinzhu/gorm"
	"github.com/stretchr/testify/assert"
)

var db *gorm.DB

func init() {
	db = database.NewDB("localhost", "5432", "postgres", "postgres", "")
	database.DropAll(db)
	database.Migrate(db)
}

func perTestSetup() *Fortress {
	database.TruncateAll(db)
	return New(db)
}

func TestThings(t *testing.T) {
	service := perTestSetup()
	t.Run("can create a thing", func(t *testing.T) {
		thing := models.FortressThing{123, 10}
		createdThing, err := service.AddThing(thing)
		assert.Nil(t, err)
		assert.Equal(t, thing, createdThing)
	})

	service = perTestSetup()
	createdThing, _ := service.AddThing(models.FortressThing{123, 10})
	t.Run("can retrieve a thing", func(t *testing.T) {
		thing, err := service.GetThing(createdThing.ID)
		assert.Nil(t, err)
		assert.Equal(t, createdThing, thing)
	})

	service = perTestSetup()
	createdThing, _ = service.AddThing(models.FortressThing{123, 10})
	t.Run("can delete a thing", func(t *testing.T) {
		err := service.DeleteThing(createdThing.ID)
		shouldNotExist, _ := service.GetThing(createdThing.ID)
		assert.Nil(t, err)
		assert.Equal(t, models.FortressThing{}, shouldNotExist)
	})
}
