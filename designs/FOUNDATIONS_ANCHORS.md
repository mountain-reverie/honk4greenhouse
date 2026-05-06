# Greenhouse Foundation and Anchor Systems

This document provides comprehensive technical guidance for selecting and implementing foundation and anchor systems for greenhouses in Canadian climate conditions. Foundation selection significantly impacts structural integrity, thermal performance, construction cost, and long-term maintenance requirements.

## Table of Contents

1. [Foundation Selection Overview](#foundation-selection-overview)
2. [Canadian Frost Depth Requirements](#canadian-frost-depth-requirements)
3. [Foundation Types](#foundation-types)
   - [Frost-Protected Shallow Foundation (FPSF)](#frost-protected-shallow-foundation-fpsf)
   - [Helical Pile Foundation](#helical-pile-foundation)
   - [Concrete Block Wall Foundation](#concrete-block-wall-foundation)
   - [Poured Concrete Slab](#poured-concrete-slab)
   - [Post Frame Foundation](#post-frame-foundation)
   - [Floating Foundation](#floating-foundation)
4. [Anchor Systems](#anchor-systems)
   - [Ground Anchors](#ground-anchors)
   - [Concrete Anchors](#concrete-anchors)
   - [Ground Screws](#ground-screws)
5. [Wind Load and Uplift Resistance](#wind-load-and-uplift-resistance)
6. [Greenhouse-Specific Considerations](#greenhouse-specific-considerations)
7. [Cost Analysis](#cost-analysis)
8. [Installation Guidelines](#installation-guidelines)
9. [Data Model](#data-model)
10. [References](#references)

---

## Foundation Selection Overview

### Key Selection Factors

| Factor | Considerations |
|--------|----------------|
| **Soil Type** | Load-bearing capacity, drainage, frost susceptibility |
| **Frost Depth** | Determines minimum foundation depth or insulation requirements |
| **Greenhouse Type** | Hoop house, wood frame, commercial structure |
| **Permanence** | Temporary/seasonal vs. permanent installation |
| **Budget** | Material costs, labor requirements, equipment needs |
| **Thermal Performance** | Heat loss through foundation, thermal bridging |
| **Local Codes** | Building permit requirements, engineering certification |

### Foundation Type Comparison

| Foundation Type | Best For | Frost Protection | Relative Cost | Complexity |
|-----------------|----------|------------------|---------------|------------|
| FPSF Grade Beam | Permanent wood frame, moderate climates | Horizontal insulation | Low-Medium | Medium |
| Helical Pile | All types, difficult soils, quick install | Below frost line | Medium | Low |
| CMU Block Wall | Permanent structures, basements | Below frost line | Medium-High | High |
| Poured Slab | Commercial, high-traffic, heavy loads | Varies by design | High | High |
| Post Frame | Budget builds, hoop houses | Below frost line or floating | Low | Low |
| Floating | Temporary structures, poor soils | Surface-level | Low | Low |

---

## Canadian Frost Depth Requirements

Frost depth determines how deep foundations must extend to prevent frost heave damage. In Canada, this varies significantly by region.

### Regional Frost Depths

| Region | Typical Frost Depth | Design Frost Depth |
|--------|---------------------|-------------------|
| Southern BC (Lower Mainland) | 0.3-0.6 m | 0.6 m |
| Vancouver Island | 0.15-0.45 m | 0.45 m |
| Southern Ontario | 0.9-1.2 m | 1.2 m |
| Central Ontario | 1.2-1.5 m | 1.5 m |
| Southern Prairies | 1.5-2.1 m | 2.1 m |
| Northern Prairies | 2.1-2.7 m | 2.7 m |
| Quebec (Montreal) | 1.2-1.5 m | 1.5 m |
| Quebec (Quebec City) | 1.5-1.8 m | 1.8 m |
| Atlantic Provinces | 0.9-1.5 m | 1.5 m |
| Northern Canada | 2.4-3.6+ m | Permafrost considerations |

### Frost Protection Strategies

**1. Extend Below Frost Line**
- Traditional approach for permanent structures
- Foundation footings placed below maximum frost depth
- Most reliable but most expensive in deep frost regions

**2. Frost-Protected Shallow Foundation (FPSF)**
- Uses horizontal insulation to reduce frost penetration
- Allows shallower foundations in cold climates
- Cost-effective for heated structures like greenhouses

**3. Helical/Screw Piles**
- Extend below frost line with minimal excavation
- Load transferred to stable soil below frost zone
- Ideal for difficult or frost-susceptible soils

---

## Foundation Types

### Frost-Protected Shallow Foundation (FPSF)

Based on UMN Farm Scale Winter Greenhouse F.01 design. FPSF uses strategic insulation placement to prevent frost from penetrating beneath the foundation, allowing shallower construction in cold climates.

#### Design Components

```
┌─────────────────────────────────────────────────────────────┐
│                    GREENHOUSE STRUCTURE                      │
├─────────────────────────────────────────────────────────────┤
│  ┌─────────────────────────────────────────────────────┐    │
│  │              4x4 Plastic Grade Beam                  │    │
│  │         (Recycled Plastic Composite)                 │    │
│  └─────────────────────────────────────────────────────┘    │
│                           │                                  │
│  ┌─────────────────────────────────────────────────────┐    │
│  │              2x8 Footing Plate                       │    │
│  └─────────────────────────────────────────────────────┘    │
│                           │                                  │
│  ════════════════════════════════════════════════════════   │
│           GRADE LEVEL                                        │
│  ════════════════════════════════════════════════════════   │
│                           │                                  │
│  ┌───────────────────────────────────────────────────────┐  │
│  │         4' (1.2m) Horizontal XPS Insulation           │  │
│  │              R-10 to R-20 (2"-4" thick)               │  │
│  └───────────────────────────────────────────────────────┘  │
│                                                              │
└─────────────────────────────────────────────────────────────┘
```

#### Specifications

| Component | Specification | Purpose |
|-----------|--------------|---------|
| Grade Beam | 4x4 (89mm x 89mm) recycled plastic lumber | Rot-resistant, supports wall loads |
| Footing | 2x8 (38mm x 184mm) pressure-treated or plastic | Load distribution |
| Horizontal Insulation | R-10 to R-20 XPS (50-100mm) | Frost protection |
| Insulation Width | 1.2-1.5m from foundation | Creates thermal "umbrella" |
| Insulation Slope | 1% away from building | Drainage |

#### FPSF Insulation Requirements by Climate

| Annual Mean Temperature | Horizontal Insulation R-Value | Width from Foundation |
|------------------------|------------------------------|----------------------|
| Above 7°C | R-5 (25mm XPS) | 0.6m |
| 5-7°C | R-7.5 (38mm XPS) | 0.9m |
| 3-5°C | R-10 (50mm XPS) | 1.2m |
| 0-3°C | R-15 (75mm XPS) | 1.2m |
| Below 0°C | R-20 (100mm XPS) | 1.5m |

#### Advantages
- Lower material costs than deep foundations
- Reduced excavation requirements
- Integrates with greenhouse thermal envelope
- Well-suited for heated structures

#### Disadvantages
- Requires heated structure (not for unheated greenhouses)
- Insulation must be protected from damage
- Less suitable for extremely cold regions (Zone 7+)
- Requires additional tie-down anchoring

#### Installation Notes
- Excavate to stable subgrade (typically 150-300mm)
- Compact base with 100mm gravel layer
- Install horizontal insulation with proper slope for drainage
- Use ground screws or auger anchors for uplift resistance
- Protect exposed insulation with backfill or rigid covering

---

### Helical Pile Foundation

Based on UMN Farm Scale Winter Greenhouse F.02 design. Helical piles (also called screw piles) are steel shafts with helical plates that are rotated into the ground, providing immediate load-bearing capacity.

#### Design Components

```
                    ┌─────────────────┐
                    │  Post Bracket   │
                    │  (Simpson type) │
                    └────────┬────────┘
                             │
    ══════════════════════════════════════════  GRADE LEVEL
                             │
                    ┌────────┴────────┐
                    │   Steel Shaft   │
                    │   (73-89mm OD)  │
                    │                 │
    ┌───────────────┼─────────────────┼───────────────┐
    │               │   4" R-20 XPS   │               │
    │               │   Perimeter     │               │
    │               │   Insulation    │               │
    └───────────────┼─────────────────┼───────────────┘
                    │                 │
                    │                 │
                    │                 │
    ────────────────┼─────────────────┼────────────────  FROST LINE
                    │                 │
                    ├─────────────────┤
                    │  Helical Plate  │
                    │  (200-350mm)    │
                    └─────────────────┘
```

#### Specifications

| Component | Specification | Purpose |
|-----------|--------------|---------|
| Shaft Diameter | 73-89mm (2.875"-3.5") OD | Structural support |
| Shaft Wall Thickness | 5-7mm | Load capacity |
| Helix Diameter | 200-350mm (8"-14") | Soil engagement |
| Helix Thickness | 10-13mm | Strength |
| Installation Depth | Below frost line + 0.3m minimum | Frost protection |
| Perimeter Insulation | R-20 XPS (100mm) vertical | Thermal performance |
| Bracket | Simpson PBS or equivalent | Post connection |

#### Load Capacity

| Pile Configuration | Typical Capacity (Compression) | Typical Capacity (Tension/Uplift) |
|-------------------|-------------------------------|----------------------------------|
| 73mm shaft, single 200mm helix | 15-30 kN | 10-20 kN |
| 89mm shaft, single 300mm helix | 40-80 kN | 25-50 kN |
| 89mm shaft, double helix | 60-120 kN | 40-80 kN |

*Actual capacity depends on soil conditions. Professional engineering recommended.*

#### Advantages
- Minimal excavation required
- Immediate load-bearing (no curing time)
- Works in difficult soils (clay, wet, rocky)
- Can be removed and relocated
- Year-round installation possible
- Minimal site disturbance

#### Disadvantages
- Requires specialized equipment for installation
- Higher per-pile cost than concrete piers
- May require engineering certification
- Corrosion considerations in aggressive soils

#### Installation Notes
- Conduct soil test or test pile to verify capacity
- Install to minimum torque specification (correlates with capacity)
- Extend minimum 0.3m below frost line
- Install perimeter insulation around pile caps
- Use galvanized or epoxy-coated piles in corrosive soils

---

### Concrete Block Wall Foundation

Based on UMN Farm Scale Winter Greenhouse F.03 design. A full perimeter foundation using concrete masonry units (CMU) provides maximum structural support and thermal mass.

#### Design Components

```
┌─────────────────────────────────────────────────────────────┐
│                    GREENHOUSE STRUCTURE                      │
├─────────────────────────────────────────────────────────────┤
│  ┌─────────────────────────────────────────────────────┐    │
│  │              Pressure-Treated Sill Plate             │    │
│  │                   with Anchor Bolts                  │    │
│  └─────────────────────────────────────────────────────┘    │
│                           │                                  │
│  ┌─────────────────────────────────────────────────────┐    │
│  │              8" (200mm) CMU Block Wall               │    │
│  │              Grouted cores with rebar                │    │
│  │              Waterproofing on exterior               │    │
│  ├─────────────────────────────────────────────────────┤    │
│  │              GRADE LEVEL                             │    │
│  ├─────────────────────────────────────────────────────┤    │
│  │              CMU continues below grade               │    │
│  │              Exterior insulation (R-10 to R-20)      │    │
│  │              Drainage board/membrane                 │    │
│  └─────────────────────────────────────────────────────┘    │
│                           │                                  │
│  ┌─────────────────────────────────────────────────────┐    │
│  │              Poured Concrete Footing                 │    │
│  │              200mm x 400mm typical                   │    │
│  │              Below frost line                        │    │
│  └─────────────────────────────────────────────────────┘    │
│                           │                                  │
│  ┌─────────────────────────────────────────────────────┐    │
│  │              Perimeter Drain Tile                    │    │
│  │              100mm perforated pipe                   │    │
│  │              Gravel bed                              │    │
│  └─────────────────────────────────────────────────────┘    │
│                                                              │
└─────────────────────────────────────────────────────────────┘
```

#### Specifications

| Component | Specification | Purpose |
|-----------|--------------|---------|
| CMU Block | 200mm x 200mm x 400mm (8" standard) | Structural wall |
| Footing Width | 400mm (16") minimum | Load distribution |
| Footing Depth | Below frost line | Frost protection |
| Reinforcement | #4 (12mm) rebar in grouted cores | Structural strength |
| Anchor Bolts | 12mm (1/2") J-bolts @ 1.2m (4') O.C. | Sill attachment |
| Waterproofing | Bituminous coating or membrane | Moisture protection |
| Exterior Insulation | R-10 to R-20 XPS | Thermal performance |
| Drain Tile | 100mm (4") perforated | Drainage |

#### Advantages
- Maximum structural capacity
- Excellent thermal mass potential
- Long service life (50+ years)
- Can incorporate basement/root cellar
- Familiar to contractors

#### Disadvantages
- Highest cost foundation option
- Requires extensive excavation
- Long construction timeline
- Weather-dependent installation
- Difficult to modify or relocate

#### Installation Notes
- Excavate to frost depth plus 200mm for footing
- Install drain tile on gravel bed
- Pour footing with keyway for wall
- Lay CMU with grouted cores and rebar
- Apply waterproofing before backfill
- Install exterior insulation with protection board
- Backfill with free-draining material

---

### Poured Concrete Slab

A monolithic or floating concrete slab provides a durable floor and foundation in one element. Common for commercial greenhouses with heavy equipment.

#### Design Types

**1. Monolithic Slab (Thickened Edge)**
- Single pour with integral footing
- Suitable for mild climates or with insulation

**2. Floating Slab**
- Slab rests on gravel base
- Not attached to frost walls
- For heated structures with perimeter insulation

**3. Slab-on-Grade with Frost Wall**
- Slab poured on grade
- Separate frost walls extend below
- Most frost-resistant option

#### Specifications

| Component | Specification | Purpose |
|-----------|--------------|---------|
| Slab Thickness | 100-150mm (4"-6") | Structural floor |
| Edge Thickness | 300-450mm (12"-18") for monolithic | Frost/load support |
| Reinforcement | Welded wire mesh (WWF 6x6 W1.4xW1.4) | Crack control |
| Concrete Strength | 25 MPa (3500 psi) minimum | Durability |
| Vapor Barrier | 6 mil (0.15mm) polyethylene | Moisture control |
| Sub-Base | 100-150mm compacted gravel | Drainage/stability |
| Underslab Insulation | R-10 to R-20 XPS (optional) | Thermal performance |

#### Advantages
- Excellent for heavy equipment/traffic
- Easy to clean and maintain
- Provides thermal mass when insulated
- Long service life

#### Disadvantages
- High material and labor cost
- Extensive site preparation
- Permanent installation
- Long curing time before use
- Cracking potential

---

### Post Frame Foundation

Post frame (pole barn) construction embeds posts directly in the ground or on concrete piers. Common for agricultural buildings and hoop houses.

#### Design Types

**1. Embedded Post**
- Posts set in concrete collar below frost line
- Pressure-treated or naturally rot-resistant wood
- Simple, economical construction

**2. Post on Pier**
- Concrete pier extends below frost line
- Post sits on pier with bracket
- Keeps wood above ground

**3. Surface-Mounted Post**
- Post on concrete pad at grade
- Ground anchors provide uplift resistance
- For temporary or portable structures

#### Specifications (Embedded Post)

| Component | Specification | Purpose |
|-----------|--------------|---------|
| Post Size | 150mm x 150mm (6x6) minimum | Structural support |
| Post Material | Pressure-treated (UC4A or UC4B) | Rot resistance |
| Embedment Depth | Frost depth + 0.6m (2') minimum | Frost/stability |
| Concrete Collar | 250-300mm diameter | Bearing/stability |
| Spacing | 1.8-3.6m (6'-12') typical | Load distribution |
| Backfill | Compacted gravel or concrete | Stability |

#### Advantages
- Low cost
- Simple construction
- Flexible sizing
- DIY-friendly

#### Disadvantages
- Posts can rot over time (15-25 year life)
- Settlement possible
- Limited load capacity
- May not meet building codes

---

### Floating Foundation

A floating or "raft" foundation sits on grade without extending below frost line. Suitable for lightweight, temporary, or seasonal structures.

#### Applications
- Seasonal high tunnels
- Temporary growing structures
- Sites with high water table
- Portable greenhouses

#### Design Components

| Component | Specification | Purpose |
|-----------|--------------|---------|
| Base Frame | 100mm x 100mm (4x4) PT lumber or steel | Structure support |
| Gravel Pad | 150-300mm compacted gravel | Drainage/leveling |
| Landscape Fabric | Commercial grade | Weed barrier |
| Anchors | Ground anchors at corners/intervals | Uplift resistance |

#### Anchor Requirements
- Minimum 4 anchors for structures up to 6m x 9m
- Additional anchors at 2.4-3.6m intervals for larger structures
- Anchor capacity must exceed calculated wind uplift

---

## Anchor Systems

Anchors provide resistance to wind uplift and lateral forces. Essential for all greenhouse types, especially in high-wind regions.

### Ground Anchors

#### Auger Anchors (Earth Anchors)

Helical steel anchors screwed into soil. Most common for greenhouse applications.

```
         ┌───┐
         │ O │  Eye/Ring
         └─┬─┘
           │   Shaft (12-19mm diameter)
           │
           │
    ═══════╪═══════  GRADE LEVEL
           │
           │
           │
         ╱─┼─╲
        ╱──┼──╲    Auger Plate (75-150mm)
       ╱───┼───╲
           │
           │
         ──┴──
```

| Type | Shaft Diameter | Auger Diameter | Length | Typical Capacity |
|------|----------------|----------------|--------|------------------|
| Light Duty | 12mm (1/2") | 75mm (3") | 0.6-0.9m | 2-4 kN |
| Medium Duty | 15mm (5/8") | 100mm (4") | 0.75-1.0m | 4-8 kN |
| Heavy Duty | 19mm (3/4") | 150mm (6") | 0.9-1.2m | 8-15 kN |

**Installation**: Drive with impact wrench or manual anchor driver. Minimum embedment 0.6m in firm soil.

#### Duckbill Anchors

Hinged plate anchors that rotate perpendicular to shaft when tensioned.

- Light duty: 2-5 kN capacity
- Heavy duty: 5-15 kN capacity
- Requires driving tool and tensioning
- Good for temporary installations

#### Arrowhead Anchors

Similar to duckbill but with arrowhead-shaped plate.

- Driven into soil with hammer
- Tensioned to lock in place
- 3-10 kN typical capacity
- Low cost, single-use

### Concrete Anchors

#### Cast-In-Place Anchors

```
    ┌───────────┐
    │   Nut     │
    ├───────────┤
    │  Washer   │
════╪═══════════╪════  CONCRETE SURFACE
    │           │
    │   J-Bolt  │
    │   or      │
    │   L-Bolt  │
    │           │
    └───────────┘
```

| Anchor Type | Diameter | Embedment | Typical Capacity (Tension) |
|-------------|----------|-----------|---------------------------|
| J-Bolt | 12mm (1/2") | 150mm | 5-8 kN |
| J-Bolt | 16mm (5/8") | 200mm | 8-12 kN |
| L-Bolt | 12mm (1/2") | 150mm | 6-10 kN |
| L-Bolt | 19mm (3/4") | 250mm | 12-18 kN |

#### Post-Installed Anchors

For anchoring to existing concrete:

| Type | Typical Capacity | Best For |
|------|------------------|----------|
| Wedge Anchor | 5-20 kN | Permanent, high load |
| Sleeve Anchor | 3-12 kN | Permanent, moderate load |
| Concrete Screw | 2-8 kN | Light-medium duty |
| Epoxy Anchor | 10-40 kN | Highest capacity, permanent |

### Ground Screws

Steel screw piles for surface-mounted structures. Combine foundation and anchor functions.

```
         ┌─────────────┐
         │ Mounting    │
         │ Bracket     │
         └──────┬──────┘
                │
    ════════════╪════════════  GRADE LEVEL
                │
           ╱────┼────╲
          ╱─────┼─────╲   Upper Helix
                │
           ╱────┼────╲
          ╱─────┼─────╲   Lower Helix (if applicable)
                │
              ──┴──
```

| Ground Screw Type | Shaft Diameter | Length | Typical Capacity |
|-------------------|----------------|--------|------------------|
| Light (greenhouse) | 60mm | 0.8-1.0m | 3-8 kN |
| Medium | 76mm | 1.0-1.5m | 8-15 kN |
| Heavy | 89mm | 1.2-2.0m | 15-30 kN |

**UMN Reference**: The Farm Scale Winter Greenhouse design uses 12 ground screws total for a 28' x 96' structure.

---

## Wind Load and Uplift Resistance

### Wind Load Calculation

Wind creates both lateral (sideways) and uplift (vertical) forces on greenhouse structures.

#### Basic Wind Pressure Formula

```
p = 0.613 × V²
Where:
p = Wind pressure (Pa)
V = Wind speed (m/s)
```

| Wind Speed | Basic Pressure | With Gust Factor (1.4) |
|------------|---------------|------------------------|
| 80 km/h (22 m/s) | 297 Pa | 416 Pa |
| 100 km/h (28 m/s) | 481 Pa | 673 Pa |
| 120 km/h (33 m/s) | 668 Pa | 935 Pa |
| 140 km/h (39 m/s) | 933 Pa | 1306 Pa |

#### Uplift Force Calculation

```
F_uplift = p × A × C_p
Where:
F_uplift = Uplift force (N)
p = Design wind pressure (Pa)
A = Projected roof area (m²)
C_p = Pressure coefficient (typically -0.7 to -1.3 for greenhouse roofs)
```

#### Example: 8m x 12m Greenhouse

- Roof area: ~105 m² (including slope)
- Design wind: 100 km/h with gusts
- Pressure coefficient: -1.0 (uplift)
- Uplift force: 673 Pa × 105 m² × 1.0 = **70.7 kN total uplift**
- Per anchor (8 anchors): 8.8 kN minimum capacity required

### Regional Wind Design Requirements (Canada)

| Region | Reference Wind Pressure (1/50 year) | Typical Design Speed |
|--------|-------------------------------------|---------------------|
| Southern BC Coast | 0.42-0.55 kPa | 90-100 km/h |
| BC Interior | 0.30-0.42 kPa | 75-90 km/h |
| Prairies | 0.38-0.50 kPa | 85-100 km/h |
| Southern Ontario | 0.36-0.48 kPa | 80-95 km/h |
| Quebec/Atlantic | 0.40-0.60 kPa | 90-110 km/h |
| Northern Canada | 0.30-0.45 kPa | 75-95 km/h |

*Reference: National Building Code of Canada, Appendix C*

### Anchor Spacing Guidelines

| Greenhouse Width | Recommended Anchor Spacing (Sidewalls) | End Wall Anchors |
|------------------|----------------------------------------|------------------|
| Up to 6m | 2.4m (8') maximum | 2 per end minimum |
| 6-9m | 1.8m (6') maximum | 3 per end minimum |
| 9-12m | 1.5m (5') maximum | 4 per end minimum |
| Over 12m | Engineering required | Engineering required |

---

## Greenhouse-Specific Considerations

### Hoop House / High Tunnel

**Typical Foundation Options:**
1. Ground posts (driven or embedded)
2. Baseboards with ground anchors
3. Concrete blocks with straps

**Recommendations:**
- Minimum 0.6m embedment for ground posts
- Anchor every bow (1.2-1.8m spacing)
- Cross-bracing at end walls
- Consider wind rating of covering material

### Gothic Arch / Quonset

**Similar to hoop house but with different load distribution:**
- Higher wind resistance due to shape
- Requires secure ridge connection
- End walls often need additional bracing

### Wood Frame Greenhouse

**Foundation Options (in order of cost/permanence):**
1. Skids on gravel (temporary)
2. Concrete blocks/piers
3. FPSF grade beam
4. Helical piles
5. CMU block wall

**Recommendations:**
- Sill plate must be pressure-treated or naturally rot-resistant
- Sill sealer gasket for air sealing
- Anchor bolts minimum 1.2m (4') O.C.
- Rim joist connections for uplift resistance

### Commercial Glass/Polycarbonate

**Typically requires:**
- Engineered foundation design
- Concrete slab or continuous footing
- Professional installation
- Building permit and inspection

### Deep Winter Greenhouse

Based on UMN designs, three recommended foundation types:

**F.01 - FPSF Grade Beam**
- Best for: Moderate budgets, DIY-friendly
- Frost protection: 1.2m horizontal XPS insulation
- Requires: Ground screw anchors for uplift

**F.02 - Helical Pile**
- Best for: Difficult soils, quick installation
- Frost protection: Piles below frost + perimeter insulation
- Requires: Professional installation or rental equipment

**F.03 - CMU Block Wall**
- Best for: Maximum durability, integrated root cellar
- Frost protection: Below frost line construction
- Requires: Contractor or experienced DIY

---

## Cost Analysis

### Material Costs (2024 Canadian Estimates)

| Foundation Type | Material Cost per Linear Meter | Total for 8m x 12m |
|-----------------|-------------------------------|-------------------|
| Ground Anchors Only | $15-30 | $600-1,200 |
| FPSF Grade Beam | $80-150 | $3,200-6,000 |
| Helical Piles (installed) | $150-300 per pile | $3,000-6,000 |
| CMU Block Wall | $200-350 | $8,000-14,000 |
| Poured Concrete Slab | $150-250/m² | $14,400-24,000 |
| Post Frame (embedded) | $50-100 | $2,000-4,000 |

### UMN Farm Scale Winter Greenhouse Reference Costs

From the design documents (28' x 96' / 8.5m x 29m structure):

| Foundation Option | Material Cost (USD) |
|-------------------|---------------------|
| F.01 + TM.1 (FPSF + Thermal Mass) | ~$35,000 |
| F.02 (Helical Pile) | Similar range |
| F.03 (CMU Wall) | Higher |

*Note: Includes thermal mass (TM.1) water wall system*

### Labor Considerations

| Foundation Type | DIY Feasible? | Professional Cost Factor |
|-----------------|---------------|-------------------------|
| Ground Anchors | Yes | 1.0x (DIY baseline) |
| FPSF Grade Beam | Yes (moderate skill) | 1.5-2.0x |
| Helical Piles | Possible with rental | 1.5-2.5x |
| CMU Block Wall | Challenging | 2.0-3.0x |
| Poured Concrete | No (usually) | 2.5-3.5x |
| Post Frame | Yes | 1.0-1.5x |

---

## Installation Guidelines

### Site Preparation

1. **Survey and Layout**
   - Mark building corners with stakes
   - Verify square using 3-4-5 triangle or diagonal measurements
   - Check for underground utilities (call before you dig)

2. **Grading**
   - Ensure positive drainage away from structure
   - Minimum 2% slope for first 3m from building
   - Remove organic material (topsoil) from foundation area

3. **Soil Assessment**
   - Visual inspection for soil type
   - Hand test for compaction
   - Professional soil test for larger projects

### Foundation-Specific Installation

#### FPSF Grade Beam
1. Excavate to stable subgrade (150-300mm)
2. Install 100mm compacted gravel base
3. Place grade beam lumber on gravel
4. Install horizontal insulation with drainage slope
5. Backfill and compact
6. Install ground screw anchors through grade beam

#### Helical Piles
1. Mark pile locations (typically at post positions)
2. Install with hydraulic driver to specified torque
3. Cut shafts to uniform height
4. Install pile caps/brackets
5. Install perimeter insulation
6. Backfill around insulation

#### CMU Block Wall
1. Excavate to frost depth plus footing
2. Install drain tile and gravel bed
3. Pour footing with keyway (let cure 7+ days)
4. Lay CMU with grouted cores and rebar
5. Install anchor bolts in top course
6. Apply waterproofing and insulation
7. Backfill with free-draining material

### Quality Checks

- [ ] Corners are square (diagonal measurements within 6mm)
- [ ] Top of foundation is level (within 6mm over 3m)
- [ ] Anchor bolts are plumb and at correct height
- [ ] Drainage slope is correct
- [ ] Insulation is continuous without gaps
- [ ] Waterproofing is complete (for below-grade)

---

## Data Model

### Go Structs for Foundation Configuration

```go
// FoundationType represents the category of foundation system
type FoundationType string

const (
    FoundationFPSF        FoundationType = "fpsf"
    FoundationHelicalPile FoundationType = "helical_pile"
    FoundationCMUWall     FoundationType = "cmu_wall"
    FoundationSlab        FoundationType = "slab"
    FoundationPostFrame   FoundationType = "post_frame"
    FoundationFloating    FoundationType = "floating"
)

// AnchorType represents the type of anchor system
type AnchorType string

const (
    AnchorAuger       AnchorType = "auger"
    AnchorDuckbill    AnchorType = "duckbill"
    AnchorGroundScrew AnchorType = "ground_screw"
    AnchorJBolt       AnchorType = "j_bolt"
    AnchorLBolt       AnchorType = "l_bolt"
    AnchorWedge       AnchorType = "wedge"
    AnchorEpoxy       AnchorType = "epoxy"
)

// FoundationConfig represents the complete foundation configuration
type FoundationConfig struct {
    ID                string            `json:"id"`
    ProjectID         string            `json:"project_id"`
    Type              FoundationType    `json:"type"`
    FrostDepth        float64           `json:"frost_depth_m"`        // Local frost depth in meters
    DesignWindSpeed   float64           `json:"design_wind_speed_kmh"` // Design wind speed in km/h

    // Type-specific configuration (one will be populated)
    FPSF              *FPSFConfig       `json:"fpsf,omitempty"`
    HelicalPile       *HelicalPileConfig `json:"helical_pile,omitempty"`
    CMUWall           *CMUWallConfig    `json:"cmu_wall,omitempty"`
    Slab              *SlabConfig       `json:"slab,omitempty"`
    PostFrame         *PostFrameConfig  `json:"post_frame,omitempty"`

    Anchors           []AnchorConfig    `json:"anchors"`

    // Calculated values
    TotalUpliftResistance float64       `json:"total_uplift_resistance_kn"`
    EstimatedCost         float64       `json:"estimated_cost_cad"`
}

// FPSFConfig represents Frost-Protected Shallow Foundation configuration
type FPSFConfig struct {
    GradeBeamMaterial   string  `json:"grade_beam_material"`   // "plastic_lumber", "pressure_treated"
    GradeBeamSize       string  `json:"grade_beam_size"`       // "4x4", "4x6", "6x6"
    FootingMaterial     string  `json:"footing_material"`
    FootingSize         string  `json:"footing_size"`          // "2x8", "2x10"

    HorizontalInsulation InsulationConfig `json:"horizontal_insulation"`
    InsulationWidth     float64 `json:"insulation_width_m"`    // Width from foundation
}

// HelicalPileConfig represents helical pile foundation configuration
type HelicalPileConfig struct {
    PileCount           int     `json:"pile_count"`
    ShaftDiameter       float64 `json:"shaft_diameter_mm"`
    HelixDiameter       float64 `json:"helix_diameter_mm"`
    InstallationDepth   float64 `json:"installation_depth_m"`
    CapacityPerPile     float64 `json:"capacity_per_pile_kn"`

    PerimeterInsulation InsulationConfig `json:"perimeter_insulation"`
}

// CMUWallConfig represents concrete block wall foundation configuration
type CMUWallConfig struct {
    BlockWidth          float64 `json:"block_width_mm"`        // 150, 200, 250, 300
    WallHeight          float64 `json:"wall_height_m"`
    FootingWidth        float64 `json:"footing_width_mm"`
    FootingDepth        float64 `json:"footing_depth_mm"`
    RebarSize           string  `json:"rebar_size"`            // "#4", "#5"
    RebarSpacing        float64 `json:"rebar_spacing_mm"`

    ExteriorInsulation  InsulationConfig `json:"exterior_insulation"`
    HasDrainTile        bool    `json:"has_drain_tile"`
    WaterproofingType   string  `json:"waterproofing_type"`
}

// SlabConfig represents concrete slab foundation configuration
type SlabConfig struct {
    SlabThickness       float64 `json:"slab_thickness_mm"`
    EdgeThickness       float64 `json:"edge_thickness_mm"`      // For monolithic
    ConcreteStrength    float64 `json:"concrete_strength_mpa"`
    ReinforcementType   string  `json:"reinforcement_type"`     // "wwf", "rebar", "fiber"

    HasVaporBarrier     bool    `json:"has_vapor_barrier"`
    UnderslabInsulation *InsulationConfig `json:"underslab_insulation,omitempty"`
    PerimeterInsulation *InsulationConfig `json:"perimeter_insulation,omitempty"`
}

// PostFrameConfig represents post frame foundation configuration
type PostFrameConfig struct {
    PostSize            string  `json:"post_size"`             // "4x4", "6x6"
    PostMaterial        string  `json:"post_material"`         // "pressure_treated", "cedar"
    PostSpacing         float64 `json:"post_spacing_m"`
    EmbedmentDepth      float64 `json:"embedment_depth_m"`
    ConcreteCollar      bool    `json:"concrete_collar"`
    CollarDiameter      float64 `json:"collar_diameter_mm"`
}

// InsulationConfig represents insulation specifications
type InsulationConfig struct {
    Material            string  `json:"material"`              // "xps", "eps", "polyiso"
    Thickness           float64 `json:"thickness_mm"`
    RValue              float64 `json:"r_value_m2kw"`          // m²·K/W
}

// AnchorConfig represents individual anchor specifications
type AnchorConfig struct {
    Type                AnchorType `json:"type"`
    Quantity            int     `json:"quantity"`
    Diameter            float64 `json:"diameter_mm"`
    Length              float64 `json:"length_m"`
    Spacing             float64 `json:"spacing_m"`
    UpliftCapacity      float64 `json:"uplift_capacity_kn"`    // Per anchor
    LateralCapacity     float64 `json:"lateral_capacity_kn"`   // Per anchor
}

// WindLoadCalculation represents wind load analysis results
type WindLoadCalculation struct {
    DesignWindSpeed     float64 `json:"design_wind_speed_kmh"`
    BasicPressure       float64 `json:"basic_pressure_pa"`
    GustFactor          float64 `json:"gust_factor"`
    DesignPressure      float64 `json:"design_pressure_pa"`

    RoofArea            float64 `json:"roof_area_m2"`
    UpliftCoefficient   float64 `json:"uplift_coefficient"`
    TotalUpliftForce    float64 `json:"total_uplift_force_kn"`

    RequiredAnchorCapacity float64 `json:"required_anchor_capacity_kn"`
    SafetyFactor        float64 `json:"safety_factor"`
}
```

### Database Schema

```sql
-- Foundation configurations table
CREATE TABLE foundation_configs (
    id TEXT PRIMARY KEY,
    project_id TEXT NOT NULL REFERENCES projects(id),
    type TEXT NOT NULL,
    frost_depth_m REAL NOT NULL,
    design_wind_speed_kmh REAL NOT NULL,
    config_json TEXT NOT NULL,  -- JSON blob for type-specific config
    total_uplift_resistance_kn REAL,
    estimated_cost_cad REAL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- Anchor configurations table
CREATE TABLE anchor_configs (
    id TEXT PRIMARY KEY,
    foundation_id TEXT NOT NULL REFERENCES foundation_configs(id),
    type TEXT NOT NULL,
    quantity INTEGER NOT NULL,
    diameter_mm REAL,
    length_m REAL,
    spacing_m REAL,
    uplift_capacity_kn REAL,
    lateral_capacity_kn REAL
);

-- Wind load calculations table
CREATE TABLE wind_load_calculations (
    id TEXT PRIMARY KEY,
    foundation_id TEXT NOT NULL REFERENCES foundation_configs(id),
    design_wind_speed_kmh REAL NOT NULL,
    basic_pressure_pa REAL NOT NULL,
    design_pressure_pa REAL NOT NULL,
    roof_area_m2 REAL NOT NULL,
    total_uplift_force_kn REAL NOT NULL,
    required_anchor_capacity_kn REAL NOT NULL,
    calculated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);
```

---

## References

### Building Codes and Standards

1. **National Building Code of Canada (NBCC)**
   - Structural design requirements
   - Climate data and design values
   - Foundation frost protection requirements

2. **CSA S16 - Design of Steel Structures**
   - Anchor bolt design
   - Connection requirements

3. **CSA A23.3 - Design of Concrete Structures**
   - Concrete foundation design
   - Reinforcement requirements

### University Research

4. **University of Minnesota - Farm Scale Winter Greenhouse**
   - Deep winter greenhouse foundation designs (F.01, F.02, F.03)
   - Thermal mass integration
   - https://extension.umn.edu/growing-systems/deep-winter-greenhouses

5. **University of Saskatchewan - Greenhouse Design**
   - Prairie climate considerations
   - Frost depth requirements

### Industry Resources

6. **Canadian Greenhouse Conference**
   - Commercial greenhouse standards
   - https://canadiangreenhouseconference.com/

7. **Helical Pile Institute**
   - Installation standards
   - Load capacity guidelines
   - https://www.helicalpileinstitute.org/

8. **Post Frame Building Design Manual (NFBA)**
   - Post frame construction standards
   - Anchor and foundation requirements

### Technical References

9. **Frost-Protected Shallow Foundations (FPSF)**
   - ICC-ES Report ESR-1607
   - NAHB Research Center publications

10. **Ground Anchor Manufacturers**
    - American Earth Anchors
    - Penetrator anchors
    - Duckbill anchor systems

11. **Simpson Strong-Tie**
    - Post bases and caps
    - Anchor bolt specifications
    - https://www.strongtie.com/

### Climate Data

12. **Environment and Climate Change Canada**
    - Frost depth data
    - Design wind speeds
    - https://climate.weather.gc.ca/

13. **National Research Council Canada**
    - Climate data for building design
    - https://nrc.canada.ca/

### Cost and Material Resources

14. **RSMeans Construction Cost Data**
    - Canadian construction cost estimates
    - Foundation material pricing

15. **Canadian Construction Association**
    - Regional cost factors
    - https://www.cca-acc.com/
