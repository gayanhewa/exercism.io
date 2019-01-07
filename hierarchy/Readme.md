## User Hierarchy

The solution is developed in a single package named hierarchy. There is no entry point for the application. You can run the test suite for the application with the following command 

```
go test -v --bench . --benchmem
```

Although there is a benchmark, it doesn't make much sense since there isn't any additional functions to compare benchmarks against. 

All the test cases can be found in the `cases_test.go` file. Feel free to add / remove and run the suite.

### Assumptions

- You have Go 1.11 running.
- RoleID's and UserID's will always be unsigned.
- The role relationships are not cyclic.
- Errors will only be thrown for invalid user id's
- The function assumes that the initial data given is valid. It does not validate the input data for the initial list of users and roles.
- I have replaced setUser / setRoles in favour of UserHierarchy.New factory method
