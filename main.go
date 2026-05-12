package main

import (
	"fmt"
	"html/template"
	"math"
	"net/http"
	"time"
)

type CountdownData struct {
	TotalDays       float64
	SchoolDays      float64
	DaysUntilNextYear float64
	Progress        int
}

func calculateData() CountdownData {
	now := time.Now()
	theTime := time.Date(2026, 5, 29, 0, 0, 0, 0, time.Local)
	nextYear := time.Date(2026, 8, 20, 0, 0, 0, 0, time.Local)
	schoolStart := time.Date(2025, 8, 20, 0, 0, 0, 0, time.Local)

	days := math.Floor(time.Until(theTime).Hours()/24) + 2
	daysUntilNextYear := math.Floor(time.Until(nextYear).Hours()/24) + 2
	weekendDays := math.Floor(days/7) * 2
	schoolDays := math.Round(days - weekendDays)

	totalSchoolSpan := theTime.Sub(schoolStart).Hours() / 24
	elapsed := totalSchoolSpan - days
	progress := int(math.Min(100, math.Max(0, math.Round((elapsed/totalSchoolSpan)*100))))

	_ = now
	return CountdownData{
		TotalDays:         math.Max(0, days),
		SchoolDays:        math.Max(0, schoolDays),
		DaysUntilNextYear: math.Max(0, daysUntilNextYear),
		Progress:          progress,
	}
}

var tmpl = template.Must(template.New("page").Parse(`
<!DOCTYPE html>
<html lang="en">
<head>
  <meta charset="UTF-8"/>
  <meta name="viewport" content="width=device-width, initial-scale=1.0"/>
  <title>School Countdown</title>
  <link rel="stylesheet" href="https://cdn.jsdelivr.net/npm/@tabler/icons-webfont@latest/tabler-icons.min.css"/>
  <style>
    body { font-family: sans-serif; background: #f5f5f0; color: #1a1a1a; display: flex; justify-content: center; padding: 3rem 1rem; margin: 0; }
    .container { width: 100%; max-width: 680px; }
    .label { font-size: 13px; color: #888; text-transform: uppercase; letter-spacing: 0.05em; margin: 0 0 1.5rem 0; }
    .cards { display: grid; grid-template-columns: repeat(auto-fit, minmax(160px, 1fr)); gap: 12px; margin-bottom: 1.5rem; }
    .card { background: #ebebeb; border-radius: 8px; padding: 1rem; }
    .card-title { font-size: 13px; color: #666; margin: 0 0 6px 0; }
    .card-number { font-size: 32px; font-weight: 500; margin: 0; color: #1a1a1a; }
    .card-sub { font-size: 12px; color: #aaa; margin: 4px 0 0 0; }
    .progress-box { background: #ebebeb; border-radius: 8px; padding: 1rem; }
    .progress-track { background: #fff; border-radius: 999px; height: 8px; overflow: hidden; }
    .progress-fill { height: 100%; border-radius: 999px; background: #1D9E75; }
    .progress-meta { display: flex; justify-content: space-between; margin-top: 6px; font-size: 12px; color: #aaa; }
    @media (prefers-color-scheme: dark) {
      body { background: #1a1a1a; color: #f0f0f0; }
      .card, .progress-box { background: #2a2a2a; }
      .card-number { color: #f0f0f0; }
    }
  </style>
</head>
<body>
  <div class="container">
    <p class="label">Countdown tracker</p>
    <div class="cards">
      <div class="card">
        <p class="card-title"><i class="ti ti-calendar"></i> Days until school's out</p>
        <p class="card-number">{{.TotalDays}}</p>
        <p class="card-sub">May 29, 2026</p>
      </div>
      <div class="card">
        <p class="card-title"><i class="ti ti-school"></i> School days left</p>
        <p class="card-number">{{.SchoolDays}}</p>
        <p class="card-sub">Excluding weekends</p>
      </div>
      <div class="card">
        <p class="card-title"><i class="ti ti-sun"></i> Days until next year</p>
        <p class="card-number">{{.DaysUntilNextYear}}</p>
        <p class="card-sub">Aug 20, 2026</p>
      </div>
    </div>
    <div class="progress-box">
      <p class="card-title">Progress to summer</p>
      <div class="progress-track">
        <div class="progress-fill" style="width: {{.Progress}}%;"></div>
      </div>
      <div class="progress-meta">
        <span>Today</span>
        <span>{{.Progress}}% done</span>
        <span>May 29</span>
      </div>
    </div>
  </div>
</body>
</html>
`))

func handler(w http.ResponseWriter, r *http.Request) {
	data := calculateData()
	tmpl.Execute(w, data)
}

func main() {
	http.HandleFunc("/", handler)
	fmt.Println("Server running at http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}
