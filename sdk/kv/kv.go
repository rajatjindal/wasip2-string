package kv

import (
	"fmt"

	keyvalue "github.com/rajatjindal/wasip2-string/internal/fermyon/spin/v2.0.0/key-value"
	"github.com/ydnar/wasm-tools-go/cm"
)

type Store struct {
	store *keyvalue.Store
}

func Open(label string) (*Store, error) {
	result := keyvalue.StoreOpen(label)
	if result.IsErr() {
		return nil, errorVariantToError(*result.Err())
	}

	return &Store{
		store: result.OK(),
	}, nil
}

func OpenDefault() (*Store, error) {
	return Open("default")
}

func (s *Store) Set(key string, value []byte) error {
	result := s.store.Set(key, cm.ToList(value))
	if result.IsErr() {
		return errorVariantToError(*result.Err())
	}

	return nil
}

func (s *Store) Get(key string) ([]byte, error) {
	result := s.store.Get(key)
	if result.IsErr() {
		return nil, errorVariantToError(*result.Err())
	}

	value := result.OK()
	if value.None() {
		return []byte(""), nil
	}

	return value.Some().Slice(), nil
}

func (s *Store) Delete(key string) error {
	result := s.store.Delete(key)
	if result.IsErr() {
		return errorVariantToError(*result.Err())
	}

	return nil
}

func (s *Store) Exists(key string) (bool, error) {
	result := s.store.Exists(key)
	if result.IsErr() {
		return false, errorVariantToError(*result.Err())
	}

	return *result.OK(), nil
}

func (s *Store) GetKeys() ([]string, error) {
	result := s.store.GetKeys()
	if result.IsErr() {
		return nil, errorVariantToError(*result.Err())
	}

	return result.OK().Slice(), nil
}

func errorVariantToError(code keyvalue.Error) error {
	switch code {
	case keyvalue.ErrorAccessDenied():
		return fmt.Errorf("access denied")
	case keyvalue.ErrorNoSuchStore():
		return fmt.Errorf("no such store")
	case keyvalue.ErrorStoreTableFull():
		return fmt.Errorf("store table full")
	default:
		if code.Other() != nil {
			return fmt.Errorf(*code.Other())
		}

		return fmt.Errorf("no error provided by host implementation")
	}
}
