package handlers

import (
	"errors"
	"github.com/bondarenkoi07/wbschool_task_2/dev/11_http/serializers"
	"github.com/bondarenkoi07/wbschool_task_2/dev/11_http/service"
	"net/http"
	"strconv"
	"strings"
)

var (
	invalidInputError = errors.New("check input data")
)

type Controller struct {
	service *service.Service
}

func NewController(service *service.Service) *Controller {
	return &Controller{service: service}
}

func (c *Controller) Create(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		err := r.ParseForm()

		if err != nil {
			http.Error(w, serializers.SerializeError(err), 500)
		}

		date := r.Form.Get("date")
		desc := r.Form.Get("desc")

		if date != "" && desc != "" {
			err = (*c).service.Create(desc, date)
			if err != nil {
				errorMsg := err.Error()
				if strings.HasPrefix(errorMsg, "InputError:") {
					http.Error(w, serializers.SerializeError(err), 400)
				} else if strings.HasPrefix(errorMsg, "Logic:") {
					http.Error(w, serializers.SerializeError(err), 503)
				} else {
					http.Error(w, serializers.SerializeError(err), 500)
				}
			} else {
				w.Header().Set("Content-Type", "application/json")
				_, err = w.Write([]byte(serializers.SerializeError(nil)))
				if err != nil {
					http.Error(w, serializers.SerializeError(err), 500)
				}
			}
		} else {
			http.Error(w, serializers.SerializeError(invalidInputError), 500)
		}
	}
}

func (c *Controller) Update(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		err := r.ParseForm()

		if err != nil {
			http.Error(w, serializers.SerializeError(err), 500)
		}

		id := r.Form.Get("id")
		date := r.Form.Get("date")
		desc := r.Form.Get("desc")

		Id, err := strconv.Atoi(id)

		if date != "" && desc != "" && err == nil && Id > 0 {
			err = (*c).service.Update(uint(Id), desc, date)
			if err != nil {
				errorMsg := err.Error()
				if strings.HasPrefix(errorMsg, "InputError:") {
					http.Error(w, serializers.SerializeError(err), 400)
				} else if strings.HasPrefix(errorMsg, "Logic:") {
					http.Error(w, serializers.SerializeError(err), 503)
				} else {
					http.Error(w, serializers.SerializeError(err), 500)
				}
			} else {
				w.Header().Set("Content-Type", "application/json")
				_, err = w.Write([]byte(serializers.SerializeError(nil)))
				if err != nil {
					http.Error(w, serializers.SerializeError(err), 500)
				}
			}
		} else {
			http.Error(w, serializers.SerializeError(invalidInputError), 500)
		}
	}
}

func (c *Controller) Delete(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		GET := r.URL.Query()
		id := GET.Get("id")

		Id, err := strconv.Atoi(id)

		if err == nil && Id > 0 {
			err = (*c).service.Delete(uint(Id))
			if err != nil {
				errorMsg := err.Error()
				if strings.HasPrefix(errorMsg, "InputError:") {
					http.Error(w, serializers.SerializeError(err), 400)
				} else if strings.HasPrefix(errorMsg, "Logic:") {
					http.Error(w, serializers.SerializeError(err), 503)
				} else {
					http.Error(w, serializers.SerializeError(err), 500)
				}
			} else {
				w.Header().Set("Content-Type", "application/json")
				_, err = w.Write([]byte(serializers.SerializeError(nil)))
				if err != nil {
					http.Error(w, serializers.SerializeError(err), 500)
				}
			}
		} else {
			http.Error(w, serializers.SerializeError(invalidInputError), 500)
		}
	}
}

func (c *Controller) Read(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		GET := r.URL.Query()
		id := GET.Get("id")

		Id, err := strconv.Atoi(id)

		if err == nil && Id > 0 {
			read, err := (*c).service.Read(uint(Id))

			if err != nil {
				errorMsg := err.Error()
				if strings.HasPrefix(errorMsg, "InputError:") {
					http.Error(w, serializers.SerializeError(err), 400)
				} else if strings.HasPrefix(errorMsg, "Logic:") {
					http.Error(w, serializers.SerializeError(err), 503)
				} else {
					http.Error(w, serializers.SerializeError(err), 500)
				}
			} else {
				w.Header().Set("Content-Type", "application/json")
				_, err = w.Write([]byte(read))
				if err != nil {
					http.Error(w, serializers.SerializeError(err), 500)
				}
			}
		} else {
			http.Error(w, serializers.SerializeError(invalidInputError), 500)
		}
	}
}

func (c *Controller) ReadAll(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		GET := r.URL.Query()
		id := GET.Get("id")

		Id, err := strconv.Atoi(id)

		if err == nil && Id > 0 {
			read := (*c).service.ReadAll()
			w.Header().Set("Content-Type", "application/json")
			_, err = w.Write([]byte(read))
			if err != nil {
				http.Error(w, serializers.SerializeError(err), 500)
			}
		} else {
			http.Error(w, serializers.SerializeError(invalidInputError), 500)
		}
	}
}

func (c *Controller) FilterByDay(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		GET := r.URL.Query()
		date := GET.Get("date")

		if date != "" {
			read, err := (*c).service.FilterByDay(date)
			if err != nil {
				errorMsg := err.Error()
				if strings.HasPrefix(errorMsg, "InputError:") {
					http.Error(w, serializers.SerializeError(err), 400)
				} else if strings.HasPrefix(errorMsg, "Logic:") {
					http.Error(w, serializers.SerializeError(err), 503)
				} else {
					http.Error(w, serializers.SerializeError(err), 500)
				}
			} else {
				w.Header().Set("Content-Type", "application/json")
				_, err = w.Write([]byte(read))
				if err != nil {
					http.Error(w, serializers.SerializeError(err), 500)
				}
			}
		} else {
			http.Error(w, serializers.SerializeError(invalidInputError), 500)
		}
	}
}

func (c *Controller) FilterByWeek(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		GET := r.URL.Query()
		week := GET.Get("date")

		if week != "" {
			read, err := (*c).service.FilterByDay(week)
			if err != nil {
				errorMsg := err.Error()
				if strings.HasPrefix(errorMsg, "InputError:") {
					http.Error(w, serializers.SerializeError(err), 400)
				} else if strings.HasPrefix(errorMsg, "Logic:") {
					http.Error(w, serializers.SerializeError(err), 503)
				} else {
					http.Error(w, serializers.SerializeError(err), 500)
				}
			} else {
				w.Header().Set("Content-Type", "application/json")
				_, err = w.Write([]byte(read))
				if err != nil {
					http.Error(w, serializers.SerializeError(err), 500)
				}
			}
		} else {
			http.Error(w, serializers.SerializeError(invalidInputError), 500)
		}
	}
}

func (c *Controller) FilterByMonth(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		GET := r.URL.Query()
		Month := GET.Get("date")

		if Month != "" {
			read, err := (*c).service.FilterByDay(Month)
			if err != nil {
				errorMsg := err.Error()
				if strings.HasPrefix(errorMsg, "InputError:") {
					http.Error(w, serializers.SerializeError(err), 400)
				} else if strings.HasPrefix(errorMsg, "Logic:") {
					http.Error(w, serializers.SerializeError(err), 503)
				} else {
					http.Error(w, serializers.SerializeError(err), 500)
				}
			} else {
				w.Header().Set("Content-Type", "application/json")
				_, err = w.Write([]byte(read))
				if err != nil {
					http.Error(w, serializers.SerializeError(err), 500)
				}
			}
		} else {
			http.Error(w, serializers.SerializeError(invalidInputError), 400)
		}
	}
}
