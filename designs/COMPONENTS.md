# Component Structure

## Directory Layout

```text
internal/components/
├── shared/
│   ├── action_menu.templ        # Reusable action menu (Share/Transfer/Delete)
│   ├── toast.templ              # Toast notification system
│   ├── skip_link.templ          # Skip to content link
│   ├── focus_trap.templ         # Focus management helpers
│   ├── tooltip_help.templ       # Help tooltip with icon
│   ├── empty_state.templ        # Empty state component
│   └── responsive_table.templ   # Table with mobile card view
├── layout/
│   ├── base.templ               # HTML document wrapper
│   ├── header.templ             # Site header with nav
│   ├── nav_tabs.templ           # Main navigation tabs
│   ├── mobile_menu.templ        # Mobile hamburger menu
│   └── language_switcher.templ  # EN/FR toggle
├── greenhouse/
│   ├── card.templ
│   ├── card_list.templ
│   └── detail.templ
├── farm/
│   ├── header.templ
│   ├── form.templ
│   ├── climate_section.templ
│   ├── project_list.templ
│   ├── add_project_button.templ
│   ├── co_owners.templ
│   └── energy_costs.templ
├── project/
│   ├── card.templ
│   ├── specs.templ
│   ├── financials.templ
│   ├── needs_charts.templ
│   └── quotes_table.templ       # Includes mobile card view
├── wizard/
│   ├── layout.templ             # Wizard shell with step indicator
│   ├── step_indicator.templ     # Progress indicator (1-9)
│   ├── step1_shape.templ
│   ├── step2_dimensions.templ
│   ├── step3_glazing.templ      # NEW: Separated glazing step
│   ├── step4_walls.templ        # NEW: Wall insulation with material selection
│   ├── step5_roof.templ         # NEW: Roof, foundation, and frost skirt
│   ├── foundation_type_selector.templ   # NEW: Foundation type selection cards
│   ├── anchor_configurator.templ        # NEW: Anchor system configuration
│   ├── wind_load_calculator.templ       # NEW: Wind uplift calculation display
│   ├── frost_depth_display.templ        # NEW: Regional frost depth info
│   ├── step6_air_seal.templ     # NEW: Air sealing with descriptions
│   ├── step7_climate_battery.templ  # NEW: Climate battery system selection
│   ├── step8_thermal_mass.templ     # NEW: Thermal mass & optimizations
│   ├── step9_heating.templ
│   ├── shape_card.templ
│   ├── orientation_compass.templ # Accessible compass with ARIA
│   ├── crop_profile.templ
│   ├── material_selector.templ  # NEW: Material + thickness selector
│   ├── climate_battery_design_card.templ  # NEW: Climate battery option card with diagram
│   ├── thermal_mass_type_selector.templ   # NEW: Water/rock/PCM/concrete selection
│   ├── water_container_selector.templ     # NEW: Container type and quantity
│   ├── pcm_selector.templ                 # NEW: PCM product selection
│   ├── thermal_mass_calculator.templ      # NEW: Interactive sizing calculator
│   ├── thermal_mass_summary.templ         # NEW: Capacity summary display
│   ├── optimization_item.templ
│   └── heating_sources.templ
├── charts/
│   ├── climate_chart.templ
│   ├── cost_chart.templ
│   ├── monthly_table.templ
│   ├── editable_table.templ
│   ├── lock_toggle.templ
│   ├── version_selector.templ
│   ├── chart_toggle.templ
│   ├── light_requirements.templ
│   └── water_requirements.templ
├── sharing/
│   ├── share_dialog.templ
│   ├── permissions.templ
│   └── showcase_approval.templ
├── publication/
│   ├── request_button.templ
│   ├── approval_status.templ
│   ├── approval_card.templ
│   └── photo_requirements.templ
├── quotes/
│   ├── quote_row.templ
│   ├── quote_card.templ         # NEW: Mobile card view
│   ├── quote_form.templ
│   └── needs_review_badge.templ
└── auth/
    ├── login_form.templ
    ├── oauth_buttons.templ
    └── user_menu.templ
```

## Templ UI Dependencies

Install from <https://templui.io/docs/components>:

```bash
templui add tabs card badge button dropdown accordion dialog form input avatar alert table chart checkbox radio tooltip spinner icon emptystate toast
```

## Shared Components

### ActionMenu

Reusable action menu component for farms and projects.

