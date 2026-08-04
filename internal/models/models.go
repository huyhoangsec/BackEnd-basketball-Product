package models

import (
	"time"
)

// User represents admins or coaches who can log in
type User struct {
	ID        string    `gorm:"primaryKey" json:"id"`
	Name      string    `json:"name"`
	Email     string    `gorm:"unique" json:"email"`
	Password  string    `json:"-"` // Don't expose password
	Role      string    `json:"role"` // "admin" or "coach"
	Avatar    string    `json:"avatar"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// Coach profile details
type Coach struct {
	ID             string    `gorm:"primaryKey" json:"id"`
	Name           string    `json:"name"`
	Email          string    `json:"email"`
	Phone          string    `json:"phone"`
	Avatar         string    `json:"avatar"`
	Specialization string    `json:"specialization"`
	Experience     string    `json:"experience"`
	Achievements   []string  `gorm:"serializer:json" json:"achievements"`
	Bio            string    `json:"bio"`
	IsActive       bool      `json:"is_active"`
	CreatedAt      time.Time `json:"created_at"`
	UpdatedAt      time.Time `json:"updated_at"`
}

// Court details
type Court struct {
	ID         string    `gorm:"primaryKey" json:"id"`
	Name       string    `json:"name"`
	Address    string    `json:"address"`
	Latitude   float64   `json:"latitude"`
	Longitude  float64   `json:"longitude"`
	Images     []string  `gorm:"serializer:json" json:"images"`
	Facilities []string  `gorm:"serializer:json" json:"facilities"`
	ClassCount int       `json:"class_count"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
}

// Class details
type Class struct {
	ID              string          `gorm:"primaryKey" json:"id"`
	Name            string          `json:"name"`
	CourtID         string          `json:"court_id"`
	CoachID         string          `json:"coach_id"`
	Level           string          `json:"level"`
	Image           string          `json:"image"`
	MaxStudents     int             `json:"max_students"`
	CurrentStudents int             `gorm:"-" json:"current_students"` // Calculated field
	TrialStudents   int             `gorm:"-" json:"trial_students"`   // Calculated field
	Court           Court           `gorm:"foreignKey:CourtID" json:"court"`
	Coach           Coach           `gorm:"foreignKey:CoachID" json:"coach"`
	Schedules       []ClassSchedule `gorm:"foreignKey:ClassID" json:"schedule"`
	CreatedAt       time.Time       `json:"created_at"`
	UpdatedAt       time.Time       `json:"updated_at"`
}

