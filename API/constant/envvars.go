package constant

import (
	"os"
)

var ATLAS_URI = os.Getenv("ATLAS_URI")
var AWS_BUCKET = os.Getenv("AWS_BUCKET")
var AWS_KEY = os.Getenv("AWS_KEY")
var AWS_ACCESS_KEY_ID = os.Getenv("AWS_ACCESS_KEY_ID")
var AWS_SECRET_ACCESS_KEY = os.Getenv("AWS_SECRET_ACCESS_KEY")
