package helper

// REQUIRED_PARAMS --------------- DEVELOPER MESSAGES -----------------//
var REQUIRED_PARAMS = "Please provide a required params."
var FETCHED_FAILED = "Failed to fetch the data from server."
var FETCHED_SUCCESS = "Data fetched successfully."

// DATA_INSERTED -------------- END USERS MESSAGES ----------------//
var DATA_INSERTED = "Operation successful."
var DATA_INSERTED_FAILED = "Operation failed, please try again."
var UPDATE_FAILED = "Update failed, please try again"
var UPDATE_SUCCESS = "Update successful."
var LOGIN_SUCCESS = "Login successful."
var DATA_NOT_FOUND = "Data not found."
var DATA_FOUND = "Data found."

// FAILED_PROCESS -------------- COMMON MESSAGES ----------------//
var FAILED_PROCESS = "failed to process the request."
var PAGINATION_INVALID = "Pagination failed, please try again ..."
var DELETE_FAILED = "failed to delete"
var DELETE_SUCCESS = "successfully deleted"
var INVALID_ID = "invalid id pass"

// REQUEST_HOST -------------- BACKEND DEV VARIABLES -----------//S
var REQUEST_HOST = ""
var CURRENT_IDX = 0
var PREVIOUS_IDX = 0
var TOTALCOUNT int64 = 0
var TOKEN_ID = ""
var USER_TYPE = ""

// DATA --------------- Response Key Words ------------//
var DATA = "data"
var USER_DATA = "user_data"
var ORDER_DATA = "order_data"

// order Status tags

var PLACED = "PLACED"
var DISPATCHED = "DISPATCHED"
var COMPLETED = "COMPLETED"
var CANCELLED = "CANCELLED"
var CART = "CART"
