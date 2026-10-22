# ffl

> **Archived.** This was an experimental side project from the 2018 NFL season. It never got past the data-gathering stage and isn't maintained. The code is left here for reference only.

## What it was meant to be

The idea was a small Go service that pulled NFL data from several sources into one place, so it could feed a daily fantasy sports (DFS) lineup optimizer. The optimizer itself was never built.

## What actually got written

A handful of loosely connected clients and data models, mostly exercised by printing results to the console from `main.go`:

| Package | What it does |
| --- | --- |
| `fantasyradar` | Client for the Sportradar odds-comparison API (trial tier); fetches a tournament schedule. |
| `mysportsfeed` | Client for the MySportsFeeds v1.2 API: box scores, cumulative player stats, schedules, roster players. |
| `footballoutsiders` | HTML scraper for Football Outsiders' weekly team DVOA ratings tables. |
| `yahoo` | Pulls the player/salary CSV export for a Yahoo DFS contest. |
| `optimizer` | A single struct mapping the columns of an optimizer spreadsheet. No optimizer logic. |
| `model` | Draft domain types (season, week, team, game, player, stat lines), intended for a MySQL store that was never wired up. |
| `tasks` | Loads the `America/New_York` time zone at startup. |

There's also a `Makefile` and multi-stage `Dockerfile` for building a static Linux binary into a `scratch` image (`make run`). The Docker build expects a `zoneinfo.zip` copied into the repo root, which is git-ignored.

## State of things

- Dependencies are vendored with `dep` (`Gopkg.toml` / `Gopkg.lock`), which predates Go modules. Building it today would likely need migrating to modules first.
- API keys were read from `SPORTSRADAR_API_KEY`, `MSF_API_KEY` and `MSF_PASSWORD`. The third-party APIs and scraped pages have almost certainly changed since 2018, so don't expect the clients to work as-is.
- Error handling is uneven (some paths `panic`), and there's very little test coverage.

## License

No license was ever added, so all rights are reserved by default.
