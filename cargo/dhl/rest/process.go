package rest

import (
	"github.com/mstgnz/shipping/config"
)

func CreateAbroad(current cargo.Current, data map[string]any) (map[string]any, error) {
	parseAbroad(data)
	return nil, nil
}

func CreateDomestic(current cargo.Current, data map[string]any) (map[string]any, error) {
	parseDomestic(data)
	return nil, nil
}

func parseAbroad(data map[string]any) {

}

func parseDomestic(data map[string]any) {

}
