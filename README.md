# go-etl
A simple golang program that implements basic ETL capabilities.


## Sources
Sources are a data's origin. Pipelines in go-etl always have a source as the starting point. Once data is in the pipeline, you can manipulate it using transformations like filter, sort, and normalize.

1. Flat-file
    - CSV only

## Targets
1. Flat-file
    -  CSV only
2. Mongo DB Atlas

## Transformations
1. Filter
2. Sorter
3. Normalize
## Data Sources
1. MORTGAGE30US.csv - ref: [30-Year Fixed Rate Mortgage Average in the United States](https://fred.stlouisfed.org/series/MORTGAGE30US)
2. DRFFRMACBS.CSV - ref: [Delinquency Rate on Single-Family Residential Mortgages, Booked in Domestic Offices, All Commercial Banks](https://fred.stlouisfed.org/series/DRSFRMACBS)