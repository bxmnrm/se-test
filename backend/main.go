package main

import (
	"log"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"

	"backend/config"
	"backend/controllers"
	entity "backend/entity"
)

func main() {
	r := gin.Default()

	// Configure CORS
	configCORS := cors.DefaultConfig()
	configCORS.AllowOrigins = []string{"http://localhost:5173", "http://127.0.0.1:5173"} // Allow frontend dev server
	configCORS.AllowMethods = []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"}
	configCORS.AllowHeaders = []string{"Origin", "Content-Type", "Authorization"}
	r.Use(cors.New(configCORS))

	// Connect to Database
	config.ConnectDB()

	// Auto-migrate the schema
	if err := config.DB.AutoMigrate(
		&entity.User{},
		&entity.Industry{},
		&entity.Sponsor{},
		&entity.SponsorContact{},
	); err != nil {
		log.Fatalf("AutoMigrate failed: %v", err)
	} else {
		log.Println("AutoMigrate completed")
	}

	SeedIndustries()

	// API routes
	api := r.Group("/api")
	{
		api.POST("/register", controllers.Register)
		api.POST("/login", controllers.Login)

		api.GET("/industries", controllers.GetIndustries)
		api.POST("/industries", controllers.CreateIndustry)

		api.GET("/sponsors", controllers.GetSponsors)
		api.POST("/sponsors", controllers.CreateSponsor)
		api.PATCH("/sponsors/:id", controllers.UpdateSponsor)
		api.DELETE("/sponsors/:id", controllers.DeleteSponsor)

		/*api.GET("/Screenings", controllers.GetScreenings)
		api.POST("/Screenings", controllers.CreateScreening)
		api.PUT("/Screenings/:id", controllers.UpdateScreening)
		api.DELETE("/Screenings/:id", controllers.DeleteScreening)

		api.GET("/NewsPosts", controllers.GetNewsPosts)
		api.GET("/NewsPosts/:newsID", controllers.GetNewsPostByID)
		api.POST("/NewsPosts", controllers.CreateNewsPost)
		api.PUT("/NewsPosts/:newsID", controllers.UpdateNewsPost)

		api.GET("/student-fav/:studentID", controllers.GetStudentFavsByStudent)
		api.POST("/student-fav", controllers.CreateStudentFav)
		api.DELETE("/student-fav/:studentID/:newsID", controllers.DeleteStudentFav)*/
	}

	r.Run() // listen and serve on 0.0.0.0:8080
}

func SeedIndustries() {
	industries := []entity.Industry{
		{Name: "Technology"},
		{Name: "Finance"},
		{Name: "Healthcare"},
		{Name: "Education"},
		{Name: "Manufacturing"},
		{Name: "Energy"},
		{Name: "Non-Profit"},
	}

	for _, i := range industries {
		result := config.DB.FirstOrCreate(&i, entity.Industry{Name: i.Name})
		if result.Error != nil {
			log.Printf("Failed to seed industry '%s': %v", i.Name, result.Error)
		}
	}

	log.Println("Seeded industries (FirstOrCreate).")
}