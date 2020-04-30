package tshirts

import "fmt"
import "github.com/Coffie/fortress/database"
import "github.com/Coffie/fortress/models"
import "github.com/jinzhu/gorm"
import "github.com/stretchr/testify/assert"
import "testing"

type TestVars struct {
	TshirtGroup models.TshirtGroup
	CreatedFlag models.Flag
}

var db *gorm.DB
var g TestVars

func init() {
	db = database.NewDB("localhost", "5432", "postgres", "postgres", "")
	database.DropAll(db)
	database.Migrate(db)
	g = TestVars{
		TshirtGroup: models.TshirtGroup{FlagID: 1, Name: "tg1"},
	}
}

func perTestSetup() *TshirtService {
	database.TruncateAll(db)
	tshirtService := New(db)
	flag := models.Flag{g.CreatedFlag.ID, "bla", "bla"}
	db.Create(&flag)
	g.CreatedFlag = flag
	return tshirtService
}

func TestTshirtGroups(t *testing.T) {
	service := perTestSetup()
	t.Run("a group can be added", func(t *testing.T) {
		g.TshirtGroup.ID = uint(100)
		res, _ := service.AddTshirtGroup(g.TshirtGroup)

		assert.IsType(t, g.TshirtGroup.ID, res.ID)
	})

	service = perTestSetup()
	t.Run("retrieving group does not exist yields nil", func(t *testing.T) {
		res, err := service.GetTshirtGroup("notexists")
		assert.Equal(t, models.TshirtGroup{}, res)
		assert.NotNil(t, err)
	})

	service = perTestSetup()
	t.Run("an added group can be retrieved", func(t *testing.T) {
		expected, _ := service.AddTshirtGroup(g.TshirtGroup)
		actual, err := service.GetTshirtGroup(expected.Name)
		assert.Equal(t, expected, actual)
		assert.Nil(t, err)
	})

	service = perTestSetup()
	t.Run("multiple groups can be retrieved", func(t *testing.T) {
		expected, _ := service.AddTshirtGroup(models.TshirtGroup{FlagID: g.CreatedFlag.ID, Name: "tg1"})
		expected2, _ := service.AddTshirtGroup(models.TshirtGroup{FlagID: g.CreatedFlag.ID, Name: "tg2"})

		actual, err := service.GetTshirtGroup(expected.Name)
		actual2, err2 := service.GetTshirtGroup(expected2.Name)

		assert.Equal(t, expected, actual)
		assert.Nil(t, err)
		assert.NotEqual(t, 0, actual.FlagID)
		assert.Nil(t, err2)
		assert.Equal(t, expected2, actual2)
	})

	service = perTestSetup()
	t.Run("no groups yields empty list", func(t *testing.T) {
		res, err := service.ListTshirtGroups()
		assert.Equal(t, []models.TshirtGroup{}, res)
		assert.Nil(t, err)
	})

	service = perTestSetup()
	service.AddTshirtGroup(g.TshirtGroup)
	service.AddTshirtGroup(models.TshirtGroup{
		FlagID: g.CreatedFlag.ID,
		Name:   "tg2",
	})
	t.Run("groups can be listed", func(t *testing.T) {
		res, err := service.ListTshirtGroups()
		assert.Len(t, res, 2)
		assert.NotEqual(t, res[0].ID, 0)
		assert.Nil(t, err)
	})

	service = perTestSetup()
	createdTshirtGroup, _ := service.AddTshirtGroup(g.TshirtGroup)
	otherTshirtGroup, _ := service.AddTshirtGroup(models.TshirtGroup{
		FlagID: g.CreatedFlag.ID,
		Name:   "",
	})
	t.Run("a group can be deleted", func(t *testing.T) {
		err := service.DeleteTshirtGroup(createdTshirtGroup.Name)
		shouldNotExist, _ := service.GetTshirtGroup(createdTshirtGroup.Name)
		assert.Nil(t, err)
		assert.Equal(t, models.TshirtGroup{}, shouldNotExist)

		shouldExist, _ := service.GetTshirtGroup(otherTshirtGroup.Name)
		assert.Equal(t, otherTshirtGroup, shouldExist)
	})
}

