package poker_test

import (
	"github.com/quzhi1/GolangPlayground/poker"
	"net/http"
	"net/http/httptest"
	"testing"
)

//server_integration_test.go
func TestRecordingWinsAndRetrievingThem(t *testing.T) {
	database, cleanDatabase := poker.CreateTempFile(t, `[]`)
	defer cleanDatabase()
	store, err := poker.NewFileSystemPlayerStore(database)
	poker.AssertNoError(t, err)
	server := poker.NewPlayerServer(store)
	player := "Pepper"

	server.ServeHTTP(httptest.NewRecorder(), poker.NewPostWinRequest(player))
	server.ServeHTTP(httptest.NewRecorder(), poker.NewPostWinRequest(player))
	server.ServeHTTP(httptest.NewRecorder(), poker.NewPostWinRequest(player))

	t.Run("get score", func(t *testing.T) {
		response := httptest.NewRecorder()
		server.ServeHTTP(response, poker.NewGetScoreRequest(player))
		poker.AssertStatus(t, response.Code, http.StatusOK)

		poker.AssertResponseBody(t, response.Body.String(), "3")
	})

	t.Run("get league", func(t *testing.T) {
		response := httptest.NewRecorder()
		server.ServeHTTP(response, poker.NewLeagueRequest())
		poker.AssertStatus(t, response.Code, http.StatusOK)

		got := poker.GetLeagueFromResponse(t, response.Body)
		want := []poker.Player{
			{"Pepper", 3},
		}
		poker.AssertLeague(t, got, want)
	})
}
