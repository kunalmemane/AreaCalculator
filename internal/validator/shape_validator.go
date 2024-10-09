package validator

import (
	"errors"

	"github.com/kunalmemane9150/AreaCalculator/internal/model"
)

func Validator(shape model.IShape) (model.IShape, error) {

	switch sh := shape.(type) {
	case *model.Square:

		if sh.Side > 0 {
			return shape, nil
		} else {
			return nil, errors.New("invalid square side")
		}

	case *model.Circle:

		if sh.Radius > 0 {
			return shape, nil
		} else {
			return nil, errors.New("invalid circle radius")
		}

	case *model.Rectangle:

		if sh.Length > 0 && sh.Breadth > 0 {
			return shape, nil
		} else {
			return nil, errors.New("invalid rectangle params")
		}

	case *model.Triangle:

		if sh.SideA > 0 && sh.SideB > 0 && sh.SideC > 0 {
			return shape, nil
		} else {
			return nil, errors.New("invalid triangle side")
		}

	default:
		return nil, errors.New("invalid shape")
	}

}
