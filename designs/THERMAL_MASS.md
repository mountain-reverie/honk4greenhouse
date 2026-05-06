# Thermal Mass for Greenhouse Heat Storage

## Overview

Thermal mass refers to materials that absorb, store, and release heat to moderate temperature fluctuations in a greenhouse. By capturing excess solar energy during the day and releasing it at night, thermal mass reduces heating costs, prevents overheating, and creates a more stable growing environment.

This document provides comprehensive technical guidance for selecting, sizing, and integrating thermal mass systems in greenhouses designed for Canadian climates.

**Key Benefits:**
- Reduces daily temperature swings by 5-15°C
- Lowers supplemental heating requirements by 20-50%
- Prevents overheating during sunny winter days
- Extends growing seasons by maintaining minimum temperatures
- Reduces energy costs and carbon footprint

**Relationship to Climate Battery Systems:**
Thermal mass and climate battery systems are complementary technologies:
- **Thermal mass** stores heat *inside* the greenhouse structure
- **Climate batteries** store heat *underground* in soil

For optimal performance, both approaches can be combined. See [CLIMATE_BATTERY_DESIGNS.md](CLIMATE_BATTERY_DESIGNS.md) for underground heat storage systems.

---

## Thermal Mass Fundamentals

### Heat Storage Principles

Thermal mass works by exploiting two physical properties:

1. **Sensible Heat Storage**: Heat absorbed causes temperature rise in the material
2. **Latent Heat Storage**: Heat absorbed causes phase change (solid↔liquid) without temperature change

**Sensible Heat Formula:**
```
Q = m × c × ΔT

Where:
Q = Heat stored (Joules or kJ)
m = Mass of material (kg)
c = Specific heat capacity (J/kg·K or kJ/kg·K)
ΔT = Temperature change (°C or K)
```

**Practical Example:**
A 200-liter (200 kg) water barrel heated from 15°C to 25°C stores:
```
Q = 200 kg × 4.18 kJ/kg·K × 10K = 8,360 kJ = 2.32 kWh
```

### Key Thermal Properties

| Property | Symbol | Units | Description |
|----------|--------|-------|-------------|
| Specific Heat Capacity | c | J/(kg·K) | Heat to raise 1 kg by 1°C |
| Volumetric Heat Capacity | Cv | MJ/(m³·K) | Heat to raise 1 m³ by 1°C |
| Thermal Conductivity | k | W/(m·K) | Rate of heat flow through material |
| Thermal Diffusivity | α | m²/s | Rate of temperature change propagation |

**Volumetric Heat Capacity Formula:**
```
Cv = ρ × c

Where:
Cv = Volumetric heat capacity (J/m³·K)
ρ = Density (kg/m³)
c = Specific heat capacity (J/kg·K)
```

### Diurnal Cycle Performance

For effective thermal mass performance in greenhouses:

1. **Charge Phase** (daytime): Material absorbs solar heat, temperature rises
2. **Discharge Phase** (nighttime): Material releases stored heat, temperature falls
3. **Cycle Period**: Ideally 12-24 hours for greenhouse applications

**Design Goal:** Size thermal mass to absorb/release close to full capacity within a single day-night cycle. Oversized mass won't charge fully; undersized mass won't moderate temperatures adequately.

---

## Material Properties Comparison

### Thermal Properties Table

| Material | Specific Heat c (J/kg·K) | Density ρ (kg/m³) | Volumetric Heat Capacity Cv (MJ/m³·K) | Thermal Conductivity k (W/m·K) |
|----------|--------------------------|-------------------|---------------------------------------|-------------------------------|
| **Water** | 4,184 | 1,000 | 4.18 | 0.60 |
| **Concrete** | 880 | 2,300 | 2.02 | 1.0-1.4 |
| **Dense Rock** | 800-900 | 2,500-2,800 | 2.0-2.5 | 1.5-3.5 |
| **Gravel (loose)** | 840 | 1,500-1,700 | 1.3-1.4 | 0.3-0.5 |
| **Brick** | 900 | 1,800-2,000 | 1.6-1.8 | 0.6-1.0 |
| **Adobe/Earth** | 860 | 1,500-1,800 | 1.3-1.5 | 0.5-0.8 |
| **Sand (dry)** | 830 | 1,500 | 1.25 | 0.25 |
| **Sand (moist)** | 1,000 | 1,700 | 1.7 | 1.5-2.0 |
| **PCM (salt hydrate)** | 2,000* | 1,500-1,800 | 3.0-3.6* | 0.5-1.0 |
| **PCM (organic/paraffin)** | 2,000* | 800-900 | 1.6-1.8* | 0.2-0.3 |

*\*PCM values include latent heat contribution over operating range*

### Heat Storage Capacity Comparison

To store 10 kWh of thermal energy with a 10°C temperature swing:

| Material | Volume Required | Mass Required | Notes |
|----------|-----------------|---------------|-------|
| Water | 0.86 m³ (860 L) | 860 kg | Most compact |
| Concrete | 1.78 m³ | 4,100 kg | 2x water volume |
| Rock | 1.5-2.0 m³ | 3,750-5,000 kg | Varies with type |
| PCM (Q25) | 0.4-0.6 m³ | 600-900 kg | Most compact, highest cost |

**Key Insight:** Water stores approximately twice as much heat per unit volume as concrete or rock, making it the most space-efficient option for greenhouse thermal mass.

### Material Selection Guide

| Material | Best For | Avoid When |
|----------|----------|------------|
| **Water** | Maximum storage/space, DIY builds, retrofits | Freezing risk, structural concerns |
| **Concrete** | New construction, flooring, north walls | Retrofits, limited budget |
| **Rock/Gravel** | Underground beds, floor systems, low cost | Space-constrained, above-ground |
| **PCM** | Maximum performance, limited space, precise temp control | Budget-constrained, DIY builds |

---

## Water-Based Thermal Mass

Water is the most common and effective thermal mass material for greenhouses due to its high heat capacity, low cost, and ease of installation.

### Container Options

#### 55-Gallon (208 L) Drums

The standard choice for greenhouse thermal mass.

**Specifications:**
| Parameter | Value |
|-----------|-------|
| Capacity | 208 L (55 US gallons) |
| Dimensions | 57 cm dia × 88 cm tall |
| Weight (full) | 210-215 kg |
| Heat capacity | 8.7 MJ per 10°C swing |
| Surface area | ~1.1 m² |

**Types:**
- **Steel drums**: More durable, better heat transfer, stackable (up to 3 high), ~$30-60 CAD
- **HDPE plastic drums**: Lighter, won't rust, food-grade available, ~$15-40 CAD
- **Blue poly drums**: Common, inexpensive, UV-resistant, ~$10-25 CAD (used)

**Pros:** Widely available, inexpensive, easy to move and arrange
**Cons:** Lower surface area per volume than smaller containers

#### IBC Totes (275 Gallon / 1,000 L)

Intermediate Bulk Containers offer higher capacity with efficient footprint.

**Specifications:**
| Parameter | Value |
|-----------|-------|
| Capacity | 1,000 L (275 US gallons) |
| Dimensions | 122 cm L × 102 cm W × 117 cm H |
| Weight (full) | ~1,050 kg |
| Heat capacity | 41.8 MJ per 10°C swing |
| Surface area | ~4.5 m² |

**Types:**
- **Food-grade HDPE**: New ~$200-400 CAD, reconditioned ~$75-150 CAD
- **Chemical-grade**: Lower cost but verify previous contents
- **Stainless steel**: Best heat transfer, ~$500-1,000 CAD

**Pros:** High capacity, stackable (4 high when full), forklift-compatible, integrated valve
**Cons:** Large footprint, expensive new, may need painting for UV protection

#### Water Tubes/Columns

Vertical tubes maximize surface area for faster heat exchange.

**Specifications:**
| Size | Capacity | Height | Diameter | Heat Capacity/10°C |
|------|----------|--------|----------|-------------------|
| 12" × 4' | 170 L (45 gal) | 122 cm | 30 cm | 7.1 MJ |
| 12" × 6' | 255 L (67 gal) | 183 cm | 30 cm | 10.7 MJ |
| 18" × 4' | 380 L (100 gal) | 122 cm | 46 cm | 15.9 MJ |

**Types:**
- **Fiberglass tubes**: Purpose-built, UV stable, ~$100-200 each
- **PVC pipe (capped)**: DIY option, requires sealing, ~$30-80/tube
- **Poly film tubes**: Lowest cost, least durable, ~$5-15/tube

**Pros:** High surface-area-to-volume ratio, fast heat exchange, vertical space efficient
**Cons:** Less stable, may tip, limited availability in Canada

#### Aquaponics Tanks / Stock Tanks

