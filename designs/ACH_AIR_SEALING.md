# Air Infiltration and Sealing for Greenhouse Applications

This document provides greenhouse-specific guidance on managing air infiltration to optimize heating efficiency while maintaining the ventilation flexibility essential for plant production. The focus is on practical techniques for Canadian climate conditions.

## Table of Contents

1. [Understanding Air Infiltration in Greenhouses](#understanding-air-infiltration-in-greenhouses)
2. [Greenhouse Envelope Leakage Points](#greenhouse-envelope-leakage-points)
3. [Air Sealing by Greenhouse Type](#air-sealing-by-greenhouse-type)
4. [Glazing-Specific Sealing Techniques](#glazing-specific-sealing-techniques)
5. [Vent and Opening Management](#vent-and-opening-management)
6. [Materials for Greenhouse Air Sealing](#materials-for-greenhouse-air-sealing)
7. [Assessing Your Greenhouse Air Leakage](#assessing-your-greenhouse-air-leakage)
8. [Target Values and Cost-Benefit Analysis](#target-values-and-cost-benefit-analysis)
9. [References](#references)

---

## Understanding Air Infiltration in Greenhouses

### Why Greenhouse Air Sealing Is Different

Greenhouses present unique challenges compared to residential buildings:

| Aspect | Residential Building | Greenhouse |
|--------|---------------------|------------|
| Envelope composition | Mostly opaque, insulated walls | Mostly transparent glazing |
| Airtightness goal | As tight as possible | Tight in winter, open in summer |
| Ventilation approach | Controlled mechanical | High-volume natural + mechanical |
| Temperature target | 20-22°C year-round | 10-25°C depending on crops |
| Humidity tolerance | 30-50% RH | 50-85% RH |
| Code requirements | Strict building codes | Agricultural exemptions |
| Envelope joints | Hundreds of meters | Thousands of meters of glazing bars |

**Key Insight:** The goal for greenhouses is not maximum airtightness, but *controlled* air exchange—tight when you need it (winter nights), open when you need it (summer days).

### The Greenhouse Air Exchange Spectrum

```
VERY LEAKY                                           VERY TIGHT
    |                                                     |
    |  Poly tunnel    Glass      Polycarbonate   Passive  |
    |  (old film)   (old, loose)  (well-sealed)   solar   |
    |                                                     |
    2-5 ACH         1-2 ACH       0.5-1 ACH      0.2-0.5 ACH
    (natural)       (natural)     (natural)      (natural)
```

**ACH (Air Changes per Hour)** measures how many times the entire air volume is replaced in one hour. For greenhouses, we typically discuss *natural* ACH (ACHnat)—the infiltration rate under normal wind and temperature conditions.

### How Air Infiltration Affects Heating Costs

#### The Infiltration Heat Loss Formula

```
Q_infiltration = ACH × V × 0.33 × ΔT

Where:
Q = Heat loss (Watts)
ACH = Natural air changes per hour
V = Greenhouse volume (m³)
0.33 = Volumetric heat capacity of air (W·h/m³·K)
ΔT = Temperature difference indoor vs outdoor (°C)
```

#### Example: Impact of Air Sealing on a 100 m² Greenhouse

**Scenario:** 100 m² greenhouse, 3 m average height (300 m³), Ottawa climate
- Indoor target: 15°C
- Design outdoor temperature: -25°C
- ΔT = 40°C

| Envelope Condition | ACHnat | Heat Loss (W) | Annual Cost* | Savings |
|-------------------|--------|---------------|--------------|---------|
| Very leaky (old poly) | 2.0 | 7,920 | $4,200 | — |
| Typical (loose glazing) | 1.0 | 3,960 | $2,100 | 50% |
| Well-sealed | 0.5 | 1,980 | $1,050 | 75% |
| Excellent (passive solar) | 0.25 | 990 | $525 | 87% |

*Assuming propane heat at $1.50/L, heating 180 days/year, 12 hours/day average

**The takeaway:** Reducing ACH from 2.0 to 0.5 can save $3,150/year in heating costs for a 100 m² greenhouse.

### The Three Drivers of Air Infiltration

1. **Wind Pressure**
   - Dominant in exposed locations
   - Creates positive pressure on windward side, negative on leeward
   - Higher wind = more infiltration
   - Windbreaks can reduce infiltration 20-40%

2. **Stack Effect (Thermal Buoyancy)**
   - Warm air rises, exits through high leaks
   - Cold air enters through low leaks
   - Greater effect with taller greenhouses
   - Dominant at night with little wind

3. **Mechanical Systems**
   - Exhaust fans create negative pressure
   - Unbalanced ventilation pulls air through leaks
   - Circulation fans can increase local infiltration at joints

---

## Greenhouse Envelope Leakage Points

### Ranked by Typical Leakage Contribution

Understanding where air leaks occur helps prioritize sealing efforts.

#### 1. Glazing Panel Joints (30-50% of total leakage)

**The Challenge:** A typical 10 m × 20 m greenhouse has 200+ meters of glazing bar joints.

| Joint Type | Typical Gap | Leakage Potential |
|------------|-------------|-------------------|
| Polycarbonate H-connector | 0-3 mm | Medium |
| Glazing bar cap | 1-5 mm | High |
| Panel overlap | 2-10 mm | Very High |
| Endwall glazing | 3-10 mm | Very High |

**Signs of leakage:**
- Condensation patterns along glazing bars
- Cold drafts near joints
- Snow melt patterns on exterior
- Dust/debris accumulation at joints

#### 2. Roof Vents and Ridge Vents (15-30%)

**The Challenge:** Vents are designed to open—making them seal well when closed is difficult.

| Vent Type | Closed Gap | Notes |
|-----------|------------|-------|
| Hinged roof vent | 2-8 mm | Warps over time |
| Continuous ridge vent | 3-15 mm | Often poorly sealed |
| Roll-up vent | 5-20 mm | Very difficult to seal |
| Louvered vent | 5-15 mm | Gaps between louvers |

**Common problems:**
- Vent opener mechanisms prevent full closure
- Weatherstripping flattens over time
- Vent frames warp from thermal cycling
- Wind lifts loose vent covers

#### 3. Doors and Access Points (10-20%)

**The Challenge:** Doors are used daily, seals wear quickly.

| Door Type | Typical Gaps | Wear Rate |
|-----------|--------------|-----------|
| Sliding door | 5-15 mm (perimeter) | Fast |
| Hinged door | 3-8 mm | Medium |
| Roll-up door | 10-30 mm | Very Fast |
| Double-door airlock | 3-5 mm each | Medium |

**High-traffic greenhouses:** Doors may be open 50+ times/day during the season.

#### 4. Foundation and Baseboard Connection (10-15%)

**The Challenge:** Transition from ground to glazing is complex.

| Foundation Type | Typical Gaps | Sealing Difficulty |
|-----------------|--------------|-------------------|
| Concrete stem wall | 2-5 mm | Medium |
| Treated timber | 3-10 mm | Medium |
| Ground-level (no foundation) | Variable | High |
| Knee wall (insulated) | 2-5 mm | Low |

**Issues:**
- Concrete is rough and irregular
- Wood shrinks and expands
- Sill plates rarely sealed properly
- Mouse entry points = air entry points

#### 5. Exhaust Fans and Louvers (5-15%)

**The Challenge:** Must open for ventilation but leak when closed.

| Opening Type | Closed Gap | Winter Solution |
|--------------|------------|-----------------|
| Gravity louver | 3-10 mm | Interior cover |
| Motorized damper | 1-3 mm | Better seal |
| Fixed louver | 5-15 mm | Remove/cap winter |
| Intake vent | 5-20 mm | Insulated cover |

#### 6. Structural Penetrations (5-10%)

**Includes:**
- Electrical conduit entries
- Water/irrigation pipe entries
- Heating pipe/duct entries
- Thermostat and sensor wiring
- Structural anchor bolts

**Rule of Thumb:** Each unsealed 25 mm diameter hole = 3-5 W heat loss at ΔT 40°C

---

## Air Sealing by Greenhouse Type

### Polyethylene Film Greenhouses (Tunnel/Hoop Houses)

**Characteristics:**
- Single or double-layer poly film
- Film attached to frame with wiggle wire/poly lock
- Roll-up or drop-down sides common
- 3-5 year film lifespan

**Air Sealing Strategy:**

| Priority | Action | Difficulty | Impact |
|----------|--------|------------|--------|
| 1 | Install double-poly with inflation blower | Medium | High |
| 2 | Seal endwalls properly | Medium | High |
| 3 | Seal base perimeter | Low | Medium |
| 4 | Install insulated endwall doors | Medium | Medium |
| 5 | Add rollup side barriers for winter | Low | Medium |

**Double-Poly Inflation:**

The single most effective improvement for poly houses.

```
                    Outside
                       |
    Film Layer 1  ============
                   ^  Air   ^
                   |  Gap   |  Inflation blower
                   v  10cm  v  (25-50 CFM)
    Film Layer 2  ============
                       |
                    Inside
```

**Benefits:**
- Creates dead air insulating layer (R-3.5 vs R-0.9)
- Inflated film resists wind infiltration
- Positive pressure prevents inward air leaks
- Reduces condensation dripping

**Cost:** $200-500 for inflation blower + $400-800 additional film

#### Endwall Sealing for Poly Houses

Endwalls are typically the leakiest part of poly greenhouses.

**Options (in order of preference):**

1. **Insulated Panel Endwall**
   - R-10 to R-20 foam sandwich panel
   - Properly framed and sealed
   - Reduces heat loss 40-60%

2. **Double-Poly Endwall**
   - Continuous with roof/sidewall film
   - Tensioned properly
   - Sealed at base and edges

3. **Polycarbonate Endwall**
   - More durable than poly
   - Better long-term seal
   - Easy to add doors/vents

### Multi-Wall Polycarbonate Greenhouses

**Characteristics:**
- 8-32 mm multi-wall polycarbonate panels
- Aluminum or steel glazing bar system
- Inherently more airtight than poly
- 10-20 year lifespan

**Air Sealing Strategy:**

| Priority | Action | Difficulty | Impact |
|----------|--------|------------|--------|
| 1 | Tape all panel ends (flute closure) | Low | High |
| 2 | Use proper gaskets in glazing bars | Low | High |
| 3 | Seal glazing bar caps | Low | Medium |
| 4 | Caulk panel-to-frame joints | Low | Medium |
| 5 | Install compression weatherstripping on vents | Medium | Medium |
| 6 | Seal foundation connection | Medium | Medium |

**Panel End Sealing (Critical):**

Polycarbonate flutes (channels) act as air pathways if unsealed.

```
    CORRECT:                    INCORRECT:

    [Solid tape - top]          [Open - top]
          |                          |
       || || ||                   || || ||
       || || ||  Air              || || || ← Air flows through
       || || ||  blocked          || || ||   flutes
       || || ||                   || || ||
          |                          |
    [Vent tape - bottom]        [Open - bottom]
```

**Top edge:** Solid aluminum tape (blocks moisture, insects, debris)
**Bottom edge:** Vent tape with micro-perforations (allows condensation drainage)

### Glass Greenhouses

**Characteristics:**
- Single, double, or insulated glass panels
- Metal (aluminum) or wood glazing bars
- Oldest greenhouse technology
- 25+ year lifespan

**Common Issues:**

1. **Glazing Compound Failure**
   - Traditional putty dries, cracks, falls out
   - Creates 3-10 mm gaps around each pane
   - Labor-intensive to repair

2. **Frame Deterioration**
   - Wood frames rot, create gaps
   - Aluminum frames bend, glazing loosens
   - Thermal expansion breaks seals

3. **Settling and Movement**
   - Structure shifts over decades
   - Glass panes no longer fit properly
   - Reglazing required

**Rehabilitation Options:**

| Option | Cost | Effectiveness | Notes |
|--------|------|---------------|-------|
| Reglaze with silicone | $$$ | Excellent | Professional job |
| Caulk all joints | $ | Good | DIY possible |
| Add interior poly layer | $$ | Very Good | Creates double envelope |
| Install insulated panels (opaque areas) | $$ | Excellent | Reduces glass area |

### Hybrid/Passive Solar Greenhouses

**Characteristics:**
- Insulated north/east/west walls
- Glazing only on south face
- Designed for maximum airtightness
- Often includes thermal mass

**Air Sealing Strategy:**

These greenhouses should achieve residential-level airtightness (0.3-0.5 ACHnat).

| Priority | Action | Difficulty | Impact |
|----------|--------|------------|--------|
| 1 | Install continuous air barrier on insulated walls | High | Critical |
| 2 | Tape all sheathing seams | Medium | High |
| 3 | Seal wall-to-glazing transition | High | High |
| 4 | Use high-quality vent weatherstripping | Medium | Medium |
| 5 | Caulk all penetrations | Low | Medium |
| 6 | Install gasketed electrical boxes | Low | Low |

**Wall-to-Glazing Transition:**

This is the most challenging detail in passive solar greenhouse design.

```
                Glazing
                   |
    [Sealant] → ===+===
                   |
    [Gasket] →  [Frame]  ← Aluminum or wood
                   |
    [Sealant] → ===+===
                   |
    [Air barrier] → [Sheathing]
                   |
                Insulation
```

**Key Principles:**
- Air barrier must be continuous from wall to glazing frame
- Use flexible sealant (allows movement)
- Gasket between glazing and frame
- Seal both interior and exterior transitions

---

## Glazing-Specific Sealing Techniques

### Polycarbonate Panel Sealing

#### End Sealing (Flute Closure)

**Materials Needed:**
- Solid aluminum tape (top edges)
- Vent tape with 1.5 mm perforations (bottom edges)
- U-channel or F-channel closure strips
- Isopropyl alcohol for cleaning

**Procedure:**

1. **Clean panel ends** with isopropyl alcohol
2. **Apply tape** overlapping onto panel faces 10-15 mm
3. **Press firmly** using roller or squeegee
4. **Install closure strip** over tape for protection

**Cost:** $2-4 per linear meter of panel edge

#### Panel-to-Frame Joints

**Glazing Tape System:**

```
    [Glazing bar cap]
           |
    [EPDM gasket] ← Creates seal
           |
    [Polycarbonate panel]
           |
    [Foam setting tape] ← Cushions and seals
           |
    [Glazing bar base]
```

**Materials:**
- Foam glazing tape: 6-10 mm thick × 10-20 mm wide
- EPDM compression gasket: Sized for glazing bar profile
- Silicone sealant (optional): For additional sealing

**Installation Tips:**
- Tape should compress 30-50% when cap installed
- Never stretch tape during application
- Apply in temperatures above 5°C
- Replace every 10-15 years

#### Panel-to-Panel Joints (H-Connectors)

**Standard H-Profile:**
- Insert foam tape into H-channel before panel installation
- Or apply bead of silicone after installation
- Ensure H-profile is continuous (no breaks)

**Cost:** $3-6 per linear meter including materials and labor

### Polyethylene Film Attachment

#### Poly-Lock Channel System

```
    [Wiggle wire]
         ↓
    ═══════════════  ← Film layer
    [  U-channel  ]  ← Attached to frame
    ═══════════════
         ↑
        Frame
```

**Sealing Tips:**
- Install channel on flat, clean surface
- Use double layer of wiggle wire for high-wind areas
- Caulk under channel before installation
- Overlap film 50-100 mm past channel

#### Base Perimeter Sealing

**Ground-Level Attachment:**

1. **Buried edge** (traditional)
   - Dig trench 150-200 mm deep
   - Extend film into trench
   - Backfill with soil or gravel
   - Simple but can rot film at soil line

2. **Timber base board**
   - Attach film to treated 38 × 140 mm board
   - Bury board or weight with soil
   - Easier to replace film

3. **Concrete curb**
   - Film attaches with poly-lock on curb face
   - Most durable option
   - Provides rodent/pest barrier

### Glass Glazing Sealing

#### Traditional Putty Replacement

**When to Replace:**
- Putty is cracked, loose, or missing
- Air leaks detectable around panes
- Condensation between panes and frame

**Modern Alternatives to Putty:**

| Material | Lifespan | Flexibility | Cost |
|----------|----------|-------------|------|
| Traditional oil putty | 10-15 years | Low | $ |
| Acrylic glazing compound | 15-20 years | Medium | $$ |
| Silicone sealant | 20-30 years | High | $$ |
| Preformed tape + cap | 25+ years | High | $$$ |

#### Dry Glazing Systems

For aluminum-frame glass greenhouses:

**Components:**
- EPDM setting blocks (support glass weight)
- EPDM wedge gaskets (compress against glass)
- Snap-on glazing cap (holds gaskets in place)

**Maintenance:**
- Check gasket compression annually
- Replace gaskets if cracked or flattened
- Ensure cap clips are secure

---

## Vent and Opening Management

### Roof Vent Sealing

#### Hinged Roof Vents

**Problem:** Gap between vent frame and greenhouse frame when closed.

**Solutions:**

1. **Compression Weatherstripping**
   - Install closed-cell EPDM on vent frame
   - Should compress 3-5 mm when closed
   - Replace every 3-5 years

2. **Bulb Seal Gasket**
   - Hollow rubber bulb compresses more easily
   - Better for warped frames
   - More expensive but longer lasting

3. **Brush Seal**
   - For vents with large gaps
   - Allows some air but blocks drafts
   - Good for transition seasons

**Automatic Vent Openers:**

Wax-cylinder vent openers can prevent full closure.

**Adjustment Tips:**
- Set minimum opening to 0 mm (full close)
- Use triple-spring models for positive closure
- Remove openers in winter if vents not needed
- Check adjustment annually

#### Continuous Ridge Vents

**Challenge:** Long linear opening very difficult to seal.

**Options:**

| Option | Effectiveness | Cost | Notes |
|--------|---------------|------|-------|
| Foam gasket on vent flap | Medium | $ | Compresses over time |
| Motorized damper inserts | High | $$$ | Best for automation |
| Winter cover panel | High | $$ | Seasonal installation |
| Interior baffle + curtain | Medium | $$ | Reduces convective loss |

### Side Vent and Roll-Up Wall Management

**Roll-Up Sides:**

Roll-up sides are inherently leaky. Winter strategies:

1. **Drop and seal for winter**
   - Unroll completely
   - Attach with poly-lock at base
   - Tape/seal overlaps
   - Creates fixed wall for winter

2. **Install interior curtain**
   - Thermal/shade curtain inside
   - Creates dead air buffer
   - Can be automated

3. **Add temporary insulated panel**
   - Foam board or bubble wrap
   - Attached inside roll-up wall
   - Remove in spring

### Door Weatherstripping

#### Entry Door Specifications

**Minimum Requirements:**

| Door Component | Material | Gap When Closed |
|----------------|----------|-----------------|
| Hinge side | EPDM compression | < 3 mm |
| Latch side | EPDM compression | < 3 mm |
| Top | EPDM or brush | < 5 mm |
| Threshold | Door sweep + sill | < 5 mm |

**Commercial Greenhouse Doors:**

Consider upgrading to:
- Insulated steel door (R-10+)
- Adjustable aluminum threshold
- Triple-seal weatherstripping
- Automatic door closer

**Vestibule/Airlock:**

For high-traffic greenhouses in cold climates:

```
    [Exterior] → [Door 1] → [Vestibule] → [Door 2] → [Interior]
                             (1-2 m²)
```

**Benefits:**
- Reduces air exchange per entry by 60-80%
- Transition zone for temperature acclimation
- Storage for tools/supplies

### Exhaust Fan Openings

#### Winter Management Options

| Option | Cost | Effectiveness | Effort |
|--------|------|---------------|--------|
| Interior insulated cover | $ | Excellent | High (install/remove) |
| Motorized insulated damper | $$$ | Excellent | Low (automated) |
| Gravity louvers + interior cover | $$ | Good | Medium |
| Remove fan, install solid panel | $ | Excellent | High |

#### Insulated Fan Cover Construction

**DIY Approach:**

Materials:
- 50 mm rigid foam board
- Plywood or OSB backing
- Foam weatherstrip tape
- Latches or wing nuts

Construction:
1. Cut foam to match fan opening + 50 mm overlap
2. Attach plywood backing
3. Apply foam tape around perimeter
4. Add latches for quick installation/removal

**Cost:** $30-50 per cover

---

## Materials for Greenhouse Air Sealing

### Tapes (UV-Resistant Required)

| Product | Type | Width | Cost (CAD) | UV Life | Best Use |
|---------|------|-------|------------|---------|----------|
| Aluminum foil tape | Foil + acrylic | 50-75 mm | $15-25/roll | 20+ years | Polycarbonate ends |
| Poly repair tape | Polyethylene | 75-100 mm | $20-35/roll | 3-5 years | Film patches |
| Greenhouse tape (white) | Polyethylene | 25-50 mm | $15-25/roll | 5-7 years | Film seams |
| Acrylic glazing tape | Acrylic adhesive | 50-75 mm | $40-80/roll | 15+ years | Glazing bars |
| EPDM seam tape | Rubber | 75 mm | $50-90/roll | 20+ years | Membrane seams |

**Critical:** Standard duct tape, masking tape, and electrical tape will fail within months in greenhouse UV exposure.

### Sealants and Caulks

| Product | Type | Temp Range | UV Resistance | Best Use |
|---------|------|------------|---------------|----------|
| Silicone (100%) | Neutral cure | -50 to +200°C | Excellent | Permanent joints |
| Silicone (hybrid) | Polyether | -40 to +100°C | Excellent | Moving joints |
| Polyurethane | PU | -40 to +80°C | Good | Adhesive sealing |
| Acrylic latex | Latex | -20 to +80°C | Fair | Interior only |
| Butyl rubber | Butyl | -30 to +100°C | Good | Tape-like application |

**Polycarbonate Compatibility:**

- ✓ Use: Neutral-cure silicone, polyurethane (check label)
- ✗ Avoid: Acetic-cure silicone (vinegar smell), solvent-based caulks

**Cost:** $8-15 per 300 ml tube, 10-15 m coverage per tube (3 mm bead)

### Gaskets and Weatherstripping

#### Glazing Gaskets

| Type | Shore Hardness | Compression Range | Cost |
|------|----------------|-------------------|------|
| EPDM solid | 60-70A | 20-30% | $3-6/m |
| EPDM sponge | 40-50A | 40-60% | $2-4/m |
| Silicone | 40-60A | 30-50% | $5-10/m |
| TPE (thermoplastic) | 50-70A | 20-40% | $2-5/m |

**Sizing:**
- Measure gap when closed
- Select gasket 20-40% larger than gap
- Test compression before full installation

#### Foam Setting Tapes

For cushioning glazing panels in frames:

| Product | Thickness | Width | Density | Cost |
|---------|-----------|-------|---------|------|
| Closed-cell PVC | 3-6 mm | 10-20 mm | Medium | $1-2/m |
| Closed-cell EPDM | 6-10 mm | 10-25 mm | Medium | $2-4/m |
| Closed-cell silicone | 3-6 mm | 10-20 mm | Low | $4-8/m |

### Poly Attachment Hardware

| System | Components | Cost | Lifespan |
|--------|------------|------|----------|
| Wiggle wire + U-channel | Wire, aluminum channel | $2-4/m | 15+ years |
| Poly-lock (spring wire) | Spring wire, aluminum channel | $3-5/m | 15+ years |
| Batten + screw | Wood batten, screws | $1-2/m | 5-10 years |
| Rope in channel | Rope, plastic channel | $1-3/m | 10-15 years |

**Inflation Blowers (Double-Poly):**

| Size | CFM | Power | Cost | Coverage |
|------|-----|-------|------|----------|
| Small | 25-50 | 15-25 W | $150-250 | Up to 200 m² |
| Medium | 50-100 | 25-50 W | $200-350 | 200-500 m² |
| Large | 100-200 | 50-100 W | $300-500 | 500+ m² |

**Features to look for:**
- Filtered air intake
- Backdraft damper
- Speed control (optional)
- Weather-resistant housing

---

## Assessing Your Greenhouse Air Leakage

### Visual Inspection Checklist

Walk through your greenhouse on a cold, calm day:

**Glazing System:**
- [ ] Panel ends sealed (polycarbonate flutes closed)?
- [ ] Glazing bar caps secure and seated?
- [ ] Visible gaps at panel-to-frame joints?
- [ ] Cracked or missing sealant?
- [ ] Damaged or loose panels?

**Vents and Openings:**
- [ ] Roof vents close fully?
- [ ] Weatherstripping intact on vents?
- [ ] Exhaust fan louvers close fully?
- [ ] Gaps around vent frames?

**Doors:**
- [ ] Door sweeps in good condition?
- [ ] Weatherstripping on all four sides?
- [ ] Visible light around door perimeter?
- [ ] Door closes and latches properly?

**Foundation/Base:**
- [ ] Gaps at base of glazing?
- [ ] Visible holes or cracks in foundation?
- [ ] Sill sealed to foundation?
- [ ] Mouse holes or pest entry points?

**Penetrations:**
- [ ] Electrical entries sealed?
- [ ] Pipe entries sealed?
- [ ] Gaps around structural posts?

### Smoke Testing (Leak Detection)

**Simple Method:**

1. Close all vents, doors, windows
2. Turn on one exhaust fan to depressurize greenhouse
3. Walk perimeter with smoke pencil or incense
4. Observe smoke movement—indicates air entry points
5. Mark leaks for later sealing

**Materials Needed:**
- Smoke pencil ($15-25) or incense sticks
- Tape to mark leak locations
- Notepad for recording findings

**Best Conditions:**
- Calm day (wind < 10 km/h)
- Temperature differential (cold outside)
- Early morning or evening

### Thermal Imaging

If you have access to a thermal camera (or smartphone attachment):

**What to Look For:**
- Cold spots at joints and edges = air leaks
- Blue/purple colors indicate heat loss
- Compare similar areas to identify problems

**Rental Options:**
- FLIR ONE smartphone attachment: $250-400 (or rent)
- Professional thermal camera: $50-100/day rental
- Energy audit service: $200-400 (includes report)

**Best Time:**
- Before sunrise on cold morning
- At least 10°C temperature difference
- No direct sun on glazing

### Estimating Your Current ACH

#### Method 1: Comparison to Similar Structures

| If Your Greenhouse Is... | Estimated ACHnat |
|--------------------------|------------------|
| New double-poly, well-sealed | 0.3-0.5 |
| New polycarbonate, well-sealed | 0.4-0.7 |
| Older polycarbonate, typical | 0.7-1.2 |
| Old glass, loose glazing | 1.5-3.0 |
| Poly with roll-up sides (closed) | 1.0-2.0 |
| Any type with obvious leaks | 2.0+ |

#### Method 2: Fuel Consumption Back-Calculation

If you know your heating fuel usage:

1. Calculate total heat loss from fuel consumption
2. Subtract estimated glazing/wall conduction loss
3. Remainder = infiltration loss
4. Back-calculate ACH

*This method is approximate but useful for tracking improvement.*

#### Method 3: CO2 Decay Test (Advanced)

1. Raise CO2 to 1,500-2,000 ppm (with burner or tank)
2. Seal greenhouse, turn off CO2 source
3. Measure CO2 decay over time
4. Natural dilution rate = air exchange rate

**Formula:**
```
ACH = ln(C1/C2) / (t2 - t1)

Where:
C1 = Initial CO2 concentration
C2 = Final CO2 concentration
t = Time in hours
```

---

## Target Values and Cost-Benefit Analysis

### Recommended ACH Targets by Greenhouse Type

| Greenhouse Type | Target ACHnat | Why This Target |
|-----------------|---------------|-----------------|
| Passive solar (year-round) | 0.2-0.4 | Maximum heating efficiency |
| Heated polycarbonate | 0.4-0.7 | Good efficiency, practical to achieve |
| Heated glass | 0.5-1.0 | Limited by glazing system |
| Double-poly heated | 0.3-0.6 | Inflation provides good seal |
| Three-season (frost protection) | 0.7-1.5 | Less critical |
| Unheated cold frame | N/A | Airtightness not a priority |

### Climate Zone Adjustment (Canada)

| Zone | HDD (base 18°C) | Heating Priority | Recommended Max ACHnat |
|------|-----------------|------------------|----------------------|
| Zone 4 | 2500-3999 | Medium | 0.8 |
| Zone 5 | 4000-4999 | Medium-High | 0.6 |
| Zone 6 | 5000-5999 | High | 0.5 |
| Zone 7+ | 6000+ | Very High | 0.4 |

### Cost-Benefit Analysis

#### Air Sealing Investment Packages

**Basic Package (DIY, ~$400-700 for 100 m² greenhouse):**

| Item | Quantity | Cost |
|------|----------|------|
| Aluminum tape for panel ends | 4 rolls | $80 |
| Foam glazing tape | 50 m | $100 |
| Silicone sealant | 6 tubes | $60 |
| Door weatherstripping | 2 doors | $60 |
| Vent weatherstripping | 4 vents | $80 |
| Miscellaneous (foam, tape) | — | $50 |
| **Total** | | **$430** |

**Expected Improvement:** Reduce ACH by 30-50%

**Comprehensive Package (DIY + some professional, ~$1,200-2,000):**

| Item | Quantity | Cost |
|------|----------|------|
| Basic package | — | $430 |
| Inflation blower (if poly) | 1 | $250 |
| Insulated door upgrade | 1 | $400 |
| Exhaust fan covers | 2 | $100 |
| Professional vent adjustment | 1 visit | $200 |
| **Total** | | **$1,380** |

**Expected Improvement:** Reduce ACH by 50-70%

#### Payback Period Calculation

**Formula:**
```
Payback (years) = Investment ($) / Annual Heating Savings ($)

Annual Savings = Original Heating Cost × % Reduction
```

**Example:**

- Investment: $1,000 in air sealing
- Original heating cost: $3,000/year
- ACH reduced from 1.5 to 0.6 (60% reduction)
- Infiltration portion of heat loss: ~40% of total
- Heating cost reduction: 40% × 60% = 24%
- Annual savings: $3,000 × 24% = $720
- Payback: $1,000 / $720 = 1.4 years

**Typical Payback by Climate Zone:**

| Zone | Typical Payback | Notes |
|------|-----------------|-------|
| Zone 4 | 2-4 years | Lower heating costs reduce savings |
| Zone 5 | 1.5-3 years | Good ROI |
| Zone 6 | 1-2 years | Excellent ROI |
| Zone 7+ | <1-1.5 years | Very high heating costs, fast payback |

### Diminishing Returns

**Important:** Beyond a certain point, additional air sealing provides little benefit.

```
Heating Cost Reduction vs. Air Sealing Investment

100%|
    |
 75%|        ___________
    |      /
 50%|    /
    |  /
 25%|/
    |__________________|____
    $0     $500    $1000   $2000
         Investment
```

**Guidelines:**
- First $300-500: High impact, 20-40% cost reduction
- Next $500-1000: Moderate impact, additional 10-20%
- Beyond $1500: Diminishing returns in most cases

**Exception:** Passive solar greenhouses benefit from investment up to $3000-5000 in envelope quality.

---

## References

### Greenhouse Heating and Ventilation

- [Manitoba Agriculture - Greenhouse Heating and Venting](https://www.gov.mb.ca/agriculture/crops/crop-management/print,heating-and-venting.html)
- [Purdue University - Calculating Greenhouse Heating Requirements](https://www.purdue.edu/hla/sites/cea/article/calculating-greenhouse-heating-requirements/)
- [University of Florida - Greenhouse Ventilation](https://edis.ifas.ufl.edu/publication/AE030)
- [ANSI/ASAE EP406.4 - Heating, Ventilating and Cooling Greenhouses](https://ceac.arizona.edu/sites/default/files/asae_-_heating_ventilating_and_cooling_greenhouses.pdf)

### Glazing and Sealing Materials

- [Polygal - Polycarbonate Installation Guide](https://polygal.com/installation-guidelines/)
- [GE Sealants - How to Seal a Greenhouse](https://gesealants.com/projects-howtos/how-to-seal-a-greenhouse-for-maximum-insulation/)
- [Farmtek - Greenhouse Poly Film Guide](https://www.farmtek.com/farm/supplies/cat1;ft_poly_film.html)

### Passive Solar Greenhouse Design

- [Verge Permaculture - Passive Solar Greenhouse Design](https://vergepermaculture.ca/designing-your-own-passive-solar-greenhouse-part-3/)
- [Ceres Greenhouse - Building Envelope Best Practices](https://ceresgs.com/building-envelope-best-practices/)
- [The Year-Round Solar Greenhouse (Lindsey Schiller, Marc Plinke)](https://www.newsociety.com/books/y/the-year-round-solar-greenhouse)

### Canadian Agricultural Resources

- [Ontario Ministry of Agriculture - Greenhouse Production](https://www.ontario.ca/page/greenhouse-production)
- [Alberta Agriculture - Greenhouse Climate Control](https://www.alberta.ca/greenhouse-climate-control)
- [Wisconsin Extension - Greenhouse Infiltration Losses](https://fyi.extension.wisc.edu/energy/greenhouses/infiltration-losses/)

### Air Sealing Products

- [Tuck Tape (Cantech)](https://www.cantechindustries.com/tuck-tape)
- [3M Industrial Tapes - Technical Data](https://www.3m.com/3M/en_US/company-us/all-3m-products/~/All-3M-Products/Industrial/Tapes/)
- [Tremco - Construction Sealants](https://www.tremcosealants.com/)

---

*Document Version: 2.0*
*Last Updated: February 2026*
*For use with honk4greenhouse passive greenhouse design application*
