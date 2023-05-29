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

## 2. Setup suite

This is the stage that we initiate our dependencies in :

```go
func (suite ServerTestSuite) SetupSuite() {

}
```

> Read about 'gomock' here . (add the link)

## 3. Setup test

I think this method is called before each unit-test . 

```go
func (suite *CarsTestSuite) SetupTest() {
	ctrl := gomock.NewController(suite.T())
	defer ctrl.Finish()

	require := suite.Require()
	var err error

	suite.db, suite.mock, err = sqlmock.New()
	require.NoError(err)

	suite.repo = &baseAPIPassengerRating{
		db: sqlx.NewDb(suite.db, "mysql"),
	}
}

```

> Put the creation of ctrl and mocks in SetupTest.

## 4. Testing

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

## 4. TearDown

```go
func (suite *DriverCancellationReasonTestSuite) TearDownSuite() { //monkey.UnpatchAll()
	suite.ctrl.Finish()
}
```

### Succes

### Failure

## 5. Run tests
In order to run all of tests , we should use these line : 
```go
func TestCars(t *testing.T) {
	suite.Run(t, new(CarsTestSuite))
}
```