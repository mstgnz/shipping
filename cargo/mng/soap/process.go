package soap

import (
	cargo "github.com/mstgnz/shipping/config"
)

func Process(mode string, current cargo.Current, data map[string]any) (map[string]any, error) {
	parse(data)
	return nil, nil
}

func parse(data map[string]any) {

}
