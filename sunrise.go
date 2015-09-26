// Copyright 2013 Travis Keep. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file or
// at http://opensource.org/licenses/BSD-3-Clause.

//This library is a modified version of the library at http://godoc.org/github.com/keep94/sunrise

// Package sunrise computes sunrises and sunsets using wikipedia article
// http://en.wikipedia.org/wiki/Sunrise_equation. Testing at my
// latitude and longitude in California shows that computed sunrises and
// sunsets can vary by as much as 2 minutes from those that NOAA reports
// at http://www.esrl.noaa.gov/gmd/grad/solcalc/sunrise.html.
package sunrise

import (
	"fmt"
	"math"
	"time"
)

const (
	jepoch = float64(2451545.0)
	uepoch = int64(946728000.0)
)

//Location gives sunrise and sunset times.
type Location struct {
	location        *time.Location
	latitude        float64
	longitude       float64
	jstar           float64
	solarNoon       float64
	hourAngleInDays float64
}

// NewLocation computes the sunrise and sunset times for latitude and longitude
// around currentTime. Generally, the computed sunrise will be no earlier
// than 24 hours before currentTime and the computed sunset will be no later
// than 24 hours after currentTime. However, these differences may exceed 24
// hours on days with more than 23 hours of daylight.
// The latitude is positive for north and negative for south. Longitude is
// positive for east and negative for west.
func NewLocation(latitude float64, longitude float64) *Location {
	l := &Location{
		location:  time.Now().Location(),
		latitude:  latitude,
		longitude: longitude,
		jstar:     jStar(longitude),
	}

	l.computeSolarNoonHourAngle()

	return l
}

//String interface to show location details
func (l *Location) String() string {
	return fmt.Sprintf("Calculation Details: Lat %.3f, Long %.3f", l.latitude, l.longitude)
}

//Today updates instance for calculation of today's sunrise and sunset
func (l *Location) Today() {
	l.jstar = jStar(l.longitude)

	l.computeSolarNoonHourAngle()
}

// AddDays computes the sunrise and sunset numDays after
// (or before if numDays is negative) the current sunrise and sunset at the
// same latitude and longitude.
func (l *Location) AddDays(numDays int) {
	l.jstar += float64(numDays)
	l.computeSolarNoonHourAngle()
}

// Sunrise returns the current computed sunrise. Returned sunrise has the same
// location as the time passed to Around.
func (l *Location) Sunrise() time.Time {
	return goTime(l.solarNoon-l.hourAngleInDays, l.location)
}

// Sunset returns the current computed sunset. Returned sunset has the same
// location as the time passed to Around.
func (l *Location) Sunset() time.Time {
	return goTime(l.solarNoon+l.hourAngleInDays, l.location)
}

func (l *Location) computeSolarNoonHourAngle() {
	ma := mod360(357.5291 + 0.98560028*(l.jstar-jepoch))
	center := 1.9148*sin(ma) + 0.02*sin(2.0*ma) + 0.0003*sin(3.0*ma)
	el := mod360(ma + 102.9372 + center + 180.0)
	l.solarNoon = l.jstar + 0.0053*sin(ma) - 0.0069*sin(2.0*el)
	declination := asin(sin(el) * sin(23.45))
	l.hourAngleInDays = acos((sin(-0.83)-sin(l.latitude)*sin(declination))/(cos(l.latitude)*cos(declination))) / 360.0
}

func julianDay(unix int64) float64 {
	return float64(unix-uepoch)/86400.0 + jepoch
}

func jStar(longitude float64) float64 {
	return math.Floor(
		julianDay(time.Now().Unix())-0.0009+longitude/360.0+0.5) + 0.0009 - longitude/360.0
}

func goTime(julianDay float64, loc *time.Location) time.Time {
	unix := uepoch + int64((julianDay-jepoch)*86400.0)
	return time.Unix(unix, 0).In(loc)
}

func sin(degrees float64) float64 {
	return math.Sin(degrees * math.Pi / 180.0)
}

func cos(degrees float64) float64 {
	return math.Cos(degrees * math.Pi / 180.0)
}

func asin(x float64) float64 {
	return math.Asin(x) * 180.0 / math.Pi
}

func acos(x float64) float64 {
	if x >= 1.0 {
		return 0.0
	}
	if x <= -1.0 {
		return 180.0
	}
	return math.Acos(x) * 180.0 / math.Pi
}

func mod360(x float64) float64 {
	return x - 360.0*math.Floor(x/360.0)
}