// ClassSchedule
type ClassSchedule struct {
	ID        string    `gorm:"primaryKey" json:"id"`
	ClassID   string    `json:"class_id"`
	DayOfWeek int       `json:"day_of_week"` // 0=Sun, 1=Mon, ..., 6=Sat
	StartTime string    `json:"start_time"`  // "17:00"
	EndTime   string    `json:"end_time"`    // "18:30"
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// Student details
type Student struct {
	ID          string    `gorm:"primaryKey" json:"id"`
	Name        string    `json:"name"`
	Avatar      string    `json:"avatar"`
	BirthYear   int       `json:"birth_year"`
	ParentName  string    `json:"parent_name"`
	ParentPhone string    `json:"parent_phone"`
	ClassID     string    `json:"class_id"`
	ClassName   string    `gorm:"-" json:"class_name"` // Computed field
	Status      string    `json:"status"` // "active", "trial", "inactive", "dropped"
	JoinDate    time.Time `json:"join_date"`
	Notes       string    `json:"notes"`
	Class       Class     `gorm:"foreignKey:ClassID" json:"class,omitempty"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

// TrialRegistration
type TrialRegistration struct {
	ID               string    `gorm:"primaryKey" json:"id"`
	ParentName       string    `json:"parent_name"`
	ParentPhone      string    `json:"parent_phone"`
	StudentName      string    `json:"student_name"`
	StudentBirthYear int       `json:"student_birth_year"`
	PreferredCourt   string    `json:"preferred_court"`
	Notes            string    `json:"notes"`
	Status           string    `json:"status"` // "pending", "approved", "rejected", "converted"
	CreatedAt        time.Time `json:"created_at"`
	UpdatedAt        time.Time `json:"updated_at"`
}

// AttendanceRecord
type AttendanceRecord struct {
	ID        string    `gorm:"primaryKey" json:"id"`
	StudentID string    `json:"student_id"`
	ClassID   string    `json:"class_id"`
	Date      time.Time `gorm:"type:date" json:"date"`
	Status    string    `json:"status"` // "present", "cancelled", "excused", "unexcused", "dropped"
	MarkedBy  string    `json:"marked_by"` // user/coach ID
	Student   Student   `gorm:"foreignKey:StudentID" json:"student,omitempty"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// Tournament
type Tournament struct {
	ID          string     `gorm:"primaryKey" json:"id"`
	Name        string     `json:"name"`
	Description string     `json:"description"`
	Date        time.Time  `json:"date"`
	EndDate     *time.Time `json:"end_date"`
	Location    string     `json:"location"`
	Banner      string     `json:"banner"`
	Status      string     `json:"status"` // "upcoming", "ongoing", "completed"
	CreatedAt   time.Time  `json:"created_at"`
	UpdatedAt   time.Time  `json:"updated_at"`
}

// TuitionPlan
type TuitionPlan struct {
	ID              string    `gorm:"primaryKey" json:"id"`
	Name            string    `json:"name"`
	Price           float64   `json:"price"`
	Duration        string    `json:"duration"`
	SessionsPerWeek int       `json:"sessions_per_week"`
	Features        []string  `gorm:"serializer:json" json:"features"`
	IsPopular       bool      `json:"is_popular"`
	CreatedAt       time.Time `json:"created_at"`
	UpdatedAt       time.Time `json:"updated_at"`
}

// FAQ
type FAQ struct {
	ID        string    `gorm:"primaryKey" json:"id"`
	Question  string    `json:"question"`
	Answer    string    `json:"answer"`
	Category  string    `json:"category"`
	Order     int       `json:"order"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// Review
type Review struct {
	ID          string    `gorm:"primaryKey" json:"id"`
	ParentName  string    `json:"parent_name"`
	Avatar      string    `json:"avatar"`
	Rating      int       `json:"rating"`
	Content     string    `json:"content"`
	StudentName string    `json:"student_name"`
	IsVisible   bool      `json:"is_visible"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

// Banner
type Banner struct {
	ID        string    `gorm:"primaryKey" json:"id"`
	Image     string    `json:"image"`
	Title     string    `json:"title"`
	Subtitle  string    `json:"subtitle"`
	CTAText   string    `json:"cta_text"`
	CTALink   string    `json:"cta_link"`
	Order     int       `json:"order"`
	IsActive  bool      `json:"is_active"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// Invoice
type Invoice struct {
	ID        string    `gorm:"primaryKey" json:"id"`
	StudentID string    `json:"student_id"`
	Amount    float64   `json:"amount"`
	Month     int       `json:"month"`
	Year      int       `json:"year"`
	Status    string    `json:"status"` // "unpaid", "paid", "overdue"
	DueDate   time.Time `gorm:"type:date" json:"due_date"`
	Student   Student   `gorm:"foreignKey:StudentID" json:"student,omitempty"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// Payment
type Payment struct {
	ID        string    `gorm:"primaryKey" json:"id"`
	InvoiceID string    `json:"invoice_id"`
	Amount    float64   `json:"amount"`
	Method    string    `json:"method"` // "cash", "transfer"
	Note      string    `json:"note"`
	Invoice   Invoice   `gorm:"foreignKey:InvoiceID" json:"invoice,omitempty"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// Contact represents a contact form submission
type Contact struct {
	ID        string    `gorm:"primaryKey" json:"id"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	Phone     string    `json:"phone"`
	Subject   string    `json:"subject"`
	Message   string    `json:"message"`
	IsRead    bool      `json:"is_read"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
