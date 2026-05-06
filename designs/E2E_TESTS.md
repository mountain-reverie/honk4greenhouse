# E2E Test Examples

Using `github.com/mountain-reverie/playwright-ci-go`

## Test Setup

```go
package e2e

import (
    "testing"
    "time"
    
    playwright "github.com/mountain-reverie/playwright-ci-go"
)

const baseURL = "http://localhost:8080"

func login(page playwright.Page, email string) {
    page.Goto(baseURL + "/login")
    page.Fill("input[name='email']", email)
    page.Click("button[type='submit']")
    page.WaitForURL("**/personal")
}

func logout(page playwright.Page) {
    page.Click("[data-testid='user-menu']")
    page.Click("text=Logout")
    page.WaitForURL("**/")
}

func getApprovalToken(email string) string {
    // Query test database for token
    return "test-token-" + email
}

func uploadTestPhoto(page playwright.Page) {
    page.SetInputFiles("input[type='file'][accept='image/*']", "fixtures/test-photo.jpg")
    page.WaitForSelector("[data-testid='photo-thumbnail']")
}
```

## Showcase Tests

```go
func TestShowcasePage_NoAuthRequired(t *testing.T) {
    pw := playwright.New(t)
    defer pw.Close()

    page := pw.NewPage()
    page.Goto(baseURL + "/showcase")
    
    // Should see greenhouse cards without login
    page.WaitForSelector("[data-testid='greenhouse-card']")
    cards := page.Locator("[data-testid='greenhouse-card']")
    
    if cards.Count() == 0 {
        t.Error("Expected greenhouse cards on showcase page")
    }
}

func TestShowcasePage_ClickCardOpensDetail(t *testing.T) {
    pw := playwright.New(t)
    defer pw.Close()

    page := pw.NewPage()
    page.Goto(baseURL + "/showcase")
    
    page.WaitForSelector("[data-testid='greenhouse-card']")
    page.Locator("[data-testid='greenhouse-card']").First().Click()
    
    page.WaitForURL("**/greenhouse/*")
    page.WaitForSelector("[data-testid='greenhouse-detail']")
}

func TestSharedTab_RequiresAuth(t *testing.T) {
    pw := playwright.New(t)
    defer pw.Close()

    page := pw.NewPage()
    page.Goto(baseURL + "/shared")
    
    // Should redirect to login
    page.WaitForURL("**/login**")
}
```

## Climate Data Tests

```go
func TestClimateChart_ToggleToTable(t *testing.T) {
    pw := playwright.New(t)
    defer pw.Close()

    page := pw.NewPage()
    page.Goto(baseURL + "/greenhouse/test-id")
    
    // Default should show chart
    page.WaitForSelector("[data-testid='climate-chart']")
    
    // Click table toggle
    page.Click("button:has-text('Table')")
    
    // Should now show table
    page.WaitForSelector("[data-testid='climate-table']")
    
    // Verify 12 months
    rows := page.Locator("[data-testid='climate-table'] tbody tr")
    if rows.Count() != 12 {
        t.Errorf("Expected 12 months, got %d", rows.Count())
    }
}

func TestClimateData_LockedByDefault(t *testing.T) {
    pw := playwright.New(t)
    defer pw.Close()

    page := pw.NewPage()
    login(page, "owner@test.com")
    page.Goto(baseURL + "/farms/test-farm")
    
    // Lock should show locked icon
    lockBtn := page.Locator("[data-testid='lock-toggle']")
    expect(lockBtn).ToHaveText("🔒")
    
    // Inputs should not exist (not in edit mode)
    expect(page.Locator("input[name='avgTemp_Jan']")).ToHaveCount(0)
}

func TestClimateData_UnlockAndEdit(t *testing.T) {
    pw := playwright.New(t)
    defer pw.Close()

    page := pw.NewPage()
    login(page, "owner@test.com")
    page.Goto(baseURL + "/farms/test-farm")
    
    // Unlock
    page.Click("[data-testid='lock-toggle']")
    page.WaitForSelector("[data-testid='lock-toggle']:has-text('🔓')")
    
    // Now can edit
    page.Fill("input[name='avgTemp_Jan']", "-15")
    
    // Save creates new version
    page.Click("button:has-text('Save as New Version')")
    
    // Check version selector shows today
    today := time.Now().Format("2006-01-02")
    expect(page.Locator("[data-testid='version-selector']")).ToContainText(today)
}

func TestClimateData_VersionHistory(t *testing.T) {
    pw := playwright.New(t)
    defer pw.Close()

    page := pw.NewPage()
    login(page, "owner@test.com")
    page.Goto(baseURL + "/farms/farm-with-versions")
    
    // Open version selector
    page.Click("[data-testid='version-selector']")
    
    // Should show multiple versions
    versions := page.Locator("[data-testid='version-item']")
    if versions.Count() < 2 {
        t.Error("Expected multiple versions in history")
    }
    
    // Click older version
    versions.Last().Click()
    
    // Data should update
    page.WaitForResponse("**/climate/**")
}
```

