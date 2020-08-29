package dbtest

/*
	Maps dependecies between sql scripts.

	key = the maindependecy name.
	value = array of other scripts that this one depends.

	The ideia is to quickly generate the minimum amount of data needed
	to run a particular test, by specifying only the direct dependencies.
*/
var dependencies = map[string][]string{
	"users":    []string{"schema"},
	"messages": []string{"schema", "users", "queues", "users_queues"},
}
