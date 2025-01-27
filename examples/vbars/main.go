// Copyright ©2018 Peter Paolucci. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"log"

	"github.com/vitalyisaev2/plotext/custplotter"
	"github.com/vitalyisaev2/plotext/examples"
	"gonum.org/v1/plot"
)

func main() {
	n := 60
	fakeTOHLCVs := examples.CreateTOHLCVExampleData(n)

	p, err := plot.New()
	if err != nil {
		log.Panic(err)
	}
	p.Title.Text = "Volume Bars"
	p.X.Label.Text = "Time"
	p.Y.Label.Text = "Volume"
	p.X.Tick.Marker = plot.TimeTicks{Format: "2006-01-02\n15:04:05"}

	bars, err := custplotter.NewVBars(fakeTOHLCVs)
	if err != nil {
		log.Panic(err)
	}

	p.Add(bars)

	err = p.Save(450, 150, "vbars.png")
	if err != nil {
		log.Panic(err)
	}
}