```go
package shared

import "context"

type ActionMenuItem struct {
    LabelKey     string   // i18n key
    Icon         string   // Icon name
    HxGet        string   // For dialog triggers
    HxDelete     string   // For delete actions
    HxTarget     string
    ConfirmKey   string   // i18n key for confirmation message
    Destructive  bool     // Red styling
}

templ ActionMenu(ctx context.Context, items []ActionMenuItem) {
    @dropdown.Dropdown() {
        @dropdown.Trigger() {
            @button.Button(
                variant="ghost",
                size="icon",
                aria-label=T(ctx, "actions.menu"),
                class="min-w-[44px] min-h-[44px]",
            ) {
                @icon.Icon(name="more-vertical", size="sm")
            }
        }
        @dropdown.Content(align="end") {
            for _, item := range items {
                if item.HxDelete != "" {
                    @dropdown.Item(
                        hx-delete=item.HxDelete,
                        hx-target=item.HxTarget,
                        hx-confirm=T(ctx, item.ConfirmKey),
                        class=if item.Destructive { "text-destructive" } else { "" },
                    ) {
                        @icon.Icon(name=item.Icon, size="sm", class="mr-2")
                        { T(ctx, item.LabelKey) }
                    }
                } else {
                    @dropdown.Item(
                        hx-get=item.HxGet,
                        hx-target="#dialog-container",
                    ) {
                        @icon.Icon(name=item.Icon, size="sm", class="mr-2")
                        { T(ctx, item.LabelKey) }
                    }
                }
            }
        }
    }
}

// Helper to create farm action menu items
func FarmActionItems(farmID string, canDelete bool) []ActionMenuItem {
    items := []ActionMenuItem{
        {LabelKey: "actions.share", Icon: "share-2", HxGet: "/api/farms/" + farmID + "/share"},
        {LabelKey: "actions.transfer", Icon: "user-plus", HxGet: "/api/farms/" + farmID + "/transfer"},
    }
    if canDelete {
        items = append(items, ActionMenuItem{
            LabelKey:    "actions.delete",
            Icon:        "trash-2",
            HxDelete:    "/farms/" + farmID,
            HxTarget:    "#main-content",
            ConfirmKey:  "actions.confirm_delete_farm",
            Destructive: true,
        })
    }
    return items
}

// Helper to create project action menu items
func ProjectActionItems(projectID string, canDelete bool) []ActionMenuItem {
    items := []ActionMenuItem{
        {LabelKey: "actions.share", Icon: "share-2", HxGet: "/api/projects/" + projectID + "/share"},
        {LabelKey: "actions.transfer", Icon: "user-plus", HxGet: "/api/projects/" + projectID + "/transfer"},
    }
    if canDelete {
        items = append(items, ActionMenuItem{
            LabelKey:    "actions.delete",
            Icon:        "trash-2",
            HxDelete:    "/projects/" + projectID,
            HxTarget:    "#main-content",
            ConfirmKey:  "actions.confirm_delete_project",
            Destructive: true,
        })
    }
    return items
}
```

### Toast Notification System

Global toast system with HTMX error handling.

```go
package shared

templ ToastContainer() {
    <div
        id="toast-container"
        class="fixed bottom-4 right-4 z-50 flex flex-col gap-2 max-w-sm"
        aria-live="polite"
        aria-atomic="true"
    ></div>
}

templ Toast(variant string, message string, dismissable bool) {
    <div
        class={
            "p-4 rounded-lg shadow-lg flex items-start gap-3 animate-slide-in",
            "bg-background border",
            variantClasses(variant),
        }
        role="alert"
    >
        @toastIcon(variant)
        <p class="flex-1 text-sm">{ message }</p>
        if dismissable {
            <button
                type="button"
                class="text-muted-foreground hover:text-foreground"
                onclick="this.parentElement.remove()"
                aria-label="Dismiss"
            >
                @icon.Icon(name="x", size="sm")
            </button>
        }
    </div>
}

templ toastIcon(variant string) {
    switch variant {
    case "success":
        @icon.Icon(name="check-circle", size="sm", class="text-green-600")
    case "error":
        @icon.Icon(name="alert-circle", size="sm", class="text-destructive")
    case "warning":
        @icon.Icon(name="alert-triangle", size="sm", class="text-yellow-600")
    default:
        @icon.Icon(name="info", size="sm", class="text-blue-600")
    }
}

func variantClasses(variant string) string {
    switch variant {
    case "success":
        return "border-green-200 bg-green-50"
    case "error":
        return "border-destructive/20 bg-destructive/5"
    case "warning":
        return "border-yellow-200 bg-yellow-50"
    default:
        return "border-blue-200 bg-blue-50"
    }
}
```

