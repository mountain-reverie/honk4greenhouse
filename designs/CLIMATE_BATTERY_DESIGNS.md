# Climate Battery System Design Options

## Overview

Climate battery systems (also known as ground-to-air heat transfer systems, earth tubes, or subterranean heating and cooling systems) leverage the stable temperature of the earth to moderate greenhouse temperatures. These systems store excess daytime heat in the soil for release at night, and can provide passive cooling in summer.

This document provides comprehensive technical documentation for climate battery system design, with a focus on Canadian climate conditions and building requirements.

**Related Documentation:**
- [THERMAL_MASS.md](THERMAL_MASS.md) - Above-ground thermal mass systems (water barrels, PCM, concrete)
- Climate batteries store heat *underground* while thermal mass stores heat *inside* the greenhouse
- For optimal performance, combine both approaches (see Option 5: Hybrid Systems)

**Terminology:**
- **Climate Battery**: General term for underground thermal storage using air circulation (used throughout this document)
- **GAHT®** (Ground-to-Air Heat Transfer): Ceres Greenhouse Solutions' trademarked closed-loop system
- **Earth Tubes**: One-way airflow systems drawing outside air through buried pipes
- **SHCS** (Subterranean Heating and Cooling System): Academic terminology for underground thermal storage

> **Note:** "GAHT" is a registered trademark of Ceres Greenhouse Solutions. This document uses the generic term "climate battery" except when specifically referring to Ceres products.

---

## Climate Battery Fundamentals

### How Climate Battery Systems Work

Climate battery systems exploit the thermal stability of underground soil to moderate greenhouse temperatures. The fundamental principle is straightforward: soil below the frost line maintains a relatively constant temperature year-round, typically 6-12°C in inhabited regions of Canada.

```
                    DAYTIME (Heating Mode)

    [Greenhouse Interior - Hot Air ~30C]
                    |
                    v
              [Intake Fan]
                    |
                    v
    ====================================
    |     Underground Pipe Network     |  <-- Soil ~10C absorbs heat
    |     (buried 1.5-2.5m deep)      |
    ====================================
                    |
                    v
         [Return Air ~18C to greenhouse]


                   NIGHTTIME (Cooling Mode)

    [Greenhouse Interior - Cool Air ~8C]
                    |
                    v
              [Intake Fan]
                    |
                    v
    ====================================
    |     Underground Pipe Network     |  <-- Soil ~10C releases heat
    |          (warmed soil)           |
    ====================================
                    |
                    v
         [Return Air ~12C to greenhouse]
```

### Thermodynamic Principles

**Heat Exchange Mechanism:**

When warm, humid greenhouse air is circulated through underground pipes, two heat transfer processes occur:

1. **Sensible Heat Transfer**: Heat flows from warm air to cooler soil via conduction through the pipe wall
2. **Latent Heat Transfer**: Water vapor condenses on cool pipe surfaces, releasing latent heat (approximately 2,260 kJ/kg)

The latent heat component can be substantial - research indicates that condensation-based heat storage can transfer roughly 100 times more energy than sensible heating alone.

**Heat Transfer Rate Formula:**

```
Q/L = 2 * pi * k * (T_pipe - T_soil) / ln(4z/D)

Where:
Q/L = Heat transfer per unit length (W/m)
k   = Soil thermal conductivity (W/m-K)
T_pipe = Pipe surface temperature (C)
T_soil = Undisturbed soil temperature (C)
z   = Burial depth (m)
D   = Pipe diameter (m)
```

**Coefficient of Performance (COP):**

Research from Appalachian State University found that a properly designed climate battery system achieved:
- Daily heat storage: 128,588 BTU (37.7 kWh thermal)
- Average COP: 2.38 (when continuously running)
- A COP of 2.0 or higher is considered effective performance

### Seasonal Operation Modes

**Winter Heating Mode:**
- Fan activates when greenhouse temperature exceeds setpoint (e.g., 21C)
- Warm air circulates through underground pipes
- Heat transfers to soil, "charging" the thermal battery
- Cooled air returns to greenhouse, preventing overheating
- At night, process reverses - soil releases stored heat

**Summer Cooling Mode:**
- Underground soil remains cooler than peak greenhouse temperatures
- System provides natural cooling by circulating air through cool soil
- Also dehumidifies air as moisture condenses on cool pipe surfaces
- Reduces or eliminates need for mechanical cooling

**Control System:**
Every climate battery fan should be operated by two thermostats:
- One turns the fan ON when greenhouse is too cold (heating mode)
- One turns the fan ON when greenhouse is too hot (cooling mode)
- Fan remains OFF when temperature is within acceptable range

Example setpoints:
- Climate battery fan heating mode: Activates at 21°C
- Climate battery fan cooling mode: Activates at 27°C
- Exhaust fans: Set 3-5°C higher than climate battery cooling setpoint (e.g., 32°C)

---

## Ground Temperature Data for Canada

### Regional Ground Temperature Variation

Undisturbed ground temperature (UGT) across Canada varies significantly by latitude:

| Region | Latitude | Ground Temperature Range |
|--------|----------|-------------------------|
| Southern Canada (inhabited) | 42-52N | 6 to 12C |
| Central Canada | 52-60N | 0 to 6C |
| Northern Canada | 60-68N | -7 to 0C |
| Arctic (permafrost) | 68N+ | -15 to -2C |

**Key observations:**
- For the same latitudes, western Canada is generally warmer than eastern Canada
- Snow cover provides significant insulation, raising ground temperatures
- South-facing slopes have higher average ground temperatures than north-facing slopes

### Ground Temperature by Major Canadian City (Estimated at 2m depth)

| City | Province | Approx. Ground Temp | Climate Zone |
|------|----------|---------------------|--------------|
| Vancouver | BC | 10-11C | Zone 4 |
| Victoria | BC | 10-12C | Zone 4 |
| Toronto | ON | 9-10C | Zone 5 |
| Ottawa | ON | 8-9C | Zone 5 |
| Montreal | QC | 7-9C | Zone 5 |
| Calgary | AB | 6-8C | Zone 6 |
| Edmonton | AB | 5-7C | Zone 6 |
| Winnipeg | MB | 4-6C | Zone 7 |
| Saskatoon | SK | 4-6C | Zone 7 |
| Halifax | NS | 7-9C | Zone 5 |

### Temperature Stability vs Depth

Ground temperature fluctuation decreases with depth:

| Depth | Temperature Variation | Notes |
|-------|----------------------|-------|
| Surface | +/- 20-30C | Follows air temperature |
| 0.5m | +/- 10-15C | Significant seasonal swing |
| 1.0m | +/- 5-8C | Reduced variation |
| 1.5m | +/- 3-5C | More stable |
| 2.0m | +/- 2-3°C | Good stability for climate batteries |
| 3.0m | +/- 1-2C | Near-constant temperature |
| 6.0m+ | +/- 0.5C | Essentially constant |

**Thermal Lag:**
At 5-6m depth, maximum ground temperature occurs approximately 6 months after surface maximum. This lag can be beneficial for greenhouse heating, as ground warmth peaks in late fall/early winter.

### Frost Depth by Region

Critical for determining minimum climate battery burial depth:

| Location | Frost Depth | Building Code Requirement |
|----------|-------------|--------------------------|
| Vancouver, BC | 0.3-0.6m | 0.5m minimum |
| Southern Ontario | 1.0-1.2m | 1.2m (Part 9 OBC) |
| Toronto | 1.0-1.2m | 1.2m typical |
| Ottawa | 1.2-1.5m | 1.4m recommended |
| Montreal | 1.2-1.5m | 1.4m typical |
| Calgary | 1.5-1.8m | 1.8m recommended |
| Edmonton | 1.8-2.1m | 2.1m recommended |
| Winnipeg | 2.1-2.4m | 2.4m recommended |
| Northern ON/Prairies | 2.4m+ | Consult local codes |

**Climate battery burial depth must exceed local frost depth by at least 0.3m to ensure stable soil temperatures.**

---

## Soil Thermal Properties

### Thermal Conductivity by Soil Type

Soil thermal conductivity significantly affects climate battery performance:

