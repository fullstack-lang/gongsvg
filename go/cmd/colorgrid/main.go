package main

import (
	"embed"
	"flag"
	"fmt"
	"io/fs"
	"log"
	"net/http"
	"os"

	"github.com/gin-contrib/cors"
	"github.com/gin-contrib/static"
	"github.com/gin-gonic/gin"

	"github.com/fullstack-lang/gongsvg/go/controllers"
	"github.com/fullstack-lang/gongsvg/go/fullstack"
	"github.com/fullstack-lang/gongsvg/go/models"

	"github.com/fullstack-lang/gongsvg"
)

var (
	logDBFlag  = flag.Bool("logDB", false, "log mode for db")
	logGINFlag = flag.Bool("logGIN", false, "log mode for gin")
)

const RectangleWidth = 160
const RectangleHeigth = 50
const numberOfRectanglePerLine = 6

func main() {

	log.SetPrefix("gongsvg: ")
	log.SetFlags(0)

	// parse program arguments
	flag.Parse()

	// setup controlers
	if !*logGINFlag {
		myfile, _ := os.Create("/tmp/server.log")
		gin.DefaultWriter = myfile
	}
	r := gin.Default()
	r.Use(cors.Default())

	// setup GORM
	stage, _ := fullstack.NewStackInstance(r, "")

	svg := (&models.SVG{Name: "SVG2"}).Stage(stage)
	svg.Display = true

	// create an array of rectangle
	for idx, color := range models.Colors {
		column := float64(idx % numberOfRectanglePerLine)
		line := float64(int(
			float64(idx) /
				float64(numberOfRectanglePerLine)))

		appendRect(stage, svg, color,
			RectangleWidth*column,
			RectangleHeigth*line)
	}

	stage.Commit()

	controllers.RegisterControllers(r)

	// provide the static route for the angular pages
	r.Use(static.Serve("/", EmbedFolder(gongsvg.NgDistNg, "ng/dist/ng")))
	r.NoRoute(func(c *gin.Context) {
		fmt.Println(c.Request.URL.Path, "doesn't exists, redirect on /")
		c.Redirect(http.StatusMovedPermanently, "/")
		c.Abort()
	})

	log.Printf("Server ready serve on localhost:8082")
	r.Run(":8082")
}

type embedFileSystem struct {
	http.FileSystem
}

func (e embedFileSystem) Exists(prefix string, path string) bool {
	_, err := e.Open(path)
	return err == nil
}

func EmbedFolder(fsEmbed embed.FS, targetPath string) static.ServeFileSystem {
	fsys, err := fs.Sub(fsEmbed, targetPath)
	if err != nil {
		panic(err)
	}
	return embedFileSystem{
		FileSystem: http.FS(fsys),
	}
}

func appendRect(stage *models.StageStruct, svg *models.SVG, color models.ColorType, x, y float64) {
	rect := (&models.Rect{Name: color.ToString()}).Stage(stage)
	rect.Color = color.ToString()
	rect.Height = RectangleHeigth
	rect.Width = RectangleWidth
	rect.X = x
	rect.Y = y
	rect.FillOpacity = 50.0

	text := (&models.Text{Name: color.ToString()}).Stage(stage)
	text.X = x + 2
	text.Y = y + RectangleHeigth/2
	text.Content = color.ToString()
	text.FillOpacity = 100.0

	svg.Texts = append(svg.Texts, text)
	svg.Rects = append(svg.Rects, rect)
}