Multi-purpose option combining thermal mass with growing systems.

**Types:**
- **Round stock tanks**: 100-1,500 gallon, galvanized steel, ~$200-800 CAD
- **Aquaponics tanks**: HDPE, various shapes, ~$300-1,000 CAD
- **Swimming pool/hot tub**: Repurposed, varies widely

**Pros:** Dual function (thermal mass + growing), large capacity
**Cons:** Takes floor space, may shade plants, requires pump/aeration

### Water Container Comparison

| Container | Capacity (L) | Cost (CAD) | Surface Area (m²) | SA/Volume Ratio |
|-----------|--------------|------------|-------------------|-----------------|
| 55-gal drum | 208 | $15-60 | 1.1 | 5.3 |
| 275-gal IBC | 1,000 | $75-400 | 4.5 | 4.5 |
| 12" × 6' tube | 255 | $30-200 | 1.7 | 6.7 |
| 4 × 1-gal jugs | 15 | $5 | 0.45 | 30.0 |

**Insight:** Smaller containers have higher surface-area-to-volume ratios and exchange heat faster, but require more labor to install and maintain.

### Container Preparation

**Painting for Heat Absorption:**
1. Clean container surface thoroughly
2. Apply primer suitable for container material (steel/plastic)
3. Paint with flat/matte black paint (absorbs ~95% of solar radiation)
4. Allow full cure before filling with water

**Paint Options:**
- Flat black latex paint: Easy application, may fade, ~$30/gallon
- BBQ/stove paint: Heat resistant, durable, ~$15/can
- Selective absorber paint: Premium absorption, expensive, ~$80+/quart

**Water Treatment:**
Add algaecide or bleach to prevent algae growth:
- Pool algaecide: 30 mL per 200 L, reapply annually
- Household bleach: 15 mL per 200 L, reapply monthly
- Copper sulfate: 5 g per 200 L, long-lasting

### Stacking and Safety

**Critical Safety Requirements:**

Water barrels are extremely heavy—a full 55-gallon drum weighs approximately 210 kg (460 lbs). Improper stacking can cause catastrophic failure.

**Stacking Guidelines:**
| Container Type | Maximum Stack | Foundation Required |
|----------------|---------------|---------------------|
| Steel drums | 3 high | Level concrete, 15 cm thick |
| Plastic drums | 2 high | Level concrete or flagstone |
| IBC totes | 4 high (per rating) | Level concrete, structural review |
| Water tubes | Not stackable | Secure to wall/frame |

