package main

import (
	"context"
	"encoding/json"
	"net/http"
	"time"

	"github.com/google/uuid"
	"github.com/shri771/Go/WebServer/internal/auth"
	"github.com/shri771/Go/WebServer/internal/database"
)

type User struct {
	ID           uuid.UUID `json:"id"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
	Email        string    `json:"email"`
	Token        string    `json:"token"`
	RefreshToken string    `json:"refresh_token"`
}

func (cfg *apiConfig) handlerUsers(w http.ResponseWriter, r *http.Request) {
	type users struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	type response struct {
		User
	}
	//Decode Json
	user := users{}
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Erro with Decoding:", err)
		return
	}
	defer r.Body.Close()

	if r.Method == "PUT" {

		// Hash Password and Add User to Database
		hashedPassword, err := auth.HashPassword(user.Password)
		if err != nil {
			respondWithError(w, http.StatusUnauthorized, "Error with hashing pswd", err)
			return
		}

		userID, err := cfg.AuthorizeUser(w, r.Header)
		if err != nil {
			respondWithError(w, http.StatusUnauthorized, "Could not retrive token from auth header", err)
			return
		}

		updatedUser, err := cfg.db.UpdateUser(context.Background(), database.UpdateUserParams{
			Email:          user.Email,
			HashedPassword: hashedPassword,
			ID:             userID,
		})

		respondWithJSON(w, http.StatusOK, response{
			User: User{
				ID:        updatedUser.ID,
				CreatedAt: updatedUser.CreatedAt,
				UpdatedAt: updatedUser.UpdatedAt,
				Email:     updatedUser.Email,
			},
		})

		return

	}
	// Hash Password and Add User to Database
	hashedPassword, err := auth.HashPassword(user.Password)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Error with hashing pswd", err)
		return
	}
	addedUser, err := cfg.db.CreateUser(context.Background(), database.CreateUserParams{
		CreatedAt:      time.Now().UTC(),
		UpdatedAt:      time.Now().UTC(),
		Email:          user.Email,
		HashedPassword: hashedPassword,
	})
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Error with Adding user:", err)
		return
	}

	// Return Json
	respondWithJSON(w, http.StatusCreated, response{
		User: User{
			ID:        addedUser.ID,
			CreatedAt: addedUser.CreatedAt,
			UpdatedAt: addedUser.UpdatedAt,
			Email:     addedUser.Email,
		},
	})
}

func (cfg *apiConfig) handlerLogin(w http.ResponseWriter, r *http.Request) {
	type users struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	// Decode Data
	loginCredintial := users{}
	err := json.NewDecoder(r.Body).Decode(&loginCredintial)
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Error with decoding json", err)
		return
	}
	defer r.Body.Close()

	// Query database
	login, err := cfg.db.GetUserByEmail(context.Background(), loginCredintial.Email)
	if err != nil {
		respondWithError(w, http.StatusUnauthorized, "Error while quering database", err)
		return
	}

	// Verify user
	err = auth.CheckPasswordHash(loginCredintial.Password, login.HashedPassword)
	if err != nil {
		respondWithError(w, http.StatusUnauthorized, "Error while quering database", err)
		return
	}

	// Create JWT
	token, err := auth.MakeJWT(login.ID, cfg.secret, 1*time.Hour)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Error while Genarating JWT token", err)
	}

	// Create and Store Refreh token
	refreshToken, err := auth.MakeRefreshToken()
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Could not genrate refrest tokens", err)
		return
	}

	refreshTokenData, err := cfg.db.CreateRefreshToken(context.Background(), database.CreateRefreshTokenParams{
		Token:     refreshToken,
		UserID:    login.ID,
		ExpiresAt: time.Now().UTC().Add(time.Hour * 24 * 60),
	})

	respondWithJSON(w, http.StatusOK, User{
		ID:           login.ID,
		CreatedAt:    login.CreatedAt,
		UpdatedAt:    login.UpdatedAt,
		Email:        login.Email,
		Token:        token,
		RefreshToken: refreshTokenData.Token,
	})

}

// Helpers
func (cfg *apiConfig) AuthorizeUser(w http.ResponseWriter, header http.Header) (uuid.UUID, error) {
	// Valdidate JWT
	accessToken, err := auth.GetBearerToken(header)
	if err != nil {
		return uuid.Nil, err
	}

	userID, err := auth.ValidateJWT(accessToken, cfg.secret)
	if err != nil {
		return uuid.Nil, err
	}

	return userID, nil
}
