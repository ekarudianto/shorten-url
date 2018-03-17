package actions

import (
	"github.com/buffalo/app/models"
	"github.com/gobuffalo/buffalo"
	"github.com/gobuffalo/pop"
)

// HomeHandler is a default handler to serve up
// a home page.
func HomeHandler(c buffalo.Context) error {
	return c.Render(200, r.HTML("index.html"))
}

// Decode url
func DecodeURL(c buffalo.Context) error {
	return c.Render(200, r.JSON(map[string]string{"version": "v1.0.0"}))
}

// Short url
func ShortenURL(c buffalo.Context) error {
	url := c.Request().FormValue("url")

	// Connect to DB
	tx, ok := c.Value("tx").(*pop.Connection)

	if !ok {
		return c.Render(500, r.JSON(map[string]string{"error": "Database connection error"}))
	}

	link := &models.Link{}

	// Find existing record and if fails
	// Create one registry
	if err := tx.Where("links.link = ?", url).First(link); err != nil {
		link.Link = url
		link.ShortLink = link.RandSeq(tx, 6)

		verrs, e := tx.ValidateAndCreate(link)

		if e != nil || verrs.HasAny() {
			return c.Render(500, r.JSON(map[string]string{"error": "Error inserting record"}))
		}

		return c.Render(200, r.JSON(link))
	}

	return c.Render(200, r.JSON(link))
}
