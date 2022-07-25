package http

import (
	"ProdReadyRestAPI1.0/internal/comment"
	"encoding/json"
	"errors"
	"fmt"
	jwt "github.com/dgrijalva/jwt-go"
	"github.com/gorilla/mux"
	log "github.com/sirupsen/logrus"
	"net/http"
	"strconv"
	"strings"
)

//Handler - stores pointer to our comments service
type Handler struct {
	Router  *mux.Router
	Service *comment.Service
}

// Response - an object to store responses from our API
type Response struct {
	Message string
	Error   string
}

// NewHandler - returns a pointer to a Handler
func NewHandler(service *comment.Service) *Handler {
	return &Handler{
		Service: service,
	}
}

//LoggingMiddleware - adds middleware around endpoints
func LoggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.WithFields(
			log.Fields{
				"Method": r.Method,
				"Path":   r.URL.Path,
			}).Info("Handled Request")
		next.ServeHTTP(w, r)
	})
}

// BasicAuth - a middlware function that will provide basic auth around specific endpoints
func BasicAuth(original func(w http.ResponseWriter, r *http.Request)) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Info("Basic Auth endpoint Hit")
		user, pass, ok := r.BasicAuth()
		if user == "admin" && pass == "password" && ok {
			original(w, r)
		} else {
			sendErrorResponse(w, "not authorized", errors.New("not authorized"))
		}
	}
}

func validateToken(accessToken string) bool {
	var mySigningKey = []byte("missionimpossible")
	token, err := jwt.Parse(accessToken, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("there has been an error")
		}
		return mySigningKey, nil
	})
	if err != nil {
		return false
	}
	return token.Valid
}

//JWTAuth - A decorator function for JWT validation for endpoints
func JWTAuth(original func(w http.ResponseWriter, r *http.Request)) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Info("jwt Auth hit")
		authHeader := r.Header["Authorization"]
		if authHeader == nil {
			sendErrorResponse(w, "not authorized", errors.New("not authorized"))
			return
		}
		//Bearer jwt-token
		authHeaderParts := strings.Split(authHeader[0], " ")
		if len(authHeaderParts) != 2 || strings.ToLower(authHeaderParts[0]) != "bearer" {
			sendErrorResponse(w, "not authorized", errors.New("not authorized"))
			return
		}
		if validateToken(authHeaderParts[1]) {
			original(w, r)
		} else {
			sendErrorResponse(w, "not authorized", errors.New("not authorized"))
			return
		}
	}
}

//SetupRoutes - Sets up all the routes for our application
func (h *Handler) SetupRoutes() {
	log.Info("Setting up Routes")
	h.Router = mux.NewRouter()
	h.Router.Use(LoggingMiddleware)

	h.Router.HandleFunc("/api/comment", JWTAuth(h.GetAllComments)).Methods("GET")
	h.Router.HandleFunc("/api/comment", JWTAuth(h.PostComment)).Methods("POST")

	h.Router.HandleFunc("/api/comment/{id}", JWTAuth(h.GetComment)).Methods("GET")
	h.Router.HandleFunc("/api/comment/{id}", JWTAuth(h.DeleteComment)).Methods("DELETE")
	h.Router.HandleFunc("/api/comment/{id}", JWTAuth(h.UpdateComment)).Methods("PUT")

	h.Router.HandleFunc("/api/health", func(w http.ResponseWriter, r *http.Request) {
		if err := sendOkResponse(w, Response{Message: "Successful Health Check"}); err != nil {
			panic(err)
		}
		err := json.NewEncoder(w).Encode(Response{Message: "I am Alive"})
		if err != nil {
			panic(err)
		}
	})
}

// GetComment - retrieve a comment by ID
func (h *Handler) GetComment(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	i, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		sendErrorResponse(w, "Unable to parse UINT from ID", err)
		return
	}

	comment, err := h.Service.GetComment(uint(i))
	if err != nil {
		sendErrorResponse(w, "Error Retrieving Comment by ID", err)
		return
	}

	if err := sendOkResponse(w, Response{Message: "Successfully Retrieved Comment"}); err != nil {
		panic(err)
	}
	err = json.NewEncoder(w).Encode(comment)
	if err != nil {
		panic(err)
	}
}

// GetAllComments - Retrieves all comments from the comment service
func (h *Handler) GetAllComments(w http.ResponseWriter, r *http.Request) {
	comments, err := h.Service.GetAllComments()
	if err != nil {
		sendErrorResponse(w, "Error to  retrieve all comments", err)
		return
	}
	if err := sendOkResponse(w, Response{Message: "Successfully Retrieved All Comments"}); err != nil {
		panic(err)
	}
	err = json.NewEncoder(w).Encode(comments)
	if err != nil {
		panic(err)
	}
}

// PostComment - adds a new comment
func (h *Handler) PostComment(w http.ResponseWriter, r *http.Request) {
	var comment comment.Comment
	err := json.NewDecoder(r.Body).Decode(&comment)
	if err != nil {
		sendErrorResponse(w, "Failed to decode JSON Body", err)
		return
	}

	comment, err = h.Service.PostComment(comment)
	if err != nil {
		sendErrorResponse(w, "Failed to post new comment", err)
		return
	}
	if err := sendOkResponse(w, Response{Message: "Successfully Posted Comment"}); err != nil {
		panic(err)
	}
	err = json.NewEncoder(w).Encode(comment)
	if err != nil {
		panic(err)
	}
}

// UpdateComment - Updates comment by ID
func (h *Handler) UpdateComment(w http.ResponseWriter, r *http.Request) {
	var comment comment.Comment
	err := json.NewDecoder(r.Body).Decode(&comment)
	if err != nil {
		sendErrorResponse(w, "Failed to decode JSON Body", err)
		return
	}

	vars := mux.Vars(r)
	id := vars["id"]
	commentID, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		sendErrorResponse(w, "Failed to parse uint from ID", err)
		return
	}

	comment, err = h.Service.UpdateComment(uint(commentID), comment)
	if err != nil {
		sendErrorResponse(w, "Failed to update comment", err)
		return
	}
	if err := sendOkResponse(w, Response{Message: "Successfully Updated Comment"}); err != nil {
		panic(err)
	}
	err = json.NewEncoder(w).Encode(comment)
	if err != nil {
		panic(err)
	}
}

// DeleteComment - deletes comment by ID
func (h *Handler) DeleteComment(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	commentID, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		sendErrorResponse(w, "Failed to parse uint from ID", err)
		return
	}

	if err := h.Service.DeleteComment(uint(commentID)); err != nil {
		sendErrorResponse(w, "Failed to delete comment by comment ID", err)
		return
	}
	if err := sendOkResponse(w, Response{Message: "Successfully Deleted Comment"}); err != nil {
		panic(err)
	}
	if err := json.NewEncoder(w).Encode(Response{Message: "Comment Successfully Deleted"}); err != nil {
		panic(err)
	}
}

func sendOkResponse(w http.ResponseWriter, resp interface{}) error {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	return json.NewEncoder(w).Encode(resp)
}

func sendErrorResponse(w http.ResponseWriter, message string, err error) {
	w.WriteHeader(http.StatusInternalServerError)
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")

	if err := json.NewEncoder(w).Encode(Response{Message: message, Error: err.Error()}); err != nil {
		log.Error(err)
	}
}
