package gridlocator

import (
	"math"
	"strconv"
	"strings"

	"github.com/pkg/errors"
)

// Coordinates contains latitude and longitude.
type Coordinates struct {
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
}

// Convert converts the specified latitude and longitude into the six
// digit Maidenhead grid locator.
func Convert(c *Coordinates) (string, error) {

	lat := c.Latitude + 90
	lng := c.Longitude + 180

	// Field
	lat = (lat / 10) // + 0.0000001;
	lng = (lng / 20) // + 0.0000001;
	val, err := n2l(int(math.Floor(lng)), true)
	if err != nil {
		return "", errors.Wrap(err, "field longitude")
	}
	locator := val
	val, err = n2l(int(math.Floor(lat)), true)

	if err != nil {
		return "", errors.Wrap(err, "field latitude")
	}
	locator += val

	// Square
	lat = 10 * (lat - math.Floor(lat))
	lng = 10 * (lng - math.Floor(lng))
	locator += strconv.Itoa(int(math.Floor(lng)))
	locator += strconv.Itoa(int(math.Floor(lat)))

	// Subsquare
	lat = 24 * (lat - math.Floor(lat))
	lng = 24 * (lng - math.Floor(lng))
	val, err = n2l(int(math.Floor(lng)), false)
	if err != nil {
		return "", errors.Wrap(err, "subsquare longitude")
	}
	locator += val
	val, err = n2l(int(math.Floor(lat)), false)
	if err != nil {
		return "", errors.Wrap(err, "subsquare latitude")
	}
	locator += val

	return locator, nil
}

// ConvertGridLocation converts a string grid location into latitude and longitude.
func ConvertGridLocation(location string) (float64, float64, error) {
	if len(location) != 4 && len(location) != 6 {
		return 0, 0, errors.New("grid location must be either 4 or 6 digits")
	}

	location = strings.ToLower(location)

	//lng = (($l[0] * 20) + ($l[2] * 2) + ($l[4]/12)  - 180);
	l := make([]int, 6)

	// Field
	var err error
	l[0], err = l2n(string(location[0]))
	if err != nil {
		return 0, 0, errors.Wrap(err, "longitude field value")
	}
	l[1], err = l2n(string(location[1]))
	if err != nil {
		return 0, 0, errors.Wrap(err, "latitude field value")
	}

	// Square
	val, err := strconv.ParseInt(string(location[2]), 10, 64)
	if err != nil {
		return 0, 0, errors.Wrap(err, "longitude sqare value")
	}
	l[2] = int(val)

	val, err = strconv.ParseInt(string(location[3]), 10, 64)
	if err != nil {
		return 0, 0, errors.Wrap(err, "latitude sqare value")
	}
	l[3] = int(val)

	if len(location) == 6 {
		// Subsquare
		l[4], err = l2n(string(location[4]))
		if err != nil {
			return 0, 0, errors.Wrap(err, "longitude subsquare value")
		}
		l[5], err = l2n(string(location[5]))
		if err != nil {
			return 0, 0, errors.Wrap(err, "latitude subsquare value")
		}
	}

	long := (float64(l[0]) * 20) + (float64(l[2]) * 2) + (float64(l[4]) / 12) - 180
	lat := (float64(l[1]) * 10) + float64(l[3]) + (float64(l[5]) / 24) - 90

	return lat, long, nil
}

// n2l checks if a given integer is in range of 0 to 23
// and converts it to an ASCII character (byte/uint8) in the range of a-x (0x61-0x78).
// An error is thrown if the number is out of bounds (>23).
func n2l(number int, uppercase bool) (string, error) {
	if number > (0x17) {
		return "", errors.New("number out of bounds")
	}

	n := number + 0x61

	if uppercase {
		return strings.ToUpper(string(n)), nil
	} else {
		return string(n), nil
	}
}

var let2num = map[string]int{
	`a`: 0,
	`b`: 1,
	`c`: 2,
	`d`: 3,
	`e`: 4,
	`f`: 5,
	`g`: 6,
	`h`: 7,
	`i`: 8,
	`j`: 9,
	`k`: 10,
	`l`: 11,
	`m`: 12,
	`n`: 13,
	`o`: 14,
	`p`: 15,
	`q`: 16,
	`r`: 17,
	`s`: 18,
	`t`: 19,
	`u`: 20,
	`v`: 21,
	`w`: 22,
	`x`: 23,
}

func l2n(letter string) (int, error) {
	val, ok := let2num[letter]
	if !ok {
		return 0, errors.New("illegal character")
	}
	return val, nil
}
