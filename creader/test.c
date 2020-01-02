#include <stdio.h>
#include <stdlib.h>
#include "rnibreader.h"

int main() {
    printf("Using rnibreader lib from C:\n");

    open();
    void *result = getListGnbIds();

    if(result == NULL){

        printf("ERROR: no data from getListGnbIds\n");
        return 1;
    }

    printf("getListGnbIds response: %s\n", (char *)result);

    close();

    free(result);

    return 0;
}