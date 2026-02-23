package repository

import (
	"cart-app/models"

	"gorm.io/gorm"
)

type CartRepository struct {
	db *gorm.DB
}

func NewCartRepository(db *gorm.DB) *CartRepository {
	return &CartRepository{db: db}
}

func (r *CartRepository) GetOrCreateActiveCart(userID string) (*models.Cart, error) {
	var cart models.Cart
	err := r.db.Where("user_id = ? AND status = ?", userID, "active").
		Preload("Items").
		First(&cart).Error

	if err == gorm.ErrRecordNotFound {
		// Create new cart
		cart = models.Cart{
			UserID: userID,
			Status: "active",
		}
		if err := r.db.Create(&cart).Error; err != nil {
			return nil, err
		}
		return &cart, nil
	}

	if err != nil {
		return nil, err
	}

	return &cart, nil
}

func (r *CartRepository) AddItem(cartID uint, item *models.CartItem) error {
	item.CartID = cartID
	return r.db.Create(item).Error
}

func (r *CartRepository) UpdateCartStatus(cartID uint, status string) error {
	return r.db.Model(&models.Cart{}).Where("id = ?", cartID).Update("status", status).Error
}

func (r *CartRepository) GetCartWithItems(cartID uint) (*models.Cart, error) {
	var cart models.Cart
	err := r.db.Preload("Items").First(&cart, cartID).Error
	return &cart, err
}

func (r *CartRepository) CalculateTotal(cartID uint) (float64, error) {
	var total float64
	err := r.db.Model(&models.CartItem{}).
		Where("cart_id = ?", cartID).
		Select("COALESCE(SUM(price * quantity), 0)").
		Scan(&total).Error
	return total, err
}
