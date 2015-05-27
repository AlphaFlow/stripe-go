package sku

import (
	"fmt"
	"math/rand"
	"testing"
	"time"

	stripe "github.com/stripe-internal/stripe-go"
	"github.com/stripe-internal/stripe-go/product"
	. "github.com/stripe-internal/stripe-go/utils"
)

func init() {
	stripe.Key = GetTestKey()
	rand.Seed(time.Now().UTC().UnixNano())
}

func TestSKUCreate(t *testing.T) {
	active := true

	p, err := product.New(&stripe.ProductParams{
		Active:    &active,
		Name:      "test name",
		Desc:      "This is a description",
		Caption:   "This is a caption",
		Attrs:     []string{"attr1", "attr2"},
		URL:       "http://example.com",
		Shippable: &active,
	})
	if err != nil {
		t.Fatalf("%+v", err)
	}

	randID := fmt.Sprintf("TEST-SKU-%v", RandSeq(16))
	sku, err := New(&stripe.SKUParams{
		ID:        randID,
		Active:    &active,
		Desc:      "This is a SKU description",
		Attrs:     map[string]string{"attr1": "val1", "attr2": "val2"},
		Price:     499,
		Currency:  "usd",
		Inventory: stripe.Inventory{Type: "bucket", Value: "scant"},
		Product:   p.ID,
	})

	if err != nil {
		t.Fatalf("%+v", err)
	}

	if sku.ID == "" {
		t.Errorf("ID is not set %v", sku.ID)
	}

	if sku.Created == 0 {
		t.Errorf("Created date is not set")
	}

	if sku.Updated == 0 {
		t.Errorf("Updated is not set")
	}

	if sku.Desc != "This is a SKU description" {
		t.Errorf("Description is invalid: %v", sku.Desc)
	}

	if len(sku.Attrs) != 2 {
		t.Errorf("Invalid attributes: %v", sku.Attrs)
	}

	if sku.Attrs["attr1"] != "val1" {
		t.Errorf("Invalid attributes: %v", sku.Attrs)
	}

	if sku.Attrs["attr2"] != "val2" {
		t.Errorf("Invalid attributes: %v", sku.Attrs)
	}

	if sku.Inventory.Type != "bucket" {
		t.Errorf("Invalid inventory type: %v", sku.Inventory.Type)
	}

	if sku.Inventory.Value != "scant" {
		t.Errorf("Invalid inventory type: %v", sku.Inventory.Value)
	}
}