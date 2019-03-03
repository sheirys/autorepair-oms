package order_test

import (
	"testing"

	"gopkg.in/mgo.v2/bson"

	"github.com/sheirys/autorepair-oms/order"
	"github.com/sheirys/autorepair-oms/types"
	"github.com/stretchr/testify/assert"
)

func TestMockOrderServiceInit(t *testing.T) {
	service := order.MockOrderService{}
	assert.NoError(t, service.Init())
}

// TestMockOrderServiceGet is used to check if MockOrderService can extract
// single order without any erros.
func TestMockOrderServiceGet(t *testing.T) {
	service := order.MockOrderService{}
	assert.NoError(t, service.Init())

	// create
	order, err := service.Create(types.Order{})
	assert.NoError(t, err)

	// extract created
	get, err := service.Get(order.ID)
	assert.NoError(t, err)
	assert.Equal(t, order, get)

	// try to extract non existing
	get, err = service.Get(bson.NewObjectId())
	assert.Error(t, err)
	assert.Equal(t, err, types.ErrDocumentNotFound)
}

// TestMockOrderServiceAll is used to text if MockOrderService can extrc all existing
// orders from database without any erros.
func TestMockOrderServiceAll(t *testing.T) {
	service := order.MockOrderService{}
	assert.NoError(t, service.Init())

	// create orders. We cannot use testTables here, because
	// create does not accept any arguments.
	for i := 0; i < 5; i++ {
		_, err := service.Create(types.Order{})
		assert.NoError(t, err)
	}

	orders, err := service.All()
	assert.NoError(t, err)
	assert.Equal(t, 5, len(orders))
}

func TestMockOrderServiceCreate(t *testing.T) {
	service := order.MockOrderService{}
	assert.NoError(t, service.Init())

	// create
	order, err := service.Create(types.Order{})
	assert.NoError(t, err)
	assert.NotNil(t, order)

	// extract created
	get, err := service.Get(order.ID)
	assert.NoError(t, err)
	assert.Equal(t, order, get)
}

// TestMockOrderServiceUpdate is used to check MockOrderService can update single
// order with provided data.
func TestMockOrderServiceUpdate(t *testing.T) {
	service := order.MockOrderService{}
	assert.NoError(t, service.Init())

	new := types.Order{}

	// create
	old, err := service.Create(types.Order{})
	assert.NoError(t, err)
	assert.NotNil(t, old)

	// as we update by ID, we need to set order id, after we create one.
	new.ID = old.ID

	// update and check if updated equals to order we wanted.
	updated, err := service.Update(new)
	assert.NoError(t, err)
	assert.Equal(t, new, updated)

	// try to update non-exiting order
	updated, err = service.Update(types.Order{})
	assert.Error(t, err)
	assert.Equal(t, types.ErrDocumentNotFound, err)
}
