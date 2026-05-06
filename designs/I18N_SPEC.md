# Internationalization (i18n) Specification

## Overview

This document defines the internationalization architecture for honk4greenhouse, enabling multilingual support with a focus on maintainability and extensibility.

## Supported Languages

### Initial Release
- **English (en)** - Default language
- **French (fr-CA)** - Canadian French (legal requirement for Canadian markets)

### Planned Extensions
- Spanish (es)
- Korean (ko)
- Simplified Chinese (zh-Hans)

## Architecture

### Translation Package Structure

```text
internal/i18n/
├── i18n.go           # Core translation functions
├── loader.go         # Translation file loader
├── middleware.go     # Gin middleware for locale detection
├── translations/
│   ├── en.yaml       # English translations
│   ├── fr-CA.yaml    # Canadian French translations
│   └── schema.yaml   # Translation key schema (for validation)
└── types.go          # Locale and translation types
```

### Core Types

```go
package i18n

import (
    "context"
    "sync"
)

// Locale represents a supported locale
type Locale string

const (
    LocaleEN   Locale = "en"
    LocaleFrCA Locale = "fr-CA"
    LocaleES   Locale = "es"
    LocaleKO   Locale = "ko"
    LocaleZH   Locale = "zh-Hans"
)

// DefaultLocale is used when no locale is specified or detected
const DefaultLocale = LocaleEN

// SupportedLocales lists all currently active locales
var SupportedLocales = []Locale{LocaleEN, LocaleFrCA}

// Translator provides translation functionality
type Translator struct {
    translations map[Locale]map[string]string
    mu           sync.RWMutex
}

// contextKey for storing locale in context
type contextKey string

const localeContextKey contextKey = "locale"

// FromContext retrieves the locale from context
func FromContext(ctx context.Context) Locale {
    if locale, ok := ctx.Value(localeContextKey).(Locale); ok {
        return locale
    }
    return DefaultLocale
}

// WithLocale adds locale to context
func WithLocale(ctx context.Context, locale Locale) context.Context {
    return context.WithValue(ctx, localeContextKey, locale)
}
```

### Translation Function

```go
package i18n

import (
    "fmt"
    "strings"
)

// Global translator instance
var globalTranslator *Translator

// Init initializes the global translator with translation files
func Init(translationsDir string) error {
    t, err := NewTranslator(translationsDir)
    if err != nil {
        return err
    }
    globalTranslator = t
    return nil
}

// T translates a key for the given locale
// Usage: i18n.T(ctx, "wizard.step1.title")
func T(ctx context.Context, key string) string {
    locale := FromContext(ctx)
    return globalTranslator.Translate(locale, key)
}

// TF translates a key with format arguments
// Usage: i18n.TF(ctx, "farm.created", farmName)
func TF(ctx context.Context, key string, args ...any) string {
    locale := FromContext(ctx)
    template := globalTranslator.Translate(locale, key)
    return fmt.Sprintf(template, args...)
}

// TP translates with pluralization
// Usage: i18n.TP(ctx, "projects.count", count)
func TP(ctx context.Context, key string, count int) string {
    locale := FromContext(ctx)

    // Try plural forms: key.zero, key.one, key.few, key.many, key.other
    pluralKey := key + "." + getPluralForm(locale, count)
    if translation := globalTranslator.Translate(locale, pluralKey); translation != pluralKey {
        return fmt.Sprintf(translation, count)
    }

    // Fallback to base key
    return fmt.Sprintf(globalTranslator.Translate(locale, key), count)
}

// Translate returns the translation for a key
func (t *Translator) Translate(locale Locale, key string) string {
    t.mu.RLock()
    defer t.mu.RUnlock()

    // Try requested locale
    if translations, ok := t.translations[locale]; ok {
        if value, ok := translations[key]; ok {
            return value
        }
    }

    // Fallback to default locale
    if locale != DefaultLocale {
        if translations, ok := t.translations[DefaultLocale]; ok {
            if value, ok := translations[key]; ok {
                return value
            }
        }
    }

    // Return key if no translation found (helps identify missing translations)
    return key
}

// getPluralForm returns the CLDR plural form for a count
func getPluralForm(locale Locale, count int) string {
    // French and English use simple one/other rules
    switch locale {
    case LocaleFrCA:
        if count == 0 || count == 1 {
            return "one"
        }
        return "other"
    default: // English
        if count == 1 {
            return "one"
        }
        return "other"
    }
}
```

