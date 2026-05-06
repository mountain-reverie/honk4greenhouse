# API Routes

## Public (No Auth)

| Method | Path | Handler | Description |
|--------|------|---------|-------------|
| GET | `/` | `showcase` | Landing page with public greenhouses |
| GET | `/showcase` | `showcase` | Alias for landing |
| GET | `/greenhouse/{id}` | `greenhouseDetail` | Public greenhouse detail |
| GET | `/login` | `loginPage` | Login form |
| POST | `/auth/oauth/{provider}` | `oauthStart` | Start OAuth flow |
| GET | `/auth/callback/{provider}` | `oauthCallback` | OAuth callback |
| GET | `/publication/approve` | `publicationApprove` | Approve via email token |
| GET | `/publication/deny` | `publicationDeny` | Deny via email token |

## Authenticated

| Method | Path | Handler | Description |
|--------|------|---------|-------------|
| GET | `/shared` | `shared` | Greenhouses shared with user |
| GET | `/personal` | `personal` | User's own farms |
| POST | `/logout` | `logout` | Logout |

## Farms

| Method | Path | Handler | Description |
|--------|------|---------|-------------|
| GET | `/farms/new` | `farmForm` | Add Farm form page |
| GET | `/farms/{id}` | `farmDetail` | Farm detail page |
| POST | `/farms` | `farmCreate` | Create farm |
| PUT | `/farms/{id}` | `farmUpdate` | Update farm |
| DELETE | `/farms/{id}` | `farmDelete` | Delete farm |

## Farm Co-Owners

| Method | Path | Handler | Description |
|--------|------|---------|-------------|
| GET | `/api/farms/{id}/co-owners` | `coOwnersGet` | List co-owners |
| POST | `/api/farms/{id}/co-owners` | `coOwnerAdd` | Add co-owner by email |
| DELETE | `/api/farms/{id}/co-owners/{userId}` | `coOwnerRemove` | Remove co-owner |

## Energy Costs

| Method | Path | Handler | Description |
|--------|------|---------|-------------|
| GET | `/api/farms/{id}/energy-costs` | `energyCostsGet` | List energy costs |
| POST | `/api/farms/{id}/energy-costs` | `energyCostAdd` | Add energy cost |
| PUT | `/api/farms/{id}/energy-costs/{costId}` | `energyCostUpdate` | Update energy cost |
| DELETE | `/api/farms/{id}/energy-costs/{costId}` | `energyCostDelete` | Delete energy cost |

## Climate Data

| Method | Path | Handler | Description |
|--------|------|---------|-------------|
| GET | `/api/farms/{id}/climate` | `climateGet` | Get current climate data |
| GET | `/api/farms/{id}/climate?view=chart` | `climateChart` | Chart partial |
| GET | `/api/farms/{id}/climate?view=table` | `climateTable` | Table partial |
| GET | `/api/farms/{id}/climate/versions` | `climateVersions` | Version list |
| GET | `/api/farms/{id}/climate/{version}` | `climateVersion` | Specific version |
| POST | `/api/farms/{id}/climate/toggle-lock` | `climateLock` | Toggle edit lock |
| POST | `/api/farms/{id}/climate/update` | `climateUpdate` | Update cell value |
| POST | `/api/farms/{id}/climate/save` | `climateSave` | Save as new version |
| POST | `/api/farms/{id}/climate/cancel` | `climateCancel` | Cancel edits |

## Projects

| Method | Path | Handler | Description |
|--------|------|---------|-------------|
| GET | `/projects/{id}` | `projectDetail` | Project detail page |
| POST | `/farms/{id}/projects` | `projectCreate` | Create project (final submit) |
| PUT | `/projects/{id}` | `projectUpdate` | Update project |
| DELETE | `/projects/{id}` | `projectDelete` | Delete project |

## Project Creation Wizard

| Method | Path | Handler | Description |
|--------|------|---------|-------------|
| GET | `/farms/{id}/projects/new` | `projectWizardStart` | Start wizard (step 1) |
| GET | `/farms/{id}/projects/new/step/{n}` | `projectWizardStep` | Get wizard step n |
| POST | `/api/projects/wizard/shapes` | `getGreenhouseShapes` | List available shapes |
| POST | `/api/projects/wizard/request-shape` | `requestNewShape` | Request custom shape |
| POST | `/api/projects/wizard/calculate-btu` | `calculateBTU` | Calculate BTU requirements |
| POST | `/api/projects/wizard/calculate-heating` | `calculateHeating` | Calculate heating needs |
| GET | `/api/crops/{type}/profile` | `getCropProfile` | Get crop requirements |
| GET | `/api/optimizations` | `listOptimizations` | List optimization options |
| POST | `/api/optimizations/{type}/configure` | `configureOptimization` | Configure optimization |

## Photos

| Method | Path | Handler | Description |
|--------|------|---------|-------------|
| POST | `/api/projects/{id}/photos` | `photoUpload` | Upload photo |
| DELETE | `/api/projects/{id}/photos/{photoId}` | `photoDelete` | Delete photo |

## Quotes

| Method | Path | Handler | Description |
|--------|------|---------|-------------|
| GET | `/api/projects/{id}/quotes` | `quotesGet` | Get quotes |
| POST | `/api/projects/{id}/quotes` | `quoteCreate` | Submit quote |
| PUT | `/api/quotes/{id}` | `quoteUpdate` | Update quote |
| DELETE | `/api/quotes/{id}` | `quoteDelete` | Delete quote |
| POST | `/api/quotes/{id}/review` | `quoteReview` | Mark as reviewed |

## Sharing

| Method | Path | Handler | Description |
|--------|------|---------|-------------|
| GET | `/api/projects/{id}/shares` | `sharesGet` | Get shares |
| POST | `/api/projects/{id}/share` | `shareCreate` | Share project |
| DELETE | `/api/shares/{id}` | `shareRevoke` | Revoke share |
| POST | `/api/projects/{id}/transfer` | `projectTransfer` | Transfer ownership |

## Publication

| Method | Path | Handler | Description |
|--------|------|---------|-------------|
| GET | `/api/projects/{id}/publication` | `publicationStatus` | Get status |
| POST | `/api/projects/{id}/publication/request` | `publicationRequest` | Request publication |
| POST | `/api/projects/{id}/publication/retrigger` | `publicationRetrigger` | Retrigger requests |
| POST | `/api/publication/approve` | `publicationApproveAPI` | Approve (with token) |
| POST | `/api/publication/deny` | `publicationDenyAPI` | Deny (with token) |

## Admin

| Method | Path | Handler | Description |
|--------|------|---------|-------------|
| POST | `/admin/greenhouse/{id}/disable` | `adminDisable` | Disable showcase visibility |
| POST | `/admin/greenhouse/{id}/enable` | `adminEnable` | Re-enable visibility |
