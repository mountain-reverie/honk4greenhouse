# Services & Infrastructure

## S3 Storage

```go
package storage

import (
    "context"
    "fmt"
    "io"
    "path"
    
    "github.com/aws/aws-sdk-go-v2/aws"
    "github.com/aws/aws-sdk-go-v2/service/s3"
)

type S3Storage struct {
    client *s3.Client
    bucket string
    cdnURL string // Optional CDN prefix
}

func NewS3Storage(client *s3.Client, bucket, cdnURL string) *S3Storage {
    return &S3Storage{
        client: client,
        bucket: bucket,
        cdnURL: cdnURL,
    }
}

// Photos

func (s *S3Storage) UploadFarmPhoto(ctx context.Context, farmID, photoID string, r io.Reader, contentType string) (string, error) {
    key := fmt.Sprintf("farms/%s/photos/%s", farmID, photoID)
    return s.upload(ctx, key, r, contentType)
}

func (s *S3Storage) UploadProjectPhoto(ctx context.Context, projectID, photoID string, r io.Reader, contentType string) (string, error) {
    key := fmt.Sprintf("projects/%s/photos/%s", projectID, photoID)
    return s.upload(ctx, key, r, contentType)
}

func (s *S3Storage) DeletePhoto(ctx context.Context, key string) error {
    _, err := s.client.DeleteObject(ctx, &s3.DeleteObjectInput{
        Bucket: aws.String(s.bucket),
        Key:    aws.String(key),
    })
    return err
}

// PDFs

func (s *S3Storage) UploadQuotePDF(ctx context.Context, projectID, quoteID string, r io.Reader) (string, error) {
    key := fmt.Sprintf("projects/%s/quotes/%s.pdf", projectID, quoteID)
    return s.upload(ctx, key, r, "application/pdf")
}

// Climate Versions

func (s *S3Storage) SaveClimateVersion(ctx context.Context, farmID, version string, data []byte) error {
    key := fmt.Sprintf("climate/%s/%s.json", farmID, version)
    _, err := s.client.PutObject(ctx, &s3.PutObjectInput{
        Bucket:      aws.String(s.bucket),
        Key:         aws.String(key),
        Body:        bytes.NewReader(data),
        ContentType: aws.String("application/json"),
    })
    return err
}

func (s *S3Storage) GetClimateVersion(ctx context.Context, farmID, version string) ([]byte, error) {
    key := fmt.Sprintf("climate/%s/%s.json", farmID, version)
    resp, err := s.client.GetObject(ctx, &s3.GetObjectInput{
        Bucket: aws.String(s.bucket),
        Key:    aws.String(key),
    })
    if err != nil {
        return nil, err
    }
    defer resp.Body.Close()
    return io.ReadAll(resp.Body)
}

// Helper

func (s *S3Storage) upload(ctx context.Context, key string, r io.Reader, contentType string) (string, error) {
    _, err := s.client.PutObject(ctx, &s3.PutObjectInput{
        Bucket:      aws.String(s.bucket),
        Key:         aws.String(key),
        Body:        r,
        ContentType: aws.String(contentType),
    })
    if err != nil {
        return "", err
    }
    
    if s.cdnURL != "" {
        return s.cdnURL + "/" + key, nil
    }
    return fmt.Sprintf("https://%s.s3.amazonaws.com/%s", s.bucket, key), nil
}
```

## Email Notifications

