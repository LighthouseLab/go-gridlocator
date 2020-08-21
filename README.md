# Repository Description

This repository contains a Golang implementation which converts regular longitude and latitude coordinates to the Maidenhead Locator System and vice versa.

Based on the algorithm published by Edmund T. Tyson, N5JTY, titled "Conversion Between Geodetic and Grid Locator Systems" in _QST_ January 1989, pp. 29-30, 43.

[![GoDoc](https://godoc.org/github.com/LighthouseLab/go-maidenhead?status.svg)](https://godoc.org/github.com/LighthouseLab/go-maidenhead)
[![Go Report Card](https://goreportcard.com/badge/github.com/LighthouseLab/go-maidenhead)](https://goreportcard.com/report/github.com/LighthouseLab/go-maidenhead)
[![Build Status](https://travis-ci.org/LighthouseLab/go-maidenhead.svg?branch=master)](https://travis-ci.org/LighthouseLab/go-maidenhead)

# What is the Maidenhead Locator System?

The Maidenhead Locator System is a grid system which divides the earth in fields, squares, subsquares and extended subsquares. The extended subsquares could be split again in extended subsquares.

![Example of Grid System](https://upload.wikimedia.org/wikipedia/commons/1/1d/Maidenhead_grid_over_Europe.svg)

If we want the Maidenhead Locator string of Londen, it would be:

| Field | Square | Subsquare | Ext. Subsq. |
| ----- | ------ | --------- | ----------- |
| IO    | 91     | xm        | 02          |