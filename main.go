package goatlassianapi

type ApiClient struct {
	baseurl string
	email   string
	token   string
}

func NewApiClient(baseurl, email, token string) *ApiClient {
	return &ApiClient{
		baseurl: baseurl,
		email:   email,
		token:   token,
	}
}

func (a *ApiClient) BaseUrl() string {
	return a.baseurl
}

func (a *ApiClient) Email() string {
	return a.email
}

func (a *ApiClient) Token() string {
	return a.token
}
