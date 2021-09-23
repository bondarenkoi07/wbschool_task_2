package pattern

import (
	"errors"
	"fmt"
	"os"
)

type User struct {
	Authed    bool
	Moderate  bool
	Superuser bool
}

type Status interface {
	Publish(text string) error
	Read() error
	Delete(id uint) error
	Next() Status
	LogOut() Status
}

type StatusLoggedOut struct {
	user User
	next *StatusLoggedIn
}

func (s *StatusLoggedOut) SetUser(user User) {
	s.user = user
}

func (s StatusLoggedOut) Publish(string) error {
	return errors.New("you are not logged in")
}

func (s StatusLoggedOut) Read() error {
	var a interface{}
	_, err := fmt.Fscan(os.Stdin, &a)
	if err != nil {
		return err
	}
	fmt.Println(a)
	return nil
}

func (s StatusLoggedOut) Delete(uint) error {
	return errors.New("you are not logged in")
}

func (s *StatusLoggedOut) Next() Status {
	if s.user.Authed {
		return s.next
	}
	return s
}

func (s *StatusLoggedOut) LogOut() Status {
	return s
}

type StatusLoggedIn struct {
	user      User
	privilege *StatusPrivileged
	moderate  *StatusModerate
	logout    *StatusLoggedOut
}

func (s *StatusLoggedIn) SetUser(user User) {
	s.user = user
}

func (s StatusLoggedIn) Publish(string) error {
	return errors.New("you don't have permission")
}

func (s StatusLoggedIn) Read() error {
	var a interface{}
	_, err := fmt.Fscan(os.Stdin, &a)
	if err != nil {
		return err
	}
	fmt.Println(a)
	return nil
}

func (s StatusLoggedIn) Delete(uint) error {
	return errors.New("you do not have permission to do that")
}

func (s StatusLoggedIn) Next() Status {
	if s.user.Superuser {
		return s.privilege
	} else if s.user.Moderate {
		return s.moderate
	}
	return s
}

func (s StatusLoggedIn) LogOut() Status {
	return s.logout
}

type StatusPrivileged struct {
	user   User
	logout *StatusLoggedOut
}

func (s *StatusPrivileged) SetUser(user User) {
	s.user = user
}

func (s StatusPrivileged) Publish(text string) error {
	_, err := fmt.Fprint(os.Stdout, text)
	if err != nil {
		return err
	}
	return nil
}

func (s StatusPrivileged) Read() error {
	var a interface{}
	_, err := fmt.Fscan(os.Stdin, &a)
	if err != nil {
		return err
	}
	fmt.Println(a)
	return nil
}

func (s StatusPrivileged) Delete(count uint) error {
	for i := count; i > 0; i-- {
		_, err := fmt.Fprint(os.Stdout, '\b')
		if err != nil {
			return err
		}
	}
	return nil
}

func (s StatusPrivileged) Next() Status {
	return s.LogOut()
}

func (s StatusPrivileged) LogOut() Status {
	return s.logout
}

type StatusModerate struct {
	user   User
	logout *StatusLoggedOut
	next   *StatusPrivileged
}

func (s *StatusModerate) SetUser(user User) {
	s.user = user
}

func (s StatusModerate) Publish(text string) error {
	_, err := fmt.Fprint(os.Stdout, text)
	if err != nil {
		return err
	}
	return nil
}

func (s StatusModerate) Read() error {
	var a interface{}
	_, err := fmt.Fscan(os.Stdin, &a)
	if err != nil {
		return err
	}
	fmt.Println(a)
	return nil
}

func (s StatusModerate) Delete(uint) error {
	return errors.New("you do not have permission to do that")
}

func (s StatusModerate) Next() Status {
	if s.user.Superuser {
		return s.next
	}
	return s.LogOut()
}

func (s StatusModerate) LogOut() Status {
	return s.logout
}
