package constants

const (
	DatabaseDriverName = "sqlite3"
	DatabasePath       = "/database/database.db"

	Root      string = "/"
	RootRoute string = "/dont/api/v1/"

	AuthRoute         string = "/auth"
	LoginRoute        string = "/login"
	RefreshTokenRoute string = "/token/refresh"

	SubjectRoute string = "/subjects"

	ThemeRoute string = "/themes"

	QuestionRoute string = "/questions"

	TestRoute string = "/tests"

	Port string = ":32678"
)
