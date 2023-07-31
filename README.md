With the aim of improving my skills as a Golang developer, I decided to create a more elaborate project, considering my lack of experience in the language. I sought an approach that would resemble the reality of a backend developer, and I would like to list the API functionalities to facilitate the identification of tasks that I can undertake as a developer.

First and foremost, the idea of the project is to create an application that allows tracking stocks in a portfolio. I chose this topic because I believe it's an area where I can explore various types of functionalities. It's essential to highlight that this project is focused on learning, so I tried to cover multiple areas at once to make the most of this learning opportunity.

----------
* Handlers

`Router:` 
All router functions are straightforward, simply adding the controllers for this group. These controllers are a custom struct specific to the project.

````
func NewHealthRouter() config.MotionController {
	return config.NewMotionController(
		"/health",
		config.NewMotionRouter(http.MethodGet, "/close", controller.NewHealthController().CloseHealth,
			middleware.Authorization(sqlDomain.Role{Name: sqlDomain.USER}, sqlDomain.Role{Name: sqlDomain.ADMIN}),
		),
		config.NewMotionRouter(http.MethodGet, "/open", controller.NewHealthController().OpenHealth),
	)
}
````

All the `controllers` and `services` are easy to understand.

For the implementation of these routers, I made a configuration in the file `internal/config/MotionGo.go`, which serves as an initial setup for the project.

----------
* Versioning
There is a function where all versions of the application are added.

```
func (m *MotionGo) AddRouter(version ...RoutersVersion) 
```
```

var version1 = config.RoutersVersion{
	Version: "v1",
	Handlers: []func() config.MotionController{
		router.NewHealthRouter,
	},
}

var version2 = config.RoutersVersion{
	Version: "v2",
	Handlers: []func() config.MotionController{
		router.NewHealthRouter,
	},
}
```
The intention is to facilitate the switching between different versions of the application.

----------
* Database
I created a database interface to simplify the insertion or modification of information. I must confess that I borrowed some ideas from Java's JPA (Java Persistence API).

```
func (s CompanyService) GetCompany(id int64) (sqlDomain.Company, error) {
	byId, err := s.companyRepository.FindById(id)
	if err != nil || byId.Status == domain.INACTIVE {
		return sqlDomain.Company{}, err
	}
	return byId, nil
}
```
MongoDB and Dependency Injection
It's great to know that the MongoDB implementation follows the same pattern, making life easier for the developer with dependency injection.


----------
* Middlewares

You've added middleware for logs and security, allowing users to have sessions.

----------
* Testes

One of the most significant challenges was creating configurations where you can use the same application to set up test environments. All tests are end-to-end (E2E) with an in-memory database for quick testing and a separate MongoDB database for more comprehensive testing. This is achieved by adding the necessary information to the `config.test.properties` file.

----------
* Next Steps:
The upcoming steps include enhancing data consistency in MongoDB, implementing Telemetry, Multi-Tenancy, and exploring Microservices architecture.
