# Implementation Plan

## Development Principles

### Accessibility First
Every phase includes accessibility requirements. Components must meet WCAG 2.1 AA standards:
- Focus management after HTMX swaps
- ARIA labels on all interactive elements
- 44px minimum touch targets on mobile
- Color is never the sole indicator of state

### Mobile First
Build mobile layouts first, then enhance for larger screens:
- Responsive breakpoints: mobile (<768px), tablet (768-1024px), desktop (>1024px)
- Card-based views for tables on mobile
- Hamburger menu for mobile navigation

### Metric Units
All calculations and displays use metric units (Celsius, Watts/kWh, meters, m2-K/W).

### Live-Update Pattern
Replace loading spinners with immediate visual feedback:
- Debounced requests (300ms for text input, immediate for selections)
- Optimistic UI updates where appropriate
- Toast notifications for errors

---

## Phase 1: Foundation

### Core Infrastructure
1. **i18n Setup** (see I18N_SPEC.md)
   - Translation middleware for Gin
   - Language detection (Accept-Language, URL param, cookie)
   - Initial translations: English (en), French (fr-CA)
   - Templ `T()` function for string translation

2. **Base Layout** (`layout/base.templ`, `header.templ`, `nav_tabs.templ`)
   - Skip link to main content
   - Global toast container with `aria-live="polite"`
   - HTMX error handler (toast on responseError/sendError)
   - `<div id="dialog-container">` for modals

3. **Shared Components** (`shared/`)
   - `ActionMenu` - Reusable dropdown for Share/Transfer/Delete
   - `Toast` - Notification system
   - `SkipLink` - Accessibility skip navigation
   - `TooltipHelp` - Help icon with tooltip
   - `ResponsiveTable` - Table with mobile card view

4. **Mobile Navigation** (`mobile_menu.templ`)
   - Hamburger menu button (44px touch target)
   - Slide-in drawer
   - Language switcher

### Accessibility Checklist
- [ ] Skip link visible on focus
- [ ] Focus trap in mobile menu drawer
- [ ] ARIA labels on all buttons
- [ ] Keyboard navigation (Tab, Enter, Escape)

---

## Phase 2: Showcase & Cards

1. **Greenhouse Card** (`greenhouse/card.templ`)
   - Thumbnail with alt text
   - Title, location, metrics
   - Badge for publication status

2. **Showcase Page** (public, no auth)
   - Responsive grid: `grid-cols-1 md:grid-cols-2 lg:grid-cols-3`
   - Empty state component
   - Lazy-loaded images

3. **Card List** (`greenhouse/card_list.templ`)
   - HTMX partial for tab switching
   - Focus management after swap

### Accessibility Checklist
- [ ] Image alt text
- [ ] Card is keyboard focusable
- [ ] Badge has icon + text (not color alone)

---

## Phase 3: Charts & Climate

1. **Climate Chart** (Chart.js via Templ UI)
   - Line chart with temperature range
   - Accessible color palette
   - Screen reader description

2. **Monthly Table** (`charts/monthly_table.templ`)
   - Proper `<th scope="col">` headers
   - Mobile: card view with `<dl>` structure

3. **Chart/Table Toggle** (`charts/chart_toggle.templ`)
   - Radio group with ARIA
   - Live swap target

4. **Climate API**
   - `GET /api/farms/{id}/climate`
   - `GET /api/farms/{id}/climate/chart`
   - `GET /api/farms/{id}/climate/table`

### Accessibility Checklist
- [ ] Chart has text alternative
- [ ] Table headers associated with cells
- [ ] Toggle buttons have `aria-pressed`

---

## Phase 4: Auth & Roles

1. **Login Page** (`auth/login_form.templ`)
   - Email input with validation
   - OAuth buttons (GitHub, Google, Facebook)
   - Error messages with `aria-describedby`

2. **OAuth Integration** (Goth)
   - Session management
   - Locale persistence through auth flow

3. **Role-Based Middleware**
   - Visitor, Owner, Co-Owner, Designer, Builder, Manufacturer, Admin
   - Permission checks on protected routes

