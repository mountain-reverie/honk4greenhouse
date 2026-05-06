# Handlers

## Showcase Handler

```go
package handlers

import (
    "net/http"
)

type ShowcaseHandler struct {
    db        *Database
    templates *TemplateRenderer
}

func (h *ShowcaseHandler) Index(w http.ResponseWriter, r *http.Request) {
    greenhouses, err := h.db.GetPublishedGreenhouses(r.Context())
    if err != nil {
        http.Error(w, "Failed to load greenhouses", http.StatusInternalServerError)
        return
    }
    
    user := GetUserFromContext(r.Context()) // nil if not authenticated
    
    // Check if HTMX request (partial render)
    if r.Header.Get("HX-Request") == "true" {
        h.templates.Render(w, "greenhouse/card_list", greenhouse.CardListProps{
            Greenhouses: greenhouses,
        })
        return
    }
    
    // Full page render
    h.templates.Render(w, "pages/showcase", ShowcasePageProps{
        User:        user,
        ActiveTab:   "showcase",
        Greenhouses: greenhouses,
    })
}
```

## Climate Handler

```go
package handlers

type ClimateHandler struct {
    db      *Database
    climate *ClimateService
}

func (h *ClimateHandler) Get(w http.ResponseWriter, r *http.Request) {
    farmID := chi.URLParam(r, "farmID")
    view := r.URL.Query().Get("view") // "chart" or "table"
    
    cv, err := h.climate.GetCurrentVersion(r.Context(), farmID)
    if err != nil {
        http.Error(w, "Climate data not found", http.StatusNotFound)
        return
    }
    
    if view == "table" {
        h.templates.Render(w, "charts/monthly_table", charts.MonthlyTableProps{
            ID:   farmID,
            Data: cv.Data,
        })
    } else {
        h.templates.Render(w, "charts/climate_chart", charts.ClimateChartProps{
            ID:   farmID,
            Data: cv.Data,
        })
    }
}

func (h *ClimateHandler) ToggleLock(w http.ResponseWriter, r *http.Request) {
    farmID := chi.URLParam(r, "farmID")
    user := GetUserFromContext(r.Context())
    
    // Verify ownership
    farm, err := h.db.GetFarm(r.Context(), farmID)
    if err != nil || farm.OwnerID != user.ID {
        http.Error(w, "Forbidden", http.StatusForbidden)
        return
    }
    
    // Get current lock state from session/cookie
    isLocked := h.getEditLockState(r, farmID)
    newState := !isLocked
    h.setEditLockState(w, farmID, newState)
    
    cv, _ := h.climate.GetCurrentVersion(r.Context(), farmID)
    
    h.templates.Render(w, "charts/editable_table_body", charts.EditableTableBodyProps{
        FarmID:   farmID,
        Data:     cv.Data,
        Editable: !newState,
    })
}

func (h *ClimateHandler) Save(w http.ResponseWriter, r *http.Request) {
    farmID := chi.URLParam(r, "farmID")
    user := GetUserFromContext(r.Context())
    
    // Parse form data
    if err := r.ParseForm(); err != nil {
        http.Error(w, "Invalid form data", http.StatusBadRequest)
        return
    }
    
    data := parseClimateFormData(r.Form)
    
    cv, err := h.climate.SaveVersion(r.Context(), farmID, user.ID, data)
    if err != nil {
        http.Error(w, "Failed to save", http.StatusInternalServerError)
        return
    }
    
    // Lock after save
    h.setEditLockState(w, farmID, true)
    
    h.templates.Render(w, "charts/editable_table", charts.EditableTableProps{
        FarmID:   farmID,
        Data:     cv.Data,
        Version:  cv.Version,
        IsOwner:  true,
        IsLocked: true,
    })
}

func (h *ClimateHandler) Versions(w http.ResponseWriter, r *http.Request) {
    farmID := chi.URLParam(r, "farmID")
    
    versions, err := h.climate.ListVersions(r.Context(), farmID)
    if err != nil {
        http.Error(w, "Failed to load versions", http.StatusInternalServerError)
        return
    }
    
    current, _ := h.climate.GetCurrentVersion(r.Context(), farmID)
    
    h.templates.Render(w, "charts/version_list", charts.VersionListProps{
        FarmID:         farmID,
        Versions:       versions,
        CurrentVersion: current.Version,
    })
}
```

## Publication Handler

