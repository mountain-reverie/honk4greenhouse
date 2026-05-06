# Earth-Sheltered Greenhouse Design Guide

This document provides comprehensive technical guidance for designing earth-sheltered, pit, and walipini-style greenhouses for Canadian climates. It covers the thermal physics of earth-contact construction, calculation methods for geothermal gain, solar geometry for roof design, and practical construction considerations.

## Table of Contents

1. [Overview](#overview)
2. [Types of Earth-Sheltered Greenhouses](#types-of-earth-sheltered-greenhouses)
3. [Thermal Physics of Earth Contact](#thermal-physics-of-earth-contact)
4. [Ground Temperature Profiles](#ground-temperature-profiles)
5. [Heat Transfer Calculations](#heat-transfer-calculations)
6. [Solar Geometry and Roof Design](#solar-geometry-and-roof-design)
7. [Sizing and Depth Parameters](#sizing-and-depth-parameters)
8. [Drainage and Waterproofing](#drainage-and-waterproofing)
9. [Construction Methods](#construction-methods)
10. [Cost Analysis](#cost-analysis)
11. [Calculator Parameters](#calculator-parameters)
12. [Case Studies](#case-studies)
13. [References](#references)

---

## Overview

Earth-sheltered greenhouses leverage the thermal mass and stable temperatures of the ground to dramatically reduce heating requirements. By burying or berming portions of the structure, these designs tap into the "thermal constant" of the earth—the stable temperature found below the frost line.

### Key Benefits

| Benefit | Description | Typical Savings |
|---------|-------------|-----------------|
| **Reduced Temperature Swing** | Earth mass dampens daily/seasonal variations | 50-70% reduction in swing |
| **Lower Heating Load** | Reduced ΔT between inside and outside | 30-50% reduction |
| **Thermal Lag** | Ground temperature lags air by 2-6 months | Warmest earth in winter |
| **Wind Protection** | Below-grade or bermed walls block wind | Eliminates infiltration on buried sides |
| **Thermal Mass** | Soil stores and releases heat | Free heat storage |

### Design Philosophy

The earth-sheltered greenhouse works by:
1. **Reducing exposed surface area** - Less glazing = less heat loss
2. **Increasing effective R-value** - Earth contact provides both insulation and thermal mass
3. **Leveraging stable ground temperatures** - Ground at 2+ m depth is warmer than winter air
4. **Maximizing solar gain** - South-facing glazing captures winter sun
5. **Storing heat in mass** - Soil, concrete, and water store daytime heat for night

---

## Types of Earth-Sheltered Greenhouses

### 1. Walipini (Pit Greenhouse)

**Origin:** Bolivia (Aymara word meaning "place of warmth")

**Description:** Rectangular excavation 1.8-2.5 m (6-8 ft) deep with a sloped transparent roof at or near ground level.

**Characteristics:**
- Floor below frost line taps stable ground temperature
- Low-angle roof suited for tropical/subtropical latitudes
- Simple construction using local materials
- **Limitation in Canada:** Low roof angle causes shading at northern latitudes

```
                    Ground Level
    ════════════════════════════════════════════════
                 ╱ Glazing (low angle) ╲
               ╱                         ╲
    ══════════╱                           ╲══════════
              │                           │
              │      Growing Area         │  1.8-2.5m
              │                           │  deep
              │                           │
    ══════════╧═══════════════════════════╧══════════
              ▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓
                     Earth Floor
```

### 2. Earth-Bermed Greenhouse

**Origin:** Popularized by Mike Oehler (Idaho, USA)

**Description:** Partially underground structure with earth bermed against north, east, and west walls. South face fully glazed with steep angle.

**Characteristics:**
- Steep south-facing glazing (50-70°) for northern latitudes
- Insulated north wall acts as thermal mass
- Earth berms on three sides
- **Best for Canada:** Addresses solar angle issues of traditional walipini

```
                        ╱│
                      ╱  │
                    ╱    │ Glazing
                  ╱      │ (50-70°)
    ▓▓▓▓▓▓▓▓▓▓▓▓╱        │
    ▓▓ Earth  ▓▓│        │
    ▓▓ Berm   ▓▓│ Growing │
    ▓▓▓▓▓▓▓▓▓▓▓▓│  Area   │▓▓▓▓▓▓▓▓▓▓▓▓
    ▓▓▓▓▓▓▓▓▓▓▓▓│        │▓▓ Earth  ▓▓
                │        │▓▓ Berm   ▓▓
    ════════════╧════════╧▓▓▓▓▓▓▓▓▓▓▓▓
                  Floor
```

### 3. Chinese Solar Greenhouse (Lean-To Style)

**Origin:** Northern China (>800,000 in use)

**Description:** South-facing glazed wall with massive insulated north wall, often partially earth-sheltered.

**Characteristics:**
- Thick north wall (0.5-1.0 m) of brick, adobe, or rammed earth
- Removable thermal blanket covers glazing at night
- Often built into south-facing slope
- Proven in climates similar to Canadian prairies

### 4. Uphill Patio Design (Oehler Method)

**Description:** Earth-sheltered structure built into hillside with an "uphill patio" cut into the slope above, allowing light to enter from above and behind.

**Characteristics:**
- Overcomes shading problems of pit greenhouses
- Creates protected outdoor growing area
- Allows light from multiple angles
- Complex excavation but superior light penetration

---

## Thermal Physics of Earth Contact

### Why Earth Contact Reduces Heat Loss

The heat loss through any surface is:
```
Q = U × A × ΔT

Where:
Q = Heat loss (W)
U = Overall heat transfer coefficient (W/m²·K)
A = Surface area (m²)
ΔT = Temperature difference (K)
```

Earth-sheltered construction reduces heat loss through **both** lower U-values AND lower ΔT:

| Surface Type | Typical U-Value (W/m²·K) | Winter ΔT (°C) | Heat Loss Factor |
|--------------|--------------------------|----------------|------------------|
| Single glazing | 6.2 | 35-45 | 1.00 (baseline) |
| Triple polycarbonate | 2.4 | 35-45 | 0.39 |
| Insulated wall (R-20) | 0.28 | 35-45 | 0.05 |
| Earth-contact wall (2m deep) | 0.3-0.5* | 5-15 | 0.01-0.03 |

*Effective U-value including soil thermal resistance

### Thermal Mass vs. Insulation

**Key Insight:** Soil is a poor insulator (R-0.25/inch) but excellent thermal mass.

| Property | EPS Insulation | Soil | Water |
|----------|---------------|------|-------|
| R-value per inch | 3.85 | 0.25 | ~0.24 |
| Density (kg/m³) | 16-32 | 1,400-2,000 | 1,000 |
| Volumetric heat capacity (kJ/m³·K) | 29-58 | 1,300-2,000 | 4,186 |
| Thermal mass benefit | Minimal | Excellent | Best |

The benefit of earth contact is not primarily insulation—it's the **combination of**:
1. Moderate thermal resistance
2. Massive thermal storage
3. Stable temperatures at depth
4. Thermal lag (warmest in winter, coolest in summer)

### Thermal Lag Effect

The ground temperature at depth lags behind surface/air temperature:

| Depth | Lag Time | Amplitude Reduction |
|-------|----------|---------------------|
| 0.5 m | ~3 weeks | 50% |
| 1.0 m | ~6 weeks | 25% |
| 2.0 m | ~3 months | 6% |
| 3.0 m | ~4.5 months | 1.5% |
| 5.0 m | ~6 months | <0.5% |

This means **the earth is warmest when you need it most** (mid-winter) and coolest in summer.

---

## Ground Temperature Profiles

### The Kusuda-Achenbach Equation

This is the standard method for calculating ground temperature at any depth and time of year:

```
T(z,t) = T_mean - T_amp × exp(-z × √(π/(365×α))) × cos(2π/365 × (t - t_shift - z/2 × √(365/(π×α))))

Where:
T(z,t) = Ground temperature at depth z and day t (°C)
T_mean = Mean annual surface temperature (°C)
T_amp = Annual surface temperature amplitude (°C)
z = Depth below surface (m)
t = Day of year (1-365)
t_shift = Day of minimum surface temperature (typically ~35 for Canada)
α = Soil thermal diffusivity (m²/day)
```

### Simplified Form

For most practical calculations:
```
T(z) = T_mean - T_amp × exp(-z/D) × cos(2π(t - t_shift)/365 - z/D)

Where:
D = Damping depth = √(2α/ω) = √(365α/π) meters
ω = Annual frequency = 2π/365 day⁻¹
```

### Damping Depth by Soil Type

The **damping depth** is where temperature amplitude drops to 37% (1/e) of surface amplitude:

| Soil Type | Thermal Diffusivity α (m²/day) | Damping Depth D (m) |
|-----------|--------------------------------|---------------------|
| Dry sand | 0.024 | 1.7 |
| Moist sand | 0.052 | 2.5 |
| Dry clay | 0.018 | 1.5 |
| Moist clay | 0.036 | 2.1 |
| Saturated clay | 0.042 | 2.2 |
| Wet loam | 0.045 | 2.3 |
| Organic/peat | 0.010 | 1.1 |

### Ground Temperatures for Canadian Cities

**Mean Annual Ground Temperature ≈ Mean Annual Air Temperature + 1-2°C**

| City | Mean Air Temp (°C) | Est. Ground Temp at 3m (°C) | Winter Surface Temp (°C) |
|------|-------------------|----------------------------|--------------------------|
| Vancouver | 10.4 | 11-12 | 2-5 |
| Victoria | 10.3 | 11-12 | 3-6 |
| Toronto | 9.4 | 10-11 | -5 to -10 |
| Ottawa | 6.6 | 8-9 | -10 to -15 |
| Montreal | 7.4 | 8-9 | -10 to -15 |
| Calgary | 4.4 | 6-7 | -10 to -20 |
| Edmonton | 2.9 | 5-6 | -15 to -25 |
| Winnipeg | 3.0 | 5-6 | -15 to -25 |
| Saskatoon | 2.8 | 5-6 | -15 to -25 |

### Practical Temperature Differences

At 2m depth in mid-winter (January):

| Location | Outside Air (°C) | Ground at 2m (°C) | ΔT Reduction |
|----------|------------------|-------------------|--------------|
| Toronto | -10 | +8 | 18°C warmer than air |
| Calgary | -20 | +5 | 25°C warmer than air |
| Winnipeg | -25 | +4 | 29°C warmer than air |

**Example benefit:** If greenhouse target is 10°C:
- Above-ground wall in Winnipeg: ΔT = 10 - (-25) = 35°C
- Earth-contact wall at 2m: ΔT = 10 - 4 = 6°C
- **Heat loss reduction: 83%** on that surface

---

## Heat Transfer Calculations

### Heat Loss Through Earth-Contact Walls

The effective U-value for earth-contact walls depends on depth and soil properties:

```
U_eff(z) = 1 / (R_inside + R_wall + R_soil(z) + R_ground)

Where:
R_inside = Inside air film resistance = 0.12 m²·K/W
R_wall = Wall construction resistance (m²·K/W)
R_soil(z) = Soil thermal resistance (varies with depth)
R_ground = Effective ground resistance to stable temperature
```

### Soil Thermal Resistance

For a wall extending from surface to depth z:

```
R_soil = z / (k_soil × shape_factor)

Simplified: R_soil ≈ 0.5 × z / k_soil  (for typical geometry)

Where:
k_soil = Soil thermal conductivity (W/m·K)
```

| Soil Type | k (W/m·K) | R per meter depth (m²·K/W) |
|-----------|-----------|---------------------------|
| Dry sand | 0.3-0.4 | 1.25-1.67 |
| Moist sand | 1.5-2.5 | 0.20-0.33 |
| Dry clay | 0.4-0.6 | 0.83-1.25 |
| Moist clay | 1.0-1.5 | 0.33-0.50 |
| Saturated clay | 1.4-1.8 | 0.28-0.36 |
| Wet loam | 1.2-1.8 | 0.28-0.42 |

### Average Earth-Contact Wall U-Value

For a wall from surface to depth z_max, use average U-value:

| Wall Depth | Uninsulated U_avg | With R-10 Insulation U_avg |
|------------|-------------------|---------------------------|
| 1.0 m | 1.5-2.0 W/m²·K | 0.45-0.55 W/m²·K |
| 1.5 m | 1.0-1.5 W/m²·K | 0.40-0.50 W/m²·K |
| 2.0 m | 0.8-1.2 W/m²·K | 0.35-0.45 W/m²·K |
| 2.5 m | 0.6-1.0 W/m²·K | 0.30-0.40 W/m²·K |

### Perimeter Heat Loss (F-Factor Method)

For floor slabs and shallow earth contact, use perimeter heat loss:

```
Q_floor = F × P × ΔT

Where:
Q_floor = Heat loss through floor/perimeter (W)
F = F-factor (W/m·K)
P = Exposed perimeter length (m)
ΔT = Inside air temp - Outside air temp (K)
```

**F-Factor Values (ASHRAE):**

| Configuration | F (W/m·K) | F (Btu/hr·ft·°F) |
|---------------|-----------|------------------|
| Uninsulated slab | 1.26 | 0.73 |
| R-5 vertical 600mm | 1.06 | 0.62 |
| R-10 vertical 600mm | 0.86 | 0.50 |
| R-15 vertical 600mm | 0.75 | 0.44 |
| R-10 horizontal 600mm | 1.04 | 0.60 |
| R-10 horizontal 1200mm | 0.95 | 0.55 |
| Fully insulated (R-10) | 0.52 | 0.30 |

### Deep Floor Heat Loss

For floors more than 2m below grade with no exposed perimeter:
```
U_floor ≈ 0.1-0.3 W/m²·K (to stable ground temperature)
```

Heat loss is minimal because:
- Large R-value through soil
- Ground temperature close to target
- Very small effective ΔT

---

## Solar Geometry and Roof Design

### The Critical Challenge for Northern Latitudes

Traditional walipini designs fail in Canada because:
1. Low winter sun angle means shallow roofs cause floor shading
2. At 45°N latitude, winter solstice sun is only 21.5° above horizon
3. A pit with 2m walls and flat roof will have 100% floor shading in winter

### Solar Altitude Angle Formula

```
α_noon = 90° - |Latitude - Declination|

Where:
α_noon = Solar altitude at solar noon (degrees)
Latitude = Your latitude (degrees)
Declination = Sun's declination (-23.45° to +23.45°)
```

**Winter Solstice (Dec 21):** Declination = -23.45°

| Latitude | Location | Winter Noon Sun Angle | Summer Noon Sun Angle |
|----------|----------|----------------------|----------------------|
| 43°N | Toronto | 23.5° | 70.5° |
| 45°N | Ottawa/Montreal | 21.5° | 68.5° |
| 49°N | Vancouver | 17.5° | 64.5° |
| 51°N | Calgary | 15.5° | 62.5° |
| 53°N | Edmonton | 13.5° | 60.5° |
| 56°N | Fort McMurray | 10.5° | 57.5° |

### Roof Angle for Maximum Winter Light

**Rule of Thumb:** Roof angle = Latitude + 15° to 25°

**Optimal angle** (perpendicular to winter sun): 90° - α_winter_solstice

| Latitude | Winter Sun Angle | Optimal Roof Angle | Practical Range |
|----------|------------------|-------------------|-----------------|
| 43°N | 23.5° | 66.5° | 55-70° |
| 45°N | 21.5° | 68.5° | 55-70° |
| 49°N | 17.5° | 72.5° | 60-75° |
| 51°N | 15.5° | 74.5° | 60-75° |
| 53°N | 13.5° | 76.5° | 65-80° |

### Preventing Floor Shading

For a pit greenhouse with vertical north wall height H and horizontal depth D:

```
Minimum roof angle = arctan(H/D) + (90° - α_winter)

To prevent ANY winter shading:
D_max = H × tan(α_winter)
```

**Example (45°N, 2m deep pit, sun angle 21.5°):**
```
D_max = 2m × tan(21.5°) = 2m × 0.394 = 0.79m
```

This means a traditional flat-roofed walipini can only be **0.79m wide** to avoid shading in Ottawa!

### Solutions for Northern Latitudes

1. **Steep glazing (55-75°):** Nearly vertical south wall
2. **Earth-bermed design:** Build up north side instead of digging down
3. **Clerestory windows:** Add high windows on north side
4. **Uphill patio:** Cut light wells into north slope
5. **Reflective surfaces:** Use reflective material on north wall to bounce light

### Shadow Length Calculator

```
Shadow length = Object height / tan(solar altitude)

Example: At 45°N on winter solstice (α = 21.5°):
2m tall north wall casts shadow = 2m / tan(21.5°) = 5.1m
```

---

## Sizing and Depth Parameters

### Recommended Depths by Climate Zone

| Climate Zone (HDD) | Min Depth | Optimal Depth | Notes |
|-------------------|-----------|---------------|-------|
| Zone 4 (<3000) | 0.6-1.0 m | 1.0-1.5 m | Shallow OK, mild winters |
| Zone 5 (3000-4000) | 1.0-1.5 m | 1.5-2.0 m | Standard depth |
| Zone 6 (4000-5000) | 1.2-1.8 m | 1.8-2.5 m | Below frost line critical |
| Zone 7 (5000-7000) | 1.5-2.0 m | 2.0-3.0 m | Deep excavation beneficial |
| Zone 8 (>7000) | 2.0-2.5 m | 2.5-3.5 m | Maximum depth for stability |

### Depth vs. Frost Line

**Critical:** Floor must be below maximum frost penetration

| Region | Typical Frost Depth | Minimum Excavation Depth |
|--------|---------------------|-------------------------|
| Coastal BC | 0.3-0.6 m | 1.0 m |
| Southern Ontario | 1.0-1.5 m | 1.8 m |
| Ottawa/Montreal | 1.2-1.8 m | 2.0 m |
| Calgary/Edmonton | 1.5-2.1 m | 2.5 m |
| Winnipeg/Saskatoon | 1.8-2.4 m | 2.7 m |
| Northern regions | 2.0-3.0 m+ | 3.0 m+ |

### Sizing Rules of Thumb

**Floor Area:**
- Minimum practical size: 3m × 4m = 12 m²
- Typical hobby size: 4m × 6m = 24 m²
- Production size: 6m × 12m = 72 m²

**Proportions (for earth-bermed style):**
- Width (N-S): 3-6 m (limited by shading)
- Length (E-W): 6-15 m (no theoretical limit)
- Ratio: 2:1 to 3:1 (length:width)

**Wall Heights:**
- North wall: 2.0-3.0 m (taller for steeper roof)
- South wall: 0.5-1.5 m (or at grade)
- Knee wall (if above grade): 0.3-0.6 m

### Excavation Volume

```
V_excavation = (L × W × D_avg) + V_berm

Where:
L = Length (m)
W = Width (m)
D_avg = Average excavation depth (m)
V_berm = Volume for earth berms (typically 30-50% of excavation)
```

**Example (4m × 6m × 2m deep):**
```
V_excavation = 4 × 6 × 2 = 48 m³
Plus berm material: ~15-25 m³ additional
Total excavation: ~65-75 m³
```

---

## Drainage and Waterproofing

### Critical Importance

**#1 failure mode for earth-sheltered greenhouses is water infiltration.**

### Site Assessment

Before construction:
1. **Water table depth:** Must be >1.5 m below floor level
2. **Soil percolation:** Percolation test for drainage capacity
3. **Surface drainage:** Ensure water flows away from site
4. **Seasonal variations:** Check spring high water levels

### Drainage System Components

**1. Perimeter French Drain**
```
                    Grade
    ════════════════════════════════════════
         ↓  ↓  ↓  ↓  ↓  Water flows to drain
    ▓▓▓▓▓│                    │▓▓▓▓▓
    ▓▓▓▓▓│      Greenhouse    │▓▓▓▓▓
    ▓▓▓▓▓│                    │▓▓▓▓▓
    ═════╧════════════════════╧═════
         ╔════════════════════╗
         ║  Gravel bed 150mm  ║
         ╠════════════════════╣
         ║ Perforated pipe    ║ → To daylight
         ╠════════════════════╣     or sump
         ║  Gravel bed 150mm  ║
         ╚════════════════════╝
```

**Specifications:**
- Pipe: 100mm perforated corrugated or rigid PVC
- Gravel: 19-38mm washed, free of fines
- Filter fabric: Wrap gravel bed
- Slope: Minimum 1% (1 cm per meter)
- Outlet: To daylight, dry well, or sump pump

**2. Sub-Floor Drainage**
```
    ════════════════════════════════════
    ║ Concrete/gravel floor  75-100mm ║
    ╠════════════════════════════════════╣
    ║ Drainage gravel layer  100-150mm ║
    ╠════════════════════════════════════╣
    ║ Filter fabric                    ║
    ╠════════════════════════════════════╣
    ║ Compacted subgrade               ║
    ════════════════════════════════════
          Collection pipe to sump
```

**3. Interior Sump Pit**
- Size: 450mm × 450mm × 600mm deep minimum
- Pump: Submersible, automatic float switch
- Capacity: Match expected inflow rate
- Backup: Battery backup or gravity overflow

### Waterproofing Options

| Method | Cost | Durability | DIY Feasibility |
|--------|------|------------|-----------------|
| Bentonite clay | $ | Good | Excellent |
| Asphalt emulsion | $ | Moderate | Good |
| Rubber membrane (EPDM) | $$ | Excellent | Good |
| Bituminous membrane | $$ | Very Good | Moderate |
| Concrete waterproofing admixture | $$ | Good | N/A (in concrete) |
| Spray-applied membrane | $$$ | Excellent | Poor (specialized) |

### Recommendations by Soil Type

| Soil Type | Primary Strategy | Secondary Strategy |
|-----------|------------------|-------------------|
| Sandy/gravel | Perimeter drain only | Sump as backup |
| Sandy loam | Perimeter + sub-floor drains | Sump required |
| Clay loam | Full waterproofing + drains | Sump required |
| Heavy clay | Waterproof membrane + drains | Dual sumps recommended |
| High water table | Consider above-grade bermed | Avoid fully buried |

---

## Construction Methods

### Excavation Options

**1. Machine Excavation**
- Backhoe: Best for most projects, ~$100-200/hr
- Mini-excavator: Tight access, ~$150-250/day rental
- Skid steer: Good for spreading/berming

**2. Hand Excavation**
- Feasible for small projects (<20 m³)
- Time: ~1 m³ per person-hour in average soil
- Tools: Spade, mattock, wheelbarrow

### Wall Construction Methods

**1. Concrete Block (CMU)**
- Standard: 200mm (8") block
- Reinforced: Vertical rebar every 1.2m, filled with concrete
- Waterproofing: Exterior membrane + drainage
- Insulation: Interior or exterior rigid foam

**2. Poured Concrete**
- Thickness: 150-200mm (6-8")
- Reinforcement: #4 rebar at 400mm o.c. both ways
- Best for: Large projects, wet sites
- Waterproofing: Integral admixture + membrane

**3. Treated Wood (Pressure-Treated)**
- Posts: 150×150mm (6×6") or 200×200mm (8×8")
- Sheathing: Pressure-treated plywood
- Waterproofing: Polyethylene sheet + drain board
- Lifespan: 25-40 years with proper treatment

**4. Insulated Concrete Forms (ICF)**
- Excellent insulation (R-22 to R-30)
- Integral waterproofing
- Easy construction
- Higher material cost

**5. Rammed Earth / Adobe**
- Traditional method, very high thermal mass
- Requires dry climate or good drainage
- Stabilized with 5-10% Portland cement for durability
- Not recommended for wet Canadian climates without modification

### Glazing Installation

**Steep Roof Requirements (>45°):**
- Structural framing: 50×100mm or larger at 600mm o.c.
- Polycarbonate: Multi-wall panels with H-channel connectors
- Sealing: Butyl tape + aluminum capping
- Ventilation: Ridge vent or operable panels at top

**Snow Load Considerations:**
- Steep angles (>60°) shed snow naturally
- Shallower angles need structural calculation
- Consider snow guards if shading is acceptable
- Ridge should be high enough that snow slide doesn't block ventilation

### Insulation Placement

**North/Side Walls (Earth-Contact):**
- Option A: Interior insulation (easier, reduces thermal mass benefit)
- Option B: Exterior insulation (better, protects waterproofing)
- Option C: No insulation on deep portions (if depth >2.5m)

**Recommended R-Values:**

| Component | Zone 5 | Zone 6 | Zone 7+ |
|-----------|--------|--------|---------|
| North wall (above grade) | R-20 | R-25 | R-30 |
| North wall (below grade) | R-10 | R-15 | R-20 |
| Side walls (above grade) | R-15 | R-20 | R-25 |
| Side walls (below grade) | R-5 to R-10 | R-10 | R-15 |
| Floor perimeter | R-10 | R-10 | R-15 |
| Floor center | None | None | R-5 |

---

## Cost Analysis

### Excavation Costs

| Method | Cost per m³ | Notes |
|--------|-------------|-------|
| Machine (soft soil) | $15-25 | Including removal |
| Machine (hard soil/rock) | $40-80 | May need breaking |
| Hand excavation | $30-50 | Labor intensive |
| Spoil disposal | $10-20 | If can't use on-site |

### Wall Construction Costs (per m² of wall)

| Method | Materials | Labor | Total |
|--------|-----------|-------|-------|
| Concrete block | $40-60 | $50-80 | $90-140 |
| Poured concrete | $50-80 | $60-100 | $110-180 |
| Treated wood | $30-50 | $40-60 | $70-110 |
| ICF | $80-120 | $40-60 | $120-180 |

### Total Project Cost Estimates

**Small Earth-Bermed Greenhouse (4m × 6m = 24 m²):**

| Component | Low Estimate | High Estimate |
|-----------|--------------|---------------|
| Excavation (50 m³) | $1,000 | $2,000 |
| Foundation/walls | $2,500 | $5,000 |
| Waterproofing/drainage | $500 | $1,500 |
| Insulation | $400 | $800 |
| Glazing (steep roof) | $1,500 | $3,000 |
| Framing/structure | $800 | $1,500 |
| Electrical/ventilation | $300 | $800 |
| Finishing | $300 | $600 |
| **Total** | **$7,300** | **$15,200** |
| **Cost per m²** | **$304** | **$633** |

**Medium Production Greenhouse (6m × 12m = 72 m²):**

| Component | Low Estimate | High Estimate |
|-----------|--------------|---------------|
| Excavation (150 m³) | $2,500 | $5,000 |
| Foundation/walls | $6,000 | $12,000 |
| Waterproofing/drainage | $1,500 | $3,500 |
| Insulation | $1,000 | $2,000 |
| Glazing | $4,000 | $8,000 |
| Framing/structure | $2,000 | $4,000 |
| Electrical/ventilation | $800 | $2,000 |
| Finishing | $500 | $1,200 |
| **Total** | **$18,300** | **$37,700** |
| **Cost per m²** | **$254** | **$524** |

### Comparison to Conventional Greenhouse

| Type | Cost per m² | Annual Heating Cost* | 10-Year Total |
|------|-------------|---------------------|---------------|
| Unheated hoop house | $50-100 | N/A | $50-100 |
| Single-wall greenhouse | $150-250 | $40-80/m² | $550-1,050 |
| Double-wall greenhouse | $200-350 | $25-50/m² | $450-850 |
| Earth-sheltered | $250-550 | $5-15/m² | $300-700 |

*Assuming Zone 6 climate, electric heat

**Break-even period vs. conventional:** 3-7 years depending on heating costs and climate severity.

---

## Calculator Parameters

### Required User Inputs

For the greenhouse design calculator, collect these parameters:

```go
type EarthShelteredParams struct {
    // Location
    Latitude           float64 // degrees N
    ClimateZone        int     // HDD zone (4-8)
    MeanAnnualTemp     float64 // °C
    DesignWinterTemp   float64 // °C (2.5% design)
    FrostDepth         float64 // meters

    // Dimensions
    Length             float64 // meters (E-W)
    Width              float64 // meters (N-S)
    ExcavationDepth    float64 // meters
    NorthWallHeight    float64 // meters
    SouthWallHeight    float64 // meters (often 0 for at-grade)

    // Construction
    DesignType         string  // "walipini", "bermed", "chinese"
    WallConstruction   string  // "concrete", "block", "wood", "icf"
    InsulationR        float64 // m²·K/W (below grade)
    InsulationRAbove   float64 // m²·K/W (above grade)

    // Soil Properties
    SoilType           string  // "sand", "loam", "clay"
    SoilMoisture       string  // "dry", "moist", "saturated"
    WaterTableDepth    float64 // meters below grade

    // Glazing
    GlazingType        string  // "single", "double_poly", "twinwall", "triplewall"
    GlazingU           float64 // W/m²·K (auto-set from type)
    RoofAngle          float64 // degrees from horizontal

    // Target Conditions
    TargetTempNight    float64 // °C (minimum acceptable)
    TargetTempDay      float64 // °C (maximum before venting)
}
```

### Calculated Outputs

```go
type EarthShelteredResults struct {
    // Ground Temperatures
    GroundTempAtDepth     float64 // °C at floor level
    GroundTempWinter      float64 // °C at floor level, Jan
    GroundTempSummer      float64 // °C at floor level, July
    DampingDepth          float64 // meters

    // Solar Analysis
    WinterSolsticeSunAngle float64 // degrees
    RecommendedRoofAngle   float64 // degrees
    MaxWidthNoShading      float64 // meters (for flat roof)
    ShadingPercentWinter   float64 // % of floor shaded

    // Heat Loss Analysis
    GlazingArea           float64 // m²
    EarthContactArea      float64 // m² (walls + floor)
    ExposedWallArea       float64 // m² (above grade)

    HeatLossGlazing       float64 // W at design temp
    HeatLossEarthContact  float64 // W at design temp
    HeatLossExposedWall   float64 // W at design temp
    HeatLossInfiltration  float64 // W at design temp
    HeatLossTotal         float64 // W at design temp

    // Comparison
    ConventionalHeatLoss  float64 // W (same size above-ground)
    HeatLossReduction     float64 // % reduction
    GeothermalGain        float64 // W (heat from ground)

    // Economics
    EstimatedConstructionCost float64 // $
    AnnualHeatingCost         float64 // $
    AnnualSavingsVsConventional float64 // $
    PaybackYears              float64

    // Excavation
    ExcavationVolume      float64 // m³
    SpoilForBerms         float64 // m³
    ExportVolume          float64 // m³ (if any)

    // Drainage
    DrainageRequired      string  // "minimal", "standard", "extensive"
    SumpRecommended       bool
    WaterproofingLevel    string  // "basic", "moderate", "full"
}
```

### Key Calculation Functions

```go
// Ground temperature at depth using Kusuda equation
func GroundTemperature(depth, dayOfYear, meanTemp, amplitude, diffusivity float64) float64 {
    // D = damping depth
    D := math.Sqrt(365.0 * diffusivity / math.Pi)

    // Phase shift (day of minimum surface temp, ~35 for Canada)
    phaseShift := 35.0

    // Kusuda equation
    T := meanTemp - amplitude * math.Exp(-depth/D) *
         math.Cos(2*math.Pi*(dayOfYear-phaseShift)/365.0 - depth/D)

    return T
}

// Winter solstice sun angle at solar noon
func WinterSunAngle(latitude float64) float64 {
    // Declination at winter solstice = -23.45°
    declination := -23.45
    return 90.0 - math.Abs(latitude - declination)
}

// Recommended roof angle for winter
func RecommendedRoofAngle(latitude float64) float64 {
    sunAngle := WinterSunAngle(latitude)
    // Perpendicular to winter sun
    return 90.0 - sunAngle
}

// Maximum width before winter floor shading (flat roof)
func MaxWidthNoShading(northWallHeight, latitude float64) float64 {
    sunAngle := WinterSunAngle(latitude)
    return northWallHeight * math.Tan(sunAngle * math.Pi / 180.0)
}

// Heat loss through earth-contact surface
func EarthContactHeatLoss(area, insR, soilK, depth, tInside, tGround float64) float64 {
    // Effective soil R-value (simplified)
    soilR := 0.5 * depth / soilK

    // Total R-value
    totalR := 0.12 + insR + soilR  // inside film + insulation + soil

    // U-value
    U := 1.0 / totalR

    // Heat loss (to ground temperature, not outside air)
    return U * area * (tInside - tGround)
}

// Thermal diffusivity by soil type
func SoilDiffusivity(soilType, moisture string) float64 {
    // Returns m²/day
    diffusivityTable := map[string]map[string]float64{
        "sand": {"dry": 0.024, "moist": 0.045, "saturated": 0.052},
        "loam": {"dry": 0.020, "moist": 0.040, "saturated": 0.045},
        "clay": {"dry": 0.018, "moist": 0.030, "saturated": 0.042},
    }
    return diffusivityTable[soilType][moisture]
}

// Soil thermal conductivity
func SoilConductivity(soilType, moisture string) float64 {
    // Returns W/m·K
    conductivityTable := map[string]map[string]float64{
        "sand": {"dry": 0.35, "moist": 1.5, "saturated": 2.2},
        "loam": {"dry": 0.4, "moist": 1.2, "saturated": 1.6},
        "clay": {"dry": 0.5, "moist": 1.0, "saturated": 1.5},
    }
    return conductivityTable[soilType][moisture]
}
```

### Sample Calculation: Ottawa Greenhouse

**Inputs:**
- Latitude: 45.4°N
- Mean annual temp: 6.6°C
- Design winter temp: -24°C
- Size: 4m × 6m
- Depth: 2.0 m
- North wall: 2.5 m
- Soil: Moist clay

**Ground Temperature (Jan 15, depth 2.0m):**
```
D = √(365 × 0.030 / π) = 1.87 m
T = 6.6 - 15 × exp(-2.0/1.87) × cos(2π(15-35)/365 - 2.0/1.87)
T ≈ 6.6 - 15 × 0.34 × cos(-0.31 - 1.07)
T ≈ 6.6 - 5.1 × cos(-1.38)
T ≈ 6.6 - 5.1 × 0.19
T ≈ 5.6°C
```

**Winter Sun Angle:**
```
α = 90° - |45.4° - (-23.45°)| = 90° - 68.85° = 21.15°
```

**Recommended Roof Angle:**
```
Roof = 90° - 21.15° = 68.85° ≈ 70°
```

**Heat Loss Comparison:**

| Component | Earth-Sheltered | Conventional |
|-----------|-----------------|--------------|
| Glazing (18 m²) | 18 × 2.4 × 39 = 1,685 W | 18 × 2.4 × 39 = 1,685 W |
| North wall (10 m²) | 10 × 0.4 × 9.4 = 38 W | 10 × 0.35 × 39 = 137 W |
| Side walls (24 m²) | 24 × 0.5 × 6 = 72 W | 24 × 0.35 × 39 = 328 W |
| Floor (24 m²) | 24 × 0.15 × 9.4 = 34 W | F × P × ΔT = 1.0 × 20 × 39 = 780 W |
| Infiltration | ~200 W | ~400 W |
| **Total** | **~2,030 W** | **~3,330 W** |
| **Reduction** | **39%** | (baseline) |

---

## Case Studies

### 1. Newfoundland Earth-Sheltered Greenhouse

**Location:** St. John's, Newfoundland (47.5°N, Zone 6)
**Size:** 4.3m × 7.3m (14' × 24')
**Construction Cost:** ~$20,000 CAD (including monitoring systems)
**Annual Heating Cost:** <$400

**Design Features:**
- Concrete back wall with 50mm rigid foam insulation
- Earth-bermed on three sides
- Polycarbonate front roof/wall
- Concrete thermal mass maintains >5°C with minimal supplemental heat
- Digital monitoring system tracks performance

**Performance:**
- Interior maintained above 5°C through winter with 6-hour winter day length
- Baseboard electric heat as backup only

### 2. Idaho Earth-Sheltered Greenhouse (Mike Oehler)

**Location:** Northern Idaho (48°N, Zone 6)
**Size:** Multiple structures, 30+ years experience
**Construction Cost:** Minimal (owner-built)
**Heating:** None required

**Design Features:**
- "Uphill patio" design for light from above
- PSP (Post/Shoring/Polyethylene) construction
- Produces tomatoes in December without supplemental heat

### 3. Siberian Underground Greenhouse/Barn

**Location:** Siberia (extreme climate, -50°C winters)
**Size:** 81 m² greenhouse + 108 m² barn
**Heating:** Two compost heat piles (planned)

**Design Features:**
- Fully underground construction
- Target temperature: 10°C in -50°C external conditions
- Compost heating provides 9-12 kW continuous for 16-20 months

### 4. Chinese Solar Greenhouse

**Location:** Northern China (similar to Canadian prairies)
**Size:** Standardized designs, typically 8m × 60m
**Heating:** Solar only (no supplemental)

**Design Features:**
- 0.5-1.0 m thick north wall (brick, rammed earth)
- Removable thermal blanket covers glazing at night
- Produces vegetables year-round in -20°C climates
- >800,000 in operation

---

## References

### Walipini and Pit Greenhouse Design

1. [Benson Institute - Walipini Construction Manual (2002)](https://wiki.opensourceecology.org/images/1/1c/Walipini.pdf) - Original technical guide
2. [Mother Earth News - Walipini Construction Tips](https://www.motherearthnews.com/organic-gardening/essential-tips-for-building-a-durable-walipini-greenhouse-zbcz1706/)
3. [Ceres Greenhouse - Walipini Considerations](https://ceresgs.com/the-walipini-low-down/) - Northern latitude adaptations
4. [BC Greenhouses - Underground Greenhouses](https://info.bcgreenhouses.com/en-ca/en-ca/walipini-underground-greenhouses)

### Earth-Sheltered Construction

5. [Mike Oehler - The Earth-Sheltered Solar Greenhouse Book](https://undergroundhousing.com/greenhouse_book.html) - Essential reference
6. [US DOE - Efficient Earth-Sheltered Homes](https://www.energy.gov/energysaver/efficient-earth-sheltered-homes)
7. [OSTI - Earth-Sheltered Construction Case Study](https://www.osti.gov/biblio/5284616) - Performance measurements

### Thermal Calculations

8. [NRC Canada - Ground Temperatures (CBD-180)](https://nrc-publications.canada.ca/eng/view/ft/?id=386ddf88-fe8d-45dd-aabb-0a55be826f3f)
9. [Build It Solar - Earth Temperatures](https://www.builditsolar.com/Projects/Cooling/EarthTemperatures.htm)
10. [Oklahoma State - Soil Temperature Theory](http://soilphysics.okstate.edu/software/SoilTemperature/document.pdf)
11. [EnergyPlus - Ground Heat Transfer Calculations](https://bigladdersoftware.com/epx/docs/8-2/engineering-reference/ground-heat-transfer-calculations.html)

### Solar Geometry

12. [NOAA Solar Position Calculator](https://gml.noaa.gov/grad/solcalc/azel.html)
13. [SunCalc](https://www.suncalc.org/) - Interactive sun position tool
14. [Ceres Greenhouse - Roof Angle Guide](https://ceresgs.com/whats-the-best-roof-angle-for-a-solar-greenhouse/)

### Canadian Climate Data

15. [Environment Canada - Historical Climate Data](https://climate.weather.gc.ca/)
16. [ClimateData.ca](https://climatedata.ca/) - Projections and design data

### Case Studies

17. [Canada's Local Gardener - Earth-Sheltered Greenhouse (Newfoundland)](https://issuu.com/pegasuspublicationsinc/docs/canadaslocalgardener_vol2_iss4_digital/s/13326100)
18. [Permies - Siberian Earth-Sheltered Greenhouse](https://permies.com/t/32195/Wonderful-AMAZING-News-Earth-Sheltered)
19. [Verge Permaculture - Passive Solar Greenhouse Design](https://vergepermaculture.ca/designing-your-own-passive-solar-greenhouse-part-3/)

### Soil Properties

20. [Abu-Hamdeh (2000) - Soil Thermal Conductivity Effects](https://acsess.onlinelibrary.wiley.com/doi/abs/10.2136/sssaj2000.6441285x)
21. [Oklahoma State - Soil Thermal Properties](https://open.library.okstate.edu/rainorshine/chapter/13-2-soil-thermal-properties/)

### Building Codes and Standards

22. [ASHRAE Standard 90.1 - Energy Standard for Buildings](https://www.ashrae.org/) - F-factor tables
23. [Energy Star - F-Factor Guide](https://www.energystar.gov/sites/default/files/F-Factor%20Guide%20v1%202023-12-01.pdf)