## Publication Approval Tests

```go
func TestPublication_RequiresPhotos(t *testing.T) {
    pw := playwright.New(t)
    defer pw.Close()

    page := pw.NewPage()
    login(page, "owner@test.com")
    page.Goto(baseURL + "/projects/project-no-photos")
    
    // Request button should be disabled
    btn := page.Locator("button:has-text('Photos Required')")
    expect(btn).ToBeDisabled()
}

func TestPublication_RequestFlow(t *testing.T) {
    pw := playwright.New(t)
    defer pw.Close()

    page := pw.NewPage()
    login(page, "owner@test.com")
    page.Goto(baseURL + "/projects/collaborative-project")
    
    // Upload required photo
    uploadTestPhoto(page)
    
    // Request publication
    page.Click("button:has-text('Request Publication')")
    
    // Should show approval status
    page.WaitForSelector("[data-testid='approval-status']")
    
    // Owner auto-approved
    expect(page.Locator("text=✅ Approved")).ToBeVisible()
    
    // Others pending
    pending := page.Locator("text=⏳ Pending")
    if pending.Count() != 2 {
        t.Errorf("Expected 2 pending approvals, got %d", pending.Count())
    }
}

func TestPublication_ApproveViaEmailLink(t *testing.T) {
    pw := playwright.New(t)
    defer pw.Close()

    page := pw.NewPage()
    
    // Simulate clicking email link
    token := getApprovalToken("designer@test.com")
    page.Goto(baseURL + "/publication/approve?token=" + token)
    
    // Should show approval card
    page.WaitForSelector("[data-testid='approval-card']")
    
    // Click approve
    page.Click("button:has-text('Approve')")
    
    // Should confirm
    page.WaitForSelector("text=approved")
}

func TestPublication_DenyWithReason(t *testing.T) {
    pw := playwright.New(t)
    defer pw.Close()

    page := pw.NewPage()
    
    token := getApprovalToken("builder@test.com")
    page.Goto(baseURL + "/publication/deny?token=" + token)
    
    // Fill reason
    page.Fill("textarea[name='reason']", "Need to update specs first")
    page.Click("button:has-text('Deny')")
    
    // Owner should see denial
    login(page, "owner@test.com")
    page.Goto(baseURL + "/projects/collaborative-project")
    
    expect(page.Locator("text=❌ Denied")).ToBeVisible()
    expect(page.Locator("text=Need to update specs")).ToBeVisible()
}

func TestPublication_Retrigger(t *testing.T) {
    pw := playwright.New(t)
    defer pw.Close()

    page := pw.NewPage()
    login(page, "owner@test.com")
    page.Goto(baseURL + "/projects/denied-project")
    
    // Should see retrigger button
    page.Click("button:has-text('Re-send Requests')")
    
    // Pending count should reset
    pending := page.Locator("text=⏳ Pending")
    expect(pending.Count()).ToBe(2)
}
```

## Quote Tests

```go
func TestQuotes_SubmitWithPDF(t *testing.T) {
    pw := playwright.New(t)
    defer pw.Close()

    page := pw.NewPage()
    login(page, "builder@test.com")
    page.Goto(baseURL + "/projects/shared-project")
    
    // Open quote form
    page.Click("button:has-text('Add Quote')")
    page.WaitForSelector("[data-testid='quote-dialog']")
    
    // Fill form
    page.Fill("input[name='builderName']", "Test Builder")
    page.Fill("input[name='price']", "95000")
    page.Fill("input[name='duration']", "6 weeks")
    page.SetInputFiles("input[name='pdf']", "fixtures/test-quote.pdf")
    
    // Submit
    page.Click("button:has-text('Submit Quote')")
    
    // Should appear in table
    page.WaitForSelector("text=Test Builder")
    page.WaitForSelector("text=$95,000")
}

func TestQuotes_MarkedNeedsReviewAfterEdit(t *testing.T) {
    pw := playwright.New(t)
    defer pw.Close()

    page := pw.NewPage()
    
    // Setup: project with quote, shared with editor
    login(page, "editor@test.com")
    page.Goto(baseURL + "/projects/project-with-quote")
    
    // Verify no review badge initially
    expect(page.Locator("[data-testid='needs-review-badge']")).ToHaveCount(0)
    
    // Make an edit
    page.Fill("input[name='expectedRevenue']", "75000")
    page.Click("button:has-text('Save')")
    
    // Quote should now need review
    page.WaitForSelector("[data-testid='needs-review-badge']")
}
```