```go
package handlers

type PublicationHandler struct {
    db          *Database
    publication *PublicationService
}

func (h *PublicationHandler) GetStatus(w http.ResponseWriter, r *http.Request) {
    projectID := chi.URLParam(r, "projectID")
    
    req, err := h.db.GetActivePublicationRequest(r.Context(), projectID)
    if err != nil {
        // No active request
        project, _ := h.db.GetProject(r.Context(), projectID)
        photos, _ := h.db.GetProjectPhotos(r.Context(), projectID)
        
        h.templates.Render(w, "publication/request_button", publication.RequestButtonProps{
            ProjectID:         projectID,
            CanRequest:        true,
            HasRequiredPhotos: len(photos) >= 1,
        })
        return
    }
    
    h.templates.Render(w, "publication/approval_status", publication.ApprovalStatusProps{
        Request: req,
    })
}

func (h *PublicationHandler) Request(w http.ResponseWriter, r *http.Request) {
    projectID := chi.URLParam(r, "projectID")
    user := GetUserFromContext(r.Context())
    
    req, err := h.publication.RequestPublication(r.Context(), projectID, user.ID)
    if err != nil {
        http.Error(w, "Failed to request publication", http.StatusInternalServerError)
        return
    }
    
    h.templates.Render(w, "publication/approval_status", publication.ApprovalStatusProps{
        Request: req,
    })
}

func (h *PublicationHandler) Retrigger(w http.ResponseWriter, r *http.Request) {
    projectID := chi.URLParam(r, "projectID")
    
    req, err := h.db.GetActivePublicationRequest(r.Context(), projectID)
    if err != nil {
        http.Error(w, "No active request", http.StatusNotFound)
        return
    }
    
    if err := h.publication.Retrigger(r.Context(), req.ID); err != nil {
        http.Error(w, "Failed to retrigger", http.StatusInternalServerError)
        return
    }
    
    req, _ = h.db.GetPublicationRequest(r.Context(), req.ID)
    h.templates.Render(w, "publication/approval_status", publication.ApprovalStatusProps{
        Request: req,
    })
}

// Email link handlers (public, token-based)

func (h *PublicationHandler) ApprovePage(w http.ResponseWriter, r *http.Request) {
    token := r.URL.Query().Get("token")
    
    req, participant, err := h.db.GetPublicationByToken(r.Context(), token)
    if err != nil {
        http.Error(w, "Invalid or expired token", http.StatusNotFound)
        return
    }
    
    project, _ := h.db.GetProject(r.Context(), req.ProjectID)
    
    h.templates.Render(w, "publication/approval_card", publication.ApprovalCardProps{
        Request:     req,
        Project:     project,
        Participant: participant,
        Token:       token,
    })
}

func (h *PublicationHandler) ApproveAPI(w http.ResponseWriter, r *http.Request) {
    token := r.URL.Query().Get("token")
    
    if err := h.publication.Approve(r.Context(), token); err != nil {
        http.Error(w, "Failed to approve", http.StatusInternalServerError)
        return
    }
    
    // Render success page
    h.templates.Render(w, "publication/approved_confirmation", nil)
}

func (h *PublicationHandler) DenyPage(w http.ResponseWriter, r *http.Request) {
    token := r.URL.Query().Get("token")
    
    req, _, err := h.db.GetPublicationByToken(r.Context(), token)
    if err != nil {
        http.Error(w, "Invalid token", http.StatusNotFound)
        return
    }
    
    h.templates.Render(w, "publication/deny_form", publication.DenyFormProps{
        Request: req,
        Token:   token,
    })
}

func (h *PublicationHandler) DenyAPI(w http.ResponseWriter, r *http.Request) {
    token := r.URL.Query().Get("token")
    reason := r.FormValue("reason")
    
    if err := h.publication.Deny(r.Context(), token, reason); err != nil {
        http.Error(w, "Failed to deny", http.StatusInternalServerError)
        return
    }
    
    h.templates.Render(w, "publication/denied_confirmation", nil)
}
```

## Quote Handler