4. **User Menu** (`auth/user_menu.templ`)
   - Avatar with dropdown
   - Sign out action

### Accessibility Checklist
- [ ] Form labels associated with inputs
- [ ] Error focus management
- [ ] OAuth buttons have descriptive labels

---

## Phase 5: Farm Management

1. **Add Farm Form** (`farm/form.templ`)
   - Name input with validation
   - Location cascading dropdowns (Country -> Province -> City)
   - Progressive enhancement (works without JS)

2. **Co-Owner Management** (`farm/co_owners.templ`)
   - Add by email (HTMX submit)
   - List with remove buttons
   - Confirmation dialog for removal

3. **Energy Costs** (`farm/energy_costs.templ`)
   - Electricity rate ($/kWh)
   - Natural gas rate ($/m3)
   - Wood cost ($/cord)
   - Custom fuel types with +Add button
   - Provincial defaults (see SERVICES.md)

4. **Farm Detail Page**
   - Header with ActionMenu
   - Climate section
   - "Add Project" button (prominent, 44px height)
   - Project list

### Accessibility Checklist
- [ ] Form validation announced
- [ ] Cascading dropdowns maintain focus
- [ ] Confirm dialogs trap focus

---

## Phase 6: Project Creation Wizard (9 Steps)

### Wizard Shell
- Step indicator showing progress (1-9)
- Back/Next buttons (44px height)
- URL updates with `hx-push-url`
- Warning before leaving with unsaved changes
- Focus moves to step heading after navigation

### Step 1: Shape Selection (`wizard/step1_shape.templ`)
- Grid of shape cards with visual icons
- Radio buttons styled as cards (works without JS)
- "Request Additional Shape" link

### Step 2: Dimensions & Crop Profile (`wizard/step2_dimensions.templ`)
- Length / Width / Height inputs (meters)
- **Orientation Compass** with proper ARIA
  - 8 direction radio buttons
  - Visual compass diagram
  - Tooltip explaining orientation importance
- Crop type dropdown
- **Crop Profile Display** (HTMX fetch on crop change)
  - Season range
  - Temperature requirements (Celsius)
  - Light/Water requirements charts

### Step 3: Glazing Selection (`wizard/step3_glazing.templ`)
- Visual cards showing glazing options
- Light transmission percentage displayed
- **Live heat loss calculation** updates as selection changes
- Options: Single/Double Poly, Twin/Triple/Five-wall PC, Single/Double Glass

### Step 4: Wall Insulation (`wizard/step4_walls.templ`)
- For each wall (North, South, East, West):
  - **Material Selector** dropdown (see INSULATION_MATERIALS.md)
  - **Thickness Selector** (populated based on material)
  - **R-Value Display** (auto-calculated, m2-K/W)
- Live heat loss update on any change

### Step 5: Roof & Foundation (`wizard/step5_roof.templ`)
- Roof insulation (North/South sections)
- Default: Glazing on south, insulated on north

**Foundation Type Selection** (`foundation_type_selector.templ`):
- **FPSF Grade Beam** - Horizontal insulation frost protection
- **Helical Pile** - Below frost line, minimal excavation
- **CMU Block Wall** - Full perimeter, maximum durability
- **Post Frame** - Embedded posts, budget-friendly
- **Floating/Surface** - Temporary structures
- Visual diagrams for each option
- Regional frost depth auto-populated from farm location

**Anchor System Configuration** (`anchor_configurator.templ`):
- Anchor type dropdown (auger, ground screw, concrete)
- **Quantity Calculator** based on:
  - Greenhouse dimensions
  - Regional design wind speed
  - Foundation type requirements
- Live uplift capacity verification display
- Spacing diagram showing anchor positions

**Frost Skirt Configuration** (for FPSF):
- Depth (cm) with recommended range tooltip
- Width (m)
- Material selection

See FOUNDATIONS_ANCHORS.md for detailed specifications.

### Step 6: Air Sealing (`wizard/step6_air_seal.templ`)
- Radio button selection with clear descriptions:
  - **Tight (0.5 ACH)** - "Professional air sealing, verified with blower door test"
  - **Average (1.5 ACH)** - "Careful construction with sealed joints"
  - **Standard (3.0 ACH)** - "Typical greenhouse construction"