**Foundation Requirements:**
- Pour minimum 10 cm (4") concrete pad for heavy thermal mass
- Extend pad 30 cm beyond barrel footprint
- Compact subgrade thoroughly before pouring
- Level to within 5 mm over pad area
- Allow 28 days cure before loading

**Securing Barrels:**
- Use ratchet straps between stacked drums
- Install angle-iron or lumber frame around perimeter
- Anchor frame to foundation with concrete screws
- Never stack on uneven surfaces

---

## Rock and Gravel Thermal Mass

Rock-based thermal mass is cost-effective for underground applications and integrates well with climate battery systems.

### Rock Types and Properties

| Rock Type | Density (kg/m³) | Specific Heat (J/kg·K) | Cv (MJ/m³·K) | Conductivity (W/m·K) |
|-----------|-----------------|------------------------|--------------|----------------------|
| Basalt | 2,800-3,000 | 840-900 | 2.4-2.7 | 1.5-2.5 |
| Granite | 2,600-2,700 | 790-830 | 2.0-2.2 | 2.5-3.5 |
| Limestone | 2,400-2,700 | 840-910 | 2.0-2.5 | 1.3-1.7 |
| Sandstone | 2,200-2,500 | 770-920 | 1.7-2.3 | 1.5-4.0 |
| River rock (rounded) | 2,500-2,700 | 800-850 | 2.0-2.3 | 1.5-2.5 |

**Recommended:** Washed river rock (rounded) is ideal for greenhouse applications—good thermal properties, excellent airflow between stones, widely available.

### Rock Bed Design

#### Under-Floor Rock Beds

Used in Deep Winter Greenhouse (DWG) designs and climate battery systems.

**Specifications:**
| Parameter | Recommended Value | Notes |
|-----------|-------------------|-------|
| Rock size | 40-100 mm (1.5-4") | Smaller = more surface area, more resistance |
| Bed depth | 0.6-1.2 m (2-4 ft) | Deeper = more capacity, higher cost |
| Void fraction | 35-45% | Air space between rocks |
| Coverage | Full floor area | Maximize thermal contact |

**Cross-Section:**
```
    Greenhouse Floor
    ═══════════════════════════════════
    [Vapor barrier - 6 mil poly]
    ───────────────────────────────────
    |                                 |
    |     Washed River Rock           |  ← 40-100mm diameter
    |     (40-45% void space)         |
    |                                 |
    0.6-1.2m ~~~~~~~~~~~~~~~~~~~~~~~~~~~  ← Air distribution (optional pipes)
    |                                 |
    |     Rock Bed                    |
    |                                 |
    ───────────────────────────────────
    [Perimeter insulation R-10]
    ═══════════════════════════════════
    Undisturbed soil
```

**Heat Capacity Calculation:**
```
Q_bed = V_rock × ρ_rock × c_rock × (1 - void_fraction) × ΔT

Example: 50 m² floor × 0.8 m deep × (1-0.40) void
Rock volume = 50 × 0.8 × 0.60 = 24 m³
Heat capacity = 24 m³ × 2,500 kg/m³ × 840 J/kg·K × 10K
             = 504 MJ = 140 kWh per 10°C swing
```

#### Perimeter Rock Beds

Rock-filled trenches around greenhouse perimeter provide insulation and thermal storage.

**Configuration:**
- Trench width: 0.6-1.0 m
- Trench depth: Below frost line (1.5-2.5 m in Canada)
- Fill with washed gravel/rock
- Insulate top surface with EPS foam

### Gravel Floors

Gravel greenhouse floors provide modest thermal mass plus excellent drainage.

**Specifications:**
| Parameter | Value |
|-----------|-------|
| Gravel size | 10-20 mm (pea gravel) |
| Depth | 10-15 cm |
| Heat capacity | ~0.5 MJ/m² per 10°C |

**Pros:** Drainage, weed suppression, some thermal storage, low cost
**Cons:** Limited heat capacity compared to other options

---

## Concrete and Masonry Thermal Mass

Concrete and masonry provide structural thermal mass integrated into greenhouse construction.

### Concrete Floors

A concrete greenhouse floor serves as both structural foundation and thermal mass.

**Specifications:**
| Parameter | Recommended | Notes |
|-----------|-------------|-------|
| Thickness | 10-15 cm (4-6") | Thicker = more mass |
| Heat capacity | 2.0-3.0 MJ/m² per 10°C | Depends on thickness |
| Thermal lag | 4-8 hours | Heat travels ~25 mm/hour |
| Insulation | R-10 minimum below | Essential in cold climates |

**Heat Capacity per Square Meter:**
```
Q/A = thickness × ρ × c × ΔT

10 cm slab: 0.10 m × 2,300 kg/m³ × 880 J/kg·K × 10K = 2.0 MJ/m²
15 cm slab: 0.15 m × 2,300 kg/m³ × 880 J/kg·K × 10K = 3.0 MJ/m²
```

**Insulation Configuration:**
```
    Greenhouse interior
    ═══════════════════════════
    [Concrete slab 10-15 cm]
    ═══════════════════════════
    [Vapor barrier]
    ───────────────────────────
    [EPS/XPS insulation R-10+]
    ───────────────────────────
    Compacted gravel 10 cm
    ───────────────────────────
    Undisturbed soil
```

**Critical:** Without insulation below the slab, heat will conduct into the earth rather than back into the greenhouse.

### Concrete Block (CMU) Walls

Concrete masonry units filled with grout provide excellent thermal mass for north walls.

**Specifications:**
| Block Type | Density (hollow) | Density (grouted) | Cv (MJ/m³·K) |
|------------|------------------|-------------------|--------------|
| Standard 8" CMU | ~1,100 kg/m³ | ~2,100 kg/m³ | 1.0 / 1.8 |
| Standard 12" CMU | ~1,200 kg/m³ | ~2,000 kg/m³ | 1.1 / 1.8 |

**Grouting Recommendation:** Always fill CMU cores with grout or sand for maximum thermal mass. Hollow blocks have less than half the heat capacity.

### Trombe Walls

A Trombe wall is a passive solar heating system using a massive south-facing wall behind glazing.

**Design Principles:**
1. Dark-colored masonry wall absorbs solar radiation
2. Air gap between wall and glazing creates convective loop
3. Heat conducts through wall over 8-12 hours
4. Interior surface radiates heat into greenhouse at night

**Specifications:**
| Parameter | Value | Notes |
|-----------|-------|-------|
| Wall thickness | 20-40 cm (8-16") | Thicker = longer lag time |
| Wall material | Concrete, CMU (grouted), adobe, brick | Dense materials preferred |
| Air gap | 3-6 cm (1-2.5") | Between wall and glazing |
| Glazing | Double-pane glass or twin-wall PC | Single glazing loses too much heat |
| Thermal lag | 8-12 hours | For 20 cm concrete wall |
| Selective surface | Optional | Black paint or selective absorber coating |

**Cross-Section:**
```
    Exterior (south-facing)
           ↓ Solar radiation
    ═══════════════════════════
    [Double glazing]
    ───────────────────────────
        Air gap 3-6 cm
        ↑↓ Convection
    ───────────────────────────
    ███████████████████████████
    ███  Concrete/Masonry   ███  ← 20-40 cm thick
    ███  (dark surface)     ███    Heat flows →→→
    ███████████████████████████
    ───────────────────────────
        Radiant heat →
    ═══════════════════════════
    Greenhouse interior
```

**Performance:**
- Heat transfer rate: ~1 inch (25 mm) per hour through concrete
- 20 cm wall: Heat absorbed at noon reaches interior by 8 PM
- Can provide 30-70% of heating needs in sunny climates
- Less effective in cloudy climates (requires solar input)

**Vented vs. Unvented:**

| Configuration | Daytime | Nighttime | Best For |
|---------------|---------|-----------|----------|
| Unvented | Wall absorbs heat | Radiant release only | Consistent heating |
| Vented (top/bottom openings) | Convective + absorption | Close vents to prevent reverse flow | Faster daytime heating |

### Adobe and Rammed Earth

Traditional earth construction provides thermal mass with very low embodied energy.

**Properties:**
| Material | Density (kg/m³) | Cv (MJ/m³·K) | Conductivity (W/m·K) |
|----------|-----------------|--------------|----------------------|
| Adobe | 1,500-1,700 | 1.3-1.5 | 0.5-0.8 |
| Rammed earth | 1,800-2,000 | 1.5-1.7 | 0.8-1.0 |
| Compressed earth block | 1,700-1,900 | 1.4-1.6 | 0.7-0.9 |

**Pros:** Low cost, local materials, high embodied carbon sequestration
**Cons:** Labor intensive, moisture sensitive, lower capacity than concrete

---

## Phase Change Materials (PCM)

Phase change materials store latent heat during melting and release it during solidification, providing 3-5× the heat storage capacity of water per unit volume.

### How PCMs Work

When a PCM reaches its melting point:
1. Material absorbs heat (latent heat of fusion)
2. Temperature remains nearly constant during phase change
3. Stores 100-300 kJ/kg during melting
4. Releases same heat during solidification

**Latent vs. Sensible Heat:**
```
Water: 4.18 kJ/kg·K sensible heat
       + 0 kJ/kg latent heat (liquid state)
       = 4.18 kJ/kg per °C

PCM:   2.0 kJ/kg·K sensible heat
       + 200 kJ/kg latent heat (at phase change)
       = Equivalent to ~100 kJ/kg per °C near melting point
```

### PCM Types for Greenhouses

#### Salt Hydrates

Inorganic PCMs with high latent heat and good thermal conductivity.

**Common Types:**
| Material | Melting Point | Latent Heat | Density | Cost |
|----------|---------------|-------------|---------|------|
| Calcium chloride hexahydrate (CaCl₂·6H₂O) | 29°C | 170-190 kJ/kg | 1,710 kg/m³ | Low |
| Sodium sulfate decahydrate (Glauber's salt) | 32°C | 251 kJ/kg | 1,460 kg/m³ | Very low |
| Sodium acetate trihydrate | 58°C | 226 kJ/kg | 1,450 kg/m³ | Low |

**Pros:** High latent heat, higher thermal conductivity, lower cost, non-flammable
**Cons:** Supercooling, phase segregation over cycles, corrosive to some metals

#### Organic PCMs (Paraffins and Bio-based)

Carbon-based PCMs with excellent cycling stability.

**Common Types:**
| Material | Melting Point | Latent Heat | Density | Cost |
|----------|---------------|-------------|---------|------|
| Paraffin wax (various grades) | 20-40°C | 150-220 kJ/kg | 800-900 kg/m³ | Medium |
| BioPCM® (fatty acid esters) | 20-30°C | 180-220 kJ/kg | 850-950 kg/m³ | Medium-High |
| Coconut oil | 24°C | 105 kJ/kg | 920 kg/m³ | Low |

**Pros:** No supercooling, excellent cycling stability, non-corrosive, tunable melting points
**Cons:** Lower thermal conductivity, flammable (requires encapsulation), lower volumetric capacity

### Commercial PCM Products

#### BioPCM® (Phase Change Energy Solutions)

Bio-based PCM derived from agricultural byproducts.

**Product Specifications:**
| Product | Melting Point | Heat Storage | Format |
|---------|---------------|--------------|--------|
| BioPCM Q21 | 21°C | ~100 BTU/ft² | Mat (2.7 kg/m²) |
| BioPCM Q23 | 23°C | ~100 BTU/ft² | Mat |
| BioPCM Q25 | 25°C | ~100 BTU/ft² | Mat |
| BioPCM Q27 | 27°C | ~100 BTU/ft² | Mat |

**Cost:** ~$3.00 USD/ft² (~$32 CAD/m²)
**Installation:** Attach to walls or ceiling with staples/adhesive

#### Infinite-R® (Insolcorp)

Salt-hydrate PCM in pouches.

**Product Specifications:**
| Product | Melting Point | Heat Storage |
|---------|---------------|--------------|
| Infinite-R 21 | 21°C | 1.1 MJ/m² (0.3 kWh/m²) |
| Infinite-R 23 | 23°C | 1.1 MJ/m² |
| Infinite-R 25 | 25°C | 1.1 MJ/m² |

**Cost:** ~$4-6 USD/ft² (~$45-65 CAD/m²)

#### PCM Products Ltd (UK)

Industrial PCM solutions including encapsulated balls and flat panels.

**Products:**
- FlatICE® panels: Modular panels for wall/ceiling installation
- TubeICE®: Encapsulated tubes for tank/container installation
- BallICE®: Spherical capsules for bulk storage

### DIY PCM Options

For budget-conscious builders, some PCMs can be sourced affordably:

| Material | Melting Point | Source | Approximate Cost |
|----------|---------------|--------|------------------|
| Paraffin wax (bulk) | 25-30°C | Candle supply | $3-5 CAD/kg |
| Coconut oil | 24°C | Food supply | $8-12 CAD/kg |
| Sodium sulfate decahydrate | 32°C | Chemical supply | $1-2 CAD/kg |

**DIY Encapsulation:**
- Fill PVC pipes (capped) with liquid PCM
- Use food-grade plastic pouches (heat-sealed)
- Fill HDPE bottles/jugs

**Caution:** DIY PCM systems require careful material selection and containment to prevent leaks and ensure fire safety.

### PCM Sizing for Greenhouses

**General Guideline:** 1-2 kg of PCM per m³ of greenhouse volume

**Calculation Example:**
```
Greenhouse: 100 m² floor × 3 m average height = 300 m³

PCM requirement: 300 m³ × 1.5 kg/m³ = 450 kg

Using BioPCM Q25 mats (2.7 kg/m²):
Area needed = 450 kg / 2.7 kg/m² = 167 m²

Heat storage capacity:
450 kg × 200 kJ/kg = 90,000 kJ = 25 kWh
```

### PCM Placement Strategies

**North Wall Installation:**
- Mount PCM panels/mats on insulated north wall
- Position to receive reflected/diffused light
- Protects from direct summer sun (prevents overheating PCM)

**Ceiling Installation:**
- Installs between rafters or on ceiling surface
- Captures rising warm air
- Effective for heat capture during sunny days

**Growing Bed Integration:**
- Bury PCM containers beneath growing beds
- Provides root zone temperature moderation
- Requires waterproof encapsulation

### PCM Performance in Greenhouses

**Research Results:**

| Study | Configuration | Results |
|-------|---------------|---------|
| Green Built Alliance (Asheville, NC) | 1,370 lbs salt hydrate PCM in PVC tubes | 30 kWh storage, provides 60 kWh daily cycling |
| ScienceDirect (2022) | PCM north wall | 7°C daytime reduction, 9°C nighttime increase |
| Chinese Solar Greenhouse Study | Composite PCM mortar | 1.5°C average increase, 80% yield improvement |

---

## DIY Phase Change Materials (25-30°C Range)

For greenhouse applications, PCMs with melting points between 25-30°C are ideal—warm enough to capture daytime heat but cool enough to release it when nighttime temperatures drop. This section covers accessible DIY formulations that can be made from readily available materials.

### Option 1: Coconut Oil (Simplest, ~24-26°C)

Coconut oil is the easiest and safest DIY PCM option. It's food-grade, non-toxic, widely available, and has excellent thermal cycling stability.

**Thermal Properties:**
| Property | Value |
|----------|-------|
| Melting point | 24-26°C (varies by grade/purity) |
| Latent heat | 100-105 kJ/kg |
| Specific heat (solid) | 3.2 kJ/kg·K |
| Specific heat (liquid) | 2.4 kJ/kg·K |
| Density | 920 kg/m³ |
| Thermal conductivity | 0.17 W/m·K |
| Cycling stability | Excellent (200+ cycles tested) |

**Advantages:**
- Food-safe and non-toxic
- Available at grocery stores ($8-15 CAD/kg)
- No additives required
- Excellent long-term stability
- Pleasant smell (unlike some chemical PCMs)
- Low fire risk (flash point ~315°C)

**Disadvantages:**
- Melting point slightly below ideal (24°C vs 27-29°C)
- Lower latent heat than salt hydrates
- Higher cost per kJ stored than other DIY options
- Low thermal conductivity (slow heat transfer)

**Best Practices:**
- Use refined (RBD) coconut oil for consistent melting point
- Virgin coconut oil has slightly higher melting point (~26°C)
- Store in opaque containers to prevent UV degradation
- Leave 10% headspace for expansion

**Heat Storage Calculation:**
```
1 kg coconut oil stores:
Latent heat: 103 kJ
Sensible heat (5°C swing): 5 × 2.8 kJ = 14 kJ
Total: ~117 kJ per kg = 0.033 kWh per kg

For 10 kWh storage: ~300 kg coconut oil (~330 liters)
Cost: ~$2,400-4,500 CAD
```

### Option 2: Calcium Chloride Hexahydrate (CaCl₂·6H₂O) (~29°C)

Calcium chloride hexahydrate is a salt hydrate PCM with high latent heat and a melting point of 29°C. It requires additives to prevent supercooling and phase separation but offers excellent heat storage at low cost.

**Thermal Properties:**
| Property | Value |
|----------|-------|
| Melting point | 29-30°C |
| Latent heat | 170-190 kJ/kg |
| Specific heat | 1.4-2.1 kJ/kg·K |
| Density | 1,710 kg/m³ |
| Thermal conductivity | 0.5-1.0 W/m·K |
| Heat storage | ~225 BTU/liter (~237 kJ/L) |

**Basic DIY Formulation:**

Based on US Patent 4,613,444 (George Lane) and subsequent research:

| Component | Amount | Purpose |
|-----------|--------|---------|
| Calcium chloride (CaCl₂) | 49-50 wt% | Base PCM component |
| Distilled water | 49-50 wt% | Forms hexahydrate crystal |
| Strontium chloride hexahydrate (SrCl₂·6H₂O) | 1-2 wt% | Nucleating agent (prevents supercooling) |
| Sodium chloride (NaCl) | 0.5-1 wt% | Stabilizer (optional) |
| Potassium chloride (KCl) | 0.5-1 wt% | Stabilizer (optional) |

**Preparation Steps:**

1. **Source materials:**
   - Calcium chloride: Pool supply store (calcium hardness increaser), ice melt products (DowFlake, PellaDow), or chemical supplier
   - Strontium chloride: Chemical supplier, fireworks/flare supply, or online (less common)
   - Distilled water: Grocery store

2. **Calculate quantities (for 1 kg batch):**
   ```
   Calcium chloride (anhydrous): 500g
   Distilled water: 480g
   Strontium chloride hexahydrate: 20g
   ```

3. **Mixing procedure:**
   - Wear safety glasses and gloves (CaCl₂ is caustic)
   - Slowly add calcium chloride to water in a heat-resistant container
   - **Caution:** Reaction is exothermic—mixture will heat significantly
   - Stir until fully dissolved
   - Add strontium chloride and stir until dissolved
   - Allow to cool to room temperature
   - Transfer to containers while still liquid

4. **Verify formation:**
   - Cool below 20°C—should solidify into white crystalline mass
   - Heat above 30°C—should melt to clear liquid
   - If it doesn't re-solidify (supercooling), add more nucleating agent

**Safety Warnings:**

⚠️ **Critical Safety Information:**

| Hazard | Precaution |
|--------|------------|
| Caustic/corrosive | Wear gloves, safety glasses, work in ventilated area |
| Exothermic mixing | Add CaCl₂ slowly, use heat-resistant container |
| Metal corrosion | Use plastic, glass, or stainless steel containers only |
| Skin irritation | Wash immediately if contact occurs |
| Temperature limit | Never heat above 40°C (destroys nucleation centers) |

**Container Compatibility:**
| Material | Compatible | Notes |
|----------|------------|-------|
| HDPE plastic | Yes | Recommended for DIY |
| Polypropylene | Yes | Good option |
| Glass | Yes | Fragile but inert |
| Stainless steel (304/316) | Yes | Expensive but durable |
| Aluminum | **No** | Corrodes rapidly |
| Carbon steel | **No** | Corrodes rapidly |
| Copper | **No** | Corrodes |

**Cost Analysis:**
```
Materials for 10 kg batch:
- Calcium chloride (5 kg pool grade): $15-25 CAD
- Strontium chloride (200g): $20-40 CAD (harder to source)
- Distilled water (5L): $5-10 CAD

Total: ~$40-75 CAD for 10 kg
Heat storage: 10 kg × 180 kJ/kg = 1,800 kJ = 0.5 kWh

Cost per kWh storage: ~$80-150 CAD
(Compare to commercial PCM: $200-400/kWh)
```

### Option 3: Fatty Acid Eutectic Mixtures (25-28°C)

Eutectic mixtures of fatty acids can be precisely tuned to specific melting points. For the 25-28°C range, capric acid + myristic acid is the best documented option.

**Capric-Myristic Acid Eutectic:**

| Property | Value |
|----------|-------|
| Composition | ~75-80% capric acid, 20-25% myristic acid (by weight) |
| Melting point | 25.6°C |
| Latent heat | 160-205 kJ/kg |
| Density | ~900 kg/m³ |
| Cycling stability | Excellent |

**Alternative Mixtures:**

| Mixture | Ratio (wt%) | Melting Point | Latent Heat |
|---------|-------------|---------------|-------------|
| Capric + Myristic | 75:25 | 25.6°C | 205 kJ/kg |
| Capric + Stearic | 83:17 | 24.8°C | 178 kJ/kg |
| Capric + Palmitic | 76:24 | 22.5°C | 171 kJ/kg |
| Lauric + Myristic | 66:34 | 32.7°C | 145 kJ/kg |

**Sourcing Fatty Acids:**

| Acid | Source | Approximate Cost |
|------|--------|------------------|
| Capric acid (C10:0) | Chemical suppliers, soap-making supply | $20-40 CAD/kg |
| Myristic acid (C14:0) | Chemical suppliers, soap-making supply | $15-30 CAD/kg |
| Lauric acid (C12:0) | Soap-making supply, coconut oil derivative | $10-25 CAD/kg |
| Stearic acid (C18:0) | Craft stores (candle making), soap supply | $8-15 CAD/kg |

**Suppliers (Canada/North America):**
- Voyageur Soap & Candle (voyageursoapandcandle.com) - fatty acids for soap making
- Bulk Apothecary (bulkapothecary.com) - various fatty acids
- Chemical suppliers (Sigma-Aldrich, Fisher Scientific) - lab grade

**Preparation Steps:**

1. **Calculate precise ratio** for desired melting point
2. **Melt components separately** in double boiler (water bath)
3. **Combine and stir thoroughly** while liquid
4. **Test melting point** with thermometer
5. **Adjust ratio** if needed and re-test
6. **Transfer to containers** while liquid

**Advantages:**
- Precise melting point control
- High latent heat
- No supercooling
- Non-corrosive
- Good cycling stability

**Disadvantages:**
- Fatty acids harder to source than coconut oil
- Higher cost than salt hydrates
- May have slight odor
- Flammable (fire precautions needed)

### Option 4: Modified Coconut Oil Blends (26-28°C)

By blending coconut oil with higher-melting-point oils or waxes, the melting point can be raised into the optimal 26-28°C range.

**Coconut + Palm Oil Blend:**

| Blend Ratio | Approximate Melting Point |
|-------------|---------------------------|
| 100% Coconut | 24°C |
| 80% Coconut + 20% Palm | 26°C |
| 70% Coconut + 30% Palm | 27°C |
| 60% Coconut + 40% Palm | 28°C |

**Preparation:**
1. Melt both oils in double boiler
2. Combine in desired ratio
3. Stir thoroughly
4. Allow to cool and verify melting point
5. Adjust ratio as needed

**Note:** These ratios are approximate. Test your specific oil sources, as natural products vary.

### DIY PCM Encapsulation Methods

#### Method 1: PVC Pipe Containers (Recommended)

**Materials:**
- Schedule 40 PVC pipe (3" or 4" diameter)
- PVC end caps
- PVC primer and cement
- PCM material

**Construction:**
```
        [Removable cap or plug]
              │
    ┌─────────┴─────────┐
    │                   │ ← 10% headspace for expansion
    │                   │
    │   PCM Material    │ ← 90% fill
    │                   │
    │                   │
    └─────────┬─────────┘
              │
        [Glued end cap]
```

**Steps:**
1. Cut PVC to desired length (typically 1-2 m)
2. Glue one end cap permanently with PVC cement
3. Fill with liquid PCM to 90% capacity
4. Install removable cap with gasket, or glue second cap
5. Label with material type and melting point
6. Mount horizontally on wall brackets

**Sizing:**
```
3" PVC (76mm ID) × 1m length:
Volume = π × (0.038)² × 1.0 = 0.0045 m³ = 4.5 L
Coconut oil mass: 4.5 L × 0.92 kg/L = 4.1 kg
Heat storage: 4.1 kg × 103 kJ/kg = 422 kJ = 0.12 kWh

For 10 kWh storage: ~85 pipes
```

#### Method 2: HDPE Bottles/Jugs

**Suitable Containers:**
- 1-gallon HDPE jugs (milk/water jugs)
- 5-gallon HDPE pails with lids
- 275-gallon IBC totes (large scale)

**Requirements:**
- Minimum wall thickness: 0.9 mm (0.035")
- Must seal completely (no moisture exchange)
- Leave 10% headspace
- Opaque containers preferred (UV protection)

**Steps:**
1. Clean container thoroughly
2. Fill with liquid PCM to 90% capacity
3. Seal tightly
4. Store horizontally or secure vertically
5. Check for leaks after first thermal cycle

#### Method 3: Vacuum-Sealed Bags (Budget Option)

**Materials:**
- Heavy-duty vacuum seal bags (150+ micron)
- Vacuum sealer
- Rigid outer container for protection

**Steps:**
1. Pour liquid PCM into bag
2. Seal with vacuum sealer (minimal vacuum—just seal)
3. Place in rigid frame or between rigid boards
4. Ensure bag cannot be punctured

**Caution:** Not recommended for salt hydrates (corrosive) or for high-traffic areas.

### DIY PCM Comparison Table

| Material | Melting Point | Latent Heat | Cost/kg | Difficulty | Safety |
|----------|---------------|-------------|---------|------------|--------|
| Coconut oil | 24-26°C | 103 kJ/kg | $8-15 | Easy | Safe |
| Coconut + palm blend | 26-28°C | 95-100 kJ/kg | $8-12 | Easy | Safe |
| CaCl₂·6H₂O | 29°C | 180 kJ/kg | $4-8 | Moderate | Caustic |
| Capric-myristic eutectic | 25.6°C | 205 kJ/kg | $25-40 | Moderate | Flammable |

### DIY PCM Best Practices

**General Guidelines:**

1. **Start small:** Test with 1-2 kg before scaling up
2. **Verify melting point:** Use thermometer to confirm phase transition temperature
3. **Test cycling:** Perform 5-10 melt/freeze cycles before full deployment
4. **Label everything:** Material type, melting point, date prepared
5. **Monitor for degradation:** Check annually for separation, leakage, or performance loss

**Container Placement:**

| Location | Advantage | Considerations |
|----------|-----------|----------------|
| North wall | Direct solar absorption | Paint containers black |
| Beneath benches | Root zone warming | Ensure good thermal contact |
| Suspended near ceiling | Captures rising warm air | Secure mounting essential |
| Surrounding growing beds | Microclimate control | May reduce growing space |

**Thermal Contact Enhancement:**

For low-conductivity PCMs (coconut oil, paraffins):
- Use thin, flat containers to maximize surface area
- Add aluminum fins or mesh to containers
- Ensure good air circulation around containers
- Consider multiple smaller containers vs. few large ones

### DIY PCM Safety Summary

| PCM Type | PPE Required | Fire Risk | Spill Cleanup |
|----------|--------------|-----------|---------------|
| Coconut oil | None | Low | Absorb with paper towels |
| CaCl₂ hexahydrate | Gloves, glasses | None | Absorb, wash with water |
| Fatty acid eutectics | Gloves | Moderate | Absorb, ventilate |

**Emergency Procedures:**

- **Skin contact (CaCl₂):** Wash immediately with plenty of water for 15 minutes
- **Eye contact (CaCl₂):** Flush with water for 15 minutes, seek medical attention
- **Coconut oil spill:** Clean with soap and water
- **Fatty acid spill:** Absorb with sand/vermiculite, ventilate area

---

## Sizing Calculations

### Rule of Thumb Guidelines

**Water Thermal Mass:**
| Climate | Water per m² Glazing | Water per m² Floor |
|---------|----------------------|--------------------|
| Mild (Zone 4) | 8-12 L/m² (2 gal/ft²) | 75-100 L/m² |
| Moderate (Zone 5-6) | 12-16 L/m² (3-4 gal/ft²) | 100-150 L/m² |
| Cold (Zone 7+) | 16-20 L/m² (4-5 gal/ft²) | 150-200 L/m² |

**Masonry Thermal Mass:**
- Approximately 2× the volume of water for equivalent capacity
- Minimum 10 cm thick concrete floor
- North wall: 20-40 cm thick masonry

**PCM Thermal Mass:**
- 1-2 kg per m³ greenhouse volume
- Or 20-40 m² of PCM mat per 100 m² floor area

### Detailed Calculation Method

#### Step 1: Determine Heat Loss Rate

Calculate overnight heat loss using the formula:
```
Q_loss = U × A × ΔT × t

Where:
Q_loss = Heat lost (kJ)
U = Overall heat transfer coefficient (W/m²·K)
A = Surface area (m²)
ΔT = Temperature difference inside to outside (K)
t = Time period (hours) × 3.6 (to convert W·h to kJ)
```

**Example:**
```
100 m² greenhouse, twin-wall polycarbonate (U = 3.0 W/m²·K)
Surface area ≈ 200 m² (walls + roof)
ΔT = 20°C (maintain 10°C inside when -10°C outside)
Duration = 14 hours overnight

Q_loss = 3.0 × 200 × 20 × 14 × 3.6 = 604,800 kJ = 168 kWh
```

#### Step 2: Determine Available Solar Gain

Estimate daytime solar heat gain:
```
Q_solar = I × A_glazing × τ × η

Where:
Q_solar = Solar heat gain (kJ)
I = Solar radiation (kJ/m²/day)
A_glazing = Glazing area (m²)
τ = Glazing transmittance (0.7-0.85)
η = Utilization factor (0.5-0.7)
```

**Example (January, Southern Ontario):**
```
I = 7,200 kJ/m²/day (2 kWh/m²/day average)
A_glazing = 80 m² (south-facing)
τ = 0.80 (twin-wall polycarbonate)
η = 0.60 (utilization factor)

Q_solar = 7,200 × 80 × 0.80 × 0.60 = 276,480 kJ = 77 kWh
```

#### Step 3: Calculate Storage Requirement

Storage needed equals heat loss minus solar gain that can be used directly:
```
Q_storage = Q_loss - Q_direct

Where Q_direct ≈ 0.3 × Q_solar (30% used directly)

Q_storage = 168 - (0.3 × 77) = 145 kWh
```

#### Step 4: Size Thermal Mass

**For Water:**
```
Volume = Q_storage / (ρ × c × ΔT_swing)

Assuming 10°C temperature swing:
Volume = 145 kWh × 3,600 kJ/kWh / (1,000 kg/m³ × 4.18 kJ/kg·K × 10K)
Volume = 522,000 / 41,800 = 12.5 m³ = 12,500 L

That's approximately 60 × 55-gallon drums
```

**For Concrete:**
```
Volume = Q_storage / (ρ × c × ΔT_swing)
Volume = 522,000 / (2,300 × 0.88 × 10) = 25.8 m³

For 15 cm thick floor: Area = 25.8 / 0.15 = 172 m²
```

### Simplified Sizing Tables

#### Water Barrels Needed by Greenhouse Size

| Floor Area | Zone 4 | Zone 5-6 | Zone 7+ |
|------------|--------|----------|---------|
| 25 m² | 8-10 drums | 10-15 drums | 15-20 drums |
| 50 m² | 15-20 drums | 20-30 drums | 30-40 drums |
| 100 m² | 25-35 drums | 40-55 drums | 55-75 drums |
| 200 m² | 45-65 drums | 75-100 drums | 100-140 drums |

*Based on 55-gallon (208 L) drums, 10°C temperature swing*

#### PCM Area Needed by Greenhouse Size

| Floor Area | BioPCM Area | Heat Storage |
|------------|-------------|--------------|
| 25 m² | 15-25 m² | 5-8 kWh |
| 50 m² | 30-50 m² | 10-17 kWh |
| 100 m² | 60-100 m² | 20-33 kWh |
| 200 m² | 120-200 m² | 40-67 kWh |

*Based on BioPCM Q25 mats, ~0.33 kWh/m²*

---

## Placement Strategies

### North Wall Placement

The north wall is the optimal location for thermal mass in passive solar greenhouses:

1. **Direct solar exposure**: South-facing glazing illuminates north wall
2. **No shading**: Doesn't block light to plants
3. **Reflective benefit**: Light-colored mass reflects light back to plants
4. **Structural integration**: Wall mass provides insulation function

**North Wall Configuration:**
```
    South glazing (60-90°)
           \
            \  ↘ Solar radiation
             \
              \________________________
              |                        |
    Plants    |    Thermal mass        |  North wall
              |    (water barrels,     |
              |    concrete, PCM)      |
              |________________________|
                                       |
                                    Insulation
                                    (exterior)
```

**Design Guidelines:**
- Paint containers flat black for maximum absorption
- Leave 5-10 cm air gap between containers and insulation
- Stack drums with tallest in back (tiered arrangement)
- Consider reflective surface above/behind mass (Reflectix®)

### South Wall / Glazing Proximity

Placing thermal mass near south glazing captures maximum solar radiation but may shade plants.

**Strategies:**
- **Low thermal mass**: Rock/gravel beds, concrete floor absorb diffuse radiation
- **Water wall**: Translucent water containers allow some light transmission
- **Selective placement**: Mass in corners or edges, not center

### Floor Integration

Greenhouse floors provide distributed thermal mass without occupying growing space.

**Options:**
| Floor Type | Heat Capacity | Cost | Notes |
|------------|---------------|------|-------|
| Concrete slab (15 cm) | 3.0 MJ/m²·10°C | $50-80/m² | Requires insulation below |
| Rock bed (1 m deep) | 14-18 MJ/m²·10°C | $40-60/m² | Excellent with climate battery |
| Gravel (15 cm) | 0.5 MJ/m²·10°C | $15-25/m² | Limited capacity, good drainage |

### Underground/Perimeter

Underground thermal mass provides long-term heat storage but slower response.

**Applications:**
- Climate battery systems (see [CLIMATE_BATTERY_DESIGNS.md](CLIMATE_BATTERY_DESIGNS.md))
- Perimeter insulation/mass trenches
- Earth-sheltered greenhouse walls

---

## Integration with Climate Battery Systems

Thermal mass and climate battery systems work synergistically:

| System | Response Time | Best For | Location |
|--------|---------------|----------|----------|
| Water barrels | Fast (hours) | Daily cycling | Inside greenhouse |
| PCM | Fast (hours) | Temperature stabilization | Walls/ceiling |
| Concrete floor | Medium (hours) | Steady heat release | Greenhouse floor |
| Climate battery | Slow (days) | Extended cold periods | Underground |

### Combined System Design

**Recommended Configuration:**

1. **Primary (fast response)**: Water barrels on north wall
   - Handles daily temperature swings
   - 50-100 L per m² floor area

2. **Secondary (medium response)**: Concrete floor or rock bed
   - Provides base-level thermal stability
   - 10-15 cm concrete or 30+ cm rock

3. **Tertiary (slow response)**: Climate battery
   - Stores heat for extended cloudy periods
   - Recovers heat over 2-5 days

**Energy Flow Diagram:**
```
    Daytime (Sunny)
    ================
    Solar radiation
          ↓
    +-----+---------+
    |     |         |
    | Water → Fast absorption (1-4 hours)
    |     |         |
    | Floor → Medium absorption (4-8 hours)
    |     |         |
    | Climate → Slow absorption (continuous)
    | Battery       |
    +---------------+

    Nighttime
    =========
    Heat demand
          ↑
    +-----+---------+
    |     |         |
    | Water → Fast release (first 4 hours)
    |     |         |
    | Floor → Medium release (4-8 hours)
    |     |         |
    | Climate → Slow release (maintains minimum)
    | Battery       |
    +---------------+
```

### Sizing for Combined Systems

**Rule of Thumb for Combined Systems:**
- Water/PCM: 25-50% of daily heat loss
- Floor/rock bed: 25-50% of daily heat loss
- Climate battery: 2-5 days of heat reserve

**Example (100 m² greenhouse, Zone 5):**
- Daily heat loss: ~150 kWh
- Water barrels: 50-75 kWh capacity → 25-35 drums
- Concrete floor: 50-75 kWh capacity → 100 m² × 15 cm
- Climate battery: 300-750 kWh → Standard two-trench design

---

## Canadian Climate Considerations

### Climate Zone Recommendations

#### Zone 4 (Coastal BC, Southern Ontario)
- **HDD:** 2,500-4,000
- Thermal mass can provide 60-80% of heating needs
- Water barrels: 8-12 L/m² glazing
- Freezing risk low—standard containers acceptable
- PCM effective year-round

#### Zone 5-6 (Central Ontario, Montreal, Calgary)
- **HDD:** 4,000-6,000
- Thermal mass provides 40-60% of heating needs
- Water barrels: 12-16 L/m² glazing
- Consider freeze protection (insulated covers, heating element backup)
- Combine with climate battery for extended cold snaps

#### Zone 7+ (Northern Ontario, Northern Prairies)
- **HDD:** 6,000+
- Thermal mass provides 20-40% of heating needs
- Water barrels: 16-20 L/m² glazing
- Freeze protection essential
- Climate battery depth must exceed deep frost line
- Consider PCM for faster response during limited solar hours

### Freeze Protection

Water thermal mass is vulnerable to freezing in unheated greenhouses.

**Protection Strategies:**

| Strategy | Effectiveness | Cost | Notes |
|----------|---------------|------|-------|
| Antifreeze additive | High | $50-100 | Reduces heat capacity slightly |
| Aquarium heaters | High | $30-50 each | Requires electricity |
| Insulated covers | Medium | $20-50 | Manual operation |
| Bottom insulation | Medium | $10-20/m² | Prevents ground cooling |
| Oversized mass | Low-Medium | Variable | Large mass resists freezing |

**Antifreeze Options:**
- RV/marine antifreeze (propylene glycol): Non-toxic, -50°C protection
- Calcium chloride solution: Lower cost, corrosive to steel
- Methanol: Effective but toxic and flammable—avoid

**Antifreeze Concentration:**
| Protection Level | Propylene Glycol % | Heat Capacity Reduction |
|------------------|-------------------|------------------------|
| -10°C | 20% | 5% |
| -20°C | 35% | 10% |
| -30°C | 45% | 15% |
| -40°C | 55% | 20% |

### Extended Cloudy Periods

Canadian winters often feature extended overcast conditions where solar gain is minimal.

**Design Considerations:**
- Don't rely solely on thermal mass in cloudy climates
- Combine with backup heating (propane, electric, wood)
- Size climate battery for 3-5 day reserve
- Consider greenhouse orientation and glazing angle for low winter sun

**Regional Solar Availability:**
| Region | Winter Sun Hours | Thermal Mass Effectiveness |
|--------|------------------|---------------------------|
| Southern BC coast | Low (cloudy) | Moderate—backup needed |
| Southern Ontario | Moderate | Good |
| Prairies | High (cold but sunny) | Excellent |
| Northern regions | Low (short days) | Limited—backup essential |

---

## Installation Guidelines

### Water Barrel Installation

**Materials Needed:**
- 55-gallon drums (steel or plastic)
- Flat black paint + primer
- Concrete pad materials (or existing floor)
- Ratchet straps or framing lumber
- Water treatment (algaecide/bleach)
- Hose and funnel for filling

**Installation Steps:**

1. **Prepare foundation**
   - Pour concrete pad or verify floor capacity
   - Level to within 5 mm
   - Allow concrete to cure 7+ days

2. **Prepare containers**
   - Clean thoroughly (especially used containers)
   - Apply primer, then flat black paint
   - Allow full cure (24-48 hours)

3. **Position containers**
   - Place on level foundation
   - Arrange in rows, tallest at back
   - Leave 5 cm gaps for air circulation

4. **Secure containers**
   - Install lumber frame around perimeter
   - Use ratchet straps between stacked drums
   - Anchor frame to floor/foundation

5. **Fill with water**
   - Fill slowly to avoid tipping
   - Add water treatment
   - Leave 5 cm headspace for expansion

6. **Monitor and maintain**
   - Check water levels monthly
   - Reapply treatment annually
   - Inspect for leaks, rust, algae

### Rock Bed Installation

**Materials Needed:**
- Washed river rock (40-100 mm)
- Landscape fabric
- Perimeter insulation (EPS/XPS)
- Vapor barrier (6 mil poly)
- Perforated drain pipe (optional)
- Excavation equipment

**Installation Steps:**

1. **Excavate to depth**
   - Mark area and excavate 0.6-1.2 m deep
   - Maintain level bottom
   - Compact subgrade

2. **Install perimeter insulation**
   - Line vertical walls with EPS/XPS (R-10 minimum)
   - Extend insulation 30 cm below grade outside perimeter

3. **Install drainage (if needed)**
   - Lay perforated pipe at lowest point
   - Route to daylight or dry well

4. **Install air distribution (for climate battery)**
   - Lay perforated pipes in grid pattern
   - Connect to manifolds at surface

5. **Add rock fill**
   - Fill in 30 cm lifts
   - Do not compact (maintain void space)
   - Level top surface

6. **Install vapor barrier**
   - Lay 6 mil poly over rock surface
   - Overlap seams 30 cm
   - Seal to perimeter insulation

7. **Install flooring**
   - Gravel, pavers, or concrete over vapor barrier

### PCM Installation

**Wall-Mounted PCM Mats:**

1. **Prepare wall surface**
   - Ensure clean, dry, smooth surface
   - Install vapor barrier if needed

2. **Attach mounting strips**
   - Use furring strips or direct attachment
   - Space per manufacturer instructions

3. **Install PCM mats**
   - Roll out mats horizontally
   - Staple or adhesive per manufacturer
   - Overlap edges as specified

4. **Install protective covering**
   - Drywall, plywood, or protective fabric
   - Ensure PCM can exchange heat with air

**DIY PCM in PVC Pipes:**

1. **Cut pipes to length** (typically 1.5-2 m)
2. **Cap one end** with PVC cement
3. **Fill with liquid PCM** (heated if necessary)
4. **Cap other end** (leave small air space)
5. **Mount horizontally** on wall supports
6. **Allow access for inspection**

---

## Cost Analysis

### Material Costs (CAD, 2025)

#### Water Thermal Mass

| Item | Unit Cost | Notes |
|------|-----------|-------|
| 55-gal steel drum (new) | $50-80 | Food grade |
| 55-gal steel drum (used) | $20-40 | Verify previous contents |
| 55-gal plastic drum (new) | $40-60 | HDPE, food grade |
| 55-gal plastic drum (used) | $10-25 | Common, inspect for damage |
| 275-gal IBC tote (reconditioned) | $75-150 | Check valve condition |
| 275-gal IBC tote (new) | $250-400 | Food grade HDPE |
| Water tubes (fiberglass) | $100-200 each | 45-100 gallon |
| Black paint (exterior latex) | $30-50/gallon | Covers ~10-15 drums |
| Algaecide | $15-30 | Annual treatment |
| Concrete pad (10 cm) | $80-120/m² | Installed price |

#### Rock/Gravel Thermal Mass

| Item | Unit Cost | Notes |
|------|-----------|-------|
| Washed river rock (40-80mm) | $40-60/m³ | Delivered |
| Pea gravel (10-20mm) | $30-50/m³ | Delivered |
| Landscape fabric | $1-2/m² | Under rock bed |
| EPS insulation (R-10) | $25-40/m² | Perimeter |
| Vapor barrier (6 mil) | $0.50-1/m² | Over rock bed |
| Excavation | $50-100/m³ | Machine excavation |

#### Concrete/Masonry Thermal Mass

| Item | Unit Cost | Notes |
|------|-----------|-------|
| Concrete slab (10 cm) | $60-100/m² | Installed |
| Concrete slab (15 cm) | $80-120/m² | Installed |
| CMU blocks (8" grouted) | $15-25/block | Includes grout |
| CMU blocks (12" grouted) | $20-35/block | Includes grout |
| Insulation below slab (R-10) | $25-40/m² | XPS recommended |

#### Phase Change Materials

| Item | Unit Cost | Notes |
|------|-----------|-------|
| BioPCM mat (Q25) | $30-40/m² | Includes shipping |
| Infinite-R panels | $45-65/m² | Salt hydrate |
| Bulk paraffin wax | $3-5/kg | DIY encapsulation required |
| PVC pipe for DIY PCM | $10-20/m | 4" diameter |

### Total System Cost Estimates

| Greenhouse Size | Water Barrels | Rock Bed | PCM System |
|-----------------|---------------|----------|------------|
| 25 m² | $400-800 | $1,500-2,500 | $1,000-1,500 |
| 50 m² | $700-1,400 | $2,500-4,000 | $1,800-2,800 |
| 100 m² | $1,200-2,500 | $4,500-7,000 | $3,500-5,500 |
| 200 m² | $2,200-4,500 | $8,000-12,000 | $6,500-10,000 |

*Costs include materials and basic installation; professional installation adds 50-100%*

### Cost-Benefit Analysis

**Payback Period Example (100 m² greenhouse, Zone 5):**

| Scenario | Heating Cost/Year | Savings | System Cost | Payback |
|----------|-------------------|---------|-------------|---------|
| No thermal mass (propane) | $3,000 | — | — | — |
| Water barrels (30% reduction) | $2,100 | $900 | $1,800 | 2 years |
| Rock bed + climate battery (60%) | $1,200 | $1,800 | $6,000 | 3.3 years |
| Combined system (75% reduction) | $750 | $2,250 | $8,000 | 3.5 years |

---

## Data Model

```go
// ThermalMassSystem represents the thermal mass configuration for a greenhouse project
type ThermalMassSystem struct {
    ID        string
    ProjectID string

    // Water-based thermal mass
    WaterContainers []WaterContainer

    // Rock/gravel thermal mass
    RockBed *RockBedConfig

    // Concrete/masonry thermal mass
    ConcreteMass *ConcreteMassConfig

    // Phase change materials
    PCMSystems []PCMConfig

    // Calculated values
    TotalHeatCapacity    float64 // kWh per 10°C swing
    TotalVolume          float64 // m³
    TotalCost            float64 // CAD estimated
    EffectiveStorageDays float64 // Days of heat reserve

    CreatedAt time.Time
    UpdatedAt time.Time
}

type WaterContainer struct {
    Type     WaterContainerType // "drum_55gal", "ibc_275gal", "tube", "custom"
    Material string             // "steel", "hdpe", "fiberglass"
    Quantity int
    Capacity float64 // Liters per container
    Location string  // "north_wall", "perimeter", "floor_level"
}

type WaterContainerType string

const (
    WCDrum55Gal  WaterContainerType = "drum_55gal"
    WCIBC275Gal  WaterContainerType = "ibc_275gal"
    WCTube       WaterContainerType = "tube"
    WCCustomTank WaterContainerType = "custom_tank"
)

type RockBedConfig struct {
    Enabled   bool
    RockType  string  // "river_rock", "gravel", "crusite"
    RockSize  float64 // mm diameter
    BedDepth  float64 // m
    BedArea   float64 // m²
    VoidRatio float64 // 0.35-0.45 typical
}

type ConcreteMassConfig struct {
    // Floor slab
    FloorEnabled   bool
    FloorThickness float64 // m
    FloorArea      float64 // m²
    FloorInsulated bool

    // Masonry wall
    WallEnabled   bool
    WallType      string  // "cmu_8in", "cmu_12in", "poured", "trombe"
    WallThickness float64 // m
    WallArea      float64 // m²
    WallGrouted   bool    // For CMU walls
}

type PCMConfig struct {
    Type         PCMType // "biopcm", "infinite_r", "salt_hydrate", "paraffin", "custom"
    MeltingPoint float64 // °C
    LatentHeat   float64 // kJ/kg
    Mass         float64 // kg
    Area         float64 // m² (for mat-type products)
    Location     string  // "north_wall", "ceiling", "growing_beds"
}

type PCMType string

const (
    PCMBioPCM      PCMType = "biopcm"
    PCMInfiniteR   PCMType = "infinite_r"
    PCMSaltHydrate PCMType = "salt_hydrate"
    PCMParaffin    PCMType = "paraffin"
    PCMCustom      PCMType = "custom"
)

// ThermalMassCalculation holds computed thermal performance values
type ThermalMassCalculation struct {
    // Input parameters
    GreenhouseVolume   float64 // m³
    GlazingArea        float64 // m²
    DesignDeltaT       float64 // °C (inside - coldest outside)
    ClimateZone        int

    // Calculated requirements
    DailyHeatLoss      float64 // kWh
    RecommendedStorage float64 // kWh
    RecommendedWaterL  float64 // Liters
    RecommendedPCMkg   float64 // kg

    // Current system capacity
    CurrentCapacity    float64 // kWh
    CapacityPercent    float64 // % of recommended
    EstimatedCoverage  float64 // % of heating needs met
}
```

---

## API Endpoints

| Method | Path | Handler | Description |
|--------|------|---------|-------------|
| GET | `/api/thermal-mass/materials` | `listThermalMassMaterials` | List available materials with properties |
| GET | `/api/thermal-mass/containers` | `listWaterContainers` | List water container options |
| GET | `/api/thermal-mass/pcm-products` | `listPCMProducts` | List PCM product options |
| POST | `/api/thermal-mass/calculate` | `calculateThermalMass` | Calculate storage capacity and recommendations |
| POST | `/api/thermal-mass/size` | `sizeThermalMass` | Size system for given greenhouse |
| GET | `/api/thermal-mass/cost-estimate` | `estimateThermalMassCost` | Estimate material and installation costs |

---

## Component Structure

```text
internal/components/wizard/
├── step8_thermal_mass.templ           # Main thermal mass selection step
├── thermal_mass_type_selector.templ   # Material type selection (water, rock, PCM)
├── water_container_selector.templ     # Container type and quantity
├── rock_bed_config.templ              # Rock bed configuration
├── pcm_selector.templ                 # PCM product selection
├── thermal_mass_calculator.templ      # Interactive sizing calculator
├── thermal_mass_placement.templ       # Placement visualization
└── thermal_mass_summary.templ         # Configuration summary with capacity
```

---

## References

### Primary Sources

- [Ceres Greenhouse Solutions - Tips on Using Water Barrels](https://ceresgs.com/tips-on-using-water-barrels-in-a-solar-greenhouse/)
- [UMass Extension - Heat Storage for Greenhouses](https://ag.umass.edu/greenhouse-floriculture/fact-sheets/heat-storage-for-greenhouses)
- [UGA Extension - Passive Solar Greenhouse Construction](https://extension.uga.edu/publications/detail.html?number=B1566)
- [YourHome (Australia) - Thermal Mass](https://www.yourhome.gov.au/passive-design/thermal-mass)

### Phase Change Materials

- [ScienceDirect - Phase Change Materials for Greenhouses: A Review](https://www.sciencedirect.com/science/article/abs/pii/S2213138822002934)
- [Green Built Alliance - The Case for Phase-Change Materials](https://www.greenbuilt.org/183-all-passive-heat-storage-is-not-created-equal-the-case-for-phase-change-materials/)
- [BioPCM Products](https://phasechange.com/products-biopcm/)
- [PCM Products Ltd](https://www.pcmproducts.net/)
- [BuildingGreen - BioPCM Review](https://www.buildinggreen.com/product-review/biopcm-finally-low-cost-practical-phase-change-material)

### Thermal Properties and Calculations

- [Engineering Toolbox - Storing Thermal Heat](https://www.engineeringtoolbox.com/sensible-heat-storage-d_1217.html)
- [Wikipedia - Volumetric Heat Capacity](https://en.wikipedia.org/wiki/Volumetric_heat_capacity)
- [Wikipedia - Table of Specific Heat Capacities](https://en.wikipedia.org/wiki/Table_of_specific_heat_capacities)
- [Chemistry LibreTexts - Thermal Mass for Heat Storage](https://chem.libretexts.org/Ancillary_Materials/Exemplars_and_Case_Studies/Exemplars/Environmental_and_Green_chemistry/Thermal_Mass_for_Heat_Storage)

### Trombe Walls and Masonry

- [Wikipedia - Trombe Wall](https://en.wikipedia.org/wiki/Trombe_wall)
- [InterNACHI - Trombe Walls](https://www.nachi.org/trombe-walls.htm)
- [Building America Solution Center - High-Thermal-Mass Construction](https://basc.pnnl.gov/resource-guides/high-thermal-mass-construction)
- [GreenBuildingAdvisor - Trombe Wall Construction](https://www.greenbuildingadvisor.com/article/trombe-wall-construction-mass-and-glass)

### Practical Guides

- [Mother Earth News - Thermal Mass Materials](https://www.motherearthnews.com/diy/thermal-mass-materials-zm0z18fmzsor/)
- [Verge Permaculture - Passive Solar Greenhouse Design Part 3](https://vergepermaculture.ca/designing-your-own-passive-solar-greenhouse-part-3/)
- [Sierra Greenhouse - Passive Solar Design](https://www.sierragreenhouse.com/blog/passive-solar-greenhouse-design/)
- [Planta Greenhouses - Guide to Thermal Mass](https://plantagreenhouses.com/blogs/learn/the-ultimate-guide-to-incorporating-thermal-mass-in-your-greenhouse)
- [Growing Spaces - Thermal Mass in Dome Greenhouses](https://growingspaces.com/geodesic-dome-greenhouse/thermal-mass/)

### Research Papers

- [IIETA - Effects of Rock-Bed Heat Storage on Solar Greenhouse Microclimate](https://www.iieta.org/journals/i2m/paper/10.18280/i2m.190608)
- [ScienceDirect - Thermal Modelling for Greenhouse Heating Using Packed Bed](https://scialert.net/fulltext/?doi=ijar.2006.373.383)
- [ScienceDirect - Climate-Responsive Thermal Mass Design for Sunspaces](https://www.sciencedirect.com/science/article/abs/pii/S0960148115301233)
- [MDPI - Active-Passive Heat Storage Wall with PCM in Chinese Solar Greenhouse](https://www.mdpi.com/2071-1050/16/7/2624)

### Container Sources

- [IBC Tanks - 275 Gallon Totes](https://www.ibctanks.com/275gallon)
- [Lexington Container Company - IBC Totes](https://www.lexingtoncontainercompany.com/275-Gallon-IBC-Totes.html)

### DIY PCM Materials and Formulations

- [Build It Solar - DIY Phase Change Material](https://www.builditsolar.com/Experimental/PCM/DIYPhaseChangeMaterial.htm)
- [MDPI - Edible Oils as Practical Phase Change Materials](https://www.mdpi.com/2076-3417/9/8/1627)
- [MDPI - Potential of Coconut Oil for Air Temperature Control](https://www.mdpi.com/2075-5309/8/8/95)
- [Springer - Coconut Oil as PCM: A Review](https://link.springer.com/article/10.1007/s10973-021-10839-7)
- [MDPI - Climate-Based Analysis for Coconut Oil PCM in Buildings](https://www.mdpi.com/2071-1050/13/19/10731)
- [ScienceDirect - Capric and Lauric Acid Mixture as PCM](https://www.sciencedirect.com/science/article/abs/pii/S0360544202000245)
- [ACS Omega - Lauric Acid-Based Eutectic Fatty Acid PCM](https://pubs.acs.org/doi/10.1021/acsomega.2c01420)
- [ACS Omega - Capric-Stearic Acid PCM](https://pubs.acs.org/doi/10.1021/acsomega.1c01705)
- [PMC - Stabilization of Low-Cost Salt Hydrate PCMs](https://pmc.ncbi.nlm.nih.gov/articles/PMC10329044/)
- [MDPI - Optimization of CaCl₂·6H₂O PCM Synthesis](https://www.mdpi.com/2075-5309/12/10/1762)
- [GreenBuildingAdvisor - Calcium Chloride Hexahydrate for Heat Storage](https://www.greenbuildingadvisor.com/question/calcium-chloride-hexahydrate-for-phase-change-heat-storage)
- [ResearchGate - Thermal Energy Storage Using Calcium Chloride Hexahydrate](https://www.researchgate.net/publication/317274757_Thermal_energy_storage_using_calcium_chloride_hexahydrate)
- [Wikipedia - Phase-Change Material](https://en.wikipedia.org/wiki/Phase-change_material)
- [Permies Forum - Phase Change Materials for Heat Sinks](https://permies.com/t/9979/phase-change-materials-heat-sinks)
- [Paraffin Wax Co - Phase Change Materials](https://paraffinwaxco.com/phase-change-materials/)

### Fatty Acid Suppliers (North America)

- [Voyageur Soap & Candle](https://www.voyageursoapandcandle.com/) - Fatty acids for soap making
- [Bulk Apothecary](https://www.bulkapothecary.com/) - Various fatty acids and carrier oils
- [Cailà & Parés - Caprylic/Capric Acid](https://cailapares.com/en/caprylic-acid-capric/)