### Skip Link

```go
package shared

templ SkipLink(ctx context.Context) {
    <a
        href="#main-content"
        class="sr-only focus:not-sr-only focus:absolute focus:top-4 focus:left-4 focus:z-50 focus:px-4 focus:py-2 focus:bg-primary focus:text-primary-foreground focus:rounded-md focus:outline-none focus:ring-2 focus:ring-ring"
    >
        { T(ctx, "nav.skip_to_content") }
    </a>
}
```

### TooltipHelp

Reusable help tooltip with consistent styling.

```go
package shared

templ TooltipHelp(ctx context.Context, contentKey string) {
    @tooltip.Tooltip() {
        @tooltip.Trigger() {
            <span class="inline-flex items-center">
                @icon.Icon(name="help-circle", size="sm", class="text-muted-foreground ml-1 cursor-help")
            </span>
        }
        @tooltip.Content(class="max-w-xs") {
            { T(ctx, contentKey) }
        }
    }
}
```

### ResponsiveTable

Table that converts to cards on mobile.

```go
package shared

type TableColumn struct {
    Key       string
    LabelKey  string
    Class     string
    MobileHide bool  // Hide this column on mobile card view
}

type TableRow struct {
    ID     string
    Values map[string]templ.Component
}

templ ResponsiveTable(ctx context.Context, columns []TableColumn, rows []TableRow) {
    // Desktop table view
    <div class="hidden md:block overflow-x-auto">
        @table.Table() {
            @table.Header() {
                @table.Row() {
                    for _, col := range columns {
                        @table.Head(class=col.Class) {
                            { T(ctx, col.LabelKey) }
                        }
                    }
                }
            }
            @table.Body() {
                for _, row := range rows {
                    @table.Row() {
                        for _, col := range columns {
                            @table.Cell(class=col.Class) {
                                @row.Values[col.Key]
                            }
                        }
                    }
                }
            }
        }
    </div>

    // Mobile card view
    <div class="md:hidden space-y-4">
        for _, row := range rows {
            @card.Card() {
                @card.Content(class="pt-4") {
                    <dl class="space-y-2">
                        for _, col := range columns {
                            if !col.MobileHide {
                                <div class="flex justify-between items-center">
                                    <dt class="text-sm text-muted-foreground">
                                        { T(ctx, col.LabelKey) }
                                    </dt>
                                    <dd class="font-medium">
                                        @row.Values[col.Key]
                                    </dd>
                                </div>
                            }
                        }
                    </dl>
                }
            }
        }
    </div>
}
```

## Component Hierarchy

```text
base.templ
├── @SkipLink()
├── header.templ
│   ├── Logo
│   ├── @LanguageSwitcher()
│   ├── auth/user_menu.templ (if logged in)
│   └── Login/SignIn buttons (if not)
├── nav_tabs.templ
│   └── Tabs: Showcase | Shared | Personal
├── <main id="main-content" tabindex="-1">
│   └── [page content]
└── @ToastContainer()
└── <div id="dialog-container"></div>
```

## HTMX Patterns

### Live Update Pattern (Debounced)

For text inputs with calculations:

```go
hx-post="/api/wizard/calculate"
hx-target="#result-display"
hx-trigger="input changed delay:300ms"
hx-swap="innerHTML"
```

### Live Update Pattern (Immediate)

For select dropdowns:

```go
hx-post="/api/wizard/calculate"
hx-target="#result-display"
hx-trigger="change"
hx-swap="innerHTML"
```

### Focus Management After Swap

```go
hx-on::after-swap="document.querySelector('#target h2')?.focus()"
```

### Global Error Handler (in base.templ)

```html
<script>
document.body.addEventListener('htmx:responseError', function(evt) {
    const toast = document.createElement('div');
    toast.innerHTML = `<div class="p-4 rounded-lg shadow-lg bg-destructive/10 border border-destructive/20" role="alert">
        <p class="text-sm text-destructive">An error occurred. Please try again.</p>
    </div>`;
    document.getElementById('toast-container').appendChild(toast.firstChild);
    setTimeout(() => toast.firstChild?.remove(), 5000);
});

document.body.addEventListener('htmx:sendError', function(evt) {
    const toast = document.createElement('div');
    toast.innerHTML = `<div class="p-4 rounded-lg shadow-lg bg-yellow-50 border border-yellow-200" role="alert">
        <p class="text-sm">Network error. Please check your connection.</p>
    </div>`;
    document.getElementById('toast-container').appendChild(toast.firstChild);
    setTimeout(() => toast.firstChild?.remove(), 5000);
});
</script>
```

