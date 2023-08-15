package main

import (
	"encoding/csv"
	"fmt"
	"github.com/gocarina/gocsv"
	"github.com/urfave/cli"
	"github.com/yokogawa-k/sysbench-output-parser/dbreport"
	"io"
	"log"
	"os"
	"strings"
)

type output struct {
	Version           string  `csv:"version"`
	Threads           uint64  `csv:"threads"`
	TPS               float64 `csv:"tps"`
	QPS               float64 `csv:"qps"`
	Read              uint64  `csv:"read"`
	Write             uint64  `csv:"write"`
	Other             uint64  `csv:"other"`
	Total             uint64  `csv:"total"`
	IgnoredErrors     uint64  `csv:"ignored-errors"`
	Recoonects        uint64  `csv:"recoonects"`
	MinLatency        float64 `csv:"Latency(min)"`
	AvgLatency        float64 `csv:"Latency(avg)"`
	MaxLatency        float64 `csv:"Latency(max)"`
	PercentileLatency string  `csv:"Latency(Percentile)"`
}

func main() {
	app := &cli.App{
		Name:   "kubeblocks",
		Action: parseAndPrint,
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name: "file",
			},
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}

func parseAndPrint(c *cli.Context) error {
	fileName := c.String("file")

	readers, err := parse(fileName)
	if err != nil {
		return err
	}

	return print2(readers)
}

func parse(fileName string) ([]io.Reader, error) {
	content, err := ReadAsString(fileName)
	if err != nil {
		return nil, err
	}
	readers, err := partitionAsReaders(content)
	if err != nil {
		return nil, err
	}
	return readers, nil
}

func print(readers []io.Reader) error {
	for _, r := range readers {
		s, err := dbreport.ParseDBReport(r)
		if err != nil {
			return err
		}

		percentileLatencyFormat := fmt.Sprintf("%.2f (%d%%)", *s.Latency.Percentile, *s.Latency.LatencyPct)
		o := []*output{
			{
				Version:           s.Version.VersionNumber,
				Threads:           s.NumThreads,
				TPS:               *s.SQLStatistics.TransactionsInfo.TransactionsPerSecond,
				QPS:               *s.SQLStatistics.Queries.QueriesPerSecond,
				Read:              *s.SQLStatistics.ReadQueries,
				Write:             *s.SQLStatistics.WriteQueries,
				Other:             *s.SQLStatistics.OtherQueries,
				Total:             *s.SQLStatistics.TotalQueries,
				IgnoredErrors:     *s.SQLStatistics.IgnoredErrors.IgnoredErrors,
				Recoonects:        *s.SQLStatistics.Reconnects.Reconnects,
				MinLatency:        *s.Latency.Minimum,
				AvgLatency:        *s.Latency.Average,
				MaxLatency:        *s.Latency.Maximum,
				PercentileLatency: percentileLatencyFormat,
			},
		}

		gocsv.SetCSVWriter(func(out io.Writer) *gocsv.SafeCSVWriter {
			writer := csv.NewWriter(out)
			writer.Comma = '\t'
			return gocsv.NewSafeCSVWriter(writer)
		})

		cs, err := gocsv.MarshalString(o)
		if err != nil {
			return err
		}

		fmt.Println(cs)
	}
	return nil
}

func print2(readers []io.Reader) error {
	// 定义CSV header
	header := []string{"Version", "Threads", "TPS", "QPS", "Read", "Write", "Other", "Total", "IgnoredErrors", "Recoonects", "MinLatency", "AvgLatency", "MaxLatency"}

	// 打印CSV header
	fmt.Println(strings.Join(header, "\t"))

	for _, r := range readers {
		s, err := dbreport.ParseDBReport(r)
		if err != nil {
			return err
		}

		o := []*output{
			{
				Version:       s.Version.VersionNumber,
				Threads:       s.NumThreads,
				TPS:           *s.SQLStatistics.TransactionsInfo.TransactionsPerSecond,
				QPS:           *s.SQLStatistics.Queries.QueriesPerSecond,
				Read:          *s.SQLStatistics.ReadQueries,
				Write:         *s.SQLStatistics.WriteQueries,
				Other:         *s.SQLStatistics.OtherQueries,
				Total:         *s.SQLStatistics.TotalQueries,
				IgnoredErrors: *s.SQLStatistics.IgnoredErrors.IgnoredErrors,
				Recoonects:    *s.SQLStatistics.Reconnects.Reconnects,
				MinLatency:    *s.Latency.Minimum,
				AvgLatency:    *s.Latency.Average,
				MaxLatency:    *s.Latency.Maximum,
			},
		}

		// 打印数据行
		for _, output := range o {
			fmt.Printf("%s\t%d\t%.2f\t%.2f\t%d\t%d\t%d\t%d\t%d\t%d\t%.2f\t%.2f\t%.2f\t\n",
				output.Version, output.Threads, output.TPS, output.QPS, output.Read, output.Write, output.Other, output.Total, output.IgnoredErrors, output.Recoonects, output.MinLatency, output.AvgLatency, output.MaxLatency)
		}
	}
	return nil
}

func ReadAsString(fileName string) (string, error) {
	// 读取文件内容到字符串
	content, err := os.ReadFile(fileName)
	if err != nil {
		return "", fmt.Errorf("Error reading file:%v", err)
	}
	return string(content), nil
}
