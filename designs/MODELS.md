# Data Models

## User

```go
type User struct {
    ID        string
    Email     string
    Name      string
    AvatarURL string
    Role      Role // owner, designer, builder, manufacturer, admin
    CreatedAt time.Time
}

type Role string

const (
    RoleOwner        Role = "owner"
    RoleCoOwner      Role = "co_owner"
    RoleDesigner     Role = "designer"
    RoleBuilder      Role = "builder"
    RoleManufacturer Role = "manufacturer"
    RoleAdmin        Role = "admin"
)
```

## Farm

```go
type Farm struct {
    ID          string
    OwnerID     string
    CoOwnerIDs  []string    // Co-owners with shared access
    Name        string
    Location    Location
    ClimateZone string      // "Zone 4a"
    EnergyCosts []EnergyCost
    CreatedAt   time.Time
    UpdatedAt   time.Time
}

type Location struct {
    Country   string
    Province  string      // Province/State
    City      string
    Latitude  float64
    Longitude float64
}

type EnergyCost struct {
    ID       string
    FarmID   string
    FuelType string  // "electricity", "wood", "natural_gas", "propane"
    Unit     string  // "kWh", "cord", "m3", "gallon"
    Cost     float64 // cost per unit
}
```

## Climate Data

```go
type ClimateVersion struct {
    ID        string
    FarmID    string
    Version   string        // "2025-12-03" (YYYY-MM-DD)
    CreatedAt time.Time
    CreatedBy string        // User ID
    Data      []MonthlyData
}

type MonthlyData struct {
    Month         string  // "Jan", "Feb", ...
    AvgTemp       float64
    MinTemp       float64 // Extreme min
    MaxTemp       float64 // Extreme max
    Rainfall      float64 // mm
    SunlightHours float64
}
```

## Project

```go
type Project struct {
    ID                 string
    FarmID             string
    Name               string
    
    // Step 1: Shape
    GreenhouseShape    string   // Shape ID/type selected
    
    // Step 2: Dimensions & Profile
    Dimensions         Dimensions
    Orientation        string   // "N", "NE", "E", "SE", "S", "SW", "W", "NW"
    CropGoal           CropProfile
    
    // Step 3: Insulation
    Insulation         InsulationConfig
    CalculatedBTU      float64  // Calculated BTU/year
    
    // Step 4: Optimization & Heating
    PeakHeatNeed       float64  // kW
    AnnualEnergyDemand float64  // kWh
    Optimizations      []Optimization
    HeatingSources     []HeatingSource
    
    // Legacy/Display
    ThermalMass        string   // "Clay 100m x 5m x 1m"
    Equipment          []string // ["4x 24\" Fan"]
    ExpectedProduction string   // "20T Tomato"
    ExpectedRevenue    float64  // 50000
    ExpectedOpCost     float64  // 6000/year
    
    PhotoURLs          []string
    IsPublished        bool
    CreatedAt          time.Time
    UpdatedAt          time.Time
}

type Dimensions struct {
    Length float64 // meters
    Width  float64
    Height float64
}

type CropProfile struct {
    CropType   string  // "Tomato", "Lettuce", etc.
    SeasonStart string // "February"
    SeasonEnd   string // "November"
    MinTemp     float64 // °C
    MaxTemp     float64 // °C
    LightNeeds  string  // Reference to light profile
    WaterNeeds  string  // Reference to water profile
}

type InsulationConfig struct {
    Glazing     string       // "double_poly", "triple_wall_poly", etc.
    WallNorth   string       // R-value, e.g. "R22"
    WallSouth   string
    WallEast    string
    WallWest    string
    RoofNorth   string
    RoofSouth   string
    FrostSkirt  FrostSkirt
    AirSeal     string       // "good-3ach", "tight-1ach", etc.
}

type FrostSkirt struct {
    Depth      float64 // cm
    Width      float64 // m
    Insulation string  // R-value
}

type Optimization struct {
    ID     string
    Type   string  // "gaht", "night_blanket", "row_cover", "thermal_mass"
    Config string  // JSON config specific to type
}

type HeatingSource struct {
    Priority   int     // 1 = primary, 2 = secondary
    FuelType   string  // "electric", "wood", "gas"
    Percentage float64 // 0-100
}

type Period struct {
    Start string // "February"
    End   string // "November"
}
```

## Quote

```go
type Quote struct {
    ID          string
    ProjectID   string
    BuilderID   string
    BuilderName string
    Price       float64
    Duration    string    // "1 Month", "2 Months"
    PDFURL      string    // S3 URL
    NeedsReview bool
    CreatedAt   time.Time
    UpdatedAt   time.Time
}
```

## Sharing

```go
type Share struct {
    ID           string
    ProjectID    string
    SharedBy     string      // User ID
    SharedWith   string      // User ID or email
    Permission   Permission
    AllowReshare bool
    AllowCopy    bool
    CreatedAt    time.Time
}

type Permission string

const (
    PermissionView Permission = "view"
    PermissionEdit Permission = "edit"
)
```

## Publication

```go
type PublicationRequest struct {
    ID           string
    ProjectID    string
    RequesterID  string
    Status       PublicationStatus
    Participants []ParticipantApproval
    CreatedAt    time.Time
    UpdatedAt    time.Time
}

type PublicationStatus string

const (
    StatusPending  PublicationStatus = "pending"
    StatusApproved PublicationStatus = "approved"
    StatusDenied   PublicationStatus = "denied"
)

type ParticipantApproval struct {
    UserID         string
    Role           Role
    ApprovalStatus string // "pending", "approved", "denied"
    DenialReason   string
    Token          string     // For email link
    RespondedAt    *time.Time
}
```