```go
package notifications

import (
    "bytes"
    "context"
    "html/template"
    
    "github.com/aws/aws-sdk-go-v2/service/ses"
)

type EmailService struct {
    client    *ses.Client
    fromEmail string
    baseURL   string
    templates map[string]*template.Template
}

func NewEmailService(client *ses.Client, fromEmail, baseURL string) *EmailService {
    svc := &EmailService{
        client:    client,
        fromEmail: fromEmail,
        baseURL:   baseURL,
        templates: make(map[string]*template.Template),
    }
    svc.loadTemplates()
    return svc
}

func (s *EmailService) loadTemplates() {
    s.templates["share_invite"] = template.Must(template.ParseFiles("templates/email/share_invite.html"))
    s.templates["publication_request"] = template.Must(template.ParseFiles("templates/email/publication_request.html"))
    s.templates["publication_approved"] = template.Must(template.ParseFiles("templates/email/publication_approved.html"))
    s.templates["publication_denied"] = template.Must(template.ParseFiles("templates/email/publication_denied.html"))
    s.templates["quote_submitted"] = template.Must(template.ParseFiles("templates/email/quote_submitted.html"))
    s.templates["quote_needs_review"] = template.Must(template.ParseFiles("templates/email/quote_needs_review.html"))
}

// Share Invitation

type ShareInviteData struct {
    RecipientName string
    SenderName    string
    ProjectName   string
    Permission    string
    ViewURL       string
}

func (s *EmailService) SendShareInvite(ctx context.Context, to string, data ShareInviteData) error {
    data.ViewURL = s.baseURL + "/shared"
    return s.send(ctx, to, "You've been invited to view a greenhouse project", "share_invite", data)
}

// Publication Request

type PublicationRequestData struct {
    RecipientName string
    RecipientRole string
    RequesterName string
    ProjectName   string
    ApproveURL    string
    DenyURL       string
    ViewURL       string
}

func (s *EmailService) SendPublicationRequest(ctx context.Context, to, token string, data PublicationRequestData) error {
    data.ApproveURL = s.baseURL + "/publication/approve?token=" + token
    data.DenyURL = s.baseURL + "/publication/deny?token=" + token
    data.ViewURL = s.baseURL + "/projects/" + data.ProjectName
    return s.send(ctx, to, "Publication approval requested", "publication_request", data)
}

// Publication Result

type PublicationResultData struct {
    RecipientName string
    ProjectName   string
    Status        string
    Reason        string
    ViewURL       string
}

func (s *EmailService) SendPublicationResult(ctx context.Context, to string, approved bool, data PublicationResultData) error {
    tmpl := "publication_approved"
    subject := "Your project is now on the showcase!"
    if !approved {
        tmpl = "publication_denied"
        subject = "Publication request update"
    }
    return s.send(ctx, to, subject, tmpl, data)
}

// Quote Notifications

type QuoteNotificationData struct {
    RecipientName string
    BuilderName   string
    ProjectName   string
    Price         string
    ViewURL       string
}

func (s *EmailService) SendQuoteSubmitted(ctx context.Context, to string, data QuoteNotificationData) error {
    return s.send(ctx, to, "New quote received for "+data.ProjectName, "quote_submitted", data)
}

func (s *EmailService) SendQuoteNeedsReview(ctx context.Context, to string, data QuoteNotificationData) error {
    return s.send(ctx, to, "Quote needs review", "quote_needs_review", data)
}

// Helper

func (s *EmailService) send(ctx context.Context, to, subject, templateName string, data any) error {
    var body bytes.Buffer
    if err := s.templates[templateName].Execute(&body, data); err != nil {
        return err
    }
    
    _, err := s.client.SendEmail(ctx, &ses.SendEmailInput{
        Source: aws.String(s.fromEmail),
        Destination: &types.Destination{
            ToAddresses: []string{to},
        },
        Message: &types.Message{
            Subject: &types.Content{Data: aws.String(subject)},
            Body: &types.Body{
                Html: &types.Content{Data: aws.String(body.String())},
            },
        },
    })
    return err
}
```

## Publication Service

```go
package services

import (
    "context"
    "crypto/rand"
    "encoding/hex"
    "time"
)

type PublicationService struct {
    db    *Database
    email *EmailService
}

func (s *PublicationService) RequestPublication(ctx context.Context, projectID, requesterID string) (*models.PublicationRequest, error) {
    // Get all participants
    participants, err := s.db.GetProjectParticipants(ctx, projectID)
    if err != nil {
        return nil, err
    }
    
    // Create request
    req := &models.PublicationRequest{
        ID:          generateID(),
        ProjectID:   projectID,
        RequesterID: requesterID,
        Status:      models.StatusPending,
        CreatedAt:   time.Now(),
        UpdatedAt:   time.Now(),
    }
    
    // Create participant approvals
    for _, p := range participants {
        token := generateToken()
        approval := models.ParticipantApproval{
            UserID:         p.UserID,
            Name:           p.Name,
            Email:          p.Email,
            Role:           p.Role,
            AvatarURL:      p.AvatarURL,
            ApprovalStatus: "pending",
            Token:          token,
        }
        
        // Auto-approve requester
        if p.UserID == requesterID {
            approval.ApprovalStatus = "approved"
            now := time.Now()
            approval.RespondedAt = &now
        }
        
        req.Participants = append(req.Participants, approval)
    }
    
    // Save to database
    if err := s.db.CreatePublicationRequest(ctx, req); err != nil {
        return nil, err
    }
    
    // Send emails to pending participants
    project, _ := s.db.GetProject(ctx, projectID)
    requester, _ := s.db.GetUser(ctx, requesterID)
    
    for _, p := range req.Participants {
        if p.ApprovalStatus == "pending" {
            s.email.SendPublicationRequest(ctx, p.Email, p.Token, notifications.PublicationRequestData{
                RecipientName: p.Name,
                RecipientRole: string(p.Role),
                RequesterName: requester.Name,
                ProjectName:   project.Name,
            })
        }
    }
    
    return req, nil
}

func (s *PublicationService) Approve(ctx context.Context, token string) error {
    req, participant, err := s.db.GetPublicationByToken(ctx, token)
    if err != nil {
        return err
    }
    
    now := time.Now()
    participant.ApprovalStatus = "approved"
    participant.RespondedAt = &now
    
    if err := s.db.UpdateParticipantApproval(ctx, req.ID, participant); err != nil {
        return err
    }
    
    // Check if all approved
    if s.checkAllApproved(ctx, req.ID) {
        s.publishProject(ctx, req.ProjectID)
    }
    
    return nil
}

func (s *PublicationService) Deny(ctx context.Context, token, reason string) error {
    req, participant, err := s.db.GetPublicationByToken(ctx, token)
    if err != nil {
        return err
    }
    
    now := time.Now()
    participant.ApprovalStatus = "denied"
    participant.DenialReason = reason
    participant.RespondedAt = &now
    
    if err := s.db.UpdateParticipantApproval(ctx, req.ID, participant); err != nil {
        return err
    }
    
    // Update request status
    req.Status = models.StatusDenied
    s.db.UpdatePublicationRequest(ctx, req)
    
    // Notify requester
    s.notifyDenial(ctx, req, participant)
    
    return nil
}

func (s *PublicationService) Retrigger(ctx context.Context, requestID string) error {
    req, err := s.db.GetPublicationRequest(ctx, requestID)
    if err != nil {
        return err
    }
    
    // Reset pending/denied to pending, generate new tokens
    for i := range req.Participants {
        p := &req.Participants[i]
        if p.ApprovalStatus != "approved" {
            p.ApprovalStatus = "pending"
            p.DenialReason = ""
            p.Token = generateToken()
            p.RespondedAt = nil
        }
    }
    
    req.Status = models.StatusPending
    req.UpdatedAt = time.Now()
    
    if err := s.db.UpdatePublicationRequest(ctx, req); err != nil {
        return err
    }
    
    // Resend emails
    s.sendPendingEmails(ctx, req)
    
    return nil
}

func generateToken() string {
    b := make([]byte, 32)
    rand.Read(b)
    return hex.EncodeToString(b)
}
```

