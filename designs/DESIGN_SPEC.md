# Honk for Greenhouse - Design Specification

## Overview

A greenhouse design sharing platform where owners, designers, builders, and manufacturers collaborate on greenhouse projects. Public showcase requires approval from all participants.

## User Roles

| Role | Capabilities |
|------|-------------|
| Visitor | View Showcase only (no auth) |
| Owner | Full control of farms/projects, manages permissions |
| Co-Owner | Shared ownership of farm, can manage projects |
| Designer | Create/edit designs, submit for quotes |
| Builder | View shared projects, submit quotes with PDF |
| Manufacturer | View specs, provide equipment quotes |
| Administrator | Disable showcase visibility, correct errors |

## Core Design Principles

### Live-Update Pattern

All calculations and data updates use a live-update pattern instead of loading spinners:

1. **User input triggers debounced server request** (300ms delay for typing, immediate for selection changes)
2. **Server calculates and returns updated values**
3. **UI updates in place without full page reload**
4. **Subtle transition animations indicate updates**

```
User Input -> Debounce -> Server Calculation -> Partial HTML Response -> DOM Update
```

This pattern applies to:
- R-value calculations when selecting materials/thickness
- BTU/heat loss calculations when changing dimensions or insulation
- Heating source percentage balancing
- Climate data chart updates
- Cost projections

### Progressive Enhancement

All core functionality works without JavaScript:
- Forms submit normally with page reload
- HTMX enhances with partial updates
- No-JS users get full functionality, just with more page loads

### Accessibility First

- All interactive elements have ARIA labels
- Focus management after dynamic updates
- Skip links for keyboard navigation
- 44px minimum touch targets on mobile
- Color is never the only indicator of state

### Internationalization

- All user-facing strings use translation keys
- Initial support: English (en), French (fr-CA)
- Designed for extension to Spanish, Korean, Chinese
- See I18N_SPEC.md for implementation details

## Pages

### 1. Landing/Showcase (Public)

- Header: "Honk for Greenhouse" + Language Switcher + Login/Sign In
- Skip link: "Skip to main content"
- Tabs: `SHOWCASE` | `SHARED` | `PERSONAL`
- Grid of greenhouse cards (public designs)
- Responsive: 1 column mobile, 2 tablet, 3 desktop
- No auth required

### 2. Shared

- Greenhouses shared with current user
- Auth required
- Empty state with guidance if no shares

### 3. Personal

- User's own farms and projects
- "Add Farm" button prominently displayed
- Auth required

### 4. Farm Detail

- Farm header with location, action menu (Share/Transfer/Delete)
- Climate/Soil section with charts
- **"Add Project" button** (44px height minimum)
- List of projects (collapsible cards)
- Each project card has its own action menu

### 5. Project Detail (Expanded Card)

- Thumbnail + Specs (Thermal Mass, Insulation, Equipment)
- Production period, Expected Production/Revenue/OpCost
- Monthly Cost chart (responsive sizing)
- Heat Need / Cooling Need charts
- Builder quotes table (card view on mobile)
- Publication approval status

### 6. Login

- Email input with validation
- OAuth: GitHub, Facebook, Google
- Language persists through auth flow

### 7. Add Farm Form

- Name input
- Location: Country -> Province/State -> City (cascading dropdowns)
- **ZIP/Postal Code**: Used for climate zone lookup
- **Climate Zone Display** (auto-populated from location):
  - Canadian zone (for Canadian locations)
  - USDA zone (for US locations, or as reference for Canada)
  - Design minimum temperature
  - Annual Heating Degree Days (HDD)
  - Frost depth for foundation design
- **Co-Owners**: Add other owners by email
- **Energy Costs**: Configurable list
  - Electricity rate ($/kWh)
  - Wood cost ($/cord)
  - Custom fuel types with +Add
- All inputs have labels and help text

See [CLIMATE_ZONES.md](CLIMATE_ZONES.md) for zone lookup implementation details.

## Project Creation Wizard (9 Steps)

The wizard is designed to reduce cognitive load by:
- Breaking complex configuration into focused steps
- Guiding material selection instead of asking for raw values
- Providing sensible defaults with explanations
- Showing live calculations as values change

### Step 1 - Select Shape

- Grid of greenhouse shape options (various roof profiles)
- Each shape is a radio button styled as a card (works without JS)
- Visual icons showing shape profiles
- "Request Additional Shape" option for custom requests
- Shape selection affects calculations in later steps

### Step 2 - Dimensions & Crop Profile

- Length / Width / Height inputs (meters, live validation)
- **Orientation compass** with proper ARIA labels
  - 8 cardinal/intercardinal directions as radio buttons
  - Visual compass diagram
  - Tooltip explaining orientation importance
- Goal: Crop type dropdown (Tomato, Lettuce, Pepper, etc.)
- **Crop Profile Display** (fetched via HTMX on crop change):
  - Season range (start -> end month)
  - Temperature requirements (Min/Max C)
  - Light requirements visualization
  - Water requirements visualization

### Step 3 - Glazing Selection

