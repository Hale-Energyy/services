package models

import "github.com/lib/pq"

type UserFitnessProfile struct {
	ID                 uint           `gorm:"primaryKey" json:"id"`
	UserID             uint           `json:"user_id"` // foreign key reference to User
	FitnessLevel       string         `json:"fitness_level"`
	FitnessGoal        string         `json:"fitness_goal"`
	WorkoutPreferences pq.StringArray `json:"workout_preferences" gorm:"type:text[]"`
	DaysPerWeek        int            `json:"days_per_week"`
	HealthCondition    string         `json:"health_condition"`
	BloodType          string         `json:"blood_type"`
}
