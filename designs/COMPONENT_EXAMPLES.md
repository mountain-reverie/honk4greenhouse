# Component Implementation Examples

This document provides implementation examples for key components with:
- Full i18n support
- Accessibility (ARIA labels, focus management)
- Mobile responsiveness (44px touch targets, responsive layouts)
- Live-update HTMX patterns (debounced where appropriate)
- Progressive enhancement (works without JS)

## Layout Components

### base.templ

```go
package layout

import (
    "context"
    . "github.com/mountain-reverie/honk4greenhouse/internal/components"
    "github.com/mountain-reverie/honk4greenhouse/internal/i18n"
)

templ Base(ctx context.Context, title string) {
    <!DOCTYPE html>
    <html lang={ i18n.GetLang(ctx) }>
    <head>
        <meta charset="UTF-8"/>
        <meta name="viewport" content="width=device-width, initial-scale=1.0"/>
        <title>{ title } | { T(ctx, "app.name") }</title>
        <script src="https://unpkg.com/htmx.org@1.9.10"></script>
        <script src="https://cdn.tailwindcss.com"></script>
        @chart.Script()
        @dialog.Script()
        @dropdown.Script()

        // Global HTMX error handling
        <script>
        document.body.addEventListener('htmx:responseError', function(evt) {
            showToast('error', document.documentElement.lang === 'fr-CA'
                ? 'Une erreur s\'est produite. Veuillez reessayer.'
                : 'An error occurred. Please try again.');
        });

        document.body.addEventListener('htmx:sendError', function(evt) {
            showToast('warning', document.documentElement.lang === 'fr-CA'
                ? 'Erreur reseau. Veuillez verifier votre connexion.'
                : 'Network error. Please check your connection.');
        });

        function showToast(variant, message) {
            const container = document.getElementById('toast-container');
            const toast = document.createElement('div');
            const bgClass = variant === 'error' ? 'bg-destructive/10 border-destructive/20'
                          : variant === 'warning' ? 'bg-yellow-50 border-yellow-200'
                          : 'bg-green-50 border-green-200';
            toast.className = `p-4 rounded-lg shadow-lg border ${bgClass} animate-slide-in`;
            toast.setAttribute('role', 'alert');
            toast.innerHTML = `<p class="text-sm">${message}</p>`;
            container.appendChild(toast);
            setTimeout(() => toast.remove(), 5000);
        }
        </script>
    </head>
    <body class="min-h-screen bg-background">
        @shared.SkipLink(ctx)
        { children... }
        @shared.ToastContainer()
        <div id="dialog-container" aria-live="polite"></div>
    </body>
    </html>
}
```

### header.templ

```go
package layout

templ Header(ctx context.Context, user *models.User) {
    <header class="border-b">
        <div class="container mx-auto px-4 h-16 flex items-center justify-between">
            <a href="/" class="text-xl font-bold flex items-center gap-2">
                <span aria-hidden="true">🌱</span>
                <span>{ T(ctx, "app.name") }</span>
            </a>

            <div class="flex items-center gap-2 md:gap-4">
                @LanguageSwitcher(ctx)

                if user != nil {
                    @auth.UserMenu(ctx, user)
                } else {
                    <div class="hidden md:flex gap-2">
                        @button.Button(
                            href="/login",
                            variant="outline",
                            class="min-h-[44px]",
                        ){ { T(ctx, "nav.login") } }
                        @button.Button(
                            href="/login?signup=true",
                            class="min-h-[44px]",
                        ){ { T(ctx, "nav.sign_in") } }
                    </div>
                    // Mobile: single login button
                    <div class="md:hidden">
                        @button.Button(
                            href="/login",
                            class="min-w-[44px] min-h-[44px]",
                        ){ { T(ctx, "nav.login") } }
                    </div>
                }
            </div>
        </div>
    </header>
}
```

### nav_tabs.templ

