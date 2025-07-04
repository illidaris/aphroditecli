/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"io"
	"math/rand"
	"os"

	"github.com/go-echarts/go-echarts/v2/charts"
	"github.com/go-echarts/go-echarts/v2/components"
	"github.com/go-echarts/go-echarts/v2/opts"
	"github.com/spf13/cobra"
)

// chartsCmd represents the charts command
var chartsCmd = &cobra.Command{
	Use:   "charts",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		//TODO example: https://github.com/go-echarts/go-echarts https://github.com/go-echarts/examples
		f, _ := os.Create("bar.html")
		defer f.Close()

		bar := charts.NewBar()

		bar.SetGlobalOptions(charts.WithTitleOpts(opts.Title{
			Title:    "My first bar chart generated by go-echarts",
			Subtitle: "It's extremely easy to use, right?",
		}))
		bar.SetXAxis([]string{"苹果", "香蕉", "橘子"}).
			AddSeries("销量", []opts.BarData{
				{Value: 10}, {Value: 20}, {Value: 30},
			})

		bar.Render(f)

		line := charts.NewLine()
		line.SetGlobalOptions(charts.WithTitleOpts(opts.Title{
			Title:    "折线图",
			Subtitle: "It's extremely easy to use, right?",
		}))
		line.SetXAxis([]string{"1月", "2月", "3月", "4月"}).
			AddSeries("销售额", []opts.LineData{
				{Value: 500}, {Value: 800}, {Value: 600}, {Value: 900},
			})
		line.Render(f)

		gauge := charts.NewGauge()
		gauge.SetGlobalOptions(charts.WithTitleOpts(opts.Title{
			Title:    "仪表盘",
			Subtitle: "It's extremely easy to use, right?",
		}))
		gauge.AddSeries("速度", []opts.GaugeData{
			{Name: "运行速度", Value: 82},
		})
		gauge.Render(f)

		page := components.NewPage()
		page.AddCharts(
			mapBase(),
			mapShowLabel(),
			mapVisualMap(),
			mapRegion(),
			mapTheme(),
		)
		page.Render(io.MultiWriter(f))

	},
}

func init() {
	rootCmd.AddCommand(chartsCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// chartsCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// chartsCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

var (
	baseMapData = []opts.MapData{
		{Name: "北京", Value: float64(rand.Intn(150))},
		{Name: "上海", Value: float64(rand.Intn(150))},
		{Name: "广东", Value: float64(rand.Intn(150))},
		{Name: "辽宁", Value: float64(rand.Intn(150))},
		{Name: "山东", Value: float64(rand.Intn(150))},
		{Name: "山西", Value: float64(rand.Intn(150))},
		{Name: "陕西", Value: float64(rand.Intn(150))},
		{Name: "新疆", Value: float64(rand.Intn(150))},
		{Name: "内蒙古", Value: float64(rand.Intn(150))},
	}

	guangdongMapData = map[string]float64{
		"深圳市": float64(rand.Intn(150)),
		"广州市": float64(rand.Intn(150)),
		"湛江市": float64(rand.Intn(150)),
		"汕头市": float64(rand.Intn(150)),
		"东莞市": float64(rand.Intn(150)),
		"佛山市": float64(rand.Intn(150)),
		"云浮市": float64(rand.Intn(150)),
		"肇庆市": float64(rand.Intn(150)),
		"梅州市": float64(rand.Intn(150)),
	}
)

func generateMapData(data map[string]float64) (items []opts.MapData) {
	items = make([]opts.MapData, 0)
	for k, v := range data {
		items = append(items, opts.MapData{Name: k, Value: v})
	}
	return
}

func mapBase() *charts.Map {
	mc := charts.NewMap()
	mc.RegisterMapType("china")
	mc.SetGlobalOptions(
		charts.WithTitleOpts(opts.Title{Title: "basic map example"}),
	)

	mc.AddSeries("map", baseMapData)
	return mc
}

func mapShowLabel() *charts.Map {
	mc := charts.NewMap()
	mc.RegisterMapType("china")
	mc.SetGlobalOptions(
		charts.WithTitleOpts(opts.Title{Title: "show label"}),
	)

	mc.AddSeries("map", baseMapData).
		SetSeriesOptions(
			charts.WithLabelOpts(opts.Label{
				Show: opts.Bool(true),
			}),
		)
	return mc
}

func mapVisualMap() *charts.Map {
	mc := charts.NewMap()
	mc.RegisterMapType("china")
	mc.SetGlobalOptions(
		charts.WithTitleOpts(opts.Title{
			Title: "VisualMap",
		}),
		charts.WithVisualMapOpts(opts.VisualMap{
			Calculable: opts.Bool(true),
		}),
	)

	mc.AddSeries("map", baseMapData)
	return mc
}

func mapRegion() *charts.Map {
	mc := charts.NewMap()
	mc.RegisterMapType("广东")
	mc.SetGlobalOptions(
		charts.WithTitleOpts(opts.Title{
			Title: "Guangdong province",
		}),
		charts.WithVisualMapOpts(opts.VisualMap{
			Calculable: opts.Bool(true),
			InRange:    &opts.VisualMapInRange{Color: []string{"#50a3ba", "#eac736", "#d94e5d"}},
		}),
	)

	mc.AddSeries("map", generateMapData(guangdongMapData))
	return mc
}

func mapTheme() *charts.Map {
	mc := charts.NewMap()
	mc.RegisterMapType("china")
	mc.SetGlobalOptions(
		charts.WithInitializationOpts(opts.Initialization{
			Theme: "macarons",
		}),
		charts.WithTitleOpts(opts.Title{
			Title: "Map-theme",
		}),
		charts.WithVisualMapOpts(opts.VisualMap{
			Calculable: opts.Bool(true),
			Max:        150,
		}),
	)

	mc.AddSeries("map", baseMapData)
	return mc
}