- Visual cards showing glazing options with light transmission %
- Options include:
  - Single/Double Polyethylene
  - Twin/Triple/Five-wall Polycarbonate
  - Single/Double Glass
- **Live heat loss calculation** updates as selection changes
- Displayed: "Estimated Annual Heat Loss: X kWh"
- Tooltip explaining glazing trade-offs (light vs insulation)

### Step 4 - Wall Insulation

**Instead of asking for R-values directly:**

For each wall (North, South, East, West):
1. **Select Material** dropdown:
   - Closed-Cell Spray Foam
   - Open-Cell Spray Foam
   - Fiberglass Batts
   - Mineral Wool
   - Rigid Foam (XPS, EPS, Polyiso)
2. **Select Thickness** dropdown (populated based on material)
3. **R-Value Display** (auto-calculated and shown)

Live update: As user changes any selection, total heat loss recalculates.

See INSULATION_MATERIALS.md for material specifications.

### Step 5 - Roof & Foundation

**Roof Insulation** (North/South sections):
- Same material/thickness selection as walls
- Default: Glazing on south, insulated on north (for passive solar)

**Foundation Type Selection** (expandable cards):

1. **Frost-Protected Shallow Foundation (FPSF)**
   - Grade beam with horizontal insulation
   - Best for: Heated structures, moderate budgets
   - Uses ground anchors for uplift resistance

2. **Helical Pile Foundation**
   - Steel piles screwed below frost line
   - Best for: Difficult soils, quick installation
   - Minimal site disturbance

3. **Concrete Block Wall (CMU)**
   - Full perimeter masonry foundation
   - Best for: Maximum durability, basement option
   - Highest cost and complexity

4. **Post Frame**
   - Embedded or pier-mounted posts
   - Best for: Budget builds, DIY-friendly
   - Common for hoop houses

5. **Floating/Surface**
   - Gravel pad with anchors
   - Best for: Temporary structures, poor soils
   - Lowest cost, seasonal use

**Anchor System Configuration**:
- Anchor type selection based on foundation
- Quantity calculator based on wind loads
- Spacing recommendations
- Uplift capacity verification

**Frost Skirt Configuration** (for FPSF):
- Depth (cm) - with recommended range tooltip
- Width (m) - horizontal extent from foundation
- Material selection (typically XPS for below-grade)
- Explanation: "Insulation around the perimeter prevents ground freezing"

See [FOUNDATIONS_ANCHORS.md](FOUNDATIONS_ANCHORS.md) for detailed specifications and Canadian frost depth requirements.

### Step 6 - Air Sealing

**Instead of asking for ACH directly:**

Radio button selection with clear descriptions:

1. **Tight (0.5 ACH)**
   - "Professional air sealing, verified with blower door test"
   - Best for: High-performance builds with professional installation

2. **Average (1.5 ACH)**
   - "Careful construction with sealed joints and penetrations"
   - Best for: Quality DIY or professional builds

3. **Standard (3 ACH)**
   - "Typical greenhouse construction without special air sealing"
   - Best for: Simple structures, can be improved later

**Blower Door Info Box**:
> "A blower door test after construction can precisely measure your air changes per hour (ACH). For now, select an estimate based on your planned construction quality. You can update this value later after construction is complete."

### Step 7 - Climate Battery System

**Pre-designed climate battery options** instead of custom configuration:

1. **No Climate Battery**
   - "Skip this optimization for simpler construction"
   - Suitable for: Mild climates, budget builds

2. **Multi-Pipe with Fan** (Recommended for small-medium)
   - Visual diagram showing pipe layout
   - "Multiple short pipes with optimized airflow"
   - Auto-sized based on greenhouse dimensions

3. **Two-Trench Design**
   - Visual diagram
   - "Parallel trenches for longer greenhouses"
   - Best for 15-30m length

4. **Single Manifold (Both Sides)**
   - Visual diagram
   - "Central manifold with balanced distribution"
   - Works well with center aisle layouts

5. **Custom / Advanced**
   - Opens detailed configuration panel
   - For experienced builders

See CLIMATE_BATTERY_DESIGNS.md for detailed specifications.

### Step 8 - Thermal Mass & Optimizations

#### Thermal Mass Configuration

Thermal mass stores excess daytime heat for release at night. The wizard guides users through selecting appropriate thermal mass for their greenhouse.

**Thermal Mass Type Selection** (expandable cards):

1. **Water Containers** (Most common)
   - Container type: 55-gallon drums, IBC totes, water tubes
   - Quantity calculator based on glazing area
   - Placement: North wall, perimeter, floor level
   - Estimated heat storage capacity shown

2. **Rock/Gravel Bed**
   - Integrates with climate battery system
   - Rock type and sizing guidance
   - Bed depth configuration

3. **Phase Change Materials (PCM)**
   - Product selection (BioPCM, Infinite-R, etc.)
   - Melting point selection based on climate
   - Area coverage calculator

4. **Concrete/Masonry**
   - Floor slab configuration
   - North wall thermal mass
   - Trombe wall option

