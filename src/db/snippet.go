package db

import (
	"bytes"
	"encoding/gob"
	"log"
	"time"

	bolt "go.etcd.io/bbolt"
)

var code_snippet_bucket = []byte("code-snippets")
var db *bolt.DB

/*
type Snippet struct {
	Name     string `json:"name"` // Use appropriate tags for encoding/gob package
	Category string `json:"category"`
	Code     string `json:"code"`
}
*/

type Snippet struct {
	Name     string
	Category string
	Code     string
}

// making a map with a key of string and value of Snippet
var snippets = make(map[string]Snippet)

func Init(dbPath string) error {
	var err error
	db, err = bolt.Open(dbPath, 0600, &bolt.Options{Timeout: 1 * time.Second})
	if err != nil {
		return err
	}
	return db.Update(func(tx *bolt.Tx) error {
		_, err := tx.CreateBucketIfNotExists(code_snippet_bucket)
		return err
	})
}

func CreateSnippet(name, code, category string) (string, error) {
	// var id int
	var key string

	err := db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket(code_snippet_bucket)
		// id64, _ := b.NextSequence()
		// id = int(id64)
		// key := itob(id)
		// return b.Put(key, []byte(task))
		// creating the key contecanating category-name
		key = category + "-" + name
		s := Snippet{Name: name, Category: category, Code: code}
		encoded, err := EncodeSnippet(s)

		if err != nil {
			return err
		}
		if err := b.Put([]byte(key), encoded); err != nil {
			return err
		}
		snippets[string(key)] = s
		return nil
	})
	if err != nil {
		//log.Println("Error adding snippet:", err)
		return "error adding snippet", err
	} else {
		return key, nil
	}
}

func ListSnippets() ([]Snippet, error) {
	var snippets []Snippet

	err := db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket(code_snippet_bucket)
		c := b.Cursor()

		for key, value := c.First(); key != nil; key, value = c.Next() {
			var s Snippet
			err := DecodeSnippet(value, &s)
			if err != nil {
				return err
			}
			snippets = append(snippets, s)
		}

		return nil
	})

	if err != nil {
		return nil, err
	}

	return snippets, nil
}

func GetSnippet(name, category string) (Snippet, error) {
	key := category + "-" + name

	code, ok := snippets[key]
	if !ok {
		log.Printf("Snippet not found in the map for key %s", key)

		retrievedSnippet, err := RetrieveSnippetFromBoltDB(key)
		log.Printf("this is the snippet received %s", retrievedSnippet)
		if err != nil {
			log.Printf("Error retrieving snippet from BoltDB for key %s: %v", key, err)
			return Snippet{}, err
		}

		// Store the retrieved snippet in the map
		//	snippets[key] = retrievedSnippet
		//	log.Printf("Added snippet to the map for key %s", key)

		// Return the retrieved snippet
		return retrievedSnippet, nil
	}

	return code, nil
}

func RetrieveSnippetFromBoltDB(key string) (Snippet, error) {
	var s Snippet

	err := db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket(code_snippet_bucket)
		value := b.Get([]byte(key))
		if value != nil {
			log.Printf("Retrieved value from BoltDB for key %s: %v", key, value)

			err := DecodeSnippet(value, &s)
			if err != nil {
				log.Printf("Error decoding snippet for key %s: %v", key, err)
				return err
			}

			log.Printf("Decoded snippet for key %s: %+v", key, s)
		}
		return nil
	})

	if err != nil {
		log.Printf("Error retrieving snippet for key %s: %v", key, err)
		return Snippet{}, err
	}

	return s, nil
}

// turning snippet strct into slice of bytes for boltdb
func EncodeSnippet(s Snippet) ([]byte, error) {
	var buf bytes.Buffer
	enc := gob.NewEncoder(&buf)
	if err := enc.Encode(s); err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}

// decoding snippets fro a slice of bytes to struct
func DecodeSnippet(data []byte, s *Snippet) error {
	buf := bytes.NewReader(data)
	dec := gob.NewDecoder(buf)
	return dec.Decode(s)
}

func RemoveSnippet(name, category string) (string, error) {
	key := category + "-" + name

	// Check if the snippet exists in the map and BoltDB
	/*
		if _, ok := snippets[key]; !ok {
			log.Printf("Snippet with key \"%s\" not found in the map", key)
			return "", fmt.Errorf("snippet not found")
		}
	*/

	// Remove the snippet from the BoltDB and the map
	err := db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket(code_snippet_bucket)
		if err := b.Delete([]byte(key)); err != nil {
			log.Printf("Error deleting snippet with key \"%s\" from BoltDB: %v", key, err)
			return err
		}

		// Delete the snippet from the map
		// delete(snippets, key)
		return nil
	})

	if err != nil {
		log.Printf("Error removing snippet with key \"%s\": %v", key, err)
		return "", err
	}

	log.Printf("Removed \"%s\" from the \"%s\" category with key \"%s\".", name, category, key)
	return key, nil
}