## HTMX Action Patterns

| Action | Trigger | Target | Swap | Notes |
|--------|---------|--------|------|-------|
| Tab switch | `hx-get` | `#main-content` | innerHTML | Push URL |
| Chart/Table toggle | `hx-get` | `#climate-view-{id}` | innerHTML | |
| Unlock climate edit | `hx-post` | `#climate-table-{id}` | innerHTML | |
| Material selection | `hx-get` | `#thickness-{surface}` | innerHTML | Populate options |
| Calculate R-value | `hx-post delay:100ms` | `#r-value-{surface}` | innerHTML | Live update |
| Calculate heat loss | `hx-post delay:300ms` | `#heat-loss-display` | innerHTML | Debounced |
| Wizard: Next step | `hx-get` | `#wizard-content` | innerHTML | Focus h2, push URL |
| Wizard: Back step | `hx-get` | `#wizard-content` | innerHTML | Focus h2, push URL |
| Select climate battery design | `hx-post` | `#climate-battery-config` | innerHTML | Show config |
| Select foundation type | `hx-post` | `#foundation-config` | innerHTML | Show type-specific options |
| Update anchor config | `hx-post delay:100ms` | `#anchor-summary` | innerHTML | Recalculate capacity |
| Calculate wind load | `hx-post` | `#wind-load-display` | innerHTML | Update uplift requirements |
| Add optimization | `hx-post` | `#optimization-list` | beforeend | |
| Update slider | `hx-post delay:100ms` | `#heating-percentages` | innerHTML | Balance % |
| Share project | `hx-post` | none | none | Close dialog, show toast |
| Delete item | `hx-delete` | `#main-content` | innerHTML | With confirm |

## Accessibility Checklist

### All Interactive Elements

- [ ] `aria-label` or visible label
- [ ] Focus visible styles
- [ ] Minimum 44x44px touch target on mobile
- [ ] Keyboard accessible (Tab, Enter, Space, Escape)

### Dynamic Content

- [ ] `aria-live="polite"` on update regions
- [ ] Focus moved to new content after swap
- [ ] Loading states announced

### Forms

- [ ] Labels associated with inputs
- [ ] Error messages with `aria-describedby`
- [ ] Required fields marked with `aria-required`
- [ ] Validation errors focused

### Navigation

- [ ] Skip link present
- [ ] Current page indicated
- [ ] Disabled items have `aria-disabled`

## Mobile Breakpoints

- **Mobile**: < 768px (md)
  - Single column layouts
  - Hamburger menu
  - Card views for tables
  - Full-width buttons

- **Tablet**: 768px - 1024px
  - Two column grids
  - Side navigation possible
  - Tables with horizontal scroll

- **Desktop**: > 1024px (lg)
  - Three column grids
  - Full navigation
  - Full tables

## Touch Target Sizes

All interactive elements must meet 44x44px minimum:

```go
// Button with minimum size
@button.Button(class="min-w-[44px] min-h-[44px]")

// Icon button
@button.Button(variant="ghost", size="icon", class="w-11 h-11")

// Compass direction buttons
<button class="w-11 h-11 md:w-8 md:h-8 rounded-full ...">
```

## Chart Responsive Sizing

```go
templ ResponsiveChart(data ChartData) {
    <div class="aspect-[4/3] md:aspect-[16/9] min-h-[200px] max-h-[400px]">
        @chart.Chart(
            variant="line",
            data=data,
            responsive=true,
            maintainAspectRatio=false,
        )
    </div>
}
```

## Foundation Components

### Foundation Type Selector

Selection cards for foundation type with visual diagrams.