```go
package layout

templ NavTabs(ctx context.Context, activeTab string, isAuthenticated bool) {
    <nav class="border-b" aria-label={ T(ctx, "nav.main") }>
        <div class="container mx-auto px-4">
            // Desktop tabs
            <div class="hidden md:block">
                @tabs.Tabs(value=activeTab) {
                    @tabs.List() {
                        @tabs.Trigger(
                            value="showcase",
                            hx-get="/showcase",
                            hx-target="#main-content",
                            hx-push-url="true",
                            hx-on::after-swap="document.querySelector('#main-content h1')?.focus()",
                            class="min-h-[44px]",
                        ) {
                            { T(ctx, "nav.showcase") }
                        }
                        if isAuthenticated {
                            @tabs.Trigger(
                                value="shared",
                                hx-get="/shared",
                                hx-target="#main-content",
                                hx-push-url="true",
                                hx-on::after-swap="document.querySelector('#main-content h1')?.focus()",
                                class="min-h-[44px]",
                            ) {
                                { T(ctx, "nav.shared") }
                            }
                            @tabs.Trigger(
                                value="personal",
                                hx-get="/personal",
                                hx-target="#main-content",
                                hx-push-url="true",
                                hx-on::after-swap="document.querySelector('#main-content h1')?.focus()",
                                class="min-h-[44px]",
                            ) {
                                { T(ctx, "nav.personal") }
                            }
                        } else {
                            @tabs.Trigger(
                                value="shared",
                                disabled=true,
                                aria-disabled="true",
                                class="opacity-50 min-h-[44px]",
                            ) {
                                { T(ctx, "nav.shared") }
                            }
                            @tabs.Trigger(
                                value="personal",
                                disabled=true,
                                aria-disabled="true",
                                class="opacity-50 min-h-[44px]",
                            ) {
                                { T(ctx, "nav.personal") }
                            }
                        }
                    }
                }
            </div>

            // Mobile dropdown
            <div class="md:hidden py-2">
                @MobileNavDropdown(ctx, activeTab, isAuthenticated)
            </div>
        </div>
    </nav>
}

templ MobileNavDropdown(ctx context.Context, activeTab string, isAuthenticated bool) {
    @dropdown.Dropdown() {
        @dropdown.Trigger() {
            @button.Button(
                variant="outline",
                class="w-full min-h-[44px] justify-between",
                aria-label={ T(ctx, "nav.current_section") },
            ) {
                { getTabLabel(ctx, activeTab) }
                @icon.Icon(name="chevron-down", size="sm")
            }
        }
        @dropdown.Content(class="w-full") {
            @dropdown.Item(
                hx-get="/showcase",
                hx-target="#main-content",
                hx-push-url="true",
                class=if activeTab == "showcase" { "bg-muted" } else { "" },
            ){ { T(ctx, "nav.showcase") } }
            if isAuthenticated {
                @dropdown.Item(
                    hx-get="/shared",
                    hx-target="#main-content",
                    hx-push-url="true",
                    class=if activeTab == "shared" { "bg-muted" } else { "" },
                ){ { T(ctx, "nav.shared") } }
                @dropdown.Item(
                    hx-get="/personal",
                    hx-target="#main-content",
                    hx-push-url="true",
                    class=if activeTab == "personal" { "bg-muted" } else { "" },
                ){ { T(ctx, "nav.personal") } }
            }
        }
    }
}
```

### language_switcher.templ

```go
package layout

templ LanguageSwitcher(ctx context.Context) {
    <div data-testid="language-switcher">
        @dropdown.Dropdown() {
            @dropdown.Trigger() {
                @button.Button(
                    variant="ghost",
                    size="sm",
                    aria-label={ T(ctx, "nav.change_language") },
                    class="min-w-[44px] min-h-[44px]",
                ) {
                    if i18n.FromContext(ctx) == i18n.LocaleFrCA {
                        FR
                    } else {
                        EN
                    }
                    @icon.Icon(name="chevron-down", size="xs", class="ml-1")
                }
            }
            @dropdown.Content(align="end") {
                @dropdown.Item(
                    href="?lang=en",
                    class=if i18n.FromContext(ctx) == i18n.LocaleEN { "bg-muted" } else { "" },
                ) {
                    English
                }
                @dropdown.Item(
                    href="?lang=fr-CA",
                    class=if i18n.FromContext(ctx) == i18n.LocaleFrCA { "bg-muted" } else { "" },
                ) {
                    Francais
                }
            }
        }
    </div>
}
```

## Wizard Components

### step_indicator.templ

```go
package wizard

templ StepIndicator(ctx context.Context, currentStep int, totalSteps int) {
    <div
        class="flex items-center justify-center gap-1 md:gap-2"
        role="progressbar"
        aria-valuenow={ fmt.Sprint(currentStep) }
        aria-valuemin="1"
        aria-valuemax={ fmt.Sprint(totalSteps) }
        aria-label={ TF(ctx, "wizard.step_of", currentStep, totalSteps) }
    >
        for i := 1; i <= totalSteps; i++ {
            <div class={
                "w-8 h-8 md:w-10 md:h-10 rounded-full flex items-center justify-center border-2 text-sm font-medium",
                if i < currentStep { "bg-primary text-primary-foreground border-primary" },
                if i == currentStep { "border-primary text-primary" },
                if i > currentStep { "border-muted text-muted-foreground" },
            }>
                if i < currentStep {
                    @icon.Icon(name="check", size="sm")
                } else {
                    { fmt.Sprint(i) }
                }
            </div>
            if i < totalSteps {
                <div class={
                    "w-4 md:w-8 h-0.5",
                    if i < currentStep { "bg-primary" } else { "bg-muted" },
                }></div>
            }
        }
    </div>

    // Mobile: text indicator
    <p class="text-center text-sm text-muted-foreground mt-2 md:hidden">
        { TF(ctx, "wizard.step_of", currentStep, totalSteps) }
    </p>
}
```

### step1_shape.templ (Progressive Enhancement)

