package engine

import (
	"context"
	"database/sql"
	"errors"
	"fmt"

	"github.com/google/uuid"
	"github.com/rohitbisht01/car-management-system/models"
)

type EngineStore struct {
	db *sql.DB
}

func New(db *sql.DB) *EngineStore {
	return &EngineStore{db: db}
}

func (e EngineStore) EngineById(ctx context.Context, id string) (models.Engine, error) {
	var engine models.Engine

	tx, err := e.db.BeginTx(ctx, nil)
	if err != nil {
		return engine, err
	}

	defer func() {
		if err != nil {
			if rbErr := tx.Rollback(); rbErr != nil {
				fmt.Printf("Transaction rollback error: %v\n", rbErr)
			}
		} else {
			if cmErr := tx.Commit(); cmErr != nil {
				fmt.Printf("Transaction commit error: %v\n", cmErr)
			}
		}
	}()

	err = tx.QueryRowContext(ctx, "SELECT id,displacement,no_of_cyclinders,car_range FROM engine WHERE id=$1", id).Scan(&engine.EngineID, &engine.Displacement, &engine.NoOfCyclinders, &engine.CarRange)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return engine, nil
		}
		return engine, err
	}
	return engine, nil

}

func (e EngineStore) CreateEngine(ctx context.Context, engineReq *models.EngineRequest) (models.Engine, error) {
	tx, err := e.db.BeginTx(ctx, nil)
	if err != nil {
		return models.Engine{}, err
	}

	defer func() {
		if err != nil {
			if rbErr := tx.Rollback(); rbErr != nil {
				fmt.Printf("Transaction rollback error: %v\n", rbErr)
			}
		} else {
			if cmErr := tx.Commit(); cmErr != nil {
				fmt.Printf("Transaction commit error: %v\n", cmErr)
			}
		}
	}()

	engineId := uuid.New()

	_, err = tx.ExecContext(ctx, "INSERT INTO engine (id,displacement,no_of_cyclinders,car_range) VALUES ($1,$2,$3,$4)", engineId, engineReq.Displacement, engineReq.NoOfCyclinders, engineReq.CarRange)

	if err != nil {
		return models.Engine{}, err
	}

	engine := models.Engine{
		EngineID:       engineId,
		Displacement:   engineReq.Displacement,
		NoOfCyclinders: engineReq.NoOfCyclinders,
		CarRange:       engineReq.CarRange,
	}

	return engine, nil
}

func (e EngineStore) EngineUpdate(ctx context.Context, id string, engineReq *models.EngineRequest) (models.Engine, error) {
	engineId, err := uuid.Parse(id)

	if err != nil {
		return models.Engine{}, fmt.Errorf("Invalid Engine ID: %w", err)
	}

	tx, err := e.db.BeginTx(ctx, nil)
	if err != nil {
		return models.Engine{}, err
	}

	defer func() {
		if err != nil {
			if rbErr := tx.Rollback(); rbErr != nil {
				fmt.Printf("Transaction rollback error: %v\n", rbErr)
			}
		} else {
			if cmErr := tx.Commit(); cmErr != nil {
				fmt.Printf("Transaction commit error: %v\n", cmErr)
			}
		}
	}()

	results, err := tx.ExecContext(ctx, "UPDATE engine SET displacement=$1, no_of_cyclinders=$2, car_range=$3 WHERE id=$4", engineReq.Displacement, engineReq.NoOfCyclinders, engineReq.CarRange, engineId)

	if err != nil {
		return models.Engine{}, err
	}

	rowsAffected, err := results.RowsAffected()
	if err != nil {
		return models.Engine{}, err
	}
	if rowsAffected == 0 {
		return models.Engine{}, errors.New("No rows were updated")
	}

	engine := models.Engine{
		EngineID:       engineId,
		Displacement:   engineReq.Displacement,
		NoOfCyclinders: engineReq.NoOfCyclinders,
		CarRange:       engineReq.CarRange,
	}

	return engine, nil
}

func (e EngineStore) EngineDelete(ctx context.Context, id string) (models.Engine, error) {
	var engine models.Engine

	tx, err := e.db.BeginTx(ctx, nil)
	if err != nil {
		return engine, err
	}

	defer func() {
		if err != nil {
			if rbErr := tx.Rollback(); rbErr != nil {
				fmt.Printf("Transaction rollback error: %v\n", rbErr)
			}
		} else {
			if cmErr := tx.Commit(); cmErr != nil {
				fmt.Printf("Transaction commit error: %v\n", cmErr)
			}
		}
	}()

	err = tx.QueryRowContext(ctx, "SELECT id,displacement,no_of_cyclinders,car_range FROM engine WHERE id=$1", id).Scan(&engine.EngineID, &engine.Displacement, &engine.NoOfCyclinders, &engine.CarRange)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return engine, nil
		}
		return engine, err
	}

	result, err := tx.ExecContext(ctx, "DELETE FROM engine WHERE id=$1", id)
	if err != nil {
		return models.Engine{}, err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return models.Engine{}, err
	}
	if rowsAffected == 0 {
		return models.Engine{}, errors.New("No row are deleted")
	}

	return engine, nil
}
