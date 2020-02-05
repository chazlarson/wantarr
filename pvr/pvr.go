package pvr

import (
	"fmt"
	"github.com/l3uddz/wantarr/config"
	"strings"
	"time"
)

var (
	pvrDefaultPageSize      = 1000
	pvrDefaultSortKey       = "airDateUtc"
	pvrDefaultSortDirection = "desc"
)

type MediaItem struct {
	ItemId     int
	AirDateUtc time.Time
	LastSearch time.Time
}

type Interface interface {
	Init() error
	GetQueueSize() (int, error)
	GetWantedMissing() ([]MediaItem, error)
	GetWantedCutoff() ([]MediaItem, error)
	SearchMediaItems([]int) (bool, error)
}

/* Public */

func Get(pvrName string, pvrType string, pvrConfig *config.Pvr) (Interface, error) {

	switch strings.ToLower(pvrType) {
	case "sonarr":
		return NewSonarr(pvrName, pvrConfig), nil
	case "radarr":
		return NewRadarr(pvrName, pvrConfig), nil
	default:
		break
	}

	return nil, fmt.Errorf("unsupported pvr type provided: %q", pvrType)
}
