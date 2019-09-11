package main

import (
	"fmt"
	"net/http"
)

func Cikti(value, value2 string) (veri string) {
	veri = value + value2
	return
}


func main() {

	type Requester interface {
		Get(url string) (resp *http.Response, err error)
	}
	
	type UserClient struct {
		BaseUrl                  string
		TaskCount, ResultPerPage int
	
		requester Requester
	}
	
	func (d *UserClient) Get() (response map[string]interface{}, err error) {
	}
	
	func (d *UserClient) FetchAll() (result []map[string]interface{}, err error) {
	}

	type UserClient struct {
		BaseUrl                  string
		TaskCount, ResultPerPage int
	}

	uc := &UserClient{
		BaseUrl:       "https://randomuser.me/api/",
		TaskCount:     20,
		ResultPerPage: 100,
		requester:     &http.Client{},
	}

	user, _ := uc.Get()
	users, _ := uc.FetchAll()
	fmt.Println(len(res)) // 2000
}

type FakeClient struct {
	Body  []byte
	Error error
}

func (c *FakeClient) Get(url string) (resp *http.Response, err error) {
	resp = &http.Response{
		Body: ioutil.NopCloser(bytes.NewBuffer(c.Body)),
	}
	return resp, c.Error
}

func TestGetUser(t *testing.T) {
	userclient := &UserClient{}

	t.Run("get userclient returns err on bad response", func(t *testing.T) {

		userclient.requester = &FakeClient{
			Body: []byte(`{"status": "ok"}`),
		}

		_, err := userclient.Get()
		if err == nil {
			t.Error("Expected to get an error")
		}
	})
}

func (c *FakeClient) Get(url string) (resp *http.Response, err error) {
	resp = &http.Response{
		Body: ioutil.NopCloser(bytes.NewBuffer(c.Body)),
	}
	return resp, c.Error
}

// func (d *UserClient) Get() (response map[string]interface{}, err error) {
// 	err := erros.New("Error")
// 	return
// }

// func (d *UserClient) Get() (response map[string]interface{}, err error) {
// 	url := fmt.Sprintf("%s?results=%d", d.BaseUrl, d.ResultPerPage)

// 	res, err := d.requester.Get(url)
// 	if err != nil {
// 		return nil, err
// 	}
// }

func (d *UserClient) Get() (response map[string]interface{}, err error) {
	url := fmt.Sprintf("%s?results=%d", d.BaseUrl, d.ResultPerPage)

	res, err := d.requester.Get(url)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()
	buf := new(bytes.Buffer)
	buf.ReadFrom(res.Body)
	newStr := buf.String()
	var data map[string]interface{}
	err = json.Unmarshal([]byte(newStr), &data)
	response = data
	return
}