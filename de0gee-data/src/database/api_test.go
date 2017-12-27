package database

import (
	"encoding/json"
	"fmt"
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAddSensor(t *testing.T) {
	var s SensorData
	err := json.Unmarshal([]byte(j), &s)
	if err != nil {
		panic(err)
	}
	db, _ := Open("testing")
	defer db.Close()
	err = db.AddSensor(s)
	assert.Nil(t, err)

	s2, err := db.GetSensorFromTime(s.Timestamp)
	assert.Nil(t, err)
	assert.Equal(t, s, s2)
	fmt.Println(s2)
}

func TestGetAllForClassification(t *testing.T) {
	var err error
	var s SensorData
	db, _ := Open("testing")
	defer db.Close()
	json.Unmarshal([]byte(j), &s)
	err = db.AddSensor(s)
	assert.Nil(t, err)
	json.Unmarshal([]byte(j2), &s)
	err = db.AddSensor(s)
	assert.Nil(t, err)

	ss, err := db.GetAllForClassification()
	assert.Equal(t, 2, len(ss))
	assert.Nil(t, err)
}

func BenchmarkAddSensor(b *testing.B) {
	var s SensorData
	json.Unmarshal([]byte(j), &s)
	db, _ := Open("testing")
	defer db.Close()
	Debug(false)

	for i := 0; i < b.N; i++ {
		s.Timestamp = float64(i)
		err := db.AddSensor(s)
		if err != nil {
			panic(err)
		}
	}
}

func BenchmarkGetSensor(b *testing.B) {
	var s SensorData
	err := json.Unmarshal([]byte(j), &s)
	if err != nil {
		panic(err)
	}
	db, _ := Open("testing")
	defer db.Close()
	Debug(false)
	err = db.AddSensor(s)
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		_, err := db.GetSensorFromTime(s.Timestamp)
		if err != nil {
			panic(err)
		}
	}
}
func BenchmarkKeystoreSet(b *testing.B) {
	db, _ := Open("testing")
	defer db.Close()
	Debug(false)
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		err := db.Set("human:"+strconv.Itoa(i), Human{"Dante", 5.4})
		if err != nil {
			panic(err)
		}
	}
}

func BenchmarkKeystoreOpenAndSet(b *testing.B) {
	for i := 0; i < b.N; i++ {
		db, _ := Open("testing")
		Debug(false)
		err := db.Set("human:"+strconv.Itoa(i), Human{"Dante", 5.4})
		if err != nil {
			panic(err)
		}
		db.Close()
	}
}

func BenchmarkKeystoreGet(b *testing.B) {
	db, _ := Open("testing")
	defer db.Close()
	Debug(false)
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		var h2 Human
		db.Get("human:"+strconv.Itoa(i), &h2)
	}
}