- Info box explaining blower door testing

### Step 7: Climate Battery System (`wizard/step7_climate_battery.templ`)
- Pre-designed options (see CLIMATE_BATTERY_DESIGNS.md):
  - No Climate Battery
  - Multi-Pipe with Fan (recommended for small-medium)
  - Two-Trench Design
  - Single Manifold (Both Sides)
  - Custom / Advanced
- Visual diagrams for each option
- Specifications table

### Step 8: Thermal Mass & Optimizations (`wizard/step8_thermal_mass.templ`)

**Thermal Mass Configuration:**
- **Type Selection** cards:
  - Water containers (drums, IBC totes, tubes)
  - Rock/gravel bed (integrates with climate battery)
  - Phase change materials (PCM)
  - Concrete/masonry (floor, walls, Trombe)
- **Sizing Calculator** (`thermal_mass_calculator.templ`)
  - Live calculation based on glazing area and climate zone
  - "Recommended: X liters water or equivalent"
  - "Current capacity: Y kWh (Z% of recommended)"
- **Container Selector** (`water_container_selector.templ`)
  - Type: 55-gal drums, 275-gal IBC, tubes
  - Quantity input with +/- buttons
  - Placement: North wall, perimeter, floor level
- **PCM Selector** (`pcm_selector.templ`)
  - Product: BioPCM, Infinite-R, custom
  - Melting point selection
  - Area coverage

See THERMAL_MASS.md for detailed specifications.

**Additional Optimizations** (checkboxes):
- Thermal Night Blanket (20-35% reduction)
  - Configure: Material type, coverage area
- Row Covers (10-20% reduction)
  - Configure: Coverage percentage, material weight
- Each shows estimated % reduction in heating needs

### Step 9: Heating Sources (`wizard/step9_heating.templ`)
- **Peak Heat Need** display (Watts)
- **Annual Energy Demand** display (kWh)
- Primary source dropdown + percentage slider
- Secondary source dropdown + percentage slider (auto-balances)
- **Live cost projection** based on farm's energy costs
- Displays: "Estimated annual heating cost: $X,XXX CAD"

### Wizard Completion
- Summary review
- Create project button
- Redirect to project detail

### Accessibility Checklist
- [ ] Step indicator has `aria-current="step"`
- [ ] Compass has `role="radiogroup"` with `aria-label`
- [ ] Sliders have `aria-valuemin`, `aria-valuemax`, `aria-valuenow`
- [ ] Material selectors update thickness options accessibly
- [ ] Live calculations announced via `aria-live`

---

## Phase 7: Project Detail & CRUD

1. **Project Card** (collapsible)
   - Accordion pattern with ARIA
   - Thumbnail + summary on collapsed
   - Full specs on expanded

2. **Specs Display**
   - Thermal Mass, Insulation, Equipment
   - Production period, Expected Production/Revenue/OpCost

3. **Monthly Cost Chart**
   - Responsive sizing with aspect ratio
   - Accessible alternative text

4. **Heat/Cooling Needs Charts**
   - Monthly breakdown
   - Unit: kWh

5. **Project ActionMenu**
   - Share, Transfer, Delete
   - Delete confirmation dialog

### Accessibility Checklist
- [ ] Accordion uses `aria-expanded`
- [ ] Charts have text descriptions
- [ ] Delete confirmation traps focus

---

## Phase 8: Climate Editing

1. **Lock Toggle** (`charts/lock_toggle.templ`)
   - "Locked" by default
   - Click to unlock editing
   - ARIA states

2. **Inline Cell Editing** (`charts/editable_table.templ`)
   - Click to edit
   - Enter to save, Escape to cancel
   - Tab to next cell

3. **Version Selector** (`charts/version_selector.templ`)
   - Dropdown of date-versioned data
   - Load selected version

4. **Save/Cancel**
   - Creates new version (YYYY-MM-DD format)
   - Backs up to S3

### Accessibility Checklist
- [ ] Editable cells announced
- [ ] Focus management during editing
- [ ] Version changes announced

---

