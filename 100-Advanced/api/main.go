package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"io"
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

 	rows, err := Db.Query("SELECT id, title, description FROM cards WHERE column_id = ?", columnID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var cards []Card
	for rows.Next() {
		var card Card
		if err := rows.Scan(&card.ID, &card.Title, &card.Description); err != nil {
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
	// Leer el cuerpo de la solicitud para obtener los datos
	body, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Error leyendo el cuerpo de la solicitud: "+err.Error(), http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	// Log para ver los datos crudos recibidos
	fmt.Printf("Cuerpo recibido: %s\n", string(body))

	// Definir la estructura de los datos que se esperan
	var moveData struct {
		CardID      int `json:"card_id"`
		ToColumnID  int `json:"to_column_id"`
		NewPosition int `json:"new_position"`
	}

	// Decodificar el JSON recibido en la variable moveData
	if err := json.Unmarshal(body, &moveData); err != nil {
		http.Error(w, "Error decodificando JSON: "+err.Error(), http.StatusBadRequest)
		return
	}

	// Log para ver los datos decodificados
	fmt.Printf("Datos decodificados: CardID=%d, ToColumnID=%d, NewPosition=%d\n", moveData.CardID, moveData.ToColumnID, moveData.NewPosition)

	// Verificar que los datos sean válidos
	if moveData.CardID == 0 || moveData.ToColumnID == 0 || moveData.NewPosition < 0 {
		http.Error(w, "Datos inválidos para mover la tarjeta.", http.StatusBadRequest)
		return
	}

	// Comenzar una transacción
	tx, err := Db.Begin()
	if err != nil {
		http.Error(w, "Error al comenzar transacción: "+err.Error(), http.StatusInternalServerError)
		return
	}
	defer tx.Rollback()

	// Obtener la columna actual de la tarjeta
	var currentColumnID int
	err = tx.QueryRow("SELECT column_id FROM cards WHERE id = ?", moveData.CardID).Scan(&currentColumnID)
	if err != nil {
		if err == sql.ErrNoRows {
			http.Error(w, "Tarjeta no encontrada", http.StatusNotFound)
			return
		}
		http.Error(w, "Error al consultar la base de datos (columna actual): "+err.Error(), http.StatusInternalServerError)
		fmt.Println("Error al consultar columna actual:", err)
		return
	}

	// Log columna actual
	fmt.Printf("Columna actual de la tarjeta: %d\n", currentColumnID)

	if currentColumnID == moveData.ToColumnID {
		// Reordenar en la misma columna
		err = reorderCardsInColumn(tx, currentColumnID, moveData.CardID, moveData.NewPosition)
		if err != nil {
			http.Error(w, "Error al reordenar las tarjetas: "+err.Error(), http.StatusInternalServerError)
			fmt.Println("Error al reordenar en la misma columna:", err)
			return
		}
	} else {
		// Ajustar posiciones en columna de origen
		_, err = tx.Exec(
			"UPDATE cards SET position = position - 1 WHERE column_id = ? AND position > (SELECT position FROM cards WHERE id = ?)",
			currentColumnID, moveData.CardID,
		)
		if err != nil {
			http.Error(w, "Error al ajustar posiciones en columna origen: "+err.Error(), http.StatusInternalServerError)
			fmt.Println("Error ajustando columna origen:", err)
			return
		}

		// Mover tarjeta a nueva columna
		_, err = tx.Exec(
			"UPDATE cards SET column_id = ?, position = ? WHERE id = ?",
			moveData.ToColumnID, moveData.NewPosition, moveData.CardID,
		)
		if err != nil {
			http.Error(w, "Error al mover la tarjeta: "+err.Error(), http.StatusInternalServerError)
			fmt.Println("Error moviendo la tarjeta:", err)
			return
		}

		// Reordenar en la nueva columna
		err = reorderCardsInColumn(tx, moveData.ToColumnID, moveData.CardID, moveData.NewPosition)
		if err != nil {
			http.Error(w, "Error al reordenar las tarjetas en nueva columna: "+err.Error(), http.StatusInternalServerError)
			fmt.Println("Error reordenando nueva columna:", err)
			return
		}
	}

	// Confirmar la transacción
	if err := tx.Commit(); err != nil {
		http.Error(w, "Error al confirmar la transacción: "+err.Error(), http.StatusInternalServerError)
		fmt.Println("Error confirmando transacción:", err)
		return
	}

	// Log de éxito
	fmt.Println("Tarjeta movida exitosamente.")

	// Depuración: Verificar las columnas actual y destino
	printCardsInColumn(currentColumnID)
	printCardsInColumn(moveData.ToColumnID)

	// Responder con éxito
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Tarjeta movida exitosamente")
}


func printCardsInColumn(columnID int) {
	rows, err := Db.Query("SELECT id, title, position FROM cards WHERE column_id = ? ORDER BY position", columnID)
	if err != nil {
		log.Println("Error al consultar las tarjetas en columna:", err)
		return
	}
	defer rows.Close()

	fmt.Printf("Tarjetas en columna %d:\n", columnID)
	for rows.Next() {
		var id, position int
		var title string
		if err := rows.Scan(&id, &title, &position); err != nil {
			log.Println("Error al escanear tarjeta:", err)
			return
		}
		fmt.Printf("ID: %d, Título: %s, Posición: %d\n", id, title, position)
	}
}


// Función para reordenar las tarjetas dentro de una columna
func reorderCardsInColumn(tx *sql.Tx, columnID, movedCardID, newPosition int) error {
	// Obtener todas las tarjetas ordenadas excepto la que se mueve
	rows, err := tx.Query("SELECT id FROM cards WHERE column_id = ? AND id != ? ORDER BY position ASC", columnID, movedCardID)
	if err != nil {
		return err
	}
	defer rows.Close()

	var cardIDs []int
	for rows.Next() {
		var id int
		if err := rows.Scan(&id); err != nil {
			return err
		}
		cardIDs = append(cardIDs, id)
	}

	// Insertar la tarjeta movida en la posición indicada
	if newPosition > len(cardIDs) {
		newPosition = len(cardIDs) // Ajustar al final si es necesario
	}
	cardIDs = append(cardIDs[:newPosition], append([]int{movedCardID}, cardIDs[newPosition:]...)...)

	// Actualizar las posiciones en la base de datos
	for idx, id := range cardIDs {
		_, err := tx.Exec("UPDATE cards SET position = ? WHERE id = ?", idx, id)
		if err != nil {
			return err
		}
	}

	return nil
}