## Climate Version Service

```go
package services

import (
    "context"
    "encoding/json"
    "time"
)

type ClimateService struct {
    db      *Database
    storage *S3Storage
}

func (s *ClimateService) GetCurrentVersion(ctx context.Context, farmID string) (*models.ClimateVersion, error) {
    return s.db.GetLatestClimateVersion(ctx, farmID)
}

func (s *ClimateService) GetVersion(ctx context.Context, farmID, version string) (*models.ClimateVersion, error) {
    return s.db.GetClimateVersion(ctx, farmID, version)
}

func (s *ClimateService) ListVersions(ctx context.Context, farmID string) ([]models.ClimateVersion, error) {
    return s.db.ListClimateVersions(ctx, farmID)
}

func (s *ClimateService) SaveVersion(ctx context.Context, farmID, userID string, data []models.MonthlyData) (*models.ClimateVersion, error) {
    version := time.Now().Format("2006-01-02")
    
    cv := &models.ClimateVersion{
        ID:        generateID(),
        FarmID:    farmID,
        Version:   version,
        CreatedAt: time.Now(),
        CreatedBy: userID,
        Data:      data,
    }
    
    // Save to database
    if err := s.db.CreateClimateVersion(ctx, cv); err != nil {
        return nil, err
    }
    
    // Backup to S3
    jsonData, _ := json.Marshal(cv)
    s.storage.SaveClimateVersion(ctx, farmID, version, jsonData)
    
    return cv, nil
}
```

## Calculation Service

The calculation service handles heat loss and heating calculations for the project creation wizard.
All calculations use **metric units** internally:
- Temperature: Celsius (C)
- Energy: Watts (W) for power, kilowatt-hours (kWh) for annual energy
- Thermal resistance: m2-K/W (SI R-value)
- Dimensions: meters (m), square meters (m2), cubic meters (m3)

