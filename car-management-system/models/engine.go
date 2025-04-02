package models

import (
	"errors"

	"github.com/google/uuid"
)

type Engine struct {
	EngineID       uuid.UUID `json:"engine_id"`
	Displacement   int64     `json:"displacement"`
	NoOfCyclinders int64     `json:"no_of_cyclinders"`
	CarRange       int64     `json:"car_range"`
}

type EngineRequest struct {
	Displacement   int64 `json:"displacement"`
	NoOfCyclinders int64 `json:"no_of_cyclinders"`
	CarRange       int64 `json:"car_range"`
}

func ValidateEngineRequest(EngineReq EngineRequest) error {
	if err := validateDisplacement(EngineReq.Displacement); err != nil {
		return err
	}
	if err := validateNoOfCyclinders(EngineReq.NoOfCyclinders); err != nil {
		return err
	}
	if err := validateCarRange(EngineReq.CarRange); err != nil {
		return err
	}
	return nil
}

func validateDisplacement(displacement int64) error {
	if displacement <= 0 {
		return errors.New("displacement must be greater than zero")
	}
	return nil
}

func validateNoOfCyclinders(noOfCyclinders int64) error {
	if noOfCyclinders <= 0 {
		return errors.New("noOfCyclinder must be greater than zero")
	}
	return nil
}

func validateCarRange(carRange int64) error {
	if carRange <= 0 {
		return errors.New("carRange must be greater that zero")
	}
	return nil
}
