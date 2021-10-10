package serializers

import (
	"fmt"
	"github.com/bondarenkoi07/wbschool_task_2/dev/11_http/calendar"
	"strings"
	"time"
)

func SerializeDate(value time.Time) string {
	return fmt.Sprintf("\"date\": %d-%d-%d", value.Year(), value.Day(), value.Month())
}

func SerializeEvent(event calendar.Event) string {
	return fmt.Sprintf("{%s, description: \"%s\"}", SerializeDate(event.Date), event.Description)
}

func SerializeSingleEvent(event calendar.Event) string {
	return fmt.Sprintf(`{"result": %s}`, SerializeEvent(event))
}

func SerializeError(err error) string {
	if err != nil {
		return fmt.Sprintf(`{"error": "%s"}`, err.Error())
	} else {
		return `{"error": "ok"}`
	}

}

func SerializeSlice(events []calendar.Event) string {
	var (
		index   int
		builder = strings.Builder{}
	)
	builder.Grow(len(`{"result": [`))
	builder.WriteString(`{"result": [`)

	for index = 0; index < len(events)-1; index++ {
		position := fmt.Sprintf("\"%d\": %s,", index, SerializeEvent(events[index]))
		builder.Grow(len(position))
		builder.WriteString(position)
	}

	position := fmt.Sprintf("\"%d\": %s,", index, SerializeEvent(events[index]))
	builder.Grow(len(position))
	builder.WriteString(position)

	builder.Grow(len(`]}`))
	builder.WriteString(`]}`)
	return builder.String()
}

func SerializeMap(events map[uint]calendar.Event) string {
	var builder = strings.Builder{}
	builder.Grow(len(`{"result": {`))
	builder.WriteString(`{"result": {`)
	for index, value := range events {
		position := fmt.Sprintf("\"%d\": %s,", index, SerializeEvent(value))
		if index == uint(len(events))-1 {
			position = position[:len(position)-1]
		}
		builder.Grow(len(position))
		builder.WriteString(position)
	}
	builder.Grow(len(`}}`))
	builder.WriteString(`}}`)
	return builder.String()
}