```go
package wizard

templ Step1Shape(ctx context.Context, farmID string, shapes []models.GreenhouseShape, selectedID string) {
    <div>
        <h2 tabindex="-1" class="text-lg font-semibold mb-2 focus:outline-none">
            { T(ctx, "wizard.step1.title") }
        </h2>
        <p class="text-muted-foreground mb-6">
            { T(ctx, "wizard.step1.subtitle") }
        </p>

        // Form works without JS (radio buttons)
        <form id="step1-form" method="get" action={ templ.SafeURL("/farms/" + farmID + "/projects/new/step/2") }>
            <fieldset>
                <legend class="sr-only">{ T(ctx, "wizard.step1.title") }</legend>

                <div class="grid grid-cols-2 md:grid-cols-3 gap-4" role="radiogroup">
                    for _, shape := range shapes {
                        @ShapeCard(ctx, shape, shape.ID == selectedID)
                    }

                    // Request additional shape
                    <div
                        class="border-2 border-dashed rounded-lg p-4 flex items-center justify-center text-center cursor-pointer hover:bg-muted min-h-[120px]"
                        hx-get="/api/projects/wizard/request-shape-form"
                        hx-target="#dialog-container"
                        role="button"
                        tabindex="0"
                        aria-label={ T(ctx, "wizard.step1.request_shape") }
                    >
                        <span class="text-sm text-muted-foreground">
                            { T(ctx, "wizard.step1.request_shape") }
                        </span>
                    </div>
                </div>
            </fieldset>

            <div class="flex justify-end mt-6">
                // No-JS: submit button
                <noscript>
                    @button.Button(type="submit", class="min-h-[44px]") {
                        { T(ctx, "wizard.next") }
                    }
                </noscript>

                // With JS: HTMX button
                <div class="js-only">
                    @button.Button(
                        type="button",
                        hx-get={ "/farms/" + farmID + "/projects/new/step/2" },
                        hx-target="#wizard-content",
                        hx-include="#step1-form",
                        hx-push-url="true",
                        hx-on::after-swap="document.querySelector('#wizard-content h2')?.focus()",
                        disabled={ selectedID == "" },
                        class="min-h-[44px]",
                    ){ { T(ctx, "wizard.next") } }
                </div>
            </div>
        </form>
    </div>
}

templ ShapeCard(ctx context.Context, shape models.GreenhouseShape, selected bool) {
    <label class="cursor-pointer">
        <input
            type="radio"
            name="shapeID"
            value={ shape.ID }
            checked?={ selected }
            class="peer sr-only"
            required
        />
        <div
            class={
                "border-2 rounded-lg p-4 transition-colors min-h-[120px]",
                "peer-checked:border-primary peer-checked:bg-primary/5",
                "peer-focus-visible:ring-2 peer-focus-visible:ring-ring peer-focus-visible:ring-offset-2",
                "hover:border-muted-foreground",
            }
            data-testid="shape-card"
        >
            <div class="aspect-square bg-muted rounded flex items-center justify-center mb-2">
                <img
                    src={ shape.IconURL }
                    alt=""
                    class="w-12 h-12 md:w-16 md:h-16"
                    aria-hidden="true"
                />
            </div>
            <p class="text-sm text-center font-medium">
                { T(ctx, "shapes." + shape.ID) }
            </p>
            <p class="text-xs text-center text-muted-foreground mt-1 hidden md:block">
                { T(ctx, "shapes." + shape.ID + "_desc") }
            </p>
        </div>
    </label>
}
```

### step2_dimensions.templ (Accessible Compass)

