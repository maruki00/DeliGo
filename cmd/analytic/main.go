package main



func main() {
	log.Println("Starting Feedback & Analytics Service deployment runtime initialization...")

	cfg := config.LoadConfig()

	// 1. Core Data Store Architecture Setups
	db := database.InitDB(cfg)
	log.Println("Executing automatic structural migrations updates rules checks...")
	if err := db.AutoMigrate(&domain.Feedback{}, &domain.Analytic{}); err != nil {
		log.Fatalf("Database structural layout binding migrations failures detected: %v", err)
	}

	// 2. Messaging Pipeline Topology Wiring
	rabbitConn, rabbitChannel := messaging.InitRabbitMQ(cfg)
	defer rabbitConn.Close()
	defer rabbitChannel.Close()

	// 3. Repository Dependency Injections Graph Mapping
	feedbackRepo := repository.NewFeedbackRepository(db)
	analyticRepo := repository.NewAnalyticRepository(db)

	publisher := messaging.NewEventPublisher(rabbitChannel, cfg)

	// 4. Start Background Message Pipeline Engine Lifecycle Listening Context
	consumer := messaging.NewEventConsumer(rabbitChannel, cfg, analyticRepo)
	consumer.StartConsuming()

	// 5. Service & Handler Core System Layers Wiring
	feedbackSvc := service.NewFeedbackService(feedbackRepo, analyticRepo, publisher)
	feedbackHandler := handler.NewFeedbackHandler(feedbackSvc)

	engine := router.SetupRouter(feedbackHandler)

	// 6. Graceful Server Lifecycle Management Shutdown Setups
	srv := &http.Server{
		Addr:    ":" + cfg.ServerPort,
		Handler: engine,
	}

	go func() {
		log.Printf("Application engine network listening up over port :%s", cfg.ServerPort)
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("Critical runtime HTTP failure error hook caught: %s\n", err)
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Println("Shutdown signal intercepted. Terminating operational endpoints down cleanly...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		log.Fatalf("Server forced to shutdown incorrectly: %v", err)
	}

	log.Println("Service closed out smoothly.")
}
