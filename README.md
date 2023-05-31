Com o objetivo de aprimorar minhas habilidades
como desenvolvedor em Golang, decidi criar um
projeto mais elaborado, levando em consideração
minha falta de experiência na linguagem. Busquei
uma abordagem que se aproximasse da realidade de um
desenvolvedor backend, e gostaria de listar as
funcionalidades da API para facilitar a identificação das
tarefas que posso desempenhar como desenvolvedor.

Antes de mais nada, a ideia do projeto é criar uma
aplicação que permita acompanhar ações em um portfólio.
Decidi abordar esse tema porque imagino que seja uma
área na qual posso explorar diversos tipos de
funcionalidades. É importante ressaltar que este
projeto é voltado para estudos, então tentei abordar
várias áreas de uma vez para aproveitar ao máximo essa
oportunidade de aprendizado.

----------
* Handlers

`Router:` todas as funcoes de routers estão simples,
apenas adicionar os controllers desse grupo, esses controllers
são uma struct propria do projeto

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

todos os `controllers` e `services` seguem facil de vizualização

para a implementação desses routers fiz uma configuração
no arquivo `internal/config/MotionGo.go` que é uma configuração
inicial do projeto

----------
* Versionamento

existe uma função aonde adiciona todas as verções da aplicação

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
o intuito é facilitar a troca de versão da aplicação


----------
* Banco de dados

fiz uma interface de banco de dados para facilitar a inserção
ou alteração de informação, confesso que copiei um pouco
do JPA do Java

```
func (s CompanyService) GetCompany(id int64) (sqlDomain.Company, error) {
	byId, err := s.companyRepository.FindById(id)
	if err != nil || byId.Status == domain.INACTIVE {
		return sqlDomain.Company{}, err
	}
	return byId, nil
}
```
lembrando que o mongo segue o mesmo padrão, isso facilita
muito a vida do desenvolvedor, tudo com injeção de dependências


----------
* Middlewares

adicionei middleware de logs e security (o usuario terá uma sessão)

----------
* Testes

Uma das maiores dificuldades foi criar configurações aonde
uso a mesma aplicaçao tambem para criar os ambientes de testes, todos os testes
sao E2E com banco de dados em memoria e no mongo um banco separado
apenas adicionando as informações no `config.test.properties`

----------
* Próximos passos: Consistência de dados no MongoDB, Telemetry, Multi Tenancy, Microservices,  