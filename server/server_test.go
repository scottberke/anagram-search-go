package server

import (
  "net/http"
  "net/http/httptest"
  "encoding/json"
  "io/ioutil"
  "testing"
  "bytes"
  "sort"

  "github.com/gorilla/mux"
)

// Function to fail test if any of the calls throw an error
func checkError(err error, t *testing.T) {
	if err != nil {
		t.Errorf("An error occurred. %v", err)
	}
}

func createAnagrams(t *testing.T, server *Server) *httptest.ResponseRecorder {
  var jsonStr = []byte(`{"words": ["read", "dear", "dare"] }`)
  request, err := http.NewRequest("POST", "/words.json", bytes.NewBuffer(jsonStr))
  checkError(err, t)
  request.Header.Add("Content-Type", "Content-Type: application/json")

  rr := httptest.NewRecorder()
  handler := http.HandlerFunc(server.createAnagramsHandler)
  handler.ServeHTTP(rr, request)

  return rr
}

func setupMuxRoutes(server *Server) *mux.Router {
  r := mux.NewRouter()

  r.HandleFunc("/anagrams/{word}.json", server.getAnagramsHandler).
    Methods("GET")

  r.HandleFunc("/anagrams/{word}.json", server.getAnagramsHandler).
    Methods("GET")

  r.HandleFunc("/words/{word}.json", server.deleteSingleAnagramHandler).
    Methods("DELETE")

  r.HandleFunc("/words.json", server.deleteAnagramsHandler).
    Methods("DELETE")

  return r
}

// Struct to unmarshal json response into
type Body struct {
    Anagrams   []string  `json:"anagrams"`
  }

func parseJsonResponse(rr *httptest.ResponseRecorder) *Body {
      response := new(Body)
      body, _ := ioutil.ReadAll(rr.Body)
      json.Unmarshal(body, &response)

      return response
}

func testEq(a, b []string) bool {
  sort.Strings(a)
  sort.Strings(b)

  // If one is nil, the other must also be nil. Same goes for size equality
  if (len(a) != len(b)) || (a == nil) != (b == nil) {
    return false
  }

  for i := range a {
    if a[i] != b[i] {
      return false
    }
  }

  return true
}


func TestCreateAnagramValidRequest(t *testing.T) {
    // Create a new server and anagrams
    done := make(chan bool)
    server := NewServer(8080, done)
    rr := createAnagrams(t, server)

    // Make sure our status code is correctly returned
    if status := rr.Code; status != http.StatusCreated {
      t.Errorf("Status code incorrect. Expected %d Got %d", http.StatusCreated, status)
    }
}

func TestGetAnagramValidRequest(t *testing.T) {
    // Create a new server and anagrams
    done := make(chan bool)
    server := NewServer(8080, done)
    createAnagrams(t, server)

    request, err := http.NewRequest("GET", "/anagrams/read.json", nil)
    checkError(err, t)

    r := setupMuxRoutes(server)
    rr := httptest.NewRecorder()

    r.ServeHTTP(rr, request)

    response := parseJsonResponse(rr)

    expectedWords := []string{ "dear", "dare"}
    if testEq(expectedWords, response.Anagrams) != true {
      t.Errorf("Response incorrect. Expected %v Got %v", expectedWords, response.Anagrams)
    }

    // Make sure our status code is correctly returned
    if status := rr.Code; status != http.StatusOK {
      t.Errorf("Status code incorrect. Expected %d Got %d", http.StatusOK, status)
    }
}

func TestGetAnagramWithLimitRequest(t *testing.T) {
    // Create a new server and anagrams
    done := make(chan bool)
    server := NewServer(8080, done)
    createAnagrams(t, server)

    request, err := http.NewRequest("GET", "/anagrams/read.json?limit=1", nil)
    checkError(err, t)

    r := setupMuxRoutes(server)

    rr := httptest.NewRecorder()

    r.ServeHTTP(rr, request)

    response := parseJsonResponse(rr)

    if len(response.Anagrams) != 1 {
      t.Errorf("Response limit incorrect. Expected %d Got %d", 1, len(response.Anagrams))
    }

    // Make sure our status code is correctly returned
    if status := rr.Code; status != http.StatusOK {
      t.Errorf("Status code incorrect. Expected %d Got %d", http.StatusOK, status)
    }
}


func TestDeleteSingleAnagramRequest(t *testing.T) {
    // Create a new server and anagrams
    done := make(chan bool)
    server := NewServer(8080, done)
    createAnagrams(t, server)

    request, err := http.NewRequest("DELETE", "/words/read.json", nil)
    checkError(err, t)

    r := setupMuxRoutes(server)

    rrDelete := httptest.NewRecorder()
    r.ServeHTTP(rrDelete, request)

    request, err = http.NewRequest("GET", "/anagrams/dear.json", nil)
    checkError(err, t)

    rrGet := httptest.NewRecorder()
    r.ServeHTTP(rrGet, request)

    response := parseJsonResponse(rrGet)

    if response.Anagrams[0] != "dare" {
      t.Errorf("Response word incorrect. Expected %s Got %s", "dare", response.Anagrams[0])
    }

    if len(response.Anagrams) != 1 {
      t.Errorf("Response size incorrect. Expected %d Got %d", 1, len(response.Anagrams))
    }

    // Make sure our status code is correctly returned
    if status := rrDelete.Code; status != http.StatusNoContent {
      t.Errorf("Status code incorrect. Expected %d Got %d", http.StatusNoContent, status)
    }
}



func TestDeleteAllAnagramsRequest(t *testing.T) {
    // Create a new server and anagrams
    done := make(chan bool)
    server := NewServer(8080, done)
    createAnagrams(t, server)

    request, err := http.NewRequest("DELETE", "/words.json", nil)
    checkError(err, t)

    r := setupMuxRoutes(server)

    rrDelete := httptest.NewRecorder()
    r.ServeHTTP(rrDelete, request)

    request, err = http.NewRequest("GET", "/anagrams/read.json", nil)
    checkError(err, t)

    rrGet := httptest.NewRecorder()
    r.ServeHTTP(rrGet, request)

    response := parseJsonResponse(rrGet)

    if len(response.Anagrams) != 0 {
      t.Errorf("Response size incorrect. Expected %d Got %d", 0, len(response.Anagrams))
    }

    // Make sure our status code is correctly returned
    if status := rrDelete.Code; status != http.StatusNoContent {
      t.Errorf("Status code incorrect. Expected %d Got %d", http.StatusNoContent, status)
    }
}
