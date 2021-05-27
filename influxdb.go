package main

import (
	//	"fmt"
	"github.com/influxdata/influxdb1-client/v2"
	"log"
	"time"
)

// CREATE USER admin WITH PASSWORD '$the_usual' WITH ALL PRIVILEGES
// create database BLAH

func influxDBClient(config Config) client.Client {
	c, err := client.NewHTTPClient(client.HTTPConfig{
		Addr:     config.DatabaseURL,
		Username: config.DatabaseUser,
		Password: config.DatabasePassword,
	})
	if err != nil {
		log.Fatalln("Error: ", err)
	}
	return c
}

func influx_push_metrics(c client.Client, config Config, sessions_total int, sessions_in_use int) {
	bp, err := client.NewBatchPoints(client.BatchPointsConfig{
		Database:  config.DatabaseDatabase,
		Precision: "s",
	})

	if err != nil {
		log.Fatalln("Error: ", err)
	}

	eventTime := time.Now()

	/*
		Using "Line Protocol", eg: cpu,host=server02,region=uswest value=3 1434055562000010000
		http://goinbigdata.com/working-with-influxdb-in-go/

		key: arris
		tags: none
		fields: sessions_total=blah, etc.
		timestamp in seconds
	*/

	key := "arris"
	tags := map[string]string{}
	fields := map[string]interface{}{
		"sessions_total": sessions_total,
	}

	point, err := client.NewPoint(key, tags, fields, eventTime)
	if err != nil {
		log.Fatalln("Error: ", err)
	}
	if sessions_total != 0 {
		bp.AddPoint(point)
	}

	fields = map[string]interface{}{
		"sessions_in_use": sessions_in_use,
	}

	point, err = client.NewPoint(key, tags, fields, eventTime)
	if err != nil {
		log.Fatalln("Error: ", err)
	}
	if sessions_in_use != 0 {
		bp.AddPoint(point)
	}

	err = c.Write(bp)
	if err != nil {
		log.Fatal(err)
	}

}

func deliver_stats_to_influxdb(c client.Client, config Config, sessions_total int, sessions_in_use int) {

	influx_push_metrics(c, config, sessions_total, sessions_in_use)
}