```go
package wizard

templ Step2Dimensions(ctx context.Context, farmID string, data *WizardData) {
    <div>
        <h2 tabindex="-1" class="text-lg font-semibold mb-2 focus:outline-none">
            { T(ctx, "wizard.step2.title") }
        </h2>
        <p class="text-muted-foreground mb-6">
            { T(ctx, "wizard.step2.subtitle") }
        </p>

        <form id="step2-form">
            <input type="hidden" name="shapeID" value={ data.ShapeID }/>

            <div class="grid grid-cols-1 md:grid-cols-2 gap-6">
                // Dimensions section
                <div class="space-y-4">
                    <div class="grid grid-cols-3 gap-2">
                        @DimensionInput(ctx, "length", data.Length, "m")
                        @DimensionInput(ctx, "width", data.Width, "m")
                        @DimensionInput(ctx, "height", data.Height, "m")
                    </div>

                    @OrientationCompass(ctx, data.Orientation)
                </div>

                // Crop profile section
                <div>
                    @form.Item() {
                        @form.Label(for="cropType") {
                            { T(ctx, "wizard.step2.goal") }
                        }
                        <select
                            id="cropType"
                            name="cropType"
                            class="w-full border rounded px-3 py-2 min-h-[44px]"
                            hx-get="/api/crops/profile"
                            hx-target="#crop-profile"
                            hx-trigger="change"
                            hx-include="[name='cropType']"
                        >
                            <option value="tomato" selected?={ data.CropType == "tomato" }>
                                { T(ctx, "crops.tomato") }
                            </option>
                            <option value="lettuce" selected?={ data.CropType == "lettuce" }>
                                { T(ctx, "crops.lettuce") }
                            </option>
                            <option value="pepper" selected?={ data.CropType == "pepper" }>
                                { T(ctx, "crops.pepper") }
                            </option>
                            <option value="cucumber" selected?={ data.CropType == "cucumber" }>
                                { T(ctx, "crops.cucumber") }
                            </option>
                        </select>
                    }

                    <div id="crop-profile" class="mt-4" aria-live="polite">
                        @CropProfile(ctx, data.CropProfile)
                    </div>
                </div>
            </div>

            @WizardNavigation(ctx, farmID, 2, 9)
        </form>
    </div>
}

templ DimensionInput(ctx context.Context, name string, value float64, unit string) {
    @form.Item() {
        @form.Label(for=name) {
            { T(ctx, "wizard.step2." + name) }
        }
        <div class="flex items-center gap-1">
            @input.Input(
                id=name,
                name=name,
                type="number",
                step="0.1",
                min="1",
                max="200",
                value={ fmt.Sprintf("%.1f", value) },
                class="w-full min-h-[44px]",
                required=true,
                hx-post="/api/wizard/calculate-heat-loss",
                hx-target="#heat-loss-display",
                hx-trigger="input changed delay:300ms",
                hx-include="[name]",
            )
            <span class="text-sm text-muted-foreground w-6">{ unit }</span>
        </div>
    }
}

templ OrientationCompass(ctx context.Context, selected string) {
    <fieldset class="mt-6">
        <legend class="font-medium mb-2 flex items-center">
            { T(ctx, "wizard.step2.orientation") }
            @shared.TooltipHelp(ctx, "wizard.step2.orientation_help")
        </legend>

        <div
            class="relative w-36 h-36 md:w-40 md:h-40 mx-auto"
            role="radiogroup"
            aria-label={ T(ctx, "wizard.step2.orientation") }
        >
            // Compass circle
            <div class="absolute inset-0 rounded-full border-2 border-muted"></div>

            // Direction buttons positioned around compass
            for _, dir := range []struct{ id, angle string }{
                {"N", "top-0 left-1/2 -translate-x-1/2 -translate-y-1/2"},
                {"NE", "top-[15%] right-[15%] translate-x-1/2 -translate-y-1/2"},
                {"E", "top-1/2 right-0 translate-x-1/2 -translate-y-1/2"},
                {"SE", "bottom-[15%] right-[15%] translate-x-1/2 translate-y-1/2"},
                {"S", "bottom-0 left-1/2 -translate-x-1/2 translate-y-1/2"},
                {"SW", "bottom-[15%] left-[15%] -translate-x-1/2 translate-y-1/2"},
                {"W", "top-1/2 left-0 -translate-x-1/2 -translate-y-1/2"},
                {"NW", "top-[15%] left-[15%] -translate-x-1/2 -translate-y-1/2"},
            } {
                @CompassButton(ctx, dir.id, dir.angle, selected)
            }

            // Hidden input for form submission
            <input type="hidden" name="orientation" id="orientation-input" value={ selected }/>
        </div>
    </fieldset>
}

templ CompassButton(ctx context.Context, dir string, positionClass string, selected string) {
    <button
        type="button"
        role="radio"
        aria-checked={ fmt.Sprintf("%t", dir == selected) }
        aria-label={ T(ctx, "directions." + dir) }
        class={
            "absolute w-11 h-11 md:w-9 md:h-9 rounded-full text-xs font-bold transition-colors",
            positionClass,
            if dir == selected {
                "bg-primary text-primary-foreground"
            } else {
                "bg-muted hover:bg-muted-foreground/20"
            },
            "focus:outline-none focus:ring-2 focus:ring-ring focus:ring-offset-2",
        }
        onclick="selectOrientation(this.innerText)"
    >
        { dir }
    </button>
}

script selectOrientation(dir) {
    // Update hidden input
    document.getElementById('orientation-input').value = dir;

    // Update button states
    document.querySelectorAll('[role="radio"]').forEach(btn => {
        const isSelected = btn.innerText === dir;
        btn.setAttribute('aria-checked', isSelected);
        btn.classList.toggle('bg-primary', isSelected);
        btn.classList.toggle('text-primary-foreground', isSelected);
        btn.classList.toggle('bg-muted', !isSelected);
    });
}
```

### step4_walls.templ (Material Selection)

