package db

import (
	"encoding/binary"
	"time"

	"github.com/boltdb/bolt"
)

var taskBucket = []byte("tasks")
var db *bolt.DB

var InitMock bool
var CreateMock bool
var AllMock bool

type Task struct {
	Key   int
	Value string
}

// Initialize a database and open or create a bucket if not exits
func Init(dbPath string) (*bolt.DB, error) {
	var err error
	db, err = bolt.Open(dbPath, 0600, &bolt.Options{Timeout: 1 * time.Second})
	if err != nil || InitMock {
		return nil, err
	}
	return db, db.Update(func(tx *bolt.Tx) error {
		_, err := tx.CreateBucketIfNotExists(taskBucket)
		return err
	})
}

// take an task as a string and create task
func CreateTask(task string) error {
	var id int
	err := db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket(taskBucket)
		id64, _ := b.NextSequence()
		id = int(id64)
		key := itob(id)
		return b.Put(key, []byte(task))
	})
	if err != nil || CreateMock {
		return err
	}
	return nil
}

// List all the tasks
func AllTasks() ([]Task, error) {
	var tasks []Task

	err := db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket(taskBucket)
		c := b.Cursor()
		for k, v := c.First(); k != nil; k, v = c.Next() {
			tasks = append(tasks, Task{
				Key:   btoi(k),
				Value: string(v),
			})
		}
		return nil
	})
	if err != nil || AllMock {
		return nil, err
	}
	return tasks, nil
}

// Delete Tasks
func DeleteTasks(key int) error {
	return db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket(taskBucket)
		return b.Delete(itob(key))
	})
}

// Take int as an input and convert it into byte slice
func itob(v int) []byte {
	b := make([]byte, 8)
	binary.BigEndian.PutUint64(b, uint64(v))
	return b
}

// Take byte slice as an input and convert it into int
func btoi(b []byte) int {
	return int(binary.BigEndian.Uint64(b))
}
