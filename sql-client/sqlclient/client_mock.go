package sqlclient

func AddMock(m Mock) {
	c := dbClient.(*clientMock)
	if c.mocks == nil {
		c.mocks = make(map[string]Mock, 0)
	}
	c.mocks[m.Query] = m
}

type clientMock struct {
	mocks map[string]Mock
}

type Mock struct {
	Query string
	Args  []interface{}
	Error error
}

func (c *clientMock) Query(query string, args ...interface{}) (rows, error) {
	return nil, nil
}
