//Data cleansing/extraction is up to the user.
//Remember, garbage in garbage out
package main

import (
	"github.com/antonholmquist/jason"
	"log"
)

func main() {

	exampleJSON :=
		`
[{
	"measurement": "measurement name",
	"tags": {
		"tagkey1": "tagvalue1",
		"tagkeyN": "tagvalueN"
	},
	"fields": {
		"fieldkey1": "fieldvalue1",
		"fieldkeyN": "fieldvalueN"
	},
	"time": "timeString or timeUNIXepoch"
}, {
	"measurement": "measurement name 2",
	"fields": {
		"idle": 53.3,
		"s1": "chasney",
		"s2": "chelsea",
		"friends": true,
		"duration": 135
	}
}]
`
	var v *jason.Value
	var err error

	v, err = jason.NewValueFromBytes([]byte(exampleJSON))

	if err != nil {
		log.Fatal(err)
	}
	//log.Println(v)

	//slice of obj measurement1 and measurement2
	var slcObj []*jason.Value
	slcObj, err = v.Array()

	if err != nil {
		log.Fatal(err)
	}
	//log.Println(slcObj)

	var measurement, time string
	var obj, fields, tags *jason.Object
	var mapFields map[string]interface{}
	var mapTags map[string]string
	var tagValue string

	//loop through array/slice
	for _, value := range slcObj {
		//log.Println(key)
		//log.Println(value)

		obj, err = value.Object()
		if err != nil {
			log.Println("Error: Expecting object in array of measurement")
		}
		log.Println("OBJECT")
		log.Println(obj)

		measurement, err = obj.GetString("measurement")
		if err != nil {
			log.Println("Error: Expecting measurement key name in JSON")
		}

		if measurement == "" {
			log.Println("Error: Blank value of measurement name")
		}
		log.Println(measurement)

		time, err = obj.GetString("time")
		//time was not specified so assume time.Now() TODO
		if err != nil {
			log.Println("Time not specified. Assumes time.Now")
		} else {
			log.Println("TIME")
			log.Println(time)
		}

		fields, err = obj.GetObject("fields")
		if err != nil {
			log.Println("Error: Parsing fields. Must have at least one field")
		} else {
			log.Println("FIELDS")
			log.Println(fields)

			mapFields = make(map[string]interface{})
			for key, value := range fields.Map() {
				//log.Println(key)
				//log.Println(value.Interface())
				mapFields[key] = value.Interface()
			}
			log.Println(mapFields)
		}

		tags, err = obj.GetObject("tags")
		if err != nil {
			//tags are optional so it's to be nil
		} else {
			log.Println("TAGS")
			log.Println(tags)

			mapTags = make(map[string]string)
			for key, value := range tags.Map() {
				//log.Println(key)
				//log.Println(value.Interface())
				tagValue, err = value.String()
				if err != nil {
					log.Println("Error parsing tag value")
				} else {
					mapTags[key] = tagValue
				}
			}
			log.Println(mapTags)
		}
	}
}
