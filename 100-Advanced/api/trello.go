package main

import (
	"database/sql"
	"log"

	_ "modernc.org/sqlite" // Driver para SQLite
)

// Db es la conexión global a la base de datos
var Db *sql.DB

// InitDB inicializa la conexión y estructura de la base de datos
func InitDB() {
	var err error
	Db, err = sql.Open("sqlite", "trello.db") // Conexión a la base de datos "trello.db"
	if err != nil {
		log.Fatal("Error al abrir la base de datos:", err)
	}

	// Crear tablas si no existen
	createTablesSQL := `
	CREATE TABLE IF NOT EXISTS projects (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		name TEXT NOT NULL,
		description TEXT
	);

	CREATE TABLE IF NOT EXISTS columns (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		project_id INTEGER NOT NULL,
		name TEXT NOT NULL,
		FOREIGN KEY(project_id) REFERENCES projects(id) ON DELETE CASCADE
	);

	CREATE TABLE IF NOT EXISTS cards (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		column_id INTEGER NOT NULL,
		title TEXT NOT NULL,
		description TEXT,
		position INTEGER DEFAULT 0,
		FOREIGN KEY(column_id) REFERENCES columns(id) ON DELETE CASCADE
	);
	`

	if _, err := Db.Exec(createTablesSQL); err != nil {
		log.Fatal("Error al crear las tablas:", err)
	}

	// Verificar si la columna "position" ya existe antes de intentar agregarla
	if !columnExists("cards", "position") {
		alterTableSQL := `
		ALTER TABLE cards ADD COLUMN position INTEGER DEFAULT 0;
		`
		if _, err := Db.Exec(alterTableSQL); err != nil {
			log.Fatal("Error al agregar la columna 'position':", err)
		} else {
			log.Println("Columna 'position' agregada exitosamente.")
		}
	}

	log.Println("Base de datos inicializada correctamente.")
}

// columnExists verifica si una columna existe en una tabla
func columnExists(tableName string, columnName string) bool {
	query := `
	SELECT 1 
	FROM pragma_table_info(?) 
	WHERE name = ?;
	`
	var exists int
	err := Db.QueryRow(query, tableName, columnName).Scan(&exists)
	return err == nil && exists == 1
}