### Gin Middleware

```go
package i18n

import (
    "strings"

    "github.com/gin-gonic/gin"
)

// Middleware detects and sets the locale for each request
func Middleware() gin.HandlerFunc {
    return func(c *gin.Context) {
        locale := detectLocale(c)
        ctx := WithLocale(c.Request.Context(), locale)
        c.Request = c.Request.WithContext(ctx)

        // Also set in Gin context for templates
        c.Set("locale", locale)
        c.Set("lang", string(locale))

        c.Next()
    }
}

// detectLocale determines the locale from various sources
func detectLocale(c *gin.Context) Locale {
    // 1. Check URL query parameter (?lang=fr-CA)
    if lang := c.Query("lang"); lang != "" {
        if locale := parseLocale(lang); locale != "" {
            setLocaleCookie(c, locale)
            return locale
        }
    }

    // 2. Check cookie
    if cookie, err := c.Cookie("locale"); err == nil {
        if locale := parseLocale(cookie); locale != "" {
            return locale
        }
    }

    // 3. Check Accept-Language header
    if acceptLang := c.GetHeader("Accept-Language"); acceptLang != "" {
        if locale := parseAcceptLanguage(acceptLang); locale != "" {
            return locale
        }
    }

    // 4. Default
    return DefaultLocale
}

// parseLocale validates and returns a supported locale
func parseLocale(lang string) Locale {
    lang = strings.ToLower(strings.TrimSpace(lang))

    // Direct match
    for _, supported := range SupportedLocales {
        if strings.ToLower(string(supported)) == lang {
            return supported
        }
    }

    // Language code match (e.g., "fr" matches "fr-CA")
    for _, supported := range SupportedLocales {
        if strings.HasPrefix(strings.ToLower(string(supported)), lang) {
            return supported
        }
    }

    return ""
}

// parseAcceptLanguage parses the Accept-Language header
func parseAcceptLanguage(header string) Locale {
    // Simple parsing - split by comma and check each
    parts := strings.Split(header, ",")
    for _, part := range parts {
        lang := strings.Split(strings.TrimSpace(part), ";")[0]
        if locale := parseLocale(lang); locale != "" {
            return locale
        }
    }
    return ""
}

func setLocaleCookie(c *gin.Context, locale Locale) {
    c.SetCookie("locale", string(locale), 60*60*24*365, "/", "", false, true)
}
```

## Translation File Format

Using YAML for readability and ease of editing by non-developers.

### English (en.yaml)

