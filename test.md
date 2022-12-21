# Test

Writing tests ( unit-tests ) in golang . 

This is the package that we are using : "github.com/stretchr/testify/suite"

## 1. Test struct

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

## 2. Setup test

This is the stage that we initiate our dependencies in :
```go
func (suite *CarsTestSuite) SetupTest() {
	require := suite.Require()
	var err error

	suite.db, suite.mock, err = sqlmock.New()
	require.NoError(err)

	suite.repo = &baseAPIPassengerRating{
		db: sqlx.NewDb(suite.db, "mysql"),
	}
}

```

## 3. Testing

Now for each test , we assign a new method to our `suite` and call our method and its dependency there . 

```go
func (suite *CarsTestSuite) TestCase1_Success() {
	require := suite.Require()

	// define the behavior of dependency mocks
	
	// call method 
	
	// define restrictions
}
```

We have two types of test cases : 
* Success : We expect the result to be something specific . 
* Failure : We expect the method returns an error .

### Succes

### Failure

## 4. Run tests
In order to run all of tests , we should use these line : 
```go
func TestCars(t *testing.T) {
	suite.Run(t, new(CarsTestSuite))
}
```