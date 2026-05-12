from http.server import HTTPServer, BaseHTTPRequestHandler
from datetime import datetime, date
import math

def calculate_data():
    now = date.today()
    the_time = date(2026, 5, 29)
    next_year = date(2026, 8, 20)
    school_start = date(2025, 8, 20)

    days = math.floor((the_time - now).days) + 2
    days_until_next_year = math.floor((next_year - now).days) + 2
    weekend_days = math.floor(days / 7) * 2
    school_days = round(days - weekend_days)

    total_school_span = (the_time - school_start).days
    elapsed = total_school_span - days
    progress = min(100, max(0, round((elapsed / total_school_span) * 100)))

    return {
        "total_days": max(0, days),
        "school_days": max(0, school_days),
        "days_until_next_year": max(0, days_until_next_year),
        "progress": progress,
    }

def render_html(data):
    return f"""<!DOCTYPE html>
<html lang="en">
<head>
  <meta charset="UTF-8"/>
  <meta name="viewport" content="width=device-width, initial-scale=1.0"/>
  <title>School Countdown</title>
  <link rel="stylesheet" href="https://cdn.jsdelivr.net/npm/@tabler/icons-webfont@latest/tabler-icons.min.css"/>
  <style>
    body {{ font-family: sans-serif; background: #f5f5f0; color: #1a1a1a; display: flex; justify-content: center; padding: 3rem 1rem; margin: 0; }}
    .container {{ width: 100%; max-width: 680px; }}
    .label {{ font-size: 13px; color: #888; text-transform: uppercase; letter-spacing: 0.05em; margin: 0 0 1.5rem 0; }}
    .cards {{ display: grid; grid-template-columns: repeat(auto-fit, minmax(160px, 1fr)); gap: 12px; margin-bottom: 1.5rem; }}
    .card {{ background: #ebebeb; border-radius: 8px; padding: 1rem; }}
    .card-title {{ font-size: 13px; color: #666; margin: 0 0 6px 0; }}
    .card-number {{ font-size: 32px; font-weight: 500; margin: 0; color: #1a1a1a; }}
    .card-sub {{ font-size: 12px; color: #aaa; margin: 4px 0 0 0; }}
    .progress-box {{ background: #ebebeb; border-radius: 8px; padding: 1rem; }}
    .progress-track {{ background: #fff; border-radius: 999px; height: 8px; overflow: hidden; }}
    .progress-fill {{ height: 100%; border-radius: 999px; background: #1D9E75; }}
    .progress-meta {{ display: flex; justify-content: space-between; margin-top: 6px; font-size: 12px; color: #aaa; }}
    @media (prefers-color-scheme: dark) {{
      body {{ background: #1a1a1a; color: #f0f0f0; }}
      .card, .progress-box {{ background: #2a2a2a; }}
      .card-number {{ color: #f0f0f0; }}
    }}
  </style>
</head>
<body>
  <div class="container">
    <p class="label">Countdown tracker</p>
    <div class="cards">
      <div class="card">
        <p class="card-title"><i class="ti ti-calendar"></i> Days until school's out</p>
        <p class="card-number">{data['total_days']}</p>
        <p class="card-sub">May 29, 2026</p>
      </div>
      <div class="card">
        <p class="card-title"><i class="ti ti-school"></i> School days left</p>
        <p class="card-number">{data['school_days']}</p>
        <p class="card-sub">Excluding weekends</p>
      </div>
      <div class="card">
        <p class="card-title"><i class="ti ti-sun"></i> Days until next year</p>
        <p class="card-number">{data['days_until_next_year']}</p>
        <p class="card-sub">Aug 20, 2026</p>
      </div>
    </div>
    <div class="progress-box">
      <p class="card-title">Progress to summer</p>
      <div class="progress-track">
        <div class="progress-fill" style="width: {data['progress']}%;"></div>
      </div>
      <div class="progress-meta">
        <span>Today</span>
        <span>{data['progress']}% done</span>
        <span>May 29</span>
      </div>
    </div>
  </div>
</body>
</html>"""

class Handler(BaseHTTPRequestHandler):
    def do_GET(self):
        data = calculate_data()
        html = render_html(data).encode("utf-8")
        self.send_response(200)
        self.send_header("Content-Type", "text/html; charset=utf-8")
        self.send_header("Content-Length", len(html))
        self.end_headers()
        self.wfile.write(html)

    def log_message(self, format, *args):
        print(f"Request: {args[0]} {args[1]}")

if __name__ == "__main__":
    server = HTTPServer(("localhost", 8080), Handler)
    print("Server running at http://localhost:8080")
    server.serve_forever()