```yaml
# Application
app:
  name: "Honk for Greenhouse"
  tagline: "Design efficient greenhouses for Canadian climates"

# Navigation
nav:
  showcase: "Showcase"
  shared: "Shared"
  personal: "Personal"
  login: "Login"
  sign_in: "Sign In"
  logout: "Logout"
  settings: "Settings"

# Authentication
auth:
  login_required: "Login Required"
  continue_with_email: "Continue with Email"
  or_continue_with: "Or continue with"
  email_placeholder: "you@example.com"

# Farm Management
farm:
  add_farm: "Add Farm"
  edit_farm: "Edit Farm"
  name: "Name"
  location: "Location"
  country: "Country"
  province: "Province/State"
  city: "City"
  other_owners: "Other Owners"
  add_owner: "Add"
  energy_cost: "Energy Cost"
  add_energy: "+ Add"
  save_farm: "Save Farm"
  cancel: "Cancel"
  created: "Farm '%s' created successfully"
  deleted: "Farm deleted"

# Energy Types
energy:
  electricity: "Electricity"
  natural_gas: "Natural Gas"
  propane: "Propane"
  wood: "Wood (Pine)"
  wood_hardwood: "Wood (Hardwood)"
  pellets: "Wood Pellets"
  units:
    kwh: "$/kWh"
    m3: "$/m3"
    gallon: "$/gallon"
    cord: "$/cord"
    tonne: "$/tonne"

# Project Creation Wizard
wizard:
  title: "Creating a New Greenhouse"
  step1:
    title: "Select Greenhouse Shape"
    subtitle: "Choose a structure type that fits your needs"
    request_shape: "Request Additional Shape"
  step2:
    title: "Dimensions & Crop Profile"
    subtitle: "Define size and what you want to grow"
    length: "Length"
    width: "Width"
    height: "Height"
    orientation: "Orientation"
    orientation_help: "The direction your greenhouse's glazed (south) face will point"
    goal: "Goal"
    season: "Season"
    min_temp: "Min Temp"
    max_temp: "Max Temp"
    light: "Light"
    water: "Water"
  step3:
    title: "Glazing Selection"
    subtitle: "Choose your transparent covering material"
    calculated_heat_loss: "Estimated Annual Heat Loss"
  step4:
    title: "Wall Insulation"
    subtitle: "Select materials for non-glazed walls"
    wall_north: "North Wall"
    wall_south: "South Wall"
    wall_east: "East Wall"
    wall_west: "West Wall"
    select_material: "Select Material"
    select_thickness: "Select Thickness"
    resulting_r_value: "R-Value"
  step5:
    title: "Roof & Foundation"
    subtitle: "Configure roof insulation and frost protection"
    roof_north: "North Roof"
    roof_south: "South Roof"
    frost_skirt: "Frost Skirt"
    frost_skirt_help: "Insulation around the perimeter to prevent ground freezing"
    depth: "Depth"
    width: "Width"
    insulation: "Insulation"
  step6:
    title: "Air Sealing"
    subtitle: "Manage air infiltration for energy efficiency"
    air_seal: "Air Tightness"
    air_seal_help: "Air changes per hour (ACH) - lower is better"
    blower_door_info: "A blower door test after construction can precisely measure your ACH. For now, select an estimate based on construction quality."
    can_refine_later: "You can update this value after construction"
    tight: "Tight (0.5 ACH)"
    tight_desc: "Professional air sealing, blower door tested"
    average: "Average (1.5 ACH)"
    average_desc: "Careful construction with sealed joints"
    standard: "Standard (3 ACH)"
    standard_desc: "Typical greenhouse construction"
  step7:
    title: "Heat Transfer System"
    subtitle: "Optional ground-coupled heating and cooling"
    climate_battery_intro: "Climate battery systems store excess daytime heat in the soil for release at night."
    select_design: "Select a Design"
    no_climate_battery: "No Climate Battery"
    no_climate_battery_desc: "Skip this optimization for simpler construction"
  step8:
    title: "Additional Optimizations"
    subtitle: "Optional features to improve thermal performance"
    night_blanket: "Thermal Night Blanket"
    night_blanket_desc: "Insulating cover deployed at night to reduce heat loss"
    row_cover: "Row Covers"
    row_cover_desc: "Floating fabric covers over plants for frost protection"
    thermal_mass: "Thermal Mass (Water Barrels)"
    thermal_mass_desc: "Water containers that absorb and release heat"
    add_optimization: "+ Add"
    reconfigure: "Configure"
  step9:
    title: "Heating Sources"
    subtitle: "Configure primary and backup heating"
    peak_heat_need: "Peak Heat Need"
    annual_energy_demand: "Annual Energy Demand"
    primary: "Primary"
    secondary: "Secondary"
    percentage: "Percentage"
    create_project: "Create Project"

  # Navigation
  next: "Next"
  back: "Back"
  step_of: "Step %d of %d"

# Greenhouse Shapes
shapes:
  quonset: "Quonset"
  quonset_desc: "Curved roof, efficient use of materials"
  gable: "Gable"
  gable_desc: "Traditional peaked roof design"
  gothic: "Gothic Arch"
  gothic_desc: "Pointed arch for snow shedding"
  shed: "Lean-to / Shed"
  shed_desc: "Attached to existing structure"
  a_frame: "A-Frame"
  a_frame_desc: "Steep roof angles for cold climates"
  chinese: "Chinese Passive Solar"
  chinese_desc: "North-insulated with thermal mass wall"

# Insulation Materials
insulation:
  materials:
    spray_foam_closed: "Closed-Cell Spray Foam"
    spray_foam_open: "Open-Cell Spray Foam"
    fiberglass_batt: "Fiberglass Batts"
    mineral_wool: "Mineral Wool (Rockwool)"
    rigid_xps: "Rigid Foam (XPS)"
    rigid_eps: "Rigid Foam (EPS)"
    rigid_polyiso: "Polyisocyanurate"
  thickness:
    mm: "%d mm"
    cm: "%d cm"
  r_value: "R-%.1f"
  r_value_help: "Higher R-value = better insulation"

# Glazing Materials
glazing:
  single_poly: "Single Layer Polyethylene"
  double_poly: "Double Layer Polyethylene (Inflated)"
  polycarbonate_twin: "Twin-Wall Polycarbonate (8mm)"
  polycarbonate_triple: "Triple-Wall Polycarbonate (16mm)"
  polycarbonate_five: "Five-Wall Polycarbonate (32mm)"
  glass_single: "Single Pane Glass"
  glass_double: "Double Pane Glass"

# Climate Battery Systems
climate_battery:
  multi_pipe:
    name: "Multi-Pipe with Fan"
    desc: "Multiple short pipes with optimized airflow. Best efficiency for moderate climates."
  two_trench:
    name: "Two-Trench Design"
    desc: "Parallel trenches with connecting pipes. Good for longer greenhouses."
  single_manifold:
    name: "Single Manifold (Both Sides)"
    desc: "Central manifold splitting to both sides. Balanced heat distribution."
  custom:
    name: "Custom / Advanced"
    desc: "Design your own system with specific pipe layout and sizing."

# Climate Data
climate:
  chart: "Chart"
  table: "Table"
  edit: "Edit Climate Data"
  locked: "Climate data is locked"
  unlock: "Click to unlock for editing"
  save_version: "Save as New Version"
  cancel: "Cancel"
  version: "Version"
  months:
    jan: "Jan"
    feb: "Feb"
    mar: "Mar"
    apr: "Apr"
    may: "May"
    jun: "Jun"
    jul: "Jul"
    aug: "Aug"
    sep: "Sep"
    oct: "Oct"
    nov: "Nov"
    dec: "Dec"
  metrics:
    avg_temp: "Avg Temp"
    min_temp: "Min"
    max_temp: "Max"
    rainfall: "Rain (mm)"
    sunlight: "Sun (hrs)"

# Projects
project:
  add_project: "+ Add Project"
  thermal_mass: "Thermal Mass"
  equipment: "Equipment"
  production_period: "Production Period"
  expected_production: "Expected Production"
  expected_revenue: "Expected Revenue"
  expected_op_cost: "Operating Cost"
  monthly_cost: "Monthly Cost"
  heat_need: "Heat Need"
  cooling_need: "Cooling Need"

# Quotes
quotes:
  title: "Builder Quotes"
  add_quote: "+ Add Quote"
  builder: "Builder"
  price: "Price"
  duration: "Duration"
  pdf: "PDF"
  status: "Status"
  needs_review: "Needs Review"
  current: "Current"
  submit_quote: "Submit Quote"

# Sharing
sharing:
  share_project: "Share Project"
  share_description: "Invite others to view or collaborate on this project."
  recipient_email: "Recipient Email"
  permission_level: "Permission Level"
  view_only: "View only"
  view_only_desc: "Can view and leave feedback"
  can_edit: "Can edit"
  can_edit_desc: "Can make changes to the project"
  allow_reshare: "Allow recipient to reshare"
  allow_copy: "Allow recipient to copy/fork"
  edit_warning: "Granting edit access will mark existing quotes as 'needs review' when changes are made."
  share: "Share"
  cancel: "Cancel"

# Publication
publication:
  request: "Request Publication"
  status: "Publication Status"
  approved: "Approved"
  pending: "Pending"
  denied: "Denied"
  retrigger: "Re-send Requests"
  photos_required: "Photos Required"
  photos_ready: "Photos Ready"

# Actions Menu
actions:
  share: "Share"
  transfer: "Transfer"
  delete: "Delete"
  confirm_delete: "Are you sure you want to delete this?"
  confirm_delete_farm: "Are you sure you want to delete this farm? All projects and climate data will also be deleted."
  confirm_delete_project: "Are you sure you want to delete this project? All quotes and photos will also be deleted."

# Units (all metric)
units:
  meters: "m"
  centimeters: "cm"
  square_meters: "m2"
  cubic_meters: "m3"
  celsius: "C"
  watts: "W"
  kilowatts: "kW"
  kilowatt_hours: "kWh"
  r_value_metric: "m2-K/W"

# Errors
errors:
  generic: "An error occurred. Please try again."
  network: "Network error. Please check your connection."
  not_found: "Not found"
  forbidden: "You don't have permission to access this"
  validation: "Please check your input and try again"
  required_field: "This field is required"
  invalid_email: "Please enter a valid email address"

# Success Messages
success:
  saved: "Changes saved"
  created: "Created successfully"
  deleted: "Deleted successfully"
  shared: "Shared successfully"
```

