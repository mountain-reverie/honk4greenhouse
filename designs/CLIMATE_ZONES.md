# Climate Zones for Greenhouse Design

This document provides comprehensive technical guidance on USDA and Canadian plant hardiness zone systems, including data sources, APIs, and lookup methods for determining zones from postal/ZIP codes. Understanding climate zones is essential for greenhouse design decisions including glazing selection, insulation requirements, heating system sizing, and crop selection.

## Table of Contents

1. [Overview](#overview)
2. [USDA Plant Hardiness Zone System](#usda-plant-hardiness-zone-system)
3. [Canadian Plant Hardiness Zone System](#canadian-plant-hardiness-zone-system)
4. [Comparison of USDA and Canadian Systems](#comparison-of-usda-and-canadian-systems)
5. [Zone Lookup by ZIP/Postal Code](#zone-lookup-by-zippostal-code)
6. [GIS Data and Map Downloads](#gis-data-and-map-downloads)
7. [Building Climate Zones (HDD)](#building-climate-zones-hdd)
8. [Data Model](#data-model)
9. [Implementation Strategy](#implementation-strategy)
10. [References](#references)

---

## Overview

Plant hardiness zones provide standardized classification systems for determining which plants can survive in specific geographic locations. For greenhouse design, these zones inform:

- **Heating system sizing** - Colder zones require higher capacity heating
- **Insulation requirements** - More extreme minimum temperatures need better insulation
- **Glazing selection** - Balance between light transmission and thermal resistance
- **Crop selection** - Matching crops to achievable growing conditions
- **Season extension potential** - Understanding baseline climate constraints

### Key Differences Between Systems

| Aspect | USDA System | Canadian System |
|--------|-------------|-----------------|
| **Range** | Zones 1a to 13b | Zones 0a to 9a |
| **Basis** | Single variable (minimum temperature) | 7 climate variables |
| **Variables** | Average annual extreme minimum temperature | Min/max temp, rainfall, snow, wind, growing season |
| **Interchangeable** | No | No |
| **Conversion** | ~1 zone difference (USDA 4 ≈ Canada 5) | Location-specific |

---

## USDA Plant Hardiness Zone System

The USDA Plant Hardiness Zone Map is the standard for determining which perennial plants can survive at a location. The 2023 map (current version) is based on 1991-2020 climate data from 13,412 weather stations.

### How USDA Zones Work

- Based solely on **average annual extreme minimum winter temperature**
- Divided into **10°F (5.6°C) zones**
- Further divided into **5°F (2.8°C) half-zones** (a and b)
- Zone 1 is coldest, Zone 13 is warmest
- "a" subzones are colder than "b" subzones within same zone

### Complete USDA Zone Temperature Table

| Zone | Temperature (°F) | Temperature (°C) |
|------|------------------|------------------|
| **1a** | -60 to -55 | -51.1 to -48.3 |
| **1b** | -55 to -50 | -48.3 to -45.6 |
| **2a** | -50 to -45 | -45.6 to -42.8 |
| **2b** | -45 to -40 | -42.8 to -40.0 |
| **3a** | -40 to -35 | -40.0 to -37.2 |
| **3b** | -35 to -30 | -37.2 to -34.4 |
| **4a** | -30 to -25 | -34.4 to -31.7 |
| **4b** | -25 to -20 | -31.7 to -28.9 |
| **5a** | -20 to -15 | -28.9 to -26.1 |
| **5b** | -15 to -10 | -26.1 to -23.3 |
| **6a** | -10 to -5 | -23.3 to -20.6 |
| **6b** | -5 to 0 | -20.6 to -17.8 |
| **7a** | 0 to 5 | -17.8 to -15.0 |
| **7b** | 5 to 10 | -15.0 to -12.2 |
| **8a** | 10 to 15 | -12.2 to -9.4 |
| **8b** | 15 to 20 | -9.4 to -6.7 |
| **9a** | 20 to 25 | -6.7 to -3.9 |
| **9b** | 25 to 30 | -3.9 to -1.1 |
| **10a** | 30 to 35 | -1.1 to 1.7 |
| **10b** | 35 to 40 | 1.7 to 4.4 |
| **11a** | 40 to 45 | 4.4 to 7.2 |
| **11b** | 45 to 50 | 7.2 to 10.0 |
| **12a** | 50 to 55 | 10.0 to 12.8 |
| **12b** | 55 to 60 | 12.8 to 15.6 |
| **13a** | 60 to 65 | 15.6 to 18.3 |
| **13b** | 65 to 70 | 18.3 to 21.1 |

### Geographic Distribution (USA)

| Region | Typical Zones |
|--------|---------------|
| Alaska (interior) | 1-3 |
| Northern Great Plains | 3-4 |
| Upper Midwest | 3-5 |
| Northeast | 4-7 |
| Pacific Northwest | 6-9 |
| Southeast | 7-9 |
| Gulf Coast | 8-10 |
| Southern California | 9-11 |
| Hawaii | 10-13 |
| Puerto Rico | 11-13 |

### 2023 Map Updates

The 2023 USDA map shows approximately half of the country shifted to the next warmer half-zone compared to the 2012 map. Key factors:

- Based on 1991-2020 data (vs 1976-2005 for 2012 map)
- Uses 13,412 weather stations (vs 7,983 in 2012)
- Improved interpolation algorithms
- Better coverage in data-sparse areas

---

## Canadian Plant Hardiness Zone System

Canada uses a multivariate system developed in the 1960s that considers multiple climate factors beyond just minimum temperature. The system was updated in 2025 using 1991-2020 climate data.

### How Canadian Zones Work

The Canadian system calculates a **Plant Hardiness Index** (PHI) using 7 climate variables:

| Variable | Symbol | Description |
|----------|--------|-------------|
| **X₁** | Min Temperature | Monthly mean of daily minimum temperatures of the coldest month (°C) |
| **X₂** | Growing Season | Length of frost-free period (days) |
| **X₃** | Rainfall Index | Rainfall from June-November: R/(R+25.4) where R is in mm |
| **X₄** | Max Temperature | Monthly mean of daily maximum temperatures of the warmest month (°C) |
| **X₅** | Winter Factor | (0 - X₁) × January rainfall (mm) |
| **X₆** | Snow Depth | Mean maximum snow depth: S/(S+25.4) where S is in mm |
| **X₇** | Wind Speed | Maximum hourly wind gust (km/h) |

### The Ouellet-Sherk Formula

```
PHI = -67.62 + 1.734(X₁) + 0.1868(X₂) + 69.77(X₃) + 1.256(X₄) + 0.006119(X₅) + 22.37(X₆) - 0.01832(X₇)
```

The resulting index value ranges from 0 to 100+, where higher values indicate milder (more favorable) growing conditions.

### Index Value to Zone Mapping

| Index Range | Zone | Index Range | Zone |
|-------------|------|-------------|------|
| < 0 | 0a | 45-49 | 4b |
| 0-4 | 0a | 50-54 | 5a |
| 5-9 | 0b | 55-59 | 5b |
| 10-14 | 1a | 60-64 | 6a |
| 15-19 | 1b | 65-69 | 6b |
| 20-24 | 2a | 70-74 | 7a |
| 25-29 | 2b | 75-79 | 7b |
| 30-34 | 3a | 80-84 | 8a |
| 35-39 | 3b | 85-89 | 8b |
| 40-44 | 4a | 90+ | 9a |

### Canadian Zone Coverage

| Zone | % of Canada | Primary Regions |
|------|-------------|-----------------|
| 0a | 50.3% | Nunavut, northern territories |
| 0b-2a | ~30% | Yukon, NWT, northern provinces |
| 2b-3b | ~10% | Northern Prairies, northern Ontario/Quebec |
| 4a-5b | ~7% | Southern Prairies, central Ontario/Quebec |
| 6a-7b | ~2.5% | Southern Ontario, Lower Mainland BC |
| 8a-9a | <0.2% | Vancouver Island, SW BC coast |

### Canadian Zone Temperature Equivalents (Approximate)

While the Canadian system uses multiple variables, approximate minimum temperature ranges are:

| Canadian Zone | Approximate Min Temp (°C) |
|---------------|---------------------------|
| 0a | Below -46 |
| 0b | -46 to -43 |
| 1a | -43 to -40 |
| 1b | -40 to -37 |
| 2a | -37 to -34 |
| 2b | -34 to -32 |
| 3a | -32 to -29 |
| 3b | -29 to -26 |
| 4a | -26 to -23 |
| 4b | -23 to -21 |
| 5a | -21 to -18 |
| 5b | -18 to -15 |
| 6a | -15 to -12 |
| 6b | -12 to -9 |
| 7a | -9 to -7 |
| 7b | -7 to -4 |
| 8a | -4 to -1 |
| 8b | -1 to 2 |
| 9a | 2 to 4 |

---

## Comparison of USDA and Canadian Systems

### Why They Cannot Be Used Interchangeably

1. **Different Variables**: USDA uses only minimum temperature; Canada uses 7 variables
2. **Different Scales**: USDA has 13 zones; Canada has 10 zones
3. **Different Methodologies**: Temperature bands vs. regression model
4. **Different Outcomes**: Two locations with the same USDA zone may have different Canadian zones

### Approximate Conversion

As a general rule of thumb, add one zone when converting from USDA to Canadian:

| USDA Zone | Approximate Canadian Zone |
|-----------|---------------------------|
| 3 | 4 |
| 4 | 5 |
| 5 | 6 |
| 6 | 7 |
| 7 | 8 |

**Important**: This is only an approximation. For accurate results, use the official Canadian planthardiness.gc.ca tool which provides both Canadian and USDA-equivalent zones for specific locations.

### Practical Implications for Greenhouse Design

For greenhouse design in Canada, we recommend:

1. **Primary Reference**: Use Canadian zones for Canadian locations
2. **USDA Reference**: Use for cross-referencing plant/equipment from US sources
3. **Design Temperature**: Use actual minimum temperatures rather than zone numbers
4. **Multiple Factors**: Consider all 7 Canadian variables for comprehensive design

---

## Zone Lookup by ZIP/Postal Code

### United States - ZIP Code Lookup

#### Option 1: Frostline API (Recommended)

Free, open-source static API providing USDA zones by ZIP code.

**API Endpoint**: `https://phzmapi.org/{ZIPCODE}.json`

**Example Request**:
```
GET https://phzmapi.org/20001.json
```

**Example Response**:
```json
{
  "zone": "8a",
  "temperature_range": "10 to 15",
  "coordinates": {
    "lat": "38.907711",
    "lon": "-77.01732"
  }
}
```

**Notes**:
- Temperature range is in Fahrenheit
- Not all ZIP codes are included (depends on PRISM source data)
- Source: https://github.com/waldoj/frostline

#### Option 2: PRISM ZIP Code CSV

Direct download of all US ZIP codes with hardiness zones.

**Download URL**: `https://prism.oregonstate.edu/phzm/data/2023/phzm_us_zipcode_2023.csv`

**Format**: CSV with columns for ZIP code, zone, and coordinates

#### Option 3: Official USDA Interactive Map

**URL**: https://planthardiness.ars.usda.gov/

Enter ZIP code in the search box for instant lookup.

### Canada - Postal Code Lookup

Canada does not provide a direct postal code API. Available options:

#### Option 1: Municipality Lookup (Official)

**URL**: https://planthardiness.gc.ca/

- Search by municipality/city name
- Returns Canadian zone, USDA equivalent, and climate data
- Most accurate method

#### Option 2: Build Custom Lookup

Combine postal code coordinates with zone raster data:

1. **Postal Code Coordinates**: Use one of these datasets:
   - [Geocoder.ca](http://www.geocoder.ca/?freedata=1) - Free crowdsourced (~925,000 postal codes)
   - [Service Objects](https://www.serviceobjects.com/blog/free-zip-code-and-postal-code-database-with-geocoordinates/) - Free download
   - [GitHub geoinfo-dataset](https://github.com/djbelieny/geoinfo-dataset) - Free CSV
   - Statistics Canada PCCF - Licensed product

2. **Zone Raster Data**: Download from Open Government Portal
3. **Spatial Query**: Point-in-polygon or raster sampling to determine zone

#### Option 3: PlantMaps Interactive

**URL**: https://www.plantmaps.com/interactive-canada-hardiness-gardening-zone-map.php

- Interactive map with click-to-query
- Provides both Canadian and USDA zones
- Not available as API

---

## GIS Data and Map Downloads

### USDA/PRISM Data (United States)

**Source**: PRISM Climate Group, Oregon State University
**URL**: https://prism.oregonstate.edu/phzm/

| Region | Grid Data | Shapefile | KML | ZIP Code CSV |
|--------|-----------|-----------|-----|--------------|
| **CONUS** | [phzm_us_grid_2023.zip](https://prism.oregonstate.edu/phzm/data/2023/phzm_us_grid_2023.zip) | [phzm_us_zones_shp_2023.zip](https://prism.oregonstate.edu/phzm/data/2023/phzm_us_zones_shp_2023.zip) | [phzm_us_zones_kml_2023.zip](https://prism.oregonstate.edu/phzm/data/2023/phzm_us_zones_kml_2023.zip) | [phzm_us_zipcode_2023.csv](https://prism.oregonstate.edu/phzm/data/2023/phzm_us_zipcode_2023.csv) |
| **Alaska** | [phzm_ak_grid_2023.zip](https://prism.oregonstate.edu/phzm/data/2023/phzm_ak_grid_2023.zip) | [phzm_ak_zones_shp_2023.zip](https://prism.oregonstate.edu/phzm/data/2023/phzm_ak_zones_shp_2023.zip) | [phzm_ak_zones_kml_2023.zip](https://prism.oregonstate.edu/phzm/data/2023/phzm_ak_zones_kml_2023.zip) | [phzm_ak_zipcode_2023.csv](https://prism.oregonstate.edu/phzm/data/2023/phzm_ak_zipcode_2023.csv) |
| **Hawaii** | [phzm_hi_grid_2023.zip](https://prism.oregonstate.edu/phzm/data/2023/phzm_hi_grid_2023.zip) | [phzm_hi_zones_shp_2023.zip](https://prism.oregonstate.edu/phzm/data/2023/phzm_hi_zones_shp_2023.zip) | [phzm_hi_zones_kml_2023.zip](https://prism.oregonstate.edu/phzm/data/2023/phzm_hi_zones_kml_2023.zip) | [phzm_hi_zipcode_2023.csv](https://prism.oregonstate.edu/phzm/data/2023/phzm_hi_zipcode_2023.csv) |
| **Puerto Rico** | [phzm_pr_grid_2023.zip](https://prism.oregonstate.edu/phzm/data/2023/phzm_pr_grid_2023.zip) | [phzm_pr_zones_shp_2023.zip](https://prism.oregonstate.edu/phzm/data/2023/phzm_pr_zones_shp_2023.zip) | [phzm_pr_zones_kml_2023.zip](https://prism.oregonstate.edu/phzm/data/2023/phzm_pr_zones_kml_2023.zip) | [phzm_pr_zipcode_2023.csv](https://prism.oregonstate.edu/phzm/data/2023/phzm_pr_zipcode_2023.csv) |

**Resolution**:
- CONUS & Alaska: 800m (30 arc-seconds)
- Hawaii & Puerto Rico: 400m (15 arc-seconds)

**License**: Free with attribution (USDA-ARS and OSU logos required on derived maps)

### Canadian Data

**Source**: Natural Resources Canada
**URL**: https://open.canada.ca/data/en/dataset/db9b4130-8893-11e0-9b96-6cf049291510

| Format | English | French |
|--------|---------|--------|
| **JP2 (JPEG 2000)** | [6400_plant_hardiness_zones.jp2](https://ftp.geogratis.gc.ca/pub/nrcan_rncan/raster/atlas_6_ed/eng/6400_plant_hardiness_zones.jp2) | [6400_zones_de_rusticite_des_plantes.jp2](https://ftp.geogratis.gc.ca/pub/nrcan_rncan/raster/atlas_6_ed/fra/6400_zones_de_rusticite_des_plantes.jp2) |
| **ZIP (PDF, JPG)** | [6400_plant_hardiness_zones.zip](https://ftp.geogratis.gc.ca/pub/nrcan_rncan/raster/atlas_6_ed/eng/6400_plant_hardiness_zones.zip) | [6400_zones_de_rusticite_des_plantes.zip](https://ftp.geogratis.gc.ca/pub/nrcan_rncan/raster/atlas_6_ed/fra/6400_zones_de_rusticite_des_plantes.zip) |

**Resolution**: ~2 km
**License**: Open Government Licence - Canada

### ArcGIS Resources

**USDA ArcGIS Hub**: https://www.usda-plant-hardiness-zone-map-usdaars.hub.arcgis.com/
- Feature layers for web mapping
- REST API endpoints
- Download in CSV, KML, GeoJSON, GeoTIFF, PNG

---

## Building Climate Zones (HDD)

Building climate zones based on Heating Degree Days (HDD) are different from plant hardiness zones but relevant for greenhouse construction and energy calculations.

### Heating Degree Days Explained

HDD measures how much heating is needed:

```
Daily HDD = max(0, 18°C - daily mean temperature)
Annual HDD = sum of daily HDD values for the year
```

### Canadian Building Climate Zones (NECB)

| Zone | HDD Range | Example Cities |
|------|-----------|----------------|
| 4 | < 3000 | Vancouver (2925) |
| 5 | 3000-3999 | Toronto (3520), Victoria (3040) |
| 6 | 4000-4999 | Ottawa (4440), Montreal (4270) |
| 7A | 5000-5999 | Edmonton (5120), Winnipeg (5670) |
| 7B | 6000-6999 | Calgary (5000), Saskatoon (5850) |
| 8 | ≥ 7000 | Yellowknife (8170), Whitehorse (6580) |

### Relationship to Hardiness Zones

| HDD Range | Typical Canadian Hardiness Zone | Typical USDA Zone |
|-----------|--------------------------------|-------------------|
| < 3000 | 8-9 | 8-9 |
| 3000-4000 | 6-7 | 6-7 |
| 4000-5000 | 5-6 | 5-6 |
| 5000-6000 | 4-5 | 4-5 |
| 6000-7000 | 3-4 | 3-4 |
| > 7000 | 0-3 | 1-3 |

### HDD Data Sources

- [Environment and Climate Change Canada - HDD Data](https://open.canada.ca/data/en/dataset/fd8efb83-b73d-5442-ab60-7987c824f5fd)
- [ClimateData.ca - Building Climate Zones](https://climatedata.ca/resource/projected-building-climate-zones/)

---

## Data Model

### Go Structs for Climate Zone Data

```go
package climate

// HardinessZoneSystem represents the classification system used
type HardinessZoneSystem string

const (
    ZoneSystemUSDA     HardinessZoneSystem = "usda"
    ZoneSystemCanadian HardinessZoneSystem = "canadian"
)

// HardinessZone represents a plant hardiness zone
type HardinessZone struct {
    System      HardinessZoneSystem `json:"system"`
    Zone        string              `json:"zone"`         // e.g., "5a", "5b"
    ZoneNumber  int                 `json:"zone_number"`  // e.g., 5
    SubZone     string              `json:"sub_zone"`     // "a" or "b"

    // Temperature ranges (for USDA, these are exact; for Canadian, approximate)
    MinTempC    float64             `json:"min_temp_c"`
    MaxTempC    float64             `json:"max_temp_c"`
    MinTempF    float64             `json:"min_temp_f"`
    MaxTempF    float64             `json:"max_temp_f"`
}

// CanadianHardinessIndex represents the full Canadian climate data
type CanadianHardinessIndex struct {
    // The 7 input variables
    X1_MinTemp        float64 `json:"x1_min_temp"`         // Monthly mean min temp, coldest month (°C)
    X2_GrowingSeason  float64 `json:"x2_growing_season"`   // Frost-free period (days)
    X3_RainfallIndex  float64 `json:"x3_rainfall_index"`   // June-Nov rainfall index
    X4_MaxTemp        float64 `json:"x4_max_temp"`         // Monthly mean max temp, warmest month (°C)
    X5_WinterFactor   float64 `json:"x5_winter_factor"`    // (0 - X1) × Jan rainfall
    X6_SnowDepth      float64 `json:"x6_snow_depth"`       // Snow depth index
    X7_WindGust       float64 `json:"x7_wind_gust"`        // Max hourly wind gust (km/h)

    // Calculated values
    IndexValue        float64 `json:"index_value"`         // PHI value (0-100+)
    CanadianZone      string  `json:"canadian_zone"`       // e.g., "5a"
    USDAEquivalent    string  `json:"usda_equivalent"`     // USDA zone for comparison
}

// CalculateCanadianIndex calculates the Plant Hardiness Index from variables
func CalculateCanadianIndex(x1, x2, x3, x4, x5, x6, x7 float64) float64 {
    return -67.62 +
           1.734*x1 +
           0.1868*x2 +
           69.77*x3 +
           1.256*x4 +
           0.006119*x5 +
           22.37*x6 -
           0.01832*x7
}

// IndexToCanadianZone converts a PHI value to a Canadian zone
func IndexToCanadianZone(index float64) string {
    switch {
    case index < 5:
        return "0a"
    case index < 10:
        return "0b"
    case index < 15:
        return "1a"
    case index < 20:
        return "1b"
    case index < 25:
        return "2a"
    case index < 30:
        return "2b"
    case index < 35:
        return "3a"
    case index < 40:
        return "3b"
    case index < 45:
        return "4a"
    case index < 50:
        return "4b"
    case index < 55:
        return "5a"
    case index < 60:
        return "5b"
    case index < 65:
        return "6a"
    case index < 70:
        return "6b"
    case index < 75:
        return "7a"
    case index < 80:
        return "7b"
    case index < 85:
        return "8a"
    case index < 90:
        return "8b"
    default:
        return "9a"
    }
}

// LocationClimate represents climate data for a specific location
type LocationClimate struct {
    // Location identifiers
    Country     string  `json:"country"`      // "CA" or "US"
    PostalCode  string  `json:"postal_code"`  // ZIP code or Canadian postal code
    City        string  `json:"city"`
    Province    string  `json:"province"`     // Province/State
    Latitude    float64 `json:"latitude"`
    Longitude   float64 `json:"longitude"`

    // Hardiness zones
    USDAZone         *HardinessZone          `json:"usda_zone"`
    CanadianZone     *HardinessZone          `json:"canadian_zone,omitempty"`
    CanadianIndex    *CanadianHardinessIndex `json:"canadian_index,omitempty"`

    // Building climate data
    AnnualHDD        float64 `json:"annual_hdd"`        // Heating degree days (base 18°C)
    BuildingZone     string  `json:"building_zone"`     // NECB zone (4-8)
    FrostDepthM      float64 `json:"frost_depth_m"`     // Design frost depth

    // Design temperatures
    DesignTempWinterC float64 `json:"design_temp_winter_c"` // 2.5% design temp
    DesignTempSummerC float64 `json:"design_temp_summer_c"` // 1% design temp

    // Climate data source
    DataSource       string  `json:"data_source"`
    DataYear         int     `json:"data_year"`
}

// USDAZoneTemperature returns temperature range for a USDA zone
type USDAZoneTemperature struct {
    Zone     string  `json:"zone"`
    MinTempF float64 `json:"min_temp_f"`
    MaxTempF float64 `json:"max_temp_f"`
    MinTempC float64 `json:"min_temp_c"`
    MaxTempC float64 `json:"max_temp_c"`
}

// USDAZoneTable contains all USDA zone temperature data
var USDAZoneTable = []USDAZoneTemperature{
    {"1a", -60, -55, -51.1, -48.3},
    {"1b", -55, -50, -48.3, -45.6},
    {"2a", -50, -45, -45.6, -42.8},
    {"2b", -45, -40, -42.8, -40.0},
    {"3a", -40, -35, -40.0, -37.2},
    {"3b", -35, -30, -37.2, -34.4},
    {"4a", -30, -25, -34.4, -31.7},
    {"4b", -25, -20, -31.7, -28.9},
    {"5a", -20, -15, -28.9, -26.1},
    {"5b", -15, -10, -26.1, -23.3},
    {"6a", -10, -5, -23.3, -20.6},
    {"6b", -5, 0, -20.6, -17.8},
    {"7a", 0, 5, -17.8, -15.0},
    {"7b", 5, 10, -15.0, -12.2},
    {"8a", 10, 15, -12.2, -9.4},
    {"8b", 15, 20, -9.4, -6.7},
    {"9a", 20, 25, -6.7, -3.9},
    {"9b", 25, 30, -3.9, -1.1},
    {"10a", 30, 35, -1.1, 1.7},
    {"10b", 35, 40, 1.7, 4.4},
    {"11a", 40, 45, 4.4, 7.2},
    {"11b", 45, 50, 7.2, 10.0},
    {"12a", 50, 55, 10.0, 12.8},
    {"12b", 55, 60, 12.8, 15.6},
    {"13a", 60, 65, 15.6, 18.3},
    {"13b", 65, 70, 18.3, 21.1},
}
```

### Database Schema

```sql
-- Climate zone lookup table (pre-populated)
CREATE TABLE usda_zones_by_zip (
    zip_code TEXT PRIMARY KEY,
    zone TEXT NOT NULL,
    zone_number INTEGER NOT NULL,
    sub_zone TEXT NOT NULL,
    min_temp_f REAL NOT NULL,
    max_temp_f REAL NOT NULL,
    min_temp_c REAL NOT NULL,
    max_temp_c REAL NOT NULL,
    latitude REAL,
    longitude REAL
);

-- Canadian postal code to coordinates mapping
CREATE TABLE canadian_postal_codes (
    postal_code TEXT PRIMARY KEY,
    city TEXT,
    province TEXT,
    latitude REAL NOT NULL,
    longitude REAL NOT NULL,
    accuracy INTEGER  -- 1=estimated, 4=geonameid, 6=centroid
);

-- Cached climate lookups
CREATE TABLE location_climate_cache (
    id TEXT PRIMARY KEY,
    country TEXT NOT NULL,
    postal_code TEXT NOT NULL,
    latitude REAL NOT NULL,
    longitude REAL NOT NULL,

    usda_zone TEXT,
    canadian_zone TEXT,
    canadian_index_value REAL,

    annual_hdd REAL,
    building_zone TEXT,
    frost_depth_m REAL,

    design_temp_winter_c REAL,
    design_temp_summer_c REAL,

    data_source TEXT,
    cached_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,

    UNIQUE(country, postal_code)
);

-- Index for fast lookups
CREATE INDEX idx_climate_cache_postal ON location_climate_cache(country, postal_code);
```

---

## Implementation Strategy

### Recommended Approach for Honk4Greenhouse

#### 1. US Zone Lookup (Simple)

```go
// Use the Frostline API for US ZIP codes
func GetUSDAZoneByZIP(zipCode string) (*HardinessZone, error) {
    resp, err := http.Get(fmt.Sprintf("https://phzmapi.org/%s.json", zipCode))
    if err != nil {
        return nil, err
    }
    defer resp.Body.Close()

    var result struct {
        Zone             string `json:"zone"`
        TemperatureRange string `json:"temperature_range"`
        Coordinates      struct {
            Lat string `json:"lat"`
            Lon string `json:"lon"`
        } `json:"coordinates"`
    }

    if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
        return nil, err
    }

    return parseZone(result.Zone, "usda"), nil
}
```

#### 2. Canadian Zone Lookup (More Complex)

Option A: **Pre-load ZIP code CSV from PRISM + Canadian postal code coordinates**

```go
// Load postal code coordinates and use raster lookup
func GetCanadianZoneByPostalCode(postalCode string) (*LocationClimate, error) {
    // 1. Look up postal code coordinates from database
    coords, err := db.GetPostalCodeCoords(postalCode)
    if err != nil {
        return nil, err
    }

    // 2. Query the planthardiness.gc.ca site (if API available)
    //    OR use pre-downloaded raster data for point query

    // 3. Return combined climate data
    return &LocationClimate{
        PostalCode:   postalCode,
        Latitude:     coords.Lat,
        Longitude:    coords.Lon,
        CanadianZone: zone,
        // ... additional data
    }, nil
}
```

Option B: **Cache municipality data from planthardiness.gc.ca**

Pre-populate a database with zone data for Canadian municipalities, then map postal codes to nearest municipality.

#### 3. Fallback Strategy

For locations where API lookup fails:
1. Use coordinates to determine approximate zone from downloaded shapefile/raster
2. Allow manual override by user
3. Default to conservative estimate based on province

### UI Integration

The wizard (Step 2 - Dimensions & Crop Profile or a new location step) should:

1. **Auto-detect zone** when user enters location
2. **Display zone badge** showing both Canadian and USDA zones
3. **Show design temperature** for heating calculations
4. **Allow override** for microclimate considerations

```go
templ ClimateZoneBadge(location LocationClimate) {
    <div class="flex items-center gap-3 p-3 rounded-lg bg-muted">
        @icon.Icon(name="thermometer", size="md")
        <div>
            <p class="font-medium">
                if location.Country == "CA" {
                    { T(ctx, "climate.canadian_zone") }: { location.CanadianZone.Zone }
                    <span class="text-muted-foreground text-sm ml-2">
                        (USDA { location.USDAZone.Zone })
                    </span>
                } else {
                    { T(ctx, "climate.usda_zone") }: { location.USDAZone.Zone }
                }
            </p>
            <p class="text-sm text-muted-foreground">
                { T(ctx, "climate.design_temp") }:
                { fmt.Sprintf("%.1f°C", location.DesignTempWinterC) }
            </p>
        </div>
    </div>
}
```

---

## References

### Official Sources

1. **USDA Plant Hardiness Zone Map** - https://planthardiness.ars.usda.gov/
2. **PRISM Climate Group** - https://prism.oregonstate.edu/phzm/
3. **Canada's Plant Hardiness Site** - https://planthardiness.gc.ca/
4. **Open Government Portal (Canada)** - https://open.canada.ca/data/en/dataset/db9b4130-8893-11e0-9b96-6cf049291510

### API and Data Sources

5. **Frostline API (US ZIP codes)** - https://phzmapi.org/ and https://github.com/waldoj/frostline
6. **USDA ArcGIS Hub** - https://www.usda-plant-hardiness-zone-map-usdaars.hub.arcgis.com/
7. **Geocoder.ca (Canadian postal codes)** - http://www.geocoder.ca/?freedata=1
8. **Service Objects Free Database** - https://www.serviceobjects.com/blog/free-zip-code-and-postal-code-database-with-geocoordinates/

### Scientific References

9. **Updated plant hardiness zones for Canada and assessment of change over time** - https://www.nature.com/articles/s41598-025-00931-5
10. **Canada's Plant Hardiness zones revisited using modern climate interpolation techniques** - https://cdnsciencepub.com/doi/10.4141/P00-030
11. **Ouellet, C.E. and Sherk, L.C. (1967)** - Original Canadian hardiness index development

### Climate Data

12. **Environment and Climate Change Canada** - https://climate.weather.gc.ca/
13. **ClimateData.ca** - https://climatedata.ca/
14. **Canada HDD Data** - https://open.canada.ca/data/en/dataset/fd8efb83-b73d-5442-ab60-7987c824f5fd

### Comparison and Conversion

15. **Wikipedia - Hardiness Zone** - https://en.wikipedia.org/wiki/Hardiness_zone
16. **Garden Betty - Canadian Hardiness Zones Conversion** - https://gardenbetty.com/canadian-hardiness-zones/
17. **Earth Undaunted - Differences Between US and Canadian Zones** - https://earthundaunted.com/the-differences-between-us-and-canadian-plant-hardiness-zones/
