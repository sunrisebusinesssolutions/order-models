# Order Models - Go

This package contains all order models from the [order-models](../../order-models) Java project, converted to Go structs with full field preservation for use with MongoDB and REST APIs.

## Project Structure

All models are designed for:
- JSON serialization/deserialization for REST APIs
- MongoDB BSON serialization/deserialization
- Exact field preservation from the original Java models

## Models Overview

### Core Domain Models

#### User
Represents a user in the Order Administration system. Includes authentication details from OAuth2 providers, roles, permissions, geolocation data from last login, and UI preferences.

**Key Fields:**
- `ID`, `Email`, `FirstName`, `LastName`, `Picture`
- `Role`, `Permissions`
- `LastLoginIP`, `LastLoginCity`, `LastLoginCountry`, `LastLoginLatitude`, `LastLoginLongitude`, `LastLoginTimestamp`, `LastLoginDeviceType`, `LastLoginUserAgent`
- `CreatedAt`, `UpdatedAt`, `Disabled`
- `ContactEmail`, `ContactPhone`, `AllowNotifications`, `Theme`

**Methods:**
- `GetFullName()` - Returns "FirstName LastName"
- `GetLastLoginLocation()` - Returns formatted location string
- `HasGeolocationData()` - Checks if geolocation data is available

#### Order
Represents a customer order with fulfillment details, customer contact information, and ordered items.

**Key Fields:**
- `ID`, `CreatedBy`, `CustomerID`, `Name`, `Email`, `Phone`
- `OrderDate`, `Type` (PICKUP/DELIVERY), `DeliverDate`
- `Approved`, `ApprovedDate`, `LocationID`, `Status`
- `CalendarEventID`, `Items`, `Notes`
- `CreatedAt`, `UpdatedAt`

#### ScheduledOrder
Extends Order to support recurring orders (weekly or monthly recurrence patterns).

**Additional Fields:**
- `WeekDay` - Day of week for weekly recurrence
- `RecurrenceType` - WEEKLY or MONTHLY
- `MonthlyPattern` - Pattern for monthly recurrence
- `MonthlyDayOfWeek` - Day of week for monthly recurrence

#### Item
Represents a line item in an order with product and customization details.

**Key Fields:**
- `ID`, `OrderID`, `ItemID`, `ItemType`, `Qty`
- `ProductName`, `ProductImageURL`, `ProductUnit`, `UnitPrice`, `LineTotal`
- `FrostingID`, `FrostingName`, `IcingID`, `IcingName`, `ToppingID`, `ToppingName`
- `Notes`, `Mix`
- `Brand`, `Size`, `Color`, `Material` (for merchandise items)

#### Customer
Stores personal and contact information for customers.

**Key Fields:**
- `ID`, `FirstName`, `MiddleName`, `LastName`
- `Address1`, `Address2`, `City`, `State`, `Zip`
- `Phone`, `Email`

### Product Models

#### Product
Base class for all product items in the catalog.

**Key Fields:**
- `ID`, `Order`, `Description`, `Price`, `SpecialPrice`, `SpecialPriceDate`
- `ItemType`, `Unit`, `URL`, `AvailableDays`
- `ImageOriginal`, `ImageMedium`, `ImageSmall`, `ImageIcon`, `ImageThumbnail`

**Methods:**
- `GetImageOriginal()`, `SetImageOriginal()`
- `GetImageMedium()`, `SetImageMedium()`
- `GetImageSmall()`, `SetImageSmall()`
- `GetImageIcon()`, `SetImageIcon()`
- `GetImageThumbnail()`, `SetImageThumbnail()`

#### Donut
Represents a donut product with customization options (frosting, icing, topping).

**Additional Fields:**
- `Frosting`, `RequiredFrosting`
- `Icing`, `RequiredIcing`
- `Topping`, `RequiredTopping`
- `AllowMix`

**Methods:** IsPastry interface methods for customization control

#### Roll
Represents a roll product (e.g., cinnamon roll) with customization options.

**Fields & Methods:** Same as Donut

#### Merchandise
Represents non-bakery merchandise items (cups, shirts, etc.).

**Additional Fields:**
- `Size`, `Color`, `Material`, `Brand`
- `HasSizeOptions`, `HasColorOptions`

### Customization Models

#### Frosting
Represents a frosting customization option with image URLs for multiple resolutions.

**Key Fields:**
- `ID`, `Description`, `URL`
- `ImageOriginal`, `ImageMedium`, `ImageSmall`, `ImageIcon`

#### Icing
Represents an icing customization option.

**Key Fields:**
- `ID`, `Description`, `URL`

#### Topping
Represents a topping customization option.

**Key Fields:**
- `ID`, `Description`, `URL`

### Configuration Models

#### Location
Represents a physical store or pickup location.

**Key Fields:**
- `ID`, `Name`
- `Address1`, `Address2`, `City`, `State`, `Zip`
- `Phone`, `Email`
- `PrimaryLocation`, `Visible`
- `Latitude`, `Longitude`

**Methods:**
- `HasGeolocation()` - Checks if coordinates are set

#### Status
Configuration for order placement status with time windows and customer messages.

