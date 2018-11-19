package healthchecker

type healthchecker interface {
	isHealthy() bool
	getInterval() int
}