**Live display:** "Total Heat Storage Capacity: X kWh (Y% of recommended)"

See [THERMAL_MASS.md](THERMAL_MASS.md) for detailed specifications and sizing calculations.

#### Additional Optimizations

Checkbox list with descriptions and "Configure" buttons:

- **Thermal Night Blanket**
  - "Insulating cover deployed at night to reduce heat loss"
  - Configure: Material type, coverage area
  - Estimated reduction: 20-35%

- **Row Covers**
  - "Floating fabric covers over plants for frost protection"
  - Configure: Coverage percentage, material weight
  - Estimated reduction: 10-20%

Each optimization shows estimated % reduction in heating needs.

### Step 9 - Heating Sources

- **Peak Heat Need** (kW) - calculated and displayed
- **Annual Energy Demand** (kWh) - calculated and displayed

**Heating source configuration:**

Primary source:
- Dropdown: Electric, Natural Gas, Propane, Wood, etc.
- Percentage slider (0-100%)

Secondary source:
- Dropdown (same options)
- Percentage slider (auto-balances with primary)

**Live cost projection** based on farm's energy costs:
- "Estimated annual heating cost: $X,XXX"

### Wizard Navigation

- Step indicator showing progress (1 of 9)
- Back/Next buttons (44px height minimum)
- State preserved when navigating back
- URL updates with `hx-push-url` for browser history
- Warning before leaving with unsaved changes

## Key Features

### Farm Management

- **Co-Ownership**: Multiple owners per farm (invite by email)
- **Energy Costs**: Configurable energy sources and rates
  - Electricity ($/kWh)
  - Wood ($/cord)
  - Custom fuels
- Used for cost calculations in projects

### Climate Data

- Source: Scraped from Environment Canada -> stored in DB
- Granularity: Monthly (with min/max extremes)
- User override: Inline spreadsheet editing
- Lock: Locked by default, explicit unlock to edit
- Versioning: Daily granularity (YYYY-MM-DD)
- Views: Chart <-> Spreadsheet toggle

### Sharing

- View-only (with feedback) or Edit access
- Edit access marks existing quotes as "needs review"
- Options: allow reshare, allow copy/fork
- Owner controls permissions

### Publication to Showcase

- Any participant can request
- Simultaneous async request to ALL participants
- Dashboard shows: Approved | Pending | Denied (with icons, not just colors)
- All must approve for publication
- Can retrigger cycle after denial
- Photos required for publication

### Builder Quotes

- User-entered with PDF attachment
- Stored in S3
- Marked "needs review" after project edits
- Mobile: Card view instead of table

### Notifications

- Email with deep links
- Types: share invite, publication request, quote submitted, quote needs review

## Templ UI Components

| UI Element | Component |
|------------|-----------|
| Navigation | Tabs |
| Greenhouse list | Card grid |
| Metrics | Badge |
| Climate data | Chart (Line/Area) + Table |
| Actions menu | Shared ActionMenu component |
| Projects | Accordion/Collapsible |
| Login | Form + Input + Button |
| Modals | Dialog |
| Status | Alert, Badge (with icons) |
| User | Avatar |
| Help | Tooltip |
| Feedback | Toast (global) |

## Shared Components

### ActionMenu

Reusable action menu for farms and projects:
- Share action
- Transfer action
- Delete action (with confirmation dialog)
- Consistent across all contexts
- See COMPONENTS.md for implementation

### Toast Notifications

Global toast system for:
- Success messages
- Error messages (including HTMX errors)
- Warning messages
- Auto-dismiss with manual close option

### Skip Links

- "Skip to main content" link
- Visible on focus
- Present on all pages

## Storage (S3)

```text
/farms/{farm_id}/photos/
/projects/{project_id}/photos/    # Required for publication
/projects/{project_id}/quotes/    # PDF attachments
/climate/{farm_id}/{version}.json # Versioned climate data
```

## Error Handling

### HTMX Error Handling

Global error handler catches all HTMX response errors:
- Network failures -> Toast with retry suggestion
- Server errors (5xx) -> Toast with generic error message
- Validation errors (4xx) -> Inline error display

### Form Validation

- Client-side validation with HTML5 attributes
- Server-side validation with error messages
- Errors displayed inline with `aria-describedby`
- Focus moved to first error field

## Testing

E2E tests using `github.com/mountain-reverie/playwright-ci-go`:
- All wizard steps
- Mobile responsive behavior
- Keyboard navigation
- Screen reader compatibility
- Language switching

See E2E_TESTS.md for test specifications.

## Mobile Responsive Design

Using Tailwind breakpoints:

- `grid-cols-1 md:grid-cols-2 lg:grid-cols-3` for card grids
- Mobile hamburger menu
- Card-based layouts for tables on mobile
- 44px minimum touch targets
- Responsive chart sizing with aspect ratios
- Stack form layouts on mobile

## Performance Considerations

- HTMX partial updates minimize data transfer
- Climate data cached in browser
- Material data preloaded at startup
- Images lazy-loaded
- Charts render with skeleton states
