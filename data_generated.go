// Code generated by go generate; DO NOT EDIT.
package main

import "text/template"

type CountdownEntry struct {
	Text string
	Time int64
}

var PageTemplate = template.Must(template.New("").Parse("<!DOCTYPE html>\n\n<html>\n<head>\n\t<style>\n\t\th1 {\n\t\t\tdisplay: flex;\n\t\t\tjustify-content: center;\n\t\t\talign-items: center;\n\t\t}\n\t\t#countdown { visibility: hidden }\n\t</style>\n</head>\n<body>\n\t<div id=\"current_event\">{{ .Text }}</div>\n\t<h1></h1>\n\t<div id=\"countdown\"><h1>{{ .Time }}</h1></div>\n\t<script type=\"text/javascript\">\n\t\t\"use strict\"\n\t\tconst countdown = document.getElementById(\"countdown\"),\n\t\t\tuntil = countdown.innerHTML\n\n\t\tfunction updateRemaining() {\n\t\t\tlet difference = until - Date.now()\n\n\t\t\tif ( difference <= 0 ) {\n\t\t\t\tif ( difference > -1000 ) {\n\t\t\t\t\tconsole.log(\"Reloading...\")\n\t\t\t\t\tsetTimeout(() => location.reload(true), 1100)\n\t\t\t\t}\n\t\t\t} else {\n\t\t\t\tsetTimeout(updateRemaining, difference % 1000)\n\t\t\t\tconst cdown = countdown.innerHTML = [\n\t\t\t\t\t\"\",\n\t\t\t\t\t[86400000, \"day\"],\n\t\t\t\t\t[3600000, \"hour\"],\n\t\t\t\t\t[60000, \"minute\"],\n\t\t\t\t\t[1000, \"second\"]\n\t\t\t\t].reduce((accum, cur, idx) => {\n\t\t\t\t\tconst val = Math.floor(difference / cur[0])\n\t\t\t\t\tif ( val || accum.length ) {\n\t\t\t\t\t\taccum += `${val} ${cur[1]}${val > 1 ? 's' : ''}${idx < 4 ? ', ' : ''}`\n\t\t\t\t\t\tdifference %= cur[0]\n\t\t\t\t\t}\n\t\t\t\t\treturn accum\n\t\t\t\t})\n\t\t\t\tcountdown.innerHTML = cdown ? `Next surprise in: ${cdown}` : \"Fetching next secret, please wait...\"\n\t\t\t\tcountdown.style.visibility = 'visible'\n\t\t\t}\n\t\t}\n\n\t\tconst now = Date.now()\n\t\tif ( until !== -1 ) {\n\t\t\tupdateRemaining()\n\t\t}\n\t\t\n\t</script>\n</body>\n</html>"))

var CountdownEntries = []CountdownEntry{
	{ "<h1>The first event to display</h1>", 0 },
	{ "<p>An event that will be displayed on August 8th, 2020, at 1:00PM in the time zone where the server is built.</p>", 1596891600 },
	{ "<div>An event that will be displayed on January 16th, 2018, at 10:13:43PM in the GMT-5h time zone (EST)", 1516158823 },
	{ "An event that will be displayed 5 hours and 10 minutes from the time the server is built.", 1516052917 },
}
