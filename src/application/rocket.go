package application

import (
	"errors"
	"fmt"
	"math"
	"math/rand/v2"
	"time"
)

type RocketInterface interface {
	isValid() (bool, error)
	fly(seconds int32) error
}

const (
	CRASH  = "CRASH"
	FLYING = "FLYING"

	HOUSE_WINS_PERCENT = 51
)

type Rocket struct {
	Amount int32
	Status string
}

func (r *Rocket) IsValid() (bool, error) {
	if r.Status == "" {
		r.Status = FLYING
	}

	if r.Amount < 0 {
		return false, errors.New("the amount should be greater than zero")
	}

	if r.Status != CRASH && r.Status != FLYING {
		return false, errors.New("the status should be CRASH or FLYING")
	}

	return true, nil
}

func (r *Rocket) Fly(seconds int32) error {
	if seconds <= 0 {
		return errors.New("seconds must be greater than zero")
	}

	for i := range seconds {
		randNum := rand.IntN(100) + 1

		if randNum < HOUSE_WINS_PERCENT {
			r.Status = CRASH
			break
		}

		expVariation := int32(math.Pow(2, float64(i)))
		r.Amount = r.Amount + expVariation

		fmt.Printf("[%d] rocket fly more %d to amount: %d and rand %d\n", i, expVariation, r.Amount, randNum)

		time.Sleep(1 * time.Second)
	}

	fmt.Printf("Rocket amount %d and status %s\n", r.Amount, r.Status)

	return nil
}