```go
package wizard

import "context"

type FoundationOption struct {
    Type        string  // "fpsf", "helical_pile", "cmu_wall", "post_frame", "floating"
    LabelKey    string  // i18n key
    DescKey     string  // i18n key for description
    DiagramURL  string  // Path to diagram SVG
    BestFor     string  // i18n key for "best for" text
    CostLevel   string  // "low", "medium", "high"
    Complexity  string  // "low", "medium", "high"
}

templ FoundationTypeSelector(ctx context.Context, selected string, options []FoundationOption) {
    <fieldset
        class="space-y-4"
        role="radiogroup"
        aria-labelledby="foundation-legend"
    >
        <legend id="foundation-legend" class="text-lg font-semibold mb-4">
            { T(ctx, "wizard.foundation.select_type") }
        </legend>

        <div class="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 gap-4">
            for _, opt := range options {
                @FoundationCard(ctx, opt, selected == opt.Type)
            }
        </div>
    </fieldset>
}

templ FoundationCard(ctx context.Context, opt FoundationOption, isSelected bool) {
    <label
        class={
            "relative flex flex-col p-4 rounded-lg border-2 cursor-pointer transition-all",
            "hover:border-primary/50 focus-within:ring-2 focus-within:ring-primary",
            "min-h-[200px]",
            if isSelected { "border-primary bg-primary/5" } else { "border-border" },
        }
    >
        <input
            type="radio"
            name="foundation_type"
            value={ opt.Type }
            checked?={ isSelected }
            class="sr-only"
            hx-post="/api/wizard/foundation/select"
            hx-target="#foundation-config"
            hx-swap="innerHTML"
        />

        // Diagram
        <div class="h-24 mb-3 flex items-center justify-center">
            <img
                src={ opt.DiagramURL }
                alt=""
                class="max-h-full max-w-full object-contain"
            />
        </div>

        // Label
        <h3 class="font-medium text-base">{ T(ctx, opt.LabelKey) }</h3>

        // Description
        <p class="text-sm text-muted-foreground mt-1 flex-1">
            { T(ctx, opt.DescKey) }
        </p>

        // Metadata badges
        <div class="flex flex-wrap gap-2 mt-3">
            @badge.Badge(variant="outline", size="sm") {
                { T(ctx, "wizard.foundation.cost." + opt.CostLevel) }
            }
            @badge.Badge(variant="outline", size="sm") {
                { T(ctx, "wizard.foundation.complexity." + opt.Complexity) }
            }
        </div>

        // Selection indicator
        if isSelected {
            <div class="absolute top-2 right-2">
                @icon.Icon(name="check-circle", class="text-primary", size="sm")
            </div>
        }
    </label>
}
```

### Anchor Configurator

Configuration UI for anchor type, quantity, and spacing.

```go
package wizard

templ AnchorConfigurator(ctx context.Context, config AnchorConfig, windLoad WindLoadResult) {
    <div class="space-y-6">
        <h3 class="text-lg font-semibold">{ T(ctx, "wizard.foundation.anchors") }</h3>

        // Anchor type selection
        <div class="space-y-2">
            <label for="anchor-type" class="text-sm font-medium">
                { T(ctx, "wizard.foundation.anchor_type") }
            </label>
            <select
                id="anchor-type"
                name="anchor_type"
                class="w-full rounded-md border border-input bg-background px-3 py-2"
                hx-post="/api/wizard/foundation/anchor-calculate"
                hx-target="#anchor-summary"
                hx-trigger="change"
            >
                <option value="auger" selected?={ config.Type == "auger" }>
                    { T(ctx, "wizard.foundation.anchor.auger") }
                </option>
                <option value="ground_screw" selected?={ config.Type == "ground_screw" }>
                    { T(ctx, "wizard.foundation.anchor.ground_screw") }
                </option>
                <option value="duckbill" selected?={ config.Type == "duckbill" }>
                    { T(ctx, "wizard.foundation.anchor.duckbill") }
                </option>
                <option value="concrete" selected?={ config.Type == "concrete" }>
                    { T(ctx, "wizard.foundation.anchor.concrete") }
                </option>
            </select>
        </div>

        // Quantity input
        <div class="space-y-2">
            <label for="anchor-qty" class="text-sm font-medium">
                { T(ctx, "wizard.foundation.anchor_quantity") }
            </label>
            <div class="flex items-center gap-2">
                <button
                    type="button"
                    class="w-11 h-11 rounded-md border flex items-center justify-center"
                    onclick="decrementAnchorQty()"
                    aria-label={ T(ctx, "wizard.foundation.decrease") }
                >-</button>
                <input
                    type="number"
                    id="anchor-qty"
                    name="anchor_quantity"
                    value={ fmt.Sprint(config.Quantity) }
                    min="4"
                    max="50"
                    class="w-20 text-center rounded-md border px-2 py-2"
                    hx-post="/api/wizard/foundation/anchor-calculate"
                    hx-target="#anchor-summary"
                    hx-trigger="change"
                />
                <button
                    type="button"
                    class="w-11 h-11 rounded-md border flex items-center justify-center"
                    onclick="incrementAnchorQty()"
                    aria-label={ T(ctx, "wizard.foundation.increase") }
                >+</button>
            </div>
            <p class="text-sm text-muted-foreground">
                { T(ctx, "wizard.foundation.recommended_qty", windLoad.RecommendedAnchors) }
            </p>
        </div>

        // Anchor summary
        <div id="anchor-summary" aria-live="polite">
            @AnchorSummary(ctx, config, windLoad)
        </div>
    </div>
}

templ AnchorSummary(ctx context.Context, config AnchorConfig, windLoad WindLoadResult) {
    <div class={
        "p-4 rounded-lg border",
        if windLoad.IsSufficient { "bg-green-50 border-green-200" } else { "bg-yellow-50 border-yellow-200" },
    }>
        <div class="flex items-start gap-3">
            if windLoad.IsSufficient {
                @icon.Icon(name="check-circle", class="text-green-600 mt-0.5", size="sm")
            } else {
                @icon.Icon(name="alert-triangle", class="text-yellow-600 mt-0.5", size="sm")
            }
            <div class="space-y-2 flex-1">
                <p class="font-medium">
                    { T(ctx, "wizard.foundation.total_capacity") }:
                    { fmt.Sprintf("%.1f kN", config.TotalUpliftCapacity()) }
                </p>
                <p class="text-sm text-muted-foreground">
                    { T(ctx, "wizard.foundation.required_capacity") }:
                    { fmt.Sprintf("%.1f kN", windLoad.RequiredCapacity) }
                </p>
                if !windLoad.IsSufficient {
                    <p class="text-sm text-yellow-700">
                        { T(ctx, "wizard.foundation.add_more_anchors") }
                    </p>
                }
            </div>
        </div>
    </div>
}
```

