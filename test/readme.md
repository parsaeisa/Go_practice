# Test
This repo is about writing unit tests for methods in golang using some tools . 

Tools : 
* Suite 
* Mockgen : to build mocks from interfaces . Explained [here](https://github.com/parsaeisa/Go_mock_try) .
* Mocks of other dependencies . 

## Suite
package : "github.com/stretchr/testify/suite"
At first you need a struct which has a `suite.Suite` to inherit from it and other properties that 
you want . 


At the end of file you should add this : 
```go
func TestBaseAPIPassengerRatingRepository(t *testing.T) {
	suite.Run(t, new(BaseAPIPassengerRatingRepositoryTestSuite))
}
```
## Testing Database
This section talks about unit tests for repositories . 

Packages : 
* database sql: The "database/sql" which I think is a built-in golang package .
* go sql mock : "github.com/DATA-DOG/go-sqlmock"

We should put expectations on : 
* Table rows
* The query which is sent to the database 
* The result 

For test suite we define a struct : 
```go
type FooTestSuite struct {
	suite.Suite
	db   *sql.DB
	mock sqlmock.Sqlmock
}
```

### Get requests
```go
rows := sqlmock.NewRows([]string{
    "id",
    "name",
    "height",
    "addr",
}).
    AddRow(
        customer.ID,
        customer.Name,
        customer.Height,
        customer.Addr,
    )

syntax := "SELECT .+ FROM customers WHERE .+"
suite.mock.ExpectQuery(syntax).
    WithArgs(customer.Height).
    WillReturnRows(rows)

expectedResult := model.ReasonCounter{
    model.Reason(counter.ReasonID): counter.Count,
}

data, err := suite.repo.GetCustomer(context.Background(), customer.ID)

require.NoError(err)
require.Equal(expectedResult, data)
require.NoError(suite.mock.ExpectationsWereMet())
```

### Insert requests