```go
package wizard

templ Step4Walls(ctx context.Context, farmID string, data *WizardData, materials []models.InsulationMaterial) {
    <div>
        <h2 tabindex="-1" class="text-lg font-semibold mb-2 focus:outline-none">
            { T(ctx, "wizard.step4.title") }
        </h2>
        <p class="text-muted-foreground mb-6">
            { T(ctx, "wizard.step4.subtitle") }
        </p>

        // Live heat loss display
        <div class="bg-primary/10 p-4 rounded-lg mb-6" aria-live="polite">
            <div class="flex items-center justify-between">
                <span class="text-sm font-medium">
                    { T(ctx, "wizard.step3.calculated_heat_loss") }
                </span>
                <span class="font-bold text-lg" id="heat-loss-display">
                    { fmt.Sprintf("%.0f kWh", data.CalculatedHeatLoss) }
                </span>
            </div>
        </div>

        <form id="step4-form">
            // Carry forward previous step data
            <input type="hidden" name="shapeID" value={ data.ShapeID }/>
            <input type="hidden" name="length" value={ fmt.Sprintf("%.1f", data.Length) }/>
            <input type="hidden" name="width" value={ fmt.Sprintf("%.1f", data.Width) }/>
            <input type="hidden" name="height" value={ fmt.Sprintf("%.1f", data.Height) }/>
            <input type="hidden" name="orientation" value={ data.Orientation }/>
            <input type="hidden" name="cropType" value={ data.CropType }/>
            <input type="hidden" name="glazing" value={ data.Glazing }/>

            <div class="space-y-6">
                // Grid for walls - 1 column mobile, 2 columns desktop
                <div class="grid grid-cols-1 md:grid-cols-2 gap-6">
                    @MaterialSelector(ctx, "wall_north", data.WallSelections["wall_north"], materials)
                    @MaterialSelector(ctx, "wall_south", data.WallSelections["wall_south"], materials)
                    @MaterialSelector(ctx, "wall_east", data.WallSelections["wall_east"], materials)
                    @MaterialSelector(ctx, "wall_west", data.WallSelections["wall_west"], materials)
                </div>
            </div>

            @WizardNavigation(ctx, farmID, 4, 9)
        </form>
    </div>
}

templ MaterialSelector(ctx context.Context, surface string, selection *models.InsulationSelection, materials []models.InsulationMaterial) {
    <div class="space-y-3 p-4 border rounded-lg" id={ "insulation-" + surface }>
        <div class="flex items-center justify-between">
            <label class="font-medium">
                { T(ctx, "wizard.step4." + surface) }
                @shared.TooltipHelp(ctx, "insulation.r_value_help")
            </label>
            <span
                class="text-sm font-mono bg-muted px-2 py-1 rounded"
                id={ "r-value-" + surface }
                aria-live="polite"
            >
                if selection != nil && selection.RValue > 0 {
                    R-{ fmt.Sprintf("%.1f", selection.RValue) }
                } else {
                    --
                }
            </span>
        </div>

        <div class="grid grid-cols-1 gap-3">
            // Material selection
            <div>
                <label for={ "material-" + surface } class="text-sm text-muted-foreground">
                    { T(ctx, "wizard.step4.select_material") }
                </label>
                <select
                    id={ "material-" + surface }
                    name={ "material_" + surface }
                    class="w-full border rounded px-3 py-2 min-h-[44px] mt-1"
                    hx-get={ "/api/insulation/thicknesses?surface=" + surface }
                    hx-target={ "#thickness-container-" + surface }
                    hx-trigger="change"
                    hx-include={ "[name='material_" + surface + "']" }
                >
                    <option value="">{ T(ctx, "wizard.step4.select_material") }</option>
                    for _, m := range materials {
                        <option
                            value={ m.ID }
                            selected?={ selection != nil && selection.MaterialID == m.ID }
                        >
                            { T(ctx, "insulation.materials." + m.ID) }
                        </option>
                    }
                </select>
            </div>

            // Thickness selection (populated via HTMX)
            <div id={ "thickness-container-" + surface }>
                <label for={ "thickness-" + surface } class="text-sm text-muted-foreground">
                    { T(ctx, "wizard.step4.select_thickness") }
                </label>
                <select
                    id={ "thickness-" + surface }
                    name={ "thickness_" + surface }
                    class="w-full border rounded px-3 py-2 min-h-[44px] mt-1"
                    disabled?={ selection == nil || selection.MaterialID == "" }
                    hx-post="/api/wizard/calculate-r-value"
                    hx-target={ "#r-value-" + surface }
                    hx-trigger="change"
                    hx-include={ "[name='material_" + surface + "'], [name='thickness_" + surface + "']" }
                >
                    <option value="">{ T(ctx, "wizard.step4.select_thickness") }</option>
                    if selection != nil {
                        @ThicknessOptions(ctx, selection.MaterialID, selection.ThicknessMM)
                    }
                </select>
            </div>
        </div>
    </div>
}

templ ThicknessOptions(ctx context.Context, materialID string, selectedThickness int) {
    // This is returned by the HTMX endpoint when material changes
    for _, opt := range getMaterialThicknesses(materialID) {
        <option
            value={ fmt.Sprint(opt.Millimeters) }
            selected?={ opt.Millimeters == selectedThickness }
        >
            { fmt.Sprintf("%d mm (R-%.1f)", opt.Millimeters, opt.RValueMetric) }
        </option>
    }
}
```

### step6_air_seal.templ (User-Friendly ACH)

```go
package wizard

templ Step6AirSeal(ctx context.Context, farmID string, data *WizardData) {
    <div>
        <h2 tabindex="-1" class="text-lg font-semibold mb-2 focus:outline-none">
            { T(ctx, "wizard.step6.title") }
        </h2>
        <p class="text-muted-foreground mb-6">
            { T(ctx, "wizard.step6.subtitle") }
        </p>

        <form id="step6-form">
            // Hidden fields for previous step data
            @WizardHiddenFields(data)

            <fieldset class="space-y-4">
                <legend class="sr-only">{ T(ctx, "wizard.step6.air_seal") }</legend>

                // Tight option
                @AirSealOption(ctx, "tight", "0.5", data.AirSeal)

                // Average option
                @AirSealOption(ctx, "average", "1.5", data.AirSeal)

                // Standard option
                @AirSealOption(ctx, "standard", "3.0", data.AirSeal)
            </fieldset>

            // Blower door info box
            @alert.Alert(variant="info", class="mt-6") {
                <div class="flex gap-3">
                    @icon.Icon(name="info", size="sm", class="flex-shrink-0 mt-0.5")
                    <div class="text-sm">
                        <p>{ T(ctx, "wizard.step6.blower_door_info") }</p>
                        <p class="mt-2 text-muted-foreground">
                            { T(ctx, "wizard.step6.can_refine_later") }
                        </p>
                    </div>
                </div>
            }

            @WizardNavigation(ctx, farmID, 6, 9)
        </form>
    </div>
}

templ AirSealOption(ctx context.Context, id string, achValue string, selected string) {
    <label class="block cursor-pointer">
        <input
            type="radio"
            name="airSeal"
            value={ id }
            checked?={ selected == id }
            class="peer sr-only"
            hx-post="/api/wizard/calculate-heat-loss"
            hx-target="#heat-loss-display"
            hx-trigger="change"
            hx-include="[name]"
        />
        <div class={
            "p-4 border-2 rounded-lg transition-colors",
            "peer-checked:border-primary peer-checked:bg-primary/5",
            "peer-focus-visible:ring-2 peer-focus-visible:ring-ring peer-focus-visible:ring-offset-2",
            "hover:border-muted-foreground",
        }>
            <div class="flex items-start justify-between">
                <div>
                    <p class="font-medium">{ T(ctx, "wizard.step6." + id) }</p>
                    <p class="text-sm text-muted-foreground mt-1">
                        { T(ctx, "wizard.step6." + id + "_desc") }
                    </p>
                </div>
                <span class="text-sm font-mono bg-muted px-2 py-1 rounded">
                    { achValue } ACH
                </span>
            </div>
        </div>
    </label>
}
```

