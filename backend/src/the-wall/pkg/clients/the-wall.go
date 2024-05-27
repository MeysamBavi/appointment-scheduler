package clients

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/url"
)

var (
	ErrUserNotFound = errors.New("user not found")
	ErrInternal     = errors.New("internal error")
)

type User struct {
	ID          uint
	PhoneNumber string
	Firstname   string
	Lastname    string
}

type TheWall struct {
	Client  http.Client
	Address string
}

func (t *TheWall) getUsers(ctx context.Context, values url.Values) (*User, error) {
	req, err := http.NewRequestWithContext(
		ctx,
		http.MethodGet,
		fmt.Sprintf("%s/users?%s", t.Address, values.Encode()),
		bytes.NewReader(nil),
	)
	if err != nil {
		return nil, err
	}
	resp, err := t.Client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("failed to call the-wall endpoint: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode == http.StatusOK {
		var u User
		err = unmarshalBody(resp.Body, &u)
		return &u, err
	}
	if resp.StatusCode == http.StatusNotFound {
		return nil, ErrUserNotFound
	}
	return nil, fmt.Errorf("%s: %w", resp.Status, ErrInternal)
}

func (t *TheWall) GetUserByPhoneNumber(ctx context.Context, phoneNumber string) (*User, error) {
	values := make(url.Values)
	values.Add("phone", phoneNumber)
	return t.getUsers(ctx, values)
}

func (t *TheWall) GetUserById(ctx context.Context, id uint) (*User, error) {
	values := make(url.Values)
	values.Add("id", fmt.Sprint(id))
	return t.getUsers(ctx, values)
}

func unmarshalBody(body io.Reader, p any) error {
	data, err := io.ReadAll(body)
	if err != nil {
		return err
	}
	return json.Unmarshal(data, p)
}