```go
package services

import "math"

type CalculationService struct {
    climateData  *ClimateDataService
    materialData *MaterialDataService
}

func NewCalculationService(climate *ClimateDataService, material *MaterialDataService) *CalculationService {
    return &CalculationService{
        climateData:  climate,
        materialData: material,
    }
}

// CalculateAnnualHeatLoss calculates the annual heat loss in kWh for a greenhouse
// based on dimensions, insulation, and climate data.
// Formula: Q = A * (Ti - To) / R, integrated over heating degree days
func (s *CalculationService) CalculateAnnualHeatLoss(state *WizardState) float64 {
    // Get climate data for the farm's location (HDD in Celsius-days)
    climate := s.climateData.GetHeatingDegreeDays(state.FarmID)

    // Calculate surface areas (m2)
    area := s.calculateSurfaceArea(state)

    // Calculate heat loss through each surface (kWh)
    var totalKWh float64

    // Wall heat loss
    totalKWh += s.calculateHeatLoss(area.WallNorth, state.Insulation.WallNorth, climate.HDD)
    totalKWh += s.calculateHeatLoss(area.WallSouth, state.Insulation.WallSouth, climate.HDD)
    totalKWh += s.calculateHeatLoss(area.WallEast, state.Insulation.WallEast, climate.HDD)
    totalKWh += s.calculateHeatLoss(area.WallWest, state.Insulation.WallWest, climate.HDD)

    // Roof heat loss (typically higher due to glazing)
    totalKWh += s.calculateGlazingHeatLoss(area.RoofNorth, state.Insulation.RoofNorth,
        state.Insulation.Glazing, climate.HDD)
    totalKWh += s.calculateGlazingHeatLoss(area.RoofSouth, state.Insulation.RoofSouth,
        state.Insulation.Glazing, climate.HDD)

    // Air infiltration losses
    volume := state.Dimensions.Length * state.Dimensions.Width * state.Dimensions.Height // m3
    totalKWh += s.calculateInfiltrationLoss(volume, state.Insulation.AirSeal, climate.HDD)

    // Frost skirt benefit (reduces perimeter heat loss)
    if state.Insulation.FrostSkirt.Depth > 0 {
        perimeter := 2 * (state.Dimensions.Length + state.Dimensions.Width) // m
        totalKWh -= s.calculateFrostSkirtSavings(perimeter, state.Insulation.FrostSkirt)
    }

    return totalKWh
}

// CalculateHeating calculates peak heat need (Watts) and annual energy demand (kWh)
// peakHeat: Maximum heating power required on coldest design day (W)
// annualEnergy: Total heating energy required over the year (kWh)
func (s *CalculationService) CalculateHeating(state *WizardState) (peakHeat float64, annualEnergy float64) {
    baseKWh := s.CalculateAnnualHeatLoss(state)

    // Apply crop profile adjustments
    cropFactor := s.getCropHeatingFactor(state.CropProfile)
    adjustedKWh := baseKWh * cropFactor

    // Apply optimization reductions
    for _, opt := range state.Optimizations {
        reduction := s.getOptimizationReduction(opt.Type)
        adjustedKWh *= (1 - reduction)
    }

    // Peak heat need (Watts) on coldest design day
    // Formula: Q = A * deltaT / R (W)
    climate := s.climateData.GetDesignTemperatures(state.FarmID)
    targetTempC := s.getCropTargetTemp(state.CropProfile) // Celsius
    deltaTK := targetTempC - climate.ColdestTempC         // Temperature difference in K (same magnitude as C)

    area := s.calculateTotalSurfaceArea(state) // m2
    avgRValue := s.calculateAverageRValue(state) // m2-K/W

    // Peak heat loss through envelope (W)
    envelopeLoss := (area * deltaTK) / avgRValue

    // Add infiltration heat loss at peak conditions (W)
    volume := state.Dimensions.Length * state.Dimensions.Width * state.Dimensions.Height // m3
    ach := s.getACH(state.Insulation.AirSeal)
    // Infiltration: Q = V * ACH * rho * Cp * deltaT / 3600
    // For air at sea level: rho * Cp ≈ 1200 J/(m3·K)
    infiltrationLoss := (volume * ach * 1200 * deltaTK) / 3600 // W

    peakHeat = envelopeLoss + infiltrationLoss

    // Annual energy demand (kWh)
    annualEnergy = adjustedKWh

    return peakHeat, annualEnergy
}

// BalanceHeatingSources ensures heating source percentages sum to 100%
func (s *CalculationService) BalanceHeatingSources(sources []HeatingSource) []HeatingSource {
    if len(sources) == 0 {
        return sources
    }
    
    // Calculate total percentage
    var total float64
    for _, src := range sources {
        total += src.Percentage
    }
    
    // If total isn't 100%, proportionally adjust
    if total != 100 && total > 0 {
        factor := 100 / total
        for i := range sources {
            sources[i].Percentage *= factor
        }
    }
    
    return sources
}

// Surface area calculation helpers (all areas in m2)
type SurfaceAreas struct {
    WallNorth  float64 // m2
    WallSouth  float64 // m2
    WallEast   float64 // m2
    WallWest   float64 // m2
    RoofNorth  float64 // m2
    RoofSouth  float64 // m2
    Floor      float64 // m2
}

func (s *CalculationService) calculateSurfaceArea(state *WizardState) SurfaceAreas {
    d := state.Dimensions // All dimensions in meters

    // Get shape-specific calculations
    shape := s.materialData.GetShape(state.ShapeID)

    switch shape.Type {
    case "quonset":
        // Curved roof calculation (semicircular cross-section)
        roofArea := math.Pi * (d.Width / 2) * d.Length // m2
        return SurfaceAreas{
            WallNorth: d.Width * d.Height / 2,  // End wall (semicircle approximation)
            WallSouth: d.Width * d.Height / 2,  // End wall
            WallEast:  0,                        // No side walls on quonset
            WallWest:  0,
            RoofNorth: roofArea / 2,
            RoofSouth: roofArea / 2,
            Floor:     d.Length * d.Width,
        }

    case "a-frame":
        // Triangular roof
        roofSlope := math.Sqrt(math.Pow(d.Width/2, 2) + math.Pow(d.Height, 2)) // m
        roofArea := roofSlope * d.Length * 2 // m2
        return SurfaceAreas{
            WallNorth: (d.Width * d.Height) / 2, // Triangle end
            WallSouth: (d.Width * d.Height) / 2,
            WallEast:  d.Height * d.Length / 2,  // Side knee wall
            WallWest:  d.Height * d.Length / 2,
            RoofNorth: roofArea / 2,
            RoofSouth: roofArea / 2,
            Floor:     d.Length * d.Width,
        }

    default: // rectangular/gable
        return SurfaceAreas{
            WallNorth: d.Width * d.Height,
            WallSouth: d.Width * d.Height,
            WallEast:  d.Length * d.Height,
            WallWest:  d.Length * d.Height,
            RoofNorth: d.Length * (d.Width / 2),
            RoofSouth: d.Length * (d.Width / 2),
            Floor:     d.Length * d.Width,
        }
    }
}

// calculateHeatLoss computes annual heat loss through a surface (kWh)
// area: surface area in m2
// rValue: thermal resistance identifier (lookup returns m2-K/W)
// hdd: heating degree days in Celsius-days
func (s *CalculationService) calculateHeatLoss(area float64, rValue string, hdd float64) float64 {
    r := s.materialData.GetRValue(rValue) // m2-K/W

    // Heat loss formula: Q = A * HDD * 24 / R (Wh)
    // Convert to kWh: divide by 1000
    // HDD is in Celsius-days, which equals K-days for differences
    return (area * hdd * 24) / (r * 1000) // kWh
}

// calculateGlazingHeatLoss computes heat loss through glazed surfaces (kWh)
func (s *CalculationService) calculateGlazingHeatLoss(area float64, rValue string, glazingType string, hdd float64) float64 {
    r := s.materialData.GetRValue(rValue) // m2-K/W
    glazingFactor := s.materialData.GetGlazingFactor(glazingType) // multiplier for glazing type

    return (area * hdd * 24 * glazingFactor) / (r * 1000) // kWh
}

// calculateInfiltrationLoss computes annual heat loss due to air infiltration (kWh)
// volume: greenhouse volume in m3
// airSeal: seal quality ("tight", "average", "standard")
// hdd: heating degree days in Celsius-days
func (s *CalculationService) calculateInfiltrationLoss(volume float64, airSeal string, hdd float64) float64 {
    ach := s.getACH(airSeal)

    // Infiltration heat loss formula:
    // Q = V * ACH * rho * Cp * deltaT * time
    // For air: rho * Cp ≈ 1200 J/(m3·K) = 1200 Wh/(m3·K·h) / 3600
    // Simplified: Q (Wh) = V * ACH * 0.33 * HDD * 24
    // Where 0.33 = 1200/3600 (W·h per m3·K)
    return (volume * ach * 0.33 * hdd * 24) / 1000 // kWh
}

// getACH returns air changes per hour based on seal quality
func (s *CalculationService) getACH(airSeal string) float64 {
    ach := map[string]float64{
        "tight":    0.5,  // Professional air sealing, blower door tested
        "average":  1.5,  // Careful construction with sealed joints
        "standard": 3.0,  // Typical greenhouse construction
    }
    if val, ok := ach[airSeal]; ok {
        return val
    }
    return 1.5 // Default to average
}

// calculateFrostSkirtSavings estimates heat savings from frost skirt (kWh)
// perimeter: greenhouse perimeter in m
// skirt: frost skirt configuration
func (s *CalculationService) calculateFrostSkirtSavings(perimeter float64, skirt FrostSkirt) float64 {
    // Simplified calculation - actual would consider soil conductivity
    r := s.materialData.GetRValue(skirt.Insulation) // m2-K/W
    effectiveArea := perimeter * (skirt.Depth / 100) // Convert depth from cm to m

    // Estimate 10-20% reduction in perimeter heat loss
    // This is a simplified model; full calculation would use ground coupling factors
    return effectiveArea * r * 0.15
}

// getOptimizationReduction returns the fractional heat loss reduction for an optimization type
func (s *CalculationService) getOptimizationReduction(optType string) float64 {
    reductions := map[string]float64{
        "gaht":          0.30, // Ground to Air Heat Transfer - 30% reduction
        "night_blanket": 0.25, // Thermal blanket - 25% reduction
        "row_cover":     0.15, // Row covers - 15% reduction
        "thermal_mass":  0.20, // Water barrels/mass - 20% reduction
    }
    if val, ok := reductions[optType]; ok {
        return val
    }
    return 0
}

// getCropHeatingFactor returns a multiplier for heating needs based on crop type
func (s *CalculationService) getCropHeatingFactor(profile CropProfile) float64 {
    // Different crops need different temperature maintenance
    factors := map[string]float64{
        "cold-hardy":  0.7,  // Lower heating needs (e.g., kale, spinach)
        "cool-season": 0.85, // Moderate heating (e.g., lettuce, peas)
        "warm-season": 1.0,  // Standard heating (e.g., tomatoes, peppers)
        "tropical":    1.3,  // Higher heating needs (e.g., basil, cucumbers)
    }
    if val, ok := factors[profile.Season]; ok {
        return val
    }
    return 1.0
}

// getCropTargetTemp returns the target growing temperature in Celsius
func (s *CalculationService) getCropTargetTemp(profile CropProfile) float64 {
    // Target temperatures in Celsius for heating calculations
    targets := map[string]float64{
        "cold-hardy":  2,   // Just above freezing (2C / 35F)
        "cool-season": 10,  // Cool but not cold (10C / 50F)
        "warm-season": 18,  // Warm growing conditions (18C / 65F)
        "tropical":    24,  // Tropical warmth (24C / 75F)
    }
    if val, ok := targets[profile.Season]; ok {
        return val
    }
    return 18 // Default to warm-season
}

// calculateTotalSurfaceArea returns the total envelope surface area in m2
func (s *CalculationService) calculateTotalSurfaceArea(state *WizardState) float64 {
    areas := s.calculateSurfaceArea(state)
    return areas.WallNorth + areas.WallSouth + areas.WallEast + areas.WallWest +
        areas.RoofNorth + areas.RoofSouth
}

// calculateAverageRValue computes the area-weighted average R-value (m2-K/W)
func (s *CalculationService) calculateAverageRValue(state *WizardState) float64 {
    areas := s.calculateSurfaceArea(state)
    ins := state.Insulation

    totalArea := s.calculateTotalSurfaceArea(state)
    if totalArea == 0 {
        return 1.0 // Prevent division by zero
    }

    // Area-weighted R-value calculation
    weightedR := areas.WallNorth*s.materialData.GetRValue(ins.WallNorth) +
        areas.WallSouth*s.materialData.GetRValue(ins.WallSouth) +
        areas.WallEast*s.materialData.GetRValue(ins.WallEast) +
        areas.WallWest*s.materialData.GetRValue(ins.WallWest) +
        areas.RoofNorth*s.materialData.GetRValue(ins.RoofNorth) +
        areas.RoofSouth*s.materialData.GetRValue(ins.RoofSouth)

    return weightedR / totalArea
}
```

