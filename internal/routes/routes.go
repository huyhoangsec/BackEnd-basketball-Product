package routes

import (
	"backend-ocean-basketball/config"
	"backend-ocean-basketball/internal/handlers"
	"backend-ocean-basketball/internal/middleware"
	"github.com/gin-gonic/gin"
)

func SetupRoutes(r *gin.RouterGroup, cfg *config.Config) {
	authHandler := handlers.NewAuthHandler(cfg)

	// Public routes
	r.POST("/auth/login", authHandler.Login)
    
	r.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{"status": "ok"})
	})

	public := r.Group("/public")
	{
		public.GET("/coaches", handlers.GetCoaches)
		public.GET("/courts", handlers.GetCourts)
		public.GET("/classes", handlers.GetClasses)
		public.GET("/tuition-plans", handlers.GetTuitionPlans)
		public.GET("/faqs", handlers.GetFAQs)
		public.GET("/reviews", handlers.GetReviews)
		public.GET("/banners", handlers.GetBanners)
		
		// Public trial registration
		public.POST("/trials", handlers.SubmitTrialRegistration)

		// Public contact form
		public.POST("/contacts", handlers.SubmitContact)

		// Public tournaments
		public.GET("/tournaments", handlers.GetTournaments)
		public.POST("/tournaments/:id/register", handlers.RegisterTournament)
	}

	// Protected Admin routes
	admin := r.Group("/admin")
	admin.Use(middleware.AuthMiddleware(cfg, "admin"))
	{
		// Profile & Password
		admin.GET("/profile", handlers.GetProfile)
		admin.PUT("/profile", handlers.UpdateProfile)
		admin.PUT("/change-password", handlers.ChangePassword(cfg))
		admin.POST("/refresh-token", handlers.RefreshToken(cfg))

		admin.GET("/trials", handlers.GetTrials)
		admin.PUT("/trials/:id/status", handlers.UpdateTrialStatus)
		
		// Upload
		admin.POST("/upload", handlers.UploadFile)
		
		// Coaches
		admin.GET("/coaches", handlers.GetAllCoaches)
		admin.POST("/coaches", handlers.CreateCoach)
		admin.PUT("/coaches/:id", handlers.UpdateCoach)
		admin.DELETE("/coaches/:id", handlers.DeleteCoach)

		// Courts
		admin.POST("/courts", handlers.CreateCourt)
		admin.PUT("/courts/:id", handlers.UpdateCourt)
		admin.DELETE("/courts/:id", handlers.DeleteCourt)

		// Classes
		admin.POST("/classes", handlers.CreateClass)
		admin.PUT("/classes/:id", handlers.UpdateClass)
		admin.DELETE("/classes/:id", handlers.DeleteClass)

		// Students
		admin.GET("/students", handlers.GetStudents)
		admin.GET("/students/:id", handlers.GetStudentProfile)
		admin.POST("/students", handlers.CreateStudent)
		admin.PUT("/students/:id", handlers.UpdateStudent)
		admin.POST("/students/:id/transfer", handlers.TransferStudent)
		admin.GET("/students/:id/attendance", handlers.GetStudentAttendanceHistory)
		admin.DELETE("/students/:id", handlers.DeleteStudent)

		// Invoices
		admin.GET("/invoices", handlers.GetInvoices)
		admin.GET("/invoices/overdue", handlers.GetOverdueInvoices)
		admin.GET("/invoices/stats", handlers.GetInvoiceStats)
		admin.POST("/invoices", handlers.CreateInvoice)
		admin.PUT("/invoices/:id/pay", handlers.PayInvoice)
		admin.POST("/invoices/generate", handlers.GenerateMonthlyInvoices)
		
		// CMS (Banners, FAQs, Reviews, TuitionPlans)
		admin.GET("/banners", handlers.GetAllBanners)
		admin.POST("/banners", handlers.CreateBanner)
		admin.PUT("/banners/:id", handlers.UpdateBanner)
		admin.DELETE("/banners/:id", handlers.DeleteBanner)
		
		admin.GET("/faqs", handlers.GetAllFAQs)
		admin.POST("/faqs", handlers.CreateFAQ)
		admin.PUT("/faqs/:id", handlers.UpdateFAQ)
		admin.DELETE("/faqs/:id", handlers.DeleteFAQ)
		
		admin.GET("/reviews", handlers.GetAllReviews)
		admin.POST("/reviews", handlers.CreateReview)
		admin.PUT("/reviews/:id", handlers.UpdateReview)
		admin.DELETE("/reviews/:id", handlers.DeleteReview)
		
		admin.GET("/tuition-plans", handlers.GetAllTuitionPlans)
		admin.POST("/tuition-plans", handlers.CreateTuitionPlan)
		admin.PUT("/tuition-plans/:id", handlers.UpdateTuitionPlan)
		admin.DELETE("/tuition-plans/:id", handlers.DeleteTuitionPlan)
		
		// Tournaments
		admin.GET("/tournaments", handlers.GetTournaments)
		admin.POST("/tournaments", handlers.CreateTournament)
		admin.PUT("/tournaments/:id", handlers.UpdateTournament)
		admin.DELETE("/tournaments/:id", handlers.DeleteTournament)
		admin.GET("/tournaments/:id/matches", handlers.GetTournamentMatches)

		// Contacts
		admin.GET("/contacts", handlers.GetContacts)
		admin.PUT("/contacts/:id/read", handlers.MarkContactRead)
		admin.DELETE("/contacts/:id", handlers.DeleteContact)

		// Coach stats
		admin.GET("/coaches/:id/workload", handlers.GetCoachWorkload)
		admin.GET("/coaches/stats", handlers.GetAllCoachesStats)

		// Stats & Analytics
		admin.GET("/stats/overview", handlers.GetStatsOverview)
		admin.GET("/stats/revenue", handlers.GetStatsEnrollment)
		admin.GET("/stats/attendance", handlers.GetStatsCoach)
		admin.GET("/stats/distribution", handlers.GetStatsDistribution)
		admin.GET("/stats/attendance-report", handlers.GetAttendanceReport)
		admin.GET("/stats/attendance-by-date", handlers.GetAttendanceByDate)
	}

	// Protected Coach routes
	coach := r.Group("/coach")
	coach.Use(middleware.AuthMiddleware(cfg, "coach", "admin")) // Admin can also access
	{
		coach.GET("/profile", handlers.GetProfile)
		coach.PUT("/profile", handlers.UpdateProfile)
		coach.PUT("/change-password", handlers.ChangePassword(cfg))

		coach.GET("/classes", handlers.GetCoachClasses)
		coach.GET("/classes/:classId/students", handlers.GetClassStudents)
		coach.POST("/classes/:classId/attendance", handlers.MarkAttendance)
		coach.GET("/classes/:classId/attendance", handlers.GetAttendanceHistory)
	}
}