### French Canadian (fr-CA.yaml)

```yaml
# Application
app:
  name: "Honk pour Serre"
  tagline: "Concevez des serres efficaces pour le climat canadien"

# Navigation
nav:
  showcase: "Vitrine"
  shared: "Partage"
  personal: "Personnel"
  login: "Connexion"
  sign_in: "S'inscrire"
  logout: "Deconnexion"
  settings: "Parametres"

# Authentication
auth:
  login_required: "Connexion requise"
  continue_with_email: "Continuer avec courriel"
  or_continue_with: "Ou continuer avec"
  email_placeholder: "vous@exemple.com"

# Farm Management
farm:
  add_farm: "Ajouter une ferme"
  edit_farm: "Modifier la ferme"
  name: "Nom"
  location: "Emplacement"
  country: "Pays"
  province: "Province/Etat"
  city: "Ville"
  other_owners: "Autres proprietaires"
  add_owner: "Ajouter"
  energy_cost: "Cout energetique"
  add_energy: "+ Ajouter"
  save_farm: "Enregistrer"
  cancel: "Annuler"
  created: "Ferme '%s' creee avec succes"
  deleted: "Ferme supprimee"

# Energy Types
energy:
  electricity: "Electricite"
  natural_gas: "Gaz naturel"
  propane: "Propane"
  wood: "Bois (Pin)"
  wood_hardwood: "Bois (Franc)"
  pellets: "Granules de bois"
  units:
    kwh: "$/kWh"
    m3: "$/m3"
    gallon: "$/gallon"
    cord: "$/corde"
    tonne: "$/tonne"

# Project Creation Wizard
wizard:
  title: "Creation d'une nouvelle serre"
  step1:
    title: "Selectionner la forme"
    subtitle: "Choisissez un type de structure adapte a vos besoins"
    request_shape: "Demander une forme supplementaire"
  step2:
    title: "Dimensions et profil de culture"
    subtitle: "Definissez la taille et ce que vous voulez cultiver"
    length: "Longueur"
    width: "Largeur"
    height: "Hauteur"
    orientation: "Orientation"
    orientation_help: "La direction vers laquelle la face vitree (sud) de votre serre pointera"
    goal: "Objectif"
    season: "Saison"
    min_temp: "Temp. min"
    max_temp: "Temp. max"
    light: "Lumiere"
    water: "Eau"
  step3:
    title: "Selection du vitrage"
    subtitle: "Choisissez votre materiau de couverture transparent"
    calculated_heat_loss: "Perte de chaleur annuelle estimee"
  step4:
    title: "Isolation des murs"
    subtitle: "Selectionnez les materiaux pour les murs non vitres"
    wall_north: "Mur nord"
    wall_south: "Mur sud"
    wall_east: "Mur est"
    wall_west: "Mur ouest"
    select_material: "Choisir le materiau"
    select_thickness: "Choisir l'epaisseur"
    resulting_r_value: "Valeur R"
  step5:
    title: "Toit et fondation"
    subtitle: "Configurez l'isolation du toit et la protection contre le gel"
    roof_north: "Toit nord"
    roof_south: "Toit sud"
    frost_skirt: "Jupe antigel"
    frost_skirt_help: "Isolation autour du perimetre pour empecher le gel du sol"
    depth: "Profondeur"
    width: "Largeur"
    insulation: "Isolation"
  step6:
    title: "Etancheite a l'air"
    subtitle: "Gerez l'infiltration d'air pour l'efficacite energetique"
    air_seal: "Etancheite"
    air_seal_help: "Changements d'air par heure (CAH) - plus bas est mieux"
    blower_door_info: "Un test d'infiltrometrie apres la construction peut mesurer precisement votre CAH. Pour l'instant, selectionnez une estimation basee sur la qualite de construction."
    can_refine_later: "Vous pourrez mettre a jour cette valeur apres la construction"
    tight: "Etanche (0.5 CAH)"
    tight_desc: "Etancheite professionnelle, teste par infiltrometrie"
    average: "Moyenne (1.5 CAH)"
    average_desc: "Construction soignee avec joints scelles"
    standard: "Standard (3 CAH)"
    standard_desc: "Construction de serre typique"
  step7:
    title: "Systeme de transfert de chaleur"
    subtitle: "Chauffage et refroidissement geothermique optionnel"
    climate_battery_intro: "Les systemes de batterie climatique stockent l'exces de chaleur diurne dans le sol pour la liberer la nuit."
    select_design: "Selectionnez un design"
    no_climate_battery: "Aucune batterie climatique"
    no_climate_battery_desc: "Ignorer cette optimisation pour une construction plus simple"
  step8:
    title: "Optimisations supplementaires"
    subtitle: "Fonctionnalites optionnelles pour ameliorer la performance thermique"
    night_blanket: "Couverture thermique nocturne"
    night_blanket_desc: "Couverture isolante deployee la nuit pour reduire les pertes de chaleur"
    row_cover: "Couvertures de rangs"
    row_cover_desc: "Tissus flottants sur les plantes pour la protection contre le gel"
    thermal_mass: "Masse thermique (barils d'eau)"
    thermal_mass_desc: "Conteneurs d'eau qui absorbent et liberent la chaleur"
    add_optimization: "+ Ajouter"
    reconfigure: "Configurer"
  step9:
    title: "Sources de chauffage"
    subtitle: "Configurez le chauffage principal et de secours"
    peak_heat_need: "Besoin de chaleur maximal"
    annual_energy_demand: "Demande energetique annuelle"
    primary: "Principal"
    secondary: "Secondaire"
    percentage: "Pourcentage"
    create_project: "Creer le projet"

  # Navigation
  next: "Suivant"
  back: "Retour"
  step_of: "Etape %d sur %d"

# Greenhouse Shapes
shapes:
  quonset: "Quonset"
  quonset_desc: "Toit courbe, utilisation efficace des materiaux"
  gable: "Pignon"
  gable_desc: "Design traditionnel a toit en pointe"
  gothic: "Arc gothique"
  gothic_desc: "Arc pointu pour l'evacuation de la neige"
  shed: "Appentis"
  shed_desc: "Attache a une structure existante"
  a_frame: "Charpente en A"
  a_frame_desc: "Angles de toit prononces pour climats froids"
  chinese: "Solaire passive chinoise"
  chinese_desc: "Isolee au nord avec mur de masse thermique"

# Insulation Materials
insulation:
  materials:
    spray_foam_closed: "Mousse a cellules fermees"
    spray_foam_open: "Mousse a cellules ouvertes"
    fiberglass_batt: "Matelas de fibre de verre"
    mineral_wool: "Laine minerale (Rockwool)"
    rigid_xps: "Mousse rigide (XPS)"
    rigid_eps: "Mousse rigide (EPS)"
    rigid_polyiso: "Polyisocyanurate"
  thickness:
    mm: "%d mm"
    cm: "%d cm"
  r_value: "R-%.1f"
  r_value_help: "Valeur R plus elevee = meilleure isolation"

# Climate Battery Systems
climate_battery:
  multi_pipe:
    name: "Multi-tuyaux avec ventilateur"
    desc: "Plusieurs tuyaux courts avec flux d'air optimise. Meilleure efficacite pour climats moderes."
  two_trench:
    name: "Design a deux tranchees"
    desc: "Tranchees paralleles avec tuyaux de connexion. Bon pour les serres longues."
  single_manifold:
    name: "Collecteur unique (deux cotes)"
    desc: "Collecteur central divisant vers les deux cotes. Distribution de chaleur equilibree."
  custom:
    name: "Personnalise / Avance"
    desc: "Concevez votre propre systeme avec disposition et dimensionnement specifiques."

# Climate Data
climate:
  chart: "Graphique"
  table: "Tableau"
  edit: "Modifier les donnees climatiques"
  locked: "Les donnees climatiques sont verrouillees"
  unlock: "Cliquez pour deverrouiller et modifier"
  save_version: "Enregistrer comme nouvelle version"
  cancel: "Annuler"
  version: "Version"
  months:
    jan: "Jan"
    feb: "Fev"
    mar: "Mar"
    apr: "Avr"
    may: "Mai"
    jun: "Juin"
    jul: "Juil"
    aug: "Aout"
    sep: "Sep"
    oct: "Oct"
    nov: "Nov"
    dec: "Dec"
  metrics:
    avg_temp: "Temp. moy."
    min_temp: "Min"
    max_temp: "Max"
    rainfall: "Pluie (mm)"
    sunlight: "Soleil (h)"

# Actions Menu
actions:
  share: "Partager"
  transfer: "Transferer"
  delete: "Supprimer"
  confirm_delete: "Etes-vous sur de vouloir supprimer ceci?"
  confirm_delete_farm: "Etes-vous sur de vouloir supprimer cette ferme? Tous les projets et donnees climatiques seront egalement supprimes."
  confirm_delete_project: "Etes-vous sur de vouloir supprimer ce projet? Tous les devis et photos seront egalement supprimes."

# Units (all metric)
units:
  meters: "m"
  centimeters: "cm"
  square_meters: "m2"
  cubic_meters: "m3"
  celsius: "C"
  watts: "W"
  kilowatts: "kW"
  kilowatt_hours: "kWh"
  r_value_metric: "m2-K/W"

# Errors
errors:
  generic: "Une erreur s'est produite. Veuillez reessayer."
  network: "Erreur reseau. Veuillez verifier votre connexion."
  not_found: "Non trouve"
  forbidden: "Vous n'avez pas la permission d'acceder a ceci"
  validation: "Veuillez verifier vos informations et reessayer"
  required_field: "Ce champ est requis"
  invalid_email: "Veuillez entrer une adresse courriel valide"

# Success Messages
success:
  saved: "Modifications enregistrees"
  created: "Cree avec succes"
  deleted: "Supprime avec succes"
  shared: "Partage avec succes"
```

