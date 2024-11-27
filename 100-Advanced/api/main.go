package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type Project struct {
	ID          int      `json:"id"`
	Name        string   `json:"name"`
	Description string   `json:"description"`
	Columns     []Column `json:"columns"`
}

type Column struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	ProjectID int   `json:"project_id"`
	Color    string `json:"color,omitempty"`
	Cards    []Card `json:"cards"`
}

type Card struct {
	ID          int    `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	ColumnID    int    `json:"column_id"`
	Position    int    `json:"position"`
}

func enableCORS(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "http://localhost:5173")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
		w.Header().Set("Access-Control-Allow-Credentials", "true")

		if r.Method == http.MethodOptions {
			w.WriteHeader(http.StatusNoContent)
			return
		}

		next.ServeHTTP(w, r)
	})
}

func main() {
	InitDB()
	defer Db.Close()

	mux := http.NewServeMux()

	mux.HandleFunc("/projects", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			getProjects(w, r)
		case http.MethodPost:
			addProject(w, r)
		case http.MethodPut:
			updateProject(w, r)
		case http.MethodDelete:
			deleteProject(w, r)
		default:
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}
	})

	mux.HandleFunc("/columns", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			getColumns(w, r)
		case http.MethodPost:
			addColumn(w, r)
		case http.MethodPut:
			updateColumnName(w, r)
		case http.MethodDelete:
			deleteColumn(w, r)
		default:
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}
	})

	mux.HandleFunc("/cards", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			getCards(w, r)
		case http.MethodPost:
			addCardToColumn(w, r)
		case http.MethodPut:
			editCardInColumn(w, r)
		case http.MethodDelete:
			deleteCardFromColumn(w, r)
		default:
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}
	})

	mux.HandleFunc("/cards/move", func(w http.ResponseWriter, r *http.Request){
		if r.Method == http.MethodPost {
			moveCard(w, r)
		}else{
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}
	})

	mux.HandleFunc("/updateCardPositions", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodPost {
			updateCardPositions(w, r)
		} else {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}
	})

	// Envolver el servidor con CORS
	log.Println("Server running on http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", enableCORS(mux)))
}

// PROJECTS //
func getProjects(w http.ResponseWriter, r *http.Request){
	rows, err := Db.Query("SELECT id, name, description FROM projects")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var projects []Project
	for rows.Next() {
		var project Project
		if err := rows.Scan(&project.ID, &project.Name, &project.Description); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		project.Columns = []Column{}
		projects = append(projects, project)
	}

	if err := rows.Err(); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(projects)
}

func addProject(w http.ResponseWriter, r *http.Request) {
	var newProject Project
	if err := json.NewDecoder(r.Body).Decode(&newProject); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	result, err := Db.Exec("INSERT INTO projects (name, description) VALUES (?, ?)", newProject.Name, newProject.Description)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	id, err := result.LastInsertId()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	newProject.ID = int(id)
	newProject.Columns = []Column{}
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(newProject)
}

func updateProject(w http.ResponseWriter, r *http.Request){
	projectID := r.URL.Query().Get("id")
	if projectID == "" {
		http.Error(w, "Project ID is required", http.StatusBadRequest)
		return
	}

	var updatedProject Project
	if err := json.NewDecoder(r.Body).Decode(&updatedProject); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	_, err := Db.Exec("UPDATE projects SET name = ?, description = ? WHERE id = ?", updatedProject.Name, updatedProject.Description, projectID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func deleteProject(w http.ResponseWriter, r *http.Request){
	projectID := r.URL.Query().Get("id")
	if projectID == "" {
		http.Error(w, "Project ID is required", http.StatusBadRequest)
		return
	}

	_, err := Db.Exec("DELETE FROM projects WHERE id = ?", projectID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

// COLUMNS //

func getColumns(w http.ResponseWriter, r *http.Request){
	projectID := r.URL.Query().Get("project_id")
	if projectID == "" {
		http.Error(w, "Project ID is required", http.StatusBadRequest)
		return
	}

	rows, err := Db.Query("SELECT id, name FROM columns WHERE project_id = ?", projectID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var columns []Column
	for rows.Next() {
		var column Column
		if err := rows.Scan(&column.ID, &column.Name); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		columns = append(columns, column)
	}

	if err := rows.Err(); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if columns == nil {
		columns = []Column{}
	}

	json.NewEncoder(w).Encode(columns)

}

func addColumn(w http.ResponseWriter, r *http.Request) {
	var newColumn Column
	if err := json.NewDecoder(r.Body).Decode(&newColumn); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	result, err := Db.Exec("INSERT INTO columns (project_id, name) VALUES (?, ?)", newColumn.ProjectID, newColumn.Name)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	id, err := result.LastInsertId()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	newColumn.ID = int(id)
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(newColumn)

}

func updateColumnName(w http.ResponseWriter, r *http.Request) {
	// Obtener el ID de la columna de la query
	columnID := r.URL.Query().Get("id")
	if columnID == "" {
		http.Error(w, "Column ID is required", http.StatusBadRequest)
		return
	}

	// Decodificar los datos enviados en el body
	var updatedData Column
	if err := json.NewDecoder(r.Body).Decode(&updatedData); err != nil {
		http.Error(w, "Invalid JSON data: "+err.Error(), http.StatusBadRequest)
		return
	}

	// Validar que el nombre no esté vacío
	if updatedData.Name == "" {
		http.Error(w, "Column name cannot be empty", http.StatusBadRequest)
		return
	}

	// Ejecutar la consulta SQL para actualizar el nombre
	result, err := Db.Exec("UPDATE columns SET name = ? WHERE id = ?", updatedData.Name, columnID)
	if err != nil {
		http.Error(w, "Error updating column: "+err.Error(), http.StatusInternalServerError)
		return
	}

	// Confirmar si se actualizó alguna fila
	rowsAffected, err := result.RowsAffected()
	fmt.Printf("Filas afectadas en la columna origen: %d\n", rowsAffected)
	if err != nil {
		http.Error(w, "Could not determine affected rows: "+err.Error(), http.StatusInternalServerError)
		return
	}
	if rowsAffected == 0 {
		http.Error(w, "No column found with the given ID", http.StatusNotFound)
		return
	}

	// Responder con éxito
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Column updated successfully"))
}


func deleteColumn(w http.ResponseWriter, r *http.Request){
	columnID := r.URL.Query().Get("id")
	if columnID == "" {
		http.Error(w, "column ID is required", http.StatusBadRequest)
		return
	}

	_, err := Db.Exec("DELETE FROM columns WHERE id = ?", columnID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

// CARDS //
func getCards(w http.ResponseWriter, r *http.Request) {
	columnID := r.URL.Query().Get("column_id")
	if columnID == "" {
		http.Error(w, "column ID is required", http.StatusBadRequest)
		return
	}

	rows, err := Db.Query("SELECT id, title, description, position FROM cards WHERE column_id = ? ORDER BY position", columnID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var cards []Card
	for rows.Next() {
		var card Card
		if err := rows.Scan(&card.ID, &card.Title, &card.Description, &card.Position); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		cards = append(cards, card)
	}

	if err := rows.Err(); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if cards == nil {
		cards = []Card{}
	}

	json.NewEncoder(w).Encode(cards)
}

func updateCardPositions(w http.ResponseWriter, r *http.Request) {
	var updatedCards []Card
	if err := json.NewDecoder(r.Body).Decode(&updatedCards); err != nil {
		http.Error(w, "Invalid JSON payload", http.StatusBadRequest)
		return
	}

	for _, card := range updatedCards {
		_, err := Db.Exec("UPDATE cards SET position = ? WHERE id = ?", card.Position, card.ID)
		if err != nil {
			http.Error(w, "Failed to update card positions", http.StatusInternalServerError)
			return
		}
	}

	w.WriteHeader(http.StatusOK)
}


func addCardToColumn(w http.ResponseWriter, r *http.Request) {
	var newCard Card
	if err := json.NewDecoder(r.Body).Decode(&newCard); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	result, err := Db.Exec("INSERT INTO cards (column_id, title, description) VALUES (?, ?, ?)", newCard.ColumnID, newCard.Title, newCard.Description)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	id, err := result.LastInsertId()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	newCard.ID = int(id)
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(newCard)
}

func editCardInColumn(w http.ResponseWriter, r *http.Request) {
    cardID := r.URL.Query().Get("id")
    if cardID == "" {
        http.Error(w, "Missing id parameter", http.StatusBadRequest)
        return
    }

    var updatedCard struct {
        Title       *string `json:"title,omitempty"`
        Description *string `json:"description,omitempty"`
    }
    if err := json.NewDecoder(r.Body).Decode(&updatedCard); err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }

    if updatedCard.Title == nil && updatedCard.Description == nil {
        http.Error(w, "No fields to update", http.StatusBadRequest)
        return
    }

    query := "UPDATE cards SET"
    params := []interface{}{}

    if updatedCard.Title != nil {
        query += " title = ?"
        params = append(params, *updatedCard.Title)
    }
    if updatedCard.Description != nil {
        if len(params) > 0 {
            query += ","
        }
        query += " description = ?"
        params = append(params, *updatedCard.Description)
    }
    query += " WHERE id = ?"
    params = append(params, cardID)

    res, err := Db.Exec(query, params...)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    rowsAffected, _ := res.RowsAffected()
    if rowsAffected == 0 {
        http.Error(w, "No card found with the given ID", http.StatusNotFound)
        return
    }

    w.WriteHeader(http.StatusOK)
}

func deleteCardFromColumn(w http.ResponseWriter, r *http.Request){
	cardID := r.URL.Query().Get("id")
	if cardID == "" {
		http.Error(w, "Missing id parameter", http.StatusBadRequest)
		return
	}

	_, err := Db.Exec("DELETE FROM cards WHERE id = ?", cardID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
func moveCard(w http.ResponseWriter, r *http.Request) {
	// Parsear el cuerpo de la solicitud
	var params struct {
		FromColumnID int `json:"fromColumnId"`
		CardID       int `json:"cardId"`
		ToColumnID   int `json:"toColumnId"`
		NewPosition  int `json:"newPosition"`
		OldPosition  int `json:"oldPosition"`
	}

	// Log para verificar los datos recibidos
	log.Println("Recibiendo datos para mover tarjeta...")

	err := json.NewDecoder(r.Body).Decode(&params)
	if err != nil {
		log.Printf("Error al decodificar el cuerpo de la solicitud: %v\n", err)
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	log.Printf("Parámetros recibidos: %+v\n", params)

	// Validar si no hay cambios en la posición o columna
	if params.FromColumnID == params.ToColumnID && params.NewPosition == params.OldPosition {
		log.Println("No se detectaron cambios en la posición o columna de la tarjeta.")
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"error": "No changes detected in position or column"})
		return
	}
	

	// Comenzar una transacción para garantizar consistencia
	tx, err := Db.Begin()
	if err != nil {
		log.Printf("Error al iniciar la transacción: %v\n", err)
		http.Error(w, "Database error", http.StatusInternalServerError)
		return
	}
	defer tx.Rollback()

	// Si la tarjeta se mueve a otra columna, actualizar la columna
	if params.FromColumnID != params.ToColumnID {
		log.Println("La tarjeta se mueve a otra columna.")

		// Actualizar columna y posición de la tarjeta
		_, err := tx.Exec("UPDATE cards SET column_id = ?, position = ? WHERE id = ?", params.ToColumnID, params.NewPosition, params.CardID)
		if err != nil {
			log.Printf("Error al actualizar columna y posición de la tarjeta: %v\n", err)
			http.Error(w, "Error updating card column", http.StatusInternalServerError)
			return
		}

		// Ajustar posiciones en la columna de origen
		_, err = tx.Exec("UPDATE cards SET position = position - 1 WHERE column_id = ? AND position > ?", params.FromColumnID, params.NewPosition)
		if err != nil {
			log.Printf("Error al ajustar posiciones en la columna de origen: %v\n", err)
			http.Error(w, "Error updating positions in source column", http.StatusInternalServerError)
			return
		}
	} else {
		// Reordenar dentro de la misma columna
		log.Println("La tarjeta se reordena dentro de la misma columna.")

		if params.NewPosition < params.OldPosition {
			// Mover hacia arriba: Incrementar entre NewPosition y OldPosition - 1
			_, err = tx.Exec(
				"UPDATE cards SET position = position + 1 WHERE column_id = ? AND position >= ? AND position < ?",
				params.ToColumnID, params.NewPosition, params.OldPosition,
			)
		} else {
			// Mover hacia abajo: Decrementar entre OldPosition + 1 y NewPosition
			_, err = tx.Exec(
				"UPDATE cards SET position = position - 1 WHERE column_id = ? AND position > ? AND position <= ?",
				params.ToColumnID, params.OldPosition, params.NewPosition,
			)
		}

		if err != nil {
			log.Printf("Error al actualizar posiciones dentro de la misma columna: %v\n", err)
			http.Error(w, "Error updating positions in same column", http.StatusInternalServerError)
			return
		}

		// Actualizar posición de la tarjeta
		_, err = tx.Exec("UPDATE cards SET position = ? WHERE id = ?", params.NewPosition, params.CardID)
		if err != nil {
			log.Printf("Error al actualizar posición de la tarjeta: %v\n", err)
			http.Error(w, "Error updating card position", http.StatusInternalServerError)
			return
		}
	}

	// Confirmar transacción
	if err := tx.Commit(); err != nil {
		log.Printf("Error al confirmar la transacción: %v\n", err)
		http.Error(w, "Transaction commit error", http.StatusInternalServerError)
		return
	}

	// Registro exitoso
	log.Printf("Tarjeta %d movida correctamente a columna %d en posición %d.\n", params.CardID, params.ToColumnID, params.NewPosition)
	w.WriteHeader(http.StatusOK)
}




