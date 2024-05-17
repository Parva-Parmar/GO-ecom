package controllers

import(

)

var (

	ErrCantFindProduct = errors.New("can't find the product")
	ErrCantDecodeProducts = errors.New("can't find the product")
	ErrUserIdIdNotValid = errors.New("this user is not valid")
	ErrCantUpdateUser = errors.New("cannot add this product to the cart")
	ErrCantReomveItemCart = errors.New("cannot remove this item from the cart")
	ErrCantGetItem = errors.New("was unable to get tge item from the cart")
	ErrCantBuyCartItem = errors.New("cannot update the purchase")

)

func AddProductToCart(){

}

func RemoveCartItem(){

}

func BuyItemFromCart(){

}

func InstantBuyer(){

}