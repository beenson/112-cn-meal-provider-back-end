package ordering

import (
	"context"
	"testing"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"

	"gitlab.winfra.cs.nycu.edu.tw/112-cn/meal-provider-back-end/pkg/ordering/model"
)

func setupService(t *testing.T) Service {
	db, err := gorm.Open(sqlite.Open("file::memory:"), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	err = db.AutoMigrate(&model.Food{}, &model.CartItem{}, &model.OrderItem{}, &model.Order{})
	if err != nil {
		panic(err)
	}

	return NewService(db)
}

func checkModelIsNotZero(t *testing.T, m *gorm.Model) {
	if m.CreatedAt.IsZero() {
		t.Fatal("m.CreatedAt is zero")
	}

	if m.UpdatedAt.IsZero() {
		t.Fatal("m.UpdatedAt is zero")
	}
}

func checkModelIsEqual(t *testing.T, m1, m2 *gorm.Model) {
	if m1.ID != m2.ID {
		t.Fatal("m1.ID != m2.ID")
	}

	if m1.CreatedAt != m2.CreatedAt {
		t.Fatal("m1.CreatedAt != m2.CreatedAt")
	}

	if m1.UpdatedAt != m2.UpdatedAt {
		t.Fatal("m1.UpdatedAt != m2.UpdatedAt")
	}

	if m1.DeletedAt != m2.DeletedAt {
		t.Fatal("m1.DeletedAt != m2.DeletedAt")
	}
}

func checkFoodIsEqual(t *testing.T, f1, f2 *model.Food) {
	checkModelIsEqual(t, &f1.Model, &f2.Model)

	if f1.Name != f2.Name {
		t.Fatal("f1.Name != f2.Name")
	}

	if f1.Description != f2.Description {
		t.Fatal("f1.Description != f2.Description")
	}

	if f1.Price != f2.Price {
		t.Fatal("f1.Price != f2.Price")
	}

	if f1.ImageURL != f2.ImageURL {
		t.Fatal("f1.ImageURL != f2.ImageURL")
	}
}

func createBurger(t *testing.T, service Service) *model.Food {
	food, err := service.CreateFood(context.Background(), CreateFoodInput{
		Name:        "Burger",
		Description: "Burger with cheese",
		Price:       10,
		ImageURL:    "",
	})

	if err != nil {
		t.Fatal(err)
	}

	return food
}

func createBread(t *testing.T, service Service) *model.Food {
	food, err := service.CreateFood(context.Background(), CreateFoodInput{
		Name:        "Bread",
		Description: "White bread",
		Price:       5,
		ImageURL:    "",
	})

	if err != nil {
		t.Fatal(err)
	}

	return food
}

func TestService_CreateFood(t *testing.T) {
	service := setupService(t)

	food, err := service.CreateFood(context.Background(), CreateFoodInput{
		Name:        "Burger",
		Description: "Burger with cheese",
		Price:       10,
		ImageURL:    "",
	})
	if err != nil {
		t.Fatal(err)
	}

	checkModelIsNotZero(t, &food.Model)

	if food.Name != "Burger" {
		t.Error("Food name should be Burger")
	}

	if food.Description != "Burger with cheese" {
		t.Error("Food description should be Burger")
	}

	if food.Price != 10 {
		t.Error("Food price should be 10")
	}

	if food.ImageURL != "" {
		t.Error("Food image url should be empty")
	}
}

func TestService_GetFood(t *testing.T) {
	service := setupService(t)

	food := createBurger(t, service)

	t.Run("should return food", func(t *testing.T) {
		f, err := service.GetFood(context.Background(), food.ID)
		if err != nil {
			t.Fatal(err)
		}

		if f == nil {
			t.Fatal("Food shouldn't be nil")
		}

		checkFoodIsEqual(t, f, food)
	})

	t.Run("should return nil if not found", func(t *testing.T) {
		f, err := service.GetFood(context.Background(), 2222)
		if err != nil {
			t.Fatal(err)
		}

		if f != nil {
			t.Fatal("Food should be nil")
		}
	})
}

func TestService_GetFoods(t *testing.T) {
	service := setupService(t)

	burger := createBurger(t, service)
	bread := createBread(t, service)

	t.Run("should return foods", func(t *testing.T) {
		foods, err := service.GetFoods(context.Background())
		if err != nil {
			t.Fatal(err)
		}

		if len(foods) != 2 {
			t.Fatal("Foods count should be 2, but ", len(foods))
		}

		checkFoodIsEqual(t, &foods[0], burger)
		checkFoodIsEqual(t, &foods[1], bread)
	})
}

func TestService_UpdateFood(t *testing.T) {
	service := setupService(t)

	burger := createBurger(t, service)

	foodName := "Banana"
	foodDesc := "Good banana"
	price := 60

	food, err := service.UpdateFood(context.Background(), burger.ID, UpdateFoodInput{
		Name:        &foodName,
		Description: &foodDesc,
		Price:       &price,
	})
	if err != nil {
		t.Fatal(err)
	}

	if food.UpdatedAt.Sub(burger.UpdatedAt) <= 0 {
		t.Error("UpdateAt should be greater than burger")
	}

	if food.Name != foodName {
		t.Error("Food name should be Banana")
	}

	if food.Description != foodDesc {
		t.Error("Food description should be Good banana")
	}

	if food.Price != price {
		t.Error("Food price should be 60")
	}

	if food.ImageURL != burger.ImageURL {
		t.Error("Food image url is not equal")
	}
}