**Key Fields:**
- `ID`, `Status` (OrderPlacementStatus)
- `StartTime`, `EndTime`
- `PendingMessage`, `PendingSubject`
- `ApprovedMessage`, `ApprovedSubject`
- `DeclinedMessage`, `DeclinedSubject`
- `CompletedMessage`, `CompletedSubject`
- `DeniedMessage`, `DeniedSubject`
- `DeletedMessage`, `DeletedSubject`
- `StaffMessage`

#### PricingSheet
Represents pricing configuration entry for a product at a location.

**Key Fields:**
- `ID`, `LocationID`, `Order`
- `ProductType`, `ProductCode`, `Description`, `Unit`, `Price`

#### ItemTypeLimit
Represents capacity limit for a product type.

**Key Fields:**
- `ID`, `ItemType`, `Capacity`, `SheetSize`

### Content Models

#### Announcement
Represents a general announcement or event for the calendar view.

**Key Fields:**
- `ID`, `Title`, `Description`
- `StartDate`, `EndDate`, `Color`
- `Active`, `CreatedAt`, `CreatedBy`

**Methods:**
- `GetEffectiveEndDate()` - Returns end date or start date if end is nil
- `IsMultiDay()` - Checks if announcement spans multiple days

#### ContactMessage
Represents a contact message submitted via the "Contact Us" form.

**Key Fields:**
- `ID`, `Name`, `Email`, `Subject`, `Message`, `CreatedAt`

#### Sweepstakes
Represents a sweepstakes contest with participant tracking.

**Key Fields:**
- `ID`, `Title`, `Description`
- `BeginDate`, `EndDate`
- `ImageURL`, `WinnerEmail`, `WinnerDate`, `TicketNumber`
- `EnteredUsers`, `CalendarEventID`
- `CreatedAt`, `UpdatedAt`

## Enumerations

### Role
- `USER` - Standard user role
- `STAFF` - Staff/operator role
- `ADMIN` - Administrator role

### ItemType
Product classifications: GLAZED_DONUT, CAKE_DONUT, RAISED_DONUT, ROLL, MUFFIN, PASTRY, COOKIE, CROISSANT, BAGEL, BEVERAGE, CUP, SHIRT, GLUTEN_FREE, VEGAN, HALAL, MISC, MIX, and more

### OrderStatus
- `PENDING` - Awaiting staff review
- `APPROVED` - Ready for production
- `COMPLETED` - Fulfilled and delivered/picked up
- `DENIED` - Denied by staff
- `DECLINED` - Failed to process
- `DELETED` - Logically deleted

### OrderType
- `PICKUP` - Customer pickup
- `DELIVERY` - Customer delivery
- `STANDARD` - Standard order

### OrderPlacementStatus
- `OFF` - Ordering disabled
- `SCHEDULED` - Ordering scheduled but not open
- `OPEN` - Ordering open
- `CLOSED` - Ordering window closed

### ProductType
- `DONUT`, `DONUT_HOLE`, `ROLL`, `MIX`, `BIG_DONUT`, `BEVERAGE`, `OTHER`

### RecurrenceType
- `WEEKLY` - Weekly recurrence
- `MONTHLY` - Monthly recurrence

### MonthlyDayPattern
- `FIRST_OCCURRENCE` - First occurrence of day in month
- `SECOND_OCCURRENCE`
- `THIRD_OCCURRENCE`
- `FOURTH_OCCURRENCE`
- `LAST_OCCURRENCE` - Last occurrence of day in month
- `LAST_BUSINESS_DAY` - Last business day (Mon-Fri) of month

### Permission
Fine-grained permissions for access control beyond roles:
- Product management: `MANAGE_PRODUCTS`
- Administrative: `MANAGE_USERS`, `MANAGE_MESSAGES`, `MANAGE_SWEEPSTAKES`, `MANAGE_LOCATIONS`, etc.
- Staff features: `MANAGE_DAILY_SPECIAL`, `VIEW_PRODUCTION`, `VIEW_BOARD`
- Statistics: `VIEW_STATS`, `VIEW_DASHBOARD`, `VIEW_REPORTS`, `VIEW_ANNOUNCEMENTS`

## Interfaces

### HasImageUrls
Implemented by Product and its subclasses to provide multiple image size variants.

### IsPastry
Implemented by Donut and Roll to provide customization option methods.

## JSON/BSON Tags

All structs use both `json` and `bson` tags for compatibility with:
- REST API JSON serialization
- MongoDB BSON storage with automatic `_id` field mapping for `ID` fields

## Time Fields

Time fields use Go's `time.Time` type with pointer variants (`*time.Time`) for optional fields that may be null in the database.

## Numeric Fields

- `BigDecimal` (Java) → `string` (Go) for pricing fields to preserve precision
- `Double` (Java) → `*float64` (Go pointer for optional values)
- `Integer` (Java) → `int` or `*int` (Go pointer for optional values)

## Inheritance Structure

Go uses struct composition instead of inheritance:
- `Donut` embeds `Product` and adds pastry-specific fields
- `Roll` embeds `Product` and adds pastry-specific fields
- `Merchandise` embeds `Product` and adds merchandise-specific fields
- `ScheduledOrder` embeds `Order` and adds recurrence-specific fields
