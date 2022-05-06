package transaction

import (
	"database/sql"
	"fita/Config"
	"fita/Controller/DTO/request"
	"fita/Controller/DTO/response"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/graphql-go/graphql"
	"github.com/stretchr/testify/assert"
	"log"
	"testing"
)

func ConnectionMock() (*sql.DB, sqlmock.Sqlmock) {
	db, mock, err := sqlmock.New()
	if err != nil {
		log.Println(err)
	}

	return db, mock
}

func TestRepositoryNew(t *testing.T) {
	var dataMockQuery = request.ReqBody{
		Mutation: "",
		Query:    "{items (name:\"Google Home\") {name, price, qty}}",
	}

	dataMockQuery = request.ReqBody{
		Mutation: "{{addToCart (sku: \"43N23P\", name:\"MacBook Pro 1\", qty: 1) {sku, name, qty}} }",
		Query: "",
	}

	var resMock = schemaMutation{
		Mutation:  dataMockQuery.Mutation,
		Object: nil,
	}

	res := RepositoryNew(dataMockQuery)
	assert.NotNil(t, &resMock)
	assert.EqualValues(t, &resMock, res)
}

func TestSchemaQuery_Result(t *testing.T) {
	db, mock := ConnectionMock()
	Config.SqlConnection = db
	defer db.Close()

	body := request.Cells{
		&request.Cell{
			Sku: "A304SD",
			Qty: 3,
		},
	}

	bodyReq := map[string]interface{}{
		"params": body,
	}

	var paramsMock = graphql.ResolveParams {
		Args: bodyReq,
	}

	var resMock = []*response.CheckingOutDTO{
		&response.CheckingOutDTO{
			Sku:        "A304SD",
			Name:       []string{"Alexa Speaker"},
			TotalPrice: "0.00",
			TotalQty: 3,
		},
	}

	query := "SELECT tc.sku, tc.name, ti.price " +
		"FROM t_cart tc " +
		"INNER JOIN t_items ti ON tc.sku = ti.sku " +
		"WHERE tc.sku = \\$1"

	rows := sqlmock.NewRows([]string{"sku", "name", "total_price"}).
		AddRow(resMock[0].Sku, resMock[0].Name[0], resMock[0].TotalPrice)

	mock.ExpectQuery(query).WithArgs("A304SD").WillReturnRows(rows)

	res, err := checkingOut(paramsMock)
	assert.Nil(t, err)
	assert.Equal(t, resMock, res)
}

func TestSchemaMutation_Result(t *testing.T) {
	db, mock := ConnectionMock()
	Config.SqlConnection = db
	defer db.Close()

	bodyReq := map[string]interface{}{
		"sku":   "43N23P",
		"name":  "MacBook Pro",
		"qty":   1,
	}

	var paramsMock = graphql.ResolveParams {
		Args: bodyReq,
	}

	resMock := response.ItemsCartDTO{
		Sku:  "43N23P",
		Name: "MacBook Pro",
		Qty:  1,
	}

	query := "INSERT INTO t_cart\\(sku, name, qty\\) VALUES\\(\\$1, \\$2, \\$3\\)"
	mock.ExpectExec(query).WithArgs(paramsMock.Args["sku"].(string), paramsMock.Args["name"].(string), paramsMock.Args["qty"].(int)).WillReturnResult(
		sqlmock.NewResult(0, 1))

	res, err := addToChart(paramsMock)
	assert.Nil(t, err)
	assert.Equal(t, &resMock, res)
}