### Wind Load Calculator Display

Real-time wind load and uplift calculations.

```go
package wizard

templ WindLoadDisplay(ctx context.Context, calc WindLoadCalculation) {
    <div class="p-4 rounded-lg bg-muted/50 space-y-3">
        <h4 class="font-medium flex items-center gap-2">
            @icon.Icon(name="wind", size="sm")
            { T(ctx, "wizard.foundation.wind_load") }
        </h4>

        <dl class="grid grid-cols-2 gap-x-4 gap-y-2 text-sm">
            <dt class="text-muted-foreground">{ T(ctx, "wizard.foundation.design_wind_speed") }</dt>
            <dd class="font-medium text-right">{ fmt.Sprintf("%.0f km/h", calc.DesignWindSpeed) }</dd>

            <dt class="text-muted-foreground">{ T(ctx, "wizard.foundation.design_pressure") }</dt>
            <dd class="font-medium text-right">{ fmt.Sprintf("%.0f Pa", calc.DesignPressure) }</dd>

            <dt class="text-muted-foreground">{ T(ctx, "wizard.foundation.roof_area") }</dt>
            <dd class="font-medium text-right">{ fmt.Sprintf("%.1f m²", calc.RoofArea) }</dd>

            <dt class="text-muted-foreground border-t pt-2">{ T(ctx, "wizard.foundation.total_uplift") }</dt>
            <dd class="font-medium text-right border-t pt-2">{ fmt.Sprintf("%.1f kN", calc.TotalUpliftForce) }</dd>
        </dl>

        <p class="text-xs text-muted-foreground">
            { T(ctx, "wizard.foundation.wind_note") }
        </p>
    </div>
}
```

### Frost Depth Display

Shows regional frost depth based on farm location.

```go
package wizard

templ FrostDepthDisplay(ctx context.Context, region string, depth float64) {
    <div class="flex items-center gap-3 p-3 rounded-lg border bg-blue-50/50 border-blue-200">
        @icon.Icon(name="thermometer-snowflake", class="text-blue-600", size="md")
        <div>
            <p class="font-medium">
                { T(ctx, "wizard.foundation.frost_depth") }: { fmt.Sprintf("%.1f m", depth) }
            </p>
            <p class="text-sm text-muted-foreground">
                { T(ctx, "wizard.foundation.region") }: { region }
            </p>
        </div>
        @TooltipHelp(T(ctx, "wizard.foundation.frost_depth_tooltip"))
    </div>
}
```

See [FOUNDATIONS_ANCHORS.md](FOUNDATIONS_ANCHORS.md) for complete foundation specifications and Canadian frost depth data.