```go
package handlers

type QuoteHandler struct {
    db      *Database
    storage *S3Storage
    email   *EmailService
}

func (h *QuoteHandler) Create(w http.ResponseWriter, r *http.Request) {
    projectID := chi.URLParam(r, "projectID")
    user := GetUserFromContext(r.Context())
    
    // Parse multipart form
    if err := r.ParseMultipartForm(10 << 20); err != nil { // 10MB max
        http.Error(w, "Form too large", http.StatusBadRequest)
        return
    }
    
    quote := &models.Quote{
        ID:          generateID(),
        ProjectID:   projectID,
        BuilderID:   user.ID,
        BuilderName: r.FormValue("builderName"),
        Price:       parseFloat(r.FormValue("price")),
        Duration:    r.FormValue("duration"),
        NeedsReview: false,
        CreatedAt:   time.Now(),
        UpdatedAt:   time.Now(),
    }
    
    // Handle PDF upload
    file, header, err := r.FormFile("pdf")
    if err == nil {
        defer file.Close()
        
        url, err := h.storage.UploadQuotePDF(r.Context(), projectID, quote.ID, file)
        if err != nil {
            http.Error(w, "Failed to upload PDF", http.StatusInternalServerError)
            return
        }
        quote.PDFURL = url
    }
    
    if err := h.db.CreateQuote(r.Context(), quote); err != nil {
        http.Error(w, "Failed to save quote", http.StatusInternalServerError)
        return
    }
    
    // Notify project owner
    project, _ := h.db.GetProject(r.Context(), projectID)
    owner, _ := h.db.GetUser(r.Context(), project.OwnerID)
    
    h.email.SendQuoteSubmitted(r.Context(), owner.Email, notifications.QuoteNotificationData{
        RecipientName: owner.Name,
        BuilderName:   quote.BuilderName,
        ProjectName:   project.Name,
        Price:         formatMoney(quote.Price),
        ViewURL:       baseURL + "/projects/" + projectID,
    })
    
    h.templates.Render(w, "quotes/quote_row", quotes.QuoteRowProps{
        Quote: quote,
    })
}

func (h *QuoteHandler) MarkNeedsReview(ctx context.Context, projectID string) error {
    return h.db.MarkQuotesNeedReview(ctx, projectID)
}
```

## Photo Handler

```go
package handlers

type PhotoHandler struct {
    db      *Database
    storage *S3Storage
}

func (h *PhotoHandler) Upload(w http.ResponseWriter, r *http.Request) {
    projectID := chi.URLParam(r, "projectID")
    
    file, header, err := r.FormFile("file")
    if err != nil {
        http.Error(w, "No file provided", http.StatusBadRequest)
        return
    }
    defer file.Close()
    
    // Validate image type
    contentType := header.Header.Get("Content-Type")
    if !strings.HasPrefix(contentType, "image/") {
        http.Error(w, "Invalid file type", http.StatusBadRequest)
        return
    }
    
    photoID := generateID()
    
    url, err := h.storage.UploadProjectPhoto(r.Context(), projectID, photoID, file, contentType)
    if err != nil {
        http.Error(w, "Failed to upload", http.StatusInternalServerError)
        return
    }
    
    photo := &models.Photo{
        ID:           photoID,
        ProjectID:    projectID,
        URL:          url,
        ThumbnailURL: url, // Could generate thumbnail
        CreatedAt:    time.Now(),
    }
    
    if err := h.db.CreatePhoto(r.Context(), photo); err != nil {
        http.Error(w, "Failed to save", http.StatusInternalServerError)
        return
    }
    
    // Return photo thumbnail HTML
    h.templates.Render(w, "publication/photo_thumbnail", publication.PhotoThumbnailProps{
        ProjectID: projectID,
        Photo:     photo,
    })
}

func (h *PhotoHandler) Delete(w http.ResponseWriter, r *http.Request) {
    projectID := chi.URLParam(r, "projectID")
    photoID := chi.URLParam(r, "photoID")
    
    photo, err := h.db.GetPhoto(r.Context(), photoID)
    if err != nil || photo.ProjectID != projectID {
        http.Error(w, "Photo not found", http.StatusNotFound)
        return
    }
    
    // Delete from S3
    key := fmt.Sprintf("projects/%s/photos/%s", projectID, photoID)
    h.storage.DeletePhoto(r.Context(), key)
    
    // Delete from DB
    h.db.DeletePhoto(r.Context(), photoID)
    
    w.WriteHeader(http.StatusOK)
}
```

## Middleware

