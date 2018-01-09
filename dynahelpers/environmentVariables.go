package dynahelpers

// EnvironmentConfig contains all configuration variables that need to be set
type EnvironmentConfig struct {
	AwsRegion      string
	DynamoEndpoint string
}

var (
	dbEC           EnvironmentConfig
	dynamoInitChan chan struct{}
)

// SetEnvironmentVariables sets demoParkDB's necessary environment variables and returns a channel to the caller which gets called at the end of init to let it know it can proceed
func SetEnvironmentVariables(awsRegion, dynamoEndpoint string) chan struct{} {
	dbEC = EnvironmentConfig{
		AwsRegion:      awsRegion,
		DynamoEndpoint: dynamoEndpoint,
	}

	dynamoInitChan = make(chan struct{})
	return dynamoInitChan
}
