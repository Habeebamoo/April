package tools

import (
	"encoding/json"
	"fmt"
	"strings"
	"time"

	"github.com/Habeebamoo/April/server/db"
)

// ==================== STRUCTS ====================

type User struct {
	UserId     string    `gorm:"primaryKey"`
	Name       string
	Email      string
	Role       string
	Verified   bool
	IsBanned   bool
	CreatedAt  time.Time
}

type Profile struct {
	UserId      string  `gorm:"primaryKey"`
	Username    string
	Bio         string
	Picture     string
	ProfileUrl  string
	Website     string
	Following   int
	Followers   int
}

type Article struct {
	ArticleId  string          
	AuthorId   string
	Title      string
	Content    json.RawMessage 
	Picture    string          
	ReadTime   string          
	Slug       string         
	CreatedAt  time.Time    
}
type Comment struct {
	CommentId  string 
	ArticleId  string  
	UserId     string 
	ReplyId    string  
	Replys     int    
	Content    string 
}

type Subscriber struct {
	SubscriberId  string 
	Email         string 
}

type Follow struct {
	FollowerId   string  
	FollowingId  string 
}


// ==================== TOOLS ====================

// 1. Get clivo users count
func QueryClivoUsers() string {
	var count int64
	result := db.ClivoDB.Table("users").Count(&count)
	if result.Error != nil {
		return "Failed to query Clivo users"
	}
	return fmt.Sprintf("%d", count)
}

// 2. Get recent articles
func GetRecentArticles(input map[string]any) string {
	limit := 5
	if l, ok := input["limit"].(float64); ok {
		limit = int(l)
	}

	var articles []Article
	result := db.ClivoDB.Order("created_at desc").Limit(limit).Find(&articles)
	if result.Error != nil {
		return "Failed to fetch recent articles"
	}

	if len(articles) == 0 {
		return "No articles found"
	}

	var lines []string
	for _, a := range articles {
		lines = append(lines, fmt.Sprintf("- [%d] %s (%s)", a.ArticleId, a.Title, a.CreatedAt.Format("Jan 2, 2006")))
	}
	return fmt.Sprintf("Recent %d articles:\n%s", len(articles), strings.Join(lines, "\n"))
}

// 3. Get recent signups
func GetRecentSignups(input map[string]any) string {
	limit := 5
	period := "week"

	if l, ok := input["limit"].(float64); ok {
		limit = int(l)
	}
	if p, ok := input["period"].(string); ok {
		period = p
	}

	var since time.Time
	switch period {
	case "today":
		since = time.Now().Truncate(24 * time.Hour)
	case "month":
		since = time.Now().AddDate(0, -1, 0)
	default: // week
		since = time.Now().AddDate(0, 0, -7)
	}

	var users []User
	result := db.ClivoDB.
		Where("created_at >= ?", since).
		Order("created_at desc").
		Limit(limit).
		Find(&users)

	if result.Error != nil {
		return "Failed to fetch recent signups"
	}

	if len(users) == 0 {
		return fmt.Sprintf("No signups in the last %s", period)
	}

	var lines []string
	for _, u := range users {
		lines = append(lines, fmt.Sprintf("- %s (%s) — joined %s", u.Name, u.Email, u.CreatedAt.Format("Jan 2")))
	}
	return fmt.Sprintf("%d signups this %s:\n%s", len(users), period, strings.Join(lines, "\n"))
}

// 4. Find user by name, email or username
func FindUser(input map[string]any) string {
	query, ok := input["query"].(string)
	if !ok || query == "" {
		return "Please provide a name, email or username to search"
	}

	search := "%" + strings.ToLower(query) + "%"

	type Result struct {
		UserID   uint
		Name     string
		Email    string
		Username string
		Verified bool
		Banned   bool
	}

	var results []Result
	db.ClivoDB.Raw(`
		SELECT u.user_id, u.name, u.email, p.username, u.verified, u.is_banned
		FROM users u
		LEFT JOIN profiles p ON p.user_id = u.user_id
		WHERE LOWER(u.name) LIKE ?
		   OR LOWER(u.email) LIKE ?
		   OR LOWER(p.username) LIKE ?
		LIMIT 5
	`, search, search, search).Scan(&results)

	if len(results) == 0 {
		return fmt.Sprintf("No user found matching '%s'", query)
	}

	var lines []string
	for _, r := range results {
		status := "active"
		if r.Banned {
			status = "banned"
		} else if !r.Verified {
			status = "unverified"
		}
		lines = append(lines, fmt.Sprintf(
			"- %s (@%s) | %s | %s",
			r.Name, r.Username, r.Email, status,
		))
	}
	return fmt.Sprintf("Found %d user(s):\n%s", len(results), strings.Join(lines, "\n"))
}

// 5. Create a comment on an article
func CreateComment(input map[string]any) string {
	articleID, ok1 := input["article_id"].(string)
	content, ok2 := input["content"].(string)

	if !ok1 || !ok2 || content == "" {
		return "article_id and content are required"
	}

	comment := Comment{
		ArticleId: articleID,
		Content:   content,
	}

	result := db.ClivoDB.Table("comments").Create(&comment)
	if result.Error != nil {
		return fmt.Sprintf("Failed to create comment: %s", result.Error.Error())
	}

	return fmt.Sprintf("Comment posted on article")
}

// 6. Get subscribers count
func GetSubscribersCount() string {
	var count int64
	result := db.ClivoDB.Table("subscribers").Count(&count)
	if result.Error != nil {
		return "Failed to get subscribers count"
	}
	return fmt.Sprintf("%d", count)
}

// 7. Get following/followers count of a user
func GetUserSocialStats(input map[string]any) string {
	username, ok := input["username"].(string)
	if !ok || username == "" {
		return "Please provide a username"
	}

	// Find the profile
	var profile Profile
	result := db.ClivoDB.
		Where("LOWER(username) = ?", strings.ToLower(username)).
		First(&profile)

	if result.Error != nil {
		return fmt.Sprintf("User @%s not found", username)
	}

	// Count following
	var following int64
	db.ClivoDB.Table("follows").
		Where("follower_id = ?", profile.UserId).
		Count(&following)

	// Count followers
	var followers int64
	db.ClivoDB.Table("follows").
		Where("following_id = ?", profile.UserId).
		Count(&followers)

	return fmt.Sprintf(
		"@%s — %d followers | %d following",
		username, followers, following,
	)
}