| Soil Type | Dry k (W/m-K) | Saturated k (W/m-K) | Notes |
|-----------|---------------|---------------------|-------|
| Sand | 0.15-0.25 | 1.5-2.5 | Highest when saturated |
| Sandy Loam | 0.19-0.35 | 1.0-1.5 | Good balance |
| Loam | 0.25-0.40 | 0.7-1.2 | Common agricultural soil |
| Clay Loam | 0.30-0.45 | 0.5-0.8 | Lower conductivity |
| Clay | 0.35-0.50 | 0.4-0.7 | Lowest, insulating |

**Key Insight:** Sandy soils have higher thermal conductivity due to quartz content. For climate battery applications, sandy loam with adequate moisture provides optimal heat transfer.

### Thermal Diffusivity

Thermal diffusivity indicates how quickly temperature changes propagate through soil:

| Soil Type | Diffusivity (m^2/s x 10^-7) |
|-----------|-----------------------------|
| Dry sand | 2.0-3.0 |
| Moist sand | 7.0-10.0 |
| Dry clay | 1.5-2.5 |
| Moist clay | 4.0-6.0 |
| Organic soil | 1.0-2.0 |

### Soil Recommendations for Climate Batteries

**Ideal Soil Characteristics:**
- Sandy loam composition (high thermal conductivity)
- Clay content below 25% (prevents insulating buildup around pipes)
- Adequate moisture retention without saturation
- Well-drained to prevent water table issues

**If native soil is clay-heavy:**
- Consider importing sandy fill around pipe network
- Mix native soil with sand at 1:1 ratio
- Ensure good drainage to maintain optimal moisture

**Volumetric Heat Capacity:**
- Dry soil: ~0.8-1.0 MJ/(m^3-K)
- Moist soil: ~2.0-2.5 MJ/(m^3-K)
- Saturated soil: ~3.0-4.0 MJ/(m^3-K)

---

## Climate Battery Design Parameters

### Pipe Material Selection

| Material | Advantages | Disadvantages | Cost | Best Use |
|----------|------------|---------------|------|----------|
| Corrugated HDPE (perforated) | Low cost, easy installation, handles condensation | Corrugations trap debris, harder to clean | $ | Budget systems |
| Corrugated HDPE (solid) | Low cost, blocks groundwater infiltration | Requires drainage slope, no condensation release | $ | High water table areas |
| Double-wall HDPE (smooth interior) | Best airflow, corrugated exterior for strength | Higher cost | $$ | Recommended for most applications |
| Smooth PVC | Excellent airflow, easy cleaning | More rigid, difficult installation | $$ | Long straight runs |
| Concrete | Excellent thermal mass, durability | Very heavy, requires professional installation | $$$ | Commercial/permanent installations |

**Recommended:** Double-wall HDPE with corrugated exterior and smooth interior (e.g., ADS N-12 or equivalent) provides the best balance of performance, durability, and cost.

### Pipe Diameter Recommendations

