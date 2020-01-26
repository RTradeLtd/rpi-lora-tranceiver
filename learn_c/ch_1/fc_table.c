#include <stdio.h>
#include <string.h>
// prints
/*
    Fahrenheit-Clesius table
*/

int main() {
    float fahr;
    float celsius;
    float lower,
    upper,
    step;

    lower = 0;
    upper = 200;
    step = 30;

    fahr = lower;
    
    while (fahr <= upper) {
	celsius = (5.0/9.0) * (fahr-32.0);
	printf("%3.0f %6.1f\n", fahr, celsius);
        fahr = fahr + step;
    }
}