## Templ Component Integration

### Translation Helper Component

```go
package components

import (
    "context"
    "github.com/mountain-reverie/honk4greenhouse/internal/i18n"
)

// T is a convenience wrapper for templates
func T(ctx context.Context, key string) string {
    return i18n.T(ctx, key)
}

// TF is a convenience wrapper for formatted translations
func TF(ctx context.Context, key string, args ...any) string {
    return i18n.TF(ctx, key, args...)
}

// GetLang returns the current language code for HTML lang attribute
func GetLang(ctx context.Context) string {
    return string(i18n.FromContext(ctx))
}
```

### Usage in Templ Templates

```go
package wizard

import (
    "context"
    . "github.com/mountain-reverie/honk4greenhouse/internal/components"
)

templ Step1Shape(ctx context.Context, farmID string, shapes []models.GreenhouseShape, selectedID string) {
    <div>
        <h2 tabindex="-1" class="text-lg font-semibold mb-2 focus:outline-none">
            { T(ctx, "wizard.step1.title") }
        </h2>
        <p class="text-muted-foreground mb-4">
            { T(ctx, "wizard.step1.subtitle") }
        </p>

        // ... shape grid

        <button type="button" class="...">
            { T(ctx, "wizard.step1.request_shape") }
        </button>
    </div>
}
```

