package options

import "bitbucket.org/Budry/availability-checker/src/sites"

type Options struct {
	Sites []sites.Site `json:"sites"`
}