### step7_climate_battery.templ (Pre-designed Climate Battery Systems)

```go
package wizard

templ Step7ClimateBattery(ctx context.Context, farmID string, data *WizardData) {
    <div>
        <h2 tabindex="-1" class="text-lg font-semibold mb-2 focus:outline-none">
            { T(ctx, "wizard.step7.title") }
        </h2>
        <p class="text-muted-foreground mb-4">
            { T(ctx, "wizard.step7.subtitle") }
        </p>

        // Info about climate battery
        <div class="bg-muted/50 p-4 rounded-lg mb-6">
            <p class="text-sm">{ T(ctx, "wizard.step7.climate_battery_intro") }</p>
        </div>

        <form id="step7-form">
            @WizardHiddenFields(data)

            <fieldset class="space-y-4">
                <legend class="font-medium mb-2">{ T(ctx, "wizard.step7.select_design") }</legend>

                // No climate battery option
                @ClimateBatteryDesignCard(ctx, "none", data.ClimateBatteryDesign)

                // Multi-pipe design
                @ClimateBatteryDesignCard(ctx, "multi_pipe", data.ClimateBatteryDesign)

                // Diagonal array design (Atmos style)
                @ClimateBatteryDesignCard(ctx, "diagonal_array", data.ClimateBatteryDesign)

                // Single manifold design
                @ClimateBatteryDesignCard(ctx, "single_manifold", data.ClimateBatteryDesign)

                // Custom/Advanced
                @ClimateBatteryDesignCard(ctx, "custom", data.ClimateBatteryDesign)
            </fieldset>

            // Configuration panel (shown when design selected)
            <div id="climate-battery-config" class="mt-6" aria-live="polite">
                if data.ClimateBatteryDesign != "" && data.ClimateBatteryDesign != "none" {
                    @ClimateBatteryConfigPanel(ctx, data.ClimateBatteryDesign, data.ClimateBatteryConfig)
                }
            </div>

            @WizardNavigation(ctx, farmID, 7, 9)
        </form>
    </div>
}

templ ClimateBatteryDesignCard(ctx context.Context, designID string, selected string) {
    <label class="block cursor-pointer">
        <input
            type="radio"
            name="climateBatteryDesign"
            value={ designID }
            checked?={ selected == designID }
            class="peer sr-only"
            hx-post={ "/api/climate-battery/select?design=" + designID }
            hx-target="#climate-battery-config"
            hx-trigger="change"
        />
        <div class={
            "p-4 border-2 rounded-lg transition-colors",
            "peer-checked:border-primary peer-checked:bg-primary/5",
            "peer-focus-visible:ring-2 peer-focus-visible:ring-ring peer-focus-visible:ring-offset-2",
            "hover:border-muted-foreground",
        }>
            <div class="flex flex-col md:flex-row md:items-center gap-4">
                // Diagram (hidden on mobile if no climate battery)
                if designID != "none" {
                    <div class="w-full md:w-32 h-24 bg-muted rounded flex items-center justify-center">
                        @ClimateBatteryDiagram(designID)
                    </div>
                }

                <div class="flex-1">
                    <p class="font-medium">{ T(ctx, "climate_battery." + designID + ".name") }</p>
                    <p class="text-sm text-muted-foreground mt-1">
                        { T(ctx, "climate_battery." + designID + ".desc") }
                    </p>
                </div>

                if designID != "none" && designID != "custom" {
                    <div class="hidden md:block">
                        @badge.Badge(variant="outline") {
                            { getClimateBatteryReduction(designID) }
                        }
                    </div>
                }
            </div>
        </div>
    </label>
}

templ ClimateBatteryDiagram(designID string) {
    // Simple SVG diagrams for each design type
    switch designID {
    case "multi_pipe":
        <svg viewBox="0 0 80 50" class="w-full h-full p-2" aria-hidden="true">
            // Multiple parallel pipes
            for i := 0; i < 4; i++ {
                <rect x="10" y={ fmt.Sprint(10 + i*10) } width="60" height="4" fill="currentColor" opacity="0.3" rx="2"/>
            }
            // Fan symbol
            <circle cx="75" cy="25" r="8" fill="none" stroke="currentColor" stroke-width="2"/>
        </svg>
    case "diagonal_array":
        <svg viewBox="0 0 80 50" class="w-full h-full p-2" aria-hidden="true">
            // Diagonal pipes with manifolds (Atmos style)
            <rect x="5" y="5" width="6" height="40" fill="currentColor" opacity="0.4" rx="1"/>
            <rect x="69" y="5" width="6" height="40" fill="currentColor" opacity="0.4" rx="1"/>
            // Diagonal tubes
            for i := 0; i < 5; i++ {
                <line x1="11" y1={ fmt.Sprint(8 + i*8) } x2="69" y2={ fmt.Sprint(12 + i*8) } stroke="currentColor" stroke-width="2" opacity="0.3"/>
            }
        </svg>
    case "single_manifold":
        <svg viewBox="0 0 80 50" class="w-full h-full p-2" aria-hidden="true">
            // Central manifold
            <rect x="35" y="5" width="10" height="40" fill="currentColor" opacity="0.4" rx="2"/>
            // Pipes radiating
            for i := 0; i < 3; i++ {
                <rect x="5" y={ fmt.Sprint(10 + i*12) } width="30" height="4" fill="currentColor" opacity="0.3" rx="2"/>
                <rect x="45" y={ fmt.Sprint(10 + i*12) } width="30" height="4" fill="currentColor" opacity="0.3" rx="2"/>
            }
        </svg>
    default:
        <span class="text-muted-foreground text-xs">Custom</span>
    }
}
```

