package dictionary


import (
  "os"
  "bufio"
  "sort"
  "log"
  s "strings"
  "runtime"
	"path"
  "fmt"
  // "strings"
)

var Dictionary *dictionary

type dictionary struct {
  Anagrams   map[string]map[string]struct{}
}

func GetInstance() *dictionary {
  if Dictionary == nil {
    Dictionary = &dictionary{}
  }

  return Dictionary
}

func (p *dictionary) IngestFromFile() {
  dict := make(map[string]map[string]struct{})

  _, filename, _, ok := runtime.Caller(0)
  if !ok {
		panic("No caller information")
	}
  var dict_path s.Builder

  fmt.Fprintf(&dict_path, "%s/dictionary.txt", path.Dir(filename))

  file, err := os.Open(dict_path.String())
  if err != nil {
    log.Fatal(err)
  }
  defer file.Close()

  scanner := bufio.NewScanner(file)
  for scanner.Scan() {
    word := scanner.Text()
    key := getAnagramKey(word)

    if _, ok := dict[key]; ok {
      dict[key][word] = struct{}{}
    } else {
      set := make(map[string]struct{})
      set[word] = struct{}{}

      dict[key] = set
    }
  }

  p.Anagrams = dict
}

func (p *dictionary) IngestFromArray(words []string) {
  if p.Anagrams == nil {
    p.Anagrams = make(map[string]map[string]struct{})
  }
  for _, word := range words {
    key := getAnagramKey(word)

    if set, ok := p.Anagrams[key]; ok {
      set[word] = struct{}{}
    } else {
      set := make(map[string]struct{})
      set[word] = struct{}{}
      p.Anagrams[key] = set
    }
  }
}

func getAnagramKey(word string) string {
  key_chars := s.Split(word,"")
  sort.Strings(key_chars)
  key := s.Join(key_chars,"")
  key = s.ToLower(key)

  return key
}


func (p *dictionary) ResetDictionary() {
  // Reset dictionary to empty map
  p.Anagrams = make(map[string]map[string]struct{})
}

func (p *dictionary) DeleteSingleWord(word string) {
    // Get our anagram key
    key := getAnagramKey(word)

    if val, ok := p.Anagrams[key]; ok {
      delete(val, word)
    }
}