## Energy Cost Service

Calculates annual heating costs based on energy demand and farm energy rates.

```go
package services

import "context"

type EnergyCostService struct {
    db          *Database
    calculation *CalculationService
}

func NewEnergyCostService(db *Database, calc *CalculationService) *EnergyCostService {
    return &EnergyCostService{db: db, calculation: calc}
}

// EnergyCost represents a fuel type and its associated cost
type EnergyCost struct {
    FuelType string  // electricity, natural_gas, propane, wood, pellets, etc.
    Unit     string  // kWh, m3, L, cord, kg
    Rate     float64 // Cost per unit in CAD
}

// HeatingCostResult contains the calculated heating costs
type HeatingCostResult struct {
    PeakHeatNeed   float64          // Peak heating power (W)
    AnnualEnergy   float64          // Annual energy demand (kWh)
    SourceCosts    []SourceCostItem // Cost breakdown by heating source
    TotalAnnualCost float64         // Total estimated annual cost (CAD)
}

type SourceCostItem struct {
    FuelType   string  // Fuel type name
    Percentage float64 // Percentage of heating from this source
    EnergyKWh  float64 // Energy from this source (kWh)
    UnitCost   float64 // Cost per unit
    TotalCost  float64 // Total cost for this source (CAD)
}

// CalculateHeatingCosts computes the annual heating costs for a greenhouse project
func (s *EnergyCostService) CalculateHeatingCosts(ctx context.Context, state *WizardState) (*HeatingCostResult, error) {
    // Get peak heat and annual energy from calculation service
    peakHeat, annualEnergy := s.calculation.CalculateHeating(state)

    // Get farm's energy costs
    farm, err := s.db.GetFarm(ctx, state.FarmID)
    if err != nil {
        return nil, err
    }

    result := &HeatingCostResult{
        PeakHeatNeed: peakHeat,
        AnnualEnergy: annualEnergy,
    }

    // Calculate cost for each heating source
    for _, source := range state.HeatingSources {
        energyFromSource := annualEnergy * (source.Percentage / 100)

        // Find the energy cost for this fuel type
        var unitCost float64
        for _, ec := range farm.EnergyCosts {
            if ec.FuelType == source.FuelType {
                unitCost = ec.Rate
                break
            }
        }

        // Convert kWh to fuel units and calculate cost
        fuelUnits := s.convertKWhToFuelUnits(energyFromSource, source.FuelType)
        totalCost := fuelUnits * unitCost

        result.SourceCosts = append(result.SourceCosts, SourceCostItem{
            FuelType:   source.FuelType,
            Percentage: source.Percentage,
            EnergyKWh:  energyFromSource,
            UnitCost:   unitCost,
            TotalCost:  totalCost,
        })
        result.TotalAnnualCost += totalCost
    }

    return result, nil
}

// convertKWhToFuelUnits converts kWh to the appropriate fuel unit
// Returns the quantity of fuel needed to produce the given kWh of heat
func (s *EnergyCostService) convertKWhToFuelUnits(kWh float64, fuelType string) float64 {
    // Efficiency factors and energy content by fuel type
    // These account for typical heating equipment efficiency
    switch fuelType {
    case "electricity":
        // Electric resistance: ~100% efficient, 1 kWh = 1 kWh
        // Heat pump: COP of 2.5-4.0, but we use direct kWh for billing
        return kWh

    case "natural_gas":
        // 1 m3 natural gas ≈ 10.5 kWh gross
        // Furnace efficiency: ~95% for high-efficiency
        // Effective: 10.5 * 0.95 ≈ 10 kWh per m3
        return kWh / 10.0

    case "propane":
        // 1 L propane ≈ 7.1 kWh gross
        // Furnace efficiency: ~92%
        // Effective: 7.1 * 0.92 ≈ 6.5 kWh per L
        return kWh / 6.5

    case "wood":
        // 1 cord of hardwood ≈ 6,400 kWh gross
        // Wood stove efficiency: ~60-80%, use 70%
        // Effective: 6400 * 0.70 ≈ 4480 kWh per cord
        return kWh / 4480

    case "pellets":
        // 1 kg wood pellets ≈ 4.8 kWh gross
        // Pellet stove efficiency: ~85%
        // Effective: 4.8 * 0.85 ≈ 4.1 kWh per kg
        return kWh / 4.1

    case "heating_oil":
        // 1 L heating oil ≈ 10.3 kWh gross
        // Furnace efficiency: ~85%
        // Effective: 10.3 * 0.85 ≈ 8.8 kWh per L
        return kWh / 8.8

    default:
        // Unknown fuel type, assume 1:1 ratio
        return kWh
    }
}

// GetDefaultEnergyCosts returns typical energy costs for a province
func (s *EnergyCostService) GetDefaultEnergyCosts(ctx context.Context, provinceCode string) []EnergyCost {
    // Default costs based on Canadian provincial averages (CAD, 2024)
    defaults := map[string][]EnergyCost{
        "ON": { // Ontario
            {FuelType: "electricity", Unit: "kWh", Rate: 0.13},
            {FuelType: "natural_gas", Unit: "m3", Rate: 0.35},
            {FuelType: "propane", Unit: "L", Rate: 0.85},
            {FuelType: "wood", Unit: "cord", Rate: 350},
            {FuelType: "pellets", Unit: "kg", Rate: 0.45},
        },
        "QC": { // Quebec
            {FuelType: "electricity", Unit: "kWh", Rate: 0.07},
            {FuelType: "natural_gas", Unit: "m3", Rate: 0.55},
            {FuelType: "propane", Unit: "L", Rate: 0.90},
            {FuelType: "wood", Unit: "cord", Rate: 300},
            {FuelType: "pellets", Unit: "kg", Rate: 0.48},
        },
        "BC": { // British Columbia
            {FuelType: "electricity", Unit: "kWh", Rate: 0.10},
            {FuelType: "natural_gas", Unit: "m3", Rate: 0.40},
            {FuelType: "propane", Unit: "L", Rate: 0.80},
            {FuelType: "wood", Unit: "cord", Rate: 280},
            {FuelType: "pellets", Unit: "kg", Rate: 0.42},
        },
        "AB": { // Alberta
            {FuelType: "electricity", Unit: "kWh", Rate: 0.15},
            {FuelType: "natural_gas", Unit: "m3", Rate: 0.25},
            {FuelType: "propane", Unit: "L", Rate: 0.75},
            {FuelType: "wood", Unit: "cord", Rate: 320},
            {FuelType: "pellets", Unit: "kg", Rate: 0.44},
        },
    }

    if costs, ok := defaults[provinceCode]; ok {
        return costs
    }

    // Return Ontario defaults if province not found
    return defaults["ON"]
}
```

