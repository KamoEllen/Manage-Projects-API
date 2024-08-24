package main

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
)

type Project struct {
	ID          int       `json:"id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	DueDate     time.Time `json:"due_date"`
	Status      string    `json:"status"` // e.g., "Not Started", "In Progress", "Completed"
}

func main() {
	db, err := sql.Open("postgres", os.Getenv("DATABASE_URL"))
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// Create the table if it doesn't exist
	_, err = db.Exec(`
		CREATE TABLE IF NOT EXISTS projects (
			id SERIAL PRIMARY KEY,
			name TEXT NOT NULL,
			description TEXT,
			due_date TIMESTAMP,
			status TEXT
		)
	`)
	if err != nil {
		log.Fatal(err)
	}

	// Create router
	router := mux.NewRouter()
	router.HandleFunc("/projects", getProjects(db)).Methods("GET")
	router.HandleFunc("/projects/{id}", getProject(db)).Methods("GET")
	router.HandleFunc("/projects", createProject(db)).Methods("POST")
	router.HandleFunc("/projects/{id}", updateProject(db)).Methods("PUT")
	router.HandleFunc("/projects/{id}", deleteProject(db)).Methods("DELETE")

	// Start server
	log.Fatal(http.ListenAndServe(":8000", jsonContentTypeMiddleware(router)))
}

func jsonContentTypeMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		next.ServeHTTP(w, r)
	})
}

// get all projects
func getProjects(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		rows, err := db.Query("SELECT * FROM projects")
		if err != nil {
			log.Fatal(err)
		}
		defer rows.Close()

		projects := []Project{}
		for rows.Next() {
			var p Project
			if err := rows.Scan(&p.ID, &p.Name, &p.Description, &p.DueDate, &p.Status); err != nil {
				log.Fatal(err)
			}
			projects = append(projects, p)
		}
		if err := rows.Err(); err != nil {
			log.Fatal(err)
		}

		json.NewEncoder(w).Encode(projects)
	}
}

// get project by id
func getProject(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		id := vars["id"]

		var p Project
		err := db.QueryRow("SELECT * FROM projects WHERE id = $1", id).Scan(&p.ID, &p.Name, &p.Description, &p.DueDate, &p.Status)
		if err != nil {
			w.WriteHeader(http.StatusNotFound)
			return
		}

		json.NewEncoder(w).Encode(p)
	}
}

// create project
func createProject(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var p Project
		if err := json.NewDecoder(r.Body).Decode(&p); err != nil {
			http.Error(w, "Bad Request", http.StatusBadRequest)
			return
		}

		err := db.QueryRow("INSERT INTO projects (name, description, due_date, status) VALUES ($1, $2, $3, $4) RETURNING id", p.Name, p.Description, p.DueDate, p.Status).Scan(&p.ID)
		if err != nil {
			log.Fatal(err)
		}

		json.NewEncoder(w).Encode(p)
	}
}

// update project
func updateProject(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var p Project
		if err := json.NewDecoder(r.Body).Decode(&p); err != nil {
			http.Error(w, "Bad Request", http.StatusBadRequest)
			return
		}

		vars := mux.Vars(r)
		id := vars["id"]

		_, err := db.Exec("UPDATE projects SET name = $1, description = $2, due_date = $3, status = $4 WHERE id = $5", p.Name, p.Description, p.DueDate, p.Status, id)
		if err != nil {
			log.Fatal(err)
		}

		json.NewEncoder(w).Encode(p)
	}
}

// delete project
func deleteProject(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		id := vars["id"]

		var p Project
		err := db.QueryRow("SELECT * FROM projects WHERE id = $1", id).Scan(&p.ID, &p.Name, &p.Description, &p.DueDate, &p.Status)
		if err != nil {
			w.WriteHeader(http.StatusNotFound)
			return
		}

		_, err = db.Exec("DELETE FROM projects WHERE id = $1", id)
		if err != nil {
			w.WriteHeader(http.StatusNotFound)
			return
		}

		json.NewEncoder(w).Encode("Project deleted")
	}
}
