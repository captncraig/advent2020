package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"

	"github.com/PuerkitoBio/goquery"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func main() {
	prometheus.Register(counts)
	prometheus.Register(collector{})
	log.Println(":2112")
	http.Handle("/metrics", promhttp.Handler())
	http.ListenAndServe(":2112", nil)
}

var counts = prometheus.NewGaugeVec(
	prometheus.GaugeOpts{
		Name: "aoc_gold_stars",
		Help: "Number of gold stars by day",
	},
	[]string{"year", "day"},
)

type collector struct{}

func (c collector) Describe(chan<- *prometheus.Desc) {}

func (c collector) Collect(ch chan<- prometheus.Metric) {
	collect("2020")
	collect("2019")
	collect("2018")
	collect("2017")
	collect("2016")
	collect("2015")
}

func collect(year string) error {
	res, err := http.Get(fmt.Sprintf("https://adventofcode.com/%s/stats", year))
	if err != nil {
		return err
	}
	defer res.Body.Close()
	if res.StatusCode != 200 {
		return fmt.Errorf("status code error: %d %s", res.StatusCode, res.Status)
	}

	// Load the HTML document
	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		return err
	}

	both := []int{}
	doc.Find(".stats-both").Each(func(i int, s *goquery.Selection) {
		if i, err := strconv.Atoi(strings.TrimSpace(s.Text())); err == nil {
			both = append(both, i)
		}
	})
	for i, n := range both {
		counts.WithLabelValues(year, fmt.Sprint(25-i)).Set(float64(n))
	}

	return nil
}