## Location Service

Handles cascading location dropdowns for farm creation.

```go
package services

type LocationService struct {
    db *Database
}

func NewLocationService(db *Database) *LocationService {
    return &LocationService{db: db}
}

// GetCountries returns all available countries
func (s *LocationService) GetCountries(ctx context.Context) ([]Country, error) {
    return s.db.GetCountries(ctx)
}

// GetProvinces returns provinces/states for a given country
func (s *LocationService) GetProvinces(ctx context.Context, countryCode string) ([]Province, error) {
    return s.db.GetProvincesByCountry(ctx, countryCode)
}

// GetCities returns cities for a given province
func (s *LocationService) GetCities(ctx context.Context, provinceCode string) ([]City, error) {
    return s.db.GetCitiesByProvince(ctx, provinceCode)
}

// ValidateLocation checks if a location combination is valid
func (s *LocationService) ValidateLocation(ctx context.Context, loc Location) error {
    // Verify country exists
    country, err := s.db.GetCountry(ctx, loc.Country)
    if err != nil {
        return fmt.Errorf("invalid country: %s", loc.Country)
    }
    
    // Verify province exists in country
    province, err := s.db.GetProvince(ctx, loc.Province)
    if err != nil || province.CountryCode != country.Code {
        return fmt.Errorf("invalid province: %s", loc.Province)
    }
    
    // Verify city exists in province
    city, err := s.db.GetCity(ctx, loc.City)
    if err != nil || city.ProvinceCode != province.Code {
        return fmt.Errorf("invalid city: %s", loc.City)
    }
    
    return nil
}

type Country struct {
    Code string
    Name string
}

type Province struct {
    Code        string
    CountryCode string
    Name        string
}

type City struct {
    Code         string
    ProvinceCode string
    Name         string
    Latitude     float64
    Longitude    float64
}
```

