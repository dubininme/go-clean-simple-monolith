//go:build integration

package tests

import (
	"database/sql"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/dubininme/go-clean-simple-monolith/internal/bootstrap"
	"github.com/dubininme/go-clean-simple-monolith/internal/config"
	_ "github.com/go-sql-driver/mysql"
	"github.com/stretchr/testify/require"
)

var testServer *httptest.Server

func TestMain(m *testing.M) {
	cfg, err := config.InitConfig()
	if err != nil {
		panic(err)
	}
	db, err := sql.Open("mysql", cfg.DB.DSN)
	if err != nil {
		panic(err)
	}
	defer db.Close()
	if err := loadFixtures(db); err != nil {
		panic(err)
	}

	echoHandler := bootstrap.NewTestServer(db, cfg)
	testServer = httptest.NewServer(echoHandler)
	code := m.Run()
	testServer.Close()
	os.Exit(code)
}

func loadFixtures(db *sql.DB) error {
	fixture, err := os.ReadFile("tests/fixtures/orders.sql")
	if err != nil {
		return err
	}
	_, err = db.Exec(string(fixture))
	return err
}

func TestGetOrder_Snapshot(t *testing.T) {
	resp, err := http.Get(fmt.Sprintf("%s/api/orders/1", testServer.URL))
	require.NoError(t, err)
	defer resp.Body.Close()
	body, _ := io.ReadAll(resp.Body)
	snaptest.AssertJSON(t, string(body))
}

func TestMarkOrderPaid_Snapshot(t *testing.T) {
	req, _ := http.NewRequest(http.MethodPost, fmt.Sprintf("%s/api/orders/1/pay", testServer.URL), nil)
	resp, err := http.DefaultClient.Do(req)
	require.NoError(t, err)
	defer resp.Body.Close()
	body, _ := io.ReadAll(resp.Body)
	snaptest.AssertJSON(t, string(body))
}