### wizard_navigation.templ (Shared Navigation)

```go
package wizard

templ WizardNavigation(ctx context.Context, farmID string, currentStep int, totalSteps int) {
    <div class="flex justify-between mt-8 pt-6 border-t">
        if currentStep > 1 {
            @button.Button(
                variant="outline",
                type="button",
                hx-get={ fmt.Sprintf("/farms/%s/projects/new/step/%d", farmID, currentStep-1) },
                hx-target="#wizard-content",
                hx-push-url="true",
                hx-on::after-swap="document.querySelector('#wizard-content h2')?.focus()",
                class="min-h-[44px] min-w-[100px]",
            ) {
                @icon.Icon(name="arrow-left", size="sm", class="mr-2")
                { T(ctx, "wizard.back") }
            }
        } else {
            <div></div>
        }

        if currentStep < totalSteps {
            @button.Button(
                type="button",
                hx-get={ fmt.Sprintf("/farms/%s/projects/new/step/%d", farmID, currentStep+1) },
                hx-target="#wizard-content",
                hx-include="form",
                hx-push-url="true",
                hx-on::after-swap="document.querySelector('#wizard-content h2')?.focus()",
                class="min-h-[44px] min-w-[100px]",
            ) {
                { T(ctx, "wizard.next") }
                @icon.Icon(name="arrow-right", size="sm", class="ml-2")
            }
        } else {
            @button.Button(
                type="button",
                hx-post={ "/farms/" + farmID + "/projects" },
                hx-include="form",
                hx-target="#main-content",
                class="min-h-[44px]",
            ) {
                { T(ctx, "wizard.create_project") }
                @icon.Icon(name="check", size="sm", class="ml-2")
            }
        }
    </div>
}
```

## Quote Components (Mobile Card View)

### quotes_table.templ

```go
package quotes

templ QuotesTable(ctx context.Context, projectID string, quotes []models.Quote, canSubmit bool) {
    @card.Card() {
        @card.Header() {
            <div class="flex items-center justify-between">
                @card.Title(){ { T(ctx, "quotes.title") } }
                if canSubmit {
                    @button.Button(
                        variant="outline",
                        size="sm",
                        hx-get={ "/api/projects/" + projectID + "/quotes/form" },
                        hx-target="#dialog-container",
                        class="min-h-[44px]",
                    ) {
                        @icon.Icon(name="plus", size="sm", class="mr-1")
                        { T(ctx, "quotes.add_quote") }
                    }
                }
            </div>
        }
        @card.Content() {
            <div id="quotes-table">
                // Desktop table
                <div class="hidden md:block overflow-x-auto">
                    @table.Table() {
                        @table.Header() {
                            @table.Row() {
                                @table.Head(){ { T(ctx, "quotes.builder") } }
                                @table.Head(class="text-right"){ { T(ctx, "quotes.price") } }
                                @table.Head(){ { T(ctx, "quotes.duration") } }
                                @table.Head(){ { T(ctx, "quotes.pdf") } }
                                @table.Head(){ { T(ctx, "quotes.status") } }
                            }
                        }
                        @table.Body() {
                            for _, q := range quotes {
                                @QuoteRow(ctx, q)
                            }
                        }
                    }
                </div>

                // Mobile cards
                <div class="md:hidden space-y-4">
                    for _, q := range quotes {
                        @QuoteCard(ctx, q)
                    }
                </div>

                if len(quotes) == 0 {
                    @emptystate.EmptyState() {
                        @emptystate.Icon() {
                            @icon.Icon(name="file-text", size="xl", class="text-muted-foreground")
                        }
                        @emptystate.Title(){ { T(ctx, "quotes.no_quotes") } }
                        @emptystate.Description(){ { T(ctx, "quotes.no_quotes_desc") } }
                    }
                }
            </div>
        }
    }
}

templ QuoteCard(ctx context.Context, q models.Quote) {
    @card.Card() {
        @card.Content(class="pt-4") {
            <div class="flex items-start justify-between mb-3">
                <div>
                    <p class="font-medium">{ q.BuilderName }</p>
                    <p class="text-2xl font-bold mt-1">${ formatMoney(q.Price) }</p>
                </div>
                @QuoteStatusBadge(ctx, q.NeedsReview)
            </div>

            <dl class="space-y-2 text-sm">
                <div class="flex justify-between">
                    <dt class="text-muted-foreground">{ T(ctx, "quotes.duration") }</dt>
                    <dd>{ q.Duration }</dd>
                </div>
                if q.PDFURL != "" {
                    <div class="flex justify-between items-center">
                        <dt class="text-muted-foreground">{ T(ctx, "quotes.pdf") }</dt>
                        <dd>
                            @button.Button(
                                href=q.PDFURL,
                                target="_blank",
                                variant="ghost",
                                size="sm",
                                class="min-h-[44px]",
                            ) {
                                @icon.Icon(name="file", size="sm", class="mr-1")
                                { T(ctx, "quotes.view_pdf") }
                            }
                        </dd>
                    </div>
                }
            </dl>
        }
    }
}

templ QuoteStatusBadge(ctx context.Context, needsReview bool) {
    if needsReview {
        @badge.Badge(variant="warning", data-testid="needs-review-badge") {
            @icon.Icon(name="alert-triangle", size="xs", class="mr-1")
            { T(ctx, "quotes.needs_review") }
        }
    } else {
        @badge.Badge(variant="success") {
            @icon.Icon(name="check", size="xs", class="mr-1")
            { T(ctx, "quotes.current") }
        }
    }
}
```

