package models

type Permission struct {
	Value       string
	DisplayName string
	Description string
	AllowedRoles []Role
}

var (
	PermissionManageProducts = Permission{
		Value:        "MANAGE_PRODUCTS",
		DisplayName:  "Manage Products (ADMIN)",
		Description:  "Edit donuts, rolls, frostings, icings, and toppings",
		AllowedRoles: []Role{RoleAdmin},
	}
	PermissionManageUsers = Permission{
		Value:        "MANAGE_USERS",
		DisplayName:  "Manage Users (ADMIN)",
		Description:  "Access user management and edit user accounts",
		AllowedRoles: []Role{RoleAdmin},
	}
	PermissionManageMessages = Permission{
		Value:        "MANAGE_MESSAGES",
		DisplayName:  "Manage Messages (ADMIN)",
		Description:  "Access and manage contact inquiries",
		AllowedRoles: []Role{RoleAdmin},
	}
	PermissionManageSweepstakes = Permission{
		Value:        "MANAGE_SWEEPSTAKES",
		DisplayName:  "Manage Sweepstakes (ADMIN)",
		Description:  "Create, edit, and manage promotional sweepstakes contests",
		AllowedRoles: []Role{RoleAdmin},
	}
	PermissionManageLocations = Permission{
		Value:        "MANAGE_LOCATIONS",
		DisplayName:  "Manage Locations (ADMIN)",
		Description:  "Add, edit, and remove store locations",
		AllowedRoles: []Role{RoleAdmin},
	}
	PermissionManageProductOrdering = Permission{
		Value:        "MANAGE_PRODUCT_ORDERING",
		DisplayName:  "Manage Product Display Order (ADMIN)",
		Description:  "Control display order of menu products",
		AllowedRoles: []Role{RoleAdmin},
	}
	PermissionManagePricing = Permission{
		Value:        "MANAGE_PRICING",
		DisplayName:  "Manage Pricing (ADMIN)",
		Description:  "View and edit the pricing sheet",
		AllowedRoles: []Role{RoleAdmin},
	}
	PermissionManageSchedule = Permission{
		Value:        "MANAGE_SCHEDULE",
		DisplayName:  "Manage Schedule (ADMIN)",
		Description:  "Access and modify scheduling configuration",
		AllowedRoles: []Role{RoleAdmin},
	}
	PermissionManageSystemMessages = Permission{
		Value:        "MANAGE_SYSTEM_MESSAGES",
		DisplayName:  "Manage System Messages (ADMIN)",
		Description:  "Access and configure system messages for order statuses",
		AllowedRoles: []Role{RoleAdmin},
	}
	PermissionManageLimits = Permission{
		Value:        "MANAGE_LIMITS",
		DisplayName:  "Manage Limits (STAFF, ADMIN)",
		Description:  "Configure item-type production capacity limits",
		AllowedRoles: []Role{RoleStaff, RoleAdmin},
	}
	PermissionReviewOrders = Permission{
		Value:        "REVIEW_ORDERS",
		DisplayName:  "Review Orders (STAFF, ADMIN)",
		Description:  "Access all orders review and management",
		AllowedRoles: []Role{RoleStaff, RoleAdmin},
	}
	PermissionManageDailySpecial = Permission{
		Value:        "MANAGE_DAILY_SPECIAL",
		DisplayName:  "Manage Daily Special (STAFF, ADMIN)",
		Description:  "Set and manage daily specials",
		AllowedRoles: []Role{RoleStaff, RoleAdmin},
	}
	PermissionViewProduction = Permission{
		Value:        "VIEW_PRODUCTION",
		DisplayName:  "View Production Day (STAFF, ADMIN)",
		Description:  "Access production day planning",
		AllowedRoles: []Role{RoleStaff, RoleAdmin},
	}
	PermissionViewBoard = Permission{
		Value:        "VIEW_BOARD",
		DisplayName:  "View Pickup/Deliver (STAFF, ADMIN)",
		Description:  "Access the pickup/delivery order board",
		AllowedRoles: []Role{RoleStaff, RoleAdmin},
	}
	PermissionViewStats = Permission{
		Value:        "VIEW_STATS",
		DisplayName:  "View Statistics (STAFF, ADMIN)",
		Description:  "Access application statistics and metrics",
		AllowedRoles: []Role{RoleStaff, RoleAdmin},
	}
	PermissionViewDashboard = Permission{
		Value:        "VIEW_DASHBOARD",
		DisplayName:  "View Dashboard (ADMIN)",
		Description:  "Access administrative dashboard with KPI charts",
		AllowedRoles: []Role{RoleAdmin},
	}
	PermissionViewReports = Permission{
		Value:        "VIEW_REPORTS",
		DisplayName:  "View Reports (ADMIN)",
		Description:  "Access system reports and data exports",
		AllowedRoles: []Role{RoleAdmin},
	}
	PermissionViewAnnouncements = Permission{
		Value:        "VIEW_ANNOUNCEMENTS",
		DisplayName:  "View Announcements (STAFF, ADMIN)",
		Description:  "Access announcements and events calendar",
		AllowedRoles: []Role{RoleStaff, RoleAdmin},
	}
)
