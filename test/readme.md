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

Writing tests ( unit-tests ) in golang . 

This is the package that we are using : "github.com/stretchr/testify/suite"

### 1. Test struct

This is a struct which we define and inherits suite.Suite . 

```go
type CarsTestSuite struct {
	suite.Suite
	db   *sql.DB
	mock sqlmock.Sqlmock
	repo *Customers
}
```

We add our dependencies in this struct .

### 2. Setup suite

This method runs just one time at the beginning.

```go
func (suite ServerTestSuite) SetupSuite() {

}
```

> Read about 'gomock' here . (add the link)

### 3. Setup test

This method is called many times, before each unit-test . 

```go
func (suite *CarsTestSuite) SetupTest() {
	ctrl := gomock.NewController(suite.T())
	defer ctrl.Finish()

	var err error

	suite.db, suite.mock, err = sqlmock.New()
	require.NoError(err)

	suite.repo = &baseAPIPassengerRating{
		db: sqlx.NewDb(suite.db, "mysql"),
	}
}

```

> Put the creation of ctrl and mocks in SetupTest.

### 4. Testing

Now for each test , we assign a new method to our `suite` and call our method and its dependency there . 

```go
func (suite *CarsTestSuite) TestCase1_Success() {
	require := suite.Require()

	// define the behavior of dependency mocks
	
	// call method 
	
	// define restrictions
}
```

This require has some usefull methods : 
* NoError
* Equal
* EqualError
* JSONEqual

We have two types of test cases : 
* Success : We expect the result to be something specific . 
* Failure : We expect the method returns an error .

### 4. TearDown

```go
func (suite *DriverCancellationReasonTestSuite) TearDownSuite() { //monkey.UnpatchAll()
	suite.ctrl.Finish()
}
```

#### Succes

#### Failure

### 5. Run tests
In order to run all of tests , we should use these line : 
```go
func TestCars(t *testing.T) {
	suite.Run(t, new(CarsTestSuite))
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

## Testing http handler

You need a `ctx` for the handler that you want to test. This context must be fake, here's a way to create a fake context using `http.Recorder` : 

```go
func SubmitDriverCancellationReasonsV3NewEchoContext(
	/* arguments in url */) (echo.Context, *httptest.ResponseRecorder) {
	
    // Creating request and it's params
	url := /* Your url */
	request := httptest.NewRequest(http.MethodPatch, url, nil)
	request.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

	response := httptest.NewRecorder()
	e := echo.New()
	ctx := e.NewContext(request, response)

    // Setting params
	ctx.SetParamNames(
		"id",
		"type",
	)
	ctx.SetParamValues(
		id,
		rejectReasonType,
	)
	ctx.Set(httpctx.UserIDContextField, userID)

	return ctx, response
}
```

