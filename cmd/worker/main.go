// cmd/worker/main.go
package main

import (
	"context"
	"encoding/json"
	"github.com/hibiken/asynq"
	log "github.com/sirupsen/logrus"
	"github.com/tsw025/web_analytics/internal/config"
	"github.com/tsw025/web_analytics/internal/database"
	"github.com/tsw025/web_analytics/internal/models"
	"github.com/tsw025/web_analytics/internal/repositories"
	"github.com/tsw025/web_analytics/internal/services"
	"github.com/tsw025/web_analytics/internal/tasks"
	"gorm.io/datatypes"
)

// Task Handler for AnalyzeWebsite
func handleAnalyzeWebsiteTask(ctx context.Context, t *asynq.Task) error {
	var payload tasks.AnalyzeWebsitePayload
	if err := json.Unmarshal(t.Payload(), &payload); err != nil {
		log.Infof("Failed to unmarshal payload: %v", err)
		return err
	}

	log.Infof("Processing AnalyzeWebsite task: URL=%s, AnalyticsID=%d", payload.URL, payload.AnalyticsID)

	// Initialize database connection
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("Failed to load configuration: %v", err)
	}

	db, err := database.ConnectToPostgres(cfg)
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	// Initialize repositories
	analyticsRepo := repositories.NewAnalyticsRepository(db)

	// Initialize worker service
	workerService := services.NewWorkerService(analyticsRepo)

	// Perform analysis
	analysisData, err := workerService.PerformAnalysis(payload.URL)
	if err != nil {
		// Update Analytics status to Failed
		if updateErr := analyticsRepo.UpdateStatus(payload.AnalyticsID, models.StatusFailed); updateErr != nil {
			log.Infof("Error updating analytics status to Failed for ID %d: %v", payload.AnalyticsID, updateErr)
		}
		log.Infof("Analysis failed for website %s: %v", payload.URL, err)
		return err
	}

	// Update Analytics with the result and set status to Completed
	if err := analyticsRepo.UpdateDataAndStatus(payload.AnalyticsID, datatypes.JSON(analysisData), models.StatusCompleted); err != nil {
		log.Infof("Error updating analytics data for ID %d: %v", payload.AnalyticsID, err)
		return err
	}

	log.Infof("Analysis completed for website %s (Analytics ID: %d)", payload.URL, payload.AnalyticsID)
	return nil
}

func main() {
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("Failed to load configuration: %v", err)
	}
	// Configure Asynq server options
	redisOpt := asynq.RedisClientOpt{
		Addr: cfg.RedisAddr, // Redis server address
		DB:   0,
	}

	mux := asynq.NewServeMux()
	mux.HandleFunc(tasks.TypeAnalyzeWebsite, handleAnalyzeWebsiteTask)

	srv := asynq.NewServer(
		redisOpt,
		asynq.Config{
			Concurrency: 10,
		},
	)

	log.Println("Worker is running...")
	if err := srv.Run(mux); err != nil {
		log.Fatalf("Could not run server: %v", err)
	}
}
