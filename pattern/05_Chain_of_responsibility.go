package pattern

import (
	"bytes"
	"encoding/csv"
	"encoding/json"
	"encoding/xml"
	"errors"
	"io"
	"log"
)

type Unmarshalled map[string]interface{}

//Цепочка обязанностей — это поведенческий паттерн проектирования,
//который позволяет передавать запросы последовательно по цепочке обработчиков.
//Каждый последующий обработчик решает, может ли он обработать запрос сам и стоит
//ли передавать запрос дальше по цепи.

// Плюсы:
//     Уменьшает зависимость между клиентом и обработчиками.
//    Реализует принцип единственной обязанности.
//    Реализует принцип открытости/закрытости.
//Минусы:
//    Запрос может остаться никем не обработанным.

// DecoderNode - интерфейс произвольного обработчика. Если обработчик смог сериализовать данные,
//то выполнение закончивается, иначе передает испольнение другому обработчику,
//если ни один из обработчиков не справится с сериализацией, то вернет сообщение об этом.
type DecoderNode interface {
	Next() (Unmarshalled, error)
	Unmarshal() (Unmarshalled, error)
}

type JsonDecoder struct {
	decoder *json.Decoder
	next    *DecoderNode
}

func (j JsonDecoder) SetNext(next *DecoderNode) {
	j.next = next
}

func (j JsonDecoder) SetVal(val []byte) {
	reader := bytes.NewReader(val)
	j.decoder = json.NewDecoder(reader)
}

func (j *JsonDecoder) Next() (Unmarshalled, error) {
	return (*j.next).Unmarshal()
}

func (j JsonDecoder) Unmarshall() (Unmarshalled, error) {
	var data = make(Unmarshalled)
	err := j.decoder.Decode(&data)
	if j.next != nil {
		if err != nil {
			return j.Next()
		} else {
			return data, nil
		}
	} else {
		return nil, errors.New("data could not be unmarshalled")
	}
}

type XmlDecoder struct {
	decoder *xml.Decoder
	next    *DecoderNode
}

func (x XmlDecoder) SetVal(val []byte) {
	reader := bytes.NewReader(val)
	x.decoder = xml.NewDecoder(reader)
}

func (x XmlDecoder) SetNext(next *DecoderNode) {
	x.next = next
}

func (x XmlDecoder) Unmarshall() (Unmarshalled, error) {
	var data = make(Unmarshalled)
	err := x.decoder.Decode(&data)
	if x.next != nil {
		if err != nil {
			return x.Next()
		} else {
			return data, nil
		}
	} else {
		return nil, errors.New("data could not be unmarshalled")
	}

}

func (x *XmlDecoder) Next() (Unmarshalled, error) {
	return (*x.next).Unmarshal()
}

type CsvDecoder struct {
	decoder *csv.Reader
	next    *DecoderNode
}

func (c CsvDecoder) SetVal(val []byte) {
	c.decoder = csv.NewReader(bytes.NewBuffer(val))
}

func (c CsvDecoder) SetNext(next *DecoderNode) {
	c.next = next
}

func (c *CsvDecoder) Next() (Unmarshalled, error) {
	return (*c.next).Unmarshal()
}

func (c CsvDecoder) Unmarshall() (Unmarshalled, error) {
	var data = make(Unmarshalled)
	_, err := c.decoder.Read() // skip first line
	if err != nil {
		if err != io.EOF {
			log.Fatalln(err)
		}
	}

	for i := 0; err != io.EOF && err == nil; i++ {
		var line []string
		line, err = c.decoder.Read()
		data[string(rune(i))] = line
	}

	if err == io.EOF {
		err = nil
	}

	if c.next != nil {
		if err != nil {
			return c.Next()
		} else {
			return data, nil
		}
	} else {
		return nil, errors.New("data could not be unmarshalled")
	}
}