| Diameter | Airflow Capacity | Heat Transfer | Best Application |
|----------|------------------|---------------|------------------|
| 100mm (4") | Low-Medium | Excellent surface/volume ratio | Small greenhouses, multiple pipe systems |
| 150mm (6") | Medium | Good balance | Most residential applications |
| 200mm (8") | High | Good | Larger greenhouses |
| 250mm+ (10"+) | Very High | Reduced efficiency | Commercial applications with turbulence features |

**Design Principle:** Multiple smaller pipes provide better heat transfer than fewer large pipes due to increased surface area to volume ratio.

Example: Four 100mm pipes provide 2x the surface area of one 200mm pipe with the same total cross-sectional area.

### Pipe Length Calculations

**Optimal Length Range:** 7.5-10.5m (25-35 feet) per run

Research indicates that 82-85% of total heat exchange occurs within the first 10m of pipe length. Beyond this, the air reaches thermal equilibrium with the soil and additional length provides diminishing returns.

**Length Guidelines:**

| Greenhouse Size | Recommended Length per Pipe |
|-----------------|----------------------------|
| < 50 m^2 | 6-8m |
| 50-100 m^2 | 8-10m |
| 100-200 m^2 | 10-12m |
| > 200 m^2 | 10-12m (use more pipes) |

**Calculation Formula:**
```
L_effective = 34m * (D / 1.0m)

Where:
L_effective = Length at which 85% of heat transfer is achieved
D = Pipe diameter (m)

For 150mm (0.15m) pipe:
L_effective = 34 * 0.15 = 5.1m (approximately)
```

### Burial Depth Recommendations

| Climate Zone | Minimum Depth | Optimal Depth | Maximum Practical |
|--------------|---------------|---------------|-------------------|
| Zone 4 (Coastal BC, Southern ON) | 1.5m | 1.8-2.2m | 2.5m |
| Zone 5 (Central ON, Montreal) | 1.8m | 2.0-2.5m | 3.0m |
| Zone 6 (Calgary, Edmonton) | 2.0m | 2.2-2.8m | 3.5m |
| Zone 7+ (Northern regions) | 2.4m | 2.5-3.0m | 4.0m |

**Critical Factors:**
- Must be below frost line + 0.3m minimum
- Must be above water table
- Deeper = more stable temperature but higher excavation cost
- Consider adding insulation above pipes if depth is limited

### Pipe Spacing

**Horizontal Spacing (between parallel pipes):**
- Minimum: 0.3m (1 ft) - prevents thermal interference
- Recommended: 0.4-0.6m (1.5-2 ft)
- Maximum: 1.0m (3 ft) - beyond this, excavation becomes inefficient

**Vertical Spacing (for multi-layer systems):**
- Minimum: 0.23m (9 inches)
- Recommended: 0.3m (12 inches)
- Allows adequate soil thermal recharge between layers

### Air Velocity and Fan Sizing

**Optimal Air Velocity:** 1.0-3.5 m/s (2.0 m/s recommended)

| Velocity | Effect |
|----------|--------|
| < 1.0 m/s | Insufficient airflow, poor greenhouse circulation |
| 1.0-2.0 m/s | Good heat transfer, low pressure drop |
| 2.0-3.5 m/s | Adequate heat transfer, moderate pressure drop |
| > 3.5 m/s | Reduced heat transfer (air doesn't equilibrate), high energy use |

**Reynolds Number Requirement:**
Maintain turbulent flow (Re > 4000) for effective heat transfer. Turbulent flow increases air-to-pipe heat transfer coefficient.

**Fan Sizing Formula:**
```
CFM = Volume (m^3) x ACH / 1.699

Where:
CFM = Required fan capacity (cubic feet per minute)
Volume = Greenhouse volume in cubic meters
ACH = Air changes per hour (5-10 recommended for climate batteries)

Example: 100 m^2 greenhouse, 3m average height = 300 m^3
CFM = 300 x 10 / 1.699 = 1,766 CFM
```

**Recommended Fan Types:**
- **Axial fans**: High volume, low pressure - good for short runs
- **Centrifugal/inline fans**: Higher pressure capability - better for long runs or multiple pipes
- Always specify fans rated for continuous operation

### Heat Exchange Efficiency Calculations

**Effectiveness-NTU Method:**
```
epsilon = Q_actual / Q_max
epsilon = 1 - exp(-NTU)  [for single-pass heat exchanger]

Where:
NTU = U * A / C_min
U = Overall heat transfer coefficient (W/m^2-K)
A = Pipe surface area (m^2)
C_min = Minimum heat capacity rate (W/K)
```

**Typical Values:**
- Overall heat transfer coefficient (U): 5-15 W/(m^2-K) for earth tubes
- Effectiveness: 0.5-0.8 for well-designed systems

**Daily Energy Storage Capacity:**
```
Q_daily = m_dot * c_p * delta_T * t

Where:
Q_daily = Daily energy storage (kWh)
m_dot = Mass flow rate of air (kg/s)
c_p = Specific heat of air (1.005 kJ/kg-K)
delta_T = Temperature difference (K)
t = Operating hours per day
```

---

## Climate Battery System Configurations

### Option 1: Multi-Pipe with Central Fan (Recommended for Small-Medium Greenhouses)

**Best for:** Greenhouses up to 100 m^2

**Configuration:**
- 6-12 parallel pipes (100-150mm diameter)
- Buried 1.8-2.2m deep
- Pipe length: 6-10m each
- Single manifold on each end
- One 200-400 CFM inline fan

**Visual Diagram:**
```
    [Greenhouse Floor]
    ==================
           |
    [Fan] -+-> [Manifold]
                  |
           +------+------+------+------+------+
           |      |      |      |      |      |
          [ ]    [ ]    [ ]    [ ]    [ ]    [ ]  <- 100mm corrugated pipes
           |      |      |      |      |      |
           +------+------+------+------+------+
                  |
            [Return Manifold]
                  |
         [Return to Greenhouse]

    Cross Section:

    Ground Level ________________________________

    0.5m        [Gravel bed for drainage]

    1.5m        O     O     O     O     O     O   <- Pipes (spaced 40cm)

    2.0m        [Undisturbed soil]
```

**Specifications:**
| Parameter | Value | Notes |
|-----------|-------|-------|
| Pipe material | Corrugated HDPE | Perforated optional for drainage |
| Pipe diameter | 100-150mm | Larger for more airflow |
| Burial depth | 1.8-2.2m | Below frost line in most of Canada |
| Pipe spacing | 40-60cm | Prevents thermal interference |
| Fan capacity | 200-400 CFM | 5-10 air changes per hour |
| Estimated cost | $800-1500 CAD | Materials only |
| Installation time | 2-3 days | With mini excavator |

**Advantages:**
- Compact footprint
- Easy installation
- Efficient airflow distribution
- Good for shorter greenhouses

**Disadvantages:**
- Limited capacity for large structures
- Requires manifold fabrication

---

### Option 2: Diagonal Multi-Layer Array (Atmos/Threefold Style)

**Best for:** Larger greenhouses (100-300 m², 9-30m width)

**Overview:**
This design, pioneered by [Atmos Greenhouse Systems](https://atmosgreenhouse.com) and documented at [Threefold Farm](https://threefold.farm/climate-battery-greenhouse), uses diagonal tubing runs across the greenhouse width in multiple vertical layers. The system is divided into separate "battery zones" for optimal heat distribution and redundancy.

**Key Design Principles:**
- **Equal tube lengths** for balanced airflow across all pipes
- **Diagonal runs** maximize tube length within greenhouse footprint
- **Multi-layer installation** (2-3 layers) increases thermal capacity
- **CFD-optimized** manifold design minimizes back-pressure and ensures even airflow
- **Perforated tubing** allows condensation drainage and latent heat transfer

**Configuration:**
- 2-4 separate battery zones distributed under greenhouse floor
- Each zone: vertical riser → horizontal manifold → diagonal tubes → manifold → riser
- 100mm (4") perforated corrugated drain tile
- 450mm (18") manifolds, 600mm (24") risers
- ~45-47 tube runs per manifold
- Burial depth: 0.6m to 2.4m (2-3 layers)

**Visual Diagram:**
```
    [Greenhouse Floor - Plan View, 9m x 29m example]
    +--------------------------------------------------+
    |  [Riser]        [Riser]        [Riser]           |
    |     |              |              |              |
    |  [=====]        [=====]        [=====]  <- 450mm Manifolds
    |   / / /          / / /          / / /            |
    |  / / /          / / /          / / /             |
    | / / /          / / /          / / /   <- 100mm diagonal tubes
    |/ / /          / / /          / / /               |
    |  [=====]        [=====]        [=====]  <- 450mm Manifolds
    |     |              |              |              |
    |  [Riser]        [Riser]        [Riser]           |
    |   [Fan]          [Fan]          [Fan]            |
    +--------------------------------------------------+
         Zone 1         Zone 2         Zone 3

    Cross Section (through one zone):

    Ground Level ________________________________
                        |                    |
    0.6m (2')     - - - O - - - O - - - O - - -   <- Layer 1 (top)
                        |                    |
    1.2m (4')     - - - O - - - O - - - O - - -   <- Layer 2 (middle)
                        |                    |
    1.8m (6')     - - - O - - - O - - - O - - -   <- Layer 3 (bottom, optional)
                        |                    |
    2.4m (8')     [Undisturbed soil / deep battery zone]

    Side View (one zone):

                   [Fan/HAF]
                      |
                   [Riser 600mm]
                      |
    Ground ________[Manifold 450mm]________
                   /  /  /  /  /  /
                  /  /  /  /  /  /    <- Diagonal 100mm tubes
                 /  /  /  /  /  /         (equal length ~9m each)
    Depth    ___[Manifold 450mm]___
    1.8-2.4m       |
                [Riser 600mm]
                   |
              [Return to greenhouse]
```

**Specifications (based on Threefold Farm 9m x 29m greenhouse):**
| Parameter | Value | Notes |
|-----------|-------|-------|
| Greenhouse size | 275 m² (30' x 96') | Reference installation |
| Total tubing | 1,295m (4,250 ft) | All zones combined |
| Tube material | 100mm (4") perforated corrugated HDPE | "Socked" drainage tile with filter fabric |
| Tubes per manifold | 45-47 runs | Equal length for balanced airflow |
| Tube length | ~9m (30') each | Diagonal across 9m width |
| Manifold pipe | 450mm (18") twin-wall HDPE | 6m (20') sections, capped ends |
| Riser pipe | 600mm (24") twin-wall HDPE | ~2m (6'8") height |
| Number of zones | 3 battery zones | Independent operation possible |
| Burial depth | 0.6-2.4m (2'-8') | Multiple layers |
| Layer spacing | 0.3-0.6m (12-24") | Between tube layers |
| Horizontal tube spacing | 0.3-0.4m (12-16") | Within each layer |
| Fan capacity | 3 x HAF fans, ~6,000 CFM each | ~2,000 CFM effective per fan after losses |
| Total power consumption | ~1,100 W | All fans running |
| Air changes/hour | 15-16 ACH | Based on greenhouse volume |
| Excavation depth | 2.4m (8') recommended | 1.8m (6') minimum |
| Estimated cost | $3,000-6,000 CAD | Materials only, larger systems |
| Installation time | 5-7 days | Full excavation required |

**Design Optimization (CFD Analysis):**

Atmos Greenhouse Systems uses Computational Fluid Dynamics (CFD) to optimize their designs. Key findings:

| Design Aspect | Poor Design | Optimized Design |
|---------------|-------------|------------------|
| Airflow rate | ~2,050 CFM | ~3,500 CFM |
| Static pressure | 0.57" | 0.33" |
| Velocity distribution | Unbalanced (rainbow pattern) | Uniform |
| Tube airflow | Some starved, some excess | Equal distribution |

**Critical Design Rules:**
1. **Keep all tubes equal length** - Unequal lengths cause airflow imbalance
2. **Minimize hard angles** - Use 45° elbows maximum, avoid 90° bends
3. **Size manifolds correctly** - Must handle total airflow without restriction
4. **Layer tubes properly** - Start at maximum depth, work upward
5. **Allow for in-ground planting** - Stop tubes 0.6m (2') below surface

**Advantages:**
- Proven design with documented 3+ year performance data
- Excellent thermal capacity (37+ kWh/day storage demonstrated)
- Balanced airflow through CFD-optimized design
- Redundancy with multiple independent zones
- Perforated tubes handle condensation naturally
- Enables in-ground planting above battery
- 60-80% heating cost reduction vs propane (Zone 6/7)

**Disadvantages:**
- Requires full excavation (not suitable for trenching)
- Higher upfront cost than simpler designs
- Minimum 1.8m (6') excavation depth required
- Not suitable for high water table sites
- Professional design consultation recommended

**Performance Data (Threefold Farm, Zone 6b/7a):**
- Maintained above -7°C (20°F) through Pennsylvania winter
- Successfully overwinters subtropical fruit trees
- System runs ~1,100W continuously during cold periods
- COP (Coefficient of Performance): 2.38 average

**References:**
- [Atmos Greenhouse Systems](https://atmosgreenhouse.com)
- [Atmos Climate Battery Design Comparison](https://atmosgreenhouse.com/blog/climate-battery-design-comparison)
- [Threefold Farm Climate Battery Greenhouse](https://threefold.farm/climate-battery-greenhouse)

---

### Option 3: Single Manifold (Both Sides)

**Best for:** Medium greenhouses with central aisle (50-150 m^2)

**Configuration:**
- Central intake manifold under walkway
- Pipes radiate to both sides
- Return manifolds along greenhouse edges
- Works well with raised bed layouts

**Visual Diagram:**
```
    [Greenhouse Floor - Plan View]
    +------------------------------------------+
    |   [Return]      [Fan]       [Return]     |
    |      |            |            |         |
    |  +---+            |            +---+     |
    |  |   |            |            |   |     |
    |  O   O   O   O    |    O   O   O   O     |  <- Pipes radiate outward
    |  |   |   |   |    |    |   |   |   |     |
    |  +---+---+---+----+----+---+---+---+     |
    |              [Central Manifold]          |
    |                                          |
    +------------------------------------------+

    Cross Section:

                    [Central Aisle]
    Ground Level _______|  |_______
                        |  |
    0.5m               [Manifold]
                       /        \
    1.5m              O          O    <- Pipes slope outward
                     /            \
    2.0m            O              O
```

**Specifications:**
| Parameter | Value | Notes |
|-----------|-------|-------|
| Pipe material | Smooth PVC | Easy connection to manifold |
| Pipe diameter | 100-150mm | Mixed sizes acceptable |
| Central manifold | 200-300mm | Sized for total airflow |
| Burial depth | 1.5-2.0m | Slight slope toward edges |
| Pipe length | 3-6m per side | From center to edge |
| Fan capacity | 400-600 CFM | Central location |
| Estimated cost | $1000-2000 CAD | Materials only |
| Installation time | 2-3 days | Less trenching than Option 2 |

**Advantages:**
- Balanced heat distribution
- Works well with central aisle layouts
- Single fan location
- Compact manifold design

**Disadvantages:**
- Requires precise manifold construction
- Central aisle must accommodate ductwork
- More complex installation

---

### Option 4: Under-Floor Plenum (Rock Bed)

**Best for:** High-capacity thermal storage, commercial applications

**Configuration:**
- Gravel or rock bed under greenhouse floor
- Perforated pipes distribute air through rock mass
- Rock provides additional thermal mass
- Often combined with radiant floor heating

**Visual Diagram:**
```
    [Greenhouse Floor - Cross Section]

    Floor Surface  ================================
                   [Vapor Barrier]
    0.0m           ________________________________
                   |                              |
    0.3m           |    Gravel/Rock Bed          |
                   |    (20-40mm washed stone)   |
                   |  O ~~~~~ O ~~~~~ O ~~~~~ O  |  <- Perforated distribution pipes
    0.6m           |                              |
                   |______________________________|
                   [Insulation Layer (EPS)]
    0.9m           ________________________________
                   |                              |
    1.5m+          |    Undisturbed Soil         |
                   |                              |
```

**Specifications:**
| Parameter | Value | Notes |
|-----------|-------|-------|
| Rock size | 20-40mm | Washed, rounded stone |
| Rock bed depth | 0.3-0.5m | Depending on thermal needs |
| Perforated pipe | 100-150mm | Spaced 0.6-1.0m apart |
| Insulation below | R-10 minimum | EPS or XPS foam board |
| Fan capacity | 600-1200 CFM | Higher capacity for rock bed |
| Estimated cost | $2000-4000 CAD | Higher due to rock volume |
| Installation time | 4-6 days | Significant excavation |

**Advantages:**
- Very high thermal mass (rock + soil)
- Even floor temperature
- Can integrate with radiant heating
- Excellent for commercial operations

**Disadvantages:**
- Highest installation complexity
- Significant excavation required
- Drainage critical to prevent water accumulation
- More expensive

#### Deep Winter Greenhouse Variant (UMN Design)

The **Deep Winter Greenhouse (DWG)** design, developed by the University of Minnesota Extension, is a specialized variant of the rock bed plenum system optimized for extreme northern climates (Zones 4-7). The design integrates passive solar architecture with underground rock bed thermal storage.

**Key Design Principles:**

1. **60° Glazing Angle**: The south-facing glazing is angled at approximately 60 degrees (adjusted for latitude) to maximize winter solar gain when the sun is low on the horizon. This steep angle captures more solar energy during critical heating months.

2. **Three Insulated Walls**: North, east, and west walls are heavily insulated (R-30 to R-40+) with reflective interior surfaces to bounce light back toward plants and thermal mass.

3. **Rock Bed Thermal Mass**: A 1.0-1.2m (3-4 feet) deep bed of washed river rock beneath the growing area stores heat absorbed from solar radiation during the day.

4. **Passive Air Circulation**: Warm air from the peak naturally convects down through the rock bed, while cooler air is drawn back up through the growing area—minimizing or eliminating the need for fans.

**Visual Diagram:**
```
    [DWG Cross Section - Facing East]

                        ╱╲
                       ╱  ╲  60° South-facing glazing
                      ╱    ╲
                     ╱      ╲
    Insulated       ╱        ╲
    North Wall  ███╱          ╲
    (R-30+)     ███            ╲
                ███  Growing    ╲
                ███   Area      ╲
                ███             ╲
    Floor Level ═══════════════════
                |               |
    0.3m        |  Washed Rock  |  ← Air circulation through rock
                |     Bed       |
    1.0-1.2m    | (40-80mm)     |
                |               |
                └───────────────┘
                [Perimeter Insulation]
```

**Specifications (UMN Design):**

| Parameter | Value | Notes |
|-----------|-------|-------|
| Glazing angle | 60° (±5° for latitude) | South-facing, optimized for winter |
| Insulation (opaque walls) | R-30 to R-40+ | North, east, west walls |
| Rock bed depth | 1.0-1.2m (3-4 ft) | Washed river rock, 40-80mm diameter |
| Rock bed coverage | Full floor area | Entire growing area footprint |
| Air circulation | Passive convection or low-power fan | Minimal energy input |
| Heat retention | ~3 days | Without solar recharge in cloudy weather |
| Minimum temperature | 4°C (40°F) | Zone 4 Minnesota winters |
| Typical size | 15-25 m² (160-270 ft²) | Optimized for home/market garden scale |
| Construction cost | $300-350/m² ($28-32/ft²) | DIY construction, ~$18,000-25,000 total |

**Performance Data (Minnesota Zone 4):**

- Maintains minimum 4°C (40°F) through Minnesota winters (-30°C ambient)
- Heat stored in rock bed lasts approximately 3 days without solar recharge
- Passive solar gain provides 80-90% of heating needs
- Minimal supplemental heat required (small electric or propane backup)
- Year-round growing of cold-hardy crops (greens, root vegetables)

**Rock Bed Specifications:**

| Parameter | Specification |
|-----------|---------------|
| Rock type | Washed river rock (rounded, no fines) |
| Rock size | 40-80mm (1.5-3 inches) diameter |
| Bed depth | 1.0-1.2m (3-4 feet) minimum |
| Air space ratio | ~40% void space between rocks |
| Drainage | Perforated drain tile at bottom |
| Vapor barrier | Required between rock and growing area |

**Advantages over Standard Under-Floor Design:**

- Passive operation reduces energy consumption to near zero
- 60° glazing optimizes winter solar collection
- Reflective north wall increases light to plants
- Proven performance in extreme cold climates
- Simpler construction than piped systems
- No fan maintenance or electrical requirements

**Disadvantages:**

- Requires specific architectural form (shed-style with angled south wall)
- Not suitable for retrofitting existing structures
- Rock bed adds significant excavation and material cost
- Limited to smaller greenhouse sizes (convection-driven)
- Orientation must be precise (south-facing ±15°)

**References:**

- [UMN Extension - Deep Winter Greenhouses](https://extension.umn.edu/growing-systems/deep-winter-greenhouses)
- [UMN College of Design - CSBR Deep Winter Greenhouse Project](https://design.umn.edu/center-sustainable-building-research/projects/deep-winter-greenhouse)
- [Northlands Winter Greenhouse Manual](https://extension.umn.edu/sites/extension.umn.edu/files/2022-11/northlands-winter-greenhouse-manual.pdf)

---

### Option 5: Hybrid Systems (Climate Battery + Thermal Mass)

**Best for:** Maximum performance, all-season growing

**Configuration:**
Combines underground pipe network with above-ground thermal mass elements:

**Components:**
1. **Underground climate battery system** (any configuration above)
2. **Water barrels on north wall** - 55-gallon drums painted black
3. **Insulated north wall** - R-20+ insulation or earth-bermed
4. **Phase change materials** (optional) - PCM panels for enhanced storage

**Integration Principles:**
- Water barrels absorb direct solar radiation during day
- Climate battery handles peak heat and stores excess underground
- Combined thermal mass moderates temperature swings
- PCM provides targeted temperature buffering at specific setpoints

**Thermal Mass Comparison:**
| Material | Heat Capacity | Volume Needed | Notes |
|----------|---------------|---------------|-------|
| Water | 4.18 kJ/(kg-K) | Lowest | Best specific heat |
| Concrete | 0.88 kJ/(kg-K) | 5x water | Structural benefit |
| Rock/Gravel | 0.84 kJ/(kg-K) | 5x water | Good drainage |
| Phase Change | Varies | 2-3x water | Targeted temperature |

**Water Barrel Guidelines:**
- Position along north wall in direct sunlight
- Paint containers flat black for maximum absorption
- Stack no more than 3 high for safety
- Allow 200-400 liters per 10 m^2 of greenhouse
- Climate battery soil provides approximately 2x storage capacity of equivalent water volume

---

### Option 6: Custom / Advanced

**Best for:** Experienced builders, unique greenhouse shapes, or specific requirements

**Description:**
This option allows users to design their own climate battery system with custom specifications. The wizard will prompt for:

- Number of pipes
- Pipe diameter (75mm - 200mm)
- Pipe material (PVC, HDPE, corrugated)
- Total pipe length
- Burial depth
- Fan configuration (single, dual, variable speed)
- Manifold design

**Design Calculator Interface:**

```
+--------------------------------------------------+
|  CUSTOM CLIMATE BATTERY DESIGN                   |
+--------------------------------------------------+
|                                                  |
|  Greenhouse Volume:  [calculated] m^3            |
|  Recommended Airflow: [calculated] CFM           |
|                                                  |
|  PIPE CONFIGURATION                              |
|  +--------------------------------------------+  |
|  | Number of pipes:    [ 8 ] [+] [-]          |  |
|  | Pipe diameter:      [100mm v]              |  |
|  | Pipe length (each): [ 8.0 ] m              |  |
|  | Burial depth:       [ 1.8 ] m              |  |
|  +--------------------------------------------+  |
|                                                  |
|  CALCULATED VALUES                               |
|  +--------------------------------------------+  |
|  | Total pipe length:  64.0 m                 |  |
|  | Soil contact area:  20.1 m^2               |  |
|  | Thermal capacity:   ~4.2 kWh/day           |  |
|  | Recommended fan:    350 CFM                |  |
|  +--------------------------------------------+  |
|                                                  |
|  [Import from template]  [Save as template]      |
|                                                  |
+--------------------------------------------------+
```

**Validation Rules:**
- Minimum burial depth: 1.2m (below frost line) - adjust for region
- Maximum burial depth: 3.0m (practical limit)
- Minimum pipe spacing: 0.3m (thermal interference)
- Airflow velocity: 1.5-3.0 m/s (optimal heat transfer)
- Total pipe length should provide 5-10 air changes per hour

---

## Comparison Table

| Feature | Multi-Pipe | Diagonal Array (Atmos) | Single Manifold | Under-Floor | Hybrid | Custom |
|---------|------------|------------------------|-----------------|-------------|--------|--------|
| Best greenhouse size | <100 m² | 100-300 m² | 50-150 m² | 100-500 m² | Any | Any |
| Installation complexity | Low | High | Medium | High | High | Variable |
| Material cost | $ | $$$ | $$ | $$$ | $$$$ | Variable |
| Excavation required | Minimal | Full excavation | Moderate | Major | Moderate | Variable |
| Thermal capacity | Good | Excellent (37+ kWh/day) | Very Good | Excellent | Maximum | Variable |
| Maintenance access | Easy | Limited (buried) | Moderate | Difficult | Easy | Variable |
| Expandability | Limited | Limited (fixed zones) | Limited | Limited | Good | Full |
| Proven performance | Good | Excellent (3+ yr data) | Good | Good | Excellent | Variable |
| CFD optimized | No | Yes | No | No | Optional | Optional |

---

## Canadian Climate Considerations

### Climate Zone Recommendations

#### Zone 4 (Southern Ontario, Lower BC Mainland, Vancouver Island)
- **HDD:** 2500-3999
- **Frost depth:** 0.5-1.0m
- Any climate battery design works well
- Shallower burial depth acceptable (1.5m)
- Smaller fan capacity sufficient
- Lowest heating demand - climate battery may provide 80%+ of heating needs

#### Zone 5-6 (Central Ontario, Montreal, Southern Prairies, Calgary)
- **HDD:** 4000-5999
- **Frost depth:** 1.2-2.0m
- Recommend Two-Trench or Single Manifold
- Burial depth: 1.8-2.5m
- Consider adding thermal mass (water barrels) as complement
- Perimeter insulation strongly recommended
- Climate battery may provide 60-80% of heating needs; backup heat required

#### Zone 7+ (Northern Ontario, Northern Prairies, Northern BC)
- **HDD:** 6000+
- **Frost depth:** 2.0-2.4m+
- Two-Trench or Under-Floor design recommended
- Maximum burial depth (2.5-3.0m)
- Larger fan capacity for faster heat recovery
- Combine with multiple thermal mass strategies
- Perimeter insulation essential (R-10 minimum)
- Climate battery provides supplemental heating; significant backup heat needed

### Frost Depth and Burial Requirements by Region

| City/Region | Frost Depth | Min. Burial | Recommended Burial |
|-------------|-------------|-------------|-------------------|
| Vancouver, BC | 0.45m | 0.75m | 1.5m |
| Victoria, BC | 0.30m | 0.60m | 1.2m |
| Kelowna, BC | 0.90m | 1.2m | 1.8m |
| Toronto, ON | 1.2m | 1.5m | 2.0m |
| Ottawa, ON | 1.4m | 1.7m | 2.2m |
| Montreal, QC | 1.4m | 1.7m | 2.2m |
| Calgary, AB | 1.8m | 2.1m | 2.5m |
| Edmonton, AB | 2.1m | 2.4m | 2.8m |
| Winnipeg, MB | 2.4m | 2.7m | 3.0m |
| Saskatoon, SK | 2.3m | 2.6m | 3.0m |
| Thunder Bay, ON | 1.8m | 2.1m | 2.5m |
| Whitehorse, YT | 2.5m+ | Consult local | Special design |

### Permafrost Considerations (Northern Regions)

**Special Challenges:**
- Permafrost begins within 1-2m in many northern areas
- Introducing warm air underground can thaw permafrost
- Thawed permafrost causes ground subsidence
- Climate batteries may not be suitable for continuous permafrost zones

**Alternative Approaches for Northern Climates:**
1. **Above-ground thermal mass:** Water barrels, rock walls, PCM
2. **Shallow earth-sheltered design:** Bermed walls without underground pipes
3. **Passive solar only:** Maximize glazing, insulation, and above-ground mass
4. **Heat recovery ventilation:** Capture and recycle exhaust heat

**Permafrost Zones in Canada:**
- Continuous: North of approximately 65N latitude
- Discontinuous: 55-65N (patches of permafrost)
- Sporadic: 50-55N (isolated permafrost in shaded areas)

### Snow Load and Inlet Protection

**Inlet Structure Requirements:**
- Elevate intake opening 0.5-1.0m above expected snow depth
- Use 45-degree elbow or vertical riser to prevent snow entry
- Install heavy-gauge screen (6mm mesh) to exclude rodents
- Consider heated inlet for extreme cold regions
- Provide secondary bypass inlet for spring/fall when ground is cooler than air

**Outlet Protection:**
- Position return inside greenhouse, elevated above floor
- Use diffuser to distribute air evenly
- Install backdraft damper to prevent reverse flow when fan is off

### Drainage and Water Table Issues

**Water Table Considerations:**
- Survey local water table depth before design
- Maintain minimum 0.5m clearance between pipes and seasonal high water table
- In high water table areas:
  - Use solid (non-perforated) HDPE pipe
  - Install French drain around pipe network
  - Consider shallower burial with enhanced insulation

**Slope and Drainage:**
- Slope all pipes minimum 1-2% toward drain point
- Install cleanout access at low points
- Provide sump or dry well for condensation drainage
- Gravel bed around pipes improves drainage

---

## Installation Guidelines

### General Requirements

1. **Site Assessment**
   - Survey soil type and water table depth
   - Verify frost depth for your specific location
   - Identify buried utilities before excavation
   - Assess excavation access and equipment requirements

2. **Excavation Planning**
   - Mark pipe layout before digging
   - Stockpile topsoil separately for backfill
   - Maintain consistent depth throughout
   - Create proper drainage slope

3. **Pipe Installation**
   - Lay pipes on 50-100mm gravel bed
   - Maintain consistent spacing between pipes
   - Avoid sharp bends (use 45-degree elbows maximum)
   - Test-fit all connections before final assembly

4. **Air Sealing**
   - All joints must be properly sealed
   - Use appropriate PVC cement or HDPE fusion
   - For corrugated HDPE, use rubber coupling bands
   - Pressure test system before backfilling

5. **Backfill and Compaction**
   - Backfill with native soil or imported sand/loam
   - Compact in 150-200mm lifts
   - Avoid heavy equipment directly over pipes
   - Restore topsoil and grade

6. **Electrical**
   - Install GFCI-protected circuit for fan(s)
   - Use outdoor-rated wiring and junction boxes
   - Consider variable speed controller for optimization
   - Install thermostat(s) for automatic control

### Recommended Tools and Equipment

**Excavation:**
- Mini excavator (1-2 ton) or trencher
- Laser level for consistent depth
- Measuring tape (minimum 30m)
- Survey stakes and string line

**Pipe Installation:**
- PVC cutter or reciprocating saw
- PVC primer and cement (for PVC)
- Heat fusion equipment (for HDPE, professional installation)
- Rubber coupling bands (for corrugated pipe)
- Pipe supports and hangers

**Electrical:**
- Wire strippers and crimpers
- Conduit bender
- Multimeter for testing
- Waterproof junction boxes

### DIY Feasibility Assessment

| Task | DIY Difficulty | Professional Recommended? |
|------|---------------|--------------------------|
| System design | Moderate | Consult for sizing |
| Excavation (small system) | Moderate | Rent equipment |
| Excavation (large system) | Difficult | Yes |
| Corrugated HDPE installation | Easy | No |
| Smooth HDPE fusion | Difficult | Yes |
| PVC installation | Easy-Moderate | No |
| Manifold fabrication | Moderate | Depends on skill |
| Electrical (fan, thermostat) | Moderate | If not comfortable |
| Backfill and compaction | Easy-Moderate | No |

**DIY-Friendly Materials:**
- Corrugated HDPE drainage pipe (press-fit connections)
- Schedule 40 PVC (solvent-welded)
- Pre-made manifold fittings (reduces fabrication)

---

## Installation Costs

### Material Costs (Estimated, CAD, 2025)

| Item | Unit Cost | Notes |
|------|-----------|-------|
| **Piping** | | |
| Corrugated HDPE 100mm | $3-5/m | Perforated or solid |
| Corrugated HDPE 150mm | $5-8/m | Double-wall preferred |
| Smooth PVC 100mm | $6-10/m | Schedule 40 |
| Smooth PVC 150mm | $10-15/m | Schedule 40 |
| ADS N-12 150mm | $12-18/m | Premium double-wall |
| **Fittings** | | |
| 90-degree elbow 150mm | $15-25 | Avoid if possible |
| 45-degree elbow 150mm | $12-20 | Preferred for turns |
| Tee fitting 150mm | $20-35 | For manifolds |
| Reducer coupling | $10-20 | Pipe to manifold |
| End cap | $8-15 | Per size |
| **Manifolds** | | |
| 200mm twin-wall (6m) | $80-120 | For multi-pipe systems |
| 250mm twin-wall (6m) | $100-150 | Larger systems |
| Custom fabricated | $150-400 | Professional welding |
| **Fans** | | |
| Inline fan 200 CFM | $80-150 | Basic residential |
| Inline fan 400 CFM | $150-250 | Medium systems |
| Inline fan 800 CFM | $250-400 | Large systems |
| Variable speed controller | $50-100 | Recommended |
| **Insulation** | | |
| EPS foam board R-5 | $15-25/m^2 | Perimeter insulation |
| EPS foam board R-10 | $25-40/m^2 | Enhanced insulation |
| **Controls** | | |
| Digital thermostat | $30-60 | Per setpoint |
| Dual-stage thermostat | $60-100 | Heating + cooling |
| Smart controller | $150-400 | Advanced automation |

### Excavation Costs

| Method | Cost (CAD) | Notes |
|--------|------------|-------|
| Hand digging (small system) | $0 (labor) | Very labor intensive |
| Mini excavator rental | $200-400/day | Plus delivery |
| Trencher rental | $150-300/day | For linear layouts |
| Professional excavation | $500-2000 | Depends on size/access |

### Total System Cost Estimates

| Greenhouse Size | System Type | Materials | Installation | Total |
|-----------------|-------------|-----------|--------------|-------|
| 25 m^2 (small hobby) | Multi-Pipe | $600-900 | $200-400 (DIY) | $800-1,300 |
| 50 m^2 (medium hobby) | Multi-Pipe | $900-1,400 | $300-600 (DIY) | $1,200-2,000 |
| 100 m^2 (large hobby) | Two-Trench | $1,400-2,200 | $800-1,500 | $2,200-3,700 |
| 200 m^2 (market garden) | Two-Trench | $2,500-4,000 | $1,500-3,000 | $4,000-7,000 |
| 100 m^2 (professional) | Under-Floor | $3,000-5,000 | $2,000-4,000 | $5,000-9,000 |

### Professional Installation

For professional Ceres GAHT® systems:
- System design consultation: $500-1,500
- GAHT® kit (materials): $2,000-8,000+ depending on size
- Professional installation: $3,000-10,000+
- Total turnkey: $6,000-20,000+

---

## Performance Data

### Case Study: Threefold Farm (Pennsylvania, Zone 6b/7a)

**System Specifications:**
- Greenhouse: 30' x 96' (275 m^2)
- Total tubing: 1,295m (4,250 ft) of 100mm corrugated perforated pipe
- Three separate battery zones at 0.6-2.4m depth
- Fans: 3 x 500W HAF fans (~2,000 CFM each effective)
- Energy consumption: ~1,100W total

**Performance Results:**
- Temperature maintained above 20F (-7C) in Zone 6b/7a winter
- Overwintering subtropical fruit trees successfully
- 60-80% reduction in heating costs vs propane

### Case Study: Penn State University Research (2018-2019)

**System Specifications:**
- 9m x 15m greenhouse with Ceres GAHT® system
- Climate Zone 5A (5,400-9,000 HDD)

**Measured Performance:**
- Daily heat storage: 128,588 BTU (37.7 kWh thermal)
- Average COP: 2.38 (continuously running)
- During coldest week (avg outside temp 20.5F/-6C): maintained 37F/3C setpoint
- System ran almost continuously during extreme cold

**Key Finding:** The climate battery maintained temperature even during extended cloudy periods when little solar input was available.

### Performance Expectations by Climate Zone

| Zone | Winter Low | Climate Battery Contribution | Backup Heat Needed |
|------|------------|-------------------|-------------------|
| Zone 4 | -10C | 70-90% | Minimal (occasional) |
| Zone 5 | -20C | 50-70% | Moderate |
| Zone 6 | -30C | 40-60% | Significant |
| Zone 7+ | -40C | 20-40% | Primary heating source |

### Energy Savings Analysis

**Comparison: Climate Battery vs Propane Heating (100 m² greenhouse)**

| Metric | Climate Battery | Propane Heat |
|--------|-------------|--------------|
| Capital cost | $3,000-5,000 | $1,000-2,000 |
| Annual operating cost | $200-400 | $1,500-3,000 |
| Annual energy (equivalent) | 500-1,000 kWh | 15,000-30,000 BTU/hr |
| 10-year total cost | $5,000-9,000 | $16,000-32,000 |
| Payback period | 2-4 years | N/A |

**Assumptions:**
- Electricity: $0.12/kWh
- Propane: $1.70/gallon
- Climate battery fan runs ~8 hours/day average
- Zone 5 climate with 180-day heating season

### Temperature Differential Achieved

**Typical Performance:**
- Summer cooling: 5-11C below ambient air temperature
- Winter heating: 10-20C above ambient (with solar gain)
- Overnight temperature stabilization: ±3-5°C vs 10-15°C swings without climate battery

**Example Data (Zone 5, mid-winter):**
- Outside air: -15C
- Greenhouse (no heat): -5C
- Greenhouse with climate battery: +5°C to +10°C (after charging)
- Soil temperature at 2m: +8C (stable)

---

## Potential Issues and Solutions

### Condensation and Drainage

**Problem:** Water accumulates in pipes from condensation, potentially blocking airflow or causing mold.

**Solutions:**
- Use perforated pipe to allow condensation to drain into soil
- Maintain minimum 1-2% slope toward drain point
- Install cleanout access at low points
- For solid pipe: provide sump or dry well at lowest point
- Run fans periodically even when not actively heating/cooling to dry system

**Perforated vs. Solid Pipe Debate:**

| Perforated Pipe | Solid Pipe |
|-----------------|------------|
| Condensation drains naturally | Requires engineered drainage |
| Releases latent heat to soil | All heat transfer via conduction |
| Potential for groundwater entry | Blocks groundwater infiltration |
| May allow radon entry | Better radon barrier |
| Works well in well-drained soil | Required for high water tables |

### Radon Mitigation

**Risk:** Climate battery systems can draw radon gas from soil into greenhouse air.

**Radon Levels in Canada:**
- Average indoor: 41 Bq/m^3
- Health Canada action level: 200 Bq/m^3
- Varies significantly by geology (higher in granite/uranium-bearing rock)

**Mitigation Strategies:**
1. **Use solid (non-perforated) pipe** - Most effective barrier
2. **Test before and after installation** - Radon test kits available at hardware stores
3. **Seal all pipe joints thoroughly**
4. **Install vapor barrier over greenhouse floor**
5. **Ensure adequate greenhouse ventilation**
6. **Consider sub-slab radon collection loop** (if building with concrete floor)

**Note:** Greenhouses are typically well-ventilated, which naturally dilutes radon. However, testing is recommended, especially in known high-radon areas (parts of Ontario, New Brunswick, Saskatchewan).

### Mold and Air Quality

**Concern:** Dark, moist underground pipes may harbor mold growth.

**Evidence from Practice:**
- Multiple operators report NO mold issues over years of operation
- Continuous airflow prevents mold establishment
- Soil microbe competition limits mold colonization
- Perforated pipes in contact with soil biota are self-regulating

**Best Practices:**
- Maintain adequate airflow (don't over-size pipes for fan capacity)
- Avoid stagnant periods - run fan periodically even off-season
- Use smooth-interior pipe when possible (less surface for growth)
- Install accessible cleanout points for inspection
- Consider UV light in return air stream for sensitive applications

### Rodent and Pest Prevention

**Problem:** Underground pipes and intakes can provide entry points for rodents and insects.

**Prevention Measures:**

| Location | Protection | Specification |
|----------|------------|---------------|
| Air intake | Heavy-gauge screen | 6mm (1/4") mesh, galvanized or stainless |
| Air outlet | Screen or damper | 6mm mesh |
| Pipe ends at manifold | Sealed joints | No gaps >3mm |
| Perimeter | Hardware cloth | Bury 150mm deep, bend outward |

**Additional Measures:**
- Inspect screens regularly and repair damage
- Use rodent bait stations around greenhouse perimeter
- Store no food or attractive materials near system
- Consider electronic rodent deterrents near intake

### Long-Term Performance Degradation

**Potential Issues Over Time:**

1. **Soil compaction around pipes**
   - Reduces air space and thermal exchange
   - Solution: Use gravel bed around pipes

2. **Uneven airflow distribution**
   - Some pipes receive more flow than others
   - Solution: Balance manifold design, use adjustable dampers

3. **Heat depletion during extended cloudy periods**
   - Soil "battery" can be drained in 3-5 days without solar input
   - Solution: Backup heat source, thermal mass integration

4. **Pipe degradation**
   - UV exposure at above-ground sections
   - Solution: Paint or cover exposed sections

5. **Fan motor wear**
   - Continuous operation stresses motors
   - Solution: Use quality fans rated for continuous duty, annual maintenance

**Maintenance Schedule:**

| Task | Frequency |
|------|-----------|
| Inspect intake/outlet screens | Monthly |
| Check fan operation | Monthly |
| Clean intake screen | Seasonally |
| Adjust thermostat setpoints | 2x per year (spring/fall) |
| Inspect accessible pipe sections | Annually |
| Test fan motor current draw | Annually |
| Replace fan motor | Every 5-10 years |

### Freeze Protection

**Risk:** In extreme cold, the cylinder of soil around pipes can freeze despite burial below frost line.

**Mitigation:**
- Bury pipes at recommended depth for climate zone
- Insulate upper portion of vertical risers
- Install insulation layer above pipe network (R-5 minimum)
- Avoid running system when outside air is below -20C (draws extreme cold underground)
- Install bypass damper to bypass climate battery in extreme cold

---

## Design Worksheets

### Quick Sizing Calculator

**Step 1: Determine Greenhouse Volume**
```
Volume = Length (m) x Width (m) x Average Height (m)
Example: 10m x 5m x 3m = 150 m^3
```

**Step 2: Calculate Required Airflow**
```
CFM = Volume (m^3) x ACH / 1.699
Where ACH = 5-10 (use 8 for typical design)
Example: 150 x 8 / 1.699 = 707 CFM
```

**Step 3: Determine Number of Pipes**
```
Pipe cross-sectional area = pi x (D/2)^2
For 150mm pipe: 0.0177 m^2

Air velocity target: 2 m/s = 118 m/min = 387 ft/min

CFM per pipe = Area (ft^2) x Velocity (ft/min)
For 150mm: 0.19 ft^2 x 387 fpm = 74 CFM per pipe

Number of pipes = Total CFM / CFM per pipe
Example: 707 / 74 = 9.5 -> use 10 pipes
```

**Step 4: Determine Pipe Length**
```
Optimal length: 8-10m per run
Maximum effective: 34m x (D/1m)
For 150mm: 34 x 0.15 = 5.1m (85% effectiveness point)
Practical: Use 8-10m runs
```

**Step 5: Calculate Burial Depth**
```
Minimum = Local frost depth + 0.3m
Optimal = Frost depth + 0.6m to 1.0m
Maximum practical = 3.0m

Example (Ottawa): 1.4m + 0.6m = 2.0m
```

**Step 6: Estimate Thermal Capacity**
```
Q_daily (kWh) = 0.00034 x CFM x delta_T (C) x hours/day
Example: 0.00034 x 707 x 15C x 8 hrs = 29 kWh/day
```

### Materials Checklist

**For 100 m^2 Greenhouse with Multi-Pipe Design:**

- [ ] Corrugated HDPE 150mm: 80m (8 pipes x 10m)
- [ ] Manifold pipe 250mm: 4m (2 x 2m sections)
- [ ] 45-degree elbows 150mm: 16
- [ ] Reducer couplings: 16
- [ ] End caps: 4
- [ ] Inline fan 400 CFM: 1
- [ ] Variable speed controller: 1
- [ ] Dual-stage thermostat: 1
- [ ] Electrical wire, conduit, boxes: as needed
- [ ] Gravel 20mm: 2 m^3
- [ ] EPS insulation R-5: 20 m^2
- [ ] Rodent screen 6mm mesh: 1 m^2

---

## Data Model

```go
type ClimateBatterySystem struct {
    ID          string
    ProjectID   string
    DesignType  ClimateBatteryDesignType  // "multi_pipe", "diagonal_array", "single_manifold", "under_floor", "hybrid", "custom"

    // Common fields
    NumPipes       int
    PipeDiameter   float64  // mm
    PipeLength     float64  // m (per pipe)
    BurialDepth    float64  // m
    PipeMaterial   string   // "pvc", "hdpe", "corrugated_hdpe", "double_wall_hdpe"
    PipePerforated bool     // true for perforated drainage pipe

    // Calculated fields
    TotalPipeLength  float64  // m
    SoilContactArea  float64  // m²
    ThermalCapacity  float64  // kWh/day (estimated)
    RecommendedFan   float64  // CFM
    AirVelocity      float64  // m/s

    // Climate/regional data
    ClimateZone      int      // 4-7+
    FrostDepth       float64  // m (local)
    GroundTemp       float64  // °C (undisturbed at depth)

    // Custom fields (for advanced option)
    CustomConfig    string   // JSON blob for custom designs

    // Cost tracking
    MaterialCost    float64  // CAD
    InstallCost     float64  // CAD
    EstimatedCost   float64  // CAD total

    CreatedAt       time.Time
    UpdatedAt       time.Time
}

type ClimateBatteryDesignType string

const (
    CBMultiPipe      ClimateBatteryDesignType = "multi_pipe"
    CBDiagonalArray  ClimateBatteryDesignType = "diagonal_array"
    CBSingleManifold ClimateBatteryDesignType = "single_manifold"
    CBUnderFloor     ClimateBatteryDesignType = "under_floor"
    CBHybrid         ClimateBatteryDesignType = "hybrid"
    CBCustom         ClimateBatteryDesignType = "custom"
    CBNone           ClimateBatteryDesignType = "none"
)
```

---

## API Endpoints

| Method | Path | Handler | Description |
|--------|------|---------|-------------|
| GET | `/api/climate-battery/designs` | `listClimateBatteryDesigns` | List available designs |
| GET | `/api/climate-battery/designs/{type}` | `getClimateBatteryDesign` | Get design details |
| POST | `/api/climate-battery/calculate` | `calculateClimateBattery` | Calculate thermal capacity |
| GET | `/api/climate-battery/recommendations` | `recommendClimateBattery` | Get recommendation for greenhouse size |
| GET | `/api/climate-battery/climate-data/{postal_code}` | `getClimateData` | Get local climate parameters |

---

## Component Structure

```text
internal/components/wizard/
├── step7_climate_battery.templ           # Main climate battery selection step
├── climate_battery_design_card.templ     # Individual design option card
├── climate_battery_diagram.templ         # SVG diagram component
├── climate_battery_specs_table.templ     # Specifications display
├── climate_battery_calculator.templ      # Custom design calculator
├── climate_battery_recommendation.templ  # AI-suggested design
└── climate_battery_climate_lookup.templ  # Regional climate data lookup
```

---

## References

### Primary Sources

- [Atmos Greenhouse Systems - Climate Battery Design & Education](https://atmosgreenhouse.com)
- [Atmos - Climate Battery Design Comparison (CFD Analysis)](https://atmosgreenhouse.com/blog/climate-battery-design-comparison)
- [Atmos - Climate Battery Blueprints](https://atmosgreenhouse.com/products-services)
- [Ceres Greenhouse Solutions - GAHT® System](https://ceresgs.com/gaht-system/) *(GAHT is a Ceres trademark)*
- [Ceres - 10 Dos and Don'ts for GAHT® Design](https://ceresgs.com/10-dos-and-donts-for-designing-a-ground-to-air-heat-transfer-system/)
- [Ceres - Setting Controls in a GAHT® System](https://ceresgs.com/setting-controls-in-a-gaht-system-or-climate-battery/)

### Research and Academic Sources

- [Appalachian State University - Climate Battery Greenhouse Energy Storage Thesis (Sinke, 2022)](https://libres.uncg.edu/ir/asu/f/Sinke_Leni_Spring%202022_Thesis.pdf)
- [Natural Resources Canada - Earth to Air Thermal Exchanger (EATEX) Design Tool](https://natural-resources.canada.ca/sites/nrcan/files/canmetenergy/pdf/ENG_EATEX_Design_Principles_and_Concept_Design_Tool.pdf)
- [NRC - Ground Temperatures (CBD-180)](http://web.mit.edu/parmstr/Public/NRCan/CanBldgDigests/cbd180_e.html)

### Case Studies and Practical Guides

- [Threefold Farm - Climate Battery Greenhouse](https://threefold.farm/climate-battery-greenhouse)
- [Threefold Farm - Penn State Climate Battery Study 2018-2019](https://www.threefold.farm/psu-climate-battery-study-2018-2019)
- [Growing Spaces - How to Build a Climate Battery Greenhouse](https://growingspaces.com/blog/climate-battery-greenhouse/)
- [One Community Global - Open Source Climate Battery Research Hub](https://onecommunityglobal.org/climate-battery/)
- [Mother Earth News - Using a Climate Battery System](https://www.motherearthnews.com/organic-gardening/a-self-heating-greenhouse-zbcz1707/)

### Deep Winter Greenhouse (UMN Extension)

- [UMN Extension - Deep Winter Greenhouses](https://extension.umn.edu/growing-systems/deep-winter-greenhouses)
- [UMN Extension - Deep Winter Greenhouse Technical Guides](https://extension.umn.edu/growing-systems/deep-winter-greenhouse-resources)
- [Northlands Winter Greenhouse Manual (PDF)](https://extension.umn.edu/sites/extension.umn.edu/files/2022-11/northlands-winter-greenhouse-manual.pdf)
- [UMN College of Design - CSBR Deep Winter Greenhouse Project](https://design.umn.edu/center-sustainable-building-research/projects/deep-winter-greenhouse)
- [Regional Sustainable Development Partnerships - DWG Initiative](https://rsdp.umn.edu/projects/deep-winter-greenhouses)

### Technical Design Resources

- [Home in the Earth - Earth Tube Design](https://www.homeintheearth.com/tech_notes/earth-tubes/earth-tube-design-for-earth-sheltered-homes/)
- [BuildItSolar - Ground Temperatures by Location](https://www.builditsolar.com/Projects/Cooling/EarthTemperatures.htm)
- [BuildItSolar - Earth Tube Design Notes](https://www.builditsolar.com/Projects/Cooling/EarthtubeNotes.htm)

### Canadian Climate Data

- [Environment and Climate Change Canada - Historical Climate Data](https://climate.weather.gc.ca/)
- [NRC - Frost Penetration Studies in Canada](https://nrc-publications.canada.ca/eng/view/ft/?id=02677feb-3262-42fc-9bec-da4bf193fe0d)
- [Natural Resources Canada - Undisturbed Ground Temperatures Map](https://geoappext.nrcan.gc.ca/arcgis/rest/services/Energy/clean_energy_wind_potential/MapServer/3)

### Soil Properties

- [Oklahoma State - Soil Thermal Properties](https://open.library.okstate.edu/rainorshine/chapter/13-2-soil-thermal-properties/)
- [Canadian Geotechnical Journal - Soil Thermal Conductivity Model (Cote & Konrad, 2005)](https://cdnsciencepub.com/doi/10.1139/t05-017)

### Northern/Permafrost Considerations

- [Natural Resources Canada - Permafrost Thaw in Northern Communities](https://natural-resources.canada.ca/stories/simply-science/permafrost-thaw-brings-major-problems-canada-s-northern-arctic-communities)
- [Canadian Journal of Earth Sciences - Climate and Ground Temperature in Northern Canada](https://cdnsciencepub.com/doi/10.1139/e11-075)