```go
package middleware

func RequireAuth(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        user := GetUserFromSession(r)
        if user == nil {
            if r.Header.Get("HX-Request") == "true" {
                w.Header().Set("HX-Redirect", "/login")
                w.WriteHeader(http.StatusUnauthorized)
                return
            }
            http.Redirect(w, r, "/login?redirect="+r.URL.Path, http.StatusFound)
            return
        }
        
        ctx := context.WithValue(r.Context(), userContextKey, user)
        next.ServeHTTP(w, r.WithContext(ctx))
    })
}

func RequireRole(roles ...models.Role) func(http.Handler) http.Handler {
    return func(next http.Handler) http.Handler {
        return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
            user := GetUserFromContext(r.Context())
            if user == nil {
                http.Error(w, "Unauthorized", http.StatusUnauthorized)
                return
            }
            
            for _, role := range roles {
                if user.Role == role {
                    next.ServeHTTP(w, r)
                    return
                }
            }
            
            http.Error(w, "Forbidden", http.StatusForbidden)
        })
    }
}

func RequireProjectAccess(db *Database) func(http.Handler) http.Handler {
    return func(next http.Handler) http.Handler {
        return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
            projectID := chi.URLParam(r, "projectID")
            user := GetUserFromContext(r.Context())
            
            hasAccess, err := db.UserHasProjectAccess(r.Context(), user.ID, projectID)
            if err != nil || !hasAccess {
                http.Error(w, "Forbidden", http.StatusForbidden)
                return
            }
            
            next.ServeHTTP(w, r)
        })
    }
}
```

## Farm Handler

```go
package handlers

type FarmHandler struct {
    db        *Database
    templates *TemplateRenderer
}

func (h *FarmHandler) NewForm(w http.ResponseWriter, r *http.Request) {
    user := GetUserFromContext(r.Context())
    
    h.templates.Render(w, "farm/form", farm.FormProps{
        Farm: nil, // New farm
        User: user,
    })
}

func (h *FarmHandler) Create(w http.ResponseWriter, r *http.Request) {
    user := GetUserFromContext(r.Context())
    
    if err := r.ParseForm(); err != nil {
        http.Error(w, "Invalid form", http.StatusBadRequest)
        return
    }
    
    farm := &models.Farm{
        ID:        generateID(),
        OwnerID:   user.ID,
        Name:      r.FormValue("name"),
        Location: models.Location{
            Country:  r.FormValue("country"),
            Province: r.FormValue("province"),
            City:     r.FormValue("city"),
        },
        CoOwnerIDs:  []string{},
        EnergyCosts: []models.EnergyCost{},
        CreatedAt:   time.Now(),
        UpdatedAt:   time.Now(),
    }
    
    // Parse co-owners
    for _, email := range r.Form["coOwners"] {
        coOwner, err := h.db.GetUserByEmail(r.Context(), email)
        if err == nil {
            farm.CoOwnerIDs = append(farm.CoOwnerIDs, coOwner.ID)
        }
    }
    
    // Parse energy costs
    fuelTypes := r.Form["fuelType[]"]
    costs := r.Form["energyCost[]"]
    for i := range fuelTypes {
        cost, _ := strconv.ParseFloat(costs[i], 64)
        farm.EnergyCosts = append(farm.EnergyCosts, models.EnergyCost{
            ID:       generateID(),
            FarmID:   farm.ID,
            FuelType: fuelTypes[i],
            Unit:     getUnitForFuelType(fuelTypes[i]),
            Cost:     cost,
        })
    }
    
    if err := h.db.CreateFarm(r.Context(), farm); err != nil {
        http.Error(w, "Failed to create farm", http.StatusInternalServerError)
        return
    }
    
    // Redirect to farm detail
    w.Header().Set("HX-Redirect", "/farms/"+farm.ID)
    w.WriteHeader(http.StatusOK)
}

func (h *FarmHandler) Detail(w http.ResponseWriter, r *http.Request) {
    farmID := chi.URLParam(r, "id")
    user := GetUserFromContext(r.Context())
    
    farm, err := h.db.GetFarm(r.Context(), farmID)
    if err != nil {
        http.Error(w, "Farm not found", http.StatusNotFound)
        return
    }
    
    // Check access
    if farm.OwnerID != user.ID && !contains(farm.CoOwnerIDs, user.ID) {
        http.Error(w, "Forbidden", http.StatusForbidden)
        return
    }
    
    projects, _ := h.db.GetProjectsByFarm(r.Context(), farmID)
    climate, _ := h.db.GetLatestClimateVersion(r.Context(), farmID)
    
    h.templates.Render(w, "pages/farm_detail", FarmDetailProps{
        Farm:        farm,
        Projects:    projects,
        ClimateData: climate,
        IsOwner:     farm.OwnerID == user.ID,
        User:        user,
    })
}
```

## Wizard Handler

