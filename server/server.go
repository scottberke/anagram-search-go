package server

import (
  "fmt"
  "log"
  "strings"
  "net/http"
  "os"
  "os/signal"
  "syscall"
  "context"
  "encoding/json"
  "sort"
  "strconv"

  "github.com/gorilla/mux"
  "github.com/scottberke/anagram_search/dictionary"
)

// Struct to hold our server and relevant attributes
type Server struct {
  http.Server

  // Channel used to broadcast a shutdown request when recieved via endpoint
  shutdown         chan bool

  // Channel connecting server to scope server is being run in so
  // that the encapsalating scope can block but be unblocked when server
  // shuts down in server package
  doneChan         chan<- bool

}

func NewServer(serverPort int, done chan<- bool) *Server {
    // Take port supplied via command line arg and build corresponding string
    // for the server to use
    var port strings.Builder
    fmt.Fprintf(&port, ":%d", serverPort)

    // Build our server
    mux := mux.NewRouter()

    server := &Server{
      Server:           http.Server{ Addr: port.String(), Handler: mux },
      shutdown:         make(chan bool),
      doneChan:         done,
    }

    // Setup our routes and handlers
    mux.HandleFunc("/anagrams/{word}.json", server.getAnagramsHandler).
      Methods("GET")

    mux.HandleFunc("/words.json", server.createAnagramsHandler).
      Methods("POST")

    mux.HandleFunc("/words.json", server.deleteAnagramsHandler).
      Methods("DELETE")

    mux.HandleFunc("/words/{word}.json", server.deleteSingleAnagramHandler).
      Methods("DELETE")

    mux.HandleFunc("/shutdown", server.shutdownHandler)

    return server
}

// Accept a string, downcase it, sort it and return
func getAnagramKey(word string) string {
  key_chars := strings.Split(word,"")
  sort.Strings(key_chars)
  key := strings.Join(key_chars,"")
  key = strings.ToLower(key)

  return key
}


func (server *Server) getAnagramsHandler(w http.ResponseWriter, r *http.Request) {
  // Grab our dictionary
  dict := dictionary.GetInstance()

  // Pull our our 'word' var from /anagrams/:word.json
  word := mux.Vars(r)["word"]

  // Get our anagram key
  key := getAnagramKey(word)

  // Build our response anagrams slice
  var anagrams []string

  // Check if we have a key matching the request word's key
  if val, ok := dict.Anagrams[key]; ok {
    // Build correct size array if we do
    anagrams = make([]string, 0, len(val))
    // Populate our response slice with all matches except request word
    for k := range val {
      if k != word {
        anagrams = append(anagrams, k)
      }
    }
  } else {
    // If we have no key, build empty array for response
    anagrams = make([]string,0,0)
  }

  // Build our response map
  res := make(map[string][]string)

  // Pull out any query strings from the request
  v := r.URL.Query()

  // If we have a limit query param, only return that many anagrams
  if val, ok := v["limit"]; ok {
    limit, _ := strconv.Atoi(val[0])
    // Return all anagrams if limit exceeds total anagrams returned
    if limit >= len(anagrams) {
      res["anagrams"] = anagrams
    } else {
      res["anagrams"] = anagrams[:limit]
    }
  } else {
    // Return all anagrams if no limit query string param
    res["anagrams"] = anagrams
  }

  // Turn our map into a marshalled JSON response
  data, err := json.Marshal(res)
  if err != nil {
    log.Printf("JSON Marshal Error: %v", err)
  }

  // Write headers and return
  w.Header().Set("Content-Type", "application/json")
  w.Write(data)
}

// Struct to decode JSON request body into
type words struct {
  Words []string
}

func (server *Server) createAnagramsHandler(w http.ResponseWriter, r *http.Request) {
  // Grab our dictionary
  dict := dictionary.GetInstance()

  // Parse our request boddy
  decoder := json.NewDecoder(r.Body)
  var new_words words
  err := decoder.Decode(&new_words)
  if err != nil {
    log.Printf("JSON Marshal Error: %v", err)
  }

  // Add our words to the dictionary and return Created
  dict.IngestFromArray(new_words.Words)
  w.WriteHeader(201)
}

func (server *Server) deleteAnagramsHandler(w http.ResponseWriter, r *http.Request) {
  // Grab our dictionary and reset to empty
  dict := dictionary.GetInstance()
  dict.ResetDictionary()

  // Respond with No Content
  w.WriteHeader(204)
}

func (server *Server) deleteSingleAnagramHandler(w http.ResponseWriter, r *http.Request) {
  // Grab our dictionary
  dict := dictionary.GetInstance()

  // Pull our our 'word' var from /words/:word.json
  word := mux.Vars(r)["word"]

  // Delete our word
  dict.DeleteSingleWord(word)

  // Respond with No Content
  w.WriteHeader(204)
}

func (server *Server) shutdownHandler(w http.ResponseWriter, r *http.Request) {
  log.Printf("Shutting down server.")
  w.Write([]byte(`{"Message": "Shutdown in progress. Requests Finishing"}`))

  // Unblock server shutdown so its actually called
  server.shutdown <- true
}


func (server *Server) ShutdownServer() {
  // Create a channel to listen to system interupts so this
  // can respond to both the shutdown endpoint and ctrl+c
  stop := make(chan os.Signal, 1)
  signal.Notify(stop, os.Interrupt, syscall.SIGTERM)

  // Block until we receive a shutdown signal to one of our channels
  select {
  case sig := <-stop:
     log.Printf("Shutdown request (signal: %v)", sig)
  case sig := <-server.shutdown:
     log.Printf("Shutdown request (/shutdown %v)", sig)
  }

  // Shutdown our server, gracefully waiting for requests to finish,
  // then close the blocking done channel in main() so the app terminates
  if err := server.Shutdown(context.Background()); err != nil {
      log.Printf("Shutdown request error: %v", err)
  }
  close(server.doneChan)

}