## Sharing Tests

```go
func TestShare_ViewOnly(t *testing.T) {
    pw := playwright.New(t)
    defer pw.Close()

    page := pw.NewPage()
    login(page, "owner@test.com")
    page.Goto(baseURL + "/projects/my-project")
    
    // Open share dialog
    page.Click("[data-testid='share-button']")
    page.WaitForSelector("[data-testid='share-dialog']")
    
    // Fill and submit
    page.Fill("input[name='email']", "viewer@test.com")
    page.Click("input[value='view']")
    page.Click("button:has-text('Share')")
    
    // Verify as viewer
    logout(page)
    login(page, "viewer@test.com")
    page.Goto(baseURL + "/shared")
    
    expect(page.Locator("text=my-project")).ToBeVisible()
    
    // Should not be able to edit
    page.Click("text=my-project")
    expect(page.Locator("input[name='expectedRevenue']")).ToBeDisabled()
}

func TestShare_EditAccess(t *testing.T) {
    pw := playwright.New(t)
    defer pw.Close()

    page := pw.NewPage()
    login(page, "owner@test.com")
    page.Goto(baseURL + "/projects/my-project")
    
    page.Click("[data-testid='share-button']")
    page.Fill("input[name='email']", "editor@test.com")
    page.Click("input[value='edit']")
    page.Click("button:has-text('Share')")
    
    // Verify as editor
    logout(page)
    login(page, "editor@test.com")
    page.Goto(baseURL + "/shared")
    page.Click("text=my-project")
    
    // Should be able to edit
    expect(page.Locator("input[name='expectedRevenue']")).ToBeEnabled()
}
```

## Mobile Tests

```go
func TestMobile_HamburgerMenu(t *testing.T) {
    pw := playwright.New(t)
    defer pw.Close()

    page := pw.NewPage()
    
    // Set mobile viewport
    page.SetViewportSize(375, 667)
    page.Goto(baseURL + "/")
    
    // Desktop nav should be hidden
    expect(page.Locator("[data-testid='desktop-nav']")).ToBeHidden()
    
    // Hamburger should be visible
    hamburger := page.Locator("[data-testid='mobile-menu-button']")
    expect(hamburger).ToBeVisible()
    
    // Open menu
    hamburger.Click()
    page.WaitForSelector("[data-testid='mobile-menu']")
    
    // Navigate
    page.Click("text=Showcase")
    page.WaitForURL("**/showcase")
}

func TestMobile_ResponsiveCardGrid(t *testing.T) {
    pw := playwright.New(t)
    defer pw.Close()

    page := pw.NewPage()
    
    // Mobile: 1 column
    page.SetViewportSize(375, 667)
    page.Goto(baseURL + "/showcase")
    page.WaitForSelector("[data-testid='greenhouse-card']")
    
    grid := page.Locator("[data-testid='card-grid']")
    // Check computed style shows single column
    
    // Tablet: 2 columns
    page.SetViewportSize(768, 1024)
    
    // Desktop: 3 columns
    page.SetViewportSize(1280, 800)
}
```

## Farm Creation Tests

