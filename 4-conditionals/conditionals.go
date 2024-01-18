package main

func main() {
	/*
		if:
			don't use parentheses around the condition
			some comparison operators:
				== equal to
				!= not equal to
				< less than
				> greather than
				<= less than or equal to
				>= greather than or equal to
	*/
	height := 4
	if height > 4 {
		println("Wow so tall...")
	} else if height < 4 {
		println("You're not tall")
	} else {
		println("You're mid.-")
	}

	/*
		short conditional variable syntax:
			used when for example you need to use a variable just for the conditional
				if INITIAL_STATEMENT; CONDITION {

				}
			variables declared inside INITIAL_STATEMENT will only work inside if block
	*/
	if length := 12; length < 10 {
		println("this never will run")
	}

	/*
		if's brace can't be put on the next line, because ; will be inserted after parentheses
			if i < f()  // wrong!
			{           // wrong!
				g()
			}
	*/
}
