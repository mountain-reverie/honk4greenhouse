# Insulation Materials Reference

## Overview

This document provides comprehensive reference data for greenhouse insulation design in Canadian climates. It covers R-value fundamentals, material properties, building code requirements, and greenhouse-specific considerations for the insulation material selection wizard.

---

## Table of Contents

1. [R-Value Fundamentals](#r-value-fundamentals)
2. [Unit Standards and Conversion](#unit-standards-and-conversion)
3. [ASTM Testing Standards](#astm-testing-standards)
4. [Effective vs Nominal R-Values](#effective-vs-nominal-r-values)
5. [Insulation Materials for Greenhouses](#insulation-materials-for-greenhouses)
   - [Spray Foam Insulation](#spray-foam-insulation)
   - [Rigid Foam Board](#rigid-foam-board)
   - [Mineral Wool / Rockwool](#mineral-wool--rockwool)
   - [Fiberglass Insulation](#fiberglass-insulation)
   - [Natural Insulation Materials](#natural-insulation-materials)
   - [Reflective / Radiant Barriers](#reflective--radiant-barriers)
6. [Canadian Building Code Requirements](#canadian-building-code-requirements)
7. [Greenhouse-Specific Considerations](#greenhouse-specific-considerations)
8. [Cost-Benefit Analysis](#cost-benefit-analysis)
9. [Material Selection Logic](#material-selection-logic)
10. [Data Model](#data-model)
11. [References](#references)

---

## R-Value Fundamentals

### What R-Value Measures

R-value (thermal resistance) measures a material's ability to resist heat flow. Higher R-values indicate better insulating performance. The R-value is the reciprocal of thermal conductance (U-value):

```
R = 1/U
R = thickness / k-value

Where:
- R = thermal resistance (m2-K/W)
- U = thermal transmittance (W/m2-K)
- k = thermal conductivity (W/m-K)
- thickness = material thickness (m)
```

### How R-Values Work

Heat naturally flows from warm areas to cold areas. Insulation slows this heat transfer by providing resistance. In Canadian greenhouse applications:

- **Winter**: Insulation resists heat escaping from the heated greenhouse interior to the cold exterior
- **Summer**: Insulation resists exterior heat entering the cooler greenhouse interior

The total R-value of an assembly is the sum of the R-values of all layers (when arranged in series):

```
R_total = R_1 + R_2 + R_3 + ... + R_n
```

---

## Unit Standards and Conversion

### SI Units (Metric) - Canadian Standard

All thermal values in this application use **metric units** as per Canadian building standards:

| Property | SI Unit | Description |
|----------|---------|-------------|
| R-value (RSI) | m2-K/W | Square meter kelvin per watt |
| U-value | W/(m2-K) | Watts per square meter kelvin |
| k-value | W/(m-K) | Thermal conductivity |
| Thickness | mm or m | Millimeters or meters |

### Imperial Units (I-P)

| Property | Imperial Unit | Description |
|----------|---------------|-------------|
| R-value | ft2-F-hr/BTU | Square foot degree Fahrenheit hour per BTU |
| U-value | BTU/(hr-ft2-F) | BTU per hour per square foot per degree Fahrenheit |

### Conversion Factors

**Critical Conversion:**
- 1 m2-K/W (RSI) = **5.678** ft2-F-hr/BTU (Imperial R-value)
- 1 ft2-F-hr/BTU = **0.176** m2-K/W

**Examples:**
| Imperial R-value | Metric RSI Value |
|------------------|------------------|
| R-10 | RSI 1.76 |
| R-20 | RSI 3.52 |
| R-30 | RSI 5.28 |
| R-40 | RSI 7.04 |
| R-50 | RSI 8.80 |
| R-60 | RSI 10.57 |

**UI Note:** The interface displays metric values (RSI) as primary, with imperial equivalents shown on hover for users familiar with that system.

---

## ASTM Testing Standards

### Primary Test Methods

R-values are determined through standardized laboratory testing:

#### ASTM C177 - Guarded Hot Plate Method

The primary absolute method for measuring steady-state heat flux and thermal transmission properties:

- **Apparatus**: Guarded-hot-plate apparatus with specimen placed between hot and cold plates
- **Procedure**: Heat flows through specimen from hot plate to cold plate; equilibrium measured when voltage and temperature readings stabilize
- **Accuracy**: Considered the reference method; does not require calibration standards
- **Application**: Best for homogeneous, low-conductivity thermal insulators
- **Standard Conditions**: Mean temperature of 24C (75F), temperature difference of 28C (50F) +/- 6C

#### ASTM C518 - Heat Flow Meter Method

A comparative method calibrated against C177:

- **Apparatus**: Heat flow meter with calibrated heat flux transducers
- **Procedure**: Material placed between temperature-controlled plates; heat flow measured by transducers
- **Accuracy**: Within +/- 2% of C177 when properly calibrated
- **Advantage**: Faster testing, suitable for production quality control
- **Operating Range**: 10C to 40C ambient, thicknesses up to 250mm

#### ASTM C1363 - Hot Box Method

For testing complete assemblies:

- **Application**: Tests R-value of systems and assemblies (walls, windows, doors)
- **Advantage**: Captures real-world thermal bridging effects
- **Use Case**: Whole-wall R-value testing

### Recent Standards Updates (2023-2025)

- **ASTM C1859-23**: Pneumatically installed loose-fill building insulation
- **ASTM C1373/C1373M-23**: Attic insulation systems under simulated winter conditions
- **ASTM C1303/C1303M-23**: Predicting long-term thermal resistance of closed-cell foam insulation
- **ASTM C687-24**: Thermal resistance of loose-fill building insulation
- **CAN/ULC S770**: Long-term thermal resistance (LTTR) for foam insulation

### Federal Regulation

The FTC R-value Rule (16 CFR Part 460) requires standardized reporting of residential insulation R-values. Canadian standards align with these requirements through NRC and provincial building codes.

---

## Effective vs Nominal R-Values

### Definitions

**Nominal R-value**: The R-value of the insulation material only, as stated on product packaging.

**Effective R-value**: The actual thermal resistance of the complete wall/roof assembly, accounting for all materials, air films, and thermal bridging.
   
**Key Insight**: Effective R-values are almost always lower than nominal R-values due to thermal bridging through framing members.

### Thermal Bridging Impact

#### Wood Stud Walls

| Assembly | Nominal R-value | Effective R-value | Loss |
|----------|-----------------|-------------------|------|
| 2x4 @ 16" o.c. with R-13 | R-13 (RSI 2.3) | R-9.9 (RSI 1.7) | ~24% |
| 2x6 @ 16" o.c. with R-20 | R-20 (RSI 3.5) | R-15.2 (RSI 2.7) | ~24% |
| 2x6 @ 16" o.c. with R-24 | R-24 (RSI 4.2) | R-18.4 (RSI 3.2) | ~23% |

#### Steel Stud Walls (Severe Thermal Bridging)

| Assembly | Nominal R-value | Effective R-value | Loss |
|----------|-----------------|-------------------|------|
| 6" steel stud @ 16" o.c. with R-21 | R-21 (RSI 3.7) | R-7.4 (RSI 1.3) | **65%** |
| 3.5" steel stud @ 16" o.c. with R-13 | R-13 (RSI 2.3) | R-4.8 (RSI 0.85) | **63%** |

### Framing Factor

The "framing factor" represents the percentage reduction in R-value due to thermal bridging:

```
Effective R = Nominal R x (1 - Framing Factor)
```

Typical framing factors:
- Wood stud walls: 20-25%
- Steel stud walls: 50-65%
- Advanced framing (24" o.c.): 15-18%

### Calculation Method (ASHRAE/NBC)

The National Building Code uses the isothermal planes (series-parallel) method:

```
1/R_effective = (Area_insulation/R_insulation + Area_framing/R_framing) / Total_Area
```

Or using U-values (weighted average):

```
U_effective = (U_insulation x Area_insulation + U_framing x Area_framing) / Total_Area
R_effective = 1 / U_effective
```

### Solutions to Improve Effective R-Value

1. **Continuous Exterior Insulation**: Place rigid foam outside the framing to eliminate thermal bridging through studs
2. **Advanced Framing**: 24" o.c. spacing, eliminate unnecessary studs, insulated headers
3. **Double-Stud Walls**: Two separated stud walls with continuous insulation between
4. **Structural Insulated Panels (SIPs)**: Foam core between sheathing, minimal framing
5. **Thermally Broken Attachments**: Use thermal break clips for cladding

**Example Improvement:**
- 2x6 wall with R-24 batts only: Effective R-18.4
- Same wall + 1" polyiso (R-6) exterior: Effective R-27 (+47% improvement)

---

## Insulation Materials for Greenhouses

### Spray Foam Insulation

#### Closed-Cell Spray Foam (ccSPF) - Recommended for Greenhouses

**R-Value:** R-6.0 to R-7.0 per inch (RSI 1.06-1.23 per 25mm)
**Long-term R-value:** R-5.5 per inch (RSI 0.97) - accounts for thermal drift

| Thickness (mm) | RSI Value (m2-K/W) | Imperial R-value | Notes |
|----------------|---------------------|------------------|-------|
| 25 | 1.0-1.1 | R-5.5-6.2 | Minimum air barrier thickness |
| 38 | 1.5-1.7 | R-8.3-9.3 | Minimum vapor retarder thickness |
| 50 | 2.0-2.2 | R-11-12.5 | Good for mild climates |
| 75 | 3.0-3.3 | R-16.5-18.7 | Recommended Zone 4-5 |
| 100 | 4.0-4.4 | R-22-25 | Recommended Zone 6+ |
| 125 | 5.0-5.5 | R-27.5-31 | High performance |
| 150 | 6.0-6.6 | R-33-37.5 | Maximum typical |

**Properties:**

| Property | Value |
|----------|-------|
| k-value | 0.023 W/(m-K) |
| Density | 32 kg/m3 (2 lb/ft3) |
| Vapor Permeance | <1 perm at 38mm (Class II vapor retarder) |
| Air Permeance | Impermeable at 38mm |
| Moisture Absorption | <3% by volume |
| Fire Rating | Class 1 (with thermal barrier) |
| Service Temperature | -40C to 80C |
| Lifespan | 80-100 years |

**Advantages for Greenhouses:**
- Acts as both air barrier and vapor retarder
- Excellent moisture resistance for high-humidity environments
- No settling or degradation from humidity
- Eliminates thermal bridging when applied continuously
- Seals all gaps and cracks automatically

**Limitations:**
- Requires professional installation with specialized equipment
- Higher upfront cost ($10-18/m2 per 25mm)
- Must be covered with thermal barrier for fire safety
- Subject to thermal drift (R-value decreases 5-10% in first 2 years)
- UV sensitive - must be protected from sunlight

**Best For:** North walls, foundation walls, roof assemblies, any high-humidity application

#### Open-Cell Spray Foam (ocSPF)

**R-Value:** R-3.5 to R-3.8 per inch (RSI 0.62-0.67 per 25mm)
**Long-term R-value:** R-3.6 per inch (RSI 0.63)

| Thickness (mm) | RSI Value (m2-K/W) | Imperial R-value | Notes |
|----------------|---------------------|------------------|-------|
| 50 | 1.2-1.3 | R-7-7.6 | Minimum useful |
| 75 | 1.8-2.0 | R-10.5-11.4 | Basic insulation |
| 89 | 2.2-2.4 | R-12.4-13.5 | 2x4 cavity fill |
| 100 | 2.5-2.7 | R-14-15.2 | Standard application |
| 140 | 3.5-3.8 | R-19.6-21.3 | 2x6 cavity fill |
| 150 | 3.7-4.0 | R-21-22.8 | Good performance |

**Properties:**

| Property | Value |
|----------|-------|
| k-value | 0.038 W/(m-K) |
| Density | 8 kg/m3 (0.5 lb/ft3) |
| Vapor Permeance | 16-35 perms (vapor permeable) |
| Air Permeance | Impermeable at 89mm |
| Moisture Absorption | Absorbs and releases moisture |
| Fire Rating | Requires thermal barrier |
| Lifespan | 80-100 years |

**Advantages:**
- Lower cost than closed-cell ($6-10/m2 per 25mm)
- Excellent air sealing
- Better sound insulation than closed-cell
- No thermal drift (air-blown)
- Allows moisture to dry in both directions

**Limitations:**
- NOT a vapor barrier - requires separate vapor retarder in cold climates
- Absorbs water if exposed to bulk moisture
- Not suitable for below-grade applications
- Lower R-value per inch

**Greenhouse Caution:** Open-cell foam is NOT recommended for greenhouse applications due to high humidity levels. Moisture absorption will reduce R-value and potentially lead to mold issues.

---

### Rigid Foam Board

#### Extruded Polystyrene (XPS) - Blue/Pink Board

**R-Value:** R-5.0 per inch initial (RSI 0.88); R-4.5 per inch long-term (RSI 0.79)

**Critical Note:** XPS experiences thermal drift as blowing agents escape, losing 10-20% of R-value over time. Use LTTR (long-term) values for design.

| Thickness (mm) | RSI Initial | RSI Long-term | Imperial Long-term | Notes |
|----------------|-------------|---------------|-------------------|-------|
| 25 | 0.88 | 0.79 | R-4.5 | Minimum |
| 38 | 1.32 | 1.19 | R-6.75 | Foundation |
| 50 | 1.76 | 1.58 | R-9 | Standard |
| 75 | 2.64 | 2.38 | R-13.5 | Good |
| 100 | 3.52 | 3.17 | R-18 | Excellent |
| 150 | 5.28 | 4.75 | R-27 | Maximum typical |

**Properties:**

| Property | Value |
|----------|-------|
| k-value | 0.029 W/(m-K) initial; 0.032 long-term |
| Density | 25-40 kg/m3 |
| Vapor Permeance | 0.3-1.2 perms (semi-impermeable) |
| Water Absorption | 0.1-0.3% by volume (excellent) |
| Compressive Strength | 100-700 kPa (15-100 psi) |
| Service Temperature | -50C to 75C |

**Below-Grade Derated Values (per Canadian FPSF standards):**
- Vertical installation: R-4.5/inch (RSI 0.79)
- Horizontal installation: R-4.0/inch (RSI 0.70)

**Advantages:**
- Excellent moisture resistance - best for below-grade
- High compressive strength for foundations
- R-value improves slightly in cold temperatures
- Stable long-term performance in wet environments
- Easy to cut and install (DIY-friendly)

**Limitations:**
- Contains HBCD or HFCs (environmental concerns)
- Subject to ant tunneling (use treated versions)
- Moderate thermal drift over time
- Must protect from UV exposure

**Best For:** Foundation walls, frost skirts, below-grade applications, exterior continuous insulation

**Canadian Products:** FOAMULAR (Owens Corning), Styrofoam (DuPont)

#### Expanded Polystyrene (EPS) - White Beadboard

**R-Value:** R-3.6 to R-4.2 per inch (RSI 0.63-0.74), depending on density
**Long-term:** R-value is stable (no thermal drift - uses air as insulating gas)

| Thickness (mm) | RSI Value | Imperial R-value | Notes |
|----------------|-----------|------------------|-------|
| 25 | 0.63-0.74 | R-3.6-4.2 | Basic |
| 50 | 1.26-1.48 | R-7.2-8.4 | Standard |
| 75 | 1.89-2.22 | R-10.8-12.6 | Good |
| 100 | 2.52-2.96 | R-14.4-16.8 | Very good |
| 150 | 3.78-4.44 | R-21.6-25.2 | Excellent |

**Density vs R-Value:**

| Type | Density (kg/m3) | R-value per inch |
|------|-----------------|------------------|
| Type I | 15 | R-3.6 |
| Type II | 22 | R-3.8 |
| Type VIII | 18 | R-4.0 |
| Type IX | 20 | R-4.2 |

**Properties:**

| Property | Value |
|----------|-------|
| k-value | 0.036-0.040 W/(m-K) |
| Vapor Permeance | 2-5 perms (semi-permeable) |
| Water Absorption | 2-4% by volume |
| Compressive Strength | 70-175 kPa (10-25 psi) |
| Service Temperature | -50C to 75C |

**Advantages:**
- Lowest cost rigid foam insulation
- No thermal drift - stable R-value over time
- Consistent performance across temperatures
- Can be recycled
- Breathable (allows some moisture diffusion)
- Pentane blowing agent has low GWP

**Limitations:**
- Higher water absorption than XPS
- Lower compressive strength
- Visible bead structure (less finished appearance)
- Lower R-value per inch than other foams

**Best For:** Budget applications, SIPs, non-structural walls, where moisture exposure is controlled

**Canadian Products:** DuroSpan (Home Depot Canada), Insulfoam

#### Polyisocyanurate (Polyiso)

**R-Value:** R-5.6 to R-6.8 per inch at 24C (RSI 0.99-1.20)
**Critical Warning:** R-value decreases significantly in cold temperatures

| Thickness (mm) | RSI at 24C | RSI at -4C | Imperial at -4C | Notes |
|----------------|------------|------------|-----------------|-------|
| 25 | 1.0-1.2 | 0.5-0.6 | R-2.8-3.4 | Minimum |
| 38 | 1.5-1.8 | 0.75-0.9 | R-4.3-5.1 | Standard |
| 50 | 2.0-2.4 | 1.0-1.2 | R-5.7-6.8 | Good |
| 75 | 3.0-3.6 | 1.5-1.8 | R-8.5-10.2 | Very good |
| 100 | 4.0-4.8 | 2.0-2.4 | R-11.4-13.6 | Excellent |

**Temperature-Dependent R-Value:**

| Mean Temperature | R-value per inch | RSI per 25mm |
|------------------|------------------|--------------|
| 24C (75F) | R-6.0-6.8 | 1.06-1.20 |
| 10C (50F) | R-5.0-5.5 | 0.88-0.97 |
| -4C (25F) | R-3.5-4.0 | 0.62-0.70 |
| -18C (0F) | R-2.5-3.0 | 0.44-0.53 |

**Properties:**

| Property | Value |
|----------|-------|
| k-value | 0.022-0.026 W/(m-K) at 24C |
| Vapor Permeance | <1 perm (foil-faced) |
| Water Absorption | Up to 3% if facers damaged |
| Fire Resistance | Good (intumescent char) |
| Service Temperature | -40C to 120C |

**Advantages:**
- Highest R-value per inch at room temperature
- Excellent fire performance (chars rather than melts)
- Foil facers act as radiant barrier
- Good for interior applications
- Wide availability

**Limitations:**
- **Not recommended for cold climate applications** - R-value drops 50%+ below freezing
- Not suitable for below-grade use
- Facers can absorb moisture if damaged
- Higher cost than EPS/XPS
- Experiences thermal drift

**Best For:** Interior applications only, roof insulation above heated spaces, warm climate zones
**NOT For:** Canadian greenhouse north walls, below-grade, any cold-side application

---

### Mineral Wool / Rockwool

**R-Value:** R-3.0 to R-4.2 per inch (RSI 0.53-0.74), depending on density

| Format | Thickness (mm) | RSI Value | Imperial R-value |
|--------|----------------|-----------|------------------|
| Batt (standard) | 89 (3.5") | 2.6 | R-15 |
| Batt (standard) | 140 (5.5") | 4.1 | R-23 |
| Batt (standard) | 184 (7.25") | 5.3 | R-30 |
| Batt (standard) | 241 (9.5") | 7.0 | R-40 |
| Rigid Board | 25 | 0.88 | R-5.0 |
| Rigid Board | 50 | 1.76 | R-10 |
| Rigid Board | 75 | 2.64 | R-15 |

**Properties:**

| Property | Value |
|----------|-------|
| k-value | 0.034-0.040 W/(m-K) |
| Density | Batts: 24-48 kg/m3; Boards: 80-200 kg/m3 |
| Vapor Permeance | ~50 perms (vapor open) |
| Water Absorption | Hydrophobic (repels liquid water) |
| Fire Resistance | Non-combustible to 1100C+ |
| Sound Transmission | Excellent (NRC 0.95-1.05) |
| Recycled Content | ~70% |

**Moisture Behavior:**

Mineral wool has a unique moisture profile ideal for greenhouses:
- **Hydrophobic**: Liquid water drains through without absorption
- **Vapor Permeable**: Water vapor passes through, allowing drying
- **Maintains R-value**: Retains 90%+ of insulating value even when exposed to moisture
- **Mold Resistant**: Inorganic material does not support mold growth

**Advantages for Greenhouses:**
- Excellent moisture tolerance - ideal for high-humidity environments
- Non-combustible - superior fire safety
- Does not lose R-value when damp
- Inorganic - will not rot, mold, or attract pests
- Excellent sound insulation
- Higher R-value than fiberglass per inch

**Limitations:**
- Higher cost than fiberglass (25-50% more)
- Heavier and denser than fiberglass
- Requires protective equipment during installation (dust/fibers)
- Not a vapor barrier (requires separate in cold climates)
- Not suitable for below-grade without protection

**Best For:** Wall cavities, fire-rated assemblies, high-humidity environments, soundproofing

**Canadian Products:** ROCKWOOL ComfortBatt, ROCKWOOL Safe'n'Sound, ROXUL

---

### Fiberglass Insulation

**R-Value:** R-2.2 to R-4.3 per inch (RSI 0.39-0.76), depending on density

#### Standard Density Fiberglass Batts

| Thickness (mm) | RSI Value | Imperial R-value | Application |
|----------------|-----------|------------------|-------------|
| 89 (3.5") | 2.3 | R-13 | 2x4 wall cavity |
| 140 (5.5") | 3.5 | R-20 | 2x6 wall cavity |
| 184 (7.25") | 4.0 | R-22 | Standard |
| 241 (9.5") | 5.3 | R-30 | 2x10 joist |
| 292 (11.5") | 6.7 | R-38 | 2x12 joist |

#### High-Density Fiberglass Batts

| Thickness (mm) | RSI Value | Imperial R-value | Notes |
|----------------|-----------|------------------|-------|
| 89 (3.5") | 2.6 | R-15 | 2x4 wall cavity |
| 140 (5.5") | 4.0 | R-23 | 2x6 wall cavity |
| 241 (9.5") | 6.2 | R-35 | High performance |

**Properties:**

| Property | Value |
|----------|-------|
| k-value | 0.035-0.045 W/(m-K) |
| Density | Standard: 10-14 kg/m3; High-density: 16-24 kg/m3 |
| Vapor Permeance | >50 perms (vapor open) |
| Water Absorption | Does not absorb, but holds water |
| Fire Resistance | Non-combustible (glass) |
| Recycled Content | 20-30% |

**Critical Moisture Warning for Greenhouses:**

Fiberglass insulation has significant limitations in high-humidity greenhouse environments:

- **R-value degradation**: A 1.5% increase in moisture content reduces R-value by up to 50%
- **Water retention**: While glass fibers don't absorb water, the batt holds water in air pockets
- **Sagging**: Wet fiberglass becomes heavy and sags away from contact surfaces, creating gaps
- **Mold risk**: Retained moisture promotes mold growth on adjacent organic materials (wood, paper facing)
- **Condensation**: In high-humidity greenhouses, vapor can condense within the insulation cavity

**Requirements for Greenhouse Use:**

If fiberglass is used in greenhouse applications, these measures are essential:
1. **Vapor barrier**: 6 mil polyethylene on the warm (interior) side
2. **Air sealing**: Complete air barrier to prevent humid air infiltration
3. **Ventilation**: Ensure wall cavities can dry to exterior
4. **Inspection**: Regular monitoring for moisture accumulation
5. **Location**: Use only in areas away from high-humidity zones

**Advantages:**
- Lowest cost insulation material
- Easy DIY installation
- Non-combustible
- Does not support mold growth (on fibers)
- Widely available in all sizes

**Limitations:**
- **Not recommended for greenhouse applications** due to moisture sensitivity
- Requires separate vapor barrier
- Loses significant R-value when wet
- Air gaps reduce effectiveness
- Skin/respiratory irritation during installation

**Best For:** Budget-conscious residential builds with proper vapor barriers
**NOT For:** High-humidity greenhouse environments

---

### Natural Insulation Materials

#### Straw Bale Insulation

**R-Value:** R-1.45 to R-2.4 per inch (RSI 0.26-0.42), depending on straw orientation

| Wall Type | Typical Thickness | Total RSI | Total Imperial R-value |
|-----------|-------------------|-----------|------------------------|
| Standard bale (laid flat) | 450mm (18") | 5.3-5.8 | R-30-33 |
| Standard bale (on edge) | 580mm (23") | 5.8-7.0 | R-33-40 |
| Straw SIPs | 300mm (12") | 5.3 | R-30 |

**Properties:**

| Property | Value |
|----------|-------|
| Density | 80-120 kg/m3 (compressed) |
| Fire Resistance | Good when plastered (2+ hour rating) |
| Vapor Permeance | High (breathable) |
| Carbon Footprint | Negative (sequesters ~129 kg CO2/m3) |
| Embodied Energy | 1/4 of fiberglass for same R-value |
| Lifespan | 100+ years when properly protected |

**Advantages:**
- Carbon-negative material
- Excellent thermal mass + insulation combination
- Low cost (agricultural waste product)
- Biodegradable at end of life
- Superior thermal performance in real-world conditions
- Fire resistant when plastered

**Limitations:**
- Thick walls required (450-580mm)
- Requires protection from bulk moisture
- Not suitable for high-humidity environments without careful design
- Labor-intensive installation
- Building code acceptance varies by jurisdiction
- Requires plastering (lime, clay, or cement)

**Best For:** Sustainable builds, earth-bermed north walls, where wall thickness is acceptable

#### Hempcrete

**R-Value:** R-1.7 to R-2.5 per inch (RSI 0.30-0.44)

| Wall Thickness | RSI Value | Imperial R-value |
|----------------|-----------|------------------|
| 200mm (8") | 2.5-3.5 | R-14-20 |
| 300mm (12") | 3.7-5.3 | R-21-30 |
| 400mm (16") | 5.0-7.0 | R-28-40 |

**Properties:**

| Property | Value |
|----------|-------|
| Density | 275-350 kg/m3 |
| Thermal Mass | Excellent (phase shift + insulation) |
| Vapor Permeance | High (breathable) |
| Carbon Sequestration | 100+ kg CO2/m3 |
| Fire Resistance | Non-combustible |
| Moisture Management | Excellent (hygroscopic) |

**Unique Thermal Behavior:**

Hempcrete combines insulation with thermal mass:
- **Phase Shift**: Absorbs and releases heat over 8-12 hours
- **Dynamic Performance**: Real-world performance exceeds static R-value predictions
- **Hygroscopic**: Absorbs and releases moisture, regulating humidity

**Advantages:**
- Carbon-negative construction
- Excellent humidity regulation (ideal for greenhouses)
- Non-combustible
- Pest and mold resistant
- Breathable walls that manage moisture
- Long lifespan (centuries)
- No VOCs or toxins

**Limitations:**
- Thick walls required for adequate R-value (300-400mm)
- Requires structural frame (not load-bearing)
- Higher labor costs
- Limited contractor experience
- Curing time (4-8 weeks)
- Material availability varies by region

**Best For:** Sustainable greenhouse builds, humidity-regulating walls, carbon-conscious projects

---

### Reflective / Radiant Barriers

**R-Value Contribution:** R-1 to R-17 equivalent, depending on installation and air gap

Reflective barriers work differently from mass insulation - they reduce radiant heat transfer rather than conductive heat transfer.

| Product Type | Emissivity | Equivalent R-value | Notes |
|--------------|------------|-------------------|-------|
| Foil-faced polyiso | 0.03-0.05 | Adds R-2-3 | Combined conductive + radiant |
| Bubble foil (single) | 0.03-0.05 | R-3-5 | Requires 19mm air gap |
| Bubble foil (double) | 0.03-0.05 | R-5-8 | Better performance |
| Perforated radiant barrier | 0.05-0.10 | R-2-4 | Allows vapor transmission |
| Reflective attic barrier | 0.03-0.05 | Variable | Primarily radiant reduction |

**How Radiant Barriers Work:**

- **Reflectivity**: Reflects 95-97% of radiant heat
- **Emissivity**: Emits only 3-5% of absorbed heat as radiation
- **Air Gap Required**: Must have 19mm+ air gap facing the reflective surface
- **Direction**: Most effective facing an air space

**Greenhouse Applications:**

1. **North Wall Interior**: Reflects radiant heat back into greenhouse
2. **Thermal Curtains**: Reduces nighttime radiant heat loss through glazing (up to 50%)
3. **Under-Roof**: Reduces summer heat gain
4. **Growing Space**: Reflects light to lower plant canopy

**Key Research Finding:**
"Up to 85% of the heat loss from a greenhouse occurs at night" - Virginia Tech University. Thermal blankets with reflective surfaces can reduce energy use by 25-50%.

**Installation Requirements:**
- Face reflective surface toward warmer space (interior in winter)
- Maintain 19-25mm air gap adjacent to reflective surface
- Keep surface clean and dust-free (dust reduces reflectivity)
- Avoid direct contact with other materials

**Advantages:**
- Lightweight and easy to install
- Cost-effective supplemental insulation
- Reflects light for better plant growth
- Summer cooling benefit (reflects solar radiation)
- Thin profile saves space

**Limitations:**
- Requires air gap to function
- Dust accumulation reduces effectiveness
- Limited conductive insulation value
- Condensation can occur on foil surface
- Not effective without proper installation

**Best For:** Supplemental insulation, thermal curtains, north wall lining, growing area optimization

---

## Canadian Building Code Requirements

### Overview

Thermal resistance requirements in Canada fall under provincial/territorial jurisdiction. The National Building Code (NBC) and National Energy Code of Canada for Buildings (NECB) provide model codes that provinces adopt and modify.

**Key Standards:**
- **NBC 2020**: Part 9 Section 9.36 (housing and small buildings)
- **NECB 2020**: Commercial and larger buildings
- **NBC 2020 Part 2**: New requirements for Group G (agricultural) occupancies

### Climate Zones and Heating Degree Days

| Zone | HDD Range | Regions |
|------|-----------|---------|
| 4 | <3000 | Coastal BC (limited areas) |
| 5 | 3000-3999 | Southern Ontario, Lower BC Mainland |
| 6 | 4000-4999 | Central Ontario, Southern Prairies |
| 7A | 5000-5999 | Northern Ontario, Central Prairies, Edmonton |
| 7B | 6000-6999 | Northern regions |
| 8 | 7000+ | Arctic regions |

### Provincial Requirements Summary

#### Ontario (SB-12)

**Effective January 1, 2025** (Ontario Building Code 2024)

| Component | Zone 1 (<5000 HDD) | Zone 2 (>5000 HDD) |
|-----------|-------------------|-------------------|
| Walls (effective) | RSI 3.87 (R-22) | RSI 4.23 (R-24) |
| Attic/Ceiling | RSI 8.81 (R-50) | RSI 10.57 (R-60) |
| Foundation Walls | RSI 2.11 (R-12) | RSI 3.52 (R-20) |
| Below-Grade Slab | RSI 1.76 (R-10) | RSI 1.76 (R-10) |

**SB-12 Compliance Notes:**
- Compliance packages allow trading between components
- Can meet nominal OR effective R-value for each component
- Window-to-wall ratio affects insulation requirements

#### British Columbia (BC Building Code + Step Code)

**Standard Code Requirements:**

| Component | Zone 4 | Zone 5 | Zone 6/7A |
|-----------|--------|--------|-----------|
| Walls (effective) | RSI 2.78 (R-15.8) | RSI 3.35 (R-19) | RSI 3.87 (R-22) |
| Attic/Ceiling | RSI 8.81 (R-50) | RSI 8.81 (R-50) | RSI 10.57 (R-60) |
| Foundation | RSI 2.11 (R-12) | RSI 2.99 (R-17) | RSI 3.87 (R-22) |

**BC Energy Step Code:**
Higher performance tiers progressively increase requirements toward net-zero ready (Step 5).

#### Quebec (Construction Code + Novoclimat)

**Standard Code (Part 11):**

| Component | <6000 HDD | >6000 HDD |
|-----------|-----------|-----------|
| Walls | RSI 4.31 (R-24.5) | RSI 4.31 (R-24.5) |
| Attic/Ceiling | RSI 7.22 (R-41) | RSI 8.81 (R-50) |
| Foundation | RSI 2.99 (R-17) | RSI 3.52 (R-20) |
| Thermal Bridge Coverage | RSI 0.70 (R-4) | RSI 0.70 (R-4) |

**Novoclimat 2.0 (Voluntary High-Performance):**

| Component | Requirement |
|-----------|-------------|
| Walls | RSI 5.19 (R-29.5) |
| Roof | RSI 10.92 (R-62) |
| Heated Slab | RSI 2.82 (R-16) |
| Air Tightness | 1.5 ACH @ 50 Pa |

#### Alberta (NBC 2023 Alberta Edition + NECB 2020)

**Effective May 1, 2024:**

| Component | Zone 6 | Zone 7A | Zone 7B |
|-----------|--------|---------|---------|
| Walls (effective) | RSI 3.87 (R-22) | RSI 4.40 (R-25) | RSI 4.93 (R-28) |
| Attic/Ceiling | RSI 8.81 (R-50) | RSI 10.57 (R-60) | RSI 10.57 (R-60) |
| Foundation | RSI 2.11 (R-12) | RSI 2.64 (R-15) | RSI 3.17 (R-18) |

### Agricultural Building Requirements

The NBC 2020 introduced Part 2 of Division B specifically for large farm buildings:

- New Group G (agricultural) occupancy classification
- Requirements address fire protection, structural design, and HVAC
- Prescriptive insulation R-values not specified for agricultural buildings
- Performance-based approach using NECB or energy modeling recommended
- Provincial adoption varies - check local authority having jurisdiction

**Practical Guidance for Greenhouses:**
Most provinces do not have prescriptive insulation requirements for greenhouses. Design should be based on:
1. Heating/cooling load calculations
2. Economic analysis of insulation investment vs energy costs
3. Target growing conditions (minimum temperatures, humidity)
4. Local climate data (HDD, design temperatures)

---

## Greenhouse-Specific Considerations

### Balancing Insulation vs Light Transmission

The fundamental greenhouse design challenge: maximizing light for plant growth while minimizing heat loss.

**Design Strategies:**

1. **Asymmetric Design**: Insulated north wall/roof, glazed south wall/roof
2. **Selective Glazing**: Multi-wall polycarbonate (insulating) vs single glass (transparent)
3. **Thermal Curtains**: Movable insulation deployed at night
4. **Strategic Opacity**: Insulate where light is not needed (north, lower walls)

**Light vs Insulation Trade-offs:**

| Glazing Type | R-value (RSI) | Light Transmission |
|--------------|---------------|-------------------|
| Single glass | 0.16 (R-0.9) | ~90% |
| Double glass | 0.35 (R-2.0) | ~80% |
| Twin-wall PC 8mm | 0.27 (R-1.54) | ~85% |
| Triple-wall PC 16mm | 0.42 (R-2.4) | ~77% |
| Five-wall PC 32mm | 0.99 (R-5.6) | ~70% |

**Rule of Thumb:**
- South-facing glazing: Prioritize light transmission with moderate insulation
- North-facing surfaces: Maximum insulation (opaque acceptable)
- East/West: Balance based on morning/afternoon sun requirements

### Moisture and Condensation Management

Greenhouses present unique moisture challenges:
- **Typical humidity**: 60-80% relative humidity
- **Transpiration**: Plants release significant water vapor
- **Irrigation**: Additional moisture source
- **Temperature swings**: Wide diurnal temperature variation

**Condensation Prevention Strategies:**

1. **Material Selection**:
   - Use closed-cell foam or mineral wool (moisture-tolerant)
   - Avoid fiberglass and open-cell foam in humid zones
   - Select vapor-impermeable materials for cold-side insulation

2. **Vapor Control**:
   - Vapor barrier/retarder on warm (interior) side
   - Allow drying to exterior (vapor-permeable exterior sheathing)
   - Continuous air barrier to prevent humid air infiltration

3. **Ventilation**:
   - Maintain air movement to prevent stagnant humid pockets
   - Use exhaust fans to remove excess humidity
   - Size vents for 60 air changes per hour peak

4. **Thermal Strategy**:
   - Maintain interior surface temperatures above dew point
   - Use continuous insulation to eliminate cold spots
   - Address thermal bridges at glazing-to-wall transitions

5. **Anti-Condensate Glazing**:
   - Treated greenhouse films cause moisture to sheet rather than drip
   - Water runs to edges instead of dripping on plants
   - Reduces disease pressure from water on foliage

### North Wall Insulation Strategies

The north wall receives no direct solar gain and should be heavily insulated.

#### Super-Insulated North Wall

Target: RSI 7.0+ (R-40+)

**Construction Options:**

1. **Double-Stud Wall**:
   - Two 2x4 stud walls with 25-50mm gap
   - Filled with dense-pack cellulose or mineral wool
   - Total RSI 7-9 (R-40-50)
   - Eliminates thermal bridging

2. **SIP Panel**:
   - Structural Insulated Panel (EPS or polyiso core)
   - 200-300mm thickness
   - RSI 5.3-8.8 (R-30-50)
   - Fast installation, excellent air sealing

3. **Closed-Cell Spray Foam**:
   - 100-150mm on interior of frame wall
   - RSI 4.4-6.6 (R-25-37)
   - Combines air/vapor barrier with insulation

4. **Exterior Rigid Foam + Cavity**:
   - 100mm XPS exterior + R-24 mineral wool cavity
   - RSI 6.4 (R-36) effective
   - Good moisture management

#### Earth-Bermed North Wall

Using earth as insulation and thermal mass:

**Benefits:**
- Constant ground temperature (8-12C at 2m depth in Canada)
- Massive thermal mass moderates temperature swings
- Natural frost protection
- Free insulation (soil)

**Design Parameters:**
- Berm soil: RSI ~0.07/25mm (R-0.4/inch) - low R-value but massive thermal mass
- Rigid insulation between wall and soil: RSI 2.5-3.5 (R-14-20)
- Below frost line (1.2-2.4m depending on location)
- Waterproofing critical (membrane + drainage)

**Equivalent Thermal Performance:**
- 2m of earth + RSI 1.76 (R-10) rigid insulation = effective RSI 7+ (R-40+) due to stable ground temperature

**Construction:**
1. Concrete or ICF wall below grade
2. Waterproof membrane (EPDM or bituminous)
3. XPS rigid insulation (50-100mm)
4. Drainage layer (gravel or drainage board)
5. Backfill with well-draining soil
6. Grade away from structure

### Foundation and Perimeter Insulation

#### Frost Protection

**Frost Depth by Region:**

| Region | Typical Frost Depth |
|--------|-------------------|
| Southern Ontario | 1.0-1.2m |
| Northern Ontario | 1.5-2.0m |
| Southern Prairies | 1.2-1.5m |
| Central Prairies | 1.5-2.0m |
| Northern BC | 1.2-1.8m |

**Frost-Protected Shallow Foundation (FPSF):**

FPSF uses horizontal wing insulation to prevent frost penetration, allowing shallower footings:

- **Vertical Insulation**: RSI 1.76-3.52 (R-10-20) to frost depth
- **Horizontal Wing**: 0.6-1.2m wide, RSI 1.76-2.64 (R-10-15)
- **Corner Insulation**: Thicker/wider at corners (heat loss concentrated)

**Derated R-Values for Below-Grade (per FPSF guidelines):**
- XPS vertical: R-4.5/inch (90% of nominal)
- XPS horizontal: R-4.0/inch (80% of nominal)
- EPS: Derate additional 10-20%

#### Foundation Wall Insulation

**Options by Climate Zone:**

| Zone | Foundation RSI | Frost Skirt RSI | Recommended Material |
|------|---------------|-----------------|---------------------|
| 5 | 2.11 (R-12) | 1.76 (R-10) | XPS 50mm + 25mm wing |
| 6 | 2.64 (R-15) | 2.11 (R-12) | XPS 75mm + 50mm wing |
| 7A | 3.17 (R-18) | 2.64 (R-15) | XPS 100mm + 50mm wing |
| 7B+ | 3.52 (R-20) | 2.64 (R-15) | XPS 100mm + 75mm wing |

### Thermal Bridging at Glazing-Wall Transitions

The junction between insulated walls and glazing is a critical thermal bridge location.

**Problems:**
- Heat escapes through frame connections
- Cold spots cause condensation
- Reduced effective R-value of wall assembly
- Potential for mold/rot at junction

**Solutions:**

1. **Align Thermal Planes**:
   - Position glazing within 50mm of wall continuous insulation
   - Avoid deep setbacks that create thermal bridges

2. **Thermally Broken Frames**:
   - Use aluminum frames with thermal break
   - Wood or fiberglass frames preferred (lower conductivity)
   - Avoid direct metal-to-metal connections

3. **Continuous Insulation Wrap**:
   - Extend rigid insulation to overlap glazing frame
   - Minimum 50mm overlap recommended
   - Seal air barrier continuously across junction

4. **Insulated Curbs**:
   - For roof glazing, use insulated curbs
   - Wrap curb exterior with rigid foam
   - Flash carefully for water management

5. **Jamb Details**:
   - Pack insulation tight to frames
   - Use low-expansion foam to seal gaps
   - Install interior trim to cover junction

**Target**: Limit thermal bridge impact to <10% reduction in wall effective R-value

---

## Cost-Benefit Analysis

### Material Costs (CAD 2024-2025)

| Material | Cost per m2 at RSI 3.52 (R-20) | $/RSI/m2 |
|----------|-------------------------------|----------|
| Fiberglass Batts | $8-15 | $2.50-4.50 |
| Mineral Wool Batts | $15-25 | $4.50-7.00 |
| EPS Rigid | $18-30 | $5.00-8.50 |
| XPS Rigid | $35-55 | $10-16 |
| Polyiso Rigid | $40-60 | $11-17 |
| Open-Cell Spray Foam | $45-70 | $13-20 |
| Closed-Cell Spray Foam | $90-140 | $25-40 |
| Straw Bale (materials) | $15-30 | $4-8 |
| Hempcrete (materials) | $80-120 | $22-35 |

**Notes:**
- Spray foam prices include professional installation
- Batt and rigid foam prices are materials only
- Regional pricing varies significantly
- Bulk/contractor pricing may reduce costs 15-30%

### Installation Costs

| Material | DIY Feasible? | Professional Cost/m2 |
|----------|---------------|---------------------|
| Fiberglass Batts | Yes | $15-25 |
| Mineral Wool | Yes | $20-35 |
| Rigid Foam | Yes | $25-40 |
| Spray Foam | No | Included in material |
| Straw Bale | With training | $50-100 |
| Hempcrete | With training | $100-180 |

### Payback Period Analysis

**Variables:**
- Energy cost ($/kWh or $/GJ)
- Climate zone (HDD)
- Current vs proposed R-value
- Heating system efficiency

**Typical Payback Periods:**

| Upgrade | Zone 5-6 | Zone 7+ |
|---------|----------|---------|
| RSI 2.1 to RSI 3.5 (R-12 to R-20) walls | 5-8 years | 3-5 years |
| RSI 5.3 to RSI 8.8 (R-30 to R-50) ceiling | 4-7 years | 3-5 years |
| Add RSI 1.76 (R-10) foundation | 8-12 years | 5-8 years |
| Air sealing + insulation | 2-4 years | 2-3 years |

**Key Finding:**
Air sealing combined with insulation improvements has the fastest payback. Insulation without air sealing may achieve only 50% of potential savings.

### Life-Cycle Cost Considerations

| Material | Initial Cost | Lifespan | Maintenance | End-of-Life |
|----------|-------------|----------|-------------|-------------|
| Fiberglass | Low | 10-25 years | Replace if wet | Landfill |
| Mineral Wool | Medium | 50+ years | Minimal | Recyclable |
| XPS/EPS | Medium | 50+ years | Minimal | Recyclable |
| Closed-Cell Foam | High | 80-100 years | Minimal | Landfill |
| Straw Bale | Low-Medium | 100+ years | Inspect annually | Compost |
| Hempcrete | High | 100+ years | Minimal | Compost |

### Energy Savings Estimates

**Per RSI (R-5.7) increase in wall insulation:**
- Zone 5: 5-8% heating energy reduction
- Zone 6: 7-10% heating energy reduction
- Zone 7: 9-12% heating energy reduction

**Example (Zone 6, 100m2 greenhouse, propane heat):**
- Upgrade walls RSI 2.1 to RSI 5.3 (R-12 to R-30)
- Annual heating reduction: ~25%
- Savings: ~$800-1,200/year at $0.70/L propane
- Investment: ~$3,000-5,000
- Simple payback: 3-5 years

---

## Material Selection Logic

### Decision Tree

```
What surface are you insulating?
|
+-- Below-Grade (Foundation, Frost Skirt)
|   |
|   +-- Use XPS (blue/pink board)
|       +-- Zone 5: 50-75mm (RSI 1.58-2.38)
|       +-- Zone 6: 75-100mm (RSI 2.38-3.17)
|       +-- Zone 7+: 100-150mm (RSI 3.17-4.75)
|
+-- North Wall
|   |
|   +-- Is this earth-bermed?
|       +-- Yes: XPS on exterior + concrete/ICF wall
|       +-- No: Super-insulate to RSI 5.3+ (R-30+)
|           |
|           +-- Budget: Double-stud with mineral wool
|           +-- Performance: Closed-cell spray foam
|           +-- Sustainable: Hempcrete or straw bale
|
+-- Roof/Ceiling (insulated section)
|   |
|   +-- Zone 4-5: RSI 7.0 (R-40) minimum
|   +-- Zone 6-7: RSI 8.8-10.6 (R-50-60) recommended
|   +-- Materials: Spray foam, mineral wool, or rigid foam
|
+-- Interior Walls (partitions)
|   |
|   +-- Low humidity zone: Fiberglass or mineral wool
|   +-- High humidity zone: Mineral wool only
|   +-- Fire separation: Mineral wool (non-combustible)
|
+-- Exterior Walls (above grade, not north)
    |
    +-- High humidity concern?
        +-- Yes: Closed-cell spray foam OR mineral wool
        +-- No: Any appropriate material
    |
    +-- Budget level?
        +-- Low: Mineral wool batts + exterior rigid foam
        +-- Medium: Mineral wool + continuous XPS
        +-- High: Closed-cell spray foam throughout
```

### Climate Zone Recommendations

#### Zone 5 (Southern Ontario, Lower BC Mainland)
- **Walls**: RSI 3.5-4.0 (R-20-23) minimum, RSI 5.0 (R-28) recommended
- **Insulated Roof**: RSI 7.0 (R-40) minimum
- **Foundation**: RSI 1.8-2.1 (R-10-12)
- **Suggested Assembly**: Mineral wool batts + 25-50mm XPS exterior

#### Zone 6 (Central Ontario, Southern Prairies)
- **Walls**: RSI 4.0-5.0 (R-23-28) minimum, RSI 6.0 (R-34) recommended
- **Insulated Roof**: RSI 8.8 (R-50) minimum
- **Foundation**: RSI 2.5-3.0 (R-14-17)
- **Suggested Assembly**: Closed-cell spray foam or double-stud with mineral wool

#### Zone 7A (Northern Ontario, Central Prairies, Edmonton)
- **Walls**: RSI 5.0-6.0 (R-28-34) minimum, RSI 7.0 (R-40) recommended
- **Insulated Roof**: RSI 10.6 (R-60) minimum
- **Foundation**: RSI 3.0-3.5 (R-17-20)
- **Suggested Assembly**: Closed-cell spray foam throughout or ICF + exterior rigid

#### Zone 7B+ (Northern Regions)
- **Walls**: RSI 6.0-7.0 (R-34-40) minimum
- **Insulated Roof**: RSI 10.6+ (R-60+)
- **Foundation**: RSI 3.5+ (R-20+)
- **Suggested Assembly**: ICF or super-insulated double-stud, closed-cell spray foam

---

## Data Model

```go
package models

// InsulationMaterial represents a type of insulation
type InsulationMaterial struct {
    ID                   string
    Name                 string                // "closed_cell_spray_foam"
    DisplayName          string                // "Closed-Cell Spray Foam"
    Category             InsulationCategory    // "spray_foam", "batt", "rigid", "natural"
    KValue               float64               // W/(m-K) - thermal conductivity
    KValueLongTerm       float64               // W/(m-K) - aged/derated value
    AvailableThicknesses []ThicknessOption
    Properties           MaterialProperties
    CostTier             CostTier              // "low", "medium", "high"
    CostPerRSI           float64               // $/m2 per RSI unit
    DIYFriendly          bool
    BestFor              []string              // ["north_walls", "foundation"]
    Cautions             []string              // ["not_for_below_grade", "cold_temp_penalty"]
    LifespanYears        int
    ThermalDrift         bool                  // true if R-value degrades over time
    ColdTempPenalty      bool                  // true if R-value drops in cold
}

type ThicknessOption struct {
    Millimeters    int
    RSI            float64   // m2-K/W (metric R-value)
    RValueImperial float64   // ft2-F-hr/BTU (for display)
    RSILongTerm    float64   // Aged/derated metric R-value
}

type MaterialProperties struct {
    VaporPermeance     float64  // perms
    VaporBarrier       bool     // true if <1 perm at typical thickness
    AirBarrier         bool     // true if air impermeable
    MoistureResistance string   // "poor", "good", "excellent"
    MoistureAbsorption float64  // % by volume
    FireResistance     string   // "poor", "good", "excellent", "non_combustible"
    FireRating         string   // e.g., "Class 1", "2-hour"
    SoundInsulation    string   // "poor", "good", "excellent"
    CompressiveStrength float64 // kPa
    BelowGrade         bool     // suitable for below-grade use
    HighHumidity       bool     // suitable for high-humidity greenhouse
}

type InsulationCategory string

const (
    CategorySprayFoam   InsulationCategory = "spray_foam"
    CategoryBatt        InsulationCategory = "batt"
    CategoryRigid       InsulationCategory = "rigid"
    CategoryReflective  InsulationCategory = "reflective"
    CategoryNatural     InsulationCategory = "natural"
)

type CostTier string

const (
    CostLow    CostTier = "low"
    CostMedium CostTier = "medium"
    CostHigh   CostTier = "high"
)

// InsulationSelection represents user's selection for a wall/surface
type InsulationSelection struct {
    Surface         string  // "wall_north", "wall_south", "roof_north", "foundation"
    MaterialID      string
    ThicknessMM     int
    RSI             float64 // Calculated m2-K/W
    RValueImperial  float64 // Calculated ft2-F-hr/BTU
    EffectiveRSI    float64 // After thermal bridging adjustment
    IsRecommended   bool    // Meets zone requirements
}

// ClimateZone represents Canadian climate zone data
type ClimateZone struct {
    ID                string  // "zone_5", "zone_6", etc.
    Name              string  // "Zone 5 (Southern Ontario)"
    HDDMin            int     // Minimum heating degree days
    HDDMax            int     // Maximum heating degree days
    FrostDepthM       float64 // Typical frost depth in meters
    DesignTempC       float64 // Winter design temperature
    MinWallRSI        float64 // Code minimum wall RSI
    MinRoofRSI        float64 // Code minimum roof RSI
    MinFoundationRSI  float64 // Code minimum foundation RSI
    RecommendedWallRSI float64 // Recommended for greenhouses
}

// ThermalBridgingFactor represents framing impact on R-value
type ThermalBridgingFactor struct {
    FrameType       string  // "wood_2x4_16oc", "wood_2x6_24oc", "steel_6in_16oc"
    FramingFactor   float64 // Percentage reduction (0.20 = 20%)
    Description     string
}
```

---

## API Endpoints

| Method | Path | Handler | Description |
|--------|------|---------|-------------|
| GET | `/api/insulation/materials` | `listMaterials` | List all materials with properties |
| GET | `/api/insulation/materials/{id}` | `getMaterial` | Get detailed material information |
| GET | `/api/insulation/thicknesses` | `getThicknesses` | Get thicknesses for selected material |
| GET | `/api/insulation/recommendations` | `getRecommendations` | Zone-based material recommendations |
| GET | `/api/insulation/zones` | `getClimateZones` | List climate zones with requirements |
| POST | `/api/insulation/calculate` | `calculateRValue` | Calculate R-value from selection |
| POST | `/api/insulation/effective` | `calculateEffective` | Calculate effective R-value with thermal bridging |

---

## References

### R-Value and Testing Standards
- [R-value (insulation) - Wikipedia](https://en.wikipedia.org/wiki/R-value_(insulation))
- [ASTM C177 Standard Test Method - Guarded Hot Plate](https://store.astm.org/c0177-19.html)
- [ASTM C518 Standard Test Method - Heat Flow Meter](https://store.astm.org/c0518-21.html)
- [16 CFR 460 - FTC R-Value Rule](https://www.law.cornell.edu/cfr/text/16/460.5)

### Spray Foam Insulation
- [Johns Manville - Open-Cell vs Closed-Cell Spray Foam](https://www.jm.com/en/blog/2020/june/spray-foam--open-cell-vs--closed-cell/)
- [Building Science Corporation - Spray Polyurethane Foam Vapor Retarders](https://buildingscience.com/documents/reports/rr-0912-spray-polyurethane-foam-need-for-vapor-retarders-in-above-grade-walls/view)
- [Fine Homebuilding - Closed-Cell Foam R-Value](https://www.finehomebuilding.com/project-guides/insulation/closed-cell-foam-studs-waste)

### Rigid Foam Insulation
- [Green Insulation Group - Polyiso vs XPS vs EPS](https://greeninsulationgroup.com/pros-and-cons-of-foam-insulation-polyiso-vs-xps-vs-eps/)
- [Green Insulation Group - XPS EPS Polyiso Moisture Performance](https://greeninsulationgroup.com/how-xps-eps-and-polyiso-insulation-perform-against-moisture/)
- [Ecohome - Thermal Drift in Foam Insulation](https://www.ecohome.net/en/guides/4129/thermal-drift-in-foam-insulation-understanding-r-value-degradation-in-polyiso-xps/)
- [DuPont Styrofoam - XPS vs Polyiso](https://www.dupont.com/building/XPS-vs-ISO.html)

### Mineral Wool
- [GreenBuildingAdvisor - Mineral Wool Insulation](https://www.greenbuildingadvisor.com/article/mineral-wool-insulation)
- [Bob Vila - Rockwool Insulation](https://www.bobvila.com/articles/rockwool-insulation/)
- [Attainable Home - Mineral Wool R-Value](https://www.attainablehome.com/what-is-the-r-value-of-rockwool-insulation/)

### Fiberglass and Moisture
- [UMass Extension - Greenhouse Insulation](https://www.umass.edu/agriculture-food-environment/greenhouse-floriculture/fact-sheets/insulation-know-how)
- [Hansen Buildings - Why Fiberglass Insulation Doesn't Work](https://www.hansenpolebuildings.com/2023/01/why-fiberglass-insulation-doesnt-work/)

### Natural Insulation
- [Ecohome - Straw Bale Homes Guide](https://www.ecohome.net/en/guides/4119/straw-bale-homes-a-complete-guide-for-homeowners-professionals/)
- [Sustainable Build Consultancy - Strawbale Insulation](https://www.sustainablebuildconsultancy.com/blog/how-insulating-is-strawbale)
- [Materials Palette - Hempcrete](https://www.materialspalette.org/hempcrete/)
- [Hempcrete Wikipedia](https://en.wikipedia.org/wiki/Hempcrete)
- [Homeland Hempcrete - R-Value Testing](https://www.homelandhempcrete.com/hempcretenews/blog-rvaluetesting)

### Reflective Barriers
- [Radiant Barrier - Greenhouse Applications](https://radiantbarrier.com/applications-greenhouse-insulation/)
- [Growing Spaces - Reflectix Insulation](https://growingspaces.com/geodesic-dome-greenhouse/insulation/)

### Canadian Building Codes
- [NRC - National Building Code of Canada 2020](https://nrc.canada.ca/en/certifications-evaluations-standards/codes-canada/codes-canada-publications/national-building-code-canada-2020)
- [NAIMA Canada - Insulation Requirements](https://www.naimacanada.ca/insulation-requirements/)
- [NAIMA Canada - Codes and Standards](https://www.naimacanada.ca/codes-standards/)
- [Alberta Energy Codes](https://www.alberta.ca/energy-codes)
- [BC Building Code Energy Efficiency Guide](https://www2.gov.bc.ca/assets/gov/farming-natural-resources-and-industry/construction-industry/building-codes-and-standards/guides/climatezone4.pdf)
- [Quebec Insulation Standards](https://www.quebec.ca/en/housing-territory/heating-energy-consumption/reducing-energy-consumption-at-home/building-insulation-management/insulation)

### Thermal Bridging
- [Sustainability Workshop - Total R-Values and Thermal Bridging](https://sustainabilityworkshop.venturewell.org/node/1027.html)
- [Construction Canada - Nominal vs Effective R-Values](https://www.constructioncanada.net/the-language-of-r-values-understanding-differences-between-nominal-and-effective/)
- [GreenBuildingAdvisor - What is Thermal Bridging](https://www.greenbuildingadvisor.com/article/what-is-thermal-bridging)
- [Ecohome - Thermal Bridge Prevention](https://www.ecohome.net/en/guides/2262/what-is-a-thermal-bridge-and-why-is-it-so-important-to-break-it/)

### Greenhouse Design
- [Ceres Greenhouse - Foundation Insulation](https://ceresgs.com/solar-greenhouse-basics-insulating-your-foundation/)
- [Verge Permaculture - Passive Solar Greenhouse](https://vergepermaculture.ca/designing-your-own-passive-solar-greenhouse-part-3/)
- [GreenBuildingAdvisor - Passive Solar Greenhouse Insulation](https://www.greenbuildingadvisor.com/question/insulation-techinques-for-a-passive-solar-greenhouse)
- [Medium - Passive Solar Greenhouse Design](https://medium.com/@rob_74123/how-to-design-a-passive-solar-greenhouse-light-insulation-and-subterranean-heating-and-cooling-7c66a27afd29)

### Frost-Protected Foundations
- [Rmax - Frost Protected Shallow Foundations](https://www.rmax.com/blog/frost-protected-shallow-foundations)
- [Fine Homebuilding - Frost-Protected Shallow Foundations](https://www.finehomebuilding.com/2010/11/11/frost-protected-shallow-foundations-2)
- [FPSF Guide Canada (PDF)](https://solutions.ca/anhwp/Docs/RevisedFPSFGuide_Sept2004.pdf)

### Cost and Energy Analysis
- [GTA Foaming - Spray Foam vs Batt Cost 2025](https://www.gtafoaming.com/blog/spray-foam-vs-batt-insulation-cost-2025)
- [RenoQuotes - Insulation Costs](https://renoquotes.com/en/blog/insulation-cost)
- [Energy Star - Insulation Savings Methodology](https://www.energystar.gov/saveathome/seal_insulate/methodology)

### Moisture Management
- [Building Science Corporation - Condensation Control](https://buildingscience.com/documents/digests/bsd-controlling-cold-weather-condensation-using-insulation)
- [MyGreenhouseStore - Reduce Greenhouse Condensation](https://mygreenhousestore.com/blogs/news/how-can-i-reduce-condensation-in-my-greenhouse)
- [DOE - Moisture Control](https://www.energy.gov/energysaver/moisture-control)