## Phase 9: Quotes

1. **Quotes Table** (`quotes/quotes_table.templ`)
   - Desktop: full table
   - Mobile: card view (`quotes/quote_card.templ`)

2. **Quote Form** (`quotes/quote_form.templ`)
   - Builder name, price, notes
   - PDF upload with drag-drop
   - File type validation

3. **S3 Integration**
   - Upload to `projects/{id}/quotes/{quote_id}.pdf`
   - Secure signed URLs

4. **"Needs Review" Badge**
   - Shown when project edited after quote submission
   - Icon + text (not color alone)

### Accessibility Checklist
- [ ] File input has proper label
- [ ] Upload progress announced
- [ ] Badge has title/aria-label

---

## Phase 10: Sharing

1. **Share Dialog** (`sharing/share_dialog.templ`)
   - Email input
   - Permission selection (View only, View + Feedback, Edit)
   - Options: Allow reshare, Allow copy/fork

2. **Share Management UI**
   - List of current shares
   - Remove share button

3. **Permission Checks**
   - Middleware validates access
   - Edit access marks quotes as "needs review"

4. **Email Notifications**
   - Share invite email with deep link
   - Uses i18n for recipient's locale

### Accessibility Checklist
- [ ] Dialog traps focus
- [ ] Radio group for permissions
- [ ] Close button accessible

---

## Phase 11: Publication

1. **Photo Requirements** (`publication/photo_requirements.templ`)
   - Checklist of required photos
   - Upload interface

2. **Request Button** (`publication/request_button.templ`)
   - Disabled until photos complete
   - Explains requirements

3. **Approval Status** (`publication/approval_status.templ`)
   - Dashboard showing all participants
   - Icons: Approved, Pending, Denied (with labels)

4. **Approve/Deny Flow**
   - Email with approve/deny links
   - Token-based authorization
   - Denial requires reason

5. **Retrigger Logic**
   - Can restart after denial
   - New tokens generated

### Accessibility Checklist
- [ ] Status icons have text labels
- [ ] Progress communicated
- [ ] Email links work without JS

---

## Phase 12: Admin

1. **Admin Dashboard**
   - List of published projects
   - Search/filter

2. **Visibility Toggle**
   - Can remove from showcase
   - Requires reason

3. **Error Correction Tools**
   - Edit project data
   - Audit logging

---

## Phase 13: Testing & QA

1. **E2E Tests** with playwright-ci-go
   - All wizard steps
   - Mobile responsive behavior
   - Keyboard navigation
   - Language switching

2. **Accessibility Testing**
   - Automated axe-core scans
   - Manual screen reader testing
   - Keyboard-only testing

3. **Test Fixtures**
   - Seed data for farms, projects, users
   - Mock climate data

4. **CI Pipeline**
   - Run tests on PR
   - Accessibility gate
   - Performance benchmarks

See E2E_TESTS.md for detailed test specifications.

---

## Deployment Phases

### Alpha (Internal)
- Phases 1-6 complete
- Basic wizard flow
- No sharing/publication

### Beta (Invited Users)
- Phases 7-10 complete
- Full feature set
- Limited publication

### Production
- Phase 11-13 complete
- Admin tools
- Full testing

---

## Dependencies

### External Services
- **Turso**: SQLite-compatible database
- **AWS S3**: Photo and PDF storage
- **AWS SES**: Email notifications
- **OAuth Providers**: GitHub, Google, Facebook

### Templ UI Components
```bash
templui add tabs card badge button dropdown accordion dialog form input avatar alert table chart checkbox radio tooltip spinner icon emptystate toast progress slider
```

### i18n Translation Files
- `locales/en.yaml`
- `locales/fr-CA.yaml`

---

## Risk Mitigation

| Risk | Mitigation |
|------|------------|
| Complex wizard flow | Break into 9 focused steps |
| Accessibility gaps | Checklist per phase, automated testing |
| i18n text expansion | Design flexible layouts, test with FR |
| Mobile usability | Mobile-first development, 44px targets |
| Live calculation performance | Debounced requests, optimistic UI |
| Data loss during wizard | Auto-save to session, warn on leave |
