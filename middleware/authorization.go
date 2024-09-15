package middleware

import (
	"context"
	"net/http"
	"strings"
)

// ContextKey es un tipo personalizado para claves de contexto
type ContextKey string

// Clave personalizada para el userID en el contexto
const UserIDKey ContextKey = "userID"

// AdminMiddleware es un middleware que verifica si el usuario es un administrador
func AdminMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		roles := r.Header.Get("Roles")
		if !strings.Contains(roles, "admin") {
			http.Error(w, "Acceso no autorizado", http.StatusForbidden)
			return
		}
		next.ServeHTTP(w, r)
	})
}

// AuthMiddleware es un middleware que verifica si el usuario está autenticado
func AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		token := r.Header.Get("Authorization")
		if token == "" {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}

		// Validación del token (esto es una simulación, debes implementar una validación real)
		if !strings.HasPrefix(token, "Bearer ") || token != "Bearer valid-token" {
			http.Error(w, "Invalid token", http.StatusUnauthorized)
			return
		}

		// Simulamos obtener el userID del token
		userID := "1234" // Este debería ser derivado del token
		ctx := context.WithValue(r.Context(), UserIDKey, userID)

		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
