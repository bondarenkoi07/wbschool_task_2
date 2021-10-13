package handlers

import (
	"github.com/bondarenkoi07/wbschool_task_2/dev/11_http/calendar"
	"github.com/bondarenkoi07/wbschool_task_2/dev/11_http/service"
	"log"
	"net/http"
	"net/http/httptest"
	"net/url"
	"strings"
	"testing"
)

var controller = NewController(service.NewService(calendar.NewCalendar()))

func TestController_Create(t *testing.T) {

	req := httptest.NewRequest(http.MethodGet, "/api?date=blablabla&desc=foobar", nil)
	res := httptest.NewRecorder()

	controller.Create(res, req)

	if res.Code != http.StatusMethodNotAllowed {
		t.Error("wrong method passed!")
	}

	data := url.Values{
		"date": {"2021-30-10"},
		"desc": {"gardener"},
	}
	ValidReq := httptest.NewRequest(http.MethodPost, "/", strings.NewReader(data.Encode()))
	ValidReq.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	res = httptest.NewRecorder()
	controller.Create(res, ValidReq)

	if res.Code != http.StatusOK {
		t.Errorf("status must be Ok, got %v, %v", res.Code, res.Body.String())
	}

	data.Set("date", "foobar")
	InvalidReq := httptest.NewRequest(http.MethodPost, "/", strings.NewReader(data.Encode()))
	InvalidReq.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	res = httptest.NewRecorder()
	controller.Create(res, InvalidReq)

	if res.Code == http.StatusOK {
		t.Errorf("status must not be Ok, got %v", res.Code)
	}
}

func TestController_UpdateAndDelete(t *testing.T) {

	req := httptest.NewRequest(http.MethodGet, "/api?date=blablabla&desc=foobar", nil)
	res := httptest.NewRecorder()

	controller.Update(res, req)

	if res.Code != http.StatusMethodNotAllowed {
		t.Error("wrong method passed!")
	}

	data := url.Values{
		"date": {"2021-30-10"},
		"desc": {"gardener"},
		"id":   {"0"},
	}
	ValidReq := httptest.NewRequest(http.MethodPost, "/", strings.NewReader(data.Encode()))
	ValidReq.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	res = httptest.NewRecorder()
	controller.Update(res, ValidReq)

	if res.Code != http.StatusOK {
		t.Errorf("status must be Ok, got %v, %v", res.Code, res.Body.String())
	}

	DeleteValidReq := httptest.NewRequest(http.MethodGet, "/", strings.NewReader(data.Encode()))
	res = httptest.NewRecorder()

	q := DeleteValidReq.URL.Query()
	for k, v := range data {
		q.Add(k, v[0])
	}
	log.Println(q.Encode())
	DeleteValidReq.URL.RawQuery = q.Encode()

	controller.Delete(res, DeleteValidReq)
	if res.Code != http.StatusOK {
		t.Errorf("status must be Ok, got %v, %v", res.Code, res.Body.String())
	}

	data.Set("id", "foobar")

	InvalidReq := httptest.NewRequest(http.MethodPost, "/", strings.NewReader(data.Encode()))
	InvalidReq.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	res = httptest.NewRecorder()
	controller.Update(res, InvalidReq)

	if res.Code == http.StatusOK {
		t.Errorf("status must not be Ok, got %v", res.Code)
	}
}
