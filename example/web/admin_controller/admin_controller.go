package admin_controller

import (
	"time"
	"errors"
	"github.com/labstack/echo"
	. "github.com/rohanthewiz/logger"
	"gopkg.in/nullbio/null.v6"
	ctx "github.com/rohanthewiz/church/web/context"
	"github.com/rohanthewiz/church/resource/user"
	"github.com/rohanthewiz/church/db"
	"github.com/rohanthewiz/church/models"
	"github.com/rohanthewiz/church/util/stringops"
)

const supertoken = "891ea66e067cfa7ef9b2f2b8522943f50a11ae3fbe7d74ed212b055de1d209d4"

func AdminHandler(c echo.Context) error {
	c.String(200, "Hello Administrator!")
	return nil
}

func CreateTestEvents(c echo.Context) error {
	d, err := db.Db()
	if err != nil {
		c.Error(err); return err
	}
	username := c.(*ctx.CustomContext).Username
	dte, err := time.Parse("01/02/2006 -07", "06/15/2017 -05")
	if err != nil {
		return errors.New("Error parsing provided event values")
	}
	evt := &models.Event{
		Title:      "Picnic", Summary: null.NewString("It's gonna be great!", true),
		Slug:       stringops.SlugWithRandomString("Picnic"),
		EventDate:  dte,
		EventTime:  "14:30pm",
		Categories: []string{ "default" },
		UpdatedBy:  username,
		Published:  true,
	}
	if err := evt.Insert(d); err != nil {
		c.Error(err); return err
	}

	dte, err = time.Parse("01/02/2006 -07", "06/12/2017 -05")
	if err != nil {
		return errors.New("Error parsing provided event values")
	}
	evt = &models.Event{
		Title:      "Retreat", Summary: null.NewString("Get refreshed!", true),
		Slug:       stringops.SlugWithRandomString("Retreat"),
		EventDate:  dte,
		EventTime:  "10:00AM",
		Categories: []string{ "default" },
		UpdatedBy:  username,
		Published:  true,
	}
	if err := evt.Insert(d); err != nil {
		c.Error(err); return err
	}
	c.String(200, "Events created")
	return nil
}

// If no superadmins exist and you have the right token in "supertok"
// create a superadministrator - todo
func SetupSuperAdmin(c echo.Context) error {
	// if no superadmins exist and c.FormValue...
	exists, err := user.SuperAdminsExist()
	if err != nil {
		return err
	}
	if exists {
		Log("info", "Superadmin already exists")
		return err
	}
	Log("info", "Superadmin does not exist. We'll create one.")

	if c.FormValue("supertok") == supertoken {
		Log("Info", "Here we would create a super administrator")
		// create a superadministrator ... todo
	} else {
		Log("Info", "Valid token not provided")
	}
	return nil
}
