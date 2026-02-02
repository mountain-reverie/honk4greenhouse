# CLAUDE.md

This file provides guidance to Claude Code (claude.ai/code) when working with code in this repository.

## Important Guidelines

**Measurement Units**: ALWAYS use metric units (meters, Celsius, Watts, cubic meters per hour, etc.) for all calculations and measurements, as appropriate for Canadian greenhouse design standards and building codes.

## Project Overview

### Technical Overview

This is a Go web application called "honk4greenhouse" - a website for designing efficient and potentially passive greenhouses optimized for Canadian climate conditions. The project prioritizes maintainability and API stability, leveraging Go-native solutions wherever possible.

#### Core Technology Stack

**Backend & Framework**

- [Gin](https://gin-gonic.com/): HTTP web framework with middleware support
- [Goth](https://github.com/markbates/goth): Multi-provider OAuth authentication ([introduction blog](https://dizzy.zone/2018/06/01/OAuth-with-Gin-and-Goth/))

**Frontend & UI**

- [Templ](https://templ.guide/): Type-safe HTML templating system
- [Templ UI](https://templui.io/): Component library built on Templ and [HTMX](https://htmx.org/)
- [HTMX](https://htmx.org/): Dynamic web interactions without JavaScript

**Database & Infrastructure**

- [Turso](https://turso.tech/): SQLite-compatible database with local replicas on each host
- [Pulumi](https://www.pulumi.com/docs/iac/languages-sdks/go/): Infrastructure as Code for setup and maintenance

**Testing & Monitoring**

- [Playwright CI Go](https://github.com/mountain-reverie/playwright-ci-go): End-to-end web testing framework
- OpenTelemetry: Comprehensive observability (logs, metrics, traces)
- [SigNoz](https://signoz.io/docs/install/self-host/): Self-hosted observability platform
- [Beszel](https://www.youtube.com/watch?v=O_9wT-5LoHM): System monitoring solution

#### Development Philosophy

**Maintainability First**: The codebase is designed for long-term stability with automatic dependency updates via Dependabot. All dependencies are selected for their proven track record of API stability and backward compatibility.

**Testing Strategy**: Comprehensive integration and end-to-end testing with performance benchmarking. Go PGO (Profile-Guided Optimization) uses CI benchmark results for optimization. Turso's local replica capability enables full end-to-end testing with isolated database instances.

**CI/CD Pipeline**: GitHub Actions handle continuous integration and deployment with strict quality gates, including benchmark regression detection to prevent performance degradation and maintain security standards.

#### Deployment Architecture

**Infrastructure**: Deployed exclusively on European and Canadian cloud providers ([OVH](https://www.ovhcloud.com/en-ca/)) across 3 machines:

- 1 monitoring server (SigNoz + Beszel)
- 2 application servers (load balanced)

**Networking**:

- Cloudflare Tunnel for secure internet exposure
- Tailscale for VPN, administrative access, and core VPC management

**Reference Videos** (for understanding Tailscale deployment patterns):

- [How to use cloud-init and Tailscale | Infrastructure as Code Series Part 1](https://www.youtube.com/watch?v=e-X5FJwrkaA)
- [Automate your Tailscale cloud deployments with Terraform | Infrastructure as Code Series Part 2](https://www.youtube.com/watch?v=PEoMmZOj6Cg)
- [An Ansible primer for Devops | Infrastructure as Code Series Part 3](https://www.youtube.com/watch?v=k5Xgt31yK2U)
- [Static site deployments made easy with Github Actions and Tailscale](https://www.youtube.com/watch?v=OQJAX-Ce1YY)

**Security**: GitHub Action environments segregate secrets with progressive exposure based on CI validation results. Performance regression gates prevent potentially compromised code from reaching production environments.

## Build and Development Commands

Since this is a Go project, use standard Go commands:

- `go run cmd/service/main.go` - Run the main service
- `go build cmd/service/main.go` - Build the service binary
- `go test ./...` - Run all tests (when tests are added)
- `go mod tidy` - Clean up module dependencies

## Architecture

- **Entry Point**: `cmd/service/main.go` - Main function and HTTP server startup
- **Module**: `github.com/mountain-reverie/honk4greenhouse` (Go 1.23.6)
- **Structure**: Standard Go project layout following best practices
  - `cmd/service/` - Application entry point
  - `internal/handlers/` - HTTP request handlers
  - `internal/server/` - Server configuration and routing

**Current Implementation**: Basic Gin HTTP server with health check and welcome endpoints. The server listens on port 8080 (configurable via PORT environment variable) and provides:
- `/` - Welcome message endpoint
- `/health` - Health check endpoint for monitoring

## Greenhouse Thermal Performance Design Resources

This section provides comprehensive technical resources for implementing greenhouse thermal performance calculations, heat gain/loss analysis, and passive solar design optimization for Canadian climate conditions.

### Heat Loss Calculation Methodology

**Primary Heat Loss Mechanisms:**
- Conduction through glazing materials (dominant factor)
- Air infiltration through gaps and ventilation
- Radiation heat loss during nighttime periods

**Core Formula for Total Heat Loss:**
```text
QT = QC + QA
Where:
QT = Total heat loss (Watts)
QC = Heat loss through conduction
QA = Heat loss through natural air exchange
```

**Conduction Heat Loss Calculation:**
```text
Q = A × (Ti - To) / R
Where:
Q = Heat loss (Watts)
A = Surface area (m²)
Ti = Inside air temperature (°C)
To = Outside air temperature (°C)
R = Thermal resistance of material (m²·K/W)
```

**References:**
- [Manitoba Agriculture - Greenhouse Heating and Venting](https://www.gov.mb.ca/agriculture/crops/crop-management/print,heating-and-venting.html)
- [Purdue University - Calculating Greenhouse Heating Requirements](https://www.purdue.edu/hla/sites/cea/article/calculating-greenhouse-heating-requirements/)
- [Greenhouse Management - Determining Heat Loss](https://www.greenhousemag.com/article/technology-determining-greenhouse-heat-loss/)

### Glazing Material Thermal Properties

**Polycarbonate (Recommended for Canadian Climates):**
- Twin-wall (8mm): R-value = 1.54, Light transmission = ~85%
- Triple-wall (16mm): R-value = 2.4, Light transmission = 77%
- Five-wall (32mm): R-value = 5.6, Light transmission = ~70%
- Thermal conductivity: 0.20 W/(k·m)

**Glass:**
- Single-pane (6mm): R-value = 0.16, Light transmission = ~90%
- Double-pane: R-value = ~2.0, Light transmission = ~80%
- Thermal conductivity: 0.80 W/(k·m)

**Polyethylene:**
- Single layer: R-value = 0.83
- Double layer (air-inflated): R-value = 2.0
- Typical lifespan: 3-4 years

**References:**
- [Ceres Greenhouse Solutions - Glazing Material Selection](https://ceresgs.com/how-to-choose-a-glazing-material-for-a-year-round-greenhouse/)
- [Greenhouse Catalog - Insulation Comparison](https://www.greenhousecatalog.com/greenhouse-insulation)
- [University of Arkansas - Glazing Materials](https://greenhouse.hosted.uark.edu/Unit03/Printer_Friendly.html)

### Passive Solar Design Principles

**Optimal Orientation:**
- South-facing orientation for maximum winter solar gain
- Long axis running east-west with glazed south face
- North wall typically insulated or earth-bermed

**Solar Heat Gain Calculation:**
- Solar radiation intensity varies by latitude, season, and time of day
- Glazing orientation factor affects total energy capture
- Shading and overhang design for summer cooling

**Thermal Mass for Heat Storage:**
- Water: 1 BTU/lb·°F (4.18 kJ/kg·K) - highest thermal capacity
- Concrete: ~0.2 BTU/lb·°F (0.84 kJ/kg·K)
- Rock/gravel: ~0.2 BTU/lb·°F (0.84 kJ/kg·K)
- Placement: North wall positioning for maximum solar exposure

**Climate Battery Systems (SHCS):**
- Subterranean heating/cooling using buried pipe networks
- Captures excess daytime heat for nighttime release
- Ground temperature stability (8-12°C at 2m depth in Canada)

**References:**
- [Verge Permaculture - Passive Solar Greenhouse Design](https://vergepermaculture.ca/designing-your-own-passive-solar-greenhouse-part-3/)
- [UGA Extension - Passive Solar Greenhouse Construction](https://extension.uga.edu/publications/detail.html?number=B1566)
- [Ceres Greenhouse - Thermal Mass Systems](https://ceresgs.com/tips-on-using-water-barrels-in-a-solar-greenhouse/)

### Canadian Climate Data Resources

**Primary Data Sources:**
- [Environment and Climate Change Canada - Historical Climate Data](https://climate.weather.gc.ca/)
- [Climate Atlas of Canada](https://climateatlas.ca/) - Regional climate projections
- [Open Government Portal - Heating Degree Days](https://open.canada.ca/data/en/dataset/fd8efb83-b73d-5442-ab60-7987c824f5fd)

**Key Climate Metrics:**
- **Heating Degree Days (HDD)**: Threshold 18°C, indicates heating requirements
- **Cooling Degree Days (CDD)**: Threshold 18°C, indicates cooling needs
- **Growing Degree Days (GDD)**: Crop-specific temperature accumulation
- **Design Temperatures**: Winter 2.5% and summer 1% values for equipment sizing

**Regional Building Climate Zones (NECB):**
- Zone 4: Southern BC Coast (< 3000 HDD)
- Zone 5: Toronto, Victoria (3000-3999 HDD)
- Zone 6: Ottawa, Montreal (4000-4999 HDD)
- Zone 7A/7B: Edmonton, Winnipeg, Calgary (5000-6999 HDD)
- Zone 8: Yellowknife, Northern regions (≥ 7000 HDD)

### Plant Hardiness Zones

Plant hardiness zones are essential for crop selection and understanding minimum winter temperatures. **Note:** USDA and Canadian systems are NOT interchangeable.

**USDA System (United States):**
- Based solely on average annual extreme minimum temperature
- Zones 1a to 13b (10°F increments, 5°F half-zones)
- Zone lookup by ZIP code: https://planthardiness.ars.usda.gov/
- Free API: https://phzmapi.org/{ZIPCODE}.json

**Canadian System:**
- Based on 7 climate variables (min/max temp, rainfall, snow depth, wind, growing season)
- Zones 0a to 9a (multivariate index calculation)
- Zone lookup by municipality: https://planthardiness.gc.ca/
- ~50% of Canada is Zone 0a (northern territories, unsuitable for most horticulture)

**Approximate Conversion:** Add 1 zone when converting USDA to Canadian (e.g., USDA Zone 4 ≈ Canadian Zone 5)

**Data Downloads:**
- USDA shapefiles & CSV: https://prism.oregonstate.edu/phzm/
- Canadian raster data: https://open.canada.ca/data/en/dataset/db9b4130-8893-11e0-9b96-6cf049291510

See [designs/CLIMATE_ZONES.md](designs/CLIMATE_ZONES.md) for comprehensive zone documentation, temperature tables, the Canadian formula, and implementation details.

### Building Energy Simulation Software

**TRNSYS (Transient System Simulation Tool):**
- Specialized for complex thermal systems and renewable energy
- Flexible graphical interface for component-based modeling
- Excellent for greenhouse-specific thermal analysis
- Website: [trnsys.com](http://www.trnsys.com/)

**EnergyPlus:**
- DOE whole-building energy simulation program
- Comprehensive HVAC and renewable energy modeling
- Free and open-source with extensive documentation
- Website: [energyplus.net](https://energyplus.net/)

**OpenStudio:**
- Graphical interface for EnergyPlus
- Cross-platform SDK with advanced visualization
- Integrated daylight analysis using Radiance
- Developed by NREL for building performance simulation

**Application to Greenhouse Modeling:**
- TRNSYS: Superior for transient greenhouse thermal behavior
- EnergyPlus: Good for HVAC integration and energy analysis
- OpenStudio: User-friendly interface for EnergyPlus modeling

**References:**
- [ScienceDirect - TRNSYS Greenhouse Modeling](https://www.sciencedirect.com/science/article/abs/pii/S2352710219306631)
- [Better Buildings Initiative - OpenStudio](https://betterbuildingssolutioncenter.energy.gov/solutions-at-a-glance/openstudio-energyplus-software-whole-building-energy-modeling)

### Economic Analysis Framework

**Capital Costs:**
- Glazing materials and structure ($160-540/m² depending on material)
- Thermal mass systems ($20-85/m² for water walls)
- Insulation and thermal curtains ($30-130/m²)
- HVAC and climate control systems (variable)

**Operating Costs:**
- Heating energy (natural gas, electricity, propane)
- Cooling energy (ventilation, evaporative cooling)
- Maintenance and glazing replacement
- Labor for thermal management

**Performance Metrics:**
- Annual energy consumption (kWh/m²)
- Peak heating/cooling loads (W/m²)
- Payback period for energy efficiency improvements
- Life-cycle cost analysis over 20-30 year period

**Cost-Benefit Analysis:**

- Energy savings vs. capital investment
- Productivity gains from improved climate control
- Season extension value for crop production
- Carbon footprint and environmental impact assessment

### Greenhouse Lighting Requirements and Calculations

**Daily Light Integral (DLI) Fundamentals:**

DLI measures the number of photosynthetically active photons (400-700 nm) delivered to a specific area over 24 hours, expressed in mol/(m²·day).

**Core DLI Calculation Formula:**

```text
DLI = PPFD × light hours per day × (3600/1,000,000)
Or: DLI (mol/(m²·day)) = 3.6×10⁻³ × PPFD (μmol/(m²·s)) × Light-hours/day
```

**Canadian Geographic Context:**

- Monthly-averaged DLI values at 60° latitude (much of Canada): 1-40 mol/(m²·day)
- Winter months (Oct-Mar): DLI becomes limiting factor for greenhouse crops
- Greenhouse structures absorb/reflect 30-70% of outside light
- Maximum greenhouse DLI rarely exceeds 30 mol/(m²·day)

**Crop-Specific DLI Requirements:**

- **Leafy Greens/Herbs**: Minimum 17 mol/(m²·day)
- **Flowering Crops (Tomatoes/Peppers)**: 20-40 mol/(m²·day)
- **Ornamental Plants**: 10-20 mol/(m²·day)
- **Cannabis**: 35-65 mol/(m²·day)

**Supplemental Lighting Calculation:**

```text
Required Supplemental PPFD = (Target DLI - Natural DLI) / (Light Hours × 3.6×10⁻³)
```

**LED vs HPS Energy Efficiency:**

- LED systems use 35-55% less energy than HPS lights
- LED recommended power: 270-430 watts per m² for flowering plants
- Typical greenhouse LED intensity: 70-90 μmol·m⁻²·s⁻¹ (ornamentals), 140+ μmol·m⁻²·s⁻¹ (vegetables)

**Energy Consumption Calculation:**

```text
Daily Power Consumption (kWh) = Power Rating (kW) × Operating Hours (h)
Monthly Cost ($) = Daily kWh × Days × Electricity Rate ($/kWh)
```

**References:**

- [NCBI - Energy Requirement for Supplemental Greenhouse Lighting](https://www.ncbi.nlm.nih.gov/pmc/articles/PMC10934181/)
- [Hortinergy - Greenhouse Lighting Calculator](https://www.hortinergy.com/features/greenhouse-calculator/greenhouse-lighting-calculator/)
- [P.L. Light Systems - Light Energy Calculator](https://pllight.com/lightenergy-calculator/)

### Solar Radiation Management and Shading

**Excessive Heat from Solar Radiation:**

- Maximum solar heat absorption: 950 W/m² on dark surfaces
- Greenhouse floor can absorb up to 950 W/m²
- Spectral distribution: UV (5%), PAR (45%), NIR (50%)
- Only PAR (400-700 nm) is useful for photosynthesis

**Shading Effectiveness:**

- Shade screens: Reflect 10-70% of incoming solar radiation
- Paint-on compounds: Block 10-50% of radiation
- 50% shading typically reduces temperature rise proportionally
- Smart controllers optimize shading based on real-time solar radiation

**Shading Calculation Methods:**

```text
Heat Reduction = Solar Input × Shading Factor × Surface Area
Where Shading Factor = 0.1 to 0.7 (10% to 70% reduction)
```

**Canadian Climate Considerations:**

- Northern Canada: Extremely long summer days require intensive shading
- Site shading can affect building energy by up to 90% for cooling
- Shading area typically calculated as 0-50% of total envelope
- Optimal shading: <50% of total greenhouse area for plant health

**Natural Ventilation Requirements:**

- Standard sizing: 60 air changes per hour for heat elimination
- Combined with evaporative cooling for enhanced effectiveness
- Ventilation must increase proportionally with solar heat gain

**Advanced Cooling Strategies:**

- Evaporative cooling systems
- Automated shade cloth deployment
- Reflective glazing materials
- Strategic building orientation

**References:**

- [UGA Extension - Greenhouse Heating, Cooling and Ventilation](https://extension.uga.edu/publications/detail.html?number=B792)
- [Manitoba Agriculture - Greenhouse Heating and Venting](https://www.gov.mb.ca/agriculture/crops/crop-management/print,heating-and-venting.html)
- [MDPI - Comprehensive Review on Climate Control and Cooling Systems](https://www.mdpi.com/2073-4395/12/3/626)

### Advanced Greenhouse Cooling Systems

**Fan Ventilation and Exhaust Systems:**

**Air Exchange Requirements:**

- Minimum 1 air change per minute during peak heat periods
- Standard sizing: 60 air changes per hour for heat elimination
- Maximum practical distance from pad to fan: 60 meters (optimal: ≤45 meters)
- Canadian greenhouse examples: 1170 m³/h (22 m² greenhouse), 3680 m³/h (64 m² greenhouse)

**Exhaust Fan Calculations:**

```text
Required m³/h = Greenhouse Volume (m³) × Air Changes per Hour
Heat Removal Rate = m³/h × 0.33 × Temperature Difference (°C)
Where 0.33 = constant for air density and specific heat (Watts per m³/h·°C)
```

**Evaporative Cooling Systems:**

**Fan and Pad Systems:**

- Can cool greenhouse 5-11°C below outside temperature
- Most effective in dry climates, beneficial anywhere
- Exhaust fans on one wall, evaporative pads on opposite wall
- Water distribution: sump pump circulates water through pad system

**High-Pressure Fogging/Misting:**

- Operating pressure: 48-69+ bar (4.8-6.9+ MPa)
- Droplet size: 10 microns (flash evaporation)
- Cooling effect: even temperature distribution through air circulation
- Cost-effective alternative to mechanical air conditioning

**Wet Wall Cooling:**

- More effective but higher cost than fogging systems
- Water flows down through pad material
- Air pulled through wet pads by exhaust fans
- Continuous water circulation maintains cooling efficiency

**Geothermal and Ground Cooling:**

**Ground Tube Systems:**

- Tubes buried 2.4 meters deep, extending 90+ meters
- Summer: hot air exhausted through ground tubes for cooling
- Winter: system reversed to extract ground heat
- Stable ground temperatures (8-12°C at 2m depth in Canada)

**Low-Grade Geothermal:**

- Can reduce energy bills by 50-60%
- Low maintenance requirements
- Cost-effective for larger operations
- Ground loop heat exchangers for year-round climate control

**Natural and Passive Cooling:**

**Night Ventilation:**

- Thermal gradient ventilation using side and ridge vents
- Cool night air enters through side vents
- Warm air exits through ridge vents via natural convection
- Most cost-effective cooling method

**Thermal Curtains/Energy Screens:**

- Reduce heat loss by up to 65% through roof glazing
- Save 35% of heat loss through side walls
- Overall energy bill reduction: 20-40%
- Semi-porous materials: alternating aluminized and clear strips

**Radiant Cooling Methods:**

- Reflective barriers to reduce radiant heat absorption
- Strategic placement of thermal mass for night cooling
- Roof ventilation for radiant heat removal
- Heat transfer calculation: Q = A × (Ti - To) / R

**Canadian Climate Adaptations:**

**Regional Cooling Requirements:**

- Extremely long summer days in northern regions require intensive cooling
- Southern Ontario (Leamington): 25% more cooling than BC Delta region
- Montreal region: 25% more cooling than Leamington
- Integration with existing heating systems for year-round efficiency

**Advanced Control Systems:**

- Smart thermostats for sequential fan operation
- Automated pad and fogging system integration
- Humidity control coordination with cooling systems
- Energy management to minimize operational costs

**Cooling System Combinations:**

**Integrated Approach:**

1. **Natural ventilation** as primary cooling method
2. **Shading systems** to reduce solar heat gain
3. **Evaporative cooling** for supplemental temperature control
4. **Thermal curtains** for energy efficiency
5. **Automated controls** for optimal system coordination

**Cost-Benefit Analysis:**

- Evaporative cooling: Low capital cost, moderate operating cost
- Geothermal systems: High capital cost, low operating cost
- Natural ventilation: Lowest cost, climate-dependent effectiveness
- Fogging systems: Moderate cost, high effectiveness in all climates

**Performance Calculations:**

```text
Cooling Load (Watts) = Sensible Heat Load + Latent Heat Load
Sensible Heat = 0.33 × m³/h × Temperature Difference (°C)
Latent Heat = 0.84 × m³/h × Humidity Difference (g/kg)
```

**References:**

- [University of Florida - Fan and Pad Evaporative Cooling](https://edis.ifas.ufl.edu/publication/AE069)
- [UMass Extension - Fan and Pad Systems](https://ag.umass.edu/greenhouse-floriculture/fact-sheets/fan-pad-evaporative-cooling-systems)
- [Greenhouse Canada - Energy Curtains](https://www.greenhousecanada.com/energy-curtains-for-vegetables-2791/)
- [BC Greenhouses - Heating and Cooling](https://www.bcgreenhouses.com/greenhouse-buying-tips/cooling-and-heating/)
