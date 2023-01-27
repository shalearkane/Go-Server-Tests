package constant

import (
	"os"
)

var ATLAS_URI = os.Getenv("ATLAS_URI")
var AWS_BUCKET = os.Getenv("AWS_BUCKET")
var AWS_KEY = os.Getenv("AWS_KEY")
