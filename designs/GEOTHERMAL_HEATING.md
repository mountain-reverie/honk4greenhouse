# Low-Cost Geothermal Heating Systems for Greenhouses

This document provides comprehensive technical guidance for designing, sizing, and installing cost-effective geothermal heating systems for greenhouses in Canadian climates. It covers both traditional ground-source heat pump (GSHP) systems and simpler climate battery (GAHT) alternatives, with emphasis on DIY-friendly approaches and cost optimization.

## Table of Contents

1. [Overview](#overview)
2. [System Types](#system-types)
3. [Ground Temperatures in Canada](#ground-temperatures-in-canada)
4. [Greenhouse Heating Load Calculations](#greenhouse-heating-load-calculations)
5. [Ground Loop Design](#ground-loop-design)
6. [Enhanced Grout Materials](#enhanced-grout-materials)
7. [Heat Pump Selection](#heat-pump-selection)
8. [Climate Battery (GAHT) Systems](#climate-battery-gaht-systems)
9. [Root Zone Heating Integration](#root-zone-heating-integration)
10. [DIY Installation Guide](#diy-installation-guide)
11. [Cost Analysis](#cost-analysis)
12. [Canadian Regulatory Requirements](#canadian-regulatory-requirements)
13. [System Sizing Examples](#system-sizing-examples)
14. [Maintenance and Troubleshooting](#maintenance-and-troubleshooting)
15. [References](#references)

---

## Overview

Geothermal heating leverages the stable temperatures found below the frost line to provide efficient heating (and cooling) for greenhouses. The ground acts as a thermal reservoir, maintaining relatively constant temperatures year-round while air temperatures swing dramatically with seasons.

### Why Geothermal for Greenhouses?

| Advantage | Description |
|-----------|-------------|
| **High Efficiency** | COP of 3-5 vs. 0.8-0.95 for combustion heating |
| **Stable Performance** | Ground temperature is constant unlike air-source systems |
| **Dual Function** | Same system provides heating and cooling |
| **Low Operating Cost** | 50-70% reduction in heating energy costs |
| **Long Lifespan** | Ground loops last 50+ years, heat pumps 20-25 years |
| **No Combustion** | No CO₂ emissions in greenhouse, no fuel storage |

### System Options by Budget

| System Type | Relative Cost | DIY Feasibility | Best For |
|-------------|---------------|-----------------|----------|
| Climate Battery (GAHT) | $ | Excellent | Small greenhouses, passive heating |
| Horizontal Closed Loop | $$ | Good | Large properties, shallow bedrock |
| Pond/Lake Loop | $$ | Good | Properties with water bodies |
| Vertical Borehole | $$$ | Limited | Limited space, any soil type |
| Open Loop (Well) | $$-$$$ | Moderate | High water table, good aquifer |

---

## System Types

### Closed Loop Systems

**Vertical Borehole:**
- Boreholes drilled 45-150 m (150-500 ft) deep
- HDPE U-tubes inserted and grouted
- Spacing: 4.5-6 m (15-20 ft) between boreholes
- Best for: Limited land area, any soil conditions
- Cost driver: Drilling at $15-50/ft ($50-165/m)

**Horizontal Loop:**
- Trenches dug 1.2-1.8 m (4-6 ft) deep
- Pipe laid in straight runs, slinky coils, or multiple layers
- Requires large land area: ~185 m² per ton (2000 ft²/ton)
- Best for: Large properties, DIY installation
- Cost driver: Trenching/excavation

**Slinky Configuration:**
- Pipe coiled in overlapping loops
- Reduces trench length by 60-70%
- Requires ~55-75 m² per ton (600-800 ft²/ton)
- Slightly less efficient than straight runs

### Open Loop Systems

**Standing Column Well:**
- Single deep well (150-450 m)
- Water drawn from bottom, returned to top
- High efficiency but limited to suitable geology
- Requires adequate aquifer yield

**Pump and Dump:**
- Water from supply well, discharged to second well or surface
- Simplest design but regulatory restrictions common
- Flow requirement: 5.7 L/min per ton (1.5 gpm/ton)

### Climate Battery / GAHT Systems

Ground to Air Heat Transfer (GAHT) systems are simpler, lower-cost alternatives:

- Buried perforated pipes circulate air, not water
- No heat pump required—fans only
- Stores daytime heat for nighttime release
- Best for: Passive/low-energy greenhouses, mild climates
- Limitation: Cannot "pump" heat, only store and release

---

## Ground Temperatures in Canada

Ground temperature is the foundation of geothermal design. Below the frost line, temperatures remain remarkably stable year-round.

### Annual Ground Temperature by Depth

| Depth | Temperature Variation | Notes |
|-------|----------------------|-------|
| Surface | Tracks air temperature | Daily/seasonal swings |
| 0.5 m | ±10-15°C seasonal | Frost zone in most of Canada |
| 1.0 m | ±5-8°C seasonal | Below frost in southern regions |
| 2.0 m | ±2-3°C seasonal | Stable in most regions |
| 3.0 m+ | ±1°C or less | Near-constant temperature |
| 10+ m | Constant | Equals mean annual air temp + ~1-2°C |

### Regional Deep Ground Temperatures

| Region | Mean Annual Air Temp | Deep Ground Temp (>3m) | Frost Depth |
|--------|---------------------|------------------------|-------------|
| Victoria, BC | 10.4°C | 11-12°C | 0.3 m |
| Vancouver, BC | 10.4°C | 11-12°C | 0.5 m |
| Toronto, ON | 9.4°C | 10-11°C | 1.2 m |
| Ottawa, ON | 6.6°C | 8-9°C | 1.5 m |
| Montreal, QC | 7.4°C | 8-9°C | 1.5 m |
| Calgary, AB | 4.4°C | 6-7°C | 1.8 m |
| Edmonton, AB | 2.9°C | 5-6°C | 2.1 m |
| Winnipeg, MB | 3.0°C | 5-6°C | 2.4 m |
| Saskatoon, SK | 2.8°C | 5-6°C | 2.3 m |
| Yellowknife, NT | -4.3°C | 0-2°C* | Permafrost |

*Permafrost zones require specialized design considerations.

### Soil Type Impact on Heat Transfer

| Soil Type | Thermal Conductivity (W/m·K) | Relative Performance |
|-----------|------------------------------|---------------------|
| Saturated sand/gravel | 2.4-2.9 | Excellent |
| Saturated clay | 1.4-1.8 | Good |
| Wet shale | 1.4-2.1 | Good |
| Granite | 2.5-3.8 | Excellent |
| Dry sand | 0.3-0.4 | Poor |
| Dry clay | 0.4-0.6 | Poor |
| Limestone | 2.2-2.8 | Very Good |

**Key Insight:** Soil moisture is critical. Dry soils have devastating impact on heat transfer—moisture content below 12.5% significantly degrades performance.

---

## Greenhouse Heating Load Calculations

Accurate heating load calculation is essential for proper system sizing. Oversizing wastes capital; undersizing leads to inadequate heating.

### Basic Heat Loss Formula

```
Q_total = Q_conduction + Q_infiltration
```

**Conduction Heat Loss:**
```
Q_c = U × A × ΔT

Where:
Q_c = Heat loss (W)
U = Overall heat transfer coefficient (W/m²·K)
A = Surface area (m²)
ΔT = Temperature difference (inside - outside) (K or °C)
```

**Infiltration Heat Loss:**
```
Q_i = 0.33 × V × ACH × ΔT

Where:
Q_i = Infiltration heat loss (W)
V = Greenhouse volume (m³)
ACH = Air changes per hour
0.33 = Volumetric heat capacity of air (W·h/m³·K)
```

### U-Values for Greenhouse Glazing

| Material | U-Value (W/m²·K) | Notes |
|----------|------------------|-------|
| Single glass | 6.2 | Poor insulation |
| Double glass | 3.5 | Standard commercial |
| Single polyethylene | 6.8 | Short lifespan |
| Double poly (inflated) | 4.0 | Common, economical |
| Twin-wall polycarbonate (8mm) | 3.3 | Good balance |
| Triple-wall polycarbonate (16mm) | 2.4 | Better insulation |
| Five-wall polycarbonate (32mm) | 1.1 | Best insulation |

### Air Change Rates (ACH)

| Construction Quality | ACH |
|---------------------|-----|
| New, tight construction | 0.5-1.0 |
| Average construction | 1.0-1.5 |
| Old or loose construction | 2.0-4.0 |
| Single-layer plastic | 1.0-1.5 |
| Double-layer plastic | 0.5-1.0 |
| Glass with good seals | 0.75-1.0 |

### Design Temperature Selection

Use the **2.5% winter design temperature** for heating system sizing—this is the temperature exceeded 97.5% of winter hours.

| City | 2.5% Design Temp (°C) |
|------|----------------------|
| Vancouver | -7 |
| Toronto | -18 |
| Ottawa | -24 |
| Montreal | -23 |
| Calgary | -27 |
| Edmonton | -31 |
| Winnipeg | -33 |

### Worked Example: Greenhouse Heat Load

**Greenhouse specifications:**
- Dimensions: 9 m × 15 m (30 × 50 ft) floor area = 135 m²
- Wall height: 2.4 m, peaked roof to 3.6 m
- Glazing: Triple-wall polycarbonate (U = 2.4 W/m²·K)
- North wall: Insulated (U = 0.5 W/m²·K)
- Location: Ottawa (design temp = -24°C)
- Inside temp: 15°C (nighttime minimum for tomatoes)

**Surface areas:**
- South glazing (roof + wall): ~185 m²
- East/West walls: ~40 m² each = 80 m²
- North insulated wall: ~45 m²
- Total glazed area: 265 m²
- Volume: ~405 m³

**Conduction losses:**
```
Q_glazing = 2.4 × 265 × (15 - (-24)) = 2.4 × 265 × 39 = 24,804 W
Q_north = 0.5 × 45 × 39 = 878 W
Q_conduction = 25,682 W
```

**Infiltration losses (ACH = 1.0):**
```
Q_infiltration = 0.33 × 405 × 1.0 × 39 = 5,213 W
```

**Total heat loss:**
```
Q_total = 25,682 + 5,213 = 30,895 W ≈ 31 kW (106,000 BTU/hr)
```

**Required capacity:** 31 kW ÷ 3.517 = **8.8 tons**

---

## Ground Loop Design

### Vertical Borehole Sizing

**Rule of Thumb (Northern/Heating-Dominated):**
- 45-60 m of borehole per ton (150-200 ft/ton)
- Double U-tube configuration recommended
- Minimum spacing: 4.5 m (15 ft) between boreholes

**More Precise Calculation:**

```
L = (Q × R_total) / (T_ground - T_fluid_min)

Where:
L = Required borehole length (m)
Q = Heat extraction rate (W)
R_total = Total thermal resistance (m·K/W)
T_ground = Undisturbed ground temperature (°C)
T_fluid_min = Minimum allowable fluid temperature (°C)
```

**Design constraints:**
- Minimum entering water temperature (EWT): -1°C (30°F)
- Maximum EWT: 32°C (90°F)
- Fluid: Water with antifreeze (propylene glycol or methanol)

### Horizontal Loop Sizing

**Rule of Thumb:**
- Straight runs: 120-150 m per ton (400-500 ft/ton)
- Slinky loops: 55-75 m of trench per ton (180-250 ft/ton)
- Burial depth: 1.2-1.8 m (4-6 ft) minimum

**Slinky Configuration:**
- Loop diameter: 0.9-1.2 m (3-4 ft)
- Pitch (center-to-center): 0.3-0.45 m (12-18 in)
- Trench width: 0.6-0.9 m (2-3 ft)

### Pipe Specifications

**Standard Material: HDPE PE100 SDR11**

| Nominal Size | OD (mm) | Wall (mm) | Pressure Rating |
|--------------|---------|-----------|-----------------|
| 25 mm (1") | 32 | 2.9 | 16 bar (PN16) |
| 32 mm (1¼") | 40 | 3.7 | 16 bar |
| 40 mm (1½") | 50 | 4.6 | 16 bar |

**PE100 Properties:**
- Thermal conductivity: 0.4 W/m·K
- Temperature range: -40°C to +60°C
- Design life: 50+ years
- Joining: Heat fusion (socket or butt)

**Flow Rate Design:**
- Target: 3 L/min per kW (3 gpm/ton)
- Velocity: 0.6-1.2 m/s in pipes
- Turbulent flow required for heat transfer

---

## Enhanced Grout Materials

Enhanced backfill materials can reduce required borehole depths by **20-37%** and significantly improve system efficiency.

### Thermal Conductivity Comparison

| Material | Thermal Conductivity (W/m·K) | Cost Index |
|----------|------------------------------|------------|
| Plain bentonite | 0.7-0.9 | 1.0× |
| Bentonite + sand | 1.0-1.7 | 1.2× |
| Bentonite + 5% graphite | 1.8-2.0 | 1.5× |
| Bentonite + 10% graphite | 2.5-3.0 | 2.0× |
| Commercial thermal grout | 2.0-3.3 | 2.5× |
| Brookhaven Mix 111 | 2.0+ | 1.0× |

### Graphite Enhancement

Natural graphite flakes are the most effective thermal enhancer:

- **5% by weight**: Best cost-benefit ratio (1.8-2.0 W/m·K)
- **10% by weight**: Near-maximum benefit (2.5-3.0 W/m·K)
- **Optimal particle size**: 1-5 mm diameter flakes
- **Avoid**: Fine powder <44 microns (increases viscosity, reduces conductivity)

**Sourcing:**
- Industrial bulk: $0.30-0.70/kg ($0.13-0.32/lb)
- Small quantities: $7-18/kg ($3-8/lb)
- Geothermal-specific: GeoPro PowerTEC, CETCO TC Booster

### Brookhaven Mix 111 (DIY Formula)

Open-source cement-sand grout from DOE-funded research:

```
Per batch (yields 72 L / 19.1 US gallons):
- Portland cement (Type I): 43 kg (94 lbs)
- Silica sand (100% passing #8, 60-80% passing #30): 91 kg (200 lbs)
- Water: 23.4 L (6.19 US gallons)
- Superplasticizer (Type F, ASTM C494): 620 mL (21 fl oz)
- Optional: Bentonite 0.47 kg (1.04 lbs) for shrinkage control

Specific gravity: 2.18
Thermal conductivity: 2.0+ W/m·K
Material cost: ~$1.60-2.10 per gallon
```

**Critical:** Superplasticizer is essential for pumpability.

### Mixing Procedure for Bentonite-Graphite Grout

1. Add measured water to mixer (15-17 gallons for graphite mixes)
2. Add bentonite powder at low speed
3. **Allow full 30-minute hydration** (critical step)
4. Add graphite flakes after bentonite is fully hydrated
5. Mix until homogeneous (2-3 minutes)
6. Pump immediately—working time is 5-15 minutes

**Equipment:**
- Positive displacement pump (gear or progressive cavity)
- Tremie pipe: 25-38 mm (1-1.5 in) diameter
- Flow rate: 20-60 L/min (5-15 gpm)
- Always grout from bottom to top

### Materials to Avoid

**Biochar:** Decreases thermal conductivity by 25-60%. Its porous structure traps air (thermal conductivity 0.024 W/m·K), creating thermal barriers.

**Carbon black:** Thermal conductivity <0.122 W/m·K—far inferior to graphite.

---

## Heat Pump Selection

### Types for Greenhouse Applications

| Type | Output | Best Application |
|------|--------|------------------|
| Water-to-Air | Forced air | Simple installations, ductwork |
| Water-to-Water | Hot water | Radiant floor, root zone, hydronic |
| Split System | Variable | Flexibility, multiple zones |

**Recommendation for Greenhouses:** Water-to-water heat pumps with hydronic distribution (radiant floor or root zone heating) provide the most efficient and plant-friendly heating.

### Performance Metrics

**Coefficient of Performance (COP):**
```
COP = Heat Output (kW) / Electrical Input (kW)
```

| Condition | Typical COP | Notes |
|-----------|-------------|-------|
| Ideal (EWT 10°C, LWT 35°C) | 4.5-5.5 | Manufacturer ratings |
| Good (EWT 5°C, LWT 40°C) | 3.5-4.5 | Most Canadian installations |
| Cold ground (EWT 0°C, LWT 45°C) | 2.8-3.5 | Northern Canada |
| Cold climate, high temp | 2.5-3.0 | Worst case |

**Energy Efficiency Ratio (EER):** Cooling performance
```
EER = Cooling Output (BTU/hr) / Electrical Input (W)
```

### Sizing Guidelines

**Rule of thumb:** 1 ton = 12,000 BTU/hr = 3.517 kW

For the Ottawa greenhouse example (31 kW heat load):
- Required: 31 kW / 3.517 = 8.8 tons
- With COP of 3.5: Electrical input = 31 / 3.5 = 8.9 kW

### Recommended Units for DIY

| Manufacturer | Model Series | Sizes | Notes |
|--------------|--------------|-------|-------|
| ClimateMaster | Tranquility | 2-6 ton | Popular DIY choice |
| WaterFurnace | 5 Series | 2-6 ton | High efficiency |
| Geocool | Eco Series | 2-5 ton | Budget-friendly |
| Arctic |DERA | 2-6 ton | Cold climate optimized |

**DIY Purchase Options:**
- [Ingrams Water & Air](https://iwae.com/) - Installation kits, wholesale pricing
- [123 Zero Energy](https://www.123zeroenergy.com/) - Design support
- [Alpine Home Air](https://www.alpinehomeair.com/) - Wide selection

---

## Climate Battery (GAHT) Systems

Climate Battery or Ground to Air Heat Transfer (GAHT) systems offer a simpler, lower-cost alternative to traditional geothermal heat pumps.

### How It Works

1. **Daytime (Cooling Mode):** Hot, humid greenhouse air is pushed through buried pipes
2. **Heat Transfer:** Air cools to ground temperature, moisture condenses (releasing latent heat)
3. **Thermal Storage:** Heat is stored in soil mass beneath greenhouse
4. **Nighttime (Heating Mode):** Cool greenhouse air circulates through warmed soil
5. **Heat Recovery:** Soil releases stored heat, warming return air

### Key Design Parameters

| Parameter | Recommended Value | Notes |
|-----------|-------------------|-------|
| Pipe depth | 0.9-1.2 m (3-4 ft) | Below frost line |
| Pipe diameter | 100-150 mm (4-6 in) | Perforated drainage pipe |
| Run length | 7.5-10.5 m (25-35 ft) | Optimal heat exchange |
| Pipe spacing | 0.45-0.6 m (18-24 in) | Prevents thermal interference |
| Total length | 60-90 m per 100 m² | Rule of thumb |
| Air velocity | 1-2 m/s (200-400 fpm) | Allows heat transfer |

### Installation Layout

```
                    Greenhouse Floor
    ═══════════════════════════════════════════
          ┌─────────────────────────────┐
    FAN → │  Plenum (supply manifold)   │
          └─────┬───┬───┬───┬───┬───────┘
                │   │   │   │   │
         ┌──────┴───┴───┴───┴───┴──────┐
         │    Buried perforated pipes   │  ← 0.9-1.2m deep
         │    in gravel bed with        │
         │    drainage to sump          │
         └──────┬───┬───┬───┬───┬──────┘
                │   │   │   │   │
          ┌─────┴───┴───┴───┴───┴───────┐
    ← FAN │  Return plenum (collector)  │
          └─────────────────────────────┘
    ═══════════════════════════════════════════
                 Insulated Perimeter
```

### Components

**Pipe:** 100 mm (4") perforated corrugated drainage pipe
- Cost: ~$0.50-1.00/m ($0.15-0.30/ft)
- Perforations allow condensate drainage

**Gravel Bed:** 19-38 mm (¾-1½") washed gravel
- Provides air channels around pipes
- Stores additional thermal mass

**Fans:** In-line duct fans
- Sizing: 2-4 m³/min per m² of greenhouse (6-12 CFM/ft²)
- Two fans recommended (supply and return)
- Energy use: 100-300 W total for small greenhouse

**Controls:** Dual thermostat system
- Heating thermostat: Activates when greenhouse too cold
- Cooling thermostat: Activates when greenhouse too hot
- Dead band between setpoints (e.g., 15-25°C)

### Performance Expectations

| Climate | Winter Temp Boost | Summer Temp Reduction |
|---------|-------------------|----------------------|
| Mild (Zone 7-8) | 5-10°C | 5-8°C |
| Moderate (Zone 5-6) | 3-7°C | 5-8°C |
| Cold (Zone 3-4) | 2-5°C | 5-8°C |

**Limitations:**
- Cannot provide active heating below ground temperature
- Insufficient as sole heat source in cold Canadian climates
- Best used to reduce (not eliminate) supplemental heating needs

### Cost Comparison

| Component | GAHT System | Full GSHP System |
|-----------|-------------|------------------|
| Ground work | $500-1,500 | $5,000-15,000 |
| Pipes/materials | $300-800 | $2,000-5,000 |
| Fans/controls | $200-500 | N/A |
| Heat pump | N/A | $3,000-8,000 |
| Total | $1,000-2,800 | $10,000-28,000 |

---

## Root Zone Heating Integration

Root zone heating warms the growing medium directly rather than the air, providing significant energy savings and better plant growth.

### Benefits

- Air temperature can be 3-6°C (5-10°F) lower while maintaining plant health
- Reduces heat loss through glazing (smaller ΔT)
- Energy savings of 25-40%
- Improved root development and plant vigor
- Faster crop production

### System Design

**Water Temperature:** 35-40°C (95-104°F) for root zone
- Much lower than radiator systems (60-80°C)
- Ideal for heat pump efficiency (higher COP)

**Pipe Layout:**
- Spacing: 15-30 cm (6-12 in) between pipes
- Depth: 5-10 cm (2-4 in) below growing surface
- Material: EPDM rubber, PEX, or polyethylene

**Heat Output:** 80-120 W/m² of bench area

### Integration with Geothermal

Water-to-water geothermal heat pumps are ideal for root zone heating:

1. Heat pump produces 35-45°C water efficiently
2. Circulating pump moves water through floor/bench loops
3. Return water at 30-38°C goes back to heat pump
4. Low temperature differential = high COP (4-5+)

**System schematic:**
```
Ground Loop → Heat Pump → Buffer Tank → Circulator → Root Zone Pipes
     ↑                                                      │
     └──────────────────────────────────────────────────────┘
```

---

## DIY Installation Guide

### Horizontal Loop Installation Steps

**1. Site Assessment**
- Soil test to determine type and moisture
- Locate underground utilities (call before you dig)
- Verify adequate area (55-150 m² per ton)
- Check frost depth for your region

**2. Trenching**
- Depth: 1.2-1.8 m (below frost line + 0.3 m)
- Width: 0.6-0.9 m for single pipe, 0.9-1.2 m for slinky
- Length per ton: 40-50 m (slinky) or 120-150 m (straight)

**3. Pipe Installation**
- Lay out HDPE pipe, avoiding kinks
- For slinky: form loops and stake in place
- Fuse all joints (socket or butt fusion)
- Pressure test to 150% of operating pressure

**4. Backfill**
- Layer 15-30 cm sand below and above pipes
- Maintain moisture during installation
- Avoid large rocks that could damage pipe
- Compact in 30 cm lifts

**5. Header Installation**
- Connect individual circuits at manifold
- Install isolation valves for each loop
- Balance flow with adjustable valves

### Vertical Borehole (Professional Recommended)

Vertical boreholes typically require professional drilling, but homeowner can:
- Prepare site access for drill rig
- Excavate header trench
- Install horizontal headers and manifold
- Connect to heat pump

**DIY Savings:** $2,000-5,000 by doing own site work and headers

### Piping Connections

**Heat Fusion (Recommended):**
- Socket fusion for small diameters (<63 mm)
- Butt fusion for larger pipes
- Requires fusion machine rental ($100-200/day)
- Creates joints as strong as pipe

**Mechanical Fittings (Alternative):**
- Compression fittings acceptable for indoor portions
- Not recommended for buried sections
- Easier but less reliable long-term

### Heat Pump Installation

1. Mount unit on vibration isolators
2. Connect ground loop (supply and return)
3. Connect load side (to radiant system or air handler)
4. Install circulation pumps (sized per manufacturer)
5. Wire electrical (requires licensed electrician)
6. Charge system, purge air
7. Commission and test

---

## Cost Analysis

### Capital Costs (CAD, 2025 estimates)

**Small Greenhouse (50 m², ~4 kW load):**

| Component | Horizontal | Vertical | GAHT |
|-----------|------------|----------|------|
| Ground loop/excavation | $3,000-5,000 | $8,000-12,000 | $800-1,500 |
| Pipe and fittings | $800-1,200 | $1,500-2,500 | $300-600 |
| Heat pump (2-ton) | $3,500-5,000 | $3,500-5,000 | N/A |
| Distribution system | $1,500-2,500 | $1,500-2,500 | $200-400 |
| Electrical/controls | $500-1,000 | $500-1,000 | $200-300 |
| **Total DIY** | **$9,300-14,700** | **$15,000-23,000** | **$1,500-2,800** |
| **Total Professional** | **$15,000-25,000** | **$25,000-40,000** | **$3,000-5,000** |

**Medium Greenhouse (135 m², ~31 kW load - Ottawa example):**

| Component | Horizontal | Vertical |
|-----------|------------|----------|
| Ground loop (9 tons) | $12,000-20,000 | $35,000-55,000 |
| Heat pump (10-ton) | $8,000-12,000 | $8,000-12,000 |
| Distribution system | $4,000-6,000 | $4,000-6,000 |
| Controls and electrical | $2,000-3,000 | $2,000-3,000 |
| **Total DIY** | **$26,000-41,000** | **$49,000-76,000** |

### Operating Costs

**Annual electricity for heat pump:**
```
Annual heating energy = Heating load × Heating hours / COP
Example (Ottawa, 31 kW system):
- Heating degree hours: ~100,000°C·hr (approximate)
- Annual heat energy: ~45,000 kWh
- With COP 3.5: Electricity = 45,000 / 3.5 = 12,857 kWh
- At $0.12/kWh: $1,543/year
```

**Comparison to alternatives:**
| Heating Method | Annual Cost (31 kW load) |
|----------------|--------------------------|
| Geothermal (COP 3.5) | $1,500-1,800 |
| Propane | $4,500-6,000 |
| Natural gas | $2,500-3,500 |
| Electric resistance | $5,400-6,500 |
| Wood | $1,500-2,500 |

### Return on Investment

**Simple payback vs. propane heating:**
```
Annual savings: $4,500 - $1,600 = $2,900
DIY horizontal system cost: $26,000
Payback: 26,000 / 2,900 = 9.0 years
```

**With incentives (Canada Greener Homes Grant, provincial programs):**
```
System cost after $5,000 rebate: $21,000
Payback: 21,000 / 2,900 = 7.2 years
```

---

## Canadian Regulatory Requirements

### Federal

- Heat pumps must meet minimum efficiency under Canada's Energy Efficiency Regulations
- No federal permits for ground loops

### Provincial (Ontario Example)

**Closed Loop Vertical:**
- Ontario Regulation 98/22 (Geothermal) requires:
  - Licensed geothermal installer
  - Proper grouting to prevent aquifer contamination
  - Well record filing

**Closed Loop Horizontal:**
- No specific regulation (treated as utility trenching)
- Call Before You Dig required
- Municipal building permit may be required

**Open Loop:**
- Wells Regulation (O. Reg. 903)
- Ontario Water Resources Act compliance
- Sewage Works Environmental Compliance Approval required
- Discharge permits for surface disposal

### Provincial Variations

| Province | Closed Vertical | Open Loop | Notes |
|----------|-----------------|-----------|-------|
| BC | Licensed driller | Heavily restricted | Water rights issues |
| AB | Licensed driller | Permit required | Energy Regulator oversight |
| SK | Licensed driller | Permit required | |
| MB | Licensed driller | Permit required | |
| ON | Licensed installer | Complex permits | Reg. 98/22 |
| QC | Licensed driller | Permit required | Environmental review |
| Atlantic | Varies | Generally restricted | Limited experience |

### Best Practices for Compliance

1. Consult local building department before starting
2. Use licensed contractors for regulated work (drilling, electrical)
3. File required well records
4. Maintain as-built documentation
5. Use approved antifreeze (propylene glycol preferred)

---

## System Sizing Examples

### Example 1: Small Hobby Greenhouse (Vancouver)

**Specifications:**
- Size: 4 m × 6 m = 24 m² floor
- Glazing: Double-wall polycarbonate (U = 3.3 W/m²·K)
- Design temp: -7°C outside, 10°C inside
- Surface area: ~75 m²

**Heat load:**
```
Q = 3.3 × 75 × 17 + infiltration ≈ 4,200 W + 500 W = 4.7 kW
```

**System recommendation:** Climate battery (GAHT)
- Ground stable at 11°C can boost night temps 5-8°C
- Supplemental heat only on coldest nights
- Cost: $1,500-2,500 DIY

### Example 2: Market Garden Greenhouse (Calgary)

**Specifications:**
- Size: 9 m × 30 m = 270 m² floor
- Glazing: Triple-wall polycarbonate (U = 2.4 W/m²·K)
- Insulated north wall, perimeter insulation
- Design temp: -27°C outside, 15°C inside
- Surface area: ~550 m²

**Heat load:**
```
Q = 2.4 × 500 × 42 + 0.5 × 50 × 42 + infiltration ≈ 50,400 + 1,050 + 8,000 = 59.5 kW
```

**Required:** 17 tons

**System recommendation:** Horizontal closed loop
- Required trench: ~750-1,000 m of slinky
- Area needed: ~1,000-1,200 m² (~¼ acre)
- Two 10-ton heat pumps or single large unit
- Root zone heating for efficiency
- Cost: $55,000-80,000 DIY

### Example 3: Year-Round Production (Zone 5)

**Specifications:**
- Size: 6 m × 12 m = 72 m² floor
- Glazing: 5-wall polycarbonate (U = 1.1 W/m²·K) south
- Super-insulated north/end walls (U = 0.3 W/m²·K)
- Thermal mass: 2,000 L water barrels
- Design: -20°C outside, 18°C inside
- Surface area: Glazed 120 m², insulated 60 m²

**Heat load:**
```
Q_glazed = 1.1 × 120 × 38 = 5,016 W
Q_insulated = 0.3 × 60 × 38 = 684 W
Q_infiltration = 0.33 × 220 × 0.5 × 38 = 1,379 W
Q_total = 7,079 W ≈ 7 kW (2 tons)
```

**System recommendation:** Small vertical or horizontal + GAHT
- Vertical: 2 boreholes × 60 m each
- Or horizontal: ~120 m slinky trench
- GAHT for summer cooling and shoulder season
- 2-ton water-to-water heat pump
- Root zone heating in raised beds
- Cost: $12,000-18,000 DIY vertical, $8,000-12,000 DIY horizontal

---

## Maintenance and Troubleshooting

### Annual Maintenance Checklist

**Ground Loop:**
- [ ] Check pressure (should be stable)
- [ ] Verify antifreeze concentration
- [ ] Inspect exposed piping for leaks
- [ ] Clean strainers/filters

**Heat Pump:**
- [ ] Clean air filters (if water-to-air)
- [ ] Check refrigerant pressures
- [ ] Inspect electrical connections
- [ ] Verify thermostat operation
- [ ] Listen for unusual noises

**Circulation System:**
- [ ] Check pump operation
- [ ] Purge air from system
- [ ] Verify flow rates
- [ ] Inspect expansion tank

### Common Problems and Solutions

| Problem | Possible Cause | Solution |
|---------|----------------|----------|
| Low heating output | Low loop flow | Check pump, purge air |
| | Undersized loop | Add supplemental heat |
| | Antifreeze too concentrated | Adjust concentration |
| High electricity use | Ground loop degradation | Check for leaks, retest |
| | Dirty coils/filters | Clean or replace |
| | Low refrigerant | Call technician |
| Short cycling | Oversized unit | Reduce output, add buffer |
| | Thermostat issue | Recalibrate or replace |
| Loop pressure drop | Leak in buried pipe | Pressure test, locate, repair |
| Pump failure | Motor burnout | Replace pump |
| | Air lock | Purge system |

### Antifreeze Guidelines

**Recommended:** Propylene glycol (food-grade)
- Concentration: 20-25% by volume
- Freeze point: approximately -10°C at 20%, -15°C at 25%
- Check annually with refractometer

**Alternative:** Methanol
- Better heat transfer
- Lower viscosity at cold temps
- More toxic—handle with care

**Avoid:** Ethylene glycol (toxic), automotive antifreeze (additives)

---

## References

### Ground Loop Design

1. [IGSHPA Calculators](https://igshpa.org/igshpa-calculators/) - International Ground Source Heat Pump Association
2. [Dandelion Energy - Borehole Depth Calculation](https://dandelionenergy.com/determining-the-appropriate-length-of-a-geothermal-borehole)
3. [GreenBuildingAdvisor - GSHP Rules of Thumb](https://www.greenbuildingadvisor.com/article/ground-source-heat-pumps-part-2-rules-of-thumb)

### Climate Battery / GAHT

4. [Ceres Greenhouse - GAHT System](https://ceresgs.com/gaht-system/)
5. [Eco Systems Design - Climate Batteries](http://www.ecosystems-design.com/climate-batteries.html)
6. [Mother Earth News - Self-Heating Greenhouse](https://www.motherearthnews.com/organic-gardening/a-self-heating-greenhouse-zbcz1707/)
7. [One Community Global - Climate Battery Research](https://onecommunityglobal.org/climate-battery/)

### Heat Pump Technology

8. [Building Science Corporation - GSHP Carbon and Efficiency](https://buildingscience.com/documents/digests/bsd-113-ground-source-heat-pumps-geothermal-for-residential-heating-and-cooling-carbon-emissions-and-efficiency)
9. [Cold Climate Housing Research Center - GSHPs in Cold Climates](https://cchrc.org/wp-content/uploads/media/Ground-Source-Heat-Pumps-in-Cold-Climates.pdf)
10. [Natural Resources Canada - Ground Source Heat Pumps](https://natural-resources.canada.ca/energy-efficiency/energy-star/products/list-certified-products/ground-source-heat-pumps)

### Enhanced Grout Materials

11. [Allan & Kavanaugh (1999) - Cement-Sand Grouts for Boreholes](https://www.researchgate.net/publication/284055449) - ASHRAE Transactions
12. [Oak Ridge National Laboratory - Thermally Enhanced Grout Analysis](https://docs.nrel.gov/docs/fy21osti/79479.pdf)
13. [Plastics Pipe Institute - Piping for Geothermal](https://plasticpipe.org/common/Uploaded%20files/1-PPI/Divisions/Building%20and%20Construction/Division%20Publications/Presentations/Plastic%20Piping%20for%20Geo%20Ground%20Loops_12-21.pdf)

### Greenhouse Heating

14. [Purdue University - Calculating Greenhouse Heating Requirements](https://www.purdue.edu/hla/sites/cea/article/calculating-greenhouse-heating-requirements/)
15. [Farm Energy - Root Zone Heating Systems](https://farm-energy.extension.org/root-zone-heating-systems-for-greenhouses/)
16. [Greenhouse Management - Determining Heat Loss](https://www.greenhousemag.com/article/technology-determining-greenhouse-heat-loss/)

### Canadian Ground Temperatures

17. [NRC - Ground Temperatures (CBD-180)](http://web.mit.edu/parmstr/Public/NRCan/CanBldgDigests/cbd180_e.html)
18. [CGS 2015 - Prediction of Ground Temperature for Prairie Provinces](https://members.cgs.ca/documents/conference2015/GeoQuebec/papers/376.pdf)

### DIY Resources

19. [GeoJerry.com - DIY Geothermal Consulting](https://geojerry.com/)
20. [123 Zero Energy - Geothermal Kits](https://www.123zeroenergy.com/pricing/geothermal-package.html)
21. [Build It Solar - DIY Geothermal + PV Project](https://www.builditsolar.com/Projects/SpaceHeating/GSHPplusPV/GSHPplusPV.htm)

### Regulations

22. [Ontario - Earth Energy Systems](https://www.ontario.ca/page/earth-energy-systems-ontario)
23. [Ottawa - Open Loop Geothermal Scoping Study](https://documents.ottawa.ca/sites/default/files/openloop_geothermal_study_en.pdf)

### Suppliers

24. [GeoPro](https://www.geoproinc.com/) - Thermal grouts, 877-580-9348
25. [CETCO](https://www.mineralstech.com/) - Geothermal grouts, 847-851-1800
26. [Ingrams Water & Air](https://iwae.com/) - DIY geothermal equipment
