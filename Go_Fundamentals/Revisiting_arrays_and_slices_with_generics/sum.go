package arrays

func SumAll(numbers ...[]int) []int {

	var sums []int

	for _, numbers := range numbers {
		sums = append(sums, GetSum(numbers))
	}

	return sums
}

func Reduce[A, B any](collection []A, f func(B, A) B, initialValue B) B {
	var result = initialValue
	for _, i := range collection {
		result = f(result, i)
	}
	return result
}

func GetSum(nums []int) int {
	sum := 0
	add := func(sum, x int) int {
		return sum + x
	}
	return Reduce(nums, add, sum)
}

func SumAllTails(numbers ...[]int) []int {

	Trail := func(sums, numbers []int) []int {
		if len(numbers) == 0 {
			sums = append(sums, 0)
		} else {
			sums = append(sums, GetSum(numbers[1:]))
		}
		return sums
	}

	return Reduce(numbers, Trail, []int{})
}

type Transaction struct {
	From string
	To   string
	Sum  float64
}

func NewTransaction(from, to Account, sum float64) Transaction {
	return Transaction{From: from.Name, To: to.Name, Sum: sum}
}

type Account struct {
	Name    string
	Balance float64
}

func NewBalanceFor(account Account, transactions []Transaction) Account {
	return Reduce(
		transactions,
		applyTransaction,
		account,
	)
}

func applyTransaction(a Account, transaction Transaction) Account {
	if transaction.From == a.Name {
		a.Balance -= transaction.Sum
	}
	if transaction.To == a.Name {
		a.Balance += transaction.Sum
	}
	return a
}

type Person struct {
	Name string
}

func Find[A any](items []A, predicate func(A) bool) (value A, found bool) {
	for _, v := range items {
		if predicate(v) {
			return v, true
		}
	}
	return
}
