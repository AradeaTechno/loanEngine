package middleware

import (
	"amarthaloan/config"
	"amarthaloan/helpers"
	"errors"
	"fmt"
	"net"
	"net/http"
	"strings"
	"time"

	"github.com/didip/tollbooth"
	"github.com/didip/tollbooth/limiter"
	"github.com/golang-jwt/jwt/v4"
	"github.com/labstack/echo/v4"
)

var appConf = config.AppConfig()

// var secretConfg = config.SecretConfig()
var lmt = tollbooth.NewLimiter(5, &limiter.ExpirableOptions{DefaultExpirationTTL: time.Minute})

func IsAuthenticated(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		tokenString := c.Request().Header.Get("Authorization")
		if tokenString == "" {
			return echo.NewHTTPError(http.StatusUnauthorized, helpers.ApiResponse(errors.New("missing token"), nil))
		}

		// Strip "Bearer " if present
		tokenString = strings.TrimPrefix(tokenString, "Bearer ")

		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			// Ensure the signing method is what we expect
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
			}
			return []byte(appConf.SIGN_SECRET), nil
		})

		if err != nil || !token.Valid {
			return echo.NewHTTPError(http.StatusUnauthorized, helpers.ApiResponse(errors.New("invalid token"), nil))
		}

		// Pass the token to the context if needed
		c.Set("user", token)

		return next(c)
	}
}

// func VerifyUser() echo.MiddlewareFunc {
// 	return func(next echo.HandlerFunc) echo.HandlerFunc {
// 		return func(c echo.Context) error {
// 			// Get the JWT token and claims

// 			user := c.Get("user").(*jwt.Token)
// 			claims := user.Claims.(jwt.MapClaims)

// 			var is_active, is_suspended, is_banned string

// 			if val, ok := claims["is_active"].(string); ok {
// 				is_active = val
// 			} else {
// 				is_active = ""
// 			}

// 			if val, ok := claims["is_suspended"].(string); ok {
// 				is_suspended = val
// 			} else {
// 				is_suspended = ""
// 			}

// 			if val, ok := claims["is_banned"].(string); ok {
// 				is_banned = val
// 			} else {
// 				is_banned = ""
// 			}

// 			// Fetch the user from the database using seed
// 			userId := c.Request().Header.Get("paramId")
// 			seedHeader := c.Request().Header.Get("seed")
// 			ipHeader := c.Request().Header.Get("userIp")
// 			// Get Data user
// 			obj, err := models.GetUserById(userId)
// 			if err != nil {
// 				// return echo.NewHTTPError(http.StatusUnauthorized, "User not found")
// 				res.Status = http.StatusUnauthorized
// 				res.Message = "Error!"
// 				res.Data = map[string]string{
// 					"error": "Invalid Credentials",
// 				}
// 				return echo.NewHTTPError(http.StatusUnauthorized, res)
// 			}
// 			// RETRIEVE DATA USER
// 			userMap, ok := obj.Data.(models.UserStruct)
// 			if !ok {
// 				// return echo.NewHTTPError(http.StatusUnauthorized, claims)
// 				res.Status = http.StatusUnauthorized
// 				res.Message = "Error!"
// 				res.Data = map[string]string{
// 					"error": "Invalid Credentials",
// 				}
// 				return echo.NewHTTPError(http.StatusUnauthorized, res)
// 			}

// 			// CREATE HASH FROM RETURN DATA
// 			hash := helpers.IdentifierHash(userMap.User_id, userMap.Username, userMap.Phone, userMap.Email, userMap.Seed, int(userMap.Modified))
// 			// ====================== CHECK ALL GIVEN TO DATABASE =========================== //
// 			if userMap.Is_active != is_active || userMap.Is_suspended != is_suspended || userMap.Is_banned != is_banned || hash != seedHeader || userMap.Last_ip_address != ipHeader {
// 				res.Status = http.StatusUnauthorized
// 				res.Message = "Error!"
// 				res.Data = map[string]string{
// 					"error": "Invalid Credentials",
// 				}
// 				return echo.NewHTTPError(http.StatusUnauthorized, res)
// 			}

// 			// Continue to the next handler if everything matches
// 			return next(c)
// 		}
// 	}
// }

func CurlAndWhitelist() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			// Step 1: Get the client IP
			clientIP := c.Request().Header.Get("userIp")

			// Step 2: Check if the client IP is in the whitelist
			isWhitelisted := false
			for _, allowedIP := range config.AppConfig().ALLOWED_IP {
				if clientIP == allowedIP {
					isWhitelisted = true
					break
				}
			}

			// Step 3: Check if the request is from curl
			userAgent := c.Request().Header.Get("User-Agent")
			isCurlRequest := strings.HasPrefix(userAgent, "curl/")

			// Step 4: Logic to allow access
			if isWhitelisted && isCurlRequest {
				// Allow access if the IP is whitelisted or the request is from cURL
				return next(c)
			}
			return c.JSON(http.StatusForbidden, helpers.ApiResponse(errors.New("forbidden access"), nil))
		}
	}
}

func RateLimitMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		// Retrieve the user IP from the request header
		userIP := c.Request().Header.Get("userIp")
		if userIP == "" {
			return echo.NewHTTPError(http.StatusBadRequest, helpers.ApiResponse(errors.New("missing user ip"), nil))
		}

		// Validate the IP format
		if net.ParseIP(userIP) == nil {
			return echo.NewHTTPError(http.StatusBadRequest, helpers.ApiResponse(errors.New("invalid ip format"), nil))
		}

		// Use LimitByKeys with user IP as the key
		httpError := tollbooth.LimitByKeys(lmt, []string{userIP})
		if httpError != nil {
			return c.JSON(http.StatusTooManyRequests, helpers.ApiResponse(errors.New("rate limit exceed"), nil))
		}

		// Continue to the next middleware or handler if rate limit not exceeded
		return next(c)
	}
}
