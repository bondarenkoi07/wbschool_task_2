package serializers

import (
	"encoding/json"
	"github.com/bondarenkoi07/wbschool_task_2/dev/11_http/calendar"
	"testing"
	"time"
)

func TestSerializeEvent(t *testing.T) {
	testData := []calendar.Event{
		{
			Date:        time.Now(),
			Description: "fooo",
		},
		{
			Date:        time.Now().AddDate(0, 1, 0),
			Description: "feeeee",
		},
		{
			Date:        time.Now().AddDate(0, 0, 1),
			Description: "fefefefe}{",
		},
	}

	var marshalledEvent = make([]string, 0, 3)
	for _, datum := range testData {
		marshalledEvent = append(marshalledEvent, SerializeEvent(datum))
	}

	for _, s := range marshalledEvent {
		var v interface{}
		err := json.Unmarshal([]byte(s), &v)
		if err != nil {
			t.Errorf("error: %v| %v", err, s)
		}
	}

	var v interface{}
	err := json.Unmarshal([]byte(SerializeSlice(testData)), &v)
	if err != nil {
		t.Errorf("error: %v| %v", err, SerializeSlice(testData))
	}

	var MapTestData = make(map[uint]calendar.Event)

	for i := 0; i < len(testData); i++ {
		MapTestData[uint(i)] = testData[i]
	}
	str := SerializeMap(MapTestData)
	err = json.Unmarshal([]byte(str), &v)
	if err != nil {
		t.Errorf("error: %v| %v", err, str)
	}
}
