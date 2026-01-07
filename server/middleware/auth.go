package middleware
import (
	"context"
	"net/http"
	"strings"
	"os"
	"github.com/golang-jwt/jwt/v5"
)
type contextKey string 
const UserIDKey contextKey = "userID"
func AuthMiddleware(next http.HandlerFunc) http.HandlerFunc{
	return func(w http.ResponseWriter, r* http.Request){
		authHeader := r.Header.Get("Authorization")
		if authHeader == "" {
			http.Error(w,"Authorization header missing",http.StatusUnauthorized)
			return
		}
		parts := strings.Split(authHeader," ")
		if len(parts) != 2 || parts[0] != "Bearer"{
			w.WriteHeader(http.StatusUnauthorized)
			return 
		}
		tokenString := parts[1]
		token , err := jwt.Parse(tokenString,func(token *jwt.Token)(interface{},error){
			return []byte(os.Getenv("JWT_SECRET")),nil
		})
		if err != nil || !token.Valid {
			http.Error (w,"Invalid token",http.StatusUnauthorized)
			return
		}
		claims, ok := token.Claims.(jwt.MapClaims)
		if !ok {
			http.Error(w,"Invalid token claims",http.StatusUnauthorized)
			return
		}
		userID, ok := claims["user_id"].(string)
		if !ok {
			http.Error (w,"Invalid user ID in token",http.StatusUnauthorized)
			return
		}
		ctx := context.WithValue(r.Context(),UserIDKey,userID)
		next.ServeHTTP (w,r.WithContext (ctx))

	}
}