```go
package handlers

type WizardHandler struct {
    db        *Database
    templates *TemplateRenderer
    calc      *CalculationService
}

type WizardState struct {
    FarmID       string
    Step         int
    ShapeID      string
    Dimensions   models.Dimensions
    Orientation  string
    CropProfile  models.CropProfile
    Insulation   models.InsulationConfig
    BTU          float64
    PeakHeat     float64
    AnnualEnergy float64
    Optimizations []models.Optimization
    HeatingSources []models.HeatingSource
}

func (h *WizardHandler) Start(w http.ResponseWriter, r *http.Request) {
    farmID := chi.URLParam(r, "id")
    
    shapes, _ := h.db.GetGreenhouseShapes(r.Context())
    
    h.templates.Render(w, "wizard/layout", wizard.LayoutProps{
        FarmID: farmID,
        Step:   1,
        Content: wizard.Step1ShapeProps{
            FarmID:     farmID,
            Shapes:     shapes,
            SelectedID: "",
        },
    })
}

func (h *WizardHandler) Step(w http.ResponseWriter, r *http.Request) {
    farmID := chi.URLParam(r, "id")
    stepNum, _ := strconv.Atoi(chi.URLParam(r, "n"))
    
    // Load state from session/form
    state := h.loadWizardState(r, farmID)
    
    switch stepNum {
    case 1:
        shapes, _ := h.db.GetGreenhouseShapes(r.Context())
        h.templates.Render(w, "wizard/step1_shape", wizard.Step1ShapeProps{
            FarmID:     farmID,
            Shapes:     shapes,
            SelectedID: state.ShapeID,
        })
    case 2:
        h.templates.Render(w, "wizard/step2_dimensions", wizard.Step2DimensionsProps{
            FarmID:      farmID,
            Dimensions:  state.Dimensions,
            Orientation: state.Orientation,
            CropProfile: state.CropProfile,
        })
    case 3:
        // Calculate BTU based on current state
        btu := h.calc.CalculateBTU(state)
        h.templates.Render(w, "wizard/step3_insulation", wizard.Step3InsulationProps{
            FarmID:        farmID,
            Insulation:    state.Insulation,
            CalculatedBTU: btu,
        })
    case 4:
        // Calculate heating needs
        peak, annual := h.calc.CalculateHeating(state)
        h.templates.Render(w, "wizard/step4_optimization", wizard.Step4OptimizationProps{
            FarmID:             farmID,
            PeakHeatNeed:       peak,
            AnnualEnergyDemand: annual,
            Optimizations:      state.Optimizations,
            HeatingSources:     state.HeatingSources,
        })
    }
}

func (h *WizardHandler) CalculateBTU(w http.ResponseWriter, r *http.Request) {
    if err := r.ParseForm(); err != nil {
        http.Error(w, "Invalid form", http.StatusBadRequest)
        return
    }
    
    state := h.parseWizardForm(r.Form)
    btu := h.calc.CalculateBTU(state)
    
    fmt.Fprintf(w, "%.0f BTU/year", btu)
}

func (h *WizardHandler) CalculateHeating(w http.ResponseWriter, r *http.Request) {
    if err := r.ParseForm(); err != nil {
        http.Error(w, "Invalid form", http.StatusBadRequest)
        return
    }
    
    state := h.parseWizardForm(r.Form)
    peak, annual := h.calc.CalculateHeating(state)
    
    // Return updated heating sliders with percentages balanced
    sources := h.balanceHeatingSources(r.Form)
    
    h.templates.Render(w, "wizard/heating_sliders", wizard.HeatingSlidersProps{
        Sources: sources,
    })
}

func (h *WizardHandler) GetCropProfile(w http.ResponseWriter, r *http.Request) {
    cropType := r.URL.Query().Get("cropType")
    
    profile, err := h.db.GetCropProfile(r.Context(), cropType)
    if err != nil {
        http.Error(w, "Crop not found", http.StatusNotFound)
        return
    }
    
    h.templates.Render(w, "wizard/crop_profile", wizard.CropProfileProps{
        Profile: profile,
    })
}

func (h *WizardHandler) RequestShape(w http.ResponseWriter, r *http.Request) {
    user := GetUserFromContext(r.Context())
    description := r.FormValue("description")
    
    request := &models.ShapeRequest{
        ID:          generateID(),
        UserID:      user.ID,
        Description: description,
        Status:      "pending",
        CreatedAt:   time.Now(),
    }
    
    if err := h.db.CreateShapeRequest(r.Context(), request); err != nil {
        http.Error(w, "Failed to submit", http.StatusInternalServerError)
        return
    }
    
    w.Write([]byte("Request submitted"))
}

func (h *WizardHandler) CreateProject(w http.ResponseWriter, r *http.Request) {
    farmID := chi.URLParam(r, "id")
    user := GetUserFromContext(r.Context())
    
    if err := r.ParseForm(); err != nil {
        http.Error(w, "Invalid form", http.StatusBadRequest)
        return
    }
    
    state := h.parseWizardForm(r.Form)
    
    project := &models.Project{
        ID:                 generateID(),
        FarmID:             farmID,
        Name:               generateProjectName(state),
        GreenhouseShape:    state.ShapeID,
        Dimensions:         state.Dimensions,
        Orientation:        state.Orientation,
        CropGoal:           state.CropProfile,
        Insulation:         state.Insulation,
        CalculatedBTU:      state.BTU,
        PeakHeatNeed:       state.PeakHeat,
        AnnualEnergyDemand: state.AnnualEnergy,
        Optimizations:      state.Optimizations,
        HeatingSources:     state.HeatingSources,
        CreatedAt:          time.Now(),
        UpdatedAt:          time.Now(),
    }
    
    if err := h.db.CreateProject(r.Context(), project); err != nil {
        http.Error(w, "Failed to create", http.StatusInternalServerError)
        return
    }
    
    // Redirect to farm detail
    w.Header().Set("HX-Redirect", "/farms/"+farmID)
    w.WriteHeader(http.StatusOK)
}

func (h *WizardHandler) loadWizardState(r *http.Request, farmID string) *WizardState {
    // Load from session or parse from form data
    state := &WizardState{FarmID: farmID}
    
    if r.Method == "GET" {
        // Parse query params for state restoration
        if shapeID := r.URL.Query().Get("shapeID"); shapeID != "" {
            state.ShapeID = shapeID
        }
    } else {
        // Parse from form
        r.ParseForm()
        state = h.parseWizardForm(r.Form)
    }
    
    return state
}

func (h *WizardHandler) parseWizardForm(form url.Values) *WizardState {
    state := &WizardState{}
    
    state.ShapeID = form.Get("shapeID")
    state.Dimensions.Length, _ = strconv.ParseFloat(form.Get("length"), 64)
    state.Dimensions.Width, _ = strconv.ParseFloat(form.Get("width"), 64)
    state.Dimensions.Height, _ = strconv.ParseFloat(form.Get("height"), 64)
    state.Orientation = form.Get("orientation")
    
    state.CropProfile.CropType = form.Get("cropType")
    
    state.Insulation.Glazing = form.Get("glazing")
    state.Insulation.WallNorth = form.Get("wallNorth")
    state.Insulation.WallSouth = form.Get("wallSouth")
    state.Insulation.WallEast = form.Get("wallEast")
    state.Insulation.WallWest = form.Get("wallWest")
    state.Insulation.RoofNorth = form.Get("roofNorth")
    state.Insulation.RoofSouth = form.Get("roofSouth")
    state.Insulation.AirSeal = form.Get("airSeal")
    
    state.Insulation.FrostSkirt.Depth, _ = strconv.ParseFloat(form.Get("frostSkirtDepth"), 64)
    state.Insulation.FrostSkirt.Width, _ = strconv.ParseFloat(form.Get("frostSkirtWidth"), 64)
    state.Insulation.FrostSkirt.Insulation = form.Get("frostSkirtInsulation")
    
    // Parse optimizations
    for _, opt := range form["optimizations"] {
        state.Optimizations = append(state.Optimizations, models.Optimization{Type: opt})
    }
    
    // Parse heating sources
    if primary := form.Get("heatingSource1"); primary != "" {
        pct, _ := strconv.ParseFloat(form.Get("heatingPercent1"), 64)
        state.HeatingSources = append(state.HeatingSources, models.HeatingSource{
            Priority:   1,
            FuelType:   primary,
            Percentage: pct,
        })
    }
    if secondary := form.Get("heatingSource2"); secondary != "" {
        pct, _ := strconv.ParseFloat(form.Get("heatingPercent2"), 64)
        state.HeatingSources = append(state.HeatingSources, models.HeatingSource{
            Priority:   2,
            FuelType:   secondary,
            Percentage: pct,
        })
    }
    
    return state
}
```
