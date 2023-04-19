package main

import "fmt"

func fillCoffeeMachine(
	waterLeft *int,
	regularMilkLeft *int,
	fancyMilkLeft *int,
	beansLeft *int,
	cupsLeft *int,
) {
	var (
		waterToAdd,
		regularMilkToAdd,
		fancyMilkToAdd,
		beanToAdd,
		cupsToAdd int
	)

	fmt.Print("Write how many ml of water you want to add:")
	fmt.Scan(&waterToAdd)

	fmt.Print("Write how many ml of regular milk you want to add:")
	fmt.Scan(&regularMilkToAdd)

	fmt.Print("Write how many ml of fancy milk you want to add:")
	fmt.Scan(&fancyMilkToAdd)

	fmt.Print("Write how many grams of coffee beans you want to add:")
	fmt.Scan(&beanToAdd)

	fmt.Print("Write how many disposable cups you want to add:")
	fmt.Scan(&cupsToAdd)

	*waterLeft += waterToAdd
	*regularMilkLeft += regularMilkToAdd
	*fancyMilkLeft += fancyMilkToAdd
	*beansLeft += beanToAdd
	*cupsLeft += cupsToAdd
}

func takeMoney(
	moneyLeft *int,
) {
	fmt.Printf("I gave you $%d", *moneyLeft)
	*moneyLeft = 0
}

func makeCoffee(
	waterPerCup int,
	milkPerCup int,
	beansPerCup int,
	price int,
	waterLeft *int,
	regularMilkLeft *int,
	fancyMilkLeft *int,
	beansLeft *int,
	moneyLeft *int,
	cupsLeft *int,
) {
	milkLeft := regularMilkLeft
	if milkPerCup > 0 {
		var milkType string
		fmt.Print("What type of milk? 1 - regular, " +
			"2 - fancy, back - return to main menu:")
		fmt.Scan(&milkType)
		switch milkType {
		case "1":
			milkLeft = regularMilkLeft
		case "2":
			milkLeft = fancyMilkLeft
		case "back":
			return
		}
	}

	if *waterLeft < waterPerCup {
		fmt.Print("Sorry, not enough water!")
	} else if *milkLeft < milkPerCup {
		fmt.Print("Sorry, not enough milk!")
	} else if *beansLeft < beansPerCup {
		fmt.Print("Sorry, not enough coffee beans!")
	} else if *cupsLeft < 1 {
		fmt.Print("Sorry, not enough disposable cups!")
	} else {
		fmt.Print("I have enough resources, making you a coffee!")
		*waterLeft -= waterPerCup
		*milkLeft -= milkPerCup
		*beansLeft -= beansPerCup
		*moneyLeft += price
		*cupsLeft -= 1
	}
}

func buyCoffee(
	coffeeType string,
	waterLeft *int,
	regularMilkLeft *int,
	fancyMilkLeft *int,
	beansLeft *int,
	moneyLeft *int,
	cupsLeft *int,
) {
	switch coffeeType {
	case "1":
		makeCoffee(
			250, 0, 16, 4,
			waterLeft, regularMilkLeft, fancyMilkLeft, beansLeft, moneyLeft, cupsLeft,
		)
	case "2":
		makeCoffee(
			350, 75, 20, 7,
			waterLeft, regularMilkLeft, fancyMilkLeft, beansLeft, moneyLeft, cupsLeft,
		)
	case "3":
		makeCoffee(
			200, 100, 12, 6,
			waterLeft, regularMilkLeft, fancyMilkLeft, beansLeft, moneyLeft, cupsLeft,
		)
	case "back":
		return
	}
}

func printState(
	waterLeft int,
	regullarMilkLeft int,
	fancyMilkLeft int,
	beansLeft int,
	moneyLeft int,
	cupsLeft int,
) {
	fmt.Println("The coffee machine has:")
	fmt.Printf("%d ml of water\n", waterLeft)
	fmt.Printf("%d ml of regular milk\n", regullarMilkLeft)
	fmt.Printf("%d ml of fancy milk\n", fancyMilkLeft)
	fmt.Printf("%d g of coffee beans\n", beansLeft)
	fmt.Printf("%d disposable cups\n", cupsLeft)
	fmt.Printf("$%d of money", moneyLeft)
}

func main() {
	var (
		waterLeft,
		regullarMilkLeft,
		fancyMilkLeft,
		beansLeft,
		cupsLeft,
		moneyLeft int = 400, 540, 540, 120, 9, 550
	)

	var action string

	for {
		fmt.Print("Write action (buy, fill, take, remaining, exit):")
		fmt.Scan(&action)
		switch action {
		case "fill":
			fillCoffeeMachine(&waterLeft, &regullarMilkLeft, &fancyMilkLeft, &beansLeft, &cupsLeft)
		case "buy":
			var coffeeType string
			fmt.Print("What do you want to buy? 1 - espresso, 2 - latte, " +
				"3 - cappuccino, back - return to main menu:")
			fmt.Scan(&coffeeType)
			buyCoffee(coffeeType, &waterLeft, &regullarMilkLeft, &fancyMilkLeft, &beansLeft, &moneyLeft, &cupsLeft)
		case "take":
			takeMoney(&moneyLeft)
		case "remaining":
			printState(waterLeft, regullarMilkLeft, fancyMilkLeft, beansLeft, moneyLeft, cupsLeft)
		case "exit":
			return
		}

		fmt.Print("\n\n")
	}
}
