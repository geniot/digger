package model

import (
	"path/filepath"
)

const (
	APP_NAME       = "Digger"
	APP_VERSION    = "0.1"
	CONF_FILE_NAME = ".digger.properties"

	WINDOW_XPOS_KEY   = "WINDOW_XPOS_KEY"
	WINDOW_YPOS_KEY   = "WINDOW_YPOS_KEY"
	WINDOW_WIDTH_KEY  = "WINDOW_WIDTH_KEY"
	WINDOW_HEIGHT_KEY = "WINDOW_HEIGHT_KEY"
)

var (
	PATH_TO_CONFIG = filepath.Join("${HOME}", CONF_FILE_NAME)
	WINDOW_TITLE   = APP_NAME + " " + APP_VERSION
)
