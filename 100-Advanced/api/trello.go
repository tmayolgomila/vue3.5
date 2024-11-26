package main

import (
	"database/sql"
	"log"
	"strings"

	_ "modernc.org/sqlite"
)

var Db *sql.DB // Db ahora es exportada.

func InitDB() { // InitDB ahora es exportada.
	var err error
	Db, err = sql.Open("sqlite", "trello.db")
	if err != nil {
		log.Fatal(err)
	}

	// Crear tablas si no existen
	createTable := `
	CREATE TABLE IF NOT EXISTS projects (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		name TEXT,
		description TEXT
	);
	CREATE TABLE IF NOT EXISTS columns (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		project_id INTEGER,
		name TEXT,
		FOREIGN KEY(project_id) REFERENCES projects(id)
	);
	CREATE TABLE IF NOT EXISTS cards (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		column_id INTEGER,
		title TEXT,
		description TEXT,
		position INTEGER,
		FOREIGN KEY(column_id) REFERENCES columns(id)
	);
	`
	if _, err := Db.Exec(createTable); err != nil {
		log.Fatal(err)
	}

	// Agregar columna "position" si no existe
	alterTable := `
	ALTER TABLE cards ADD COLUMN position INTEGER DEFAULT 0;
	`
	if _, err := Db.Exec(alterTable); err != nil {
		if !strings.Contains(err.Error(), "duplicate column name") { // Ignora el error si ya existe
			log.Fatal(err)
		}
	}

	log.Println("Base de datos inicializada correctamente.")
}