## Publication Status (Icons + Text)

### approval_status.templ

```go
package publication

templ ApprovalStatus(ctx context.Context, req models.PublicationRequest) {
    <div id="publication-status">
        @card.Card() {
            @card.Header() {
                <div class="flex items-center justify-between">
                    @card.Title(){ { T(ctx, "publication.status") } }
                    @StatusBadge(ctx, req.Status)
                </div>
            }
            @card.Content() {
                <div class="space-y-3">
                    for _, p := range req.Participants {
                        @ParticipantRow(ctx, p)
                    }
                </div>
            }
            @card.Footer() {
                if req.CanRetrigger() {
                    @button.Button(
                        variant="outline",
                        hx-post={ "/api/projects/" + req.ProjectID + "/publication/retrigger" },
                        hx-target="#publication-status",
                        class="min-h-[44px]",
                    ) {
                        @icon.Icon(name="refresh-cw", size="sm", class="mr-2")
                        { T(ctx, "publication.retrigger") }
                    }
                }
            }
        }
    </div>
}

templ StatusBadge(ctx context.Context, status models.PublicationStatus) {
    switch status {
    case models.StatusPending:
        @badge.Badge(variant="warning") {
            @icon.Icon(name="clock", size="xs", class="mr-1")
            { T(ctx, "publication.pending") }
        }
    case models.StatusApproved:
        @badge.Badge(variant="success") {
            @icon.Icon(name="check-circle", size="xs", class="mr-1")
            { T(ctx, "publication.approved") }
        }
    case models.StatusDenied:
        @badge.Badge(variant="destructive") {
            @icon.Icon(name="x-circle", size="xs", class="mr-1")
            { T(ctx, "publication.denied") }
        }
    }
}

templ ParticipantRow(ctx context.Context, p models.ParticipantApproval) {
    <div class="flex items-center justify-between p-3 rounded bg-muted/50">
        <div class="flex items-center gap-3">
            @avatar.Avatar(src=p.AvatarURL, size="sm", fallback=string(p.Name[0]))
            <div>
                <span class="font-medium">{ p.Name }</span>
                <span class="text-sm text-muted-foreground ml-2">{ string(p.Role) }</span>
            </div>
        </div>
        <div class="flex items-center gap-2">
            switch p.ApprovalStatus {
            case "approved":
                @icon.Icon(name="check-circle", size="sm", class="text-green-600")
                <span class="text-sm text-green-600">{ T(ctx, "publication.approved") }</span>
            case "denied":
                @icon.Icon(name="x-circle", size="sm", class="text-destructive")
                <span class="text-sm text-destructive">{ T(ctx, "publication.denied") }</span>
                if p.DenialReason != "" {
                    @tooltip.Tooltip() {
                        @tooltip.Trigger() {
                            @icon.Icon(name="info", size="xs", class="text-muted-foreground cursor-help")
                        }
                        @tooltip.Content() {
                            { p.DenialReason }
                        }
                    }
                }
            default:
                @icon.Icon(name="clock", size="sm", class="text-yellow-600")
                <span class="text-sm text-yellow-600">{ T(ctx, "publication.pending") }</span>
            }
        </div>
    </div>
}
```

## Responsive Chart Component

### climate_chart.templ

```go
package charts

templ ClimateChart(ctx context.Context, id string, data []models.MonthlyData) {
    @ChartToggle(ctx, id, "chart")
    <div
        class="aspect-[4/3] md:aspect-[16/9] min-h-[200px] max-h-[400px]"
        data-testid="climate-chart"
    >
        @chart.Chart(
            variant="line",
            data=buildClimateChartData(ctx, data),
            responsive=true,
            maintainAspectRatio=false,
            plugins=chart.Plugins{
                legend: chart.Legend{
                    position: "bottom",
                    labels: chart.LegendLabels{
                        boxWidth: 12,
                        padding: 16,
                    },
                },
            },
        )
    </div>
}

func buildClimateChartData(ctx context.Context, data []models.MonthlyData) chart.Data {
    labels := make([]string, len(data))
    temps := make([]float64, len(data))
    rain := make([]float64, len(data))

    for i, m := range data {
        labels[i] = T(ctx, "climate.months." + strings.ToLower(m.Month))
        temps[i] = m.AvgTemp
        rain[i] = m.Rainfall
    }

    return chart.Data{
        Labels: labels,
        Datasets: []chart.Dataset{
            {
                Label:       T(ctx, "climate.metrics.avg_temp") + " (C)",
                Data:        temps,
                BorderColor: "#ef4444",
                Fill:        false,
            },
            {
                Label:       T(ctx, "climate.metrics.rainfall"),
                Data:        rain,
                BorderColor: "#3b82f6",
                Fill:        false,
            },
        },
    }
}
```
