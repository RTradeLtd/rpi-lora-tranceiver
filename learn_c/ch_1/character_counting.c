#include <stdio.h>
/* v1
int main() {
    long nc;
    nc = 0;
    while (getchar() != EOF) {
        ++nc; // increment by 1
        printf("%ld\n", nc);
    }
    printf("%ld\n", nc);
}
*/


int main() {
    double nc;
    for (nc = 0; getchar() != EOF; ++nc) ;
    printf("%0.f\n", nc);
}