window.BENCHMARK_DATA = {
  "lastUpdate": 1650684665753,
  "repoUrl": "https://github.com/posteris/client-service",
  "entries": {
    "Benchmark": [
      {
        "commit": {
          "author": {
            "email": "gsdenys@gmail.com",
            "name": "Denys G. Santos",
            "username": "gsdenys"
          },
          "committer": {
            "email": "gsdenys@gmail.com",
            "name": "Denys G. Santos",
            "username": "gsdenys"
          },
          "distinct": true,
          "id": "097d7d249e3cc31475009e7c660b293db8313075",
          "message": "linting",
          "timestamp": "2022-04-22T23:45:52-03:00",
          "tree_id": "b59765593ee57a7fbe6a60f886701296d8b4955e",
          "url": "https://github.com/posteris/client-service/commit/097d7d249e3cc31475009e7c660b293db8313075"
        },
        "date": 1650683437960,
        "tool": "go",
        "benches": [
          {
            "name": "BenchmarkInitDatabase/Clickhouse",
            "value": 0.01883,
            "unit": "ns/op",
            "extra": "1000000000 times\n2 procs"
          },
          {
            "name": "BenchmarkInitDatabase/MSSQL",
            "value": 0.01442,
            "unit": "ns/op",
            "extra": "1000000000 times\n2 procs"
          },
          {
            "name": "BenchmarkInitDatabase/MySQL",
            "value": 0.007839,
            "unit": "ns/op",
            "extra": "1000000000 times\n2 procs"
          },
          {
            "name": "BenchmarkInitDatabase/PostgreSQL",
            "value": 0.05442,
            "unit": "ns/op",
            "extra": "1000000000 times\n2 procs"
          },
          {
            "name": "BenchmarkInitDatabase/SQLite",
            "value": 0.001401,
            "unit": "ns/op",
            "extra": "1000000000 times\n2 procs"
          }
        ]
      },
      {
        "commit": {
          "author": {
            "email": "gsdenys@gmail.com",
            "name": "Denys G. Santos",
            "username": "gsdenys"
          },
          "committer": {
            "email": "gsdenys@gmail.com",
            "name": "Denys G. Santos",
            "username": "gsdenys"
          },
          "distinct": true,
          "id": "2135a6676f8e481c22b9dcc4a28784d94993f0da",
          "message": "add basic doc",
          "timestamp": "2022-04-23T00:19:44-03:00",
          "tree_id": "505d2a04ae0dbde3bb32b0b07df98dd5367ae61d",
          "url": "https://github.com/posteris/client-service/commit/2135a6676f8e481c22b9dcc4a28784d94993f0da"
        },
        "date": 1650684665287,
        "tool": "go",
        "benches": [
          {
            "name": "BenchmarkInitDatabase/Clickhouse",
            "value": 0.01269,
            "unit": "ns/op",
            "extra": "1000000000 times\n2 procs"
          },
          {
            "name": "BenchmarkInitDatabase/MSSQL",
            "value": 0.01115,
            "unit": "ns/op",
            "extra": "1000000000 times\n2 procs"
          },
          {
            "name": "BenchmarkInitDatabase/MySQL",
            "value": 0.006802,
            "unit": "ns/op",
            "extra": "1000000000 times\n2 procs"
          },
          {
            "name": "BenchmarkInitDatabase/PostgreSQL",
            "value": 0.05069,
            "unit": "ns/op",
            "extra": "1000000000 times\n2 procs"
          },
          {
            "name": "BenchmarkInitDatabase/SQLite",
            "value": 0.0008383,
            "unit": "ns/op",
            "extra": "1000000000 times\n2 procs"
          }
        ]
      }
    ]
  }
}