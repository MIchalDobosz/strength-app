package database

import "github.com/jmoiron/sqlx"

const MigrationIndex = 0 // temporary solution

func getMigrations() []func(*sqlx.DB) error {
	return []func(*sqlx.DB) error{
		createExercisesTable,
		createMacrocyclesTable,
		createMesosyclesTable,
		createMicrocyclesTable,
		createSessionsTable,
		createSlotsTable,
		createSetsTable,
	}
}

func createExercisesTable(db *sqlx.DB) error {
	sql := `
		CREATE TABLE exercises (
			id INT AUTO_INCREMENT UNIQUE NOT NULL,
			name VARCHAR(128) NOT NULL,
			PRIMARY KEY (id)
		)
	`
	_, err := db.Exec(sql)
	if err != nil {
		return err
	}

	return nil
}

func createMacrocyclesTable(db *sqlx.DB) error {
	sql := `
		CREATE TABLE macrocycles (
			id INT AUTO_INCREMENT UNIQUE NOT NULL,
			name VARCHAR(128) NOT NULL,
			PRIMARY KEY (id)
		)
	`
	_, err := db.Exec(sql)
	if err != nil {
		return err
	}

	return nil
}

func createMesosyclesTable(db *sqlx.DB) error {
	sql := `
		CREATE TABLE mesocycles (
			id INT AUTO_INCREMENT UNIQUE NOT NULL,
			macrocycle_id INT NOT NULL,
			name VARCHAR(128) NOT NULL,
			PRIMARY KEY (id),
			FOREIGN KEY (macrocycle_id) REFERENCES macrocycles(id) ON DELETE CASCADE
		)
	`
	_, err := db.Exec(sql)
	if err != nil {
		return err
	}

	return nil
}

func createMicrocyclesTable(db *sqlx.DB) error {
	sql := `
		CREATE TABLE microcycles (
			id INT AUTO_INCREMENT UNIQUE NOT NULL,
			mesocycle_id INT NOT NULL,
			name VARCHAR(128) NOT NULL,
			PRIMARY KEY (id),
			FOREIGN KEY (mesocycle_id) REFERENCES mesocycles(id) ON DELETE CASCADE
		)
	`
	_, err := db.Exec(sql)
	if err != nil {
		return err
	}

	return nil
}

func createSessionsTable(db *sqlx.DB) error {
	sql := `
		CREATE TABLE sessions (
			id INT AUTO_INCREMENT UNIQUE NOT NULL,
			microcycle_id INT NOT NULL,
			name VARCHAR(128) NOT NULL,
			planned_date DATE,
			performed_date DATE,
			completed BOOL NOT NULL,
			PRIMARY KEY (id),
			FOREIGN KEY (microcycle_id) REFERENCES microcycles(id) ON DELETE CASCADE
		)
	`
	_, err := db.Exec(sql)
	if err != nil {
		return err
	}

	return nil
}

func createSlotsTable(db *sqlx.DB) error {
	sql := `
		CREATE TABLE slots (
			id INT AUTO_INCREMENT UNIQUE NOT NULL,
			session_id INT NOT NULL,
			planned_exercise_id INT NOT NULL,
			performed_exercise_id INT,
			completed BOOL NOT NULL,
			PRIMARY KEY (id),
			FOREIGN KEY (session_id) REFERENCES sessions(id) ON DELETE CASCADE,
			FOREIGN KEY (planned_exercise_id) REFERENCES exercises(id) ON DELETE CASCADE,
			FOREIGN KEY (performed_exercise_id) REFERENCES exercises(id) ON DELETE CASCADE
		)
	`
	_, err := db.Exec(sql)
	if err != nil {
		return err
	}

	return nil
}

func createSetsTable(db *sqlx.DB) error {
	sql := `
		CREATE TABLE sets (
			id INT AUTO_INCREMENT UNIQUE NOT NULL,
			slot_id INT NOT NULL,
			set_no INT NOT NULL,
			planned_reps INT NOT NULL,
			planned_reps_unit TINYINT NOT NULL,
			performed_reps INT NOT NULL,
			performed_reps_unit TINYINT NOT NULL,
			planned_percent INT NOT NULL,
			planned_rpe INT NOT NULL,
			performed_rpe INT NOT NULL,
			planned_weight DECIMAL(5, 1) NOT NULL,
			performed_weight DECIMAL(5, 1) NOT NULL,
			completed BOOL NOT NULL,
			PRIMARY KEY (id),
			FOREIGN KEY (slot_id) REFERENCES slots(id) ON DELETE CASCADE
		)
	`
	_, err := db.Exec(sql)
	if err != nil {
		return err
	}

	return nil
}
