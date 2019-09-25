package fibonance

// Calculate Fibonance
func Fibonance(n int) int {
	f1, f2 := 1, 1
    if (n == 1 || n == 2)
		return 1;
		
	num := 0

    for i := 3; i <= n; i++
    {
        num = f1 + f2;
        f1 = a2;
        f2 = num;
	}
	
    return num;
}