### Language Switcher Component

```go
package layout

templ LanguageSwitcher(ctx context.Context, currentLocale string) {
    <div class="relative" data-testid="language-switcher">
        @dropdown.Dropdown() {
            @dropdown.Trigger() {
                @button.Button(variant="ghost", size="sm", aria-label=T(ctx, "nav.language")) {
                    if currentLocale == "fr-CA" {
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
                    class=if currentLocale == "en" { "bg-muted" } else { "" },
                ) {
                    English
                }
                @dropdown.Item(
                    href="?lang=fr-CA",
                    class=if currentLocale == "fr-CA" { "bg-muted" } else { "" },
                ) {
                    Francais
                }
            }
        }
    </div>
}
```

## Adding New Languages

### Process

1. Create new translation file (e.g., `es.yaml`)
2. Add locale constant to `types.go`
3. Add locale to `SupportedLocales` slice
4. Update plural rules if needed in `getPluralForm()`
5. Test with `?lang=es` query parameter

### Translation File Template

```yaml
# [Language Name] ([locale code])
# Copy this template and translate all values

app:
  name: ""  # Keep "Honk for Greenhouse" or translate
  tagline: ""

nav:
  showcase: ""
  shared: ""
  # ... continue for all keys
```

## Testing

### Translation Coverage Test