func TestTshirts(t *testing.T) {
	service := perTestSetup()
	t.Run("a tshirt can be added", func(t *testing.T) {
		createdTshirtGroup, _ := service.AddTshirtGroup(g.TshirtGroup)
		expected := models.Tshirt{1, createdTshirtGroup.ID, "XL", "red"}
		actual, err := service.AddTshirt(expected)
		assert.Equal(t, expected, actual)
		assert.Nil(t, err)
	})

	service = perTestSetup()
	t.Run("adding a tshirt to a missing group returns an error", func(t *testing.T) {
		expected := models.Tshirt{1, 1000, "XL", "red"}
		actual, err := service.AddTshirt(expected)
		assert.Equal(t, models.Tshirt{}, actual)
		assert.NotNil(t, err)
	})

	service = perTestSetup()
	t.Run("listing from missing tshirt group returns error", func(t *testing.T) {
		actual, err := service.ListTshirts("notexist")
		fmt.Printf("%+v", err)
		assert.Equal(t, []models.Tshirt{}, actual)
		assert.NotNil(t, err)
	})

	service = perTestSetup()
	service.AddTshirtGroup(g.TshirtGroup)
	t.Run("listing no tshirts yields empty list", func(t *testing.T) {
		actual, err := service.ListTshirts(g.TshirtGroup.Name)
		assert.Equal(t, []models.Tshirt{}, actual)
		assert.Nil(t, err)
	})

	service = perTestSetup()
	createdTshirtGroup, _ := service.AddTshirtGroup(g.TshirtGroup)
	otherTshirtGroup, _ := service.AddTshirtGroup(models.TshirtGroup{
		FlagID: 1,
		Name:   "otherGroup",
	})
	t.Run("listing tshirts in group returns only tshirts in that group", func(t *testing.T) {
		ts1, _ := service.AddTshirt(models.Tshirt{TshirtGroupID: createdTshirtGroup.ID, Size: "XL", Color: "red"})
		ts2, _ := service.AddTshirt(models.Tshirt{TshirtGroupID: createdTshirtGroup.ID, Size: "XL", Color: "red"})
		service.AddTshirt(models.Tshirt{TshirtGroupID: otherTshirtGroup.ID, Size: "XL", Color: "red"})
		actual, err := service.ListTshirts(createdTshirtGroup.Name)
		assert.Equal(t, []models.Tshirt{ts1, ts2}, actual)
		assert.Nil(t, err)
	})

	service = perTestSetup()
	createdTshirtGroup, _ = service.AddTshirtGroup(g.TshirtGroup)
	createdTshirt, _ := service.AddTshirt(models.Tshirt{
		TshirtGroupID: createdTshirtGroup.ID,
		Size:          "xxs",
		Color:         "magenta",
	})
	otherTshirt, _ := service.AddTshirt(models.Tshirt{
		TshirtGroupID: createdTshirtGroup.ID,
		Size:          "xl",
		Color:         "magenta",
	})
	t.Run("a tshirt can be deleted", func(t *testing.T) {
		err := service.DeleteTshirt(createdTshirtGroup.Name, createdTshirt.Size, createdTshirt.Color)
		tshirts, _ := service.ListTshirts(createdTshirtGroup.Name)
		assert.Nil(t, err)
		assert.Contains(t, tshirts, otherTshirt)
		assert.NotContains(t, tshirts, createdTshirt)
	})
}

func TestFlags(t *testing.T) {
	service := perTestSetup()
	t.Run("a flag can be added", func(t *testing.T) {
		expected := models.Flag{100, "bloop", "doop"}
		actual, err := service.AddFlag(expected)
		assert.Nil(t, err)
		assert.Equal(t, expected, actual)
	})
	t.Helper()
	service = perTestSetup()
	t.Run("a flag can be retrieved", func(t *testing.T) {
		actual, err := service.GetFlag(g.CreatedFlag.Name)
		assert.Equal(t, g.CreatedFlag, actual)
		assert.Nil(t, err)
	})

	service = perTestSetup()
	db.Delete(&models.Flag{})
	t.Run("no flags yields empty list", func(t *testing.T) {
		actual, err := service.ListFlags()
		assert.Nil(t, err)
		assert.Equal(t, []models.Flag{}, actual)
	})

	service = perTestSetup()
	t.Run("can list all flags", func(t *testing.T) {
		actual, err := service.ListFlags()
		assert.Nil(t, err)
		assert.Equal(t, []models.Flag{g.CreatedFlag}, actual)
	})

	service = perTestSetup()
	t.Run("a flag can be deleted", func(t *testing.T) {
		err := service.DeleteFlag(g.CreatedFlag.Name)
		assert.Nil(t, err)
		flagList, _ := service.ListFlags()
		assert.Equal(t, []models.Flag{}, flagList)
	})
}
