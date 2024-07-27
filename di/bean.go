package di

import "reflect"

type InitApplicationBean interface {
	NewInstance() reflect.Type
}
