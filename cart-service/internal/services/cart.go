package repository

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"log"

	"github.com/google/uuid"
	"github.com/wycliff-ochieng/cart-service/internal/database"
	"github.com/wycliff-ochieng/cart-service/internal/models"
)

var (
	ErrItemNotFound    = errors.New("item not found")
	ErrInvalidQuantity = errors.New("invlaid quantity")
	ErrCartNotFound    = errors.New("cart not found")
)

type CartService struct {
	l  *log.Logger
	db database.DBInterface
}

func NewCartService(l *log.Logger, db database.DBInterface) *CartService {
	return &CartService{l: l, db: db}
}

func (c *CartService) GetCart(ctx context.Context, userID uuid.UUID) (*models.Cart, error) {
	//TODO: implement database logic to get cart
	var cart models.Cart

	var cartItems []models.CartItem

	query := `SELECT FROM cart WHERE userID=$1`

	err := c.db.QueryRowContext(ctx, query).Scan(&cart.ID, &cart.UserID, &cart.SessionId, &cart.Status, &cart.Total, &cart.CreatedAt, &cart.UpdatedAt)
	if err != nil {
		return nil, ErrCartNotFound
	}

	itemsQuery := `SELECT id,cartID,productID,quantity,createdat FROM cart_items`

	rows, err := c.db.QueryContext(ctx, itemsQuery)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var Item models.CartItem

		if err := rows.Scan(&Item.CartID, &Item.ProductID, &Item.Price, &Item.Quantity, &Item.Subtotal); err != nil {
			return nil, err
		}
		cartItems = append(cartItems, Item)
	}

	return &models.Cart{}, nil
}

func (c *CartService) AddToCart(ctx context.Context, userID uuid.UUID, productID int32, quantity int) error {
	//TODO : implement database logic to add items

	if quantity <= 0 {
		return ErrInvalidQuantity
	}

	//look for cart for a user if not create new cart

	var cartID uuid.UUID

	err := c.db.QueryRowContext(ctx, "SELECT id FROM carts WHERE userID = $1", userID).Scan(&cartID)
	if err == sql.ErrNoRows {
		cartID = uuid.New()
		_, err = c.db.ExecContext(ctx, "INSERT INTO carts() VALUES()", cartID, userID)
		if err != nil {
			return fmt.Errorf("cannot return create cart ")
		}
	} else if err != nil {
		return fmt.Errorf("could not insert into:")
	}

	//check if product exist in the cart else add product

	var existsQ int

	err = c.db.QueryRowContext(ctx, "SELECT quantity FROM cart where cartID = $1 AND productID = $2", cartID, productID).Scan(&existsQ)
	if err == sql.ErrNoRows {
		_, err = c.db.ExecContext(ctx, "INSERT INTO cart_items VALUES()")
		if err != nil {
			return fmt.Errorf("could not insert item into table:%v", err)
		}
	} else if err == nil {
		newQ := existsQ + quantity
		_, err = c.db.ExecContext(ctx, "UPDATE cart_items SET quantity = $1  WHERE cartID = $2 AND productID = $3, ", newQ, cartID, productID)
		if err != nil {
			return fmt.Errorf("something Happened: %v", err)
		}
	} else {
		return fmt.Errorf("something unexoextced happened : %v", err)
	}

	return nil
}

func (c *CartService) UpdateCartContent(ctx context.Context, quantity int, productID int, cartID uuid.UUID) error {
	if quantity <= 0 {
		return ErrInvalidQuantity
	}
	return nil
	//TODO: implement update functionality
}