## Co-Owner Service

Manages farm co-ownership operations.

```go
package services

type CoOwnerService struct {
    db    *Database
    email *EmailService
}

func NewCoOwnerService(db *Database, email *EmailService) *CoOwnerService {
    return &CoOwnerService{db: db, email: email}
}

// AddCoOwner adds a co-owner to a farm
func (s *CoOwnerService) AddCoOwner(ctx context.Context, farmID string, ownerID string, coOwnerEmail string) error {
    // Verify requester is the farm owner
    farm, err := s.db.GetFarm(ctx, farmID)
    if err != nil {
        return err
    }
    
    if farm.OwnerID != ownerID {
        return ErrNotAuthorized
    }
    
    // Find user by email
    coOwner, err := s.db.GetUserByEmail(ctx, coOwnerEmail)
    if err != nil {
        return fmt.Errorf("user not found: %s", coOwnerEmail)
    }
    
    // Check if already a co-owner
    if contains(farm.CoOwnerIDs, coOwner.ID) {
        return ErrAlreadyCoOwner
    }
    
    // Add co-owner
    farm.CoOwnerIDs = append(farm.CoOwnerIDs, coOwner.ID)
    farm.UpdatedAt = time.Now()
    
    if err := s.db.UpdateFarm(ctx, farm); err != nil {
        return err
    }
    
    // Send notification email
    s.email.SendCoOwnerInvite(coOwner.Email, farm.Name, farm.ID)
    
    return nil
}

// RemoveCoOwner removes a co-owner from a farm
func (s *CoOwnerService) RemoveCoOwner(ctx context.Context, farmID string, ownerID string, coOwnerID string) error {
    farm, err := s.db.GetFarm(ctx, farmID)
    if err != nil {
        return err
    }
    
    if farm.OwnerID != ownerID {
        return ErrNotAuthorized
    }
    
    // Remove co-owner
    farm.CoOwnerIDs = remove(farm.CoOwnerIDs, coOwnerID)
    farm.UpdatedAt = time.Now()
    
    return s.db.UpdateFarm(ctx, farm)
}

// ListCoOwners returns all co-owners for a farm
func (s *CoOwnerService) ListCoOwners(ctx context.Context, farmID string) ([]User, error) {
    farm, err := s.db.GetFarm(ctx, farmID)
    if err != nil {
        return nil, err
    }
    
    var coOwners []User
    for _, id := range farm.CoOwnerIDs {
        user, err := s.db.GetUser(ctx, id)
        if err == nil {
            coOwners = append(coOwners, *user)
        }
    }
    
    return coOwners, nil
}

var (
    ErrNotAuthorized  = errors.New("not authorized")
    ErrAlreadyCoOwner = errors.New("user is already a co-owner")
)
```
