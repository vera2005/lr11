package api

import (
	"errors"
	"fmt"
	"net/http"
	"strings"

	"github.com/golang-jwt/jwt/v4"
	"github.com/labstack/echo/v4"
)

var ErrEmailAlreadyTaken = errors.New("email уже занят")

//middleware - функция, которая при обработке запросов может выполнять доп действия до или после вызова обработчика

// проверяет наличие и валидность JWT в заголовке Authorization каждого запроса
// next - следдующий обработчик
// возвращает новую функцию, соответствующую типу echo.HandlerFunc, которая будет использоваться как middleware

func JWTMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	//контекст содержит информацию о текущем HTTP-запросе и позволяет взаимодействовать с ним
	return func(c echo.Context) error {
		token := c.Request().Header.Get("Authorization")
		if token == "" {
			return c.JSON(http.StatusUnauthorized, map[string]string{"error": "missing token"})
		}

		token = strings.TrimPrefix(token, "Bearer ") //удаление префикса
		claims := &Claims{}                          // Claims - это структура, которая содержит данные о пользователе
		tkn, err := jwt.ParseWithClaims(token, claims, func(token *jwt.Token) (interface{}, error) {
			return []byte("your_secret_key"), nil
		})

		if err != nil || !tkn.Valid {
			return c.JSON(http.StatusUnauthorized, map[string]string{"error": "invalid token"})
		}

		c.Set("userId", claims.UserId) // Сохраняем ID пользователя в контексте
		return next(c)
	}
}

// signUp - обработчик для регистрации пользователя
func (srv *Server) signUp(c echo.Context) error {
	var user User
	if err := c.Bind(&user); err != nil {
		return c.JSON(http.StatusBadRequest, Response{
			Message: "некорекктное считывание данных",
		})
	}
	fmt.Println("user signUp from api")
	// Вызываем бизнес-логику для регистрации пользователя
	token, err := srv.uc.SignUp(user)
	if err != nil {
		fmt.Println(err.Error(), errors.Is(err, ErrEmailAlreadyTaken))
		if errors.Is(err, ErrEmailAlreadyTaken) { // Проверяем конкретную ошибку
			fmt.Println("rrr")
			return c.JSON(http.StatusConflict, Response{
				Message: err.Error(),
			})
		}
		return c.JSON(http.StatusInternalServerError, Response{
			Message: "ошибка сервера",
		})
	}
	fmt.Println("signUp sucs")
	return c.JSON(http.StatusCreated, Response{
		Message: token,
	})
}

// Обработчик для аутентификации пользователя
func (srv *Server) signIn(c echo.Context) error {
	var credentials Credentials
	if err := c.Bind(&credentials); err != nil {
		fmt.Println("Invalid input")
		return c.JSON(http.StatusBadRequest, Response{
			Message: " Ошибка передачи параметров",
		})
	}
	token, err := srv.uc.SignIn(credentials)
	if err != nil {
		fmt.Println("Authentication failed")
		return c.JSON(http.StatusUnauthorized, Response{
			Message: " Ошибка сервера",
		})
	}

	return c.JSON(http.StatusOK, Response{
		Message: token,
	})
}