```go
func TestTranslationCoverage(t *testing.T) {
    translator := i18n.NewTranslator("translations/")

    // Get all keys from English (source of truth)
    enKeys := translator.GetAllKeys(i18n.LocaleEN)

    for _, locale := range i18n.SupportedLocales {
        if locale == i18n.LocaleEN {
            continue
        }

        t.Run(string(locale), func(t *testing.T) {
            localeKeys := translator.GetAllKeys(locale)

            for _, key := range enKeys {
                if _, exists := localeKeys[key]; !exists {
                    t.Errorf("Missing translation for key %q in locale %s", key, locale)
                }
            }
        })
    }
}
```

### E2E Language Switching Test

```go
func TestLanguageSwitching(t *testing.T) {
    pw := playwright.New(t)
    defer pw.Close()

    page := pw.NewPage()

    // Start in English
    page.Goto(baseURL + "/")
    expect(page.Locator("h1")).ToContainText("Honk for Greenhouse")

    // Switch to French
    page.Click("[data-testid='language-switcher']")
    page.Click("text=Francais")

    // Verify French content
    expect(page.Locator("h1")).ToContainText("Honk pour Serre")

    // Verify cookie persistence
    page.Reload()
    expect(page.Locator("h1")).ToContainText("Honk pour Serre")
}
```