```go
func TestAddFarm_Form(t *testing.T) {
    pw := playwright.New(t)
    defer pw.Close()

    page := pw.NewPage()
    login(page, "owner@test.com")
    
    page.Goto(baseURL + "/farms/new")
    
    // Fill basic info
    page.Fill("input[name='name']", "Fresh Pals Farm")
    
    // Cascading location selects
    page.SelectOption("select[name='country']", "CA")
    page.WaitForSelector("select[name='province'] option:not(:first-child)")
    page.SelectOption("select[name='province']", "AB")
    page.WaitForSelector("select[name='city'] option:not(:first-child)")
    page.SelectOption("select[name='city']", "Olds")
    
    // Add co-owner
    page.Fill("#co-owner-email", "partner@test.com")
    page.Click("button:has-text('Add')")
    page.WaitForSelector("#co-owner-list li:has-text('partner@test.com')")
    
    // Add energy cost
    page.Click("button:has-text('+ Add')")
    page.SelectOption("select[name='fuelType[]']", "electricity")
    page.Fill("input[name='energyCost[]']", "0.15")
    
    // Submit
    page.Click("button:has-text('Save Farm')")
    page.WaitForURL("**/farms/*")
}

func TestFarm_CoOwnerManagement(t *testing.T) {
    pw := playwright.New(t)
    defer pw.Close()

    page := pw.NewPage()
    login(page, "owner@test.com")
    page.Goto(baseURL + "/farms/test-farm/edit")
    
    // Add co-owner
    page.Fill("#co-owner-email", "newpartner@test.com")
    page.Click("button:has-text('Add')")
    
    // Should appear in list
    expect(page.Locator("#co-owner-list")).ToContainText("newpartner@test.com")
    
    // Remove co-owner
    page.Click("#co-owner-list li:has-text('newpartner@test.com') button")
    expect(page.Locator("#co-owner-list")).Not().ToContainText("newpartner@test.com")
}
```

## Project Creation Wizard Tests

```go
func TestWizard_Step1_SelectShape(t *testing.T) {
    pw := playwright.New(t)
    defer pw.Close()

    page := pw.NewPage()
    login(page, "owner@test.com")
    page.Goto(baseURL + "/farms/test-farm/projects/new")
    
    // Should show step 1 - shape selection
    expect(page.Locator("h2")).ToHaveText("Select Greenhouse Shape")
    
    // Select a shape
    page.Click("[data-testid='shape-card']:first-child")
    
    // Next button should enable
    expect(page.Locator("button:has-text('Next')")).ToBeEnabled()
    
    // Go to step 2
    page.Click("button:has-text('Next')")
    page.WaitForSelector("h2:has-text('Dimensions')")
}

func TestWizard_Step2_DimensionsAndCrop(t *testing.T) {
    pw := playwright.New(t)
    defer pw.Close()

    page := pw.NewPage()
    login(page, "owner@test.com")
    navigateToWizardStep(page, 2)
    
    // Fill dimensions
    page.Fill("input[name='length']", "100")
    page.Fill("input[name='width']", "10")
    page.Fill("input[name='height']", "4")
    
    // Select orientation (click South on compass)
    page.Click("button:has-text('S')")
    
    // Select crop - should load profile
    page.SelectOption("select[name='cropType']", "tomato")
    page.WaitForSelector("#crop-profile:has-text('February')")
    
    // Verify crop profile shows
    expect(page.Locator("#crop-profile")).ToContainText("10°C")
    expect(page.Locator("#crop-profile")).ToContainText("30°C")
    
    page.Click("button:has-text('Next')")
    page.WaitForSelector("h2:has-text('Insulation')")
}

func TestWizard_Step3_Insulation(t *testing.T) {
    pw := playwright.New(t)
    defer pw.Close()

    page := pw.NewPage()
    login(page, "owner@test.com")
    navigateToWizardStep(page, 3)
    
    // Should show calculated BTU
    expect(page.Locator("#btu-display")).ToContainText("BTU/year")
    
    // Select glazing
    page.SelectOption("select[name='glazing']", "double_poly")
    
    // Fill wall insulation
    page.Fill("input[name='wallNorth']", "R22")
    page.Fill("input[name='wallSouth']", "R22")
    page.Fill("input[name='wallEast']", "R22")
    page.Fill("input[name='wallWest']", "R22")
    
    // Fill frost skirt
    page.Fill("input[name='frostSkirtDepth']", "10")
    page.Fill("input[name='frostSkirtWidth']", "1")
    page.Fill("input[name='frostSkirtInsulation']", "R22")
    
    // Select air seal
    page.SelectOption("select[name='airSeal']", "good-3ach")
    
    page.Click("button:has-text('Next')")
    page.WaitForSelector("h2:has-text('Optimization')")
}

func TestWizard_Step4_OptimizationAndHeating(t *testing.T) {
    pw := playwright.New(t)
    defer pw.Close()

    page := pw.NewPage()
    login(page, "owner@test.com")
    navigateToWizardStep(page, 4)
    
    // Should show calculated values
    expect(page.Locator("#peak-heat")).ToContainText("kW")
    expect(page.Locator("#annual-energy")).ToContainText("kWh")
    
    // Add optimization
    page.Click("button:has-text('+ Add')")
    page.WaitForSelector("#optimization-list div")
    
    // Configure heating sources
    page.SelectOption("select[name='heatingSource1']", "electric")
    page.Fill("input[name='heatingPercent1']", "60")
    
    page.SelectOption("select[name='heatingSource2']", "wood")
    
    // Submit wizard
    page.Click("button:has-text('Create Project')")
    page.WaitForURL("**/farms/test-farm")
    
    // Should see new project in list
    expect(page.Locator("[data-testid='project-card']")).ToBeVisible()
}

func TestWizard_RequestAdditionalShape(t *testing.T) {
    pw := playwright.New(t)
    defer pw.Close()

    page := pw.NewPage()
    login(page, "owner@test.com")
    page.Goto(baseURL + "/farms/test-farm/projects/new")
    
    // Click request additional shape
    page.Click("text=Request Additional Shape")
    
    // Should open dialog
    page.WaitForSelector("[data-testid='request-shape-dialog']")
    
    // Fill and submit request
    page.Fill("textarea[name='description']", "Chinese-style passive solar greenhouse")
    page.Click("button:has-text('Submit Request')")
    
    // Should confirm
    page.WaitForSelector("text=Request submitted")
}

func TestWizard_BackNavigation(t *testing.T) {
    pw := playwright.New(t)
    defer pw.Close()

    page := pw.NewPage()
    login(page, "owner@test.com")
    navigateToWizardStep(page, 3)
    
    // Go back to step 2
    page.Click("button:has-text('Back')")
    page.WaitForSelector("h2:has-text('Dimensions')")
    
    // Values should persist
    expect(page.Locator("input[name='length']")).ToHaveValue("100")
}

func navigateToWizardStep(page playwright.Page, step int) {
    page.Goto(baseURL + "/farms/test-farm/projects/new")
    
    if step >= 2 {
        page.Click("[data-testid='shape-card']:first-child")
        page.Click("button:has-text('Next')")
        page.WaitForSelector("h2:has-text('Dimensions')")
    }
    if step >= 3 {
        page.Fill("input[name='length']", "100")
        page.Fill("input[name='width']", "10")
        page.Fill("input[name='height']", "4")
        page.Click("button:has-text('Next')")
        page.WaitForSelector("h2:has-text('Insulation')")
    }
    if step >= 4 {
        page.Fill("input[name='wallNorth']", "R22")
        page.Click("button:has-text('Next')")
        page.WaitForSelector("h2:has-text('Optimization')")
    }
}
```

