package services

import (
	"context"
	"time"

	"simple-go-api/config"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type AnalyticsService struct {
	collection *mongo.Collection
}

func NewAnalyticsService() *AnalyticsService {
	return &AnalyticsService{
		collection: config.MongoClient.Database(config.DB_NAME).Collection(config.EMPLOYEES_COLLECTION),
	}
}

// GetDepartmentDistribution returns employee count per department
func (s *AnalyticsService) GetDepartmentDistribution() ([]bson.M, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	pipeline := mongo.Pipeline{
		bson.D{{"$match", bson.D{{"employmentStatus", "active"}}}},
		bson.D{{"$group", bson.D{{"_id", "$department"}, {"count", bson.D{{"$sum", 1}}}, {"avgSalary", bson.D{{"$avg", "$salary"}}}}}},
		bson.D{{"$sort", bson.D{{"count", -1}}}},
	}

	cursor, err := s.collection.Aggregate(ctx, pipeline)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	var results []bson.M
	if err = cursor.All(ctx, &results); err != nil {
		return nil, err
	}
	return results, nil
}

// GetNewJoineesCount returns count of employees who joined in the given period
func (s *AnalyticsService) GetNewJoineesCount(startDate, endDate time.Time) (int64, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	pipeline := mongo.Pipeline{
		bson.D{
			{"$match",
				bson.D{
					{"joiningDate",
						bson.D{
							{"$gte", startDate},
							{"$lte", endDate},
						},
					},
				},
			},
		},

		bson.D{
			{"$count", "newJoinees"},
		},
	}

	cursor, err := s.collection.Aggregate(ctx, pipeline)
	if err != nil {
		return 0, err
	}
	defer cursor.Close(ctx)

	var results []bson.M
	if err = cursor.All(ctx, &results); err != nil {
		return 0, err
	}

	if len(results) == 0 {
		return 0, nil
	}
	return results[0]["newJoinees"].(int64), nil
}
