Łukasz Kozik - little proffesional experience with GoLang 

GO client library to access our fake [account API](http://api-docs.form3.tech/api.html#organisation-accounts) service.

# General account information
Account represents a bank account that is registered with Form3. It is used to validate and allocate inbound payments.
Account contains a list of attribute fields. The availibility of each field depends on the API call and scheme.

# ADR
REST client library should be simple in use.
Integration tests going to be done over account API service. 
Other unexcpected behaviour of account service API going to be covered with Unit tests with http service mock.
Library code coverage should be as high as possible - not lower than 90%

Because of relatively not much time for doing this. In first iteration gonna be covered most important functionalities
and most critical parts. Next will do refactoring for code clarity and cover uncovered parts of code.


# Faced issues 
For some reason when service gets build accountapi service runs and gets down for a moment and tests fails because of missing 
connection to the account service

```bash
tests_1       |     client_test.go:13: error returned: Get http://accountapi:8080/v1/organisation/accounts: dial tcp 192.168.128.4:8080: connect: connection refused
```