## Add Project Button Tests

```go
func TestFarmDetail_AddProjectButton(t *testing.T) {
    pw := playwright.New(t)
    defer pw.Close()

    page := pw.NewPage()
    login(page, "owner@test.com")
    page.Goto(baseURL + "/farms/test-farm")
    
    // Should see Add Project button
    expect(page.Locator("[data-testid='add-project-button']")).ToBeVisible()
    
    // Click should navigate to wizard
    page.Click("[data-testid='add-project-button']")
    page.WaitForURL("**/farms/test-farm/projects/new")
}

func TestFarmDetail_BurgerMenu(t *testing.T) {
    pw := playwright.New(t)
    defer pw.Close()

    page := pw.NewPage()
    login(page, "owner@test.com")
    page.Goto(baseURL + "/farms/test-farm")
    
    // Open burger menu
    page.Click("[data-testid='farm-burger-menu']")
    
    // Should show options
    expect(page.Locator("text=Share")).ToBeVisible()
    expect(page.Locator("text=Transfer")).ToBeVisible()
    expect(page.Locator("text=Delete")).ToBeVisible()
}

func TestProjectCard_BurgerMenu(t *testing.T) {
    pw := playwright.New(t)
    defer pw.Close()

    page := pw.NewPage()
    login(page, "owner@test.com")
    page.Goto(baseURL + "/farms/test-farm")
    
    // Expand project card
    page.Click("[data-testid='project-card']:first-child")
    
    // Open project burger menu
    page.Click("[data-testid='project-burger-menu']:first-child")
    
    // Should show options
    expect(page.Locator("text=Share")).ToBeVisible()
    expect(page.Locator("text=Transfer")).ToBeVisible()
    expect(page.Locator("text=Delete")).ToBeVisible()
}